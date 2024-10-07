/*
Service Manager

Service Manager provides REST APIs that are responsible for the creation and consumption of service instances in any connected runtime environment.   Use the Service Manager APIs to perform various operations related to your platforms, service brokers, service instances, and service bindings.  Get service plans and service offerings associated with your environment.    #### Platforms   Platforms are OSBAPI-enabled software systems on which applications and services are hosted.   With the Service Manager, you can now register your platform and enable it to consume the SAP BTP services from your native environment.   This registration results in a returned set of credentials that are needed to deploy the Service Manager agent.     #### Service Brokers   Service brokers act as brokers between the Service Manager and a platform’s marketplace to advertise catalogues of service offerings and service plans.  They also receive and process the requests from the marketplace to provision, bind, unbind, and deprovision these offerings and plans.    #### Service Instances   Service instances are instantiations of service plans that make the functionality of those service plans available for consumption.    #### Service Bindings   Service bindings provide access details to existing service instances.  The access details are part of the service bindings' ‘credentials’ property, and typically include access URLs and credentials.    #### Service Plans   Service plans represent sets of capabilities provided by a service offering.  For example, database service offerings provide different plans for different database versions or sizes, while the Service Manager plans offer different data access levels.    #### Service Offerings   Service offerings are advertisements of the services that are supported by a service broker.  For example, software that you can consume in the subaccount.  Service offerings are related to one or more service plans.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ServicePlanMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServicePlanMetadata{}

// ServicePlanMetadata struct for ServicePlanMetadata
type ServicePlanMetadata struct {
	// The latest supported OSB version.
	SupportedMaxOSBVersion *string `json:"supportedMaxOSBVersion,omitempty"`
	// The earliest supported OSB version.
	SupportedMinOSBVersion *string `json:"supportedMinOSBVersion,omitempty"`
	// Platforms supported by the service plan.
	SupportedPlatforms []string `json:"supportedPlatforms,omitempty"`
}

// NewServicePlanMetadata instantiates a new ServicePlanMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServicePlanMetadata() *ServicePlanMetadata {
	this := ServicePlanMetadata{}
	return &this
}

// NewServicePlanMetadataWithDefaults instantiates a new ServicePlanMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServicePlanMetadataWithDefaults() *ServicePlanMetadata {
	this := ServicePlanMetadata{}
	return &this
}

// GetSupportedMaxOSBVersion returns the SupportedMaxOSBVersion field value if set, zero value otherwise.
func (o *ServicePlanMetadata) GetSupportedMaxOSBVersion() string {
	if o == nil || IsNil(o.SupportedMaxOSBVersion) {
		var ret string
		return ret
	}
	return *o.SupportedMaxOSBVersion
}

// GetSupportedMaxOSBVersionOk returns a tuple with the SupportedMaxOSBVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePlanMetadata) GetSupportedMaxOSBVersionOk() (*string, bool) {
	if o == nil || IsNil(o.SupportedMaxOSBVersion) {
		return nil, false
	}
	return o.SupportedMaxOSBVersion, true
}

// HasSupportedMaxOSBVersion returns a boolean if a field has been set.
func (o *ServicePlanMetadata) HasSupportedMaxOSBVersion() bool {
	if o != nil && !IsNil(o.SupportedMaxOSBVersion) {
		return true
	}

	return false
}

// SetSupportedMaxOSBVersion gets a reference to the given string and assigns it to the SupportedMaxOSBVersion field.
func (o *ServicePlanMetadata) SetSupportedMaxOSBVersion(v string) {
	o.SupportedMaxOSBVersion = &v
}

// GetSupportedMinOSBVersion returns the SupportedMinOSBVersion field value if set, zero value otherwise.
func (o *ServicePlanMetadata) GetSupportedMinOSBVersion() string {
	if o == nil || IsNil(o.SupportedMinOSBVersion) {
		var ret string
		return ret
	}
	return *o.SupportedMinOSBVersion
}

// GetSupportedMinOSBVersionOk returns a tuple with the SupportedMinOSBVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePlanMetadata) GetSupportedMinOSBVersionOk() (*string, bool) {
	if o == nil || IsNil(o.SupportedMinOSBVersion) {
		return nil, false
	}
	return o.SupportedMinOSBVersion, true
}

// HasSupportedMinOSBVersion returns a boolean if a field has been set.
func (o *ServicePlanMetadata) HasSupportedMinOSBVersion() bool {
	if o != nil && !IsNil(o.SupportedMinOSBVersion) {
		return true
	}

	return false
}

// SetSupportedMinOSBVersion gets a reference to the given string and assigns it to the SupportedMinOSBVersion field.
func (o *ServicePlanMetadata) SetSupportedMinOSBVersion(v string) {
	o.SupportedMinOSBVersion = &v
}

// GetSupportedPlatforms returns the SupportedPlatforms field value if set, zero value otherwise.
func (o *ServicePlanMetadata) GetSupportedPlatforms() []string {
	if o == nil || IsNil(o.SupportedPlatforms) {
		var ret []string
		return ret
	}
	return o.SupportedPlatforms
}

// GetSupportedPlatformsOk returns a tuple with the SupportedPlatforms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePlanMetadata) GetSupportedPlatformsOk() ([]string, bool) {
	if o == nil || IsNil(o.SupportedPlatforms) {
		return nil, false
	}
	return o.SupportedPlatforms, true
}

// HasSupportedPlatforms returns a boolean if a field has been set.
func (o *ServicePlanMetadata) HasSupportedPlatforms() bool {
	if o != nil && !IsNil(o.SupportedPlatforms) {
		return true
	}

	return false
}

// SetSupportedPlatforms gets a reference to the given []string and assigns it to the SupportedPlatforms field.
func (o *ServicePlanMetadata) SetSupportedPlatforms(v []string) {
	o.SupportedPlatforms = v
}

func (o ServicePlanMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServicePlanMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SupportedMaxOSBVersion) {
		toSerialize["supportedMaxOSBVersion"] = o.SupportedMaxOSBVersion
	}
	if !IsNil(o.SupportedMinOSBVersion) {
		toSerialize["supportedMinOSBVersion"] = o.SupportedMinOSBVersion
	}
	if !IsNil(o.SupportedPlatforms) {
		toSerialize["supportedPlatforms"] = o.SupportedPlatforms
	}
	return toSerialize, nil
}

type NullableServicePlanMetadata struct {
	value *ServicePlanMetadata
	isSet bool
}

func (v NullableServicePlanMetadata) Get() *ServicePlanMetadata {
	return v.value
}

func (v *NullableServicePlanMetadata) Set(val *ServicePlanMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableServicePlanMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableServicePlanMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServicePlanMetadata(val *ServicePlanMetadata) *NullableServicePlanMetadata {
	return &NullableServicePlanMetadata{value: val, isSet: true}
}

func (v NullableServicePlanMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServicePlanMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


