// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

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

type ClustersInitParameters struct {

	// (String) The UID of the host cluster.
	// The UID of the host cluster.
	ClusterUID *string `json:"clusterUid,omitempty" tf:"cluster_uid,omitempty"`

	// (String) The host DNS wildcard for the cluster. i.e. *.dev or *test.com
	// The host DNS wildcard for the cluster. i.e. `*.dev` or `*test.com`
	HostDNS *string `json:"hostDns,omitempty" tf:"host_dns,omitempty"`
}

type ClustersObservation struct {

	// (String) The UID of the host cluster.
	// The UID of the host cluster.
	ClusterUID *string `json:"clusterUid,omitempty" tf:"cluster_uid,omitempty"`

	// (String) The host DNS wildcard for the cluster. i.e. *.dev or *test.com
	// The host DNS wildcard for the cluster. i.e. `*.dev` or `*test.com`
	HostDNS *string `json:"hostDns,omitempty" tf:"host_dns,omitempty"`
}

type ClustersParameters struct {

	// (String) The UID of the host cluster.
	// The UID of the host cluster.
	// +kubebuilder:validation:Optional
	ClusterUID *string `json:"clusterUid" tf:"cluster_uid,omitempty"`

	// (String) The host DNS wildcard for the cluster. i.e. *.dev or *test.com
	// The host DNS wildcard for the cluster. i.e. `*.dev` or `*test.com`
	// +kubebuilder:validation:Optional
	HostDNS *string `json:"hostDns,omitempty" tf:"host_dns,omitempty"`
}

type ConfigInitParameters struct {

	// (Number) The CPU limit in millicores.
	// The CPU limit in millicores.
	CPUMillicore *float64 `json:"cpuMillicore,omitempty" tf:"cpu_millicore,omitempty"`

	// (String) The host endpoint type. Allowed values are 'Ingress' or 'LoadBalancer'. Defaults to 'Ingress'.
	// The host endpoint type. Allowed values are 'Ingress' or 'LoadBalancer'. Defaults to 'Ingress'.
	HostEndpointType *string `json:"hostEndpointType,omitempty" tf:"host_endpoint_type,omitempty"`

	// (String) The Kubernetes distribution, allowed values are k3s and cncf_k8s.
	// The Kubernetes distribution, allowed values are `k3s` and `cncf_k8s`.
	K8SDistribution *string `json:"k8sDistribution,omitempty" tf:"k8s_distribution,omitempty"`

	// (Number) The memory limit in megabytes (MB).
	// The memory limit in megabytes (MB).
	MemoryInMb *float64 `json:"memoryInMb,omitempty" tf:"memory_in_mb,omitempty"`

	// (Number) The allowed oversubscription percentage.
	// The allowed oversubscription percentage.
	OversubscriptionPercent *float64 `json:"oversubscriptionPercent,omitempty" tf:"oversubscription_percent,omitempty"`

	// (Number) The storage limit in gigabytes (GB).
	// The storage limit in gigabytes (GB).
	StorageInGb *float64 `json:"storageInGb,omitempty" tf:"storage_in_gb,omitempty"`

	// (String)
	Values *string `json:"values,omitempty" tf:"values,omitempty"`
}

type ConfigObservation struct {

	// (Number) The CPU limit in millicores.
	// The CPU limit in millicores.
	CPUMillicore *float64 `json:"cpuMillicore,omitempty" tf:"cpu_millicore,omitempty"`

	// (String) The host endpoint type. Allowed values are 'Ingress' or 'LoadBalancer'. Defaults to 'Ingress'.
	// The host endpoint type. Allowed values are 'Ingress' or 'LoadBalancer'. Defaults to 'Ingress'.
	HostEndpointType *string `json:"hostEndpointType,omitempty" tf:"host_endpoint_type,omitempty"`

	// (String) The Kubernetes distribution, allowed values are k3s and cncf_k8s.
	// The Kubernetes distribution, allowed values are `k3s` and `cncf_k8s`.
	K8SDistribution *string `json:"k8sDistribution,omitempty" tf:"k8s_distribution,omitempty"`

	// (Number) The memory limit in megabytes (MB).
	// The memory limit in megabytes (MB).
	MemoryInMb *float64 `json:"memoryInMb,omitempty" tf:"memory_in_mb,omitempty"`

	// (Number) The allowed oversubscription percentage.
	// The allowed oversubscription percentage.
	OversubscriptionPercent *float64 `json:"oversubscriptionPercent,omitempty" tf:"oversubscription_percent,omitempty"`

	// (Number) The storage limit in gigabytes (GB).
	// The storage limit in gigabytes (GB).
	StorageInGb *float64 `json:"storageInGb,omitempty" tf:"storage_in_gb,omitempty"`

	// (String)
	Values *string `json:"values,omitempty" tf:"values,omitempty"`
}

