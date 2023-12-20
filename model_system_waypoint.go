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

// checks if the SystemWaypoint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SystemWaypoint{}

// SystemWaypoint struct for SystemWaypoint
type SystemWaypoint struct {
	// The symbol of the waypoint.
	Symbol string `json:"symbol"`
	Type WaypointType `json:"type"`
	// Relative position of the waypoint on the system's x axis. This is not an absolute position in the universe.
	X int32 `json:"x"`
	// Relative position of the waypoint on the system's y axis. This is not an absolute position in the universe.
	Y int32 `json:"y"`
	// Waypoints that orbit this waypoint.
	Orbitals []WaypointOrbital `json:"orbitals"`
	// The symbol of the parent waypoint, if this waypoint is in orbit around another waypoint. Otherwise this value is undefined.
	Orbits *string `json:"orbits,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SystemWaypoint SystemWaypoint

// NewSystemWaypoint instantiates a new SystemWaypoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemWaypoint(symbol string, type_ WaypointType, x int32, y int32, orbitals []WaypointOrbital) *SystemWaypoint {
	this := SystemWaypoint{}
	this.Symbol = symbol
	this.Type = type_
	this.X = x
	this.Y = y
	this.Orbitals = orbitals
	return &this
}

// NewSystemWaypointWithDefaults instantiates a new SystemWaypoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemWaypointWithDefaults() *SystemWaypoint {
	this := SystemWaypoint{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *SystemWaypoint) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *SystemWaypoint) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *SystemWaypoint) SetSymbol(v string) {
	o.Symbol = v
}

// GetType returns the Type field value
func (o *SystemWaypoint) GetType() WaypointType {
	if o == nil {
		var ret WaypointType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SystemWaypoint) GetTypeOk() (*WaypointType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SystemWaypoint) SetType(v WaypointType) {
	o.Type = v
}

// GetX returns the X field value
func (o *SystemWaypoint) GetX() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.X
}

// GetXOk returns a tuple with the X field value
// and a boolean to check if the value has been set.
func (o *SystemWaypoint) GetXOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.X, true
}

// SetX sets field value
func (o *SystemWaypoint) SetX(v int32) {
	o.X = v
}

// GetY returns the Y field value
func (o *SystemWaypoint) GetY() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Y
}

// GetYOk returns a tuple with the Y field value
// and a boolean to check if the value has been set.
func (o *SystemWaypoint) GetYOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Y, true
}

// SetY sets field value
func (o *SystemWaypoint) SetY(v int32) {
	o.Y = v
}

// GetOrbitals returns the Orbitals field value
func (o *SystemWaypoint) GetOrbitals() []WaypointOrbital {
	if o == nil {
		var ret []WaypointOrbital
		return ret
	}

	return o.Orbitals
}

// GetOrbitalsOk returns a tuple with the Orbitals field value
// and a boolean to check if the value has been set.
func (o *SystemWaypoint) GetOrbitalsOk() ([]WaypointOrbital, bool) {
	if o == nil {
		return nil, false
	}
	return o.Orbitals, true
}

// SetOrbitals sets field value
func (o *SystemWaypoint) SetOrbitals(v []WaypointOrbital) {
	o.Orbitals = v
}

// GetOrbits returns the Orbits field value if set, zero value otherwise.
func (o *SystemWaypoint) GetOrbits() string {
	if o == nil || IsNil(o.Orbits) {
		var ret string
		return ret
	}
	return *o.Orbits
}

// GetOrbitsOk returns a tuple with the Orbits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemWaypoint) GetOrbitsOk() (*string, bool) {
	if o == nil || IsNil(o.Orbits) {
		return nil, false
	}
	return o.Orbits, true
}

// HasOrbits returns a boolean if a field has been set.
func (o *SystemWaypoint) HasOrbits() bool {
	if o != nil && !IsNil(o.Orbits) {
		return true
	}

	return false
}

// SetOrbits gets a reference to the given string and assigns it to the Orbits field.
func (o *SystemWaypoint) SetOrbits(v string) {
	o.Orbits = &v
}

func (o SystemWaypoint) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SystemWaypoint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol
	toSerialize["type"] = o.Type
	toSerialize["x"] = o.X
	toSerialize["y"] = o.Y
	toSerialize["orbitals"] = o.Orbitals
	if !IsNil(o.Orbits) {
		toSerialize["orbits"] = o.Orbits
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SystemWaypoint) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"symbol",
		"type",
		"x",
		"y",
		"orbitals",
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

	varSystemWaypoint := _SystemWaypoint{}

	err = json.Unmarshal(bytes, &varSystemWaypoint)

	if err != nil {
		return err
	}

	*o = SystemWaypoint(varSystemWaypoint)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "type")
		delete(additionalProperties, "x")
		delete(additionalProperties, "y")
		delete(additionalProperties, "orbitals")
		delete(additionalProperties, "orbits")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSystemWaypoint struct {
	value *SystemWaypoint
	isSet bool
}

func (v NullableSystemWaypoint) Get() *SystemWaypoint {
	return v.value
}

func (v *NullableSystemWaypoint) Set(val *SystemWaypoint) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemWaypoint) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemWaypoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemWaypoint(val *SystemWaypoint) *NullableSystemWaypoint {
	return &NullableSystemWaypoint{value: val, isSet: true}
}

func (v NullableSystemWaypoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemWaypoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

