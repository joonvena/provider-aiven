/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type DatabaseObservation struct {

	// The name of the service database. This property cannot be changed, doing so forces recreation of the resource.
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Default string sort order (`LC_COLLATE`) of the database. The default value is `en_US.UTF-8`. This property cannot be changed, doing so forces recreation of the resource.
	LcCollate *string `json:"lcCollate,omitempty" tf:"lc_collate,omitempty"`

	// Default character classification (`LC_CTYPE`) of the database. The default value is `en_US.UTF-8`. This property cannot be changed, doing so forces recreation of the resource.
	LcCtype *string `json:"lcCtype,omitempty" tf:"lc_ctype,omitempty"`

	// Identifies the project this resource belongs to. To set up proper dependencies please refer to this variable as a reference. This property cannot be changed, doing so forces recreation of the resource.
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Specifies the name of the service that this resource belongs to. To set up proper dependencies please refer to this variable as a reference. This property cannot be changed, doing so forces recreation of the resource.
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// It is recommended to enable this for any production databases containing critical data. The default value is `false`.
	TerminationProtection *bool `json:"terminationProtection,omitempty" tf:"termination_protection,omitempty"`
}

type DatabaseParameters struct {

	// The name of the service database. This property cannot be changed, doing so forces recreation of the resource.
	// +kubebuilder:validation:Optional
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`

	// Default string sort order (`LC_COLLATE`) of the database. The default value is `en_US.UTF-8`. This property cannot be changed, doing so forces recreation of the resource.
	// +kubebuilder:validation:Optional
	LcCollate *string `json:"lcCollate,omitempty" tf:"lc_collate,omitempty"`

	// Default character classification (`LC_CTYPE`) of the database. The default value is `en_US.UTF-8`. This property cannot be changed, doing so forces recreation of the resource.
	// +kubebuilder:validation:Optional
	LcCtype *string `json:"lcCtype,omitempty" tf:"lc_ctype,omitempty"`

	// Identifies the project this resource belongs to. To set up proper dependencies please refer to this variable as a reference. This property cannot be changed, doing so forces recreation of the resource.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Specifies the name of the service that this resource belongs to. To set up proper dependencies please refer to this variable as a reference. This property cannot be changed, doing so forces recreation of the resource.
	// +crossplane:generate:reference:type=Service
	// +crossplane:generate:reference:extractor=github.com/joonvena/provider-aiven/config.ServiceNameExtractor()
	// +kubebuilder:validation:Optional
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// Reference to a Service to populate serviceName.
	// +kubebuilder:validation:Optional
	ServiceNameRef *v1.Reference `json:"serviceNameRef,omitempty" tf:"-"`

	// Selector for a Service to populate serviceName.
	// +kubebuilder:validation:Optional
	ServiceNameSelector *v1.Selector `json:"serviceNameSelector,omitempty" tf:"-"`

	// It is recommended to enable this for any production databases containing critical data. The default value is `false`.
	// +kubebuilder:validation:Optional
	TerminationProtection *bool `json:"terminationProtection,omitempty" tf:"termination_protection,omitempty"`
}

// DatabaseSpec defines the desired state of Database
type DatabaseSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DatabaseParameters `json:"forProvider"`
}

// DatabaseStatus defines the observed state of Database.
type DatabaseStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DatabaseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Database is the Schema for the Databases API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aiven}
type Database struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.databaseName)",message="databaseName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.project)",message="project is a required parameter"
	Spec   DatabaseSpec   `json:"spec"`
	Status DatabaseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DatabaseList contains a list of Databases
type DatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Database `json:"items"`
}

// Repository type metadata.
var (
	Database_Kind             = "Database"
	Database_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Database_Kind}.String()
	Database_KindAPIVersion   = Database_Kind + "." + CRDGroupVersion.String()
	Database_GroupVersionKind = CRDGroupVersion.WithKind(Database_Kind)
)

func init() {
	SchemeBuilder.Register(&Database{}, &DatabaseList{})
}
