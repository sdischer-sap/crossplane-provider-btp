/*
SAP XSUAA REST API

Provides access to RoleTemplates, Roles, RoleCollection etc. using the XSUAA REST API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the ScimUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScimUser{}

// ScimUser struct for ScimUser
type ScimUser struct {
	Id *string `json:"id,omitempty"`
	ExternalId *string `json:"externalId,omitempty"`
	Meta *ScimMeta `json:"meta,omitempty"`
	UserName *string `json:"userName,omitempty"`
	Name *Name `json:"name,omitempty"`
	Emails []Email `json:"emails,omitempty"`
	Groups []Group `json:"groups,omitempty"`
	Approvals []Approval `json:"approvals,omitempty"`
	PhoneNumbers []PhoneNumber `json:"phoneNumbers,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	NickName *string `json:"nickName,omitempty"`
	ProfileUrl *string `json:"profileUrl,omitempty"`
	Title *string `json:"title,omitempty"`
	UserType *string `json:"userType,omitempty"`
	PreferredLanguage *string `json:"preferredLanguage,omitempty"`
	Locale *string `json:"locale,omitempty"`
	Timezone *string `json:"timezone,omitempty"`
	Active *bool `json:"active,omitempty"`
	Verified *bool `json:"verified,omitempty"`
	Origin *string `json:"origin,omitempty"`
	ZoneId *string `json:"zoneId,omitempty"`
	Salt *string `json:"salt,omitempty"`
	PasswordLastModified *time.Time `json:"passwordLastModified,omitempty"`
	PreviousLogonTime *int64 `json:"previousLogonTime,omitempty"`
	LastLogonTime *int64 `json:"lastLogonTime,omitempty"`
	Password *string `json:"password,omitempty"`
	Schemas []string `json:"schemas,omitempty"`
}

// NewScimUser instantiates a new ScimUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScimUser() *ScimUser {
	this := ScimUser{}
	return &this
}

// NewScimUserWithDefaults instantiates a new ScimUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScimUserWithDefaults() *ScimUser {
	this := ScimUser{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ScimUser) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimUser) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ScimUser) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ScimUser) SetId(v string) {
	o.Id = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *ScimUser) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimUser) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *ScimUser) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *ScimUser) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ScimUser) GetMeta() ScimMeta {
	if o == nil || IsNil(o.Meta) {
		var ret ScimMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimUser) GetMetaOk() (*ScimMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ScimUser) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ScimMeta and assigns it to the Meta field.
func (o *ScimUser) SetMeta(v ScimMeta) {
	o.Meta = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *ScimUser) GetUserName() string {
	if o == nil || IsNil(o.UserName) {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimUser) GetUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.UserName) {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *ScimUser) HasUserName() bool {
	if o != nil && !IsNil(o.UserName) {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *ScimUser) SetUserName(v string) {
	o.UserName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ScimUser) GetName() Name {
	if o == nil || IsNil(o.Name) {
		var ret Name
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimUser) GetNameOk() (*Name, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ScimUser) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given Name and assigns it to the Name field.
func (o *ScimUser) SetName(v Name) {
	o.Name = &v
}

// GetEmails returns the Emails field value if set, zero value otherwise.
func (o *ScimUser) GetEmails() []Email {
	if o == nil || IsNil(o.Emails) {
		var ret []Email
		return ret
	}
	return o.Emails
}

// GetEmailsOk returns a tuple with the Emails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimUser) GetEmailsOk() ([]Email, bool) {
	if o == nil || IsNil(o.Emails) {
		return nil, false
	}
	return o.Emails, true
}

// HasEmails returns a boolean if a field has been set.
func (o *ScimUser) HasEmails() bool {
	if o != nil && !IsNil(o.Emails) {
		return true
	}

	return false
}

// SetEmails gets a reference to the given []Email and assigns it to the Emails field.
func (o *ScimUser) SetEmails(v []Email) {
	o.Emails = v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *ScimUser) GetGroups() []Group {
	if o == nil || IsNil(o.Groups) {
		var ret []Group
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimUser) GetGroupsOk() ([]Group, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *ScimUser) HasGroups() bool {
	if o != nil && !IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []Group and assigns it to the Groups field.
func (o *ScimUser) SetGroups(v []Group) {
	o.Groups = v
}

// GetApprovals returns the Approvals field value if set, zero value otherwise.
func (o *ScimUser) GetApprovals() []Approval {
	if o == nil || IsNil(o.Approvals) {
		var ret []Approval
		return ret
	}
	return o.Approvals
}

// GetApprovalsOk returns a tuple with the Approvals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimUser) GetApprovalsOk() ([]Approval, bool) {
	if o == nil || IsNil(o.Approvals) {
		return nil, false
	}
	return o.Approvals, true
}

// HasApprovals returns a boolean if a field has been set.
func (o *ScimUser) HasApprovals() bool {
	if o != nil && !IsNil(o.Approvals) {
		return true
	}

	return false
}

// SetApprovals gets a reference to the given []Approval and assigns it to the Approvals field.
func (o *ScimUser) SetApprovals(v []Approval) {
	o.Approvals = v
}

// GetPhoneNumbers returns the PhoneNumbers field value if set, zero value otherwise.
func (o *ScimUser) GetPhoneNumbers() []PhoneNumber {
	if o == nil || IsNil(o.PhoneNumbers) {
		var ret []PhoneNumber
		return ret
	}
	return o.PhoneNumbers
}

// GetPhoneNumbersOk returns a tuple with the PhoneNumbers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimUser) GetPhoneNumbersOk() ([]PhoneNumber, bool) {
	if o == nil || IsNil(o.PhoneNumbers) {
		return nil, false
	}
	return o.PhoneNumbers, true
}

// HasPhoneNumbers returns a boolean if a field has been set.
func (o *ScimUser) HasPhoneNumbers() bool {
	if o != nil && !IsNil(o.PhoneNumbers) {
		return true
	}

	return false
}

// SetPhoneNumbers gets a reference to the given []PhoneNumber and assigns it to the PhoneNumbers field.
func (o *ScimUser) SetPhoneNumbers(v []PhoneNumber) {
	o.PhoneNumbers = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ScimUser) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimUser) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ScimUser) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ScimUser) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetNickName returns the NickName field value if set, zero value otherwise.
func (o *ScimUser) GetNickName() string {
	if o == nil || IsNil(o.NickName) {
		var ret string
		return ret
	}
	return *o.NickName
}

// GetNickNameOk returns a tuple with the NickName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimUser) GetNickNameOk() (*string, bool) {
	if o == nil || IsNil(o.NickName) {
		return nil, false
	}
	return o.NickName, true
}

// HasNickName returns a boolean if a field has been set.
func (o *ScimUser) HasNickName() bool {
	if o != nil && !IsNil(o.NickName) {
		return true
	}

	return false
}

// SetNickName gets a reference to the given string and assigns it to the NickName field.
func (o *ScimUser) SetNickName(v string) {
	o.NickName = &v
}

// GetProfileUrl returns the ProfileUrl field value if set, zero value otherwise.
func (o *ScimUser) GetProfileUrl() string {
	if o == nil || IsNil(o.ProfileUrl) {
		var ret string
		return ret
	}
	return *o.ProfileUrl
}

// GetProfileUrlOk returns a tuple with the ProfileUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimUser) GetProfileUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ProfileUrl) {
		return nil, false
	}
	return o.ProfileUrl, true
}

// HasProfileUrl returns a boolean if a field has been set.
func (o *ScimUser) HasProfileUrl() bool {
	if o != nil && !IsNil(o.ProfileUrl) {
		return true
	}

	return false
}

// SetProfileUrl gets a reference to the given string and assigns it to the ProfileUrl field.
func (o *ScimUser) SetProfileUrl(v string) {
	o.ProfileUrl = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ScimUser) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimUser) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ScimUser) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ScimUser) SetTitle(v string) {
	o.Title = &v
}

// GetUserType returns the UserType field value if set, zero value otherwise.
func (o *ScimUser) GetUserType() string {
	if o == nil || IsNil(o.UserType) {
		var ret string
		return ret
	}
	return *o.UserType
}

// GetUserTypeOk returns a tuple with the UserType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimUser) GetUserTypeOk() (*string, bool) {
	if o == nil || IsNil(o.UserType) {
		return nil, false
	}
	return o.UserType, true
}

// HasUserType returns a boolean if a field has been set.
func (o *ScimUser) HasUserType() bool {
	if o != nil && !IsNil(o.UserType) {
		return true
	}

	return false
}

// SetUserType gets a reference to the given string and assigns it to the UserType field.
func (o *ScimUser) SetUserType(v string) {
	o.UserType = &v
}

// GetPreferredLanguage returns the PreferredLanguage field value if set, zero value otherwise.
func (o *ScimUser) GetPreferredLanguage() string {
	if o == nil || IsNil(o.PreferredLanguage) {
		var ret string
		return ret
	}
	return *o.PreferredLanguage
}

// GetPreferredLanguageOk returns a tuple with the PreferredLanguage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimUser) GetPreferredLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.PreferredLanguage) {
		return nil, false
	}
	return o.PreferredLanguage, true
}

// HasPreferredLanguage returns a boolean if a field has been set.
func (o *ScimUser) HasPreferredLanguage() bool {
	if o != nil && !IsNil(o.PreferredLanguage) {
		return true
	}

	return false
}

// SetPreferredLanguage gets a reference to the given string and assigns it to the PreferredLanguage field.
func (o *ScimUser) SetPreferredLanguage(v string) {
	o.PreferredLanguage = &v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *ScimUser) GetLocale() string {
	if o == nil || IsNil(o.Locale) {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimUser) GetLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.Locale) {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *ScimUser) HasLocale() bool {
	if o != nil && !IsNil(o.Locale) {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *ScimUser) SetLocale(v string) {
	o.Locale = &v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *ScimUser) GetTimezone() string {
	if o == nil || IsNil(o.Timezone) {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimUser) GetTimezoneOk() (*string, bool) {
	if o == nil || IsNil(o.Timezone) {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *ScimUser) HasTimezone() bool {
	if o != nil && !IsNil(o.Timezone) {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *ScimUser) SetTimezone(v string) {
	o.Timezone = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *ScimUser) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimUser) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *ScimUser) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *ScimUser) SetActive(v bool) {
	o.Active = &v
}

// GetVerified returns the Verified field value if set, zero value otherwise.
func (o *ScimUser) GetVerified() bool {
	if o == nil || IsNil(o.Verified) {
		var ret bool
		return ret
	}
	return *o.Verified
}

// GetVerifiedOk returns a tuple with the Verified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimUser) GetVerifiedOk() (*bool, bool) {
	if o == nil || IsNil(o.Verified) {
		return nil, false
	}
	return o.Verified, true
}

// HasVerified returns a boolean if a field has been set.
func (o *ScimUser) HasVerified() bool {
	if o != nil && !IsNil(o.Verified) {
		return true
	}

	return false
}

// SetVerified gets a reference to the given bool and assigns it to the Verified field.
func (o *ScimUser) SetVerified(v bool) {
	o.Verified = &v
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *ScimUser) GetOrigin() string {
	if o == nil || IsNil(o.Origin) {
		var ret string
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimUser) GetOriginOk() (*string, bool) {
	if o == nil || IsNil(o.Origin) {
		return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *ScimUser) HasOrigin() bool {
	if o != nil && !IsNil(o.Origin) {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given string and assigns it to the Origin field.
func (o *ScimUser) SetOrigin(v string) {
	o.Origin = &v
}

// GetZoneId returns the ZoneId field value if set, zero value otherwise.
func (o *ScimUser) GetZoneId() string {
	if o == nil || IsNil(o.ZoneId) {
		var ret string
		return ret
	}
	return *o.ZoneId
}

// GetZoneIdOk returns a tuple with the ZoneId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimUser) GetZoneIdOk() (*string, bool) {
	if o == nil || IsNil(o.ZoneId) {
		return nil, false
	}
	return o.ZoneId, true
}

// HasZoneId returns a boolean if a field has been set.
func (o *ScimUser) HasZoneId() bool {
	if o != nil && !IsNil(o.ZoneId) {
		return true
	}

	return false
}

// SetZoneId gets a reference to the given string and assigns it to the ZoneId field.
func (o *ScimUser) SetZoneId(v string) {
	o.ZoneId = &v
}

// GetSalt returns the Salt field value if set, zero value otherwise.
func (o *ScimUser) GetSalt() string {
	if o == nil || IsNil(o.Salt) {
		var ret string
		return ret
	}
	return *o.Salt
}

// GetSaltOk returns a tuple with the Salt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimUser) GetSaltOk() (*string, bool) {
	if o == nil || IsNil(o.Salt) {
		return nil, false
	}
	return o.Salt, true
}

// HasSalt returns a boolean if a field has been set.
func (o *ScimUser) HasSalt() bool {
	if o != nil && !IsNil(o.Salt) {
		return true
	}

	return false
}

// SetSalt gets a reference to the given string and assigns it to the Salt field.
func (o *ScimUser) SetSalt(v string) {
	o.Salt = &v
}

// GetPasswordLastModified returns the PasswordLastModified field value if set, zero value otherwise.
func (o *ScimUser) GetPasswordLastModified() time.Time {
	if o == nil || IsNil(o.PasswordLastModified) {
		var ret time.Time
		return ret
	}
	return *o.PasswordLastModified
}

// GetPasswordLastModifiedOk returns a tuple with the PasswordLastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimUser) GetPasswordLastModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PasswordLastModified) {
		return nil, false
	}
	return o.PasswordLastModified, true
}

// HasPasswordLastModified returns a boolean if a field has been set.
func (o *ScimUser) HasPasswordLastModified() bool {
	if o != nil && !IsNil(o.PasswordLastModified) {
		return true
	}

	return false
}

// SetPasswordLastModified gets a reference to the given time.Time and assigns it to the PasswordLastModified field.
func (o *ScimUser) SetPasswordLastModified(v time.Time) {
	o.PasswordLastModified = &v
}

// GetPreviousLogonTime returns the PreviousLogonTime field value if set, zero value otherwise.
func (o *ScimUser) GetPreviousLogonTime() int64 {
	if o == nil || IsNil(o.PreviousLogonTime) {
		var ret int64
		return ret
	}
	return *o.PreviousLogonTime
}

// GetPreviousLogonTimeOk returns a tuple with the PreviousLogonTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimUser) GetPreviousLogonTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.PreviousLogonTime) {
		return nil, false
	}
	return o.PreviousLogonTime, true
}

// HasPreviousLogonTime returns a boolean if a field has been set.
func (o *ScimUser) HasPreviousLogonTime() bool {
	if o != nil && !IsNil(o.PreviousLogonTime) {
		return true
	}

	return false
}

// SetPreviousLogonTime gets a reference to the given int64 and assigns it to the PreviousLogonTime field.
func (o *ScimUser) SetPreviousLogonTime(v int64) {
	o.PreviousLogonTime = &v
}

// GetLastLogonTime returns the LastLogonTime field value if set, zero value otherwise.
func (o *ScimUser) GetLastLogonTime() int64 {
	if o == nil || IsNil(o.LastLogonTime) {
		var ret int64
		return ret
	}
	return *o.LastLogonTime
}

// GetLastLogonTimeOk returns a tuple with the LastLogonTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimUser) GetLastLogonTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.LastLogonTime) {
		return nil, false
	}
	return o.LastLogonTime, true
}

// HasLastLogonTime returns a boolean if a field has been set.
func (o *ScimUser) HasLastLogonTime() bool {
	if o != nil && !IsNil(o.LastLogonTime) {
		return true
	}

	return false
}

// SetLastLogonTime gets a reference to the given int64 and assigns it to the LastLogonTime field.
func (o *ScimUser) SetLastLogonTime(v int64) {
	o.LastLogonTime = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *ScimUser) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimUser) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *ScimUser) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *ScimUser) SetPassword(v string) {
	o.Password = &v
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *ScimUser) GetSchemas() []string {
	if o == nil || IsNil(o.Schemas) {
		var ret []string
		return ret
	}
	return o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScimUser) GetSchemasOk() ([]string, bool) {
	if o == nil || IsNil(o.Schemas) {
		return nil, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *ScimUser) HasSchemas() bool {
	if o != nil && !IsNil(o.Schemas) {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given []string and assigns it to the Schemas field.
func (o *ScimUser) SetSchemas(v []string) {
	o.Schemas = v
}

func (o ScimUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScimUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ExternalId) {
		toSerialize["externalId"] = o.ExternalId
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.UserName) {
		toSerialize["userName"] = o.UserName
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Emails) {
		toSerialize["emails"] = o.Emails
	}
	if !IsNil(o.Groups) {
		toSerialize["groups"] = o.Groups
	}
	if !IsNil(o.Approvals) {
		toSerialize["approvals"] = o.Approvals
	}
	if !IsNil(o.PhoneNumbers) {
		toSerialize["phoneNumbers"] = o.PhoneNumbers
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.NickName) {
		toSerialize["nickName"] = o.NickName
	}
	if !IsNil(o.ProfileUrl) {
		toSerialize["profileUrl"] = o.ProfileUrl
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.UserType) {
		toSerialize["userType"] = o.UserType
	}
	if !IsNil(o.PreferredLanguage) {
		toSerialize["preferredLanguage"] = o.PreferredLanguage
	}
	if !IsNil(o.Locale) {
		toSerialize["locale"] = o.Locale
	}
	if !IsNil(o.Timezone) {
		toSerialize["timezone"] = o.Timezone
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.Verified) {
		toSerialize["verified"] = o.Verified
	}
	if !IsNil(o.Origin) {
		toSerialize["origin"] = o.Origin
	}
	if !IsNil(o.ZoneId) {
		toSerialize["zoneId"] = o.ZoneId
	}
	if !IsNil(o.Salt) {
		toSerialize["salt"] = o.Salt
	}
	if !IsNil(o.PasswordLastModified) {
		toSerialize["passwordLastModified"] = o.PasswordLastModified
	}
	if !IsNil(o.PreviousLogonTime) {
		toSerialize["previousLogonTime"] = o.PreviousLogonTime
	}
	if !IsNil(o.LastLogonTime) {
		toSerialize["lastLogonTime"] = o.LastLogonTime
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.Schemas) {
		toSerialize["schemas"] = o.Schemas
	}
	return toSerialize, nil
}

type NullableScimUser struct {
	value *ScimUser
	isSet bool
}

func (v NullableScimUser) Get() *ScimUser {
	return v.value
}

func (v *NullableScimUser) Set(val *ScimUser) {
	v.value = val
	v.isSet = true
}

func (v NullableScimUser) IsSet() bool {
	return v.isSet
}

func (v *NullableScimUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScimUser(val *ScimUser) *NullableScimUser {
	return &NullableScimUser{value: val, isSet: true}
}

func (v NullableScimUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScimUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


