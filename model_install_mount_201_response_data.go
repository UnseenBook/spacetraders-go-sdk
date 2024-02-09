/*
SpaceTraders API

SpaceTraders is an open-universe game and learning platform that offers a set of HTTP endpoints to control a fleet of ships and explore a multiplayer universe.  The API is documented using [OpenAPI](https://github.com/SpaceTradersAPI/api-docs). You can send your first request right here in your browser to check the status of the game server.  ```json http {   \"method\": \"GET\",   \"url\": \"https://api.spacetraders.io/v2\", } ```  Unlike a traditional game, SpaceTraders does not have a first-party client or app to play the game. Instead, you can use the API to build your own client, write a script to automate your ships, or try an app built by the community.  We have a [Discord channel](https://discord.com/invite/jh6zurdWk5) where you can share your projects, ask questions, and get help from other players.   

API version: 2.0.0
Contact: joel@spacetraders.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package stc

import (
	"encoding/json"
	"fmt"
)

// checks if the InstallMount201ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstallMount201ResponseData{}

// InstallMount201ResponseData struct for InstallMount201ResponseData
type InstallMount201ResponseData struct {
	Agent Agent `json:"agent"`
	// List of installed mounts after the installation of the new mount.
	Mounts []ShipMount `json:"mounts"`
	Cargo ShipCargo `json:"cargo"`
	Transaction ShipModificationTransaction `json:"transaction"`
	AdditionalProperties map[string]interface{}
}

type _InstallMount201ResponseData InstallMount201ResponseData

// NewInstallMount201ResponseData instantiates a new InstallMount201ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstallMount201ResponseData(agent Agent, mounts []ShipMount, cargo ShipCargo, transaction ShipModificationTransaction) *InstallMount201ResponseData {
	this := InstallMount201ResponseData{}
	this.Agent = agent
	this.Mounts = mounts
	this.Cargo = cargo
	this.Transaction = transaction
	return &this
}

// NewInstallMount201ResponseDataWithDefaults instantiates a new InstallMount201ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstallMount201ResponseDataWithDefaults() *InstallMount201ResponseData {
	this := InstallMount201ResponseData{}
	return &this
}

// GetAgent returns the Agent field value
func (o *InstallMount201ResponseData) GetAgent() Agent {
	if o == nil {
		var ret Agent
		return ret
	}

	return o.Agent
}

// GetAgentOk returns a tuple with the Agent field value
// and a boolean to check if the value has been set.
func (o *InstallMount201ResponseData) GetAgentOk() (*Agent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Agent, true
}

// SetAgent sets field value
func (o *InstallMount201ResponseData) SetAgent(v Agent) {
	o.Agent = v
}

// GetMounts returns the Mounts field value
func (o *InstallMount201ResponseData) GetMounts() []ShipMount {
	if o == nil {
		var ret []ShipMount
		return ret
	}

	return o.Mounts
}

// GetMountsOk returns a tuple with the Mounts field value
// and a boolean to check if the value has been set.
func (o *InstallMount201ResponseData) GetMountsOk() ([]ShipMount, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mounts, true
}

// SetMounts sets field value
func (o *InstallMount201ResponseData) SetMounts(v []ShipMount) {
	o.Mounts = v
}

// GetCargo returns the Cargo field value
func (o *InstallMount201ResponseData) GetCargo() ShipCargo {
	if o == nil {
		var ret ShipCargo
		return ret
	}

	return o.Cargo
}

// GetCargoOk returns a tuple with the Cargo field value
// and a boolean to check if the value has been set.
func (o *InstallMount201ResponseData) GetCargoOk() (*ShipCargo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cargo, true
}

// SetCargo sets field value
func (o *InstallMount201ResponseData) SetCargo(v ShipCargo) {
	o.Cargo = v
}

// GetTransaction returns the Transaction field value
func (o *InstallMount201ResponseData) GetTransaction() ShipModificationTransaction {
	if o == nil {
		var ret ShipModificationTransaction
		return ret
	}

	return o.Transaction
}

// GetTransactionOk returns a tuple with the Transaction field value
// and a boolean to check if the value has been set.
func (o *InstallMount201ResponseData) GetTransactionOk() (*ShipModificationTransaction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Transaction, true
}

// SetTransaction sets field value
func (o *InstallMount201ResponseData) SetTransaction(v ShipModificationTransaction) {
	o.Transaction = v
}

func (o InstallMount201ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InstallMount201ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["agent"] = o.Agent
	toSerialize["mounts"] = o.Mounts
	toSerialize["cargo"] = o.Cargo
	toSerialize["transaction"] = o.Transaction

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InstallMount201ResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"agent",
		"mounts",
		"cargo",
		"transaction",
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

	varInstallMount201ResponseData := _InstallMount201ResponseData{}

	err = json.Unmarshal(data, &varInstallMount201ResponseData)

	if err != nil {
		return err
	}

	*o = InstallMount201ResponseData(varInstallMount201ResponseData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "agent")
		delete(additionalProperties, "mounts")
		delete(additionalProperties, "cargo")
		delete(additionalProperties, "transaction")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInstallMount201ResponseData struct {
	value *InstallMount201ResponseData
	isSet bool
}

func (v NullableInstallMount201ResponseData) Get() *InstallMount201ResponseData {
	return v.value
}

func (v *NullableInstallMount201ResponseData) Set(val *InstallMount201ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableInstallMount201ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableInstallMount201ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstallMount201ResponseData(val *InstallMount201ResponseData) *NullableInstallMount201ResponseData {
	return &NullableInstallMount201ResponseData{value: val, isSet: true}
}

func (v NullableInstallMount201ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstallMount201ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


