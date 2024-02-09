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

// checks if the ShipyardShipCrew type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipyardShipCrew{}

// ShipyardShipCrew struct for ShipyardShipCrew
type ShipyardShipCrew struct {
	Required int32 `json:"required"`
	Capacity int32 `json:"capacity"`
	AdditionalProperties map[string]interface{}
}

type _ShipyardShipCrew ShipyardShipCrew

// NewShipyardShipCrew instantiates a new ShipyardShipCrew object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipyardShipCrew(required int32, capacity int32) *ShipyardShipCrew {
	this := ShipyardShipCrew{}
	this.Required = required
	this.Capacity = capacity
	return &this
}

// NewShipyardShipCrewWithDefaults instantiates a new ShipyardShipCrew object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipyardShipCrewWithDefaults() *ShipyardShipCrew {
	this := ShipyardShipCrew{}
	return &this
}

// GetRequired returns the Required field value
func (o *ShipyardShipCrew) GetRequired() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Required
}

// GetRequiredOk returns a tuple with the Required field value
// and a boolean to check if the value has been set.
func (o *ShipyardShipCrew) GetRequiredOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Required, true
}

// SetRequired sets field value
func (o *ShipyardShipCrew) SetRequired(v int32) {
	o.Required = v
}

// GetCapacity returns the Capacity field value
func (o *ShipyardShipCrew) GetCapacity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value
// and a boolean to check if the value has been set.
func (o *ShipyardShipCrew) GetCapacityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Capacity, true
}

// SetCapacity sets field value
func (o *ShipyardShipCrew) SetCapacity(v int32) {
	o.Capacity = v
}

func (o ShipyardShipCrew) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShipyardShipCrew) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["required"] = o.Required
	toSerialize["capacity"] = o.Capacity

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ShipyardShipCrew) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"required",
		"capacity",
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

	varShipyardShipCrew := _ShipyardShipCrew{}

	err = json.Unmarshal(data, &varShipyardShipCrew)

	if err != nil {
		return err
	}

	*o = ShipyardShipCrew(varShipyardShipCrew)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "required")
		delete(additionalProperties, "capacity")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableShipyardShipCrew struct {
	value *ShipyardShipCrew
	isSet bool
}

func (v NullableShipyardShipCrew) Get() *ShipyardShipCrew {
	return v.value
}

func (v *NullableShipyardShipCrew) Set(val *ShipyardShipCrew) {
	v.value = val
	v.isSet = true
}

func (v NullableShipyardShipCrew) IsSet() bool {
	return v.isSet
}

func (v *NullableShipyardShipCrew) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipyardShipCrew(val *ShipyardShipCrew) *NullableShipyardShipCrew {
	return &NullableShipyardShipCrew{value: val, isSet: true}
}

func (v NullableShipyardShipCrew) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipyardShipCrew) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


