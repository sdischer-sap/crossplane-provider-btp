/*
Entitlements Service

The Entitlements service provides REST APIs that manage the assignments of entitlements and quotas to subaccounts and directories.   Entitlements and their quota are automatically assigned to the global account when a customer order is fulfilled. Use the APIs in this service to manage the distribution of this global quota to your directories and subaccounts.   NOTE: These APIs are relevant only for cloud management tools feature set B. For details and information about whether this applies to your global account, see [Cloud Management Tools - Feature Set Overview](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/Cloud/en-US/caf4e4e23aef4666ad8f125af393dfb2.html).  See also: * [Authorization](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/latest/en-US/3670474a58c24ac2b082e76cbbd9dc19.html) * [Rate Limiting](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/latest/en-US/77b217b3f57a45b987eb7fbc3305ce1e.html) * [Error Response Format](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/latest/en-US/77fef2fb104b4b1795e2e6cee790e8b8.html) * [Asynchronous Jobs](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/latest/en-US/0a0a6ab0ad114d72a6611c1c6b21683e.html)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the CommercialModelResponseObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommercialModelResponseObject{}

// CommercialModelResponseObject The commercial model type assigned to the entitlement.
type CommercialModelResponseObject struct {
	// Whether a customer pays only for services that they actually use (consumption-based) or pays for subscribed services at a fixed cost irrespective of consumption (subscription-based). <b>True:</b> Consumption-based commercial model.<b>False:</b> Subscription-based commercial model.
	ConsumptionBased *bool `json:"consumptionBased,omitempty"`
	// Directly contained commercial models.
	ContainedCommercialModels []CommercialModelResponseObject `json:"containedCommercialModels,omitempty"`
	// A description of the commercial model
	Description *string `json:"description,omitempty"`
	// A descriptive name of the commercial model for customer-facing UIs.
	DisplayName *string `json:"displayName,omitempty"`
	// Technical name of the commercial model.
	Name *string `json:"name,omitempty"`
}

// NewCommercialModelResponseObject instantiates a new CommercialModelResponseObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommercialModelResponseObject() *CommercialModelResponseObject {
	this := CommercialModelResponseObject{}
	return &this
}

// NewCommercialModelResponseObjectWithDefaults instantiates a new CommercialModelResponseObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommercialModelResponseObjectWithDefaults() *CommercialModelResponseObject {
	this := CommercialModelResponseObject{}
	return &this
}

// GetConsumptionBased returns the ConsumptionBased field value if set, zero value otherwise.
func (o *CommercialModelResponseObject) GetConsumptionBased() bool {
	if o == nil || IsNil(o.ConsumptionBased) {
		var ret bool
		return ret
	}
	return *o.ConsumptionBased
}

// GetConsumptionBasedOk returns a tuple with the ConsumptionBased field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommercialModelResponseObject) GetConsumptionBasedOk() (*bool, bool) {
	if o == nil || IsNil(o.ConsumptionBased) {
		return nil, false
	}
	return o.ConsumptionBased, true
}

// HasConsumptionBased returns a boolean if a field has been set.
func (o *CommercialModelResponseObject) HasConsumptionBased() bool {
	if o != nil && !IsNil(o.ConsumptionBased) {
		return true
	}

	return false
}

// SetConsumptionBased gets a reference to the given bool and assigns it to the ConsumptionBased field.
func (o *CommercialModelResponseObject) SetConsumptionBased(v bool) {
	o.ConsumptionBased = &v
}

// GetContainedCommercialModels returns the ContainedCommercialModels field value if set, zero value otherwise.
func (o *CommercialModelResponseObject) GetContainedCommercialModels() []CommercialModelResponseObject {
	if o == nil || IsNil(o.ContainedCommercialModels) {
		var ret []CommercialModelResponseObject
		return ret
	}
	return o.ContainedCommercialModels
}

// GetContainedCommercialModelsOk returns a tuple with the ContainedCommercialModels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommercialModelResponseObject) GetContainedCommercialModelsOk() ([]CommercialModelResponseObject, bool) {
	if o == nil || IsNil(o.ContainedCommercialModels) {
		return nil, false
	}
	return o.ContainedCommercialModels, true
}

// HasContainedCommercialModels returns a boolean if a field has been set.
func (o *CommercialModelResponseObject) HasContainedCommercialModels() bool {
	if o != nil && !IsNil(o.ContainedCommercialModels) {
		return true
	}

	return false
}

// SetContainedCommercialModels gets a reference to the given []CommercialModelResponseObject and assigns it to the ContainedCommercialModels field.
func (o *CommercialModelResponseObject) SetContainedCommercialModels(v []CommercialModelResponseObject) {
	o.ContainedCommercialModels = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CommercialModelResponseObject) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommercialModelResponseObject) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CommercialModelResponseObject) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CommercialModelResponseObject) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *CommercialModelResponseObject) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommercialModelResponseObject) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *CommercialModelResponseObject) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *CommercialModelResponseObject) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CommercialModelResponseObject) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommercialModelResponseObject) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CommercialModelResponseObject) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CommercialModelResponseObject) SetName(v string) {
	o.Name = &v
}

func (o CommercialModelResponseObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommercialModelResponseObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ConsumptionBased) {
		toSerialize["consumptionBased"] = o.ConsumptionBased
	}
	if !IsNil(o.ContainedCommercialModels) {
		toSerialize["containedCommercialModels"] = o.ContainedCommercialModels
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableCommercialModelResponseObject struct {
	value *CommercialModelResponseObject
	isSet bool
}

func (v NullableCommercialModelResponseObject) Get() *CommercialModelResponseObject {
	return v.value
}

func (v *NullableCommercialModelResponseObject) Set(val *CommercialModelResponseObject) {
	v.value = val
	v.isSet = true
}

func (v NullableCommercialModelResponseObject) IsSet() bool {
	return v.isSet
}

func (v *NullableCommercialModelResponseObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommercialModelResponseObject(val *CommercialModelResponseObject) *NullableCommercialModelResponseObject {
	return &NullableCommercialModelResponseObject{value: val, isSet: true}
}

func (v NullableCommercialModelResponseObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommercialModelResponseObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

