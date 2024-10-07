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

// checks if the RoleCollectionListDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoleCollectionListDto{}

// RoleCollectionListDto struct for RoleCollectionListDto
type RoleCollectionListDto struct {
	RoleCollections []string `json:"roleCollections,omitempty"`
	RoleCollectionsBySamlAssignment []string `json:"roleCollectionsBySamlAssignment,omitempty"`
}

// NewRoleCollectionListDto instantiates a new RoleCollectionListDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleCollectionListDto() *RoleCollectionListDto {
	this := RoleCollectionListDto{}
	return &this
}

// NewRoleCollectionListDtoWithDefaults instantiates a new RoleCollectionListDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleCollectionListDtoWithDefaults() *RoleCollectionListDto {
	this := RoleCollectionListDto{}
	return &this
}

// GetRoleCollections returns the RoleCollections field value if set, zero value otherwise.
func (o *RoleCollectionListDto) GetRoleCollections() []string {
	if o == nil || IsNil(o.RoleCollections) {
		var ret []string
		return ret
	}
	return o.RoleCollections
}

// GetRoleCollectionsOk returns a tuple with the RoleCollections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleCollectionListDto) GetRoleCollectionsOk() ([]string, bool) {
	if o == nil || IsNil(o.RoleCollections) {
		return nil, false
	}
	return o.RoleCollections, true
}

// HasRoleCollections returns a boolean if a field has been set.
func (o *RoleCollectionListDto) HasRoleCollections() bool {
	if o != nil && !IsNil(o.RoleCollections) {
		return true
	}

	return false
}

// SetRoleCollections gets a reference to the given []string and assigns it to the RoleCollections field.
func (o *RoleCollectionListDto) SetRoleCollections(v []string) {
	o.RoleCollections = v
}

// GetRoleCollectionsBySamlAssignment returns the RoleCollectionsBySamlAssignment field value if set, zero value otherwise.
func (o *RoleCollectionListDto) GetRoleCollectionsBySamlAssignment() []string {
	if o == nil || IsNil(o.RoleCollectionsBySamlAssignment) {
		var ret []string
		return ret
	}
	return o.RoleCollectionsBySamlAssignment
}

// GetRoleCollectionsBySamlAssignmentOk returns a tuple with the RoleCollectionsBySamlAssignment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleCollectionListDto) GetRoleCollectionsBySamlAssignmentOk() ([]string, bool) {
	if o == nil || IsNil(o.RoleCollectionsBySamlAssignment) {
		return nil, false
	}
	return o.RoleCollectionsBySamlAssignment, true
}

// HasRoleCollectionsBySamlAssignment returns a boolean if a field has been set.
func (o *RoleCollectionListDto) HasRoleCollectionsBySamlAssignment() bool {
	if o != nil && !IsNil(o.RoleCollectionsBySamlAssignment) {
		return true
	}

	return false
}

// SetRoleCollectionsBySamlAssignment gets a reference to the given []string and assigns it to the RoleCollectionsBySamlAssignment field.
func (o *RoleCollectionListDto) SetRoleCollectionsBySamlAssignment(v []string) {
	o.RoleCollectionsBySamlAssignment = v
}

func (o RoleCollectionListDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoleCollectionListDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RoleCollections) {
		toSerialize["roleCollections"] = o.RoleCollections
	}
	if !IsNil(o.RoleCollectionsBySamlAssignment) {
		toSerialize["roleCollectionsBySamlAssignment"] = o.RoleCollectionsBySamlAssignment
	}
	return toSerialize, nil
}

type NullableRoleCollectionListDto struct {
	value *RoleCollectionListDto
	isSet bool
}

func (v NullableRoleCollectionListDto) Get() *RoleCollectionListDto {
	return v.value
}

func (v *NullableRoleCollectionListDto) Set(val *RoleCollectionListDto) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleCollectionListDto) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleCollectionListDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleCollectionListDto(val *RoleCollectionListDto) *NullableRoleCollectionListDto {
	return &NullableRoleCollectionListDto{value: val, isSet: true}
}

func (v NullableRoleCollectionListDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleCollectionListDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