type ConfigParameters struct {

	// (Number) The CPU limit in millicores.
	// The CPU limit in millicores.
	// +kubebuilder:validation:Optional
	CPUMillicore *float64 `json:"cpuMillicore,omitempty" tf:"cpu_millicore,omitempty"`

	// (String) The host endpoint type. Allowed values are 'Ingress' or 'LoadBalancer'. Defaults to 'Ingress'.
	// The host endpoint type. Allowed values are 'Ingress' or 'LoadBalancer'. Defaults to 'Ingress'.
	// +kubebuilder:validation:Optional
	HostEndpointType *string `json:"hostEndpointType,omitempty" tf:"host_endpoint_type,omitempty"`

	// (String) The Kubernetes distribution, allowed values are k3s and cncf_k8s.
	// The Kubernetes distribution, allowed values are `k3s` and `cncf_k8s`.
	// +kubebuilder:validation:Optional
	K8SDistribution *string `json:"k8sDistribution,omitempty" tf:"k8s_distribution,omitempty"`

	// (Number) The memory limit in megabytes (MB).
	// The memory limit in megabytes (MB).
	// +kubebuilder:validation:Optional
	MemoryInMb *float64 `json:"memoryInMb,omitempty" tf:"memory_in_mb,omitempty"`

	// (Number) The allowed oversubscription percentage.
	// The allowed oversubscription percentage.
	// +kubebuilder:validation:Optional
	OversubscriptionPercent *float64 `json:"oversubscriptionPercent,omitempty" tf:"oversubscription_percent,omitempty"`

	// (Number) The storage limit in gigabytes (GB).
	// The storage limit in gigabytes (GB).
	// +kubebuilder:validation:Optional
	StorageInGb *float64 `json:"storageInGb,omitempty" tf:"storage_in_gb,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Values *string `json:"values,omitempty" tf:"values,omitempty"`
}

type GroupClusterProfileInitParameters struct {

	// (String) The ID of this resource.
	// The ID of the cluster profile.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Block List) For packs of type spectro, helm, and manifest, at least one pack must be specified. (see below for nested schema)
	// For packs of type `spectro`, `helm`, and `manifest`, at least one pack must be specified.
	Pack []GroupClusterProfilePackInitParameters `json:"pack,omitempty" tf:"pack,omitempty"`
}

type GroupClusterProfileObservation struct {

	// (String) The ID of this resource.
	// The ID of the cluster profile.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Block List) For packs of type spectro, helm, and manifest, at least one pack must be specified. (see below for nested schema)
	// For packs of type `spectro`, `helm`, and `manifest`, at least one pack must be specified.
	Pack []GroupClusterProfilePackObservation `json:"pack,omitempty" tf:"pack,omitempty"`
}

type GroupClusterProfilePackInitParameters struct {

	// (Block List) (see below for nested schema)
	Manifest []GroupClusterProfilePackManifestInitParameters `json:"manifest,omitempty" tf:"manifest,omitempty"`

	// (String) Name of the cluster group
	// The name of the pack. The name must be unique within the cluster profile.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The registry UID of the pack. The registry UID is the unique identifier of the registry. This attribute is required if there is more than one registry that contains a pack with the same name.
	// The registry UID of the pack. The registry UID is the unique identifier of the registry. This attribute is required if there is more than one registry that contains a pack with the same name.
	RegistryUID *string `json:"registryUid,omitempty" tf:"registry_uid,omitempty"`

	// (String) The tag of the pack. The tag is the version of the pack. This attribute is required if the pack type is spectro or helm.
	// The tag of the pack. The tag is the version of the pack. This attribute is required if the pack type is `spectro` or `helm`.
	Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`

	// (String) The type of the pack. Allowed values are spectro, manifest or helm. The default value is spectro.
	// The type of the pack. Allowed values are `spectro`, `manifest` or `helm`. The default value is `spectro`.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// (String) The unique identifier of the pack. The value can be looked up using the spectrocloud_pack data source. This value is required if the pack type is spectro.
	// The unique identifier of the pack. The value can be looked up using the [`spectrocloud_pack`](https://registry.io/providers/spectrocloud/spectrocloud/latest/docs/data-sources/pack) data source. This value is required if the pack type is `spectro`.
	UID *string `json:"uid,omitempty" tf:"uid,omitempty"`

	// (String)
	// The values of the pack. The values are the configuration values of the pack. The values are specified in YAML format.
	Values *string `json:"values,omitempty" tf:"values,omitempty"`
}

