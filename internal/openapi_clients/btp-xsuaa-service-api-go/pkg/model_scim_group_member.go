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

// checks if the ScimGroupMember type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScimGroupMember{}

// ScimGroupMember struct for ScimGroupMember
type ScimGroupMember struct {
	Origin *string `json:"origin,omitempty"`
	Operation *string `json:"operation,omitempty"`
	Type *string `json:"type,omitempty"`
	Entity *ScimCore `json:"entity,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewScimGroupMember instantiates a new ScimGroupMember object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScimGroupMember() *ScimGroupMember {
	this := ScimGroupMember{}
	return &this
}

// NewScimGroupMemberWithDefaults instantiates a new ScimGroupMember object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScimGroupMemberWithDefaults() *ScimGroupMember {
	this := ScimGroupMember{}
	return &this
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *ScimGroupMember) GetOrigin() string {
	if o == nil || IsNil(o.Origin) {
		var ret string
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimGroupMember) GetOriginOk() (*string, bool) {
	if o == nil || IsNil(o.Origin) {
		return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *ScimGroupMember) HasOrigin() bool {
	if o != nil && !IsNil(o.Origin) {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given string and assigns it to the Origin field.
func (o *ScimGroupMember) SetOrigin(v string) {
	o.Origin = &v
}

// GetOperation returns the Operation field value if set, zero value otherwise.
func (o *ScimGroupMember) GetOperation() string {
	if o == nil || IsNil(o.Operation) {
		var ret string
		return ret
	}
	return *o.Operation
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimGroupMember) GetOperationOk() (*string, bool) {
	if o == nil || IsNil(o.Operation) {
		return nil, false
	}
	return o.Operation, true
}

// HasOperation returns a boolean if a field has been set.
func (o *ScimGroupMember) HasOperation() bool {
	if o != nil && !IsNil(o.Operation) {
		return true
	}

	return false
}

// SetOperation gets a reference to the given string and assigns it to the Operation field.
func (o *ScimGroupMember) SetOperation(v string) {
	o.Operation = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ScimGroupMember) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimGroupMember) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ScimGroupMember) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ScimGroupMember) SetType(v string) {
	o.Type = &v
}

// GetEntity returns the Entity field value if set, zero value otherwise.
func (o *ScimGroupMember) GetEntity() ScimCore {
	if o == nil || IsNil(o.Entity) {
		var ret ScimCore
		return ret
	}
	return *o.Entity
}

// GetEntityOk returns a tuple with the Entity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimGroupMember) GetEntityOk() (*ScimCore, bool) {
	if o == nil || IsNil(o.Entity) {
		return nil, false
	}
	return o.Entity, true
}

// HasEntity returns a boolean if a field has been set.
func (o *ScimGroupMember) HasEntity() bool {
	if o != nil && !IsNil(o.Entity) {
		return true
	}

	return false
}

// SetEntity gets a reference to the given ScimCore and assigns it to the Entity field.
func (o *ScimGroupMember) SetEntity(v ScimCore) {
	o.Entity = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ScimGroupMember) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimGroupMember) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ScimGroupMember) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ScimGroupMember) SetValue(v string) {
	o.Value = &v
}

func (o ScimGroupMember) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScimGroupMember) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Origin) {
		toSerialize["origin"] = o.Origin
	}
	if !IsNil(o.Operation) {
		toSerialize["operation"] = o.Operation
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Entity) {
		toSerialize["entity"] = o.Entity
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableScimGroupMember struct {
	value *ScimGroupMember
	isSet bool
}

func (v NullableScimGroupMember) Get() *ScimGroupMember {
	return v.value
}

func (v *NullableScimGroupMember) Set(val *ScimGroupMember) {
	v.value = val
	v.isSet = true
}

func (v NullableScimGroupMember) IsSet() bool {
	return v.isSet
}

func (v *NullableScimGroupMember) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScimGroupMember(val *ScimGroupMember) *NullableScimGroupMember {
	return &NullableScimGroupMember{value: val, isSet: true}
}

func (v NullableScimGroupMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScimGroupMember) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

