package psmdb

import (
	corev1 "k8s.io/api/core/v1"
)

func EntrypointInitContainer(initImageName string, pullPolicy corev1.PullPolicy) corev1.Container {
	// Use root for init-container to bypass ceph pv's permission limit
	var root int64 = 0
	return corev1.Container{
		VolumeMounts: []corev1.VolumeMount{
			{
				Name:      MongodDataVolClaimName,
				MountPath: "/data/db",
			},
		},
		Image:           initImageName,
		Name:            "mongo-init",
		Command:         []string{"/init-entrypoint.sh"},
		ImagePullPolicy: pullPolicy,
		SecurityContext: &corev1.SecurityContext{
			RunAsUser: &root,
		},
	}
}
