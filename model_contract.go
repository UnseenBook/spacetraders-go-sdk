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
	"time"
	"fmt"
)

// checks if the Contract type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Contract{}

// Contract Contract details.
type Contract struct {
	// ID of the contract.
	Id string `json:"id"`
	// The symbol of the faction that this contract is for.
	FactionSymbol string `json:"factionSymbol"`
	// Type of contract.
	Type string `json:"type"`
	Terms ContractTerms `json:"terms"`
	// Whether the contract has been accepted by the agent
	Accepted bool `json:"accepted"`
	// Whether the contract has been fulfilled
	Fulfilled bool `json:"fulfilled"`
	// Deprecated in favor of deadlineToAccept
	// Deprecated
	Expiration time.Time `json:"expiration"`
	// The time at which the contract is no longer available to be accepted
	DeadlineToAccept *time.Time `json:"deadlineToAccept,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Contract Contract

// NewContract instantiates a new Contract object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContract(id string, factionSymbol string, type_ string, terms ContractTerms, accepted bool, fulfilled bool, expiration time.Time) *Contract {
	this := Contract{}
	this.Id = id
	this.FactionSymbol = factionSymbol
	this.Type = type_
	this.Terms = terms
	this.Accepted = accepted
	this.Fulfilled = fulfilled
	this.Expiration = expiration
	return &this
}

// NewContractWithDefaults instantiates a new Contract object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractWithDefaults() *Contract {
	this := Contract{}
	var accepted bool = false
	this.Accepted = accepted
	var fulfilled bool = false
	this.Fulfilled = fulfilled
	return &this
}

// GetId returns the Id field value
func (o *Contract) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Contract) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Contract) SetId(v string) {
	o.Id = v
}

// GetFactionSymbol returns the FactionSymbol field value
func (o *Contract) GetFactionSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FactionSymbol
}

// GetFactionSymbolOk returns a tuple with the FactionSymbol field value
// and a boolean to check if the value has been set.
func (o *Contract) GetFactionSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FactionSymbol, true
}

// SetFactionSymbol sets field value
func (o *Contract) SetFactionSymbol(v string) {
	o.FactionSymbol = v
}

// GetType returns the Type field value
func (o *Contract) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Contract) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Contract) SetType(v string) {
	o.Type = v
}

// GetTerms returns the Terms field value
func (o *Contract) GetTerms() ContractTerms {
	if o == nil {
		var ret ContractTerms
		return ret
	}

	return o.Terms
}

// GetTermsOk returns a tuple with the Terms field value
// and a boolean to check if the value has been set.
func (o *Contract) GetTermsOk() (*ContractTerms, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Terms, true
}

// SetTerms sets field value
func (o *Contract) SetTerms(v ContractTerms) {
	o.Terms = v
}

// GetAccepted returns the Accepted field value
func (o *Contract) GetAccepted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Accepted
}

// GetAcceptedOk returns a tuple with the Accepted field value
// and a boolean to check if the value has been set.
func (o *Contract) GetAcceptedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Accepted, true
}

// SetAccepted sets field value
func (o *Contract) SetAccepted(v bool) {
	o.Accepted = v
}

// GetFulfilled returns the Fulfilled field value
func (o *Contract) GetFulfilled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Fulfilled
}

// GetFulfilledOk returns a tuple with the Fulfilled field value
// and a boolean to check if the value has been set.
func (o *Contract) GetFulfilledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fulfilled, true
}

// SetFulfilled sets field value
func (o *Contract) SetFulfilled(v bool) {
	o.Fulfilled = v
}

// GetExpiration returns the Expiration field value
// Deprecated
func (o *Contract) GetExpiration() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *Contract) GetExpirationOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expiration, true
}

// SetExpiration sets field value
// Deprecated
func (o *Contract) SetExpiration(v time.Time) {
	o.Expiration = v
}

// GetDeadlineToAccept returns the DeadlineToAccept field value if set, zero value otherwise.
func (o *Contract) GetDeadlineToAccept() time.Time {
	if o == nil || IsNil(o.DeadlineToAccept) {
		var ret time.Time
		return ret
	}
	return *o.DeadlineToAccept
}

// GetDeadlineToAcceptOk returns a tuple with the DeadlineToAccept field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Contract) GetDeadlineToAcceptOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DeadlineToAccept) {
		return nil, false
	}
	return o.DeadlineToAccept, true
}

// HasDeadlineToAccept returns a boolean if a field has been set.
func (o *Contract) HasDeadlineToAccept() bool {
	if o != nil && !IsNil(o.DeadlineToAccept) {
		return true
	}

	return false
}

// SetDeadlineToAccept gets a reference to the given time.Time and assigns it to the DeadlineToAccept field.
func (o *Contract) SetDeadlineToAccept(v time.Time) {
	o.DeadlineToAccept = &v
}

func (o Contract) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Contract) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["factionSymbol"] = o.FactionSymbol
	toSerialize["type"] = o.Type
	toSerialize["terms"] = o.Terms
	toSerialize["accepted"] = o.Accepted
	toSerialize["fulfilled"] = o.Fulfilled
	toSerialize["expiration"] = o.Expiration
	if !IsNil(o.DeadlineToAccept) {
		toSerialize["deadlineToAccept"] = o.DeadlineToAccept
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Contract) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"factionSymbol",
		"type",
		"terms",
		"accepted",
		"fulfilled",
		"expiration",
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

	varContract := _Contract{}

	err = json.Unmarshal(data, &varContract)

	if err != nil {
		return err
	}

	*o = Contract(varContract)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "factionSymbol")
		delete(additionalProperties, "type")
		delete(additionalProperties, "terms")
		delete(additionalProperties, "accepted")
		delete(additionalProperties, "fulfilled")
		delete(additionalProperties, "expiration")
		delete(additionalProperties, "deadlineToAccept")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableContract struct {
	value *Contract
	isSet bool
}

func (v NullableContract) Get() *Contract {
	return v.value
}

func (v *NullableContract) Set(val *Contract) {
	v.value = val
	v.isSet = true
}

func (v NullableContract) IsSet() bool {
	return v.isSet
}

func (v *NullableContract) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContract(val *Contract) *NullableContract {
	return &NullableContract{value: val, isSet: true}
}

func (v NullableContract) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContract) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


