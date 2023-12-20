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

// checks if the NegotiateContract200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NegotiateContract200ResponseData{}

// NegotiateContract200ResponseData struct for NegotiateContract200ResponseData
type NegotiateContract200ResponseData struct {
	Contract Contract `json:"contract"`
	AdditionalProperties map[string]interface{}
}

type _NegotiateContract200ResponseData NegotiateContract200ResponseData

// NewNegotiateContract200ResponseData instantiates a new NegotiateContract200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNegotiateContract200ResponseData(contract Contract) *NegotiateContract200ResponseData {
	this := NegotiateContract200ResponseData{}
	this.Contract = contract
	return &this
}

// NewNegotiateContract200ResponseDataWithDefaults instantiates a new NegotiateContract200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNegotiateContract200ResponseDataWithDefaults() *NegotiateContract200ResponseData {
	this := NegotiateContract200ResponseData{}
	return &this
}

// GetContract returns the Contract field value
func (o *NegotiateContract200ResponseData) GetContract() Contract {
	if o == nil {
		var ret Contract
		return ret
	}

	return o.Contract
}

// GetContractOk returns a tuple with the Contract field value
// and a boolean to check if the value has been set.
func (o *NegotiateContract200ResponseData) GetContractOk() (*Contract, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Contract, true
}

// SetContract sets field value
func (o *NegotiateContract200ResponseData) SetContract(v Contract) {
	o.Contract = v
}

func (o NegotiateContract200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NegotiateContract200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["contract"] = o.Contract

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NegotiateContract200ResponseData) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"contract",
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

	varNegotiateContract200ResponseData := _NegotiateContract200ResponseData{}

	err = json.Unmarshal(bytes, &varNegotiateContract200ResponseData)

	if err != nil {
		return err
	}

	*o = NegotiateContract200ResponseData(varNegotiateContract200ResponseData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "contract")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNegotiateContract200ResponseData struct {
	value *NegotiateContract200ResponseData
	isSet bool
}

func (v NullableNegotiateContract200ResponseData) Get() *NegotiateContract200ResponseData {
	return v.value
}

func (v *NullableNegotiateContract200ResponseData) Set(val *NegotiateContract200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableNegotiateContract200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableNegotiateContract200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNegotiateContract200ResponseData(val *NegotiateContract200ResponseData) *NullableNegotiateContract200ResponseData {
	return &NullableNegotiateContract200ResponseData{value: val, isSet: true}
}

func (v NullableNegotiateContract200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNegotiateContract200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


