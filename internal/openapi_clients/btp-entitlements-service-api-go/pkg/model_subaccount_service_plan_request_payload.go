/*
Entitlements Service

The Entitlements service provides REST APIs that manage the assignments of entitlements and quotas to subaccounts and directories.   Entitlements and their quota are automatically assigned to the global account when a customer order is fulfilled. Use the APIs in this service to manage the distribution of this global quota to your directories and subaccounts.   NOTE: These APIs are relevant only for cloud management tools feature set B. For details and information about whether this applies to your global account, see [Cloud Management Tools - Feature Set Overview](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/Cloud/en-US/caf4e4e23aef4666ad8f125af393dfb2.html).  See also: * [Authorization](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/latest/en-US/3670474a58c24ac2b082e76cbbd9dc19.html) * [Rate Limiting](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/latest/en-US/77b217b3f57a45b987eb7fbc3305ce1e.html) * [Error Response Format](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/latest/en-US/77fef2fb104b4b1795e2e6cee790e8b8.html) * [Asynchronous Jobs](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/latest/en-US/0a0a6ab0ad114d72a6611c1c6b21683e.html)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the SubaccountServicePlanRequestPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubaccountServicePlanRequestPayload{}

// SubaccountServicePlanRequestPayload List of assigned entitlements and their specifications.
type SubaccountServicePlanRequestPayload struct {
	// The quantity of the plan that is assigned to the specified subaccount. Relevant and mandatory only for plans that have a numeric quota. Do not set if enable=TRUE is specified.
	Amount *float32 `json:"amount,omitempty"`
	// Whether to enable the service plan assignment to the specified subaccount without quantity restrictions. Relevant and mandatory only for plans that do not have a numeric quota. Do not set if amount is specified.
	Enable *bool `json:"enable,omitempty"`
	// External resources to assign to subaccount
	Resources []ExternalResourceRequestPayload `json:"resources,omitempty"`
	// The unique ID of the subaccount to which to assign a service plan.
	SubaccountGUID string `json:"subaccountGUID"`
}

type _SubaccountServicePlanRequestPayload SubaccountServicePlanRequestPayload

// NewSubaccountServicePlanRequestPayload instantiates a new SubaccountServicePlanRequestPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubaccountServicePlanRequestPayload(subaccountGUID string) *SubaccountServicePlanRequestPayload {
	this := SubaccountServicePlanRequestPayload{}
	this.SubaccountGUID = subaccountGUID
	return &this
}

// NewSubaccountServicePlanRequestPayloadWithDefaults instantiates a new SubaccountServicePlanRequestPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubaccountServicePlanRequestPayloadWithDefaults() *SubaccountServicePlanRequestPayload {
	this := SubaccountServicePlanRequestPayload{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *SubaccountServicePlanRequestPayload) GetAmount() float32 {
	if o == nil || IsNil(o.Amount) {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubaccountServicePlanRequestPayload) GetAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *SubaccountServicePlanRequestPayload) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *SubaccountServicePlanRequestPayload) SetAmount(v float32) {
	o.Amount = &v
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *SubaccountServicePlanRequestPayload) GetEnable() bool {
	if o == nil || IsNil(o.Enable) {
		var ret bool
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubaccountServicePlanRequestPayload) GetEnableOk() (*bool, bool) {
	if o == nil || IsNil(o.Enable) {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *SubaccountServicePlanRequestPayload) HasEnable() bool {
	if o != nil && !IsNil(o.Enable) {
		return true
	}

	return false
}

// SetEnable gets a reference to the given bool and assigns it to the Enable field.
func (o *SubaccountServicePlanRequestPayload) SetEnable(v bool) {
	o.Enable = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *SubaccountServicePlanRequestPayload) GetResources() []ExternalResourceRequestPayload {
	if o == nil || IsNil(o.Resources) {
		var ret []ExternalResourceRequestPayload
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubaccountServicePlanRequestPayload) GetResourcesOk() ([]ExternalResourceRequestPayload, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *SubaccountServicePlanRequestPayload) HasResources() bool {
	if o != nil && !IsNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given []ExternalResourceRequestPayload and assigns it to the Resources field.
func (o *SubaccountServicePlanRequestPayload) SetResources(v []ExternalResourceRequestPayload) {
	o.Resources = v
}

// GetSubaccountGUID returns the SubaccountGUID field value
func (o *SubaccountServicePlanRequestPayload) GetSubaccountGUID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubaccountGUID
}

// GetSubaccountGUIDOk returns a tuple with the SubaccountGUID field value
// and a boolean to check if the value has been set.
func (o *SubaccountServicePlanRequestPayload) GetSubaccountGUIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubaccountGUID, true
}

// SetSubaccountGUID sets field value
func (o *SubaccountServicePlanRequestPayload) SetSubaccountGUID(v string) {
	o.SubaccountGUID = v
}

func (o SubaccountServicePlanRequestPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubaccountServicePlanRequestPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.Enable) {
		toSerialize["enable"] = o.Enable
	}
	if !IsNil(o.Resources) {
		toSerialize["resources"] = o.Resources
	}
	toSerialize["subaccountGUID"] = o.SubaccountGUID
	return toSerialize, nil
}

func (o *SubaccountServicePlanRequestPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"subaccountGUID",
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

	varSubaccountServicePlanRequestPayload := _SubaccountServicePlanRequestPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSubaccountServicePlanRequestPayload)

	if err != nil {
		return err
	}

	*o = SubaccountServicePlanRequestPayload(varSubaccountServicePlanRequestPayload)

	return err
}

type NullableSubaccountServicePlanRequestPayload struct {
	value *SubaccountServicePlanRequestPayload
	isSet bool
}

func (v NullableSubaccountServicePlanRequestPayload) Get() *SubaccountServicePlanRequestPayload {
	return v.value
}

func (v *NullableSubaccountServicePlanRequestPayload) Set(val *SubaccountServicePlanRequestPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableSubaccountServicePlanRequestPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableSubaccountServicePlanRequestPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubaccountServicePlanRequestPayload(val *SubaccountServicePlanRequestPayload) *NullableSubaccountServicePlanRequestPayload {
	return &NullableSubaccountServicePlanRequestPayload{value: val, isSet: true}
}

func (v NullableSubaccountServicePlanRequestPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubaccountServicePlanRequestPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


