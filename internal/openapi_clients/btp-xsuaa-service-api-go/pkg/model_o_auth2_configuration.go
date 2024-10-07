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

// checks if the OAuth2Configuration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OAuth2Configuration{}

// OAuth2Configuration struct for OAuth2Configuration
type OAuth2Configuration struct {
	TokenValidity *int32 `json:"token-validity,omitempty"`
	RefreshTokenValidity *int32 `json:"refresh-token-validity,omitempty"`
	Autoapprove *bool `json:"autoapprove,omitempty"`
	GrantTypes []string `json:"grant-types,omitempty"`
	SystemAttributes []string `json:"system-attributes,omitempty"`
	Allowedproviders []string `json:"allowedproviders,omitempty"`
	RedirectUris []string `json:"redirect-uris,omitempty"`
	CredentialTypes []string `json:"credential-types,omitempty"`
}

// NewOAuth2Configuration instantiates a new OAuth2Configuration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth2Configuration() *OAuth2Configuration {
	this := OAuth2Configuration{}
	return &this
}

// NewOAuth2ConfigurationWithDefaults instantiates a new OAuth2Configuration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuth2ConfigurationWithDefaults() *OAuth2Configuration {
	this := OAuth2Configuration{}
	return &this
}

// GetTokenValidity returns the TokenValidity field value if set, zero value otherwise.
func (o *OAuth2Configuration) GetTokenValidity() int32 {
	if o == nil || IsNil(o.TokenValidity) {
		var ret int32
		return ret
	}
	return *o.TokenValidity
}

// GetTokenValidityOk returns a tuple with the TokenValidity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2Configuration) GetTokenValidityOk() (*int32, bool) {
	if o == nil || IsNil(o.TokenValidity) {
		return nil, false
	}
	return o.TokenValidity, true
}

// HasTokenValidity returns a boolean if a field has been set.
func (o *OAuth2Configuration) HasTokenValidity() bool {
	if o != nil && !IsNil(o.TokenValidity) {
		return true
	}

	return false
}

// SetTokenValidity gets a reference to the given int32 and assigns it to the TokenValidity field.
func (o *OAuth2Configuration) SetTokenValidity(v int32) {
	o.TokenValidity = &v
}

// GetRefreshTokenValidity returns the RefreshTokenValidity field value if set, zero value otherwise.
func (o *OAuth2Configuration) GetRefreshTokenValidity() int32 {
	if o == nil || IsNil(o.RefreshTokenValidity) {
		var ret int32
		return ret
	}
	return *o.RefreshTokenValidity
}

// GetRefreshTokenValidityOk returns a tuple with the RefreshTokenValidity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2Configuration) GetRefreshTokenValidityOk() (*int32, bool) {
	if o == nil || IsNil(o.RefreshTokenValidity) {
		return nil, false
	}
	return o.RefreshTokenValidity, true
}

// HasRefreshTokenValidity returns a boolean if a field has been set.
func (o *OAuth2Configuration) HasRefreshTokenValidity() bool {
	if o != nil && !IsNil(o.RefreshTokenValidity) {
		return true
	}

	return false
}

// SetRefreshTokenValidity gets a reference to the given int32 and assigns it to the RefreshTokenValidity field.
func (o *OAuth2Configuration) SetRefreshTokenValidity(v int32) {
	o.RefreshTokenValidity = &v
}

// GetAutoapprove returns the Autoapprove field value if set, zero value otherwise.
func (o *OAuth2Configuration) GetAutoapprove() bool {
	if o == nil || IsNil(o.Autoapprove) {
		var ret bool
		return ret
	}
	return *o.Autoapprove
}

// GetAutoapproveOk returns a tuple with the Autoapprove field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2Configuration) GetAutoapproveOk() (*bool, bool) {
	if o == nil || IsNil(o.Autoapprove) {
		return nil, false
	}
	return o.Autoapprove, true
}

// HasAutoapprove returns a boolean if a field has been set.
func (o *OAuth2Configuration) HasAutoapprove() bool {
	if o != nil && !IsNil(o.Autoapprove) {
		return true
	}

	return false
}

// SetAutoapprove gets a reference to the given bool and assigns it to the Autoapprove field.
func (o *OAuth2Configuration) SetAutoapprove(v bool) {
	o.Autoapprove = &v
}

// GetGrantTypes returns the GrantTypes field value if set, zero value otherwise.
func (o *OAuth2Configuration) GetGrantTypes() []string {
	if o == nil || IsNil(o.GrantTypes) {
		var ret []string
		return ret
	}
	return o.GrantTypes
}

// GetGrantTypesOk returns a tuple with the GrantTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2Configuration) GetGrantTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.GrantTypes) {
		return nil, false
	}
	return o.GrantTypes, true
}

// HasGrantTypes returns a boolean if a field has been set.
func (o *OAuth2Configuration) HasGrantTypes() bool {
	if o != nil && !IsNil(o.GrantTypes) {
		return true
	}

	return false
}

// SetGrantTypes gets a reference to the given []string and assigns it to the GrantTypes field.
func (o *OAuth2Configuration) SetGrantTypes(v []string) {
	o.GrantTypes = v
}

