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

type MaintenanceWindowInitParameters struct {

	// Whether targets must be registered with the Maintenance Window before tasks can be defined for those targets.
	AllowUnassociatedTargets *bool `json:"allowUnassociatedTargets,omitempty" tf:"allow_unassociated_targets,omitempty"`

	// The number of hours before the end of the Maintenance Window that Systems Manager stops scheduling new tasks for execution.
	Cutoff *float64 `json:"cutoff,omitempty" tf:"cutoff,omitempty"`

	// A description for the maintenance window.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The duration of the Maintenance Window in hours.
	Duration *float64 `json:"duration,omitempty" tf:"duration,omitempty"`

	// Whether the maintenance window is enabled. Default: true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Timestamp in ISO-8601 extended format when to no longer run the maintenance window.
	EndDate *string `json:"endDate,omitempty" tf:"end_date,omitempty"`

	// The name of the maintenance window.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The schedule of the Maintenance Window in the form of a cron or rate expression.
	Schedule *string `json:"schedule,omitempty" tf:"schedule,omitempty"`

	// The number of days to wait after the date and time specified by a CRON expression before running the maintenance window.
	ScheduleOffset *float64 `json:"scheduleOffset,omitempty" tf:"schedule_offset,omitempty"`

	// Timezone for schedule in Internet Assigned Numbers Authority (IANA) Time Zone Database format. For example: America/Los_Angeles, etc/UTC, or Asia/Seoul.
	ScheduleTimezone *string `json:"scheduleTimezone,omitempty" tf:"schedule_timezone,omitempty"`

	// Timestamp in ISO-8601 extended format when to begin the maintenance window.
	StartDate *string `json:"startDate,omitempty" tf:"start_date,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type MaintenanceWindowObservation struct {

	// Whether targets must be registered with the Maintenance Window before tasks can be defined for those targets.
	AllowUnassociatedTargets *bool `json:"allowUnassociatedTargets,omitempty" tf:"allow_unassociated_targets,omitempty"`

	// The number of hours before the end of the Maintenance Window that Systems Manager stops scheduling new tasks for execution.
	Cutoff *float64 `json:"cutoff,omitempty" tf:"cutoff,omitempty"`

	// A description for the maintenance window.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The duration of the Maintenance Window in hours.
	Duration *float64 `json:"duration,omitempty" tf:"duration,omitempty"`

	// Whether the maintenance window is enabled. Default: true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Timestamp in ISO-8601 extended format when to no longer run the maintenance window.
	EndDate *string `json:"endDate,omitempty" tf:"end_date,omitempty"`

	// The ID of the maintenance window.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the maintenance window.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The schedule of the Maintenance Window in the form of a cron or rate expression.
	Schedule *string `json:"schedule,omitempty" tf:"schedule,omitempty"`

	// The number of days to wait after the date and time specified by a CRON expression before running the maintenance window.
	ScheduleOffset *float64 `json:"scheduleOffset,omitempty" tf:"schedule_offset,omitempty"`

	// Timezone for schedule in Internet Assigned Numbers Authority (IANA) Time Zone Database format. For example: America/Los_Angeles, etc/UTC, or Asia/Seoul.
	ScheduleTimezone *string `json:"scheduleTimezone,omitempty" tf:"schedule_timezone,omitempty"`

	// Timestamp in ISO-8601 extended format when to begin the maintenance window.
	StartDate *string `json:"startDate,omitempty" tf:"start_date,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type MaintenanceWindowParameters struct {

	// Whether targets must be registered with the Maintenance Window before tasks can be defined for those targets.
	// +kubebuilder:validation:Optional
	AllowUnassociatedTargets *bool `json:"allowUnassociatedTargets,omitempty" tf:"allow_unassociated_targets,omitempty"`

	// The number of hours before the end of the Maintenance Window that Systems Manager stops scheduling new tasks for execution.
	// +kubebuilder:validation:Optional
	Cutoff *float64 `json:"cutoff,omitempty" tf:"cutoff,omitempty"`

	// A description for the maintenance window.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The duration of the Maintenance Window in hours.
	// +kubebuilder:validation:Optional
	Duration *float64 `json:"duration,omitempty" tf:"duration,omitempty"`

	// Whether the maintenance window is enabled. Default: true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Timestamp in ISO-8601 extended format when to no longer run the maintenance window.
	// +kubebuilder:validation:Optional
	EndDate *string `json:"endDate,omitempty" tf:"end_date,omitempty"`

	// The name of the maintenance window.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The schedule of the Maintenance Window in the form of a cron or rate expression.
	// +kubebuilder:validation:Optional
	Schedule *string `json:"schedule,omitempty" tf:"schedule,omitempty"`

	// The number of days to wait after the date and time specified by a CRON expression before running the maintenance window.
	// +kubebuilder:validation:Optional
	ScheduleOffset *float64 `json:"scheduleOffset,omitempty" tf:"schedule_offset,omitempty"`

	// Timezone for schedule in Internet Assigned Numbers Authority (IANA) Time Zone Database format. For example: America/Los_Angeles, etc/UTC, or Asia/Seoul.
	// +kubebuilder:validation:Optional
	ScheduleTimezone *string `json:"scheduleTimezone,omitempty" tf:"schedule_timezone,omitempty"`

	// Timestamp in ISO-8601 extended format when to begin the maintenance window.
	// +kubebuilder:validation:Optional
	StartDate *string `json:"startDate,omitempty" tf:"start_date,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// MaintenanceWindowSpec defines the desired state of MaintenanceWindow
type MaintenanceWindowSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MaintenanceWindowParameters `json:"forProvider"`
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
	InitProvider MaintenanceWindowInitParameters `json:"initProvider,omitempty"`
}

// MaintenanceWindowStatus defines the observed state of MaintenanceWindow.
type MaintenanceWindowStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MaintenanceWindowObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// MaintenanceWindow is the Schema for the MaintenanceWindows API. Provides an SSM Maintenance Window resource
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type MaintenanceWindow struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.cutoff) || (has(self.initProvider) && has(self.initProvider.cutoff))",message="spec.forProvider.cutoff is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.duration) || (has(self.initProvider) && has(self.initProvider.duration))",message="spec.forProvider.duration is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.schedule) || (has(self.initProvider) && has(self.initProvider.schedule))",message="spec.forProvider.schedule is a required parameter"
	Spec   MaintenanceWindowSpec   `json:"spec"`
	Status MaintenanceWindowStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MaintenanceWindowList contains a list of MaintenanceWindows
type MaintenanceWindowList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MaintenanceWindow `json:"items"`
}

// Repository type metadata.
var (
	MaintenanceWindow_Kind             = "MaintenanceWindow"
	MaintenanceWindow_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MaintenanceWindow_Kind}.String()
	MaintenanceWindow_KindAPIVersion   = MaintenanceWindow_Kind + "." + CRDGroupVersion.String()
	MaintenanceWindow_GroupVersionKind = CRDGroupVersion.WithKind(MaintenanceWindow_Kind)
)

func init() {
	SchemeBuilder.Register(&MaintenanceWindow{}, &MaintenanceWindowList{})
}
