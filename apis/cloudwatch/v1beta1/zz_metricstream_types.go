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

type ExcludeFilterInitParameters struct {

	// An array that defines the metrics you want to exclude for this metric namespace
	// +listType=set
	MetricNames []*string `json:"metricNames,omitempty" tf:"metric_names,omitempty"`

	// Name of the metric namespace in the filter.
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`
}

type ExcludeFilterObservation struct {

	// An array that defines the metrics you want to exclude for this metric namespace
	// +listType=set
	MetricNames []*string `json:"metricNames,omitempty" tf:"metric_names,omitempty"`

	// Name of the metric namespace in the filter.
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`
}

type ExcludeFilterParameters struct {

	// An array that defines the metrics you want to exclude for this metric namespace
	// +kubebuilder:validation:Optional
	// +listType=set
	MetricNames []*string `json:"metricNames,omitempty" tf:"metric_names,omitempty"`

	// Name of the metric namespace in the filter.
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace" tf:"namespace,omitempty"`
}

type IncludeFilterInitParameters struct {

	// An array that defines the metrics you want to include for this metric namespace
	// +listType=set
	MetricNames []*string `json:"metricNames,omitempty" tf:"metric_names,omitempty"`

	// Name of the metric namespace in the filter.
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`
}

type IncludeFilterObservation struct {

	// An array that defines the metrics you want to include for this metric namespace
	// +listType=set
	MetricNames []*string `json:"metricNames,omitempty" tf:"metric_names,omitempty"`

	// Name of the metric namespace in the filter.
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`
}

type IncludeFilterParameters struct {

	// An array that defines the metrics you want to include for this metric namespace
	// +kubebuilder:validation:Optional
	// +listType=set
	MetricNames []*string `json:"metricNames,omitempty" tf:"metric_names,omitempty"`

	// Name of the metric namespace in the filter.
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace" tf:"namespace,omitempty"`
}

type IncludeMetricInitParameters struct {

	// The name of the metric.
	MetricName *string `json:"metricName,omitempty" tf:"metric_name,omitempty"`

	// The namespace of the metric.
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`
}

type IncludeMetricObservation struct {

	// The name of the metric.
	MetricName *string `json:"metricName,omitempty" tf:"metric_name,omitempty"`

	// The namespace of the metric.
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`
}

type IncludeMetricParameters struct {

	// The name of the metric.
	// +kubebuilder:validation:Optional
	MetricName *string `json:"metricName" tf:"metric_name,omitempty"`

	// The namespace of the metric.
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace" tf:"namespace,omitempty"`
}

