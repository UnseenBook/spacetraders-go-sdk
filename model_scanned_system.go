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

// checks if the ScannedSystem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScannedSystem{}

// ScannedSystem Details of a system was that scanned.
type ScannedSystem struct {
	// Symbol of the system.
	Symbol string `json:"symbol"`
	// Symbol of the system's sector.
	SectorSymbol string `json:"sectorSymbol"`
	Type SystemType `json:"type"`
	// Position in the universe in the x axis.
	X int32 `json:"x"`
	// Position in the universe in the y axis.
	Y int32 `json:"y"`
	// The system's distance from the scanning ship.
	Distance int32 `json:"distance"`
	AdditionalProperties map[string]interface{}
}

type _ScannedSystem ScannedSystem

// NewScannedSystem instantiates a new ScannedSystem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScannedSystem(symbol string, sectorSymbol string, type_ SystemType, x int32, y int32, distance int32) *ScannedSystem {
	this := ScannedSystem{}
	this.Symbol = symbol
	this.SectorSymbol = sectorSymbol
	this.Type = type_
	this.X = x
	this.Y = y
	this.Distance = distance
	return &this
}

// NewScannedSystemWithDefaults instantiates a new ScannedSystem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScannedSystemWithDefaults() *ScannedSystem {
	this := ScannedSystem{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *ScannedSystem) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *ScannedSystem) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *ScannedSystem) SetSymbol(v string) {
	o.Symbol = v
}

// GetSectorSymbol returns the SectorSymbol field value
func (o *ScannedSystem) GetSectorSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SectorSymbol
}

// GetSectorSymbolOk returns a tuple with the SectorSymbol field value
// and a boolean to check if the value has been set.
func (o *ScannedSystem) GetSectorSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SectorSymbol, true
}

// SetSectorSymbol sets field value
func (o *ScannedSystem) SetSectorSymbol(v string) {
	o.SectorSymbol = v
}

// GetType returns the Type field value
func (o *ScannedSystem) GetType() SystemType {
	if o == nil {
		var ret SystemType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ScannedSystem) GetTypeOk() (*SystemType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ScannedSystem) SetType(v SystemType) {
	o.Type = v
}

// GetX returns the X field value
func (o *ScannedSystem) GetX() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.X
}

// GetXOk returns a tuple with the X field value
// and a boolean to check if the value has been set.
func (o *ScannedSystem) GetXOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.X, true
}

// SetX sets field value
func (o *ScannedSystem) SetX(v int32) {
	o.X = v
}

// GetY returns the Y field value
func (o *ScannedSystem) GetY() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Y
}

// GetYOk returns a tuple with the Y field value
// and a boolean to check if the value has been set.
func (o *ScannedSystem) GetYOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Y, true
}

// SetY sets field value
func (o *ScannedSystem) SetY(v int32) {
	o.Y = v
}

// GetDistance returns the Distance field value
func (o *ScannedSystem) GetDistance() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Distance
}

// GetDistanceOk returns a tuple with the Distance field value
// and a boolean to check if the value has been set.
func (o *ScannedSystem) GetDistanceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Distance, true
}

// SetDistance sets field value
func (o *ScannedSystem) SetDistance(v int32) {
	o.Distance = v
}

func (o ScannedSystem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScannedSystem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol
	toSerialize["sectorSymbol"] = o.SectorSymbol
	toSerialize["type"] = o.Type
	toSerialize["x"] = o.X
	toSerialize["y"] = o.Y
	toSerialize["distance"] = o.Distance

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ScannedSystem) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"symbol",
		"sectorSymbol",
		"type",
		"x",
		"y",
		"distance",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varScannedSystem := _ScannedSystem{}

	err = json.Unmarshal(bytes, &varScannedSystem)

	if err != nil {
		return err
	}

	*o = ScannedSystem(varScannedSystem)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "sectorSymbol")
		delete(additionalProperties, "type")
		delete(additionalProperties, "x")
		delete(additionalProperties, "y")
		delete(additionalProperties, "distance")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableScannedSystem struct {
	value *ScannedSystem
	isSet bool
}

func (v NullableScannedSystem) Get() *ScannedSystem {
	return v.value
}

func (v *NullableScannedSystem) Set(val *ScannedSystem) {
	v.value = val
	v.isSet = true
}

func (v NullableScannedSystem) IsSet() bool {
	return v.isSet
}

func (v *NullableScannedSystem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScannedSystem(val *ScannedSystem) *NullableScannedSystem {
	return &NullableScannedSystem{value: val, isSet: true}
}

func (v NullableScannedSystem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScannedSystem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


