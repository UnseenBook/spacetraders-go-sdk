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

// GetSystemWaypointsTraitsParameter - struct for GetSystemWaypointsTraitsParameter
type GetSystemWaypointsTraitsParameter struct {
	WaypointTraitSymbol *WaypointTraitSymbol
	ArrayOfWaypointTraitSymbol *[]WaypointTraitSymbol
}

// WaypointTraitSymbolAsGetSystemWaypointsTraitsParameter is a convenience function that returns WaypointTraitSymbol wrapped in GetSystemWaypointsTraitsParameter
func WaypointTraitSymbolAsGetSystemWaypointsTraitsParameter(v *WaypointTraitSymbol) GetSystemWaypointsTraitsParameter {
	return GetSystemWaypointsTraitsParameter{
		WaypointTraitSymbol: v,
	}
}

// []WaypointTraitSymbolAsGetSystemWaypointsTraitsParameter is a convenience function that returns []WaypointTraitSymbol wrapped in GetSystemWaypointsTraitsParameter
func ArrayOfWaypointTraitSymbolAsGetSystemWaypointsTraitsParameter(v *[]WaypointTraitSymbol) GetSystemWaypointsTraitsParameter {
	return GetSystemWaypointsTraitsParameter{
		ArrayOfWaypointTraitSymbol: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetSystemWaypointsTraitsParameter) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into WaypointTraitSymbol
	err = newStrictDecoder(data).Decode(&dst.WaypointTraitSymbol)
	if err == nil {
		jsonWaypointTraitSymbol, _ := json.Marshal(dst.WaypointTraitSymbol)
		if string(jsonWaypointTraitSymbol) == "{}" { // empty struct
			dst.WaypointTraitSymbol = nil
		} else {
			match++
		}
	} else {
		dst.WaypointTraitSymbol = nil
	}

	// try to unmarshal data into ArrayOfWaypointTraitSymbol
	err = newStrictDecoder(data).Decode(&dst.ArrayOfWaypointTraitSymbol)
	if err == nil {
		jsonArrayOfWaypointTraitSymbol, _ := json.Marshal(dst.ArrayOfWaypointTraitSymbol)
		if string(jsonArrayOfWaypointTraitSymbol) == "{}" { // empty struct
			dst.ArrayOfWaypointTraitSymbol = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfWaypointTraitSymbol = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.WaypointTraitSymbol = nil
		dst.ArrayOfWaypointTraitSymbol = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetSystemWaypointsTraitsParameter)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetSystemWaypointsTraitsParameter)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetSystemWaypointsTraitsParameter) MarshalJSON() ([]byte, error) {
	if src.WaypointTraitSymbol != nil {
		return json.Marshal(&src.WaypointTraitSymbol)
	}

	if src.ArrayOfWaypointTraitSymbol != nil {
		return json.Marshal(&src.ArrayOfWaypointTraitSymbol)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetSystemWaypointsTraitsParameter) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.WaypointTraitSymbol != nil {
		return obj.WaypointTraitSymbol
	}

	if obj.ArrayOfWaypointTraitSymbol != nil {
		return obj.ArrayOfWaypointTraitSymbol
	}

	// all schemas are nil
	return nil
}

type NullableGetSystemWaypointsTraitsParameter struct {
	value *GetSystemWaypointsTraitsParameter
	isSet bool
}

func (v NullableGetSystemWaypointsTraitsParameter) Get() *GetSystemWaypointsTraitsParameter {
	return v.value
}

func (v *NullableGetSystemWaypointsTraitsParameter) Set(val *GetSystemWaypointsTraitsParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSystemWaypointsTraitsParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSystemWaypointsTraitsParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSystemWaypointsTraitsParameter(val *GetSystemWaypointsTraitsParameter) *NullableGetSystemWaypointsTraitsParameter {
	return &NullableGetSystemWaypointsTraitsParameter{value: val, isSet: true}
}

func (v NullableGetSystemWaypointsTraitsParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSystemWaypointsTraitsParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


