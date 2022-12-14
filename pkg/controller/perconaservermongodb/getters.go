package perconaservermongodb

import (
	"context"

	api "github.com/percona/percona-server-mongodb-operator/pkg/apis/psmdb/v1"
	"github.com/percona/percona-server-mongodb-operator/pkg/psmdb"
	"github.com/pkg/errors"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func (r *ReconcilePerconaServerMongoDB) getMongodPods(cr *api.PerconaServerMongoDB) (corev1.PodList, error) {
	mongodPods := corev1.PodList{}
	err := r.client.List(context.TODO(),
		&mongodPods,
		&client.ListOptions{
			Namespace:     cr.Namespace,
			LabelSelector: labels.SelectorFromSet(mongodLabels(cr)),
		},
	)

	return mongodPods, err
}

func (r *ReconcilePerconaServerMongoDB) getMongosDeployment(cr *api.PerconaServerMongoDB) (appsv1.Deployment, error) {
	mongos := appsv1.Deployment{}

	err := r.client.Get(context.TODO(), cr.MongosNamespacedName(), &mongos)

	return mongos, err

}

func (r *ReconcilePerconaServerMongoDB) getMongosPods(cr *api.PerconaServerMongoDB) (corev1.PodList, error) {
	mongosPods := corev1.PodList{}
	err := r.client.List(context.TODO(),
		&mongosPods,
		&client.ListOptions{
			Namespace:     cr.Namespace,
			LabelSelector: labels.SelectorFromSet(mongosLabels(cr)),
		},
	)

	return mongosPods, err
}

func (r *ReconcilePerconaServerMongoDB) getRSPods(cr *api.PerconaServerMongoDB, rsName string) (corev1.PodList, error) {
	pods := corev1.PodList{}
	err := r.client.List(context.TODO(),
		&pods,
		&client.ListOptions{
			Namespace:     cr.Namespace,
			LabelSelector: labels.SelectorFromSet(rsLabels(cr, rsName)),
		},
	)

	return pods, err
}

func (r *ReconcilePerconaServerMongoDB) getRSSvcs(cr *api.PerconaServerMongoDB, replset *api.ReplsetSpec) (*corev1.ServiceList, error) {
	svcList := &corev1.ServiceList{}
	err := r.client.List(context.TODO(),
		svcList,
		&client.ListOptions{
			Namespace:     cr.Namespace,
			LabelSelector: labels.SelectorFromSet(psmdb.GetCommonServiceLabels(cr, replset)),
		},
	)
	return svcList, err
}

func (r *ReconcilePerconaServerMongoDB) getArbiterStatefulset(cr *api.PerconaServerMongoDB, rs string) (appsv1.StatefulSet, error) {
	list := appsv1.StatefulSetList{}

	l := arbiterLabels(cr)
	l["app.kubernetes.io/replset"] = rs

	err := r.client.List(context.TODO(),
		&list,
		&client.ListOptions{
			Namespace:     cr.Namespace,
			LabelSelector: labels.SelectorFromSet(l),
		},
	)

	if len(list.Items) != 1 {
		return appsv1.StatefulSet{}, errors.Errorf("invalid sfs arbiter count: %d", len(list.Items))
	}

	return list.Items[0], err
}

func (r *ReconcilePerconaServerMongoDB) getRsStatefulset(cr *api.PerconaServerMongoDB, rs string) (appsv1.StatefulSet, error) {
	sts := appsv1.StatefulSet{}

	err := r.client.Get(context.TODO(), cr.StatefulsetNamespacedName(rs), &sts)

	return sts, err
}

func (r *ReconcilePerconaServerMongoDB) getArbiterStatefulsets(cr *api.PerconaServerMongoDB) (appsv1.StatefulSetList, error) {
	list := appsv1.StatefulSetList{}

	err := r.client.List(context.TODO(),
		&list,
		&client.ListOptions{
			Namespace:     cr.Namespace,
			LabelSelector: labels.SelectorFromSet(arbiterLabels(cr)),
		},
	)

	return list, err
}

func (r *ReconcilePerconaServerMongoDB) getMongodStatefulsets(cr *api.PerconaServerMongoDB) (appsv1.StatefulSetList, error) {
	list := appsv1.StatefulSetList{}

	err := r.client.List(context.TODO(),
		&list,
		&client.ListOptions{
			Namespace:     cr.Namespace,
			LabelSelector: labels.SelectorFromSet(mongodLabels(cr)),
		},
	)

	return list, err
}

func (r *ReconcilePerconaServerMongoDB) getAllstatefulsets(cr *api.PerconaServerMongoDB) (appsv1.StatefulSetList, error) {
	list := appsv1.StatefulSetList{}
	filteredList := appsv1.StatefulSetList{}

	err := r.client.List(context.TODO(),
		&list,
		&client.ListOptions{
			Namespace:     cr.Namespace,
			LabelSelector: labels.SelectorFromSet(clusterLabels(cr)),
		},
	)

	for _, sts := range list.Items {
		if metav1.IsControlledBy(&sts, cr) {
			filteredList.Items = append(filteredList.Items, sts)
		}
	}

	return filteredList, err
}

func (r *ReconcilePerconaServerMongoDB) getCfgStatefulset(cr *api.PerconaServerMongoDB) (appsv1.StatefulSet, error) {
	sts := appsv1.StatefulSet{}
	err := r.client.Get(context.TODO(),
		types.NamespacedName{Name: cr.Name + "-" + api.ConfigReplSetName, Namespace: cr.Namespace}, &sts)

	return sts, err
}

func (r *ReconcilePerconaServerMongoDB) getAllPVCs(cr *api.PerconaServerMongoDB) (corev1.PersistentVolumeClaimList, error) {
	list := corev1.PersistentVolumeClaimList{}

	err := r.client.List(context.TODO(),
		&list,
		&client.ListOptions{
			Namespace:     cr.Namespace,
			LabelSelector: labels.SelectorFromSet(clusterLabels(cr)),
		},
	)

	return list, err
}

func (r *ReconcilePerconaServerMongoDB) isReplsetReadyToInit(cr *api.PerconaServerMongoDB, replset *api.ReplsetSpec) bool {
	pods, err := r.getRSPods(cr, replset.Name)
	if err != nil {
		return false
	}

	if replset.Arbiter.Enabled && len(pods.Items) < int(replset.Size+replset.Arbiter.Size) {
		return false
	} else if !replset.Arbiter.Enabled && len(pods.Items) < int(replset.Size) {
		return false
	}

	for _, pod := range pods.Items {
		for _, cs := range pod.Status.ContainerStatuses {
			if !cs.Ready {
				log.Info("Container unhealthy", "pod name", pod.Name, "container name", cs.Name)
				return false
			}
		}
	}

	if replset.Expose.Enabled {
		// In above logic, use pod's name to create svc
		// So here to make sure when pod ready, its svc are ready too
		svcs, err := r.getRSSvcs(cr, replset)
		if err != nil {
			return false
		}
		svcNames := make(map[string]struct{}, len(svcs.Items))
		for _, svc := range svcs.Items {
			svcNames[svc.Name] = struct{}{}
		}
		for _, pod := range pods.Items {
			if _, ok := svcNames[pod.Name]; !ok {
				return false
			}
		}
	}
	return true
}

func clusterLabels(cr *api.PerconaServerMongoDB) map[string]string {
	return map[string]string{
		"app.kubernetes.io/name":       "percona-server-mongodb",
		"app.kubernetes.io/instance":   cr.Name,
		"app.kubernetes.io/managed-by": "percona-server-mongodb-operator",
		"app.kubernetes.io/part-of":    "percona-server-mongodb",
	}
}

func rsLabels(cr *api.PerconaServerMongoDB, rsName string) map[string]string {
	lbls := clusterLabels(cr)
	lbls["app.kubernetes.io/replset"] = rsName
	return lbls
}

func mongodLabels(cr *api.PerconaServerMongoDB) map[string]string {
	lbls := clusterLabels(cr)
	lbls["app.kubernetes.io/component"] = "mongod"
	return lbls
}

func arbiterLabels(cr *api.PerconaServerMongoDB) map[string]string {
	lbls := clusterLabels(cr)
	lbls["app.kubernetes.io/component"] = "arbiter"
	return lbls
}

func mongosLabels(cr *api.PerconaServerMongoDB) map[string]string {
	lbls := clusterLabels(cr)
	lbls["app.kubernetes.io/component"] = "mongos"
	return lbls
}
