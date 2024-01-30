/*
SpaceTraders API

SpaceTraders is an open-universe game and learning platform that offers a set of HTTP endpoints to control a fleet of ships and explore a multiplayer universe.  The API is documented using [OpenAPI](https://github.com/SpaceTradersAPI/api-docs). You can send your first request right here in your browser to check the status of the game server.  ```json http {   \"method\": \"GET\",   \"url\": \"https://api.spacetraders.io/v2\", } ```  Unlike a traditional game, SpaceTraders does not have a first-party client or app to play the game. Instead, you can use the API to build your own client, write a script to automate your ships, or try an app built by the community.  We have a [Discord channel](https://discord.com/invite/jh6zurdWk5) where you can share your projects, ask questions, and get help from other players.   

API version: 2.0.0
Contact: joel@spacetraders.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package stsdk

import (
	"encoding/json"
	"fmt"
)

// checks if the ContractPayment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContractPayment{}

// ContractPayment Payments for the contract.
type ContractPayment struct {
	// The amount of credits received up front for accepting the contract.
	OnAccepted int32 `json:"onAccepted"`
	// The amount of credits received when the contract is fulfilled.
	OnFulfilled int32 `json:"onFulfilled"`
	AdditionalProperties map[string]interface{}
}

type _ContractPayment ContractPayment

// NewContractPayment instantiates a new ContractPayment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContractPayment(onAccepted int32, onFulfilled int32) *ContractPayment {
	this := ContractPayment{}
	this.OnAccepted = onAccepted
	this.OnFulfilled = onFulfilled
	return &this
}

// NewContractPaymentWithDefaults instantiates a new ContractPayment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractPaymentWithDefaults() *ContractPayment {
	this := ContractPayment{}
	return &this
}

// GetOnAccepted returns the OnAccepted field value
func (o *ContractPayment) GetOnAccepted() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OnAccepted
}

// GetOnAcceptedOk returns a tuple with the OnAccepted field value
// and a boolean to check if the value has been set.
func (o *ContractPayment) GetOnAcceptedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OnAccepted, true
}

// SetOnAccepted sets field value
func (o *ContractPayment) SetOnAccepted(v int32) {
	o.OnAccepted = v
}

// GetOnFulfilled returns the OnFulfilled field value
func (o *ContractPayment) GetOnFulfilled() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OnFulfilled
}

// GetOnFulfilledOk returns a tuple with the OnFulfilled field value
// and a boolean to check if the value has been set.
func (o *ContractPayment) GetOnFulfilledOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OnFulfilled, true
}

// SetOnFulfilled sets field value
func (o *ContractPayment) SetOnFulfilled(v int32) {
	o.OnFulfilled = v
}

func (o ContractPayment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContractPayment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["onAccepted"] = o.OnAccepted
	toSerialize["onFulfilled"] = o.OnFulfilled

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ContractPayment) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"onAccepted",
		"onFulfilled",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varContractPayment := _ContractPayment{}

	err = json.Unmarshal(data, &varContractPayment)

	if err != nil {
		return err
	}

	*o = ContractPayment(varContractPayment)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "onAccepted")
		delete(additionalProperties, "onFulfilled")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableContractPayment struct {
	value *ContractPayment
	isSet bool
}

func (v NullableContractPayment) Get() *ContractPayment {
	return v.value
}

func (v *NullableContractPayment) Set(val *ContractPayment) {
	v.value = val
	v.isSet = true
}

func (v NullableContractPayment) IsSet() bool {
	return v.isSet
}

func (v *NullableContractPayment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractPayment(val *ContractPayment) *NullableContractPayment {
	return &NullableContractPayment{value: val, isSet: true}
}

func (v NullableContractPayment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractPayment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


