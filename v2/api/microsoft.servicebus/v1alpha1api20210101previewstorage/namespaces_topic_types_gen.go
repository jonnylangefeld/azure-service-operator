// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210101previewstorage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Storage version of v1alpha1api20210101preview.NamespacesTopic
//Generated from: https://schema.management.azure.com/schemas/2021-01-01-preview/Microsoft.ServiceBus.json#/resourceDefinitions/namespaces_topics
type NamespacesTopic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NamespacesTopics_Spec `json:"spec,omitempty"`
	Status            SBTopic_Status        `json:"status,omitempty"`
}

var _ conditions.Conditioner = &NamespacesTopic{}

// GetConditions returns the conditions of the resource
func (namespacesTopic *NamespacesTopic) GetConditions() conditions.Conditions {
	return namespacesTopic.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (namespacesTopic *NamespacesTopic) SetConditions(conditions conditions.Conditions) {
	namespacesTopic.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &NamespacesTopic{}

// AzureName returns the Azure name of the resource
func (namespacesTopic *NamespacesTopic) AzureName() string {
	return namespacesTopic.Spec.AzureName
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (namespacesTopic *NamespacesTopic) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(namespacesTopic.Spec)
	return &genruntime.ResourceReference{Group: group, Kind: kind, Namespace: namespacesTopic.Namespace, Name: namespacesTopic.Spec.Owner.Name}
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (namespacesTopic *NamespacesTopic) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: namespacesTopic.Spec.OriginalVersion,
		Kind:    "NamespacesTopic",
	}
}

// +kubebuilder:object:root=true
//Storage version of v1alpha1api20210101preview.NamespacesTopic
//Generated from: https://schema.management.azure.com/schemas/2021-01-01-preview/Microsoft.ServiceBus.json#/resourceDefinitions/namespaces_topics
type NamespacesTopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NamespacesTopic `json:"items"`
}

//Storage version of v1alpha1api20210101preview.NamespacesTopics_Spec
type NamespacesTopics_Spec struct {
	AutoDeleteOnIdle *string `json:"autoDeleteOnIdle,omitempty"`

	// +kubebuilder:validation:MinLength=1
	//AzureName: The name of the resource in Azure. This is often the same as the name
	//of the resource in Kubernetes but it doesn't have to be.
	AzureName                           string  `json:"azureName"`
	DefaultMessageTimeToLive            *string `json:"defaultMessageTimeToLive,omitempty"`
	DuplicateDetectionHistoryTimeWindow *string `json:"duplicateDetectionHistoryTimeWindow,omitempty"`
	EnableBatchedOperations             *bool   `json:"enableBatchedOperations,omitempty"`
	EnableExpress                       *bool   `json:"enableExpress,omitempty"`
	EnablePartitioning                  *bool   `json:"enablePartitioning,omitempty"`
	Location                            *string `json:"location,omitempty"`
	MaxSizeInMegabytes                  *int    `json:"maxSizeInMegabytes,omitempty"`
	OriginalVersion                     string  `json:"originalVersion"`

	// +kubebuilder:validation:Required
	Owner                      genruntime.KnownResourceReference `group:"microsoft.servicebus.azure.com" json:"owner" kind:"Namespace"`
	PropertyBag                genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	RequiresDuplicateDetection *bool                             `json:"requiresDuplicateDetection,omitempty"`
	SupportOrdering            *bool                             `json:"supportOrdering,omitempty"`
	Tags                       map[string]string                 `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &NamespacesTopics_Spec{}

// ConvertSpecFrom populates our NamespacesTopics_Spec from the provided source
func (namespacesTopicsSpec *NamespacesTopics_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == namespacesTopicsSpec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(namespacesTopicsSpec)
}

// ConvertSpecTo populates the provided destination from our NamespacesTopics_Spec
func (namespacesTopicsSpec *NamespacesTopics_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == namespacesTopicsSpec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(namespacesTopicsSpec)
}

//Storage version of v1alpha1api20210101preview.SBTopic_Status
//Generated from:
type SBTopic_Status struct {
	AccessedAt                          *string                     `json:"accessedAt,omitempty"`
	AutoDeleteOnIdle                    *string                     `json:"autoDeleteOnIdle,omitempty"`
	Conditions                          []conditions.Condition      `json:"conditions,omitempty"`
	CountDetails                        *MessageCountDetails_Status `json:"countDetails,omitempty"`
	CreatedAt                           *string                     `json:"createdAt,omitempty"`
	DefaultMessageTimeToLive            *string                     `json:"defaultMessageTimeToLive,omitempty"`
	DuplicateDetectionHistoryTimeWindow *string                     `json:"duplicateDetectionHistoryTimeWindow,omitempty"`
	EnableBatchedOperations             *bool                       `json:"enableBatchedOperations,omitempty"`
	EnableExpress                       *bool                       `json:"enableExpress,omitempty"`
	EnablePartitioning                  *bool                       `json:"enablePartitioning,omitempty"`
	Id                                  *string                     `json:"id,omitempty"`
	MaxSizeInMegabytes                  *int                        `json:"maxSizeInMegabytes,omitempty"`
	Name                                *string                     `json:"name,omitempty"`
	PropertyBag                         genruntime.PropertyBag      `json:"$propertyBag,omitempty"`
	RequiresDuplicateDetection          *bool                       `json:"requiresDuplicateDetection,omitempty"`
	SizeInBytes                         *int                        `json:"sizeInBytes,omitempty"`
	Status                              *string                     `json:"status,omitempty"`
	SubscriptionCount                   *int                        `json:"subscriptionCount,omitempty"`
	SupportOrdering                     *bool                       `json:"supportOrdering,omitempty"`
	SystemData                          *SystemData_Status          `json:"systemData,omitempty"`
	Type                                *string                     `json:"type,omitempty"`
	UpdatedAt                           *string                     `json:"updatedAt,omitempty"`
}

var _ genruntime.ConvertibleStatus = &SBTopic_Status{}

// ConvertStatusFrom populates our SBTopic_Status from the provided source
func (sbTopicStatus *SBTopic_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == sbTopicStatus {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(sbTopicStatus)
}

// ConvertStatusTo populates the provided destination from our SBTopic_Status
func (sbTopicStatus *SBTopic_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == sbTopicStatus {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(sbTopicStatus)
}

func init() {
	SchemeBuilder.Register(&NamespacesTopic{}, &NamespacesTopicList{})
}
