/*
Entitlements Service

The Entitlements service provides REST APIs that manage the assignments of entitlements and quotas to subaccounts and directories.   Entitlements and their quota are automatically assigned to the global account when a customer order is fulfilled. Use the APIs in this service to manage the distribution of this global quota to your directories and subaccounts.   NOTE: These APIs are relevant only for cloud management tools feature set B. For details and information about whether this applies to your global account, see [Cloud Management Tools - Feature Set Overview](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/Cloud/en-US/caf4e4e23aef4666ad8f125af393dfb2.html).  See also: * [Authorization](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/latest/en-US/3670474a58c24ac2b082e76cbbd9dc19.html) * [Rate Limiting](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/latest/en-US/77b217b3f57a45b987eb7fbc3305ce1e.html) * [Error Response Format](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/latest/en-US/77fef2fb104b4b1795e2e6cee790e8b8.html) * [Asynchronous Jobs](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/latest/en-US/0a0a6ab0ad114d72a6611c1c6b21683e.html)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the AssignedServicePlanSubaccountDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssignedServicePlanSubaccountDTO{}

// AssignedServicePlanSubaccountDTO Assignment detailed information
type AssignedServicePlanSubaccountDTO struct {
	// The quantity of the entitlement that is assigned to the root global account or directory.
	Amount *float32 `json:"amount,omitempty"`
	// Whether the plan is automatically distributed to the subaccounts that are located in the directory.
	AutoAssign *bool `json:"autoAssign,omitempty"`
	// Specifies if the plan was automatically assigned regardless of any action by an admin. This applies to entitlements that are always available to subaccounts and cannot be removed.
	AutoAssigned *bool `json:"autoAssigned,omitempty"`
	// The amount of the entitlement to automatically assign to subaccounts that are added in the future to the entitlement's assigned directory.   Requires that autoAssign is set to TRUE, and there is remaining quota for the entitlement. To automatically distribute to subaccounts that are added in the future to the directory, distribute must be set to TRUE.
	AutoDistributeAmount *int32 `json:"autoDistributeAmount,omitempty"`
	// Date the subaccount has been created. Dates and times are in UTC format.
	CreatedDate *float64 `json:"createdDate,omitempty"`
	// The unique ID of the global account or directory to which the entitlement is assigned.
	EntityId *string `json:"entityId,omitempty"`
	// The current state of the service plan assignment. * <b>STARTED:</b> CRUD operation on an entity has started. * <b>PROCESSING:</b> A series of operations related to the entity is in progress. * <b>PROCESSING_FAILED:</b> The processing operations failed. * <b>OK:</b> The CRUD operation or series of operations completed successfully.
	EntityState *string `json:"entityState,omitempty"`
	// The type of entity to which the entitlement is assigned. * <b>SUBACCOUNT:</b> The entitlement is assigned to a subaccount. * <b>GLOBAL_ACCOUNT:</b> The entitlement is assigned to a root global account. * <b>DIRECTORY:</b> The entitlement is assigned to a directory.
	EntityType *string `json:"entityType,omitempty"`
	// Date the subaccount has been modified. Dates and times are in UTC format.
	ModifiedDate *float64 `json:"modifiedDate,omitempty"`
	ParentAmount *float32 `json:"parentAmount,omitempty"`
	ParentId *string `json:"parentId,omitempty"`
	ParentRemainingAmount *float32 `json:"parentRemainingAmount,omitempty"`
	ParentType *string `json:"parentType,omitempty"`
	// The requested amount when it is different from the actual amount because the request state is still in process or failed.
	RequestedAmount *float32 `json:"requestedAmount,omitempty"`
	// Global account resource details
	Resources []ExternalResourceRequestPayload `json:"resources,omitempty"`
	// Information about the current state.
	StateMessage *string `json:"stateMessage,omitempty"`
	// True, if an unlimited quota of this service plan assigned to the directory or subaccount in the global account. False, if the service plan is assigned to the directory or subaccount with a limited numeric quota, even if the service plan has an unlimited usage entitled on the level of the global account.
	UnlimitedAmountAssigned *bool `json:"unlimitedAmountAssigned,omitempty"`
}

// NewAssignedServicePlanSubaccountDTO instantiates a new AssignedServicePlanSubaccountDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssignedServicePlanSubaccountDTO() *AssignedServicePlanSubaccountDTO {
	this := AssignedServicePlanSubaccountDTO{}
	return &this
}

// NewAssignedServicePlanSubaccountDTOWithDefaults instantiates a new AssignedServicePlanSubaccountDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssignedServicePlanSubaccountDTOWithDefaults() *AssignedServicePlanSubaccountDTO {
	this := AssignedServicePlanSubaccountDTO{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *AssignedServicePlanSubaccountDTO) GetAmount() float32 {
	if o == nil || IsNil(o.Amount) {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignedServicePlanSubaccountDTO) GetAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *AssignedServicePlanSubaccountDTO) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *AssignedServicePlanSubaccountDTO) SetAmount(v float32) {
	o.Amount = &v
}

// GetAutoAssign returns the AutoAssign field value if set, zero value otherwise.
func (o *AssignedServicePlanSubaccountDTO) GetAutoAssign() bool {
	if o == nil || IsNil(o.AutoAssign) {
		var ret bool
		return ret
	}
	return *o.AutoAssign
}

// GetAutoAssignOk returns a tuple with the AutoAssign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignedServicePlanSubaccountDTO) GetAutoAssignOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoAssign) {
		return nil, false
	}
	return o.AutoAssign, true
}

// HasAutoAssign returns a boolean if a field has been set.
func (o *AssignedServicePlanSubaccountDTO) HasAutoAssign() bool {
	if o != nil && !IsNil(o.AutoAssign) {
		return true
	}

	return false
}

// SetAutoAssign gets a reference to the given bool and assigns it to the AutoAssign field.
func (o *AssignedServicePlanSubaccountDTO) SetAutoAssign(v bool) {
	o.AutoAssign = &v
}

// GetAutoAssigned returns the AutoAssigned field value if set, zero value otherwise.
func (o *AssignedServicePlanSubaccountDTO) GetAutoAssigned() bool {
	if o == nil || IsNil(o.AutoAssigned) {
		var ret bool
		return ret
	}
	return *o.AutoAssigned
}

// GetAutoAssignedOk returns a tuple with the AutoAssigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignedServicePlanSubaccountDTO) GetAutoAssignedOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoAssigned) {
		return nil, false
	}
	return o.AutoAssigned, true
}

// HasAutoAssigned returns a boolean if a field has been set.
func (o *AssignedServicePlanSubaccountDTO) HasAutoAssigned() bool {
	if o != nil && !IsNil(o.AutoAssigned) {
		return true
	}

	return false
}

// SetAutoAssigned gets a reference to the given bool and assigns it to the AutoAssigned field.
func (o *AssignedServicePlanSubaccountDTO) SetAutoAssigned(v bool) {
	o.AutoAssigned = &v
}

// GetAutoDistributeAmount returns the AutoDistributeAmount field value if set, zero value otherwise.
func (o *AssignedServicePlanSubaccountDTO) GetAutoDistributeAmount() int32 {
	if o == nil || IsNil(o.AutoDistributeAmount) {
		var ret int32
		return ret
	}
	return *o.AutoDistributeAmount
}

// GetAutoDistributeAmountOk returns a tuple with the AutoDistributeAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignedServicePlanSubaccountDTO) GetAutoDistributeAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.AutoDistributeAmount) {
		return nil, false
	}
	return o.AutoDistributeAmount, true
}

// HasAutoDistributeAmount returns a boolean if a field has been set.
func (o *AssignedServicePlanSubaccountDTO) HasAutoDistributeAmount() bool {
	if o != nil && !IsNil(o.AutoDistributeAmount) {
		return true
	}

	return false
}

// SetAutoDistributeAmount gets a reference to the given int32 and assigns it to the AutoDistributeAmount field.
func (o *AssignedServicePlanSubaccountDTO) SetAutoDistributeAmount(v int32) {
	o.AutoDistributeAmount = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *AssignedServicePlanSubaccountDTO) GetCreatedDate() float64 {
	if o == nil || IsNil(o.CreatedDate) {
		var ret float64
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignedServicePlanSubaccountDTO) GetCreatedDateOk() (*float64, bool) {
	if o == nil || IsNil(o.CreatedDate) {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *AssignedServicePlanSubaccountDTO) HasCreatedDate() bool {
	if o != nil && !IsNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given float64 and assigns it to the CreatedDate field.
func (o *AssignedServicePlanSubaccountDTO) SetCreatedDate(v float64) {
	o.CreatedDate = &v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *AssignedServicePlanSubaccountDTO) GetEntityId() string {
	if o == nil || IsNil(o.EntityId) {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignedServicePlanSubaccountDTO) GetEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.EntityId) {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *AssignedServicePlanSubaccountDTO) HasEntityId() bool {
	if o != nil && !IsNil(o.EntityId) {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *AssignedServicePlanSubaccountDTO) SetEntityId(v string) {
	o.EntityId = &v
}

// GetEntityState returns the EntityState field value if set, zero value otherwise.
func (o *AssignedServicePlanSubaccountDTO) GetEntityState() string {
	if o == nil || IsNil(o.EntityState) {
		var ret string
		return ret
	}
	return *o.EntityState
}

// GetEntityStateOk returns a tuple with the EntityState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignedServicePlanSubaccountDTO) GetEntityStateOk() (*string, bool) {
	if o == nil || IsNil(o.EntityState) {
		return nil, false
	}
	return o.EntityState, true
}

// HasEntityState returns a boolean if a field has been set.
func (o *AssignedServicePlanSubaccountDTO) HasEntityState() bool {
	if o != nil && !IsNil(o.EntityState) {
		return true
	}

	return false
}

// SetEntityState gets a reference to the given string and assigns it to the EntityState field.
func (o *AssignedServicePlanSubaccountDTO) SetEntityState(v string) {
	o.EntityState = &v
}

// GetEntityType returns the EntityType field value if set, zero value otherwise.
func (o *AssignedServicePlanSubaccountDTO) GetEntityType() string {
	if o == nil || IsNil(o.EntityType) {
		var ret string
		return ret
	}
	return *o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignedServicePlanSubaccountDTO) GetEntityTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EntityType) {
		return nil, false
	}
	return o.EntityType, true
}

// HasEntityType returns a boolean if a field has been set.
func (o *AssignedServicePlanSubaccountDTO) HasEntityType() bool {
	if o != nil && !IsNil(o.EntityType) {
		return true
	}

	return false
}

// SetEntityType gets a reference to the given string and assigns it to the EntityType field.
func (o *AssignedServicePlanSubaccountDTO) SetEntityType(v string) {
	o.EntityType = &v
}

// GetModifiedDate returns the ModifiedDate field value if set, zero value otherwise.
func (o *AssignedServicePlanSubaccountDTO) GetModifiedDate() float64 {
	if o == nil || IsNil(o.ModifiedDate) {
		var ret float64
		return ret
	}
	return *o.ModifiedDate
}

// GetModifiedDateOk returns a tuple with the ModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignedServicePlanSubaccountDTO) GetModifiedDateOk() (*float64, bool) {
	if o == nil || IsNil(o.ModifiedDate) {
		return nil, false
	}
	return o.ModifiedDate, true
}

// HasModifiedDate returns a boolean if a field has been set.
func (o *AssignedServicePlanSubaccountDTO) HasModifiedDate() bool {
	if o != nil && !IsNil(o.ModifiedDate) {
		return true
	}

	return false
}

// SetModifiedDate gets a reference to the given float64 and assigns it to the ModifiedDate field.
func (o *AssignedServicePlanSubaccountDTO) SetModifiedDate(v float64) {
	o.ModifiedDate = &v
}

// GetParentAmount returns the ParentAmount field value if set, zero value otherwise.
func (o *AssignedServicePlanSubaccountDTO) GetParentAmount() float32 {
	if o == nil || IsNil(o.ParentAmount) {
		var ret float32
		return ret
	}
	return *o.ParentAmount
}

// GetParentAmountOk returns a tuple with the ParentAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignedServicePlanSubaccountDTO) GetParentAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.ParentAmount) {
		return nil, false
	}
	return o.ParentAmount, true
}

// HasParentAmount returns a boolean if a field has been set.
func (o *AssignedServicePlanSubaccountDTO) HasParentAmount() bool {
	if o != nil && !IsNil(o.ParentAmount) {
		return true
	}

	return false
}

// SetParentAmount gets a reference to the given float32 and assigns it to the ParentAmount field.
func (o *AssignedServicePlanSubaccountDTO) SetParentAmount(v float32) {
	o.ParentAmount = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *AssignedServicePlanSubaccountDTO) GetParentId() string {
	if o == nil || IsNil(o.ParentId) {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignedServicePlanSubaccountDTO) GetParentIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentId) {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *AssignedServicePlanSubaccountDTO) HasParentId() bool {
	if o != nil && !IsNil(o.ParentId) {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *AssignedServicePlanSubaccountDTO) SetParentId(v string) {
	o.ParentId = &v
}

// GetParentRemainingAmount returns the ParentRemainingAmount field value if set, zero value otherwise.
func (o *AssignedServicePlanSubaccountDTO) GetParentRemainingAmount() float32 {
	if o == nil || IsNil(o.ParentRemainingAmount) {
		var ret float32
		return ret
	}
	return *o.ParentRemainingAmount
}

// GetParentRemainingAmountOk returns a tuple with the ParentRemainingAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignedServicePlanSubaccountDTO) GetParentRemainingAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.ParentRemainingAmount) {
		return nil, false
	}
	return o.ParentRemainingAmount, true
}

// HasParentRemainingAmount returns a boolean if a field has been set.
func (o *AssignedServicePlanSubaccountDTO) HasParentRemainingAmount() bool {
	if o != nil && !IsNil(o.ParentRemainingAmount) {
		return true
	}

	return false
}

// SetParentRemainingAmount gets a reference to the given float32 and assigns it to the ParentRemainingAmount field.
func (o *AssignedServicePlanSubaccountDTO) SetParentRemainingAmount(v float32) {
	o.ParentRemainingAmount = &v
}

// GetParentType returns the ParentType field value if set, zero value otherwise.
func (o *AssignedServicePlanSubaccountDTO) GetParentType() string {
	if o == nil || IsNil(o.ParentType) {
		var ret string
		return ret
	}
	return *o.ParentType
}

// GetParentTypeOk returns a tuple with the ParentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignedServicePlanSubaccountDTO) GetParentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ParentType) {
		return nil, false
	}
	return o.ParentType, true
}

// HasParentType returns a boolean if a field has been set.
func (o *AssignedServicePlanSubaccountDTO) HasParentType() bool {
	if o != nil && !IsNil(o.ParentType) {
		return true
	}

	return false
}

// SetParentType gets a reference to the given string and assigns it to the ParentType field.
func (o *AssignedServicePlanSubaccountDTO) SetParentType(v string) {
	o.ParentType = &v
}

// GetRequestedAmount returns the RequestedAmount field value if set, zero value otherwise.
func (o *AssignedServicePlanSubaccountDTO) GetRequestedAmount() float32 {
	if o == nil || IsNil(o.RequestedAmount) {
		var ret float32
		return ret
	}
	return *o.RequestedAmount
}

// GetRequestedAmountOk returns a tuple with the RequestedAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignedServicePlanSubaccountDTO) GetRequestedAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.RequestedAmount) {
		return nil, false
	}
	return o.RequestedAmount, true
}

// HasRequestedAmount returns a boolean if a field has been set.
func (o *AssignedServicePlanSubaccountDTO) HasRequestedAmount() bool {
	if o != nil && !IsNil(o.RequestedAmount) {
		return true
	}

	return false
}

// SetRequestedAmount gets a reference to the given float32 and assigns it to the RequestedAmount field.
func (o *AssignedServicePlanSubaccountDTO) SetRequestedAmount(v float32) {
	o.RequestedAmount = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *AssignedServicePlanSubaccountDTO) GetResources() []ExternalResourceRequestPayload {
	if o == nil || IsNil(o.Resources) {
		var ret []ExternalResourceRequestPayload
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignedServicePlanSubaccountDTO) GetResourcesOk() ([]ExternalResourceRequestPayload, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *AssignedServicePlanSubaccountDTO) HasResources() bool {
	if o != nil && !IsNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given []ExternalResourceRequestPayload and assigns it to the Resources field.
func (o *AssignedServicePlanSubaccountDTO) SetResources(v []ExternalResourceRequestPayload) {
	o.Resources = v
}

// GetStateMessage returns the StateMessage field value if set, zero value otherwise.
func (o *AssignedServicePlanSubaccountDTO) GetStateMessage() string {
	if o == nil || IsNil(o.StateMessage) {
		var ret string
		return ret
	}
	return *o.StateMessage
}

// GetStateMessageOk returns a tuple with the StateMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignedServicePlanSubaccountDTO) GetStateMessageOk() (*string, bool) {
	if o == nil || IsNil(o.StateMessage) {
		return nil, false
	}
	return o.StateMessage, true
}

// HasStateMessage returns a boolean if a field has been set.
func (o *AssignedServicePlanSubaccountDTO) HasStateMessage() bool {
	if o != nil && !IsNil(o.StateMessage) {
		return true
	}

	return false
}

// SetStateMessage gets a reference to the given string and assigns it to the StateMessage field.
func (o *AssignedServicePlanSubaccountDTO) SetStateMessage(v string) {
	o.StateMessage = &v
}

// GetUnlimitedAmountAssigned returns the UnlimitedAmountAssigned field value if set, zero value otherwise.
func (o *AssignedServicePlanSubaccountDTO) GetUnlimitedAmountAssigned() bool {
	if o == nil || IsNil(o.UnlimitedAmountAssigned) {
		var ret bool
		return ret
	}
	return *o.UnlimitedAmountAssigned
}

// GetUnlimitedAmountAssignedOk returns a tuple with the UnlimitedAmountAssigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignedServicePlanSubaccountDTO) GetUnlimitedAmountAssignedOk() (*bool, bool) {
	if o == nil || IsNil(o.UnlimitedAmountAssigned) {
		return nil, false
	}
	return o.UnlimitedAmountAssigned, true
}

// HasUnlimitedAmountAssigned returns a boolean if a field has been set.
func (o *AssignedServicePlanSubaccountDTO) HasUnlimitedAmountAssigned() bool {
	if o != nil && !IsNil(o.UnlimitedAmountAssigned) {
		return true
	}

	return false
}

// SetUnlimitedAmountAssigned gets a reference to the given bool and assigns it to the UnlimitedAmountAssigned field.
func (o *AssignedServicePlanSubaccountDTO) SetUnlimitedAmountAssigned(v bool) {
	o.UnlimitedAmountAssigned = &v
}

func (o AssignedServicePlanSubaccountDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssignedServicePlanSubaccountDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.AutoAssign) {
		toSerialize["autoAssign"] = o.AutoAssign
	}
	if !IsNil(o.AutoAssigned) {
		toSerialize["autoAssigned"] = o.AutoAssigned
	}
	if !IsNil(o.AutoDistributeAmount) {
		toSerialize["autoDistributeAmount"] = o.AutoDistributeAmount
	}
	if !IsNil(o.CreatedDate) {
		toSerialize["createdDate"] = o.CreatedDate
	}
	if !IsNil(o.EntityId) {
		toSerialize["entityId"] = o.EntityId
	}
	if !IsNil(o.EntityState) {
		toSerialize["entityState"] = o.EntityState
	}
	if !IsNil(o.EntityType) {
		toSerialize["entityType"] = o.EntityType
	}
	if !IsNil(o.ModifiedDate) {
		toSerialize["modifiedDate"] = o.ModifiedDate
	}
	if !IsNil(o.ParentAmount) {
		toSerialize["parentAmount"] = o.ParentAmount
	}
	if !IsNil(o.ParentId) {
		toSerialize["parentId"] = o.ParentId
	}
	if !IsNil(o.ParentRemainingAmount) {
		toSerialize["parentRemainingAmount"] = o.ParentRemainingAmount
	}
	if !IsNil(o.ParentType) {
		toSerialize["parentType"] = o.ParentType
	}
	if !IsNil(o.RequestedAmount) {
		toSerialize["requestedAmount"] = o.RequestedAmount
	}
	if !IsNil(o.Resources) {
		toSerialize["resources"] = o.Resources
	}
	if !IsNil(o.StateMessage) {
		toSerialize["stateMessage"] = o.StateMessage
	}
	if !IsNil(o.UnlimitedAmountAssigned) {
		toSerialize["unlimitedAmountAssigned"] = o.UnlimitedAmountAssigned
	}
	return toSerialize, nil
}

type NullableAssignedServicePlanSubaccountDTO struct {
	value *AssignedServicePlanSubaccountDTO
	isSet bool
}

func (v NullableAssignedServicePlanSubaccountDTO) Get() *AssignedServicePlanSubaccountDTO {
	return v.value
}

func (v *NullableAssignedServicePlanSubaccountDTO) Set(val *AssignedServicePlanSubaccountDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableAssignedServicePlanSubaccountDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableAssignedServicePlanSubaccountDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssignedServicePlanSubaccountDTO(val *AssignedServicePlanSubaccountDTO) *NullableAssignedServicePlanSubaccountDTO {
	return &NullableAssignedServicePlanSubaccountDTO{value: val, isSet: true}
}

func (v NullableAssignedServicePlanSubaccountDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssignedServicePlanSubaccountDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

