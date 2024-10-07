/*
Provisioning Service

The Provisioning service provides REST APIs that are responsible for the provisioning and deprovisioning of environment instances and tenants in the corresponding region.  Provisioning is performed after validation by the Entitlements service. Use the APIs in this service to manage and create environment instances, such as a Cloud Foundry org, in a subaccount and to retrieve the plans and quota assignments for a subaccount.  NOTE: These APIs are relevant only for cloud management tools feature set B. For details and information about whether this applies to your global account, see [Cloud Management Tools - Feature Set Overview](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/Cloud/en-US/caf4e4e23aef4666ad8f125af393dfb2.html).  See also: * [Authorization](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/latest/en-US/3670474a58c24ac2b082e76cbbd9dc19.html) * [Rate Limiting](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/latest/en-US/77b217b3f57a45b987eb7fbc3305ce1e.html) * [Error Response Format](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/latest/en-US/77fef2fb104b4b1795e2e6cee790e8b8.html) * [Asynchronous Jobs](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/latest/en-US/0a0a6ab0ad114d72a6611c1c6b21683e.html)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the CreatedEnvironmentInstanceResponseObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreatedEnvironmentInstanceResponseObject{}

// CreatedEnvironmentInstanceResponseObject struct for CreatedEnvironmentInstanceResponseObject
type CreatedEnvironmentInstanceResponseObject struct {
	// ID of the created environment instance
	Id *string `json:"id,omitempty"`
}

// NewCreatedEnvironmentInstanceResponseObject instantiates a new CreatedEnvironmentInstanceResponseObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatedEnvironmentInstanceResponseObject() *CreatedEnvironmentInstanceResponseObject {
	this := CreatedEnvironmentInstanceResponseObject{}
	return &this
}

// NewCreatedEnvironmentInstanceResponseObjectWithDefaults instantiates a new CreatedEnvironmentInstanceResponseObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatedEnvironmentInstanceResponseObjectWithDefaults() *CreatedEnvironmentInstanceResponseObject {
	this := CreatedEnvironmentInstanceResponseObject{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreatedEnvironmentInstanceResponseObject) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatedEnvironmentInstanceResponseObject) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreatedEnvironmentInstanceResponseObject) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CreatedEnvironmentInstanceResponseObject) SetId(v string) {
	o.Id = &v
}

func (o CreatedEnvironmentInstanceResponseObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreatedEnvironmentInstanceResponseObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableCreatedEnvironmentInstanceResponseObject struct {
	value *CreatedEnvironmentInstanceResponseObject
	isSet bool
}

func (v NullableCreatedEnvironmentInstanceResponseObject) Get() *CreatedEnvironmentInstanceResponseObject {
	return v.value
}

func (v *NullableCreatedEnvironmentInstanceResponseObject) Set(val *CreatedEnvironmentInstanceResponseObject) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatedEnvironmentInstanceResponseObject) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatedEnvironmentInstanceResponseObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatedEnvironmentInstanceResponseObject(val *CreatedEnvironmentInstanceResponseObject) *NullableCreatedEnvironmentInstanceResponseObject {
	return &NullableCreatedEnvironmentInstanceResponseObject{value: val, isSet: true}
}

func (v NullableCreatedEnvironmentInstanceResponseObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatedEnvironmentInstanceResponseObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

