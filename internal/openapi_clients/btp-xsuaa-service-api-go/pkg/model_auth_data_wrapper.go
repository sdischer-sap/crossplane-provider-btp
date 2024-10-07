/*
SAP XSUAA REST API

Provides access to RoleTemplates, Roles, RoleCollection etc. using the XSUAA REST API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the AuthDataWrapper type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthDataWrapper{}

// AuthDataWrapper struct for AuthDataWrapper
type AuthDataWrapper struct {
	AppId *string `json:"appId,omitempty"`
	AuthorizationData []AuthData `json:"authorizationData,omitempty"`
}

// NewAuthDataWrapper instantiates a new AuthDataWrapper object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthDataWrapper() *AuthDataWrapper {
	this := AuthDataWrapper{}
	return &this
}

// NewAuthDataWrapperWithDefaults instantiates a new AuthDataWrapper object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthDataWrapperWithDefaults() *AuthDataWrapper {
	this := AuthDataWrapper{}
	return &this
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *AuthDataWrapper) GetAppId() string {
	if o == nil || IsNil(o.AppId) {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthDataWrapper) GetAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.AppId) {
		return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *AuthDataWrapper) HasAppId() bool {
	if o != nil && !IsNil(o.AppId) {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *AuthDataWrapper) SetAppId(v string) {
	o.AppId = &v
}

// GetAuthorizationData returns the AuthorizationData field value if set, zero value otherwise.
func (o *AuthDataWrapper) GetAuthorizationData() []AuthData {
	if o == nil || IsNil(o.AuthorizationData) {
		var ret []AuthData
		return ret
	}
	return o.AuthorizationData
}

// GetAuthorizationDataOk returns a tuple with the AuthorizationData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthDataWrapper) GetAuthorizationDataOk() ([]AuthData, bool) {
	if o == nil || IsNil(o.AuthorizationData) {
		return nil, false
	}
	return o.AuthorizationData, true
}

// HasAuthorizationData returns a boolean if a field has been set.
func (o *AuthDataWrapper) HasAuthorizationData() bool {
	if o != nil && !IsNil(o.AuthorizationData) {
		return true
	}

	return false
}

// SetAuthorizationData gets a reference to the given []AuthData and assigns it to the AuthorizationData field.
func (o *AuthDataWrapper) SetAuthorizationData(v []AuthData) {
	o.AuthorizationData = v
}

func (o AuthDataWrapper) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthDataWrapper) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppId) {
		toSerialize["appId"] = o.AppId
	}
	if !IsNil(o.AuthorizationData) {
		toSerialize["authorizationData"] = o.AuthorizationData
	}
	return toSerialize, nil
}

type NullableAuthDataWrapper struct {
	value *AuthDataWrapper
	isSet bool
}

func (v NullableAuthDataWrapper) Get() *AuthDataWrapper {
	return v.value
}

func (v *NullableAuthDataWrapper) Set(val *AuthDataWrapper) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthDataWrapper) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthDataWrapper) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthDataWrapper(val *AuthDataWrapper) *NullableAuthDataWrapper {
	return &NullableAuthDataWrapper{value: val, isSet: true}
}

func (v NullableAuthDataWrapper) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthDataWrapper) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


