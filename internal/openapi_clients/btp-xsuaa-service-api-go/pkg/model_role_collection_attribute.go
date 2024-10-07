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

// checks if the RoleCollectionAttribute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoleCollectionAttribute{}

// RoleCollectionAttribute struct for RoleCollectionAttribute
type RoleCollectionAttribute struct {
	RoleCollectionName *string `json:"roleCollectionName,omitempty"`
	RoleCollectionIdentityZone *string `json:"roleCollectionIdentityZone,omitempty"`
	SamlAttrName *string `json:"samlAttrName,omitempty"`
	AttributeName *string `json:"attributeName,omitempty"`
	SamlAttributeValue *string `json:"samlAttributeValue,omitempty"`
	AttributeValue *string `json:"attributeValue,omitempty"`
	ComparisonOperator *string `json:"comparisonOperator,omitempty"`
	SamlEntityId *string `json:"samlEntityId,omitempty"`
	IdpId *string `json:"idpId,omitempty"`
	IdpDisplayName *string `json:"idpDisplayName,omitempty"`
}

// NewRoleCollectionAttribute instantiates a new RoleCollectionAttribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleCollectionAttribute() *RoleCollectionAttribute {
	this := RoleCollectionAttribute{}
	return &this
}

// NewRoleCollectionAttributeWithDefaults instantiates a new RoleCollectionAttribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleCollectionAttributeWithDefaults() *RoleCollectionAttribute {
	this := RoleCollectionAttribute{}
	return &this
}

// GetRoleCollectionName returns the RoleCollectionName field value if set, zero value otherwise.
func (o *RoleCollectionAttribute) GetRoleCollectionName() string {
	if o == nil || IsNil(o.RoleCollectionName) {
		var ret string
		return ret
	}
	return *o.RoleCollectionName
}

// GetRoleCollectionNameOk returns a tuple with the RoleCollectionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleCollectionAttribute) GetRoleCollectionNameOk() (*string, bool) {
	if o == nil || IsNil(o.RoleCollectionName) {
		return nil, false
	}
	return o.RoleCollectionName, true
}

// HasRoleCollectionName returns a boolean if a field has been set.
func (o *RoleCollectionAttribute) HasRoleCollectionName() bool {
	if o != nil && !IsNil(o.RoleCollectionName) {
		return true
	}

	return false
}

// SetRoleCollectionName gets a reference to the given string and assigns it to the RoleCollectionName field.
func (o *RoleCollectionAttribute) SetRoleCollectionName(v string) {
	o.RoleCollectionName = &v
}

// GetRoleCollectionIdentityZone returns the RoleCollectionIdentityZone field value if set, zero value otherwise.
func (o *RoleCollectionAttribute) GetRoleCollectionIdentityZone() string {
	if o == nil || IsNil(o.RoleCollectionIdentityZone) {
		var ret string
		return ret
	}
	return *o.RoleCollectionIdentityZone
}

// GetRoleCollectionIdentityZoneOk returns a tuple with the RoleCollectionIdentityZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleCollectionAttribute) GetRoleCollectionIdentityZoneOk() (*string, bool) {
	if o == nil || IsNil(o.RoleCollectionIdentityZone) {
		return nil, false
	}
	return o.RoleCollectionIdentityZone, true
}

// HasRoleCollectionIdentityZone returns a boolean if a field has been set.
func (o *RoleCollectionAttribute) HasRoleCollectionIdentityZone() bool {
	if o != nil && !IsNil(o.RoleCollectionIdentityZone) {
		return true
	}

	return false
}

// SetRoleCollectionIdentityZone gets a reference to the given string and assigns it to the RoleCollectionIdentityZone field.
func (o *RoleCollectionAttribute) SetRoleCollectionIdentityZone(v string) {
	o.RoleCollectionIdentityZone = &v
}

// GetSamlAttrName returns the SamlAttrName field value if set, zero value otherwise.
func (o *RoleCollectionAttribute) GetSamlAttrName() string {
	if o == nil || IsNil(o.SamlAttrName) {
		var ret string
		return ret
	}
	return *o.SamlAttrName
}

// GetSamlAttrNameOk returns a tuple with the SamlAttrName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleCollectionAttribute) GetSamlAttrNameOk() (*string, bool) {
	if o == nil || IsNil(o.SamlAttrName) {
		return nil, false
	}
	return o.SamlAttrName, true
}

// HasSamlAttrName returns a boolean if a field has been set.
func (o *RoleCollectionAttribute) HasSamlAttrName() bool {
	if o != nil && !IsNil(o.SamlAttrName) {
		return true
	}

	return false
}

