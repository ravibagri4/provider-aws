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

type EnablerInitParameters struct {

	// Set of account IDs.
	// Can contain one of: the Organization's Administrator Account, or one or more Member Accounts.
	// +listType=set
	AccountIds []*string `json:"accountIds,omitempty" tf:"account_ids,omitempty"`

	// Type of resources to scan.
	// Valid values are EC2, ECR, LAMBDA and LAMBDA_CODE.
	// At least one item is required.
	// +listType=set
	ResourceTypes []*string `json:"resourceTypes,omitempty" tf:"resource_types,omitempty"`
}

type EnablerObservation struct {

	// Set of account IDs.
	// Can contain one of: the Organization's Administrator Account, or one or more Member Accounts.
	// +listType=set
	AccountIds []*string `json:"accountIds,omitempty" tf:"account_ids,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Type of resources to scan.
	// Valid values are EC2, ECR, LAMBDA and LAMBDA_CODE.
	// At least one item is required.
	// +listType=set
	ResourceTypes []*string `json:"resourceTypes,omitempty" tf:"resource_types,omitempty"`
}

type EnablerParameters struct {

	// Set of account IDs.
	// Can contain one of: the Organization's Administrator Account, or one or more Member Accounts.
	// +kubebuilder:validation:Optional
	// +listType=set
	AccountIds []*string `json:"accountIds,omitempty" tf:"account_ids,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Type of resources to scan.
	// Valid values are EC2, ECR, LAMBDA and LAMBDA_CODE.
	// At least one item is required.
	// +kubebuilder:validation:Optional
	// +listType=set
	ResourceTypes []*string `json:"resourceTypes,omitempty" tf:"resource_types,omitempty"`
}

// EnablerSpec defines the desired state of Enabler
type EnablerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EnablerParameters `json:"forProvider"`
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
	InitProvider EnablerInitParameters `json:"initProvider,omitempty"`
}

// EnablerStatus defines the observed state of Enabler.
type EnablerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EnablerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Enabler is the Schema for the Enablers API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Enabler struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.accountIds) || (has(self.initProvider) && has(self.initProvider.accountIds))",message="spec.forProvider.accountIds is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.resourceTypes) || (has(self.initProvider) && has(self.initProvider.resourceTypes))",message="spec.forProvider.resourceTypes is a required parameter"
	Spec   EnablerSpec   `json:"spec"`
	Status EnablerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EnablerList contains a list of Enablers
type EnablerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Enabler `json:"items"`
}

// Repository type metadata.
var (
	Enabler_Kind             = "Enabler"
	Enabler_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Enabler_Kind}.String()
	Enabler_KindAPIVersion   = Enabler_Kind + "." + CRDGroupVersion.String()
	Enabler_GroupVersionKind = CRDGroupVersion.WithKind(Enabler_Kind)
)

func init() {
	SchemeBuilder.Register(&Enabler{}, &EnablerList{})
}
