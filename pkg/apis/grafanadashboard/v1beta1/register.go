// Package v1 contains API Schema definitions for the grafanadashboards.ait.alauda.io v1 API group
// +kubebuilder:object:generate=true
// +groupName=grafanadashboards.ait.alauda.io
package v1beta1

import (
	gdv1beta1 "gomod.alauda.cn/ait-apis/grafanadashboard/v1beta1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

var (
	// GroupVersion is group version used to register these objects
	SchemeGroupVersion = schema.GroupVersion{Group: "ait.alauda.io", Version: "v1beta1"}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: SchemeGroupVersion}

	// AddToScheme adds the types in this group-version to the given scheme.
	AddToScheme = SchemeBuilder.AddToScheme
)

func init() {
	SchemeBuilder.Register(&gdv1beta1.GrafanaDashboard{})
}
