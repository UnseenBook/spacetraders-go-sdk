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

// checks if the RefuelShip200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RefuelShip200Response{}

// RefuelShip200Response 
type RefuelShip200Response struct {
	Data RefuelShip200ResponseData `json:"data"`
	AdditionalProperties map[string]interface{}
}

type _RefuelShip200Response RefuelShip200Response

// NewRefuelShip200Response instantiates a new RefuelShip200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefuelShip200Response(data RefuelShip200ResponseData) *RefuelShip200Response {
	this := RefuelShip200Response{}
	this.Data = data
	return &this
}

// NewRefuelShip200ResponseWithDefaults instantiates a new RefuelShip200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefuelShip200ResponseWithDefaults() *RefuelShip200Response {
	this := RefuelShip200Response{}
	return &this
}

// GetData returns the Data field value
func (o *RefuelShip200Response) GetData() RefuelShip200ResponseData {
	if o == nil {
		var ret RefuelShip200ResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *RefuelShip200Response) GetDataOk() (*RefuelShip200ResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *RefuelShip200Response) SetData(v RefuelShip200ResponseData) {
	o.Data = v
}

func (o RefuelShip200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RefuelShip200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RefuelShip200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varRefuelShip200Response := _RefuelShip200Response{}

	err = json.Unmarshal(data, &varRefuelShip200Response)

	if err != nil {
		return err
	}

	*o = RefuelShip200Response(varRefuelShip200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRefuelShip200Response struct {
	value *RefuelShip200Response
	isSet bool
}

func (v NullableRefuelShip200Response) Get() *RefuelShip200Response {
	return v.value
}

func (v *NullableRefuelShip200Response) Set(val *RefuelShip200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableRefuelShip200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableRefuelShip200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefuelShip200Response(val *RefuelShip200Response) *NullableRefuelShip200Response {
	return &NullableRefuelShip200Response{value: val, isSet: true}
}

func (v NullableRefuelShip200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefuelShip200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


