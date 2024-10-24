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

type SubaccountTrustConfigurationInitParameters struct {

	// (Boolean) Determines that any user from the tenant can log in. If not set, only the ones who already have a shadow user can log in.
	// Determines that any user from the tenant can log in. If not set, only the ones who already have a shadow user can log in.
	AutoCreateShadowUsers *bool `json:"autoCreateShadowUsers,omitempty" tf:"auto_create_shadow_users,omitempty"`

	// (Boolean) Determines that end users can choose the trust configuration for login. If not set, the trust configuration can remain active, however only application users that explicitly specify the origin key can use if for login.
	// Determines that end users can choose the trust configuration for login. If not set, the trust configuration can remain active, however only application users that explicitly specify the origin key can use if for login.
	AvailableForUserLogon *bool `json:"availableForUserLogon,omitempty" tf:"available_for_user_logon,omitempty"`

	// (String) Description of the trust configuration.
	// Description of the trust configuration.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The tenant's domain which should be used for user logon.
	// The tenant's domain which should be used for user logon.
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// (String) The name of the Identity Authentication tenant that you want to connect to the subaccount.
	// The name of the Identity Authentication tenant that you want to connect to the subaccount.
	IdentityProvider *string `json:"identityProvider,omitempty" tf:"identity_provider,omitempty"`

	// (String) Short string that helps users to identify the link for login.
	// Short string that helps users to identify the link for login.
	LinkText *string `json:"linkText,omitempty" tf:"link_text,omitempty"`

	// (String) The display name of the trust configuration.
	// The display name of the trust configuration.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) Determines whether the identity provider is currently 'active' or 'inactive'.
	// Determines whether the identity provider is currently 'active' or 'inactive'.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// (String) The ID of the subaccount.
	// The ID of the subaccount.
	// +crossplane:generate:reference:type=github.com/sap/crossplane-provider-btp/apis/account/v1alpha1.Subaccount
	// +crossplane:generate:reference:extractor=github.com/sap/crossplane-provider-btp/apis/account/v1alpha1.SubaccountUuid()
	// +crossplane:generate:reference:refFieldName=SubaccountRef
	// +crossplane:generate:reference:selectorFieldName=SubaccountSelector
	SubaccountID *string `json:"subaccountId,omitempty" tf:"subaccount_id,omitempty"`

	// Reference to a Subaccount in account to populate subaccountId.
	// +kubebuilder:validation:Optional
	SubaccountRef *v1.Reference `json:"subaccountRef,omitempty" tf:"-"`

	// Selector for a Subaccount in account to populate subaccountId.
	// +kubebuilder:validation:Optional
	SubaccountSelector *v1.Selector `json:"subaccountSelector,omitempty" tf:"-"`
}

