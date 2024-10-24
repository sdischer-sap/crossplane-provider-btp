/*
Copyright 2022 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type DirectoryEntitlementInitParameters struct {

	// (Number) The quota assigned to the directory.
	// The quota assigned to the directory.
	Amount *float64 `json:"amount,omitempty" tf:"amount,omitempty"`

	// (Boolean) Determines whether the plans of entitlements that have a numeric quota with the amount specified in auto_distribute_amount are automatically allocated to any new subaccount that is added to the directory in the future. For entitlements without a numeric quota, it shows if the plan are assigned to any new subaccount that is added to the directory in the future (auto_distribute_amount is not needed). If the distribute parameter is set, the same assignment is also made to all subaccounts currently in the directory. Entitlements are subject to available quota in the directory.
	// Determines whether the plans of entitlements that have a numeric quota with the amount specified in `auto_distribute_amount` are automatically allocated to any new subaccount that is added to the directory in the future. For entitlements without a numeric quota, it shows if the plan are assigned to any new subaccount that is added to the directory in the future (`auto_distribute_amount` is not needed). If the `distribute` parameter is set, the same assignment is also made to all subaccounts currently in the directory. Entitlements are subject to available quota in the directory.
	AutoAssign *bool `json:"autoAssign,omitempty" tf:"auto_assign,omitempty"`

	// (Number) The quota of the specified plan automatically allocated to any new subaccount that is created in the future in the directory. When applying this option, auto_assign and/or distribute must also be set. Applies only to entitlements that have a numeric quota.
	// The quota of the specified plan automatically allocated to any new subaccount that is created in the future in the directory. When applying this option, `auto_assign` and/or `distribute` must also be set. Applies only to entitlements that have a numeric quota.
	AutoDistributeAmount *float64 `json:"autoDistributeAmount,omitempty" tf:"auto_distribute_amount,omitempty"`

	// (String) The ID of the directory.
	// The ID of the directory.
	// +crossplane:generate:reference:type=github.com/sap/crossplane-provider-btp/apis/account/v1alpha1.Directory
	// +crossplane:generate:reference:extractor=github.com/sap/crossplane-provider-btp/apis/account/v1alpha1.DirectoryUuid()
	// +crossplane:generate:reference:refFieldName=DirectoryRef
	// +crossplane:generate:reference:selectorFieldName=DirectorySelector
	DirectoryID *string `json:"directoryId,omitempty" tf:"directory_id,omitempty"`

	// Reference to a Directory in account to populate directoryId.
	// +kubebuilder:validation:Optional
	DirectoryRef *v1.Reference `json:"directoryRef,omitempty" tf:"-"`

	// Selector for a Directory in account to populate directoryId.
	// +kubebuilder:validation:Optional
	DirectorySelector *v1.Selector `json:"directorySelector,omitempty" tf:"-"`

	// (Boolean) Defines the assignment of the plan with the quota specified in auto_distribute_amount to subaccounts currently located in the specified directory. For entitlements without a numeric quota, the plan is assigned to the subaccounts currently located in the directory (auto_distribute_amount is not needed). When applying this option, auto_assign must also be set.
	// Defines the assignment of the plan with the quota specified in `auto_distribute_amount` to subaccounts currently located in the specified directory. For entitlements without a numeric quota, the plan is assigned to the subaccounts currently located in the directory (`auto_distribute_amount` is not needed). When applying this option, `auto_assign` must also be set.
	Distribute *bool `json:"distribute,omitempty" tf:"distribute,omitempty"`

	// (String) The name of the entitled service plan.
	// The name of the entitled service plan.
	PlanName *string `json:"planName,omitempty" tf:"plan_name,omitempty"`

	// (String) The name of the entitled service.
	// The name of the entitled service.
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`
}

type DirectoryEntitlementObservation struct {

	// (Number) The quota assigned to the directory.
	// The quota assigned to the directory.
	Amount *float64 `json:"amount,omitempty" tf:"amount,omitempty"`

	// (Boolean) Determines whether the plans of entitlements that have a numeric quota with the amount specified in auto_distribute_amount are automatically allocated to any new subaccount that is added to the directory in the future. For entitlements without a numeric quota, it shows if the plan are assigned to any new subaccount that is added to the directory in the future (auto_distribute_amount is not needed). If the distribute parameter is set, the same assignment is also made to all subaccounts currently in the directory. Entitlements are subject to available quota in the directory.
	// Determines whether the plans of entitlements that have a numeric quota with the amount specified in `auto_distribute_amount` are automatically allocated to any new subaccount that is added to the directory in the future. For entitlements without a numeric quota, it shows if the plan are assigned to any new subaccount that is added to the directory in the future (`auto_distribute_amount` is not needed). If the `distribute` parameter is set, the same assignment is also made to all subaccounts currently in the directory. Entitlements are subject to available quota in the directory.
	AutoAssign *bool `json:"autoAssign,omitempty" tf:"auto_assign,omitempty"`

	// (Number) The quota of the specified plan automatically allocated to any new subaccount that is created in the future in the directory. When applying this option, auto_assign and/or distribute must also be set. Applies only to entitlements that have a numeric quota.
	// The quota of the specified plan automatically allocated to any new subaccount that is created in the future in the directory. When applying this option, `auto_assign` and/or `distribute` must also be set. Applies only to entitlements that have a numeric quota.
	AutoDistributeAmount *float64 `json:"autoDistributeAmount,omitempty" tf:"auto_distribute_amount,omitempty"`

	// (String) The current state of the entitlement. Possible values are:
	// The current state of the entitlement. Possible values are:
	//
	// | value | description |
	// | --- | --- |
	// | `PLATFORM` |  A service required for using a specific platform; for example, Application Runtime is required for the Cloud Foundry platform. |
	// | `SERVICE` | A commercial or technical service. that has a numeric quota (amount) when entitled or assigned to a resource. When assigning entitlements of this type, use the 'amount' option. |
	// | `ELASTIC_SERVICE` | A commercial or technical service that has no numeric quota (amount) when entitled or assigned to a resource. Generally this type of service can be as many times as needed when enabled, but may in some cases be restricted by the service owner. |
	// | `ELASTIC_LIMITED` | An elastic service that can be enabled for only one subaccount per global account. |
	// | `APPLICATION` | A multitenant application to which consumers can subscribe. As opposed to applications defined as a 'QUOTA_BASED_APPLICATION', these applications do not have a numeric quota and are simply enabled or disabled as entitlements per subaccount. |
	// | `QUOTA_BASED_APPLICATION` | A multitenant application to which consumers can subscribe. As opposed to applications defined as 'APPLICATION', these applications have an numeric quota that limits consumer usage of the subscribed application per subaccount. |
	// | `ENVIRONMENT` |  An environment service; for example, Cloud Foundry. |
	Category *string `json:"category,omitempty" tf:"category,omitempty"`

	// (String) The ID of the directory.
	// The ID of the directory.
	DirectoryID *string `json:"directoryId,omitempty" tf:"directory_id,omitempty"`

	// (Boolean) Defines the assignment of the plan with the quota specified in auto_distribute_amount to subaccounts currently located in the specified directory. For entitlements without a numeric quota, the plan is assigned to the subaccounts currently located in the directory (auto_distribute_amount is not needed). When applying this option, auto_assign must also be set.
	// Defines the assignment of the plan with the quota specified in `auto_distribute_amount` to subaccounts currently located in the specified directory. For entitlements without a numeric quota, the plan is assigned to the subaccounts currently located in the directory (`auto_distribute_amount` is not needed). When applying this option, `auto_assign` must also be set.
	Distribute *bool `json:"distribute,omitempty" tf:"distribute,omitempty"`

	// (String) The ID of the entitled service plan.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The ID of the entitled service plan.
	// The ID of the entitled service plan.
	PlanID *string `json:"planId,omitempty" tf:"plan_id,omitempty"`

	// (String) The name of the entitled service plan.
	// The name of the entitled service plan.
	PlanName *string `json:"planName,omitempty" tf:"plan_name,omitempty"`

	// (String) The name of the entitled service.
	// The name of the entitled service.
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`
}

type DirectoryEntitlementParameters struct {

	// (Number) The quota assigned to the directory.
	// The quota assigned to the directory.
	// +kubebuilder:validation:Optional
	Amount *float64 `json:"amount,omitempty" tf:"amount,omitempty"`

	// (Boolean) Determines whether the plans of entitlements that have a numeric quota with the amount specified in auto_distribute_amount are automatically allocated to any new subaccount that is added to the directory in the future. For entitlements without a numeric quota, it shows if the plan are assigned to any new subaccount that is added to the directory in the future (auto_distribute_amount is not needed). If the distribute parameter is set, the same assignment is also made to all subaccounts currently in the directory. Entitlements are subject to available quota in the directory.
	// Determines whether the plans of entitlements that have a numeric quota with the amount specified in `auto_distribute_amount` are automatically allocated to any new subaccount that is added to the directory in the future. For entitlements without a numeric quota, it shows if the plan are assigned to any new subaccount that is added to the directory in the future (`auto_distribute_amount` is not needed). If the `distribute` parameter is set, the same assignment is also made to all subaccounts currently in the directory. Entitlements are subject to available quota in the directory.
	// +kubebuilder:validation:Optional
	AutoAssign *bool `json:"autoAssign,omitempty" tf:"auto_assign,omitempty"`

	// (Number) The quota of the specified plan automatically allocated to any new subaccount that is created in the future in the directory. When applying this option, auto_assign and/or distribute must also be set. Applies only to entitlements that have a numeric quota.
	// The quota of the specified plan automatically allocated to any new subaccount that is created in the future in the directory. When applying this option, `auto_assign` and/or `distribute` must also be set. Applies only to entitlements that have a numeric quota.
	// +kubebuilder:validation:Optional
	AutoDistributeAmount *float64 `json:"autoDistributeAmount,omitempty" tf:"auto_distribute_amount,omitempty"`

	// (String) The ID of the directory.
	// The ID of the directory.
	// +crossplane:generate:reference:type=github.com/sap/crossplane-provider-btp/apis/account/v1alpha1.Directory
	// +crossplane:generate:reference:extractor=github.com/sap/crossplane-provider-btp/apis/account/v1alpha1.DirectoryUuid()
	// +crossplane:generate:reference:refFieldName=DirectoryRef
	// +crossplane:generate:reference:selectorFieldName=DirectorySelector
	// +kubebuilder:validation:Optional
	DirectoryID *string `json:"directoryId,omitempty" tf:"directory_id,omitempty"`

	// Reference to a Directory in account to populate directoryId.
	// +kubebuilder:validation:Optional
	DirectoryRef *v1.Reference `json:"directoryRef,omitempty" tf:"-"`

	// Selector for a Directory in account to populate directoryId.
	// +kubebuilder:validation:Optional
	DirectorySelector *v1.Selector `json:"directorySelector,omitempty" tf:"-"`

	// (Boolean) Defines the assignment of the plan with the quota specified in auto_distribute_amount to subaccounts currently located in the specified directory. For entitlements without a numeric quota, the plan is assigned to the subaccounts currently located in the directory (auto_distribute_amount is not needed). When applying this option, auto_assign must also be set.
	// Defines the assignment of the plan with the quota specified in `auto_distribute_amount` to subaccounts currently located in the specified directory. For entitlements without a numeric quota, the plan is assigned to the subaccounts currently located in the directory (`auto_distribute_amount` is not needed). When applying this option, `auto_assign` must also be set.
	// +kubebuilder:validation:Optional
	Distribute *bool `json:"distribute,omitempty" tf:"distribute,omitempty"`

	// (String) The name of the entitled service plan.
	// The name of the entitled service plan.
	// +kubebuilder:validation:Optional
	PlanName *string `json:"planName,omitempty" tf:"plan_name,omitempty"`

	// (String) The name of the entitled service.
	// The name of the entitled service.
	// +kubebuilder:validation:Optional
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`
}

// DirectoryEntitlementSpec defines the desired state of DirectoryEntitlement
type DirectoryEntitlementSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DirectoryEntitlementParameters `json:"forProvider"`
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
	InitProvider DirectoryEntitlementInitParameters `json:"initProvider,omitempty"`
}

// DirectoryEntitlementStatus defines the observed state of DirectoryEntitlement.
type DirectoryEntitlementStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DirectoryEntitlementObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// DirectoryEntitlement is the Schema for the DirectoryEntitlements API. Assigns the entitlement plan of a service, multitenant application, or environment, to a directory. Note that some environments, such as Cloud Foundry, are available by default to all global accounts and their directorys, and therefore are not made available as entitlements. Tip: You must be assigned to the admin role of the global account or the directory. Further documentation: https://help.sap.com/docs/btp/sap-business-technology-platform/entitlements-and-quotas
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,account}
type DirectoryEntitlement struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.planName) || (has(self.initProvider) && has(self.initProvider.planName))",message="spec.forProvider.planName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.serviceName) || (has(self.initProvider) && has(self.initProvider.serviceName))",message="spec.forProvider.serviceName is a required parameter"
	Spec   DirectoryEntitlementSpec   `json:"spec"`
	Status DirectoryEntitlementStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DirectoryEntitlementList contains a list of DirectoryEntitlements
type DirectoryEntitlementList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DirectoryEntitlement `json:"items"`
}

// Repository type metadata.
var (
	DirectoryEntitlement_Kind             = "DirectoryEntitlement"
	DirectoryEntitlement_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DirectoryEntitlement_Kind}.String()
	DirectoryEntitlement_KindAPIVersion   = DirectoryEntitlement_Kind + "." + CRDGroupVersion.String()
	DirectoryEntitlement_GroupVersionKind = CRDGroupVersion.WithKind(DirectoryEntitlement_Kind)
)

func init() {
	SchemeBuilder.Register(&DirectoryEntitlement{}, &DirectoryEntitlementList{})
}
