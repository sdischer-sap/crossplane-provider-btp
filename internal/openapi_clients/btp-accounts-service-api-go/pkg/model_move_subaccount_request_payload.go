/*
Accounts Service

The Accounts service provides REST APIs that are responsible for the management of global accounts, and the creation and management of directories, subaccounts, and their custom properties/tags.  Global accounts represent a business entity and contain contract information, including customer details and purchased entitlements. The global account is the context for billing each customer.  Use the subaccount APIs to structure your global account according to your organization's and project's requirements regarding members, authorizations, and quotas. This service also provides you with APIs for creating and managing directories. While the use of directories is optional, they allow you to further organize and manage your subaccounts according to your specific technical and business needs. The service also lets you manage the custom properties/tags that you associate with your directories and subaccounts. NOTE: These APIs are relevant only for cloud management tools feature set B. For details and information about whether this applies to your global account, see [Cloud Management Tools - Feature Set Overview](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/Cloud/en-US/caf4e4e23aef4666ad8f125af393dfb2.html).  See also: * [Authorization](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/latest/en-US/3670474a58c24ac2b082e76cbbd9dc19.html) * [Rate Limiting](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/latest/en-US/77b217b3f57a45b987eb7fbc3305ce1e.html) * [Error Response Format](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/latest/en-US/77fef2fb104b4b1795e2e6cee790e8b8.html) * [Asynchronous Jobs](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/latest/en-US/0a0a6ab0ad114d72a6611c1c6b21683e.html)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the MoveSubaccountRequestPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MoveSubaccountRequestPayload{}

// MoveSubaccountRequestPayload Details of where to move the subaccount to.
type MoveSubaccountRequestPayload struct {
	// The GUID of the new location of the subaccount. To move to a directory, enter the GUID of the directory. To move out of a directory to the root global account, enter the GUID of the global account.
	TargetAccountGUID string `json:"targetAccountGUID"`
}

type _MoveSubaccountRequestPayload MoveSubaccountRequestPayload

// NewMoveSubaccountRequestPayload instantiates a new MoveSubaccountRequestPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMoveSubaccountRequestPayload(targetAccountGUID string) *MoveSubaccountRequestPayload {
	this := MoveSubaccountRequestPayload{}
	this.TargetAccountGUID = targetAccountGUID
	return &this
}

// NewMoveSubaccountRequestPayloadWithDefaults instantiates a new MoveSubaccountRequestPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMoveSubaccountRequestPayloadWithDefaults() *MoveSubaccountRequestPayload {
	this := MoveSubaccountRequestPayload{}
	return &this
}

// GetTargetAccountGUID returns the TargetAccountGUID field value
func (o *MoveSubaccountRequestPayload) GetTargetAccountGUID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetAccountGUID
}

// GetTargetAccountGUIDOk returns a tuple with the TargetAccountGUID field value
// and a boolean to check if the value has been set.
func (o *MoveSubaccountRequestPayload) GetTargetAccountGUIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetAccountGUID, true
}

// SetTargetAccountGUID sets field value
func (o *MoveSubaccountRequestPayload) SetTargetAccountGUID(v string) {
	o.TargetAccountGUID = v
}

func (o MoveSubaccountRequestPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MoveSubaccountRequestPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["targetAccountGUID"] = o.TargetAccountGUID
	return toSerialize, nil
}

func (o *MoveSubaccountRequestPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"targetAccountGUID",
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

	varMoveSubaccountRequestPayload := _MoveSubaccountRequestPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMoveSubaccountRequestPayload)

	if err != nil {
		return err
	}

	*o = MoveSubaccountRequestPayload(varMoveSubaccountRequestPayload)

	return err
}

type NullableMoveSubaccountRequestPayload struct {
	value *MoveSubaccountRequestPayload
	isSet bool
}

func (v NullableMoveSubaccountRequestPayload) Get() *MoveSubaccountRequestPayload {
	return v.value
}

func (v *NullableMoveSubaccountRequestPayload) Set(val *MoveSubaccountRequestPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableMoveSubaccountRequestPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableMoveSubaccountRequestPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMoveSubaccountRequestPayload(val *MoveSubaccountRequestPayload) *NullableMoveSubaccountRequestPayload {
	return &NullableMoveSubaccountRequestPayload{value: val, isSet: true}
}

func (v NullableMoveSubaccountRequestPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMoveSubaccountRequestPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

