package stub

import (
	"crypto/rand"
	"encoding/base64"

	"github.com/Percona-Lab/percona-server-mongodb-operator/pkg/apis/psmdb/v1alpha1"

	"github.com/operator-framework/operator-sdk/pkg/sdk"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	mongoDBSecretsDirName    = "mongodb-secrets"
	mongoDBSecretsDir        = "/etc/mongodb-secrets"
	mongoDbSecretMongoKeyVal = "mongodb-key"
)

// generateMongoDBKey generates a 1024 byte-length random key for MongoDB Internal Authentication
//
// See: https://docs.mongodb.com/manual/core/security-internal-authentication/#keyfiles
//
func generateMongoDBKey() (string, error) {
	b := make([]byte, 768)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), err
}

// newPSMDBMongoKeySecret returns a Core API Secret structure containing a new MongoDB Internal
// Authentication key
func newPSMDBMongoKeySecret(m *v1alpha1.PerconaServerMongoDB) (*corev1.Secret, error) {
	key, err := generateMongoDBKey()
	if err != nil {
		return nil, err
	}
	secret := &corev1.Secret{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Secret",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      m.Spec.Secrets.Key,
			Namespace: m.Namespace,
		},
		StringData: map[string]string{
			mongoDbSecretMongoKeyVal: key,
		},
	}
	addOwnerRefToObject(secret, asOwner(m))
	return secret, nil
}

// getPSMDBSecret retrieves a Kubernetes Secret
func getPSMDBSecret(m *v1alpha1.PerconaServerMongoDB, secretName string) (*corev1.Secret, error) {
	secret := &corev1.Secret{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Secret",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      secretName,
			Namespace: m.Namespace,
		},
	}
	err := sdk.Get(secret)
	return secret, err
}