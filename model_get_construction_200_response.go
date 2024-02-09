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

// checks if the GetConstruction200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetConstruction200Response{}

// GetConstruction200Response 
type GetConstruction200Response struct {
	Data Construction `json:"data"`
	AdditionalProperties map[string]interface{}
}

type _GetConstruction200Response GetConstruction200Response

// NewGetConstruction200Response instantiates a new GetConstruction200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetConstruction200Response(data Construction) *GetConstruction200Response {
	this := GetConstruction200Response{}
	this.Data = data
	return &this
}

// NewGetConstruction200ResponseWithDefaults instantiates a new GetConstruction200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetConstruction200ResponseWithDefaults() *GetConstruction200Response {
	this := GetConstruction200Response{}
	return &this
}

// GetData returns the Data field value
func (o *GetConstruction200Response) GetData() Construction {
	if o == nil {
		var ret Construction
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetConstruction200Response) GetDataOk() (*Construction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetConstruction200Response) SetData(v Construction) {
	o.Data = v
}

func (o GetConstruction200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetConstruction200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GetConstruction200Response) UnmarshalJSON(data []byte) (err error) {
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

	varGetConstruction200Response := _GetConstruction200Response{}

	err = json.Unmarshal(data, &varGetConstruction200Response)

	if err != nil {
		return err
	}

	*o = GetConstruction200Response(varGetConstruction200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetConstruction200Response struct {
	value *GetConstruction200Response
	isSet bool
}

func (v NullableGetConstruction200Response) Get() *GetConstruction200Response {
	return v.value
}

func (v *NullableGetConstruction200Response) Set(val *GetConstruction200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetConstruction200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetConstruction200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetConstruction200Response(val *GetConstruction200Response) *NullableGetConstruction200Response {
	return &NullableGetConstruction200Response{value: val, isSet: true}
}

func (v NullableGetConstruction200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetConstruction200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


