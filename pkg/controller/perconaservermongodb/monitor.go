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

func (r *ReconcilePerconaServerMongoDB) reconcileMonitor() error {
	sm := monitor.GenerateServiceMonitor()
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Error(err, "error loading kubernetes configuration inside cluster")
		return err
	}
	p, err := promClient.NewForConfig(config)
	m, _ := p.MonitoringV1().ServiceMonitors(WatchNamespace).Get(sm.Name, v1.GetOptions{})
	if m.Name != "" {
		log.Info("ServiceMonitor %s exist", sm.Name)
	} else {
		_, err = p.MonitoringV1().ServiceMonitors(WatchNamespace).Create(sm)
		if err != nil {
			log.Error(err, "create servicemonitor fail")
		}
	}
	grafana := monitor.GenerateGrafana("")
	err = r.client.Get(context.TODO(),
		types.NamespacedName{Name: grafana.Name,
			Namespace: grafana.Namespace},
		grafana)
	if err != nil {
		if errors.IsNotFound(err) {
			err = r.client.Create(context.TODO(), grafana)
			if err != nil {
				log.Error(err, "Cannot create grafana dashboard")
			}
		} else {
			log.Error(err, "get grafana dashboard fail")
		}
	}
	return err
}
