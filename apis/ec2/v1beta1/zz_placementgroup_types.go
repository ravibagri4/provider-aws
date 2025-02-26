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

type PlacementGroupInitParameters struct {

	// The number of partitions to create in the
	// placement group.  Can only be specified when the strategy is set to
	// partition.  Valid values are 1 - 7 (default is 2).
	PartitionCount *float64 `json:"partitionCount,omitempty" tf:"partition_count,omitempty"`

	// Determines how placement groups spread instances. Can only be used
	// when the strategy is set to spread. Can be host or rack. host can only be used for Outpost placement groups. Defaults to rack.
	SpreadLevel *string `json:"spreadLevel,omitempty" tf:"spread_level,omitempty"`

	// The placement strategy. Can be cluster, partition or spread.
	Strategy *string `json:"strategy,omitempty" tf:"strategy,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type PlacementGroupObservation struct {

	// Amazon Resource Name (ARN) of the placement group.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The name of the placement group.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The number of partitions to create in the
	// placement group.  Can only be specified when the strategy is set to
	// partition.  Valid values are 1 - 7 (default is 2).
	PartitionCount *float64 `json:"partitionCount,omitempty" tf:"partition_count,omitempty"`

	// The ID of the placement group.
	PlacementGroupID *string `json:"placementGroupId,omitempty" tf:"placement_group_id,omitempty"`

	// Determines how placement groups spread instances. Can only be used
	// when the strategy is set to spread. Can be host or rack. host can only be used for Outpost placement groups. Defaults to rack.
	SpreadLevel *string `json:"spreadLevel,omitempty" tf:"spread_level,omitempty"`

	// The placement strategy. Can be cluster, partition or spread.
	Strategy *string `json:"strategy,omitempty" tf:"strategy,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type PlacementGroupParameters struct {

	// The number of partitions to create in the
	// placement group.  Can only be specified when the strategy is set to
	// partition.  Valid values are 1 - 7 (default is 2).
	// +kubebuilder:validation:Optional
	PartitionCount *float64 `json:"partitionCount,omitempty" tf:"partition_count,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Determines how placement groups spread instances. Can only be used
	// when the strategy is set to spread. Can be host or rack. host can only be used for Outpost placement groups. Defaults to rack.
	// +kubebuilder:validation:Optional
	SpreadLevel *string `json:"spreadLevel,omitempty" tf:"spread_level,omitempty"`

	// The placement strategy. Can be cluster, partition or spread.
	// +kubebuilder:validation:Optional
	Strategy *string `json:"strategy,omitempty" tf:"strategy,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// PlacementGroupSpec defines the desired state of PlacementGroup
type PlacementGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PlacementGroupParameters `json:"forProvider"`
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
	InitProvider PlacementGroupInitParameters `json:"initProvider,omitempty"`
}

// PlacementGroupStatus defines the observed state of PlacementGroup.
type PlacementGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PlacementGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// PlacementGroup is the Schema for the PlacementGroups API. Provides an EC2 placement group.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type PlacementGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.strategy) || (has(self.initProvider) && has(self.initProvider.strategy))",message="spec.forProvider.strategy is a required parameter"
	Spec   PlacementGroupSpec   `json:"spec"`
	Status PlacementGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PlacementGroupList contains a list of PlacementGroups
type PlacementGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PlacementGroup `json:"items"`
}

// Repository type metadata.
var (
	PlacementGroup_Kind             = "PlacementGroup"
	PlacementGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PlacementGroup_Kind}.String()
	PlacementGroup_KindAPIVersion   = PlacementGroup_Kind + "." + CRDGroupVersion.String()
	PlacementGroup_GroupVersionKind = CRDGroupVersion.WithKind(PlacementGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&PlacementGroup{}, &PlacementGroupList{})
}