// GetSystemAttributes returns the SystemAttributes field value if set, zero value otherwise.
func (o *OAuth2Configuration) GetSystemAttributes() []string {
	if o == nil || IsNil(o.SystemAttributes) {
		var ret []string
		return ret
	}
	return o.SystemAttributes
}

// GetSystemAttributesOk returns a tuple with the SystemAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2Configuration) GetSystemAttributesOk() ([]string, bool) {
	if o == nil || IsNil(o.SystemAttributes) {
		return nil, false
	}
	return o.SystemAttributes, true
}

// HasSystemAttributes returns a boolean if a field has been set.
func (o *OAuth2Configuration) HasSystemAttributes() bool {
	if o != nil && !IsNil(o.SystemAttributes) {
		return true
	}

	return false
}

// SetSystemAttributes gets a reference to the given []string and assigns it to the SystemAttributes field.
func (o *OAuth2Configuration) SetSystemAttributes(v []string) {
	o.SystemAttributes = v
}

// GetAllowedproviders returns the Allowedproviders field value if set, zero value otherwise.
func (o *OAuth2Configuration) GetAllowedproviders() []string {
	if o == nil || IsNil(o.Allowedproviders) {
		var ret []string
		return ret
	}
	return o.Allowedproviders
}

// GetAllowedprovidersOk returns a tuple with the Allowedproviders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2Configuration) GetAllowedprovidersOk() ([]string, bool) {
	if o == nil || IsNil(o.Allowedproviders) {
		return nil, false
	}
	return o.Allowedproviders, true
}

// HasAllowedproviders returns a boolean if a field has been set.
func (o *OAuth2Configuration) HasAllowedproviders() bool {
	if o != nil && !IsNil(o.Allowedproviders) {
		return true
	}

	return false
}

// SetAllowedproviders gets a reference to the given []string and assigns it to the Allowedproviders field.
func (o *OAuth2Configuration) SetAllowedproviders(v []string) {
	o.Allowedproviders = v
}

// GetRedirectUris returns the RedirectUris field value if set, zero value otherwise.
func (o *OAuth2Configuration) GetRedirectUris() []string {
	if o == nil || IsNil(o.RedirectUris) {
		var ret []string
		return ret
	}
	return o.RedirectUris
}

// GetRedirectUrisOk returns a tuple with the RedirectUris field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2Configuration) GetRedirectUrisOk() ([]string, bool) {
	if o == nil || IsNil(o.RedirectUris) {
		return nil, false
	}
	return o.RedirectUris, true
}

// HasRedirectUris returns a boolean if a field has been set.
func (o *OAuth2Configuration) HasRedirectUris() bool {
	if o != nil && !IsNil(o.RedirectUris) {
		return true
	}

	return false
}

// SetRedirectUris gets a reference to the given []string and assigns it to the RedirectUris field.
func (o *OAuth2Configuration) SetRedirectUris(v []string) {
	o.RedirectUris = v
}

// GetCredentialTypes returns the CredentialTypes field value if set, zero value otherwise.
func (o *OAuth2Configuration) GetCredentialTypes() []string {
	if o == nil || IsNil(o.CredentialTypes) {
		var ret []string
		return ret
	}
	return o.CredentialTypes
}

// GetCredentialTypesOk returns a tuple with the CredentialTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth2Configuration) GetCredentialTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.CredentialTypes) {
		return nil, false
	}
	return o.CredentialTypes, true
}

// HasCredentialTypes returns a boolean if a field has been set.
func (o *OAuth2Configuration) HasCredentialTypes() bool {
	if o != nil && !IsNil(o.CredentialTypes) {
		return true
	}

	return false
}

// SetCredentialTypes gets a reference to the given []string and assigns it to the CredentialTypes field.
func (o *OAuth2Configuration) SetCredentialTypes(v []string) {
	o.CredentialTypes = v
}

func (o OAuth2Configuration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OAuth2Configuration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TokenValidity) {
		toSerialize["token-validity"] = o.TokenValidity
	}
	if !IsNil(o.RefreshTokenValidity) {
		toSerialize["refresh-token-validity"] = o.RefreshTokenValidity
	}
	if !IsNil(o.Autoapprove) {
		toSerialize["autoapprove"] = o.Autoapprove
	}
	if !IsNil(o.GrantTypes) {
		toSerialize["grant-types"] = o.GrantTypes
	}
	if !IsNil(o.SystemAttributes) {
		toSerialize["system-attributes"] = o.SystemAttributes
	}
	if !IsNil(o.Allowedproviders) {
		toSerialize["allowedproviders"] = o.Allowedproviders
	}
	if !IsNil(o.RedirectUris) {
		toSerialize["redirect-uris"] = o.RedirectUris
	}
	if !IsNil(o.CredentialTypes) {
		toSerialize["credential-types"] = o.CredentialTypes
	}
	return toSerialize, nil
}

type NullableOAuth2Configuration struct {
	value *OAuth2Configuration
	isSet bool
}

func (v NullableOAuth2Configuration) Get() *OAuth2Configuration {
	return v.value
}

func (v *NullableOAuth2Configuration) Set(val *OAuth2Configuration) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth2Configuration) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth2Configuration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth2Configuration(val *OAuth2Configuration) *NullableOAuth2Configuration {
	return &NullableOAuth2Configuration{value: val, isSet: true}
}

func (v NullableOAuth2Configuration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth2Configuration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


