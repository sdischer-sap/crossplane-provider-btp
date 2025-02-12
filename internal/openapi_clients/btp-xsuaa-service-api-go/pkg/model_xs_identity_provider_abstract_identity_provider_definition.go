/*
SAP XSUAA REST API

Provides access to RoleTemplates, Roles, RoleCollection etc. using the XSUAA REST API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the XSIdentityProviderAbstractIdentityProviderDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &XSIdentityProviderAbstractIdentityProviderDefinition{}

// XSIdentityProviderAbstractIdentityProviderDefinition struct for XSIdentityProviderAbstractIdentityProviderDefinition
type XSIdentityProviderAbstractIdentityProviderDefinition struct {
	Id *string `json:"id,omitempty"`
	OriginKey string `json:"originKey"`
	Name string `json:"name"`
	Type string `json:"type"`
	Config *AbstractIdentityProviderDefinition `json:"config,omitempty"`
	Version *int32 `json:"version,omitempty"`
	Created *time.Time `json:"created,omitempty"`
	Active *bool `json:"active,omitempty"`
	IdentityZoneId *string `json:"identityZoneId,omitempty"`
	LastModified *time.Time `json:"last_modified,omitempty"`
}

type _XSIdentityProviderAbstractIdentityProviderDefinition XSIdentityProviderAbstractIdentityProviderDefinition

// NewXSIdentityProviderAbstractIdentityProviderDefinition instantiates a new XSIdentityProviderAbstractIdentityProviderDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewXSIdentityProviderAbstractIdentityProviderDefinition(originKey string, name string, type_ string) *XSIdentityProviderAbstractIdentityProviderDefinition {
	this := XSIdentityProviderAbstractIdentityProviderDefinition{}
	this.OriginKey = originKey
	this.Name = name
	this.Type = type_
	return &this
}

// NewXSIdentityProviderAbstractIdentityProviderDefinitionWithDefaults instantiates a new XSIdentityProviderAbstractIdentityProviderDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewXSIdentityProviderAbstractIdentityProviderDefinitionWithDefaults() *XSIdentityProviderAbstractIdentityProviderDefinition {
	this := XSIdentityProviderAbstractIdentityProviderDefinition{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *XSIdentityProviderAbstractIdentityProviderDefinition) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XSIdentityProviderAbstractIdentityProviderDefinition) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *XSIdentityProviderAbstractIdentityProviderDefinition) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *XSIdentityProviderAbstractIdentityProviderDefinition) SetId(v string) {
	o.Id = &v
}

// GetOriginKey returns the OriginKey field value
func (o *XSIdentityProviderAbstractIdentityProviderDefinition) GetOriginKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OriginKey
}

// GetOriginKeyOk returns a tuple with the OriginKey field value
// and a boolean to check if the value has been set.
func (o *XSIdentityProviderAbstractIdentityProviderDefinition) GetOriginKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OriginKey, true
}

// SetOriginKey sets field value
func (o *XSIdentityProviderAbstractIdentityProviderDefinition) SetOriginKey(v string) {
	o.OriginKey = v
}

// GetName returns the Name field value
func (o *XSIdentityProviderAbstractIdentityProviderDefinition) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *XSIdentityProviderAbstractIdentityProviderDefinition) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *XSIdentityProviderAbstractIdentityProviderDefinition) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *XSIdentityProviderAbstractIdentityProviderDefinition) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *XSIdentityProviderAbstractIdentityProviderDefinition) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *XSIdentityProviderAbstractIdentityProviderDefinition) SetType(v string) {
	o.Type = v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *XSIdentityProviderAbstractIdentityProviderDefinition) GetConfig() AbstractIdentityProviderDefinition {
	if o == nil || IsNil(o.Config) {
		var ret AbstractIdentityProviderDefinition
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XSIdentityProviderAbstractIdentityProviderDefinition) GetConfigOk() (*AbstractIdentityProviderDefinition, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *XSIdentityProviderAbstractIdentityProviderDefinition) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given AbstractIdentityProviderDefinition and assigns it to the Config field.
func (o *XSIdentityProviderAbstractIdentityProviderDefinition) SetConfig(v AbstractIdentityProviderDefinition) {
	o.Config = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *XSIdentityProviderAbstractIdentityProviderDefinition) GetVersion() int32 {
	if o == nil || IsNil(o.Version) {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XSIdentityProviderAbstractIdentityProviderDefinition) GetVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *XSIdentityProviderAbstractIdentityProviderDefinition) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *XSIdentityProviderAbstractIdentityProviderDefinition) SetVersion(v int32) {
	o.Version = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *XSIdentityProviderAbstractIdentityProviderDefinition) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XSIdentityProviderAbstractIdentityProviderDefinition) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *XSIdentityProviderAbstractIdentityProviderDefinition) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *XSIdentityProviderAbstractIdentityProviderDefinition) SetCreated(v time.Time) {
	o.Created = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *XSIdentityProviderAbstractIdentityProviderDefinition) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XSIdentityProviderAbstractIdentityProviderDefinition) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *XSIdentityProviderAbstractIdentityProviderDefinition) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *XSIdentityProviderAbstractIdentityProviderDefinition) SetActive(v bool) {
	o.Active = &v
}

// GetIdentityZoneId returns the IdentityZoneId field value if set, zero value otherwise.
func (o *XSIdentityProviderAbstractIdentityProviderDefinition) GetIdentityZoneId() string {
	if o == nil || IsNil(o.IdentityZoneId) {
		var ret string
		return ret
	}
	return *o.IdentityZoneId
}

// GetIdentityZoneIdOk returns a tuple with the IdentityZoneId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XSIdentityProviderAbstractIdentityProviderDefinition) GetIdentityZoneIdOk() (*string, bool) {
	if o == nil || IsNil(o.IdentityZoneId) {
		return nil, false
	}
	return o.IdentityZoneId, true
}

// HasIdentityZoneId returns a boolean if a field has been set.
func (o *XSIdentityProviderAbstractIdentityProviderDefinition) HasIdentityZoneId() bool {
	if o != nil && !IsNil(o.IdentityZoneId) {
		return true
	}

	return false
}

// SetIdentityZoneId gets a reference to the given string and assigns it to the IdentityZoneId field.
func (o *XSIdentityProviderAbstractIdentityProviderDefinition) SetIdentityZoneId(v string) {
	o.IdentityZoneId = &v
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *XSIdentityProviderAbstractIdentityProviderDefinition) GetLastModified() time.Time {
	if o == nil || IsNil(o.LastModified) {
		var ret time.Time
		return ret
	}
	return *o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XSIdentityProviderAbstractIdentityProviderDefinition) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModified) {
		return nil, false
	}
	return o.LastModified, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *XSIdentityProviderAbstractIdentityProviderDefinition) HasLastModified() bool {
	if o != nil && !IsNil(o.LastModified) {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given time.Time and assigns it to the LastModified field.
func (o *XSIdentityProviderAbstractIdentityProviderDefinition) SetLastModified(v time.Time) {
	o.LastModified = &v
}

func (o XSIdentityProviderAbstractIdentityProviderDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o XSIdentityProviderAbstractIdentityProviderDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["originKey"] = o.OriginKey
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.IdentityZoneId) {
		toSerialize["identityZoneId"] = o.IdentityZoneId
	}
	if !IsNil(o.LastModified) {
		toSerialize["last_modified"] = o.LastModified
	}
	return toSerialize, nil
}

func (o *XSIdentityProviderAbstractIdentityProviderDefinition) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"originKey",
		"name",
		"type",
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

	varXSIdentityProviderAbstractIdentityProviderDefinition := _XSIdentityProviderAbstractIdentityProviderDefinition{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varXSIdentityProviderAbstractIdentityProviderDefinition)

	if err != nil {
		return err
	}

	*o = XSIdentityProviderAbstractIdentityProviderDefinition(varXSIdentityProviderAbstractIdentityProviderDefinition)

	return err
}

type NullableXSIdentityProviderAbstractIdentityProviderDefinition struct {
	value *XSIdentityProviderAbstractIdentityProviderDefinition
	isSet bool
}

func (v NullableXSIdentityProviderAbstractIdentityProviderDefinition) Get() *XSIdentityProviderAbstractIdentityProviderDefinition {
	return v.value
}

func (v *NullableXSIdentityProviderAbstractIdentityProviderDefinition) Set(val *XSIdentityProviderAbstractIdentityProviderDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableXSIdentityProviderAbstractIdentityProviderDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableXSIdentityProviderAbstractIdentityProviderDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableXSIdentityProviderAbstractIdentityProviderDefinition(val *XSIdentityProviderAbstractIdentityProviderDefinition) *NullableXSIdentityProviderAbstractIdentityProviderDefinition {
	return &NullableXSIdentityProviderAbstractIdentityProviderDefinition{value: val, isSet: true}
}

func (v NullableXSIdentityProviderAbstractIdentityProviderDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableXSIdentityProviderAbstractIdentityProviderDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


