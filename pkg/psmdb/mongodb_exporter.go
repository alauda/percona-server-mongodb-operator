package psmdb

import (
	corev1 "k8s.io/api/core/v1"

	api "github.com/percona/percona-server-mongodb-operator/pkg/apis/psmdb/v1"
)

// PMMContainer returns a pmm container from given spec
func MongodbExporterContainer(spec api.MongodbExporterSpec, secrets string) corev1.Container {
	exportPorts := []corev1.ContainerPort{
		{
			ContainerPort: 9104,
			Name:          "exporter-http",
		},
	}

	dbArgsEnv := []corev1.EnvVar{
		{
			Name: "MONGODB_USER",
			ValueFrom: &corev1.EnvVarSource{
				SecretKeyRef: &corev1.SecretKeySelector{
					Key: "MONGODB_CLUSTER_MONITOR_USER",
					LocalObjectReference: corev1.LocalObjectReference{
						Name: secrets,
					},
				},
			},
		},
		{
			Name: "MONGODB_PASSWORD",
			ValueFrom: &corev1.EnvVarSource{
				SecretKeyRef: &corev1.SecretKeySelector{
					Key: "MONGODB_CLUSTER_MONITOR_PASSWORD",
					LocalObjectReference: corev1.LocalObjectReference{
						Name: secrets,
					},
				},
			},
		},
	}

	ExporterArgs := []string{
		"--web.listen-address=:9104",
		"--mongodb.uri",
		"mongodb://$(MONGODB_USER):$(MONGODB_PASSWORD)@127.0.0.1:27017/",
	}

	mongodb_exporter := corev1.Container{
		Name:            "mongod-exporter",
		Image:           spec.Image,
		ImagePullPolicy: corev1.PullIfNotPresent,
		Env:             dbArgsEnv,
		Args:            ExporterArgs,
		Ports:           exportPorts,
	}

	return mongodb_exporter
}
