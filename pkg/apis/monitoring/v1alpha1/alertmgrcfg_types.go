package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AlertMgrCfgSpec defines the desired state of AlertMgrCfg
// +k8s:openapi-gen=true
type AlertMgrCfgSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	Type   string  `json:"type"`
	Params []Param `json:"params,omitempty"`
}

// Param is a list of alerting receivers.
// +k8s:openapi-gen=true
type Param struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// AlertMgrCfgStatus defines the observed state of AlertMgrCfg
// +k8s:openapi-gen=true
type AlertMgrCfgStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AlertMgrCfg is the Schema for the alertmgrcfgs API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type AlertMgrCfg struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AlertMgrCfgSpec   `json:"spec,omitempty"`
	Status AlertMgrCfgStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AlertMgrCfgList contains a list of AlertMgrCfg
type AlertMgrCfgList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AlertMgrCfg `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AlertMgrCfg{}, &AlertMgrCfgList{})
}
