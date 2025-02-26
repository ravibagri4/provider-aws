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

type ResourcePolicyInitParameters struct {

	// Details of the resource policy, including the identity of the principal that is enabled to put logs to this account. This is formatted as a JSON string. Maximum length of 5120 characters.
	PolicyDocument *string `json:"policyDocument,omitempty" tf:"policy_document,omitempty"`
}

type ResourcePolicyObservation struct {

	// The name of the CloudWatch log resource policy
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Details of the resource policy, including the identity of the principal that is enabled to put logs to this account. This is formatted as a JSON string. Maximum length of 5120 characters.
	PolicyDocument *string `json:"policyDocument,omitempty" tf:"policy_document,omitempty"`
}

type ResourcePolicyParameters struct {

	// Details of the resource policy, including the identity of the principal that is enabled to put logs to this account. This is formatted as a JSON string. Maximum length of 5120 characters.
	// +kubebuilder:validation:Optional
	PolicyDocument *string `json:"policyDocument,omitempty" tf:"policy_document,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// ResourcePolicySpec defines the desired state of ResourcePolicy
type ResourcePolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ResourcePolicyParameters `json:"forProvider"`
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
	InitProvider ResourcePolicyInitParameters `json:"initProvider,omitempty"`
}

// ResourcePolicyStatus defines the observed state of ResourcePolicy.
type ResourcePolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ResourcePolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ResourcePolicy is the Schema for the ResourcePolicys API. Provides a resource to manage a CloudWatch log resource policy
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ResourcePolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.policyDocument) || (has(self.initProvider) && has(self.initProvider.policyDocument))",message="spec.forProvider.policyDocument is a required parameter"
	Spec   ResourcePolicySpec   `json:"spec"`
	Status ResourcePolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ResourcePolicyList contains a list of ResourcePolicys
type ResourcePolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResourcePolicy `json:"items"`
}

// Repository type metadata.
var (
	ResourcePolicy_Kind             = "ResourcePolicy"
	ResourcePolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ResourcePolicy_Kind}.String()
	ResourcePolicy_KindAPIVersion   = ResourcePolicy_Kind + "." + CRDGroupVersion.String()
	ResourcePolicy_GroupVersionKind = CRDGroupVersion.WithKind(ResourcePolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&ResourcePolicy{}, &ResourcePolicyList{})
}
