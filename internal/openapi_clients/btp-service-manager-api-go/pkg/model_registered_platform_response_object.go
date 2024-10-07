/*
Service Manager

Service Manager provides REST APIs that are responsible for the creation and consumption of service instances in any connected runtime environment.   Use the Service Manager APIs to perform various operations related to your platforms, service brokers, service instances, and service bindings.  Get service plans and service offerings associated with your environment.    #### Platforms   Platforms are OSBAPI-enabled software systems on which applications and services are hosted.   With the Service Manager, you can now register your platform and enable it to consume the SAP BTP services from your native environment.   This registration results in a returned set of credentials that are needed to deploy the Service Manager agent.     #### Service Brokers   Service brokers act as brokers between the Service Manager and a platform’s marketplace to advertise catalogues of service offerings and service plans.  They also receive and process the requests from the marketplace to provision, bind, unbind, and deprovision these offerings and plans.    #### Service Instances   Service instances are instantiations of service plans that make the functionality of those service plans available for consumption.    #### Service Bindings   Service bindings provide access details to existing service instances.  The access details are part of the service bindings' ‘credentials’ property, and typically include access URLs and credentials.    #### Service Plans   Service plans represent sets of capabilities provided by a service offering.  For example, database service offerings provide different plans for different database versions or sizes, while the Service Manager plans offer different data access levels.    #### Service Offerings   Service offerings are advertisements of the services that are supported by a service broker.  For example, software that you can consume in the subaccount.  Service offerings are related to one or more service plans.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the RegisteredPlatformResponseObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegisteredPlatformResponseObject{}

// RegisteredPlatformResponseObject struct for RegisteredPlatformResponseObject
type RegisteredPlatformResponseObject struct {
	// The time the platform was created. <br/>In ISO 8601 format:</br> YYYY-MM-DDThh:mm:ssTZD
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Credentials *Credentials `json:"credentials,omitempty"`
	// The description of the platform.
	Description *string `json:"description,omitempty"`
	// The ID of the platform.<br>You can use this ID to get details about the platform, to update, or to delete the platform.<br/> See the GET, PATCH, and DELETE APIs for the **Platforms** group.
	Id *string `json:"id,omitempty"`
	// Additional data associated with the resource entity. <br><br>Can be an empty object.
	Labels *map[string][]string `json:"labels,omitempty"`
	// The technical name of the platform.
	Name *string `json:"name,omitempty"`
	// Whether the platform is ready for consumption.
	Ready *bool `json:"ready,omitempty"`
	// The type of the platform.<br/><br/> Possible values: 
	Type *string `json:"type,omitempty"`
	// The last time the platform was updated. <br/>In ISO 8601 format. </br> 
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewRegisteredPlatformResponseObject instantiates a new RegisteredPlatformResponseObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegisteredPlatformResponseObject() *RegisteredPlatformResponseObject {
	this := RegisteredPlatformResponseObject{}
	return &this
}

// NewRegisteredPlatformResponseObjectWithDefaults instantiates a new RegisteredPlatformResponseObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegisteredPlatformResponseObjectWithDefaults() *RegisteredPlatformResponseObject {
	this := RegisteredPlatformResponseObject{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *RegisteredPlatformResponseObject) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisteredPlatformResponseObject) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *RegisteredPlatformResponseObject) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *RegisteredPlatformResponseObject) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *RegisteredPlatformResponseObject) GetCredentials() Credentials {
	if o == nil || IsNil(o.Credentials) {
		var ret Credentials
		return ret
	}
	return *o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisteredPlatformResponseObject) GetCredentialsOk() (*Credentials, bool) {
	if o == nil || IsNil(o.Credentials) {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *RegisteredPlatformResponseObject) HasCredentials() bool {
	if o != nil && !IsNil(o.Credentials) {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given Credentials and assigns it to the Credentials field.
func (o *RegisteredPlatformResponseObject) SetCredentials(v Credentials) {
	o.Credentials = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RegisteredPlatformResponseObject) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisteredPlatformResponseObject) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RegisteredPlatformResponseObject) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RegisteredPlatformResponseObject) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RegisteredPlatformResponseObject) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisteredPlatformResponseObject) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RegisteredPlatformResponseObject) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RegisteredPlatformResponseObject) SetId(v string) {
	o.Id = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *RegisteredPlatformResponseObject) GetLabels() map[string][]string {
	if o == nil || IsNil(o.Labels) {
		var ret map[string][]string
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisteredPlatformResponseObject) GetLabelsOk() (*map[string][]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *RegisteredPlatformResponseObject) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string][]string and assigns it to the Labels field.
func (o *RegisteredPlatformResponseObject) SetLabels(v map[string][]string) {
	o.Labels = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RegisteredPlatformResponseObject) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisteredPlatformResponseObject) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RegisteredPlatformResponseObject) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RegisteredPlatformResponseObject) SetName(v string) {
	o.Name = &v
}

// GetReady returns the Ready field value if set, zero value otherwise.
func (o *RegisteredPlatformResponseObject) GetReady() bool {
	if o == nil || IsNil(o.Ready) {
		var ret bool
		return ret
	}
	return *o.Ready
}

// GetReadyOk returns a tuple with the Ready field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisteredPlatformResponseObject) GetReadyOk() (*bool, bool) {
	if o == nil || IsNil(o.Ready) {
		return nil, false
	}
	return o.Ready, true
}

// HasReady returns a boolean if a field has been set.
func (o *RegisteredPlatformResponseObject) HasReady() bool {
	if o != nil && !IsNil(o.Ready) {
		return true
	}

	return false
}

// SetReady gets a reference to the given bool and assigns it to the Ready field.
func (o *RegisteredPlatformResponseObject) SetReady(v bool) {
	o.Ready = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RegisteredPlatformResponseObject) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisteredPlatformResponseObject) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RegisteredPlatformResponseObject) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RegisteredPlatformResponseObject) SetType(v string) {
	o.Type = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *RegisteredPlatformResponseObject) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisteredPlatformResponseObject) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *RegisteredPlatformResponseObject) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *RegisteredPlatformResponseObject) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o RegisteredPlatformResponseObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegisteredPlatformResponseObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Credentials) {
		toSerialize["credentials"] = o.Credentials
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Ready) {
		toSerialize["ready"] = o.Ready
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableRegisteredPlatformResponseObject struct {
	value *RegisteredPlatformResponseObject
	isSet bool
}

func (v NullableRegisteredPlatformResponseObject) Get() *RegisteredPlatformResponseObject {
	return v.value
}

func (v *NullableRegisteredPlatformResponseObject) Set(val *RegisteredPlatformResponseObject) {
	v.value = val
	v.isSet = true
}

func (v NullableRegisteredPlatformResponseObject) IsSet() bool {
	return v.isSet
}

func (v *NullableRegisteredPlatformResponseObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegisteredPlatformResponseObject(val *RegisteredPlatformResponseObject) *NullableRegisteredPlatformResponseObject {
	return &NullableRegisteredPlatformResponseObject{value: val, isSet: true}
}

func (v NullableRegisteredPlatformResponseObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegisteredPlatformResponseObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