type GroupClusterProfilePackManifestInitParameters struct {

	// (String) The content of the manifest. The content is the YAML content of the manifest.
	// The content of the manifest. The content is the YAML content of the manifest.
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// (String) Name of the cluster group
	// The name of the manifest. The name must be unique within the pack.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type GroupClusterProfilePackManifestObservation struct {

	// (String) The content of the manifest. The content is the YAML content of the manifest.
	// The content of the manifest. The content is the YAML content of the manifest.
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// (String) Name of the cluster group
	// The name of the manifest. The name must be unique within the pack.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The unique identifier of the pack. The value can be looked up using the spectrocloud_pack data source. This value is required if the pack type is spectro.
	UID *string `json:"uid,omitempty" tf:"uid,omitempty"`
}

type GroupClusterProfilePackManifestParameters struct {

	// (String) The content of the manifest. The content is the YAML content of the manifest.
	// The content of the manifest. The content is the YAML content of the manifest.
	// +kubebuilder:validation:Optional
	Content *string `json:"content" tf:"content,omitempty"`

	// (String) Name of the cluster group
	// The name of the manifest. The name must be unique within the pack.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`
}

type GroupClusterProfilePackObservation struct {

	// (Block List) (see below for nested schema)
	Manifest []GroupClusterProfilePackManifestObservation `json:"manifest,omitempty" tf:"manifest,omitempty"`

	// (String) Name of the cluster group
	// The name of the pack. The name must be unique within the cluster profile.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The registry UID of the pack. The registry UID is the unique identifier of the registry. This attribute is required if there is more than one registry that contains a pack with the same name.
	// The registry UID of the pack. The registry UID is the unique identifier of the registry. This attribute is required if there is more than one registry that contains a pack with the same name.
	RegistryUID *string `json:"registryUid,omitempty" tf:"registry_uid,omitempty"`

	// (String) The tag of the pack. The tag is the version of the pack. This attribute is required if the pack type is spectro or helm.
	// The tag of the pack. The tag is the version of the pack. This attribute is required if the pack type is `spectro` or `helm`.
	Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`

	// (String) The type of the pack. Allowed values are spectro, manifest or helm. The default value is spectro.
	// The type of the pack. Allowed values are `spectro`, `manifest` or `helm`. The default value is `spectro`.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// (String) The unique identifier of the pack. The value can be looked up using the spectrocloud_pack data source. This value is required if the pack type is spectro.
	// The unique identifier of the pack. The value can be looked up using the [`spectrocloud_pack`](https://registry.io/providers/spectrocloud/spectrocloud/latest/docs/data-sources/pack) data source. This value is required if the pack type is `spectro`.
	UID *string `json:"uid,omitempty" tf:"uid,omitempty"`

	// (String)
	// The values of the pack. The values are the configuration values of the pack. The values are specified in YAML format.
	Values *string `json:"values,omitempty" tf:"values,omitempty"`
}

type GroupClusterProfilePackParameters struct {

	// (Block List) (see below for nested schema)
	// +kubebuilder:validation:Optional
	Manifest []GroupClusterProfilePackManifestParameters `json:"manifest,omitempty" tf:"manifest,omitempty"`

	// (String) Name of the cluster group
	// The name of the pack. The name must be unique within the cluster profile.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// (String) The registry UID of the pack. The registry UID is the unique identifier of the registry. This attribute is required if there is more than one registry that contains a pack with the same name.
	// The registry UID of the pack. The registry UID is the unique identifier of the registry. This attribute is required if there is more than one registry that contains a pack with the same name.
	// +kubebuilder:validation:Optional
	RegistryUID *string `json:"registryUid,omitempty" tf:"registry_uid,omitempty"`

	// (String) The tag of the pack. The tag is the version of the pack. This attribute is required if the pack type is spectro or helm.
	// The tag of the pack. The tag is the version of the pack. This attribute is required if the pack type is `spectro` or `helm`.
	// +kubebuilder:validation:Optional
	Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`

	// (String) The type of the pack. Allowed values are spectro, manifest or helm. The default value is spectro.
	// The type of the pack. Allowed values are `spectro`, `manifest` or `helm`. The default value is `spectro`.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// (String) The unique identifier of the pack. The value can be looked up using the spectrocloud_pack data source. This value is required if the pack type is spectro.
	// The unique identifier of the pack. The value can be looked up using the [`spectrocloud_pack`](https://registry.io/providers/spectrocloud/spectrocloud/latest/docs/data-sources/pack) data source. This value is required if the pack type is `spectro`.
	// +kubebuilder:validation:Optional
	UID *string `json:"uid,omitempty" tf:"uid,omitempty"`

	// (String)
	// The values of the pack. The values are the configuration values of the pack. The values are specified in YAML format.
	// +kubebuilder:validation:Optional
	Values *string `json:"values,omitempty" tf:"values,omitempty"`
}

