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

// checks if the ShipCrew type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipCrew{}

// ShipCrew The ship's crew service and maintain the ship's systems and equipment.
type ShipCrew struct {
	// The current number of crew members on the ship.
	Current int32 `json:"current"`
	// The minimum number of crew members required to maintain the ship.
	Required int32 `json:"required"`
	// The maximum number of crew members the ship can support.
	Capacity int32 `json:"capacity"`
	// The rotation of crew shifts. A stricter shift improves the ship's performance. A more relaxed shift improves the crew's morale.
	Rotation string `json:"rotation"`
	// A rough measure of the crew's morale. A higher morale means the crew is happier and more productive. A lower morale means the ship is more prone to accidents.
	Morale int32 `json:"morale"`
	// The amount of credits per crew member paid per hour. Wages are paid when a ship docks at a civilized waypoint.
	Wages int32 `json:"wages"`
	AdditionalProperties map[string]interface{}
}

type _ShipCrew ShipCrew

// NewShipCrew instantiates a new ShipCrew object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipCrew(current int32, required int32, capacity int32, rotation string, morale int32, wages int32) *ShipCrew {
	this := ShipCrew{}
	this.Current = current
	this.Required = required
	this.Capacity = capacity
	this.Rotation = rotation
	this.Morale = morale
	this.Wages = wages
	return &this
}

// NewShipCrewWithDefaults instantiates a new ShipCrew object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipCrewWithDefaults() *ShipCrew {
	this := ShipCrew{}
	var rotation string = "STRICT"
	this.Rotation = rotation
	return &this
}

// GetCurrent returns the Current field value
func (o *ShipCrew) GetCurrent() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Current
}

// GetCurrentOk returns a tuple with the Current field value
// and a boolean to check if the value has been set.
func (o *ShipCrew) GetCurrentOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Current, true
}

// SetCurrent sets field value
func (o *ShipCrew) SetCurrent(v int32) {
	o.Current = v
}

// GetRequired returns the Required field value
func (o *ShipCrew) GetRequired() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Required
}

// GetRequiredOk returns a tuple with the Required field value
// and a boolean to check if the value has been set.
func (o *ShipCrew) GetRequiredOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Required, true
}

// SetRequired sets field value
func (o *ShipCrew) SetRequired(v int32) {
	o.Required = v
}

// GetCapacity returns the Capacity field value
func (o *ShipCrew) GetCapacity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value
// and a boolean to check if the value has been set.
func (o *ShipCrew) GetCapacityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Capacity, true
}

// SetCapacity sets field value
func (o *ShipCrew) SetCapacity(v int32) {
	o.Capacity = v
}

// GetRotation returns the Rotation field value
func (o *ShipCrew) GetRotation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Rotation
}

// GetRotationOk returns a tuple with the Rotation field value
// and a boolean to check if the value has been set.
func (o *ShipCrew) GetRotationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rotation, true
}

// SetRotation sets field value
func (o *ShipCrew) SetRotation(v string) {
	o.Rotation = v
}

// GetMorale returns the Morale field value
func (o *ShipCrew) GetMorale() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Morale
}

// GetMoraleOk returns a tuple with the Morale field value
// and a boolean to check if the value has been set.
func (o *ShipCrew) GetMoraleOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Morale, true
}

// SetMorale sets field value
func (o *ShipCrew) SetMorale(v int32) {
	o.Morale = v
}

// GetWages returns the Wages field value
func (o *ShipCrew) GetWages() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Wages
}

// GetWagesOk returns a tuple with the Wages field value
// and a boolean to check if the value has been set.
func (o *ShipCrew) GetWagesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Wages, true
}

// SetWages sets field value
func (o *ShipCrew) SetWages(v int32) {
	o.Wages = v
}

func (o ShipCrew) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShipCrew) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["current"] = o.Current
	toSerialize["required"] = o.Required
	toSerialize["capacity"] = o.Capacity
	toSerialize["rotation"] = o.Rotation
	toSerialize["morale"] = o.Morale
	toSerialize["wages"] = o.Wages

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ShipCrew) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"current",
		"required",
		"capacity",
		"rotation",
		"morale",
		"wages",
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

	varShipCrew := _ShipCrew{}

	err = json.Unmarshal(data, &varShipCrew)

	if err != nil {
		return err
	}

	*o = ShipCrew(varShipCrew)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "current")
		delete(additionalProperties, "required")
		delete(additionalProperties, "capacity")
		delete(additionalProperties, "rotation")
		delete(additionalProperties, "morale")
		delete(additionalProperties, "wages")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableShipCrew struct {
	value *ShipCrew
	isSet bool
}

func (v NullableShipCrew) Get() *ShipCrew {
	return v.value
}

func (v *NullableShipCrew) Set(val *ShipCrew) {
	v.value = val
	v.isSet = true
}

func (v NullableShipCrew) IsSet() bool {
	return v.isSet
}

func (v *NullableShipCrew) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipCrew(val *ShipCrew) *NullableShipCrew {
	return &NullableShipCrew{value: val, isSet: true}
}

func (v NullableShipCrew) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipCrew) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


