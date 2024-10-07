/*
Accounts Service

The Accounts service provides REST APIs that are responsible for the management of global accounts, and the creation and management of directories, subaccounts, and their custom properties/tags.  Global accounts represent a business entity and contain contract information, including customer details and purchased entitlements. The global account is the context for billing each customer.  Use the subaccount APIs to structure your global account according to your organization's and project's requirements regarding members, authorizations, and quotas. This service also provides you with APIs for creating and managing directories. While the use of directories is optional, they allow you to further organize and manage your subaccounts according to your specific technical and business needs. The service also lets you manage the custom properties/tags that you associate with your directories and subaccounts. NOTE: These APIs are relevant only for cloud management tools feature set B. For details and information about whether this applies to your global account, see [Cloud Management Tools - Feature Set Overview](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/Cloud/en-US/caf4e4e23aef4666ad8f125af393dfb2.html).  See also: * [Authorization](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/latest/en-US/3670474a58c24ac2b082e76cbbd9dc19.html) * [Rate Limiting](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/latest/en-US/77b217b3f57a45b987eb7fbc3305ce1e.html) * [Error Response Format](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/latest/en-US/77fef2fb104b4b1795e2e6cee790e8b8.html) * [Asynchronous Jobs](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/latest/en-US/0a0a6ab0ad114d72a6611c1c6b21683e.html)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the UpdateDirectoryTypeRequestPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateDirectoryTypeRequestPayload{}

// UpdateDirectoryTypeRequestPayload The request payload containing the list of features to be enabled in the directory.
type UpdateDirectoryTypeRequestPayload struct {
	// Additional admins of the directory. Applies only to directories that have the user authorization management feature enabled. Do not add yourself as you are assigned as a directory admin by default. Example: [\"admin1@example.com\", \"admin2@example.com\"]
	DirectoryAdmins []string `json:"directoryAdmins,omitempty"`
	// <b>The features to be enabled in the directory. The available features are:</b> - <b>DEFAULT</b>: (Mandatory) All directories provide the following basic features: (1) Group and filter subaccounts for reports and filters, (2) monitor usage and costs on a directory level (costs only available for contracts that use the consumption-based commercial model), and (3) set custom properties and tags to the directory for identification and reporting purposes. - <b>ENTITLEMENTS</b>: (Optional) Enables the assignment of a quota for services and applications to the directory from the global account quota for distribution to the subaccounts under this directory.  - <b>AUTHORIZATIONS</b>: (Optional) Allows you to assign users as administrators or viewers of this directory. You must apply this feature in combination with the ENTITLEMENTS feature.   IMPORTANT: Your multi-level account hierarchy can have more than one directory enabled with user authorization and/or entitlement management; however, only one directory in any directory path can have these features enabled. In other words, other directories above or below this directory in the same path can only have the default features specified. If you are not sure which features to enable, we recommend that you set only the default features, and then add features later on as they are needed.  <br/><b>Valid values:</b>  [DEFAULT] [DEFAULT,ENTITLEMENTS] [DEFAULT,ENTITLEMENTS,AUTHORIZATIONS]<br/>
	DirectoryFeatures []string `json:"directoryFeatures,omitempty"`
	// Applies only to directories that have the user authorization management feature enabled.  The subdomain becomes part of the path used to access the authorization tenant of the directory. Must be unique within the defined region. Use only letters (a-z), digits (0-9), and hyphens (not at start or end). Maximum length is 63 characters. Cannot be changed after the directory has been created.
	Subdomain *string `json:"subdomain,omitempty"`
}

// NewUpdateDirectoryTypeRequestPayload instantiates a new UpdateDirectoryTypeRequestPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDirectoryTypeRequestPayload() *UpdateDirectoryTypeRequestPayload {
	this := UpdateDirectoryTypeRequestPayload{}
	return &this
}

// NewUpdateDirectoryTypeRequestPayloadWithDefaults instantiates a new UpdateDirectoryTypeRequestPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDirectoryTypeRequestPayloadWithDefaults() *UpdateDirectoryTypeRequestPayload {
	this := UpdateDirectoryTypeRequestPayload{}
	return &this
}

// GetDirectoryAdmins returns the DirectoryAdmins field value if set, zero value otherwise.
func (o *UpdateDirectoryTypeRequestPayload) GetDirectoryAdmins() []string {
	if o == nil || IsNil(o.DirectoryAdmins) {
		var ret []string
		return ret
	}
	return o.DirectoryAdmins
}

// GetDirectoryAdminsOk returns a tuple with the DirectoryAdmins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDirectoryTypeRequestPayload) GetDirectoryAdminsOk() ([]string, bool) {
	if o == nil || IsNil(o.DirectoryAdmins) {
		return nil, false
	}
	return o.DirectoryAdmins, true
}

// HasDirectoryAdmins returns a boolean if a field has been set.
func (o *UpdateDirectoryTypeRequestPayload) HasDirectoryAdmins() bool {
	if o != nil && !IsNil(o.DirectoryAdmins) {
		return true
	}

	return false
}

// SetDirectoryAdmins gets a reference to the given []string and assigns it to the DirectoryAdmins field.
func (o *UpdateDirectoryTypeRequestPayload) SetDirectoryAdmins(v []string) {
	o.DirectoryAdmins = v
}

// GetDirectoryFeatures returns the DirectoryFeatures field value if set, zero value otherwise.
func (o *UpdateDirectoryTypeRequestPayload) GetDirectoryFeatures() []string {
	if o == nil || IsNil(o.DirectoryFeatures) {
		var ret []string
		return ret
	}
	return o.DirectoryFeatures
}

// GetDirectoryFeaturesOk returns a tuple with the DirectoryFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDirectoryTypeRequestPayload) GetDirectoryFeaturesOk() ([]string, bool) {
	if o == nil || IsNil(o.DirectoryFeatures) {
		return nil, false
	}
	return o.DirectoryFeatures, true
}

// HasDirectoryFeatures returns a boolean if a field has been set.
func (o *UpdateDirectoryTypeRequestPayload) HasDirectoryFeatures() bool {
	if o != nil && !IsNil(o.DirectoryFeatures) {
		return true
	}

	return false
}

// SetDirectoryFeatures gets a reference to the given []string and assigns it to the DirectoryFeatures field.
func (o *UpdateDirectoryTypeRequestPayload) SetDirectoryFeatures(v []string) {
	o.DirectoryFeatures = v
}

// GetSubdomain returns the Subdomain field value if set, zero value otherwise.
func (o *UpdateDirectoryTypeRequestPayload) GetSubdomain() string {
	if o == nil || IsNil(o.Subdomain) {
		var ret string
		return ret
	}
	return *o.Subdomain
}

// GetSubdomainOk returns a tuple with the Subdomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDirectoryTypeRequestPayload) GetSubdomainOk() (*string, bool) {
	if o == nil || IsNil(o.Subdomain) {
		return nil, false
	}
	return o.Subdomain, true
}

// HasSubdomain returns a boolean if a field has been set.
func (o *UpdateDirectoryTypeRequestPayload) HasSubdomain() bool {
	if o != nil && !IsNil(o.Subdomain) {
		return true
	}

	return false
}

// SetSubdomain gets a reference to the given string and assigns it to the Subdomain field.
func (o *UpdateDirectoryTypeRequestPayload) SetSubdomain(v string) {
	o.Subdomain = &v
}

func (o UpdateDirectoryTypeRequestPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateDirectoryTypeRequestPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DirectoryAdmins) {
		toSerialize["directoryAdmins"] = o.DirectoryAdmins
	}
	if !IsNil(o.DirectoryFeatures) {
		toSerialize["directoryFeatures"] = o.DirectoryFeatures
	}
	if !IsNil(o.Subdomain) {
		toSerialize["subdomain"] = o.Subdomain
	}
	return toSerialize, nil
}

type NullableUpdateDirectoryTypeRequestPayload struct {
	value *UpdateDirectoryTypeRequestPayload
	isSet bool
}

func (v NullableUpdateDirectoryTypeRequestPayload) Get() *UpdateDirectoryTypeRequestPayload {
	return v.value
}

func (v *NullableUpdateDirectoryTypeRequestPayload) Set(val *UpdateDirectoryTypeRequestPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDirectoryTypeRequestPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDirectoryTypeRequestPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDirectoryTypeRequestPayload(val *UpdateDirectoryTypeRequestPayload) *NullableUpdateDirectoryTypeRequestPayload {
	return &NullableUpdateDirectoryTypeRequestPayload{value: val, isSet: true}
}

func (v NullableUpdateDirectoryTypeRequestPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDirectoryTypeRequestPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

