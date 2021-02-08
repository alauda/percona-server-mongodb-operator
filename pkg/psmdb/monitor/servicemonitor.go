package monitor

import (
	v1 "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const ServiceMonitorName = "mongo-servicemonitor"

// GenerateServiceMonitor is a method that will generate a ServiceMonitor interface
func GenerateServiceMonitor(owner metav1.OwnerReference) *v1.ServiceMonitor {
	sm := &v1.ServiceMonitor{
		ObjectMeta: metav1.ObjectMeta{
			Name: ServiceMonitorName,
			Labels: map[string]string{
				"prometheus": "kube-prometheus",
			},
			OwnerReferences: []metav1.OwnerReference{owner},
		},
		Spec: v1.ServiceMonitorSpec{
			Selector: metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app.kubernetes.io/component": "mongod",
				},
			},
			NamespaceSelector: v1.NamespaceSelector{
				Any: true,
			},
			Endpoints: []v1.Endpoint{
				{
					Port:     "mongod-exporter",
					Path:     "/metrics",
					Interval: "10s",
				},
			},
		},
	}
	return sm
}
