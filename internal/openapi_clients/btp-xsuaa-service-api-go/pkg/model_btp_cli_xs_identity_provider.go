/*
SAP XSUAA REST API

Provides access to RoleTemplates, Roles, RoleCollection etc. using the XSUAA REST API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the BtpCliXsIdentityProvider type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BtpCliXsIdentityProvider{}

// BtpCliXsIdentityProvider struct for BtpCliXsIdentityProvider
type BtpCliXsIdentityProvider struct {
	Name *string `json:"name,omitempty"`
	OriginKey *string `json:"originKey,omitempty"`
	TypeOfTrust *string `json:"typeOfTrust,omitempty"`
	Status *string `json:"status,omitempty"`
	Description *string `json:"description,omitempty"`
	IdentityProvider *string `json:"identityProvider,omitempty"`
	Domain *string `json:"domain,omitempty"`
	LinkTextForUserLogon *string `json:"linkTextForUserLogon,omitempty"`
	AvailableForUserLogon *string `json:"availableForUserLogon,omitempty"`
	CreateShadowUsersDuringLogon *string `json:"createShadowUsersDuringLogon,omitempty"`
	SapBtpCli *string `json:"sapBtpCli,omitempty"`
	Protocol *string `json:"protocol,omitempty"`
	ReadOnly *bool `json:"readOnly,omitempty"`
}

// NewBtpCliXsIdentityProvider instantiates a new BtpCliXsIdentityProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBtpCliXsIdentityProvider() *BtpCliXsIdentityProvider {
	this := BtpCliXsIdentityProvider{}
	return &this
}

// NewBtpCliXsIdentityProviderWithDefaults instantiates a new BtpCliXsIdentityProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBtpCliXsIdentityProviderWithDefaults() *BtpCliXsIdentityProvider {
	this := BtpCliXsIdentityProvider{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BtpCliXsIdentityProvider) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BtpCliXsIdentityProvider) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BtpCliXsIdentityProvider) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BtpCliXsIdentityProvider) SetName(v string) {
	o.Name = &v
}

// GetOriginKey returns the OriginKey field value if set, zero value otherwise.
func (o *BtpCliXsIdentityProvider) GetOriginKey() string {
	if o == nil || IsNil(o.OriginKey) {
		var ret string
		return ret
	}
	return *o.OriginKey
}

// GetOriginKeyOk returns a tuple with the OriginKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BtpCliXsIdentityProvider) GetOriginKeyOk() (*string, bool) {
	if o == nil || IsNil(o.OriginKey) {
		return nil, false
	}
	return o.OriginKey, true
}

// HasOriginKey returns a boolean if a field has been set.
func (o *BtpCliXsIdentityProvider) HasOriginKey() bool {
	if o != nil && !IsNil(o.OriginKey) {
		return true
	}

	return false
}

// SetOriginKey gets a reference to the given string and assigns it to the OriginKey field.
func (o *BtpCliXsIdentityProvider) SetOriginKey(v string) {
	o.OriginKey = &v
}

// GetTypeOfTrust returns the TypeOfTrust field value if set, zero value otherwise.
func (o *BtpCliXsIdentityProvider) GetTypeOfTrust() string {
	if o == nil || IsNil(o.TypeOfTrust) {
		var ret string
		return ret
	}
	return *o.TypeOfTrust
}

// GetTypeOfTrustOk returns a tuple with the TypeOfTrust field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BtpCliXsIdentityProvider) GetTypeOfTrustOk() (*string, bool) {
	if o == nil || IsNil(o.TypeOfTrust) {
		return nil, false
	}
	return o.TypeOfTrust, true
}

// HasTypeOfTrust returns a boolean if a field has been set.
func (o *BtpCliXsIdentityProvider) HasTypeOfTrust() bool {
	if o != nil && !IsNil(o.TypeOfTrust) {
		return true
	}

	return false
}

// SetTypeOfTrust gets a reference to the given string and assigns it to the TypeOfTrust field.
func (o *BtpCliXsIdentityProvider) SetTypeOfTrust(v string) {
	o.TypeOfTrust = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BtpCliXsIdentityProvider) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BtpCliXsIdentityProvider) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BtpCliXsIdentityProvider) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *BtpCliXsIdentityProvider) SetStatus(v string) {
	o.Status = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BtpCliXsIdentityProvider) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BtpCliXsIdentityProvider) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BtpCliXsIdentityProvider) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BtpCliXsIdentityProvider) SetDescription(v string) {
	o.Description = &v
}

// GetIdentityProvider returns the IdentityProvider field value if set, zero value otherwise.
func (o *BtpCliXsIdentityProvider) GetIdentityProvider() string {
	if o == nil || IsNil(o.IdentityProvider) {
		var ret string
		return ret
	}
	return *o.IdentityProvider
}

// GetIdentityProviderOk returns a tuple with the IdentityProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BtpCliXsIdentityProvider) GetIdentityProviderOk() (*string, bool) {
	if o == nil || IsNil(o.IdentityProvider) {
		return nil, false
	}
	return o.IdentityProvider, true
}

// HasIdentityProvider returns a boolean if a field has been set.
func (o *BtpCliXsIdentityProvider) HasIdentityProvider() bool {
	if o != nil && !IsNil(o.IdentityProvider) {
		return true
	}

	return false
}

// SetIdentityProvider gets a reference to the given string and assigns it to the IdentityProvider field.
func (o *BtpCliXsIdentityProvider) SetIdentityProvider(v string) {
	o.IdentityProvider = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *BtpCliXsIdentityProvider) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BtpCliXsIdentityProvider) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *BtpCliXsIdentityProvider) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *BtpCliXsIdentityProvider) SetDomain(v string) {
	o.Domain = &v
}

// GetLinkTextForUserLogon returns the LinkTextForUserLogon field value if set, zero value otherwise.
func (o *BtpCliXsIdentityProvider) GetLinkTextForUserLogon() string {
	if o == nil || IsNil(o.LinkTextForUserLogon) {
		var ret string
		return ret
	}
	return *o.LinkTextForUserLogon
}

// GetLinkTextForUserLogonOk returns a tuple with the LinkTextForUserLogon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BtpCliXsIdentityProvider) GetLinkTextForUserLogonOk() (*string, bool) {
	if o == nil || IsNil(o.LinkTextForUserLogon) {
		return nil, false
	}
	return o.LinkTextForUserLogon, true
}

// HasLinkTextForUserLogon returns a boolean if a field has been set.
func (o *BtpCliXsIdentityProvider) HasLinkTextForUserLogon() bool {
	if o != nil && !IsNil(o.LinkTextForUserLogon) {
		return true
	}

	return false
}

// SetLinkTextForUserLogon gets a reference to the given string and assigns it to the LinkTextForUserLogon field.
func (o *BtpCliXsIdentityProvider) SetLinkTextForUserLogon(v string) {
	o.LinkTextForUserLogon = &v
}

// GetAvailableForUserLogon returns the AvailableForUserLogon field value if set, zero value otherwise.
func (o *BtpCliXsIdentityProvider) GetAvailableForUserLogon() string {
	if o == nil || IsNil(o.AvailableForUserLogon) {
		var ret string
		return ret
	}
	return *o.AvailableForUserLogon
}

// GetAvailableForUserLogonOk returns a tuple with the AvailableForUserLogon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BtpCliXsIdentityProvider) GetAvailableForUserLogonOk() (*string, bool) {
	if o == nil || IsNil(o.AvailableForUserLogon) {
		return nil, false
	}
	return o.AvailableForUserLogon, true
}

// HasAvailableForUserLogon returns a boolean if a field has been set.
func (o *BtpCliXsIdentityProvider) HasAvailableForUserLogon() bool {
	if o != nil && !IsNil(o.AvailableForUserLogon) {
		return true
	}

	return false
}

// SetAvailableForUserLogon gets a reference to the given string and assigns it to the AvailableForUserLogon field.
func (o *BtpCliXsIdentityProvider) SetAvailableForUserLogon(v string) {
	o.AvailableForUserLogon = &v
}

// GetCreateShadowUsersDuringLogon returns the CreateShadowUsersDuringLogon field value if set, zero value otherwise.
func (o *BtpCliXsIdentityProvider) GetCreateShadowUsersDuringLogon() string {
	if o == nil || IsNil(o.CreateShadowUsersDuringLogon) {
		var ret string
		return ret
	}
	return *o.CreateShadowUsersDuringLogon
}

// GetCreateShadowUsersDuringLogonOk returns a tuple with the CreateShadowUsersDuringLogon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BtpCliXsIdentityProvider) GetCreateShadowUsersDuringLogonOk() (*string, bool) {
	if o == nil || IsNil(o.CreateShadowUsersDuringLogon) {
		return nil, false
	}
	return o.CreateShadowUsersDuringLogon, true
}

// HasCreateShadowUsersDuringLogon returns a boolean if a field has been set.
func (o *BtpCliXsIdentityProvider) HasCreateShadowUsersDuringLogon() bool {
	if o != nil && !IsNil(o.CreateShadowUsersDuringLogon) {
		return true
	}

	return false
}

// SetCreateShadowUsersDuringLogon gets a reference to the given string and assigns it to the CreateShadowUsersDuringLogon field.
func (o *BtpCliXsIdentityProvider) SetCreateShadowUsersDuringLogon(v string) {
	o.CreateShadowUsersDuringLogon = &v
}

// GetSapBtpCli returns the SapBtpCli field value if set, zero value otherwise.
func (o *BtpCliXsIdentityProvider) GetSapBtpCli() string {
	if o == nil || IsNil(o.SapBtpCli) {
		var ret string
		return ret
	}
	return *o.SapBtpCli
}

// GetSapBtpCliOk returns a tuple with the SapBtpCli field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BtpCliXsIdentityProvider) GetSapBtpCliOk() (*string, bool) {
	if o == nil || IsNil(o.SapBtpCli) {
		return nil, false
	}
	return o.SapBtpCli, true
}

// HasSapBtpCli returns a boolean if a field has been set.
func (o *BtpCliXsIdentityProvider) HasSapBtpCli() bool {
	if o != nil && !IsNil(o.SapBtpCli) {
		return true
	}

	return false
}

// SetSapBtpCli gets a reference to the given string and assigns it to the SapBtpCli field.
func (o *BtpCliXsIdentityProvider) SetSapBtpCli(v string) {
	o.SapBtpCli = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *BtpCliXsIdentityProvider) GetProtocol() string {
	if o == nil || IsNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BtpCliXsIdentityProvider) GetProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *BtpCliXsIdentityProvider) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *BtpCliXsIdentityProvider) SetProtocol(v string) {
	o.Protocol = &v
}

// GetReadOnly returns the ReadOnly field value if set, zero value otherwise.
func (o *BtpCliXsIdentityProvider) GetReadOnly() bool {
	if o == nil || IsNil(o.ReadOnly) {
		var ret bool
		return ret
	}
	return *o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BtpCliXsIdentityProvider) GetReadOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.ReadOnly) {
		return nil, false
	}
	return o.ReadOnly, true
}

// HasReadOnly returns a boolean if a field has been set.
func (o *BtpCliXsIdentityProvider) HasReadOnly() bool {
	if o != nil && !IsNil(o.ReadOnly) {
		return true
	}

	return false
}

// SetReadOnly gets a reference to the given bool and assigns it to the ReadOnly field.
func (o *BtpCliXsIdentityProvider) SetReadOnly(v bool) {
	o.ReadOnly = &v
}

func (o BtpCliXsIdentityProvider) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BtpCliXsIdentityProvider) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.OriginKey) {
		toSerialize["originKey"] = o.OriginKey
	}
	if !IsNil(o.TypeOfTrust) {
		toSerialize["typeOfTrust"] = o.TypeOfTrust
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.IdentityProvider) {
		toSerialize["identityProvider"] = o.IdentityProvider
	}
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	if !IsNil(o.LinkTextForUserLogon) {
		toSerialize["linkTextForUserLogon"] = o.LinkTextForUserLogon
	}
	if !IsNil(o.AvailableForUserLogon) {
		toSerialize["availableForUserLogon"] = o.AvailableForUserLogon
	}
	if !IsNil(o.CreateShadowUsersDuringLogon) {
		toSerialize["createShadowUsersDuringLogon"] = o.CreateShadowUsersDuringLogon
	}
	if !IsNil(o.SapBtpCli) {
		toSerialize["sapBtpCli"] = o.SapBtpCli
	}
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !IsNil(o.ReadOnly) {
		toSerialize["readOnly"] = o.ReadOnly
	}
	return toSerialize, nil
}

type NullableBtpCliXsIdentityProvider struct {
	value *BtpCliXsIdentityProvider
	isSet bool
}

func (v NullableBtpCliXsIdentityProvider) Get() *BtpCliXsIdentityProvider {
	return v.value
}

func (v *NullableBtpCliXsIdentityProvider) Set(val *BtpCliXsIdentityProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableBtpCliXsIdentityProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableBtpCliXsIdentityProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBtpCliXsIdentityProvider(val *BtpCliXsIdentityProvider) *NullableBtpCliXsIdentityProvider {
	return &NullableBtpCliXsIdentityProvider{value: val, isSet: true}
}

func (v NullableBtpCliXsIdentityProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBtpCliXsIdentityProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