type SubaccountTrustConfigurationObservation struct {

	// (Boolean) Determines that any user from the tenant can log in. If not set, only the ones who already have a shadow user can log in.
	// Determines that any user from the tenant can log in. If not set, only the ones who already have a shadow user can log in.
	AutoCreateShadowUsers *bool `json:"autoCreateShadowUsers,omitempty" tf:"auto_create_shadow_users,omitempty"`

	// (Boolean) Determines that end users can choose the trust configuration for login. If not set, the trust configuration can remain active, however only application users that explicitly specify the origin key can use if for login.
	// Determines that end users can choose the trust configuration for login. If not set, the trust configuration can remain active, however only application users that explicitly specify the origin key can use if for login.
	AvailableForUserLogon *bool `json:"availableForUserLogon,omitempty" tf:"available_for_user_logon,omitempty"`

	// (String) Description of the trust configuration.
	// Description of the trust configuration.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The tenant's domain which should be used for user logon.
	// The tenant's domain which should be used for user logon.
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// (String, Deprecated) The origin of the identity provider.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The name of the Identity Authentication tenant that you want to connect to the subaccount.
	// The name of the Identity Authentication tenant that you want to connect to the subaccount.
	IdentityProvider *string `json:"identityProvider,omitempty" tf:"identity_provider,omitempty"`

	// (String) Short string that helps users to identify the link for login.
	// Short string that helps users to identify the link for login.
	LinkText *string `json:"linkText,omitempty" tf:"link_text,omitempty"`

	// (String) The display name of the trust configuration.
	// The display name of the trust configuration.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The origin of the identity provider.
	// The origin of the identity provider.
	Origin *string `json:"origin,omitempty" tf:"origin,omitempty"`

	// (String) The protocol used to establish trust with the identity provider.
	// The protocol used to establish trust with the identity provider.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// (Boolean) Shows whether the trust configuration can be modified.
	// Shows whether the trust configuration can be modified.
	ReadOnly *bool `json:"readOnly,omitempty" tf:"read_only,omitempty"`

	// (String) Determines whether the identity provider is currently 'active' or 'inactive'.
	// Determines whether the identity provider is currently 'active' or 'inactive'.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// (String) The ID of the subaccount.
	// The ID of the subaccount.
	SubaccountID *string `json:"subaccountId,omitempty" tf:"subaccount_id,omitempty"`

	// (String) The trust type.
	// The trust type.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type SubaccountTrustConfigurationParameters struct {

	// (Boolean) Determines that any user from the tenant can log in. If not set, only the ones who already have a shadow user can log in.
	// Determines that any user from the tenant can log in. If not set, only the ones who already have a shadow user can log in.
	// +kubebuilder:validation:Optional
	AutoCreateShadowUsers *bool `json:"autoCreateShadowUsers,omitempty" tf:"auto_create_shadow_users,omitempty"`

	// (Boolean) Determines that end users can choose the trust configuration for login. If not set, the trust configuration can remain active, however only application users that explicitly specify the origin key can use if for login.
	// Determines that end users can choose the trust configuration for login. If not set, the trust configuration can remain active, however only application users that explicitly specify the origin key can use if for login.
	// +kubebuilder:validation:Optional
	AvailableForUserLogon *bool `json:"availableForUserLogon,omitempty" tf:"available_for_user_logon,omitempty"`

	// (String) Description of the trust configuration.
	// Description of the trust configuration.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The tenant's domain which should be used for user logon.
	// The tenant's domain which should be used for user logon.
	// +kubebuilder:validation:Optional
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// (String) The name of the Identity Authentication tenant that you want to connect to the subaccount.
	// The name of the Identity Authentication tenant that you want to connect to the subaccount.
	// +kubebuilder:validation:Optional
	IdentityProvider *string `json:"identityProvider,omitempty" tf:"identity_provider,omitempty"`

	// (String) Short string that helps users to identify the link for login.
	// Short string that helps users to identify the link for login.
	// +kubebuilder:validation:Optional
	LinkText *string `json:"linkText,omitempty" tf:"link_text,omitempty"`

	// (String) The display name of the trust configuration.
	// The display name of the trust configuration.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) Determines whether the identity provider is currently 'active' or 'inactive'.
	// Determines whether the identity provider is currently 'active' or 'inactive'.
	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// (String) The ID of the subaccount.
	// The ID of the subaccount.
	// +crossplane:generate:reference:type=github.com/sap/crossplane-provider-btp/apis/account/v1alpha1.Subaccount
	// +crossplane:generate:reference:extractor=github.com/sap/crossplane-provider-btp/apis/account/v1alpha1.SubaccountUuid()
	// +crossplane:generate:reference:refFieldName=SubaccountRef
	// +crossplane:generate:reference:selectorFieldName=SubaccountSelector
	// +kubebuilder:validation:Optional
	SubaccountID *string `json:"subaccountId,omitempty" tf:"subaccount_id,omitempty"`

	// Reference to a Subaccount in account to populate subaccountId.
	// +kubebuilder:validation:Optional
	SubaccountRef *v1.Reference `json:"subaccountRef,omitempty" tf:"-"`

	// Selector for a Subaccount in account to populate subaccountId.
	// +kubebuilder:validation:Optional
	SubaccountSelector *v1.Selector `json:"subaccountSelector,omitempty" tf:"-"`
}

// SubaccountTrustConfigurationSpec defines the desired state of SubaccountTrustConfiguration
type SubaccountTrustConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SubaccountTrustConfigurationParameters `json:"forProvider"`
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
	InitProvider SubaccountTrustConfigurationInitParameters `json:"initProvider,omitempty"`
}

// SubaccountTrustConfigurationStatus defines the observed state of SubaccountTrustConfiguration.
type SubaccountTrustConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SubaccountTrustConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SubaccountTrustConfiguration is the Schema for the SubaccountTrustConfigurations API. Establishes trust from a subaccount to an Identity Authentication tenant. Tip: You must be assigned to the admin role of the subaccount. Further documentation: https://help.sap.com/docs/btp/sap-business-technology-platform/trust-and-federation-with-identity-providers
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,account}
type SubaccountTrustConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.identityProvider) || (has(self.initProvider) && has(self.initProvider.identityProvider))",message="spec.forProvider.identityProvider is a required parameter"
	Spec   SubaccountTrustConfigurationSpec   `json:"spec"`
	Status SubaccountTrustConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SubaccountTrustConfigurationList contains a list of SubaccountTrustConfigurations
type SubaccountTrustConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SubaccountTrustConfiguration `json:"items"`
}

// Repository type metadata.
var (
	SubaccountTrustConfiguration_Kind             = "SubaccountTrustConfiguration"
	SubaccountTrustConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SubaccountTrustConfiguration_Kind}.String()
	SubaccountTrustConfiguration_KindAPIVersion   = SubaccountTrustConfiguration_Kind + "." + CRDGroupVersion.String()
	SubaccountTrustConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(SubaccountTrustConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&SubaccountTrustConfiguration{}, &SubaccountTrustConfigurationList{})
}
