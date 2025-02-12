/*
Accounts Service

The Accounts service provides REST APIs that are responsible for the management of global accounts, and the creation and management of directories, subaccounts, and their custom properties/tags.  Global accounts represent a business entity and contain contract information, including customer details and purchased entitlements. The global account is the context for billing each customer.  Use the subaccount APIs to structure your global account according to your organization's and project's requirements regarding members, authorizations, and quotas. This service also provides you with APIs for creating and managing directories. While the use of directories is optional, they allow you to further organize and manage your subaccounts according to your specific technical and business needs. The service also lets you manage the custom properties/tags that you associate with your directories and subaccounts. NOTE: These APIs are relevant only for cloud management tools feature set B. For details and information about whether this applies to your global account, see [Cloud Management Tools - Feature Set Overview](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/Cloud/en-US/caf4e4e23aef4666ad8f125af393dfb2.html).  See also: * [Authorization](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/latest/en-US/3670474a58c24ac2b082e76cbbd9dc19.html) * [Rate Limiting](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/latest/en-US/77b217b3f57a45b987eb7fbc3305ce1e.html) * [Error Response Format](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/latest/en-US/77fef2fb104b4b1795e2e6cee790e8b8.html) * [Asynchronous Jobs](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/latest/en-US/0a0a6ab0ad114d72a6611c1c6b21683e.html)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the LegalLinksDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LegalLinksDTO{}

// LegalLinksDTO struct for LegalLinksDTO
type LegalLinksDTO struct {
	Privacy *string `json:"privacy,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LegalLinksDTO LegalLinksDTO

// NewLegalLinksDTO instantiates a new LegalLinksDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLegalLinksDTO() *LegalLinksDTO {
	this := LegalLinksDTO{}
	return &this
}

// NewLegalLinksDTOWithDefaults instantiates a new LegalLinksDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLegalLinksDTOWithDefaults() *LegalLinksDTO {
	this := LegalLinksDTO{}
	return &this
}

// GetPrivacy returns the Privacy field value if set, zero value otherwise.
func (o *LegalLinksDTO) GetPrivacy() string {
	if o == nil || IsNil(o.Privacy) {
		var ret string
		return ret
	}
	return *o.Privacy
}

// GetPrivacyOk returns a tuple with the Privacy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegalLinksDTO) GetPrivacyOk() (*string, bool) {
	if o == nil || IsNil(o.Privacy) {
		return nil, false
	}
	return o.Privacy, true
}

// HasPrivacy returns a boolean if a field has been set.
func (o *LegalLinksDTO) HasPrivacy() bool {
	if o != nil && !IsNil(o.Privacy) {
		return true
	}

	return false
}

// SetPrivacy gets a reference to the given string and assigns it to the Privacy field.
func (o *LegalLinksDTO) SetPrivacy(v string) {
	o.Privacy = &v
}

func (o LegalLinksDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LegalLinksDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Privacy) {
		toSerialize["privacy"] = o.Privacy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LegalLinksDTO) UnmarshalJSON(data []byte) (err error) {
	varLegalLinksDTO := _LegalLinksDTO{}

	err = json.Unmarshal(data, &varLegalLinksDTO)

	if err != nil {
		return err
	}

	*o = LegalLinksDTO(varLegalLinksDTO)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "privacy")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLegalLinksDTO struct {
	value *LegalLinksDTO
	isSet bool
}

func (v NullableLegalLinksDTO) Get() *LegalLinksDTO {
	return v.value
}

func (v *NullableLegalLinksDTO) Set(val *LegalLinksDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableLegalLinksDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableLegalLinksDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLegalLinksDTO(val *LegalLinksDTO) *NullableLegalLinksDTO {
	return &NullableLegalLinksDTO{value: val, isSet: true}
}

func (v NullableLegalLinksDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLegalLinksDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


