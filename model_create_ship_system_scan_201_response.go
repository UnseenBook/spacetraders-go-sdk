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

// checks if the CreateShipSystemScan201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateShipSystemScan201Response{}

// CreateShipSystemScan201Response struct for CreateShipSystemScan201Response
type CreateShipSystemScan201Response struct {
	Data CreateShipSystemScan201ResponseData `json:"data"`
	AdditionalProperties map[string]interface{}
}

type _CreateShipSystemScan201Response CreateShipSystemScan201Response

// NewCreateShipSystemScan201Response instantiates a new CreateShipSystemScan201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateShipSystemScan201Response(data CreateShipSystemScan201ResponseData) *CreateShipSystemScan201Response {
	this := CreateShipSystemScan201Response{}
	this.Data = data
	return &this
}

// NewCreateShipSystemScan201ResponseWithDefaults instantiates a new CreateShipSystemScan201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateShipSystemScan201ResponseWithDefaults() *CreateShipSystemScan201Response {
	this := CreateShipSystemScan201Response{}
	return &this
}

// GetData returns the Data field value
func (o *CreateShipSystemScan201Response) GetData() CreateShipSystemScan201ResponseData {
	if o == nil {
		var ret CreateShipSystemScan201ResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CreateShipSystemScan201Response) GetDataOk() (*CreateShipSystemScan201ResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CreateShipSystemScan201Response) SetData(v CreateShipSystemScan201ResponseData) {
	o.Data = v
}

func (o CreateShipSystemScan201Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateShipSystemScan201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateShipSystemScan201Response) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varCreateShipSystemScan201Response := _CreateShipSystemScan201Response{}

	err = json.Unmarshal(bytes, &varCreateShipSystemScan201Response)

	if err != nil {
		return err
	}

	*o = CreateShipSystemScan201Response(varCreateShipSystemScan201Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateShipSystemScan201Response struct {
	value *CreateShipSystemScan201Response
	isSet bool
}

func (v NullableCreateShipSystemScan201Response) Get() *CreateShipSystemScan201Response {
	return v.value
}

func (v *NullableCreateShipSystemScan201Response) Set(val *CreateShipSystemScan201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateShipSystemScan201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateShipSystemScan201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateShipSystemScan201Response(val *CreateShipSystemScan201Response) *NullableCreateShipSystemScan201Response {
	return &NullableCreateShipSystemScan201Response{value: val, isSet: true}
}

func (v NullableCreateShipSystemScan201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateShipSystemScan201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


