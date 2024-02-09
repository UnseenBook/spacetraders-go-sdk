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

// checks if the Faction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Faction{}

// Faction Faction details.
type Faction struct {
	Symbol FactionSymbol `json:"symbol"`
	// Name of the faction.
	Name string `json:"name"`
	// Description of the faction.
	Description string `json:"description"`
	// The waypoint in which the faction's HQ is located in.
	Headquarters string `json:"headquarters"`
	// List of traits that define this faction.
	Traits []FactionTrait `json:"traits"`
	// Whether or not the faction is currently recruiting new agents.
	IsRecruiting bool `json:"isRecruiting"`
	AdditionalProperties map[string]interface{}
}

type _Faction Faction

// NewFaction instantiates a new Faction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFaction(symbol FactionSymbol, name string, description string, headquarters string, traits []FactionTrait, isRecruiting bool) *Faction {
	this := Faction{}
	this.Symbol = symbol
	this.Name = name
	this.Description = description
	this.Headquarters = headquarters
	this.Traits = traits
	this.IsRecruiting = isRecruiting
	return &this
}

// NewFactionWithDefaults instantiates a new Faction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFactionWithDefaults() *Faction {
	this := Faction{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *Faction) GetSymbol() FactionSymbol {
	if o == nil {
		var ret FactionSymbol
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *Faction) GetSymbolOk() (*FactionSymbol, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *Faction) SetSymbol(v FactionSymbol) {
	o.Symbol = v
}

// GetName returns the Name field value
func (o *Faction) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Faction) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Faction) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *Faction) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *Faction) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *Faction) SetDescription(v string) {
	o.Description = v
}

// GetHeadquarters returns the Headquarters field value
func (o *Faction) GetHeadquarters() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Headquarters
}

// GetHeadquartersOk returns a tuple with the Headquarters field value
// and a boolean to check if the value has been set.
func (o *Faction) GetHeadquartersOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Headquarters, true
}

// SetHeadquarters sets field value
func (o *Faction) SetHeadquarters(v string) {
	o.Headquarters = v
}

// GetTraits returns the Traits field value
func (o *Faction) GetTraits() []FactionTrait {
	if o == nil {
		var ret []FactionTrait
		return ret
	}

	return o.Traits
}

// GetTraitsOk returns a tuple with the Traits field value
// and a boolean to check if the value has been set.
func (o *Faction) GetTraitsOk() ([]FactionTrait, bool) {
	if o == nil {
		return nil, false
	}
	return o.Traits, true
}

// SetTraits sets field value
func (o *Faction) SetTraits(v []FactionTrait) {
	o.Traits = v
}

// GetIsRecruiting returns the IsRecruiting field value
func (o *Faction) GetIsRecruiting() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsRecruiting
}

// GetIsRecruitingOk returns a tuple with the IsRecruiting field value
// and a boolean to check if the value has been set.
func (o *Faction) GetIsRecruitingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsRecruiting, true
}

// SetIsRecruiting sets field value
func (o *Faction) SetIsRecruiting(v bool) {
	o.IsRecruiting = v
}

func (o Faction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Faction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	toSerialize["headquarters"] = o.Headquarters
	toSerialize["traits"] = o.Traits
	toSerialize["isRecruiting"] = o.IsRecruiting

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Faction) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"symbol",
		"name",
		"description",
		"headquarters",
		"traits",
		"isRecruiting",
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

	varFaction := _Faction{}

	err = json.Unmarshal(data, &varFaction)

	if err != nil {
		return err
	}

	*o = Faction(varFaction)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "symbol")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "headquarters")
		delete(additionalProperties, "traits")
		delete(additionalProperties, "isRecruiting")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFaction struct {
	value *Faction
	isSet bool
}

func (v NullableFaction) Get() *Faction {
	return v.value
}

func (v *NullableFaction) Set(val *Faction) {
	v.value = val
	v.isSet = true
}

func (v NullableFaction) IsSet() bool {
	return v.isSet
}

func (v *NullableFaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFaction(val *Faction) *NullableFaction {
	return &NullableFaction{value: val, isSet: true}
}

func (v NullableFaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


