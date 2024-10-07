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

// checks if the SaasAsyncSubscriptionRequestPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SaasAsyncSubscriptionRequestPayload{}

// SaasAsyncSubscriptionRequestPayload The details of the subaccount tenant subscription an app provider sends to users.
type SaasAsyncSubscriptionRequestPayload struct {
	// Additional details accompanying the subscription process. Relates mostly to the cases when the subscription process status is FAILED. 
	Message *string `json:"message,omitempty"`
	// Status of the subscription job.
	Status *string `json:"status,omitempty"`
	// The URL the multitenant application is exposing for a subscription.
	SubscriptionUrl *string `json:"subscriptionUrl,omitempty"`
}

// NewSaasAsyncSubscriptionRequestPayload instantiates a new SaasAsyncSubscriptionRequestPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSaasAsyncSubscriptionRequestPayload() *SaasAsyncSubscriptionRequestPayload {
	this := SaasAsyncSubscriptionRequestPayload{}
	return &this
}

// NewSaasAsyncSubscriptionRequestPayloadWithDefaults instantiates a new SaasAsyncSubscriptionRequestPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSaasAsyncSubscriptionRequestPayloadWithDefaults() *SaasAsyncSubscriptionRequestPayload {
	this := SaasAsyncSubscriptionRequestPayload{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *SaasAsyncSubscriptionRequestPayload) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaasAsyncSubscriptionRequestPayload) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *SaasAsyncSubscriptionRequestPayload) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *SaasAsyncSubscriptionRequestPayload) SetMessage(v string) {
	o.Message = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SaasAsyncSubscriptionRequestPayload) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaasAsyncSubscriptionRequestPayload) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SaasAsyncSubscriptionRequestPayload) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SaasAsyncSubscriptionRequestPayload) SetStatus(v string) {
	o.Status = &v
}

// GetSubscriptionUrl returns the SubscriptionUrl field value if set, zero value otherwise.
func (o *SaasAsyncSubscriptionRequestPayload) GetSubscriptionUrl() string {
	if o == nil || IsNil(o.SubscriptionUrl) {
		var ret string
		return ret
	}
	return *o.SubscriptionUrl
}

// GetSubscriptionUrlOk returns a tuple with the SubscriptionUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaasAsyncSubscriptionRequestPayload) GetSubscriptionUrlOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionUrl) {
		return nil, false
	}
	return o.SubscriptionUrl, true
}

// HasSubscriptionUrl returns a boolean if a field has been set.
func (o *SaasAsyncSubscriptionRequestPayload) HasSubscriptionUrl() bool {
	if o != nil && !IsNil(o.SubscriptionUrl) {
		return true
	}

	return false
}

// SetSubscriptionUrl gets a reference to the given string and assigns it to the SubscriptionUrl field.
func (o *SaasAsyncSubscriptionRequestPayload) SetSubscriptionUrl(v string) {
	o.SubscriptionUrl = &v
}

func (o SaasAsyncSubscriptionRequestPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SaasAsyncSubscriptionRequestPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.SubscriptionUrl) {
		toSerialize["subscriptionUrl"] = o.SubscriptionUrl
	}
	return toSerialize, nil
}

type NullableSaasAsyncSubscriptionRequestPayload struct {
	value *SaasAsyncSubscriptionRequestPayload
	isSet bool
}

func (v NullableSaasAsyncSubscriptionRequestPayload) Get() *SaasAsyncSubscriptionRequestPayload {
	return v.value
}

func (v *NullableSaasAsyncSubscriptionRequestPayload) Set(val *SaasAsyncSubscriptionRequestPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableSaasAsyncSubscriptionRequestPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableSaasAsyncSubscriptionRequestPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSaasAsyncSubscriptionRequestPayload(val *SaasAsyncSubscriptionRequestPayload) *NullableSaasAsyncSubscriptionRequestPayload {
	return &NullableSaasAsyncSubscriptionRequestPayload{value: val, isSet: true}
}

func (v NullableSaasAsyncSubscriptionRequestPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSaasAsyncSubscriptionRequestPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