type MetricStreamInitParameters struct {

	// List of exclusive metric filters. If you specify this parameter, the stream sends metrics from all metric namespaces except for the namespaces and the conditional metric names that you specify here. If you don't specify metric names or provide empty metric names whole metric namespace is excluded. Conflicts with include_filter.
	ExcludeFilter []ExcludeFilterInitParameters `json:"excludeFilter,omitempty" tf:"exclude_filter,omitempty"`

	// ARN of the Amazon Kinesis Firehose delivery stream to use for this metric stream.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/firehose/v1beta1.DeliveryStream
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",false)
	FirehoseArn *string `json:"firehoseArn,omitempty" tf:"firehose_arn,omitempty"`

	// Reference to a DeliveryStream in firehose to populate firehoseArn.
	// +kubebuilder:validation:Optional
	FirehoseArnRef *v1.Reference `json:"firehoseArnRef,omitempty" tf:"-"`

	// Selector for a DeliveryStream in firehose to populate firehoseArn.
	// +kubebuilder:validation:Optional
	FirehoseArnSelector *v1.Selector `json:"firehoseArnSelector,omitempty" tf:"-"`

	// List of inclusive metric filters. If you specify this parameter, the stream sends only the conditional metric names from the metric namespaces that you specify here. If you don't specify metric names or provide empty metric names whole metric namespace is included. Conflicts with exclude_filter.
	IncludeFilter []IncludeFilterInitParameters `json:"includeFilter,omitempty" tf:"include_filter,omitempty"`

	// account observability.
	IncludeLinkedAccountsMetrics *bool `json:"includeLinkedAccountsMetrics,omitempty" tf:"include_linked_accounts_metrics,omitempty"`

	// Friendly name of the metric stream. Conflicts with name_prefix.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Output format for the stream. Possible values are json, opentelemetry0.7, and opentelemetry1.0. For more information about output formats, see Metric streams output formats.
	OutputFormat *string `json:"outputFormat,omitempty" tf:"output_format,omitempty"`

	// ARN of the IAM role that this metric stream will use to access Amazon Kinesis Firehose resources. For more information about role permissions, see Trust between CloudWatch and Kinesis Data Firehose.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// Reference to a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnRef *v1.Reference `json:"roleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnSelector *v1.Selector `json:"roleArnSelector,omitempty" tf:"-"`

	// For each entry in this array, you specify one or more metrics and the list of additional statistics to stream for those metrics. The additional statistics that you can stream depend on the stream's output_format. If the OutputFormat is json, you can stream any additional statistic that is supported by CloudWatch, listed in CloudWatch statistics definitions. If the OutputFormat is opentelemetry0.7 or opentelemetry1.0, you can stream percentile statistics (p99 etc.). See details below.
	StatisticsConfiguration []StatisticsConfigurationInitParameters `json:"statisticsConfiguration,omitempty" tf:"statistics_configuration,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type MetricStreamObservation struct {

	// ARN of the metric stream.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Date and time in RFC3339 format that the metric stream was created.
	CreationDate *string `json:"creationDate,omitempty" tf:"creation_date,omitempty"`

	// List of exclusive metric filters. If you specify this parameter, the stream sends metrics from all metric namespaces except for the namespaces and the conditional metric names that you specify here. If you don't specify metric names or provide empty metric names whole metric namespace is excluded. Conflicts with include_filter.
	ExcludeFilter []ExcludeFilterObservation `json:"excludeFilter,omitempty" tf:"exclude_filter,omitempty"`

	// ARN of the Amazon Kinesis Firehose delivery stream to use for this metric stream.
	FirehoseArn *string `json:"firehoseArn,omitempty" tf:"firehose_arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// List of inclusive metric filters. If you specify this parameter, the stream sends only the conditional metric names from the metric namespaces that you specify here. If you don't specify metric names or provide empty metric names whole metric namespace is included. Conflicts with exclude_filter.
	IncludeFilter []IncludeFilterObservation `json:"includeFilter,omitempty" tf:"include_filter,omitempty"`

	// account observability.
	IncludeLinkedAccountsMetrics *bool `json:"includeLinkedAccountsMetrics,omitempty" tf:"include_linked_accounts_metrics,omitempty"`

	// Date and time in RFC3339 format that the metric stream was last updated.
	LastUpdateDate *string `json:"lastUpdateDate,omitempty" tf:"last_update_date,omitempty"`

	// Friendly name of the metric stream. Conflicts with name_prefix.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Output format for the stream. Possible values are json, opentelemetry0.7, and opentelemetry1.0. For more information about output formats, see Metric streams output formats.
	OutputFormat *string `json:"outputFormat,omitempty" tf:"output_format,omitempty"`

	// ARN of the IAM role that this metric stream will use to access Amazon Kinesis Firehose resources. For more information about role permissions, see Trust between CloudWatch and Kinesis Data Firehose.
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// State of the metric stream. Possible values are running and stopped.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// For each entry in this array, you specify one or more metrics and the list of additional statistics to stream for those metrics. The additional statistics that you can stream depend on the stream's output_format. If the OutputFormat is json, you can stream any additional statistic that is supported by CloudWatch, listed in CloudWatch statistics definitions. If the OutputFormat is opentelemetry0.7 or opentelemetry1.0, you can stream percentile statistics (p99 etc.). See details below.
	StatisticsConfiguration []StatisticsConfigurationObservation `json:"statisticsConfiguration,omitempty" tf:"statistics_configuration,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type MetricStreamParameters struct {

	// List of exclusive metric filters. If you specify this parameter, the stream sends metrics from all metric namespaces except for the namespaces and the conditional metric names that you specify here. If you don't specify metric names or provide empty metric names whole metric namespace is excluded. Conflicts with include_filter.
	// +kubebuilder:validation:Optional
	ExcludeFilter []ExcludeFilterParameters `json:"excludeFilter,omitempty" tf:"exclude_filter,omitempty"`

	// ARN of the Amazon Kinesis Firehose delivery stream to use for this metric stream.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/firehose/v1beta1.DeliveryStream
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",false)
	// +kubebuilder:validation:Optional
	FirehoseArn *string `json:"firehoseArn,omitempty" tf:"firehose_arn,omitempty"`

	// Reference to a DeliveryStream in firehose to populate firehoseArn.
	// +kubebuilder:validation:Optional
	FirehoseArnRef *v1.Reference `json:"firehoseArnRef,omitempty" tf:"-"`

	// Selector for a DeliveryStream in firehose to populate firehoseArn.
	// +kubebuilder:validation:Optional
	FirehoseArnSelector *v1.Selector `json:"firehoseArnSelector,omitempty" tf:"-"`

	// List of inclusive metric filters. If you specify this parameter, the stream sends only the conditional metric names from the metric namespaces that you specify here. If you don't specify metric names or provide empty metric names whole metric namespace is included. Conflicts with exclude_filter.
	// +kubebuilder:validation:Optional
	IncludeFilter []IncludeFilterParameters `json:"includeFilter,omitempty" tf:"include_filter,omitempty"`

	// account observability.
	// +kubebuilder:validation:Optional
	IncludeLinkedAccountsMetrics *bool `json:"includeLinkedAccountsMetrics,omitempty" tf:"include_linked_accounts_metrics,omitempty"`

	// Friendly name of the metric stream. Conflicts with name_prefix.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Output format for the stream. Possible values are json, opentelemetry0.7, and opentelemetry1.0. For more information about output formats, see Metric streams output formats.
	// +kubebuilder:validation:Optional
	OutputFormat *string `json:"outputFormat,omitempty" tf:"output_format,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// ARN of the IAM role that this metric stream will use to access Amazon Kinesis Firehose resources. For more information about role permissions, see Trust between CloudWatch and Kinesis Data Firehose.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// Reference to a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnRef *v1.Reference `json:"roleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnSelector *v1.Selector `json:"roleArnSelector,omitempty" tf:"-"`

	// For each entry in this array, you specify one or more metrics and the list of additional statistics to stream for those metrics. The additional statistics that you can stream depend on the stream's output_format. If the OutputFormat is json, you can stream any additional statistic that is supported by CloudWatch, listed in CloudWatch statistics definitions. If the OutputFormat is opentelemetry0.7 or opentelemetry1.0, you can stream percentile statistics (p99 etc.). See details below.
	// +kubebuilder:validation:Optional
	StatisticsConfiguration []StatisticsConfigurationParameters `json:"statisticsConfiguration,omitempty" tf:"statistics_configuration,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type StatisticsConfigurationInitParameters struct {

	// The additional statistics to stream for the metrics listed in include_metrics.
	// +listType=set
	AdditionalStatistics []*string `json:"additionalStatistics,omitempty" tf:"additional_statistics,omitempty"`

	// An array that defines the metrics that are to have additional statistics streamed. See details below.
	IncludeMetric []IncludeMetricInitParameters `json:"includeMetric,omitempty" tf:"include_metric,omitempty"`
}

type StatisticsConfigurationObservation struct {

	// The additional statistics to stream for the metrics listed in include_metrics.
	// +listType=set
	AdditionalStatistics []*string `json:"additionalStatistics,omitempty" tf:"additional_statistics,omitempty"`

	// An array that defines the metrics that are to have additional statistics streamed. See details below.
	IncludeMetric []IncludeMetricObservation `json:"includeMetric,omitempty" tf:"include_metric,omitempty"`
}

type StatisticsConfigurationParameters struct {

	// The additional statistics to stream for the metrics listed in include_metrics.
	// +kubebuilder:validation:Optional
	// +listType=set
	AdditionalStatistics []*string `json:"additionalStatistics" tf:"additional_statistics,omitempty"`

	// An array that defines the metrics that are to have additional statistics streamed. See details below.
	// +kubebuilder:validation:Optional
	IncludeMetric []IncludeMetricParameters `json:"includeMetric" tf:"include_metric,omitempty"`
}

// MetricStreamSpec defines the desired state of MetricStream
type MetricStreamSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MetricStreamParameters `json:"forProvider"`
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
	InitProvider MetricStreamInitParameters `json:"initProvider,omitempty"`
}

// MetricStreamStatus defines the observed state of MetricStream.
type MetricStreamStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MetricStreamObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// MetricStream is the Schema for the MetricStreams API. Provides a CloudWatch Metric Stream resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type MetricStream struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.outputFormat) || (has(self.initProvider) && has(self.initProvider.outputFormat))",message="spec.forProvider.outputFormat is a required parameter"
	Spec   MetricStreamSpec   `json:"spec"`
	Status MetricStreamStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MetricStreamList contains a list of MetricStreams
type MetricStreamList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MetricStream `json:"items"`
}

// Repository type metadata.
var (
	MetricStream_Kind             = "MetricStream"
	MetricStream_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MetricStream_Kind}.String()
	MetricStream_KindAPIVersion   = MetricStream_Kind + "." + CRDGroupVersion.String()
	MetricStream_GroupVersionKind = CRDGroupVersion.WithKind(MetricStream_Kind)
)

func init() {
	SchemeBuilder.Register(&MetricStream{}, &MetricStreamList{})
}
