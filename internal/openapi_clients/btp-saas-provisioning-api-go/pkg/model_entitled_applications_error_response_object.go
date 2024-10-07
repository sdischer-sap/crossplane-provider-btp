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

// checks if the EntitledApplicationsErrorResponseObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitledApplicationsErrorResponseObject{}

// EntitledApplicationsErrorResponseObject The response object that contains details about an error in case the subscription failed.
type EntitledApplicationsErrorResponseObject struct {
	// A response object that contains details about the error an app provider returns to the subscriber. It contains the error code, a user-friendly, customer-oriented error message, technical details about the error, and more.
	AppError *string `json:"appError,omitempty"`
	// The message that describes the error that occurred during the subscription. 
	ErrorMessage *string `json:"errorMessage,omitempty"`
}

// NewEntitledApplicationsErrorResponseObject instantiates a new EntitledApplicationsErrorResponseObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitledApplicationsErrorResponseObject() *EntitledApplicationsErrorResponseObject {
	this := EntitledApplicationsErrorResponseObject{}
	return &this
}

// NewEntitledApplicationsErrorResponseObjectWithDefaults instantiates a new EntitledApplicationsErrorResponseObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitledApplicationsErrorResponseObjectWithDefaults() *EntitledApplicationsErrorResponseObject {
	this := EntitledApplicationsErrorResponseObject{}
	return &this
}

// GetAppError returns the AppError field value if set, zero value otherwise.
func (o *EntitledApplicationsErrorResponseObject) GetAppError() string {
	if o == nil || IsNil(o.AppError) {
		var ret string
		return ret
	}
	return *o.AppError
}

// GetAppErrorOk returns a tuple with the AppError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitledApplicationsErrorResponseObject) GetAppErrorOk() (*string, bool) {
	if o == nil || IsNil(o.AppError) {
		return nil, false
	}
	return o.AppError, true
}

// HasAppError returns a boolean if a field has been set.
func (o *EntitledApplicationsErrorResponseObject) HasAppError() bool {
	if o != nil && !IsNil(o.AppError) {
		return true
	}

	return false
}

// SetAppError gets a reference to the given string and assigns it to the AppError field.
func (o *EntitledApplicationsErrorResponseObject) SetAppError(v string) {
	o.AppError = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *EntitledApplicationsErrorResponseObject) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitledApplicationsErrorResponseObject) GetErrorMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *EntitledApplicationsErrorResponseObject) HasErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *EntitledApplicationsErrorResponseObject) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

func (o EntitledApplicationsErrorResponseObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitledApplicationsErrorResponseObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppError) {
		toSerialize["appError"] = o.AppError
	}
	if !IsNil(o.ErrorMessage) {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	return toSerialize, nil
}

type NullableEntitledApplicationsErrorResponseObject struct {
	value *EntitledApplicationsErrorResponseObject
	isSet bool
}

func (v NullableEntitledApplicationsErrorResponseObject) Get() *EntitledApplicationsErrorResponseObject {
	return v.value
}

func (v *NullableEntitledApplicationsErrorResponseObject) Set(val *EntitledApplicationsErrorResponseObject) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitledApplicationsErrorResponseObject) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitledApplicationsErrorResponseObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitledApplicationsErrorResponseObject(val *EntitledApplicationsErrorResponseObject) *NullableEntitledApplicationsErrorResponseObject {
	return &NullableEntitledApplicationsErrorResponseObject{value: val, isSet: true}
}

func (v NullableEntitledApplicationsErrorResponseObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitledApplicationsErrorResponseObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

