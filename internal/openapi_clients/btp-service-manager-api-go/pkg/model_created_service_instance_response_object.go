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

// checks if the CreatedServiceInstanceResponseObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreatedServiceInstanceResponseObject{}

// CreatedServiceInstanceResponseObject struct for CreatedServiceInstanceResponseObject
type CreatedServiceInstanceResponseObject struct {
	// Contextual data for the resource.
	Context *map[string]string `json:"context,omitempty"`
	// The time the service instance was created. <br/>In ISO 8601 format:</br> YYYY-MM-DDThh:mm:ssTZD
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The URL of the web-based management UI for the service instance.
	DashboardUrl *string `json:"dashboard_url,omitempty"`
	// The ID of the service instance.
	Id *string `json:"id,omitempty"`
	// Additional data associated with the resource entity. <br><br>Can be an empty object.
	Labels *map[string][]string `json:"labels,omitempty"`
	// The maintenance info associated with the service instance.
	MaintenanceInfo *map[string]string `json:"maintenance_info,omitempty"`
	// The name of the service instance.
	Name *string `json:"name,omitempty"`
	// The ID of the platform associated with the service instance.
	PlatformId *string `json:"platform_id,omitempty"`
	// Whether the service instance is ready for consumption.
	Ready *bool `json:"ready,omitempty"`
	// The ID of the instance to which the service instance refers.
	ReferencedInstanceId *string `json:"referenced_instance_id,omitempty"`
	// The ID of the serivce plan associated with the service instance.
	ServicePlanId *string `json:"service_plan_id,omitempty"`
	// Whether the service instance is shared.
	Shared *bool `json:"shared,omitempty"`
	// The last time the service instance was updated.<br/>In ISO 8601 format.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// Whether the resource can be used.
	Usable *bool `json:"usable,omitempty"`
}

// NewCreatedServiceInstanceResponseObject instantiates a new CreatedServiceInstanceResponseObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatedServiceInstanceResponseObject() *CreatedServiceInstanceResponseObject {
	this := CreatedServiceInstanceResponseObject{}
	return &this
}

// NewCreatedServiceInstanceResponseObjectWithDefaults instantiates a new CreatedServiceInstanceResponseObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatedServiceInstanceResponseObjectWithDefaults() *CreatedServiceInstanceResponseObject {
	this := CreatedServiceInstanceResponseObject{}
	return &this
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *CreatedServiceInstanceResponseObject) GetContext() map[string]string {
	if o == nil || IsNil(o.Context) {
		var ret map[string]string
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatedServiceInstanceResponseObject) GetContextOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Context) {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *CreatedServiceInstanceResponseObject) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given map[string]string and assigns it to the Context field.
func (o *CreatedServiceInstanceResponseObject) SetContext(v map[string]string) {
	o.Context = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CreatedServiceInstanceResponseObject) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatedServiceInstanceResponseObject) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CreatedServiceInstanceResponseObject) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *CreatedServiceInstanceResponseObject) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDashboardUrl returns the DashboardUrl field value if set, zero value otherwise.
func (o *CreatedServiceInstanceResponseObject) GetDashboardUrl() string {
	if o == nil || IsNil(o.DashboardUrl) {
		var ret string
		return ret
	}
	return *o.DashboardUrl
}

// GetDashboardUrlOk returns a tuple with the DashboardUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatedServiceInstanceResponseObject) GetDashboardUrlOk() (*string, bool) {
	if o == nil || IsNil(o.DashboardUrl) {
		return nil, false
	}
	return o.DashboardUrl, true
}

// HasDashboardUrl returns a boolean if a field has been set.
func (o *CreatedServiceInstanceResponseObject) HasDashboardUrl() bool {
	if o != nil && !IsNil(o.DashboardUrl) {
		return true
	}

	return false
}

// SetDashboardUrl gets a reference to the given string and assigns it to the DashboardUrl field.
func (o *CreatedServiceInstanceResponseObject) SetDashboardUrl(v string) {
	o.DashboardUrl = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreatedServiceInstanceResponseObject) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatedServiceInstanceResponseObject) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreatedServiceInstanceResponseObject) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CreatedServiceInstanceResponseObject) SetId(v string) {
	o.Id = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *CreatedServiceInstanceResponseObject) GetLabels() map[string][]string {
	if o == nil || IsNil(o.Labels) {
		var ret map[string][]string
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatedServiceInstanceResponseObject) GetLabelsOk() (*map[string][]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *CreatedServiceInstanceResponseObject) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string][]string and assigns it to the Labels field.
func (o *CreatedServiceInstanceResponseObject) SetLabels(v map[string][]string) {
	o.Labels = &v
}

