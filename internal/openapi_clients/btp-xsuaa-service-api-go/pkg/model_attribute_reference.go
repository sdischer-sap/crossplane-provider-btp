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

// checks if the AttributeReference type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttributeReference{}

// AttributeReference struct for AttributeReference
type AttributeReference struct {
	Name *string `json:"name,omitempty"`
	ValueRequired *bool `json:"valueRequired,omitempty"`
	DefaultValues []string `json:"default-values,omitempty"`
	ValueType *string `json:"valueType,omitempty"`
	Description *string `json:"description,omitempty"`
}

// NewAttributeReference instantiates a new AttributeReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttributeReference() *AttributeReference {
	this := AttributeReference{}
	return &this
}

// NewAttributeReferenceWithDefaults instantiates a new AttributeReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttributeReferenceWithDefaults() *AttributeReference {
	this := AttributeReference{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AttributeReference) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributeReference) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AttributeReference) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AttributeReference) SetName(v string) {
	o.Name = &v
}

// GetValueRequired returns the ValueRequired field value if set, zero value otherwise.
func (o *AttributeReference) GetValueRequired() bool {
	if o == nil || IsNil(o.ValueRequired) {
		var ret bool
		return ret
	}
	return *o.ValueRequired
}

// GetValueRequiredOk returns a tuple with the ValueRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributeReference) GetValueRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.ValueRequired) {
		return nil, false
	}
	return o.ValueRequired, true
}

// HasValueRequired returns a boolean if a field has been set.
func (o *AttributeReference) HasValueRequired() bool {
	if o != nil && !IsNil(o.ValueRequired) {
		return true
	}

	return false
}

// SetValueRequired gets a reference to the given bool and assigns it to the ValueRequired field.
func (o *AttributeReference) SetValueRequired(v bool) {
	o.ValueRequired = &v
}

// GetDefaultValues returns the DefaultValues field value if set, zero value otherwise.
func (o *AttributeReference) GetDefaultValues() []string {
	if o == nil || IsNil(o.DefaultValues) {
		var ret []string
		return ret
	}
	return o.DefaultValues
}

// GetDefaultValuesOk returns a tuple with the DefaultValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributeReference) GetDefaultValuesOk() ([]string, bool) {
	if o == nil || IsNil(o.DefaultValues) {
		return nil, false
	}
	return o.DefaultValues, true
}

// HasDefaultValues returns a boolean if a field has been set.
func (o *AttributeReference) HasDefaultValues() bool {
	if o != nil && !IsNil(o.DefaultValues) {
		return true
	}

	return false
}

// SetDefaultValues gets a reference to the given []string and assigns it to the DefaultValues field.
func (o *AttributeReference) SetDefaultValues(v []string) {
	o.DefaultValues = v
}

// GetValueType returns the ValueType field value if set, zero value otherwise.
func (o *AttributeReference) GetValueType() string {
	if o == nil || IsNil(o.ValueType) {
		var ret string
		return ret
	}
	return *o.ValueType
}

// GetValueTypeOk returns a tuple with the ValueType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributeReference) GetValueTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ValueType) {
		return nil, false
	}
	return o.ValueType, true
}

// HasValueType returns a boolean if a field has been set.
func (o *AttributeReference) HasValueType() bool {
	if o != nil && !IsNil(o.ValueType) {
		return true
	}

	return false
}

// SetValueType gets a reference to the given string and assigns it to the ValueType field.
func (o *AttributeReference) SetValueType(v string) {
	o.ValueType = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AttributeReference) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributeReference) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AttributeReference) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AttributeReference) SetDescription(v string) {
	o.Description = &v
}

func (o AttributeReference) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttributeReference) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ValueRequired) {
		toSerialize["valueRequired"] = o.ValueRequired
	}
	if !IsNil(o.DefaultValues) {
		toSerialize["default-values"] = o.DefaultValues
	}
	if !IsNil(o.ValueType) {
		toSerialize["valueType"] = o.ValueType
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableAttributeReference struct {
	value *AttributeReference
	isSet bool
}

func (v NullableAttributeReference) Get() *AttributeReference {
	return v.value
}

func (v *NullableAttributeReference) Set(val *AttributeReference) {
	v.value = val
	v.isSet = true
}

func (v NullableAttributeReference) IsSet() bool {
	return v.isSet
}

func (v *NullableAttributeReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttributeReference(val *AttributeReference) *NullableAttributeReference {
	return &NullableAttributeReference{value: val, isSet: true}
}

func (v NullableAttributeReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttributeReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


