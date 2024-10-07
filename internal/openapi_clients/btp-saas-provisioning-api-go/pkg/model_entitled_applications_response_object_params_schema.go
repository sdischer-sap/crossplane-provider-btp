/*
SaaS Provisioning Service

The SAP SaaS Provisioning service provides REST APIs that are responsible for the registration and provisioning of multitenant applications and services.   Use the APIs in this service to perform various operations related to your multitenant applications and services. For example, to get application registration details, subscribe a tenant to your application, unsubscribe a tenant from your application, retrieve all your application subscriptions, update subscription dependencies, and to get subscription job information.  See also: * [Authorization](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/latest/en-US/3670474a58c24ac2b082e76cbbd9dc19.html) * [Rate Limiting](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/latest/en-US/77b217b3f57a45b987eb7fbc3305ce1e.html) * [Error Response Format](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/latest/en-US/77fef2fb104b4b1795e2e6cee790e8b8.html) * [Asynchronous Jobs](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/latest/en-US/0a0a6ab0ad114d72a6611c1c6b21683e.html)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the EntitledApplicationsResponseObjectParamsSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitledApplicationsResponseObjectParamsSchema{}

// EntitledApplicationsResponseObjectParamsSchema JSON schema that describes the parameters that the consumer must provide when subscribing to the application.
type EntitledApplicationsResponseObjectParamsSchema struct {
	Empty *bool `json:"empty,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EntitledApplicationsResponseObjectParamsSchema EntitledApplicationsResponseObjectParamsSchema

// NewEntitledApplicationsResponseObjectParamsSchema instantiates a new EntitledApplicationsResponseObjectParamsSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitledApplicationsResponseObjectParamsSchema() *EntitledApplicationsResponseObjectParamsSchema {
	this := EntitledApplicationsResponseObjectParamsSchema{}
	return &this
}

// NewEntitledApplicationsResponseObjectParamsSchemaWithDefaults instantiates a new EntitledApplicationsResponseObjectParamsSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitledApplicationsResponseObjectParamsSchemaWithDefaults() *EntitledApplicationsResponseObjectParamsSchema {
	this := EntitledApplicationsResponseObjectParamsSchema{}
	return &this
}

// GetEmpty returns the Empty field value if set, zero value otherwise.
func (o *EntitledApplicationsResponseObjectParamsSchema) GetEmpty() bool {
	if o == nil || IsNil(o.Empty) {
		var ret bool
		return ret
	}
	return *o.Empty
}

// GetEmptyOk returns a tuple with the Empty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitledApplicationsResponseObjectParamsSchema) GetEmptyOk() (*bool, bool) {
	if o == nil || IsNil(o.Empty) {
		return nil, false
	}
	return o.Empty, true
}

// HasEmpty returns a boolean if a field has been set.
func (o *EntitledApplicationsResponseObjectParamsSchema) HasEmpty() bool {
	if o != nil && !IsNil(o.Empty) {
		return true
	}

	return false
}

// SetEmpty gets a reference to the given bool and assigns it to the Empty field.
func (o *EntitledApplicationsResponseObjectParamsSchema) SetEmpty(v bool) {
	o.Empty = &v
}

func (o EntitledApplicationsResponseObjectParamsSchema) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitledApplicationsResponseObjectParamsSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Empty) {
		toSerialize["empty"] = o.Empty
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EntitledApplicationsResponseObjectParamsSchema) UnmarshalJSON(data []byte) (err error) {
	varEntitledApplicationsResponseObjectParamsSchema := _EntitledApplicationsResponseObjectParamsSchema{}

	err = json.Unmarshal(data, &varEntitledApplicationsResponseObjectParamsSchema)

	if err != nil {
		return err
	}

	*o = EntitledApplicationsResponseObjectParamsSchema(varEntitledApplicationsResponseObjectParamsSchema)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "empty")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntitledApplicationsResponseObjectParamsSchema struct {
	value *EntitledApplicationsResponseObjectParamsSchema
	isSet bool
}

func (v NullableEntitledApplicationsResponseObjectParamsSchema) Get() *EntitledApplicationsResponseObjectParamsSchema {
	return v.value
}

func (v *NullableEntitledApplicationsResponseObjectParamsSchema) Set(val *EntitledApplicationsResponseObjectParamsSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitledApplicationsResponseObjectParamsSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitledApplicationsResponseObjectParamsSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitledApplicationsResponseObjectParamsSchema(val *EntitledApplicationsResponseObjectParamsSchema) *NullableEntitledApplicationsResponseObjectParamsSchema {
	return &NullableEntitledApplicationsResponseObjectParamsSchema{value: val, isSet: true}
}

func (v NullableEntitledApplicationsResponseObjectParamsSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitledApplicationsResponseObjectParamsSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


