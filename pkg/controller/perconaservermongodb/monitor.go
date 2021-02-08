package perconaservermongodb

import (
	"context"
	promClient "github.com/coreos/prometheus-operator/pkg/client/versioned"
	"github.com/percona/percona-server-mongodb-operator/pkg/psmdb/monitor"
	gdv1beta1 "gomod.alauda.cn/ait-apis/grafanadashboard/v1beta1"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/rest"
	"os"
	"reflect"
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
	owner, err := generateOwnerReference(r)
	if err != nil {
		log.Error(err, "Cannot generate OwnerReference")
		return err
	}
	sm := monitor.GenerateServiceMonitor(owner)
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
	m, err := p.MonitoringV1().ServiceMonitors(WatchNamespace).Get(sm.Name, v1.GetOptions{})
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
		} else if !reflect.DeepEqual(m.Spec, sm.Spec) {
			_, err = p.MonitoringV1().ServiceMonitors(WatchNamespace).Update(sm)
			if err != nil {
				log.Error(err, "update servicemonitor fail")
			}
		}
	}
	return err
}

func (r *ReconcilePerconaServerMongoDB) reconcileGrafanaDashboards(enabled bool) error {
	old_gd := &gdv1beta1.GrafanaDashboard{}
	owner, err := generateOwnerReference(r)
	if err != nil {
		log.Error(err, "Generate owner reference fail")
		return err
	}
	gd := monitor.GenerateGrafana(WatchNamespace, owner)
	err = r.client.Get(context.TODO(),
		types.NamespacedName{Name: gd.Name,
			Namespace: gd.Namespace},
		old_gd)
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
		} else if !reflect.DeepEqual(old_gd.Spec, gd.Spec) {
			old_gd.Spec = gd.Spec
			err = r.client.Update(context.TODO(), old_gd)
			if err != nil {
				log.Error(err, "Cannot update grafana dashboard")
			}
		}
	}
	return err
}

func generateOwnerReference(r *ReconcilePerconaServerMongoDB) (v1.OwnerReference, error) {
	deploy_name := os.Getenv("OPERATOR_NAME")
	deploy := &appsv1.Deployment{}
	err := r.client.Get(context.TODO(),
		types.NamespacedName{Name: deploy_name,
			Namespace: WatchNamespace},
		deploy)
	if err != nil {
		log.Error(err, "Cannot get Deployment")
		return v1.OwnerReference{}, err
	}
	owner := v1.OwnerReference{
		APIVersion: appsv1.SchemeGroupVersion.String(),
		Kind:       "Deployment",
		Name:       deploy_name,
		UID:        deploy.UID,
	}
	return owner, err
}