// SetSamlAttrName gets a reference to the given string and assigns it to the SamlAttrName field.
func (o *RoleCollectionAttribute) SetSamlAttrName(v string) {
	o.SamlAttrName = &v
}

// GetAttributeName returns the AttributeName field value if set, zero value otherwise.
func (o *RoleCollectionAttribute) GetAttributeName() string {
	if o == nil || IsNil(o.AttributeName) {
		var ret string
		return ret
	}
	return *o.AttributeName
}

// GetAttributeNameOk returns a tuple with the AttributeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleCollectionAttribute) GetAttributeNameOk() (*string, bool) {
	if o == nil || IsNil(o.AttributeName) {
		return nil, false
	}
	return o.AttributeName, true
}

// HasAttributeName returns a boolean if a field has been set.
func (o *RoleCollectionAttribute) HasAttributeName() bool {
	if o != nil && !IsNil(o.AttributeName) {
		return true
	}

	return false
}

// SetAttributeName gets a reference to the given string and assigns it to the AttributeName field.
func (o *RoleCollectionAttribute) SetAttributeName(v string) {
	o.AttributeName = &v
}

// GetSamlAttributeValue returns the SamlAttributeValue field value if set, zero value otherwise.
func (o *RoleCollectionAttribute) GetSamlAttributeValue() string {
	if o == nil || IsNil(o.SamlAttributeValue) {
		var ret string
		return ret
	}
	return *o.SamlAttributeValue
}

// GetSamlAttributeValueOk returns a tuple with the SamlAttributeValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleCollectionAttribute) GetSamlAttributeValueOk() (*string, bool) {
	if o == nil || IsNil(o.SamlAttributeValue) {
		return nil, false
	}
	return o.SamlAttributeValue, true
}

// HasSamlAttributeValue returns a boolean if a field has been set.
func (o *RoleCollectionAttribute) HasSamlAttributeValue() bool {
	if o != nil && !IsNil(o.SamlAttributeValue) {
		return true
	}

	return false
}

// SetSamlAttributeValue gets a reference to the given string and assigns it to the SamlAttributeValue field.
func (o *RoleCollectionAttribute) SetSamlAttributeValue(v string) {
	o.SamlAttributeValue = &v
}

// GetAttributeValue returns the AttributeValue field value if set, zero value otherwise.
func (o *RoleCollectionAttribute) GetAttributeValue() string {
	if o == nil || IsNil(o.AttributeValue) {
		var ret string
		return ret
	}
	return *o.AttributeValue
}

// GetAttributeValueOk returns a tuple with the AttributeValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleCollectionAttribute) GetAttributeValueOk() (*string, bool) {
	if o == nil || IsNil(o.AttributeValue) {
		return nil, false
	}
	return o.AttributeValue, true
}

// HasAttributeValue returns a boolean if a field has been set.
func (o *RoleCollectionAttribute) HasAttributeValue() bool {
	if o != nil && !IsNil(o.AttributeValue) {
		return true
	}

	return false
}

// SetAttributeValue gets a reference to the given string and assigns it to the AttributeValue field.
func (o *RoleCollectionAttribute) SetAttributeValue(v string) {
	o.AttributeValue = &v
}

// GetComparisonOperator returns the ComparisonOperator field value if set, zero value otherwise.
func (o *RoleCollectionAttribute) GetComparisonOperator() string {
	if o == nil || IsNil(o.ComparisonOperator) {
		var ret string
		return ret
	}
	return *o.ComparisonOperator
}

// GetComparisonOperatorOk returns a tuple with the ComparisonOperator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleCollectionAttribute) GetComparisonOperatorOk() (*string, bool) {
	if o == nil || IsNil(o.ComparisonOperator) {
		return nil, false
	}
	return o.ComparisonOperator, true
}

// HasComparisonOperator returns a boolean if a field has been set.
func (o *RoleCollectionAttribute) HasComparisonOperator() bool {
	if o != nil && !IsNil(o.ComparisonOperator) {
		return true
	}

	return false
}

// SetComparisonOperator gets a reference to the given string and assigns it to the ComparisonOperator field.
func (o *RoleCollectionAttribute) SetComparisonOperator(v string) {
	o.ComparisonOperator = &v
}

// GetSamlEntityId returns the SamlEntityId field value if set, zero value otherwise.
func (o *RoleCollectionAttribute) GetSamlEntityId() string {
	if o == nil || IsNil(o.SamlEntityId) {
		var ret string
		return ret
	}
	return *o.SamlEntityId
}

