package monitor

import (
	gdv1beta1 "gomod.alauda.cn/ait-apis/grafanadashboard/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type gDashboard struct {
	name   string
	define string
}

var gDashboards = []gDashboard{gDashboard{"mongodb-cluster-summary", MongoClusterSummary},
	gDashboard{"mongodb-inmemory-detail", MongoInmemoryDetail},
	gDashboard{"mongodb-instance-overview", MongoInstanceOverview},
	gDashboard{"mongodb-instance-summary", MongoInstanceSummary},
	gDashboard{"mongodb-replset-summary", MongoReplsetSummary},
	gDashboard{"mongodb-wiredtiger-details", MongoWiredtigerDetails}}

func GenerateGrafana(ns string, owner metav1.OwnerReference) []gdv1beta1.GrafanaDashboard {
	dashboards := []gdv1beta1.GrafanaDashboard{}
	for _, gd := range gDashboards {
		dashboard := gdv1beta1.GrafanaDashboard{
			ObjectMeta: metav1.ObjectMeta{
				Name:            gd.name,
				Namespace:       ns,
				Finalizers:      []string{"grafana.dashboard.finalizers.ait.cpaas.io"},
				OwnerReferences: []metav1.OwnerReference{owner},
			},
			Spec: gdv1beta1.GrafanaDashboardSpec{
				Folder: "MongoDB",
				Json:   gd.define,
			},
		}
		dashboards = append(dashboards, dashboard)
	}
	return dashboards
}
