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

// checks if the UpdateDirectoryRequestPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateDirectoryRequestPayload{}

// UpdateDirectoryRequestPayload Details of the directory properties to update.
type UpdateDirectoryRequestPayload struct {
	// (Deprecated) User-defined labels to assign, update, and remove as key-value pairs from the directory.  NOTE: This field is deprecated. \"customProperties\" supports only single values per key and is now replaced by the string array \"labels\", which supports multiple values per key. The \"customProperties\" field overwrites any labels with the same key. We recommend transitioning to the \"labels\" field. Do not include \"customProperties\" and \"labels\" together in the same request (customProperties will be ignored). 
	// Deprecated
	CustomProperties []UpdatePropertyRequestPayload `json:"customProperties,omitempty"`
	// The description of the directory for the customer-facing UIs.
	Description *string `json:"description,omitempty"`
	// The new descriptive name of the directory.
	DisplayName *string `json:"displayName,omitempty" validate:"regexp=^((?![\\/]).)*$"`
	// JSON array of up to 10 user-defined labels to assign as key-value pairs to the directory. Each label has a name (key) that you specify, and to which you can assign up to 10 corresponding values or leave empty. Keys and values are each limited to 63 characters.  Label keys and values are case-sensitive. Try to avoid creating duplicate variants of the same keys or values with a different casing (example: \"myValue\" and \"MyValue\").  Example:  {   \"Cost Center\": [\"19700626\"],   \"Department\": [\"Sales\"],   \"Contacts\": [\"name1@example.com\",\"name2@example.com\"],   \"EMEA\":[] }  IMPORTANT: The JSON array overwrites any labels that are currently assigned to the subaccount. In the request, you must include not only new and updated key-value pairs, but also existing ones that you do not want changed. Omit keys/values that you want removed as labels from the subaccount.   Any labels previously set using the deprecated \"customProperties\" field are also overwritten. 
	Labels *map[string][]string `json:"labels,omitempty"`
}

// NewUpdateDirectoryRequestPayload instantiates a new UpdateDirectoryRequestPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDirectoryRequestPayload() *UpdateDirectoryRequestPayload {
	this := UpdateDirectoryRequestPayload{}
	return &this
}

// NewUpdateDirectoryRequestPayloadWithDefaults instantiates a new UpdateDirectoryRequestPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDirectoryRequestPayloadWithDefaults() *UpdateDirectoryRequestPayload {
	this := UpdateDirectoryRequestPayload{}
	return &this
}

// GetCustomProperties returns the CustomProperties field value if set, zero value otherwise.
// Deprecated
func (o *UpdateDirectoryRequestPayload) GetCustomProperties() []UpdatePropertyRequestPayload {
	if o == nil || IsNil(o.CustomProperties) {
		var ret []UpdatePropertyRequestPayload
		return ret
	}
	return o.CustomProperties
}

// GetCustomPropertiesOk returns a tuple with the CustomProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *UpdateDirectoryRequestPayload) GetCustomPropertiesOk() ([]UpdatePropertyRequestPayload, bool) {
	if o == nil || IsNil(o.CustomProperties) {
		return nil, false
	}
	return o.CustomProperties, true
}

// HasCustomProperties returns a boolean if a field has been set.
func (o *UpdateDirectoryRequestPayload) HasCustomProperties() bool {
	if o != nil && !IsNil(o.CustomProperties) {
		return true
	}

	return false
}

// SetCustomProperties gets a reference to the given []UpdatePropertyRequestPayload and assigns it to the CustomProperties field.
// Deprecated
func (o *UpdateDirectoryRequestPayload) SetCustomProperties(v []UpdatePropertyRequestPayload) {
	o.CustomProperties = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateDirectoryRequestPayload) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDirectoryRequestPayload) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateDirectoryRequestPayload) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateDirectoryRequestPayload) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *UpdateDirectoryRequestPayload) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDirectoryRequestPayload) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *UpdateDirectoryRequestPayload) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *UpdateDirectoryRequestPayload) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *UpdateDirectoryRequestPayload) GetLabels() map[string][]string {
	if o == nil || IsNil(o.Labels) {
		var ret map[string][]string
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDirectoryRequestPayload) GetLabelsOk() (*map[string][]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *UpdateDirectoryRequestPayload) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string][]string and assigns it to the Labels field.
func (o *UpdateDirectoryRequestPayload) SetLabels(v map[string][]string) {
	o.Labels = &v
}

func (o UpdateDirectoryRequestPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateDirectoryRequestPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CustomProperties) {
		toSerialize["customProperties"] = o.CustomProperties
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	return toSerialize, nil
}

type NullableUpdateDirectoryRequestPayload struct {
	value *UpdateDirectoryRequestPayload
	isSet bool
}

func (v NullableUpdateDirectoryRequestPayload) Get() *UpdateDirectoryRequestPayload {
	return v.value
}

func (v *NullableUpdateDirectoryRequestPayload) Set(val *UpdateDirectoryRequestPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDirectoryRequestPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDirectoryRequestPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDirectoryRequestPayload(val *UpdateDirectoryRequestPayload) *NullableUpdateDirectoryRequestPayload {
	return &NullableUpdateDirectoryRequestPayload{value: val, isSet: true}
}

func (v NullableUpdateDirectoryRequestPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDirectoryRequestPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