// GetMaintenanceInfo returns the MaintenanceInfo field value if set, zero value otherwise.
func (o *CreatedServiceInstanceResponseObject) GetMaintenanceInfo() map[string]string {
	if o == nil || IsNil(o.MaintenanceInfo) {
		var ret map[string]string
		return ret
	}
	return *o.MaintenanceInfo
}

// GetMaintenanceInfoOk returns a tuple with the MaintenanceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatedServiceInstanceResponseObject) GetMaintenanceInfoOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.MaintenanceInfo) {
		return nil, false
	}
	return o.MaintenanceInfo, true
}

// HasMaintenanceInfo returns a boolean if a field has been set.
func (o *CreatedServiceInstanceResponseObject) HasMaintenanceInfo() bool {
	if o != nil && !IsNil(o.MaintenanceInfo) {
		return true
	}

	return false
}

// SetMaintenanceInfo gets a reference to the given map[string]string and assigns it to the MaintenanceInfo field.
func (o *CreatedServiceInstanceResponseObject) SetMaintenanceInfo(v map[string]string) {
	o.MaintenanceInfo = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreatedServiceInstanceResponseObject) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatedServiceInstanceResponseObject) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreatedServiceInstanceResponseObject) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreatedServiceInstanceResponseObject) SetName(v string) {
	o.Name = &v
}

// GetPlatformId returns the PlatformId field value if set, zero value otherwise.
func (o *CreatedServiceInstanceResponseObject) GetPlatformId() string {
	if o == nil || IsNil(o.PlatformId) {
		var ret string
		return ret
	}
	return *o.PlatformId
}

// GetPlatformIdOk returns a tuple with the PlatformId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatedServiceInstanceResponseObject) GetPlatformIdOk() (*string, bool) {
	if o == nil || IsNil(o.PlatformId) {
		return nil, false
	}
	return o.PlatformId, true
}

// HasPlatformId returns a boolean if a field has been set.
func (o *CreatedServiceInstanceResponseObject) HasPlatformId() bool {
	if o != nil && !IsNil(o.PlatformId) {
		return true
	}

	return false
}

// SetPlatformId gets a reference to the given string and assigns it to the PlatformId field.
func (o *CreatedServiceInstanceResponseObject) SetPlatformId(v string) {
	o.PlatformId = &v
}

// GetReady returns the Ready field value if set, zero value otherwise.
func (o *CreatedServiceInstanceResponseObject) GetReady() bool {
	if o == nil || IsNil(o.Ready) {
		var ret bool
		return ret
	}
	return *o.Ready
}

// GetReadyOk returns a tuple with the Ready field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatedServiceInstanceResponseObject) GetReadyOk() (*bool, bool) {
	if o == nil || IsNil(o.Ready) {
		return nil, false
	}
	return o.Ready, true
}

// HasReady returns a boolean if a field has been set.
func (o *CreatedServiceInstanceResponseObject) HasReady() bool {
	if o != nil && !IsNil(o.Ready) {
		return true
	}

	return false
}

// SetReady gets a reference to the given bool and assigns it to the Ready field.
func (o *CreatedServiceInstanceResponseObject) SetReady(v bool) {
	o.Ready = &v
}

// GetReferencedInstanceId returns the ReferencedInstanceId field value if set, zero value otherwise.
func (o *CreatedServiceInstanceResponseObject) GetReferencedInstanceId() string {
	if o == nil || IsNil(o.ReferencedInstanceId) {
		var ret string
		return ret
	}
	return *o.ReferencedInstanceId
}

// GetReferencedInstanceIdOk returns a tuple with the ReferencedInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatedServiceInstanceResponseObject) GetReferencedInstanceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ReferencedInstanceId) {
		return nil, false
	}
	return o.ReferencedInstanceId, true
}

// HasReferencedInstanceId returns a boolean if a field has been set.
func (o *CreatedServiceInstanceResponseObject) HasReferencedInstanceId() bool {
	if o != nil && !IsNil(o.ReferencedInstanceId) {
		return true
	}

	return false
}

// SetReferencedInstanceId gets a reference to the given string and assigns it to the ReferencedInstanceId field.
func (o *CreatedServiceInstanceResponseObject) SetReferencedInstanceId(v string) {
	o.ReferencedInstanceId = &v
}

// GetServicePlanId returns the ServicePlanId field value if set, zero value otherwise.
func (o *CreatedServiceInstanceResponseObject) GetServicePlanId() string {
	if o == nil || IsNil(o.ServicePlanId) {
		var ret string
		return ret
	}
	return *o.ServicePlanId
}

