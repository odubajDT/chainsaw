package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Delete is a reference to an object that should be deleted
type Delete struct {
	// Timeout for the operation. Overrides the global timeout set in the Configuration.
	// +optional
	Timeout *metav1.Duration `json:"timeout,omitempty"`

	// Bindings defines additional binding key/values.
	// +optional
	Bindings []Binding `json:"bindings,omitempty"`

	// Cluster defines the target cluster (default cluster will be used if not specified and/or overridden).
	// +optional
	Cluster string `json:"cluster,omitempty"`

	// Clusters holds a registry to clusters to support multi-cluster tests.
	// +optional
	Clusters map[string]Cluster `json:"clusters,omitempty"`

	// Template determines whether resources should be considered for templating.
	// +optional
	Template *bool `json:"template,omitempty"`

	// ObjectReference determines objects to be deleted.
	ObjectReference `json:"ref"`

	// Expect defines a list of matched checks to validate the operation outcome.
	// +optional
	Expect []Expectation `json:"expect,omitempty"`

	// DeletionPropagationPolicy decides if a deletion will propagate to the dependents of
	// the object, and how the garbage collector will handle the propagation.
	// Overrides the deletion propagation policy set in the Configuration, the Test and the TestStep.
	// +optional
	// +kubebuilder:validation:Enum:=Orphan;Background;Foreground
	DeletionPropagationPolicy *metav1.DeletionPropagation `json:"deletionPropagationPolicy,omitempty"`
}
