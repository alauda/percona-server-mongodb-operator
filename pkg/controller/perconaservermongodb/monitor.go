package perconaservermongodb

import (
	"context"
	promClient "github.com/coreos/prometheus-operator/pkg/client/versioned"
	"github.com/percona/percona-server-mongodb-operator/pkg/psmdb/monitor"
	"k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/rest"
	"os"
)

var WatchNamespace = os.Getenv("WATCH_NAMESPACE")

func init() {
	if WatchNamespace == "" {
		WatchNamespace = "default"
	}
}

func (r *ReconcilePerconaServerMongoDB) reconcileMonitor(enabled bool) error {
	err := r.reconcileServiceMonitors(enabled)
	if err != nil {
		return err
	}
	return r.reconcileGrafanaDashboards(enabled)
}

func (r *ReconcilePerconaServerMongoDB) reconcileServiceMonitors(enabled bool) error {
	sm := monitor.GenerateServiceMonitor()
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Error(err, "error loading kubernetes configuration inside cluster")
		return err
	}
	p, err := promClient.NewForConfig(config)
	if err != nil {
		log.Error(err, "error create prometheus client")
		return err
	}
	_, err = p.MonitoringV1().ServiceMonitors(WatchNamespace).Get(sm.Name, v1.GetOptions{})
	if err != nil && !errors.IsNotFound(err) {
		log.Error(err, "get servicemonitor fail")
		return err
	}
	if enabled {
		if errors.IsNotFound(err) {
			_, err = p.MonitoringV1().ServiceMonitors(WatchNamespace).Create(sm)
			if err != nil {
				log.Error(err, "create servicemonitor fail")
			}
		}
	} else {
		if err == nil {
			err = p.MonitoringV1().ServiceMonitors(WatchNamespace).Delete(sm.Name, &v1.DeleteOptions{})
			if err != nil {
				log.Error(err, "delete servicemonitor fail")
			}
		}
	}
	return err
}

func (r *ReconcilePerconaServerMongoDB) reconcileGrafanaDashboards(enabled bool) error {
	gd := monitor.GenerateGrafana(WatchNamespace)
	err := r.client.Get(context.TODO(),
		types.NamespacedName{Name: gd.Name,
			Namespace: gd.Namespace},
		gd)
	if err != nil && !errors.IsNotFound(err) {
		log.Error(err, "Get grafana dashboards fail")
		return err
	}
	if enabled {
		if errors.IsNotFound(err) {
			err = r.client.Create(context.TODO(), gd)
			if err != nil {
				log.Error(err, "Cannot create grafana dashboard")
			}
		}
	} else {
		if err == nil {
			err = r.client.Delete(context.TODO(), gd)
			if err != nil {
				log.Error(err, "Cannot delete grafana dashboard")
			}
		} else if errors.IsNotFound(err) {
			// expected result achieved
			err = nil
		}
	}
	return err
}
