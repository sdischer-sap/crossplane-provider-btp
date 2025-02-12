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

// checks if the EntitlementAmountResponseObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitlementAmountResponseObject{}

// EntitlementAmountResponseObject Relevant entitlements for the source that added the product.
type EntitlementAmountResponseObject struct {
	// The quantity of the entitlement that is assigned to the root global account or directory.
	Amount *float32 `json:"amount,omitempty"`
	// Specifies if a plan associated with this entitlement will be automatically assigned by the system to any new subaccount. For example, free plans that are available to all subaccounts.
	AutoAssign *bool `json:"autoAssign,omitempty"`
	CommercialModel *CommercialModelResponseObject `json:"commercialModel,omitempty"`
	// The technical name of the product.
	EntitlementName *string `json:"entitlementName,omitempty"`
	// The unique ID of the material (product) to which the assigned service plan belongs.
	ProductId *string `json:"productId,omitempty"`
}

// NewEntitlementAmountResponseObject instantiates a new EntitlementAmountResponseObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementAmountResponseObject() *EntitlementAmountResponseObject {
	this := EntitlementAmountResponseObject{}
	return &this
}

// NewEntitlementAmountResponseObjectWithDefaults instantiates a new EntitlementAmountResponseObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementAmountResponseObjectWithDefaults() *EntitlementAmountResponseObject {
	this := EntitlementAmountResponseObject{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *EntitlementAmountResponseObject) GetAmount() float32 {
	if o == nil || IsNil(o.Amount) {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementAmountResponseObject) GetAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *EntitlementAmountResponseObject) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *EntitlementAmountResponseObject) SetAmount(v float32) {
	o.Amount = &v
}

// GetAutoAssign returns the AutoAssign field value if set, zero value otherwise.
func (o *EntitlementAmountResponseObject) GetAutoAssign() bool {
	if o == nil || IsNil(o.AutoAssign) {
		var ret bool
		return ret
	}
	return *o.AutoAssign
}

// GetAutoAssignOk returns a tuple with the AutoAssign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementAmountResponseObject) GetAutoAssignOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoAssign) {
		return nil, false
	}
	return o.AutoAssign, true
}

// HasAutoAssign returns a boolean if a field has been set.
func (o *EntitlementAmountResponseObject) HasAutoAssign() bool {
	if o != nil && !IsNil(o.AutoAssign) {
		return true
	}

	return false
}

// SetAutoAssign gets a reference to the given bool and assigns it to the AutoAssign field.
func (o *EntitlementAmountResponseObject) SetAutoAssign(v bool) {
	o.AutoAssign = &v
}

// GetCommercialModel returns the CommercialModel field value if set, zero value otherwise.
func (o *EntitlementAmountResponseObject) GetCommercialModel() CommercialModelResponseObject {
	if o == nil || IsNil(o.CommercialModel) {
		var ret CommercialModelResponseObject
		return ret
	}
	return *o.CommercialModel
}

// GetCommercialModelOk returns a tuple with the CommercialModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementAmountResponseObject) GetCommercialModelOk() (*CommercialModelResponseObject, bool) {
	if o == nil || IsNil(o.CommercialModel) {
		return nil, false
	}
	return o.CommercialModel, true
}

// HasCommercialModel returns a boolean if a field has been set.
func (o *EntitlementAmountResponseObject) HasCommercialModel() bool {
	if o != nil && !IsNil(o.CommercialModel) {
		return true
	}

	return false
}

// SetCommercialModel gets a reference to the given CommercialModelResponseObject and assigns it to the CommercialModel field.
func (o *EntitlementAmountResponseObject) SetCommercialModel(v CommercialModelResponseObject) {
	o.CommercialModel = &v
}

// GetEntitlementName returns the EntitlementName field value if set, zero value otherwise.
func (o *EntitlementAmountResponseObject) GetEntitlementName() string {
	if o == nil || IsNil(o.EntitlementName) {
		var ret string
		return ret
	}
	return *o.EntitlementName
}

// GetEntitlementNameOk returns a tuple with the EntitlementName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementAmountResponseObject) GetEntitlementNameOk() (*string, bool) {
	if o == nil || IsNil(o.EntitlementName) {
		return nil, false
	}
	return o.EntitlementName, true
}

// HasEntitlementName returns a boolean if a field has been set.
func (o *EntitlementAmountResponseObject) HasEntitlementName() bool {
	if o != nil && !IsNil(o.EntitlementName) {
		return true
	}

	return false
}

// SetEntitlementName gets a reference to the given string and assigns it to the EntitlementName field.
func (o *EntitlementAmountResponseObject) SetEntitlementName(v string) {
	o.EntitlementName = &v
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *EntitlementAmountResponseObject) GetProductId() string {
	if o == nil || IsNil(o.ProductId) {
		var ret string
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementAmountResponseObject) GetProductIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProductId) {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *EntitlementAmountResponseObject) HasProductId() bool {
	if o != nil && !IsNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given string and assigns it to the ProductId field.
func (o *EntitlementAmountResponseObject) SetProductId(v string) {
	o.ProductId = &v
}

func (o EntitlementAmountResponseObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitlementAmountResponseObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.AutoAssign) {
		toSerialize["autoAssign"] = o.AutoAssign
	}
	if !IsNil(o.CommercialModel) {
		toSerialize["commercialModel"] = o.CommercialModel
	}
	if !IsNil(o.EntitlementName) {
		toSerialize["entitlementName"] = o.EntitlementName
	}
	if !IsNil(o.ProductId) {
		toSerialize["productId"] = o.ProductId
	}
	return toSerialize, nil
}

type NullableEntitlementAmountResponseObject struct {
	value *EntitlementAmountResponseObject
	isSet bool
}

func (v NullableEntitlementAmountResponseObject) Get() *EntitlementAmountResponseObject {
	return v.value
}

func (v *NullableEntitlementAmountResponseObject) Set(val *EntitlementAmountResponseObject) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementAmountResponseObject) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementAmountResponseObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementAmountResponseObject(val *EntitlementAmountResponseObject) *NullableEntitlementAmountResponseObject {
	return &NullableEntitlementAmountResponseObject{value: val, isSet: true}
}

func (v NullableEntitlementAmountResponseObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementAmountResponseObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