// GetServicePlanIdOk returns a tuple with the ServicePlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatedServiceInstanceResponseObject) GetServicePlanIdOk() (*string, bool) {
	if o == nil || IsNil(o.ServicePlanId) {
		return nil, false
	}
	return o.ServicePlanId, true
}

// HasServicePlanId returns a boolean if a field has been set.
func (o *CreatedServiceInstanceResponseObject) HasServicePlanId() bool {
	if o != nil && !IsNil(o.ServicePlanId) {
		return true
	}

	return false
}

// SetServicePlanId gets a reference to the given string and assigns it to the ServicePlanId field.
func (o *CreatedServiceInstanceResponseObject) SetServicePlanId(v string) {
	o.ServicePlanId = &v
}

// GetShared returns the Shared field value if set, zero value otherwise.
func (o *CreatedServiceInstanceResponseObject) GetShared() bool {
	if o == nil || IsNil(o.Shared) {
		var ret bool
		return ret
	}
	return *o.Shared
}

// GetSharedOk returns a tuple with the Shared field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatedServiceInstanceResponseObject) GetSharedOk() (*bool, bool) {
	if o == nil || IsNil(o.Shared) {
		return nil, false
	}
	return o.Shared, true
}

// HasShared returns a boolean if a field has been set.
func (o *CreatedServiceInstanceResponseObject) HasShared() bool {
	if o != nil && !IsNil(o.Shared) {
		return true
	}

	return false
}

// SetShared gets a reference to the given bool and assigns it to the Shared field.
func (o *CreatedServiceInstanceResponseObject) SetShared(v bool) {
	o.Shared = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *CreatedServiceInstanceResponseObject) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatedServiceInstanceResponseObject) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *CreatedServiceInstanceResponseObject) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *CreatedServiceInstanceResponseObject) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetUsable returns the Usable field value if set, zero value otherwise.
func (o *CreatedServiceInstanceResponseObject) GetUsable() bool {
	if o == nil || IsNil(o.Usable) {
		var ret bool
		return ret
	}
	return *o.Usable
}

// GetUsableOk returns a tuple with the Usable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatedServiceInstanceResponseObject) GetUsableOk() (*bool, bool) {
	if o == nil || IsNil(o.Usable) {
		return nil, false
	}
	return o.Usable, true
}

// HasUsable returns a boolean if a field has been set.
func (o *CreatedServiceInstanceResponseObject) HasUsable() bool {
	if o != nil && !IsNil(o.Usable) {
		return true
	}

	return false
}

// SetUsable gets a reference to the given bool and assigns it to the Usable field.
func (o *CreatedServiceInstanceResponseObject) SetUsable(v bool) {
	o.Usable = &v
}

func (o CreatedServiceInstanceResponseObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreatedServiceInstanceResponseObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Context) {
		toSerialize["context"] = o.Context
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.DashboardUrl) {
		toSerialize["dashboard_url"] = o.DashboardUrl
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.MaintenanceInfo) {
		toSerialize["maintenance_info"] = o.MaintenanceInfo
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PlatformId) {
		toSerialize["platform_id"] = o.PlatformId
	}
	if !IsNil(o.Ready) {
		toSerialize["ready"] = o.Ready
	}
	if !IsNil(o.ReferencedInstanceId) {
		toSerialize["referenced_instance_id"] = o.ReferencedInstanceId
	}
	if !IsNil(o.ServicePlanId) {
		toSerialize["service_plan_id"] = o.ServicePlanId
	}
	if !IsNil(o.Shared) {
		toSerialize["shared"] = o.Shared
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.Usable) {
		toSerialize["usable"] = o.Usable
	}
	return toSerialize, nil
}

type NullableCreatedServiceInstanceResponseObject struct {
	value *CreatedServiceInstanceResponseObject
	isSet bool
}

func (v NullableCreatedServiceInstanceResponseObject) Get() *CreatedServiceInstanceResponseObject {
	return v.value
}

func (v *NullableCreatedServiceInstanceResponseObject) Set(val *CreatedServiceInstanceResponseObject) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatedServiceInstanceResponseObject) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatedServiceInstanceResponseObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatedServiceInstanceResponseObject(val *CreatedServiceInstanceResponseObject) *NullableCreatedServiceInstanceResponseObject {
	return &NullableCreatedServiceInstanceResponseObject{value: val, isSet: true}
}

func (v NullableCreatedServiceInstanceResponseObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatedServiceInstanceResponseObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


