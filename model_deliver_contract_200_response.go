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

// checks if the DeliverContract200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeliverContract200Response{}

// DeliverContract200Response 
type DeliverContract200Response struct {
	Data DeliverContract200ResponseData `json:"data"`
	AdditionalProperties map[string]interface{}
}

type _DeliverContract200Response DeliverContract200Response

// NewDeliverContract200Response instantiates a new DeliverContract200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliverContract200Response(data DeliverContract200ResponseData) *DeliverContract200Response {
	this := DeliverContract200Response{}
	this.Data = data
	return &this
}

// NewDeliverContract200ResponseWithDefaults instantiates a new DeliverContract200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliverContract200ResponseWithDefaults() *DeliverContract200Response {
	this := DeliverContract200Response{}
	return &this
}

// GetData returns the Data field value
func (o *DeliverContract200Response) GetData() DeliverContract200ResponseData {
	if o == nil {
		var ret DeliverContract200ResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *DeliverContract200Response) GetDataOk() (*DeliverContract200ResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *DeliverContract200Response) SetData(v DeliverContract200ResponseData) {
	o.Data = v
}

func (o DeliverContract200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeliverContract200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeliverContract200Response) UnmarshalJSON(bytes []byte) (err error) {
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

	varDeliverContract200Response := _DeliverContract200Response{}

	err = json.Unmarshal(bytes, &varDeliverContract200Response)

	if err != nil {
		return err
	}

	*o = DeliverContract200Response(varDeliverContract200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeliverContract200Response struct {
	value *DeliverContract200Response
	isSet bool
}

func (v NullableDeliverContract200Response) Get() *DeliverContract200Response {
	return v.value
}

func (v *NullableDeliverContract200Response) Set(val *DeliverContract200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliverContract200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliverContract200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliverContract200Response(val *DeliverContract200Response) *NullableDeliverContract200Response {
	return &NullableDeliverContract200Response{value: val, isSet: true}
}

func (v NullableDeliverContract200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliverContract200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