type GroupClusterProfileParameters struct {

	// (String) The ID of this resource.
	// The ID of the cluster profile.
	// +kubebuilder:validation:Optional
	ID *string `json:"id" tf:"id,omitempty"`

	// (Block List) For packs of type spectro, helm, and manifest, at least one pack must be specified. (see below for nested schema)
	// For packs of type `spectro`, `helm`, and `manifest`, at least one pack must be specified.
	// +kubebuilder:validation:Optional
	Pack []GroupClusterProfilePackParameters `json:"pack,omitempty" tf:"pack,omitempty"`
}

type GroupInitParameters struct {

	// (Block List) (see below for nested schema)
	ClusterProfile []GroupClusterProfileInitParameters `json:"clusterProfile,omitempty" tf:"cluster_profile,omitempty"`

	// (Block List) A list of clusters to include in the cluster group. (see below for nested schema)
	// A list of clusters to include in the cluster group.
	Clusters []ClustersInitParameters `json:"clusters,omitempty" tf:"clusters,omitempty"`

	// (Block List, Min: 1, Max: 1) (see below for nested schema)
	Config []ConfigInitParameters `json:"config,omitempty" tf:"config,omitempty"`

	// (String) The context of the Cluster group. Allowed values are project or tenant. Defaults to tenant. If  the project context is specified, the project name will sourced from the provider configuration parameter project_name.
	// The context of the Cluster group. Allowed values are `project` or `tenant`. Defaults to `tenant`. If  the `project` context is specified, the project name will sourced from the provider configuration parameter [`project_name`](https://registry.io/providers/spectrocloud/spectrocloud/latest/docs#schema).
	Context *string `json:"context,omitempty" tf:"context,omitempty"`

	// (String) The description of the cluster. Default value is empty string.
	// The description of the cluster. Default value is empty string.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Set of String) A list of tags to be applied to the cluster group. Tags must be in the form of key:value.
	// A list of tags to be applied to the cluster group. Tags must be in the form of `key:value`.
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type GroupObservation struct {

	// (Block List) (see below for nested schema)
	ClusterProfile []GroupClusterProfileObservation `json:"clusterProfile,omitempty" tf:"cluster_profile,omitempty"`

	// (Block List) A list of clusters to include in the cluster group. (see below for nested schema)
	// A list of clusters to include in the cluster group.
	Clusters []ClustersObservation `json:"clusters,omitempty" tf:"clusters,omitempty"`

	// (Block List, Min: 1, Max: 1) (see below for nested schema)
	Config []ConfigObservation `json:"config,omitempty" tf:"config,omitempty"`

	// (String) The context of the Cluster group. Allowed values are project or tenant. Defaults to tenant. If  the project context is specified, the project name will sourced from the provider configuration parameter project_name.
	// The context of the Cluster group. Allowed values are `project` or `tenant`. Defaults to `tenant`. If  the `project` context is specified, the project name will sourced from the provider configuration parameter [`project_name`](https://registry.io/providers/spectrocloud/spectrocloud/latest/docs#schema).
	Context *string `json:"context,omitempty" tf:"context,omitempty"`

	// (String) The description of the cluster. Default value is empty string.
	// The description of the cluster. Default value is empty string.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Set of String) A list of tags to be applied to the cluster group. Tags must be in the form of key:value.
	// A list of tags to be applied to the cluster group. Tags must be in the form of `key:value`.
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type GroupParameters struct {

	// (Block List) (see below for nested schema)
	// +kubebuilder:validation:Optional
	ClusterProfile []GroupClusterProfileParameters `json:"clusterProfile,omitempty" tf:"cluster_profile,omitempty"`

	// (Block List) A list of clusters to include in the cluster group. (see below for nested schema)
	// A list of clusters to include in the cluster group.
	// +kubebuilder:validation:Optional
	Clusters []ClustersParameters `json:"clusters,omitempty" tf:"clusters,omitempty"`

	// (Block List, Min: 1, Max: 1) (see below for nested schema)
	// +kubebuilder:validation:Optional
	Config []ConfigParameters `json:"config,omitempty" tf:"config,omitempty"`

	// (String) The context of the Cluster group. Allowed values are project or tenant. Defaults to tenant. If  the project context is specified, the project name will sourced from the provider configuration parameter project_name.
	// The context of the Cluster group. Allowed values are `project` or `tenant`. Defaults to `tenant`. If  the `project` context is specified, the project name will sourced from the provider configuration parameter [`project_name`](https://registry.io/providers/spectrocloud/spectrocloud/latest/docs#schema).
	// +kubebuilder:validation:Optional
	Context *string `json:"context,omitempty" tf:"context,omitempty"`

	// (String) The description of the cluster. Default value is empty string.
	// The description of the cluster. Default value is empty string.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Set of String) A list of tags to be applied to the cluster group. Tags must be in the form of key:value.
	// A list of tags to be applied to the cluster group. Tags must be in the form of `key:value`.
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// GroupSpec defines the desired state of Group
type GroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GroupParameters `json:"forProvider"`
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
	InitProvider GroupInitParameters `json:"initProvider,omitempty"`
}

// GroupStatus defines the observed state of Group.
type GroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Group is the Schema for the Groups API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,palette}
type Group struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.config) || (has(self.initProvider) && has(self.initProvider.config))",message="spec.forProvider.config is a required parameter"
	Spec   GroupSpec   `json:"spec"`
	Status GroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GroupList contains a list of Groups
type GroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Group `json:"items"`
}

// Repository type metadata.
var (
	Group_Kind             = "Group"
	Group_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Group_Kind}.String()
	Group_KindAPIVersion   = Group_Kind + "." + CRDGroupVersion.String()
	Group_GroupVersionKind = CRDGroupVersion.WithKind(Group_Kind)
)

func init() {
	SchemeBuilder.Register(&Group{}, &GroupList{})
}
