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

// checks if the WaypointModifier type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WaypointModifier{}

// WaypointModifier struct for WaypointModifier
type WaypointModifier struct {
	Symbol WaypointModifierSymbol `json:"symbol"`
	// The name of the trait.
	Name string `json:"name"`
	// A description of the trait.
	Description string `json:"description"`
	AdditionalProperties map[string]interface{}
}

type _WaypointModifier WaypointModifier

// NewWaypointModifier instantiates a new WaypointModifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWaypointModifier(symbol WaypointModifierSymbol, name string, description string) *WaypointModifier {
	this := WaypointModifier{}
	this.Symbol = symbol
	this.Name = name
	this.Description = description
	return &this
}

// NewWaypointModifierWithDefaults instantiates a new WaypointModifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWaypointModifierWithDefaults() *WaypointModifier {
	this := WaypointModifier{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *WaypointModifier) GetSymbol() WaypointModifierSymbol {
	if o == nil {
		var ret WaypointModifierSymbol
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *WaypointModifier) GetSymbolOk() (*WaypointModifierSymbol, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *WaypointModifier) SetSymbol(v WaypointModifierSymbol) {
	o.Symbol = v
}

// GetName returns the Name field value
func (o *WaypointModifier) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WaypointModifier) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WaypointModifier) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *WaypointModifier) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *WaypointModifier) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *WaypointModifier) SetDescription(v string) {
	o.Description = v
}

func (o WaypointModifier) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WaypointModifier) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WaypointModifier) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"symbol",
		"name",
		"description",
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

	varWaypointModifier := _WaypointModifier{}

	err = json.Unmarshal(data, &varWaypointModifier)

	if err != nil {
		return err
	}

	*o = WaypointModifier(varWaypointModifier)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWaypointModifier struct {
	value *WaypointModifier
	isSet bool
}

func (v NullableWaypointModifier) Get() *WaypointModifier {
	return v.value
}

func (v *NullableWaypointModifier) Set(val *WaypointModifier) {
	v.value = val
	v.isSet = true
}

func (v NullableWaypointModifier) IsSet() bool {
	return v.isSet
}

func (v *NullableWaypointModifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWaypointModifier(val *WaypointModifier) *NullableWaypointModifier {
	return &NullableWaypointModifier{value: val, isSet: true}
}

func (v NullableWaypointModifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWaypointModifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


