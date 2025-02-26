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

type StandardsSubscriptionInitParameters struct {

	// The ARN of a standard - see below.
	StandardsArn *string `json:"standardsArn,omitempty" tf:"standards_arn,omitempty"`
}

type StandardsSubscriptionObservation struct {

	// The ARN of a resource that represents your subscription to a supported standard.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ARN of a standard - see below.
	StandardsArn *string `json:"standardsArn,omitempty" tf:"standards_arn,omitempty"`
}

type StandardsSubscriptionParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The ARN of a standard - see below.
	// +kubebuilder:validation:Optional
	StandardsArn *string `json:"standardsArn,omitempty" tf:"standards_arn,omitempty"`
}

// StandardsSubscriptionSpec defines the desired state of StandardsSubscription
type StandardsSubscriptionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     StandardsSubscriptionParameters `json:"forProvider"`
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
	InitProvider StandardsSubscriptionInitParameters `json:"initProvider,omitempty"`
}

// StandardsSubscriptionStatus defines the observed state of StandardsSubscription.
type StandardsSubscriptionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        StandardsSubscriptionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// StandardsSubscription is the Schema for the StandardsSubscriptions API. Subscribes to a Security Hub standard.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type StandardsSubscription struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.standardsArn) || (has(self.initProvider) && has(self.initProvider.standardsArn))",message="spec.forProvider.standardsArn is a required parameter"
	Spec   StandardsSubscriptionSpec   `json:"spec"`
	Status StandardsSubscriptionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StandardsSubscriptionList contains a list of StandardsSubscriptions
type StandardsSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StandardsSubscription `json:"items"`
}

// Repository type metadata.
var (
	StandardsSubscription_Kind             = "StandardsSubscription"
	StandardsSubscription_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: StandardsSubscription_Kind}.String()
	StandardsSubscription_KindAPIVersion   = StandardsSubscription_Kind + "." + CRDGroupVersion.String()
	StandardsSubscription_GroupVersionKind = CRDGroupVersion.WithKind(StandardsSubscription_Kind)
)

func init() {
	SchemeBuilder.Register(&StandardsSubscription{}, &StandardsSubscriptionList{})
}