// GetSamlEntityIdOk returns a tuple with the SamlEntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleCollectionAttribute) GetSamlEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.SamlEntityId) {
		return nil, false
	}
	return o.SamlEntityId, true
}

// HasSamlEntityId returns a boolean if a field has been set.
func (o *RoleCollectionAttribute) HasSamlEntityId() bool {
	if o != nil && !IsNil(o.SamlEntityId) {
		return true
	}

	return false
}

// SetSamlEntityId gets a reference to the given string and assigns it to the SamlEntityId field.
func (o *RoleCollectionAttribute) SetSamlEntityId(v string) {
	o.SamlEntityId = &v
}

// GetIdpId returns the IdpId field value if set, zero value otherwise.
func (o *RoleCollectionAttribute) GetIdpId() string {
	if o == nil || IsNil(o.IdpId) {
		var ret string
		return ret
	}
	return *o.IdpId
}

// GetIdpIdOk returns a tuple with the IdpId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleCollectionAttribute) GetIdpIdOk() (*string, bool) {
	if o == nil || IsNil(o.IdpId) {
		return nil, false
	}
	return o.IdpId, true
}

// HasIdpId returns a boolean if a field has been set.
func (o *RoleCollectionAttribute) HasIdpId() bool {
	if o != nil && !IsNil(o.IdpId) {
		return true
	}

	return false
}

// SetIdpId gets a reference to the given string and assigns it to the IdpId field.
func (o *RoleCollectionAttribute) SetIdpId(v string) {
	o.IdpId = &v
}

// GetIdpDisplayName returns the IdpDisplayName field value if set, zero value otherwise.
func (o *RoleCollectionAttribute) GetIdpDisplayName() string {
	if o == nil || IsNil(o.IdpDisplayName) {
		var ret string
		return ret
	}
	return *o.IdpDisplayName
}

// GetIdpDisplayNameOk returns a tuple with the IdpDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleCollectionAttribute) GetIdpDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.IdpDisplayName) {
		return nil, false
	}
	return o.IdpDisplayName, true
}

// HasIdpDisplayName returns a boolean if a field has been set.
func (o *RoleCollectionAttribute) HasIdpDisplayName() bool {
	if o != nil && !IsNil(o.IdpDisplayName) {
		return true
	}

	return false
}

// SetIdpDisplayName gets a reference to the given string and assigns it to the IdpDisplayName field.
func (o *RoleCollectionAttribute) SetIdpDisplayName(v string) {
	o.IdpDisplayName = &v
}

func (o RoleCollectionAttribute) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoleCollectionAttribute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RoleCollectionName) {
		toSerialize["roleCollectionName"] = o.RoleCollectionName
	}
	if !IsNil(o.RoleCollectionIdentityZone) {
		toSerialize["roleCollectionIdentityZone"] = o.RoleCollectionIdentityZone
	}
	if !IsNil(o.SamlAttrName) {
		toSerialize["samlAttrName"] = o.SamlAttrName
	}
	if !IsNil(o.AttributeName) {
		toSerialize["attributeName"] = o.AttributeName
	}
	if !IsNil(o.SamlAttributeValue) {
		toSerialize["samlAttributeValue"] = o.SamlAttributeValue
	}
	if !IsNil(o.AttributeValue) {
		toSerialize["attributeValue"] = o.AttributeValue
	}
	if !IsNil(o.ComparisonOperator) {
		toSerialize["comparisonOperator"] = o.ComparisonOperator
	}
	if !IsNil(o.SamlEntityId) {
		toSerialize["samlEntityId"] = o.SamlEntityId
	}
	if !IsNil(o.IdpId) {
		toSerialize["idpId"] = o.IdpId
	}
	if !IsNil(o.IdpDisplayName) {
		toSerialize["idpDisplayName"] = o.IdpDisplayName
	}
	return toSerialize, nil
}

type NullableRoleCollectionAttribute struct {
	value *RoleCollectionAttribute
	isSet bool
}

func (v NullableRoleCollectionAttribute) Get() *RoleCollectionAttribute {
	return v.value
}

func (v *NullableRoleCollectionAttribute) Set(val *RoleCollectionAttribute) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleCollectionAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleCollectionAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleCollectionAttribute(val *RoleCollectionAttribute) *NullableRoleCollectionAttribute {
	return &NullableRoleCollectionAttribute{value: val, isSet: true}
}

func (v NullableRoleCollectionAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleCollectionAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

