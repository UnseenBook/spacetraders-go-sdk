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

// checks if the ShipEngine type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipEngine{}

// ShipEngine The engine determines how quickly a ship travels between waypoints.
type ShipEngine struct {
	// The symbol of the engine.
	Symbol string `json:"symbol"`
	// The name of the engine.
	Name string `json:"name"`
	// The description of the engine.
	Description string `json:"description"`
	// Condition is a range of 0 to 100 where 0 is completely worn out and 100 is brand new.
	Condition *int32 `json:"condition,omitempty"`
	// The speed stat of this engine. The higher the speed, the faster a ship can travel from one point to another. Reduces the time of arrival when navigating the ship.
	Speed int32 `json:"speed"`
	Requirements ShipRequirements `json:"requirements"`
	AdditionalProperties map[string]interface{}
}

type _ShipEngine ShipEngine

// NewShipEngine instantiates a new ShipEngine object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipEngine(symbol string, name string, description string, speed int32, requirements ShipRequirements) *ShipEngine {
	this := ShipEngine{}
	this.Symbol = symbol
	this.Name = name
	this.Description = description
	this.Speed = speed
	this.Requirements = requirements
	return &this
}

// NewShipEngineWithDefaults instantiates a new ShipEngine object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipEngineWithDefaults() *ShipEngine {
	this := ShipEngine{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *ShipEngine) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *ShipEngine) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *ShipEngine) SetSymbol(v string) {
	o.Symbol = v
}

// GetName returns the Name field value
func (o *ShipEngine) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ShipEngine) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ShipEngine) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *ShipEngine) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ShipEngine) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ShipEngine) SetDescription(v string) {
	o.Description = v
}

// GetCondition returns the Condition field value if set, zero value otherwise.
func (o *ShipEngine) GetCondition() int32 {
	if o == nil || IsNil(o.Condition) {
		var ret int32
		return ret
	}
	return *o.Condition
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipEngine) GetConditionOk() (*int32, bool) {
	if o == nil || IsNil(o.Condition) {
		return nil, false
	}
	return o.Condition, true
}

// HasCondition returns a boolean if a field has been set.
func (o *ShipEngine) HasCondition() bool {
	if o != nil && !IsNil(o.Condition) {
		return true
	}

	return false
}

// SetCondition gets a reference to the given int32 and assigns it to the Condition field.
func (o *ShipEngine) SetCondition(v int32) {
	o.Condition = &v
}

// GetSpeed returns the Speed field value
func (o *ShipEngine) GetSpeed() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Speed
}

// GetSpeedOk returns a tuple with the Speed field value
// and a boolean to check if the value has been set.
func (o *ShipEngine) GetSpeedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Speed, true
}

// SetSpeed sets field value
func (o *ShipEngine) SetSpeed(v int32) {
	o.Speed = v
}

// GetRequirements returns the Requirements field value
func (o *ShipEngine) GetRequirements() ShipRequirements {
	if o == nil {
		var ret ShipRequirements
		return ret
	}

	return o.Requirements
}

// GetRequirementsOk returns a tuple with the Requirements field value
// and a boolean to check if the value has been set.
func (o *ShipEngine) GetRequirementsOk() (*ShipRequirements, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Requirements, true
}

// SetRequirements sets field value
func (o *ShipEngine) SetRequirements(v ShipRequirements) {
	o.Requirements = v
}

func (o ShipEngine) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShipEngine) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	if !IsNil(o.Condition) {
		toSerialize["condition"] = o.Condition
	}
	toSerialize["speed"] = o.Speed
	toSerialize["requirements"] = o.Requirements

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ShipEngine) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"symbol",
		"name",
		"description",
		"speed",
		"requirements",
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

	varShipEngine := _ShipEngine{}

	err = json.Unmarshal(data, &varShipEngine)

	if err != nil {
		return err
	}

	*o = ShipEngine(varShipEngine)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "condition")
		delete(additionalProperties, "speed")
		delete(additionalProperties, "requirements")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableShipEngine struct {
	value *ShipEngine
	isSet bool
}

func (v NullableShipEngine) Get() *ShipEngine {
	return v.value
}

func (v *NullableShipEngine) Set(val *ShipEngine) {
	v.value = val
	v.isSet = true
}

func (v NullableShipEngine) IsSet() bool {
	return v.isSet
}

func (v *NullableShipEngine) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipEngine(val *ShipEngine) *NullableShipEngine {
	return &NullableShipEngine{value: val, isSet: true}
}

func (v NullableShipEngine) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipEngine) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


