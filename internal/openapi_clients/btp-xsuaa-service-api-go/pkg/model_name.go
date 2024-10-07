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

// checks if the Name type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Name{}

// Name struct for Name
type Name struct {
	Formatted *string `json:"formatted,omitempty"`
	FamilyName *string `json:"familyName,omitempty"`
	GivenName *string `json:"givenName,omitempty"`
	MiddleName *string `json:"middleName,omitempty"`
	HonorificPrefix *string `json:"honorificPrefix,omitempty"`
	HonorificSuffix *string `json:"honorificSuffix,omitempty"`
}

// NewName instantiates a new Name object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewName() *Name {
	this := Name{}
	return &this
}

// NewNameWithDefaults instantiates a new Name object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNameWithDefaults() *Name {
	this := Name{}
	return &this
}

// GetFormatted returns the Formatted field value if set, zero value otherwise.
func (o *Name) GetFormatted() string {
	if o == nil || IsNil(o.Formatted) {
		var ret string
		return ret
	}
	return *o.Formatted
}

// GetFormattedOk returns a tuple with the Formatted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Name) GetFormattedOk() (*string, bool) {
	if o == nil || IsNil(o.Formatted) {
		return nil, false
	}
	return o.Formatted, true
}

// HasFormatted returns a boolean if a field has been set.
func (o *Name) HasFormatted() bool {
	if o != nil && !IsNil(o.Formatted) {
		return true
	}

	return false
}

// SetFormatted gets a reference to the given string and assigns it to the Formatted field.
func (o *Name) SetFormatted(v string) {
	o.Formatted = &v
}

// GetFamilyName returns the FamilyName field value if set, zero value otherwise.
func (o *Name) GetFamilyName() string {
	if o == nil || IsNil(o.FamilyName) {
		var ret string
		return ret
	}
	return *o.FamilyName
}

// GetFamilyNameOk returns a tuple with the FamilyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Name) GetFamilyNameOk() (*string, bool) {
	if o == nil || IsNil(o.FamilyName) {
		return nil, false
	}
	return o.FamilyName, true
}

// HasFamilyName returns a boolean if a field has been set.
func (o *Name) HasFamilyName() bool {
	if o != nil && !IsNil(o.FamilyName) {
		return true
	}

	return false
}

// SetFamilyName gets a reference to the given string and assigns it to the FamilyName field.
func (o *Name) SetFamilyName(v string) {
	o.FamilyName = &v
}

// GetGivenName returns the GivenName field value if set, zero value otherwise.
func (o *Name) GetGivenName() string {
	if o == nil || IsNil(o.GivenName) {
		var ret string
		return ret
	}
	return *o.GivenName
}

// GetGivenNameOk returns a tuple with the GivenName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Name) GetGivenNameOk() (*string, bool) {
	if o == nil || IsNil(o.GivenName) {
		return nil, false
	}
	return o.GivenName, true
}

// HasGivenName returns a boolean if a field has been set.
func (o *Name) HasGivenName() bool {
	if o != nil && !IsNil(o.GivenName) {
		return true
	}

	return false
}

// SetGivenName gets a reference to the given string and assigns it to the GivenName field.
func (o *Name) SetGivenName(v string) {
	o.GivenName = &v
}

// GetMiddleName returns the MiddleName field value if set, zero value otherwise.
func (o *Name) GetMiddleName() string {
	if o == nil || IsNil(o.MiddleName) {
		var ret string
		return ret
	}
	return *o.MiddleName
}

// GetMiddleNameOk returns a tuple with the MiddleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Name) GetMiddleNameOk() (*string, bool) {
	if o == nil || IsNil(o.MiddleName) {
		return nil, false
	}
	return o.MiddleName, true
}

// HasMiddleName returns a boolean if a field has been set.
func (o *Name) HasMiddleName() bool {
	if o != nil && !IsNil(o.MiddleName) {
		return true
	}

	return false
}

// SetMiddleName gets a reference to the given string and assigns it to the MiddleName field.
func (o *Name) SetMiddleName(v string) {
	o.MiddleName = &v
}

// GetHonorificPrefix returns the HonorificPrefix field value if set, zero value otherwise.
func (o *Name) GetHonorificPrefix() string {
	if o == nil || IsNil(o.HonorificPrefix) {
		var ret string
		return ret
	}
	return *o.HonorificPrefix
}

// GetHonorificPrefixOk returns a tuple with the HonorificPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Name) GetHonorificPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.HonorificPrefix) {
		return nil, false
	}
	return o.HonorificPrefix, true
}

// HasHonorificPrefix returns a boolean if a field has been set.
func (o *Name) HasHonorificPrefix() bool {
	if o != nil && !IsNil(o.HonorificPrefix) {
		return true
	}

	return false
}

// SetHonorificPrefix gets a reference to the given string and assigns it to the HonorificPrefix field.
func (o *Name) SetHonorificPrefix(v string) {
	o.HonorificPrefix = &v
}

// GetHonorificSuffix returns the HonorificSuffix field value if set, zero value otherwise.
func (o *Name) GetHonorificSuffix() string {
	if o == nil || IsNil(o.HonorificSuffix) {
		var ret string
		return ret
	}
	return *o.HonorificSuffix
}

// GetHonorificSuffixOk returns a tuple with the HonorificSuffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Name) GetHonorificSuffixOk() (*string, bool) {
	if o == nil || IsNil(o.HonorificSuffix) {
		return nil, false
	}
	return o.HonorificSuffix, true
}

// HasHonorificSuffix returns a boolean if a field has been set.
func (o *Name) HasHonorificSuffix() bool {
	if o != nil && !IsNil(o.HonorificSuffix) {
		return true
	}

	return false
}

// SetHonorificSuffix gets a reference to the given string and assigns it to the HonorificSuffix field.
func (o *Name) SetHonorificSuffix(v string) {
	o.HonorificSuffix = &v
}

func (o Name) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Name) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Formatted) {
		toSerialize["formatted"] = o.Formatted
	}
	if !IsNil(o.FamilyName) {
		toSerialize["familyName"] = o.FamilyName
	}
	if !IsNil(o.GivenName) {
		toSerialize["givenName"] = o.GivenName
	}
	if !IsNil(o.MiddleName) {
		toSerialize["middleName"] = o.MiddleName
	}
	if !IsNil(o.HonorificPrefix) {
		toSerialize["honorificPrefix"] = o.HonorificPrefix
	}
	if !IsNil(o.HonorificSuffix) {
		toSerialize["honorificSuffix"] = o.HonorificSuffix
	}
	return toSerialize, nil
}

type NullableName struct {
	value *Name
	isSet bool
}

func (v NullableName) Get() *Name {
	return v.value
}

func (v *NullableName) Set(val *Name) {
	v.value = val
	v.isSet = true
}

func (v NullableName) IsSet() bool {
	return v.isSet
}

func (v *NullableName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableName(val *Name) *NullableName {
	return &NullableName{value: val, isSet: true}
}

func (v NullableName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

