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

type GraphInitParameters struct {

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type GraphObservation struct {

	// Date and time, in UTC and extended RFC 3339 format, when the Amazon Detective Graph was created.
	CreatedTime *string `json:"createdTime,omitempty" tf:"created_time,omitempty"`

	// ARN of the Detective Graph.
	GraphArn *string `json:"graphArn,omitempty" tf:"graph_arn,omitempty"`

	// ARN of the Detective Graph.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type GraphParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// GraphSpec defines the desired state of Graph
type GraphSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GraphParameters `json:"forProvider"`
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
	InitProvider GraphInitParameters `json:"initProvider,omitempty"`
}

// GraphStatus defines the observed state of Graph.
type GraphStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GraphObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Graph is the Schema for the Graphs API. Provides a resource to manage an Amazon Detective graph.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Graph struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GraphSpec   `json:"spec"`
	Status            GraphStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GraphList contains a list of Graphs
type GraphList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Graph `json:"items"`
}

// Repository type metadata.
var (
	Graph_Kind             = "Graph"
	Graph_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Graph_Kind}.String()
	Graph_KindAPIVersion   = Graph_Kind + "." + CRDGroupVersion.String()
	Graph_GroupVersionKind = CRDGroupVersion.WithKind(Graph_Kind)
)

func init() {
	SchemeBuilder.Register(&Graph{}, &GraphList{})
}
