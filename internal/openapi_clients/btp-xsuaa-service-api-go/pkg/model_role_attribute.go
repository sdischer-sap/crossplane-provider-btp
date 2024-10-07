/*
SAP XSUAA REST API

Provides access to RoleTemplates, Roles, RoleCollection etc. using the XSUAA REST API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the RoleAttribute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoleAttribute{}

// RoleAttribute struct for RoleAttribute
type RoleAttribute struct {
	AttributeName string `json:"attributeName"`
	AttributeValueOrigin string `json:"attributeValueOrigin"`
	AttributeValues []string `json:"attributeValues"`
	ValueRequired *bool `json:"valueRequired,omitempty"`
	Description *string `json:"description,omitempty"`
}

type _RoleAttribute RoleAttribute

// NewRoleAttribute instantiates a new RoleAttribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleAttribute(attributeName string, attributeValueOrigin string, attributeValues []string) *RoleAttribute {
	this := RoleAttribute{}
	this.AttributeName = attributeName
	this.AttributeValueOrigin = attributeValueOrigin
	this.AttributeValues = attributeValues
	return &this
}

// NewRoleAttributeWithDefaults instantiates a new RoleAttribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleAttributeWithDefaults() *RoleAttribute {
	this := RoleAttribute{}
	return &this
}

// GetAttributeName returns the AttributeName field value
func (o *RoleAttribute) GetAttributeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AttributeName
}

// GetAttributeNameOk returns a tuple with the AttributeName field value
// and a boolean to check if the value has been set.
func (o *RoleAttribute) GetAttributeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttributeName, true
}

// SetAttributeName sets field value
func (o *RoleAttribute) SetAttributeName(v string) {
	o.AttributeName = v
}

// GetAttributeValueOrigin returns the AttributeValueOrigin field value
func (o *RoleAttribute) GetAttributeValueOrigin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AttributeValueOrigin
}

// GetAttributeValueOriginOk returns a tuple with the AttributeValueOrigin field value
// and a boolean to check if the value has been set.
func (o *RoleAttribute) GetAttributeValueOriginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttributeValueOrigin, true
}

// SetAttributeValueOrigin sets field value
func (o *RoleAttribute) SetAttributeValueOrigin(v string) {
	o.AttributeValueOrigin = v
}

// GetAttributeValues returns the AttributeValues field value
func (o *RoleAttribute) GetAttributeValues() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AttributeValues
}

// GetAttributeValuesOk returns a tuple with the AttributeValues field value
// and a boolean to check if the value has been set.
func (o *RoleAttribute) GetAttributeValuesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AttributeValues, true
}

// SetAttributeValues sets field value
func (o *RoleAttribute) SetAttributeValues(v []string) {
	o.AttributeValues = v
}

// GetValueRequired returns the ValueRequired field value if set, zero value otherwise.
func (o *RoleAttribute) GetValueRequired() bool {
	if o == nil || IsNil(o.ValueRequired) {
		var ret bool
		return ret
	}
	return *o.ValueRequired
}

// GetValueRequiredOk returns a tuple with the ValueRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleAttribute) GetValueRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.ValueRequired) {
		return nil, false
	}
	return o.ValueRequired, true
}

// HasValueRequired returns a boolean if a field has been set.
func (o *RoleAttribute) HasValueRequired() bool {
	if o != nil && !IsNil(o.ValueRequired) {
		return true
	}

	return false
}

// SetValueRequired gets a reference to the given bool and assigns it to the ValueRequired field.
func (o *RoleAttribute) SetValueRequired(v bool) {
	o.ValueRequired = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RoleAttribute) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleAttribute) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RoleAttribute) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RoleAttribute) SetDescription(v string) {
	o.Description = &v
}

func (o RoleAttribute) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoleAttribute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["attributeName"] = o.AttributeName
	toSerialize["attributeValueOrigin"] = o.AttributeValueOrigin
	toSerialize["attributeValues"] = o.AttributeValues
	if !IsNil(o.ValueRequired) {
		toSerialize["valueRequired"] = o.ValueRequired
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

func (o *RoleAttribute) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"attributeName",
		"attributeValueOrigin",
		"attributeValues",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varRoleAttribute := _RoleAttribute{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRoleAttribute)

	if err != nil {
		return err
	}

	*o = RoleAttribute(varRoleAttribute)

	return err
}

type NullableRoleAttribute struct {
	value *RoleAttribute
	isSet bool
}

func (v NullableRoleAttribute) Get() *RoleAttribute {
	return v.value
}

func (v *NullableRoleAttribute) Set(val *RoleAttribute) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleAttribute(val *RoleAttribute) *NullableRoleAttribute {
	return &NullableRoleAttribute{value: val, isSet: true}
}

func (v NullableRoleAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

