// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type SchemaInitParameters struct {

	// The schema specification. Must be a valid Open API 3.0 spec.
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// The description of the schema. Maximum of 256 characters.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the schema. Maximum of 385 characters consisting of lower case letters, upper case letters, ., -, _, @.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of the registry in which this schema belongs.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/schemas/v1beta1.Registry
	RegistryName *string `json:"registryName,omitempty" tf:"registry_name,omitempty"`

	// Reference to a Registry in schemas to populate registryName.
	// +kubebuilder:validation:Optional
	RegistryNameRef *v1.Reference `json:"registryNameRef,omitempty" tf:"-"`

	// Selector for a Registry in schemas to populate registryName.
	// +kubebuilder:validation:Optional
	RegistryNameSelector *v1.Selector `json:"registryNameSelector,omitempty" tf:"-"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The type of the schema. Valid values: OpenApi3 or JSONSchemaDraft4.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type SchemaObservation struct {

	// The Amazon Resource Name (ARN) of the discoverer.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The schema specification. Must be a valid Open API 3.0 spec.
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// The description of the schema. Maximum of 256 characters.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The last modified date of the schema.
	LastModified *string `json:"lastModified,omitempty" tf:"last_modified,omitempty"`

	// The name of the schema. Maximum of 385 characters consisting of lower case letters, upper case letters, ., -, _, @.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of the registry in which this schema belongs.
	RegistryName *string `json:"registryName,omitempty" tf:"registry_name,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The type of the schema. Valid values: OpenApi3 or JSONSchemaDraft4.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The version of the schema.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`

	// The created date of the version of the schema.
	VersionCreatedDate *string `json:"versionCreatedDate,omitempty" tf:"version_created_date,omitempty"`
}

type SchemaParameters struct {

	// The schema specification. Must be a valid Open API 3.0 spec.
	// +kubebuilder:validation:Optional
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// The description of the schema. Maximum of 256 characters.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the schema. Maximum of 385 characters consisting of lower case letters, upper case letters, ., -, _, @.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The name of the registry in which this schema belongs.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/schemas/v1beta1.Registry
	// +kubebuilder:validation:Optional
	RegistryName *string `json:"registryName,omitempty" tf:"registry_name,omitempty"`

	// Reference to a Registry in schemas to populate registryName.
	// +kubebuilder:validation:Optional
	RegistryNameRef *v1.Reference `json:"registryNameRef,omitempty" tf:"-"`

	// Selector for a Registry in schemas to populate registryName.
	// +kubebuilder:validation:Optional
	RegistryNameSelector *v1.Selector `json:"registryNameSelector,omitempty" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The type of the schema. Valid values: OpenApi3 or JSONSchemaDraft4.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// SchemaSpec defines the desired state of Schema
type SchemaSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SchemaParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider SchemaInitParameters `json:"initProvider,omitempty"`
}

// SchemaStatus defines the observed state of Schema.
type SchemaStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SchemaObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Schema is the Schema for the Schemas API. Provides an EventBridge Schema resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Schema struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.content) || (has(self.initProvider) && has(self.initProvider.content))",message="spec.forProvider.content is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	Spec   SchemaSpec   `json:"spec"`
	Status SchemaStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SchemaList contains a list of Schemas
type SchemaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Schema `json:"items"`
}

// Repository type metadata.
var (
	Schema_Kind             = "Schema"
	Schema_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Schema_Kind}.String()
	Schema_KindAPIVersion   = Schema_Kind + "." + CRDGroupVersion.String()
	Schema_GroupVersionKind = CRDGroupVersion.WithKind(Schema_Kind)
)

func init() {
	SchemeBuilder.Register(&Schema{}, &SchemaList{})
}
