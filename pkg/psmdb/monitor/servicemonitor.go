package monitor

import (
	v1 "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const ServiceMonitorName = "redis-sentinel"

// GenerateServiceMonitor is a method that will generate a ServiceMonitor interface
func GenerateServiceMonitor() *v1.ServiceMonitor {
	sm := &v1.ServiceMonitor{
		ObjectMeta: metav1.ObjectMeta{
			Name: ServiceMonitorName,
			Labels: map[string]string{
				"prometheus": "kube-prometheus",
			},
		},
		Spec: v1.ServiceMonitorSpec{
			Selector: metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app.kubernetes.io/part-of": "redis-failover",
				},
			},
			NamespaceSelector: v1.NamespaceSelector{
				Any: true,
			},
			Endpoints: []v1.Endpoint{
				{
					HonorLabels:   true,
					Port:          "http-metrics",
					Path:          "/metrics",
					Interval:      "10s",
					ScrapeTimeout: "10s",
				},
				{
					HonorLabels:   true,
					Port:          "metrics",
					Path:          "/metrics",
					Interval:      "10s",
					ScrapeTimeout: "10s",
				},
			},
		},
	}
	return sm
}
