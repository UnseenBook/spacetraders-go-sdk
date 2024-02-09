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

// checks if the Meta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Meta{}

// Meta Meta details for pagination.
type Meta struct {
	// Shows the total amount of items of this kind that exist.
	Total int32 `json:"total"`
	// A page denotes an amount of items, offset from the first item. Each page holds an amount of items equal to the `limit`.
	Page int32 `json:"page"`
	// The amount of items in each page. Limits how many items can be fetched at once.
	Limit int32 `json:"limit"`
	AdditionalProperties map[string]interface{}
}

type _Meta Meta

// NewMeta instantiates a new Meta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMeta(total int32, page int32, limit int32) *Meta {
	this := Meta{}
	this.Total = total
	this.Page = page
	this.Limit = limit
	return &this
}

// NewMetaWithDefaults instantiates a new Meta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetaWithDefaults() *Meta {
	this := Meta{}
	var page int32 = 1
	this.Page = page
	var limit int32 = 10
	this.Limit = limit
	return &this
}

// GetTotal returns the Total field value
func (o *Meta) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *Meta) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *Meta) SetTotal(v int32) {
	o.Total = v
}

// GetPage returns the Page field value
func (o *Meta) GetPage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *Meta) GetPageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value
func (o *Meta) SetPage(v int32) {
	o.Page = v
}

// GetLimit returns the Limit field value
func (o *Meta) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *Meta) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *Meta) SetLimit(v int32) {
	o.Limit = v
}

func (o Meta) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Meta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["total"] = o.Total
	toSerialize["page"] = o.Page
	toSerialize["limit"] = o.Limit

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Meta) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"total",
		"page",
		"limit",
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

	varMeta := _Meta{}

	err = json.Unmarshal(data, &varMeta)

	if err != nil {
		return err
	}

	*o = Meta(varMeta)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "total")
		delete(additionalProperties, "page")
		delete(additionalProperties, "limit")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMeta struct {
	value *Meta
	isSet bool
}

func (v NullableMeta) Get() *Meta {
	return v.value
}

func (v *NullableMeta) Set(val *Meta) {
	v.value = val
	v.isSet = true
}

func (v NullableMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMeta(val *Meta) *NullableMeta {
	return &NullableMeta{value: val, isSet: true}
}

func (v NullableMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


