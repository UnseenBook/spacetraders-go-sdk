# Go API client for stsdk

SpaceTraders is an open-universe game and learning platform that offers a set of HTTP endpoints to control a fleet of ships and explore a multiplayer universe.

The API is documented using [OpenAPI](https://github.com/SpaceTradersAPI/api-docs). You can send your first request right here in your browser to check the status of the game server.

```json http
{
  \"method\": \"GET\",
  \"url\": \"https://api.spacetraders.io/v2\",
}
```

Unlike a traditional game, SpaceTraders does not have a first-party client or app to play the game. Instead, you can use the API to build your own client, write a script to automate your ships, or try an app built by the community.

We have a [Discord channel](https://discord.com/invite/jh6zurdWk5) where you can share your projects, ask questions, and get help from other players.




## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 2.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import stsdk "github.com/UnseenBook/spacetraders-go-sdk"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `stsdk.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), stsdk.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `stsdk.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), stsdk.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `stsdk.ContextOperationServerIndices` and `stsdk.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), stsdk.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), stsdk.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.spacetraders.io/v2*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AgentsAPI* | [**GetAgent**](docs/AgentsAPI.md#getagent) | **Get** /agents/{agentSymbol} | Get Public Agent
*AgentsAPI* | [**GetAgents**](docs/AgentsAPI.md#getagents) | **Get** /agents | List Agents
*AgentsAPI* | [**GetMyAgent**](docs/AgentsAPI.md#getmyagent) | **Get** /my/agent | Get Agent
*ContractsAPI* | [**AcceptContract**](docs/ContractsAPI.md#acceptcontract) | **Post** /my/contracts/{contractId}/accept | Accept Contract
*ContractsAPI* | [**DeliverContract**](docs/ContractsAPI.md#delivercontract) | **Post** /my/contracts/{contractId}/deliver | Deliver Cargo to Contract
*ContractsAPI* | [**FulfillContract**](docs/ContractsAPI.md#fulfillcontract) | **Post** /my/contracts/{contractId}/fulfill | Fulfill Contract
*ContractsAPI* | [**GetContract**](docs/ContractsAPI.md#getcontract) | **Get** /my/contracts/{contractId} | Get Contract
*ContractsAPI* | [**GetContracts**](docs/ContractsAPI.md#getcontracts) | **Get** /my/contracts | List Contracts
*DefaultAPI* | [**GetStatus**](docs/DefaultAPI.md#getstatus) | **Get** / | Get Status
*DefaultAPI* | [**Register**](docs/DefaultAPI.md#register) | **Post** /register | Register New Agent
*FactionsAPI* | [**GetFaction**](docs/FactionsAPI.md#getfaction) | **Get** /factions/{factionSymbol} | Get Faction
*FactionsAPI* | [**GetFactions**](docs/FactionsAPI.md#getfactions) | **Get** /factions | List Factions
*FleetAPI* | [**CreateChart**](docs/FleetAPI.md#createchart) | **Post** /my/ships/{shipSymbol}/chart | Create Chart
*FleetAPI* | [**CreateShipShipScan**](docs/FleetAPI.md#createshipshipscan) | **Post** /my/ships/{shipSymbol}/scan/ships | Scan Ships
*FleetAPI* | [**CreateShipSystemScan**](docs/FleetAPI.md#createshipsystemscan) | **Post** /my/ships/{shipSymbol}/scan/systems | Scan Systems
*FleetAPI* | [**CreateShipWaypointScan**](docs/FleetAPI.md#createshipwaypointscan) | **Post** /my/ships/{shipSymbol}/scan/waypoints | Scan Waypoints
*FleetAPI* | [**CreateSurvey**](docs/FleetAPI.md#createsurvey) | **Post** /my/ships/{shipSymbol}/survey | Create Survey
*FleetAPI* | [**DockShip**](docs/FleetAPI.md#dockship) | **Post** /my/ships/{shipSymbol}/dock | Dock Ship
*FleetAPI* | [**ExtractResources**](docs/FleetAPI.md#extractresources) | **Post** /my/ships/{shipSymbol}/extract | Extract Resources
*FleetAPI* | [**ExtractResourcesWithSurvey**](docs/FleetAPI.md#extractresourceswithsurvey) | **Post** /my/ships/{shipSymbol}/extract/survey | Extract Resources with Survey
*FleetAPI* | [**GetMounts**](docs/FleetAPI.md#getmounts) | **Get** /my/ships/{shipSymbol}/mounts | Get Mounts
*FleetAPI* | [**GetMyShip**](docs/FleetAPI.md#getmyship) | **Get** /my/ships/{shipSymbol} | Get Ship
*FleetAPI* | [**GetMyShipCargo**](docs/FleetAPI.md#getmyshipcargo) | **Get** /my/ships/{shipSymbol}/cargo | Get Ship Cargo
*FleetAPI* | [**GetMyShips**](docs/FleetAPI.md#getmyships) | **Get** /my/ships | List Ships
*FleetAPI* | [**GetShipCooldown**](docs/FleetAPI.md#getshipcooldown) | **Get** /my/ships/{shipSymbol}/cooldown | Get Ship Cooldown
*FleetAPI* | [**GetShipNav**](docs/FleetAPI.md#getshipnav) | **Get** /my/ships/{shipSymbol}/nav | Get Ship Nav
*FleetAPI* | [**InstallMount**](docs/FleetAPI.md#installmount) | **Post** /my/ships/{shipSymbol}/mounts/install | Install Mount
*FleetAPI* | [**Jettison**](docs/FleetAPI.md#jettison) | **Post** /my/ships/{shipSymbol}/jettison | Jettison Cargo
*FleetAPI* | [**JumpShip**](docs/FleetAPI.md#jumpship) | **Post** /my/ships/{shipSymbol}/jump | Jump Ship
*FleetAPI* | [**NavigateShip**](docs/FleetAPI.md#navigateship) | **Post** /my/ships/{shipSymbol}/navigate | Navigate Ship
*FleetAPI* | [**NegotiateContract**](docs/FleetAPI.md#negotiatecontract) | **Post** /my/ships/{shipSymbol}/negotiate/contract | Negotiate Contract
*FleetAPI* | [**OrbitShip**](docs/FleetAPI.md#orbitship) | **Post** /my/ships/{shipSymbol}/orbit | Orbit Ship
*FleetAPI* | [**PatchShipNav**](docs/FleetAPI.md#patchshipnav) | **Patch** /my/ships/{shipSymbol}/nav | Patch Ship Nav
*FleetAPI* | [**PurchaseCargo**](docs/FleetAPI.md#purchasecargo) | **Post** /my/ships/{shipSymbol}/purchase | Purchase Cargo
*FleetAPI* | [**PurchaseShip**](docs/FleetAPI.md#purchaseship) | **Post** /my/ships | Purchase Ship
*FleetAPI* | [**RefuelShip**](docs/FleetAPI.md#refuelship) | **Post** /my/ships/{shipSymbol}/refuel | Refuel Ship
*FleetAPI* | [**RemoveMount**](docs/FleetAPI.md#removemount) | **Post** /my/ships/{shipSymbol}/mounts/remove | Remove Mount
*FleetAPI* | [**SellCargo**](docs/FleetAPI.md#sellcargo) | **Post** /my/ships/{shipSymbol}/sell | Sell Cargo
*FleetAPI* | [**ShipRefine**](docs/FleetAPI.md#shiprefine) | **Post** /my/ships/{shipSymbol}/refine | Ship Refine
*FleetAPI* | [**SiphonResources**](docs/FleetAPI.md#siphonresources) | **Post** /my/ships/{shipSymbol}/siphon | Siphon Resources
*FleetAPI* | [**TransferCargo**](docs/FleetAPI.md#transfercargo) | **Post** /my/ships/{shipSymbol}/transfer | Transfer Cargo
*FleetAPI* | [**WarpShip**](docs/FleetAPI.md#warpship) | **Post** /my/ships/{shipSymbol}/warp | Warp Ship
*SystemsAPI* | [**GetConstruction**](docs/SystemsAPI.md#getconstruction) | **Get** /systems/{systemSymbol}/waypoints/{waypointSymbol}/construction | Get Construction Site
*SystemsAPI* | [**GetJumpGate**](docs/SystemsAPI.md#getjumpgate) | **Get** /systems/{systemSymbol}/waypoints/{waypointSymbol}/jump-gate | Get Jump Gate
*SystemsAPI* | [**GetMarket**](docs/SystemsAPI.md#getmarket) | **Get** /systems/{systemSymbol}/waypoints/{waypointSymbol}/market | Get Market
*SystemsAPI* | [**GetShipyard**](docs/SystemsAPI.md#getshipyard) | **Get** /systems/{systemSymbol}/waypoints/{waypointSymbol}/shipyard | Get Shipyard
*SystemsAPI* | [**GetSystem**](docs/SystemsAPI.md#getsystem) | **Get** /systems/{systemSymbol} | Get System
*SystemsAPI* | [**GetSystemWaypoints**](docs/SystemsAPI.md#getsystemwaypoints) | **Get** /systems/{systemSymbol}/waypoints | List Waypoints in System
*SystemsAPI* | [**GetSystems**](docs/SystemsAPI.md#getsystems) | **Get** /systems | List Systems
*SystemsAPI* | [**GetWaypoint**](docs/SystemsAPI.md#getwaypoint) | **Get** /systems/{systemSymbol}/waypoints/{waypointSymbol} | Get Waypoint
*SystemsAPI* | [**SupplyConstruction**](docs/SystemsAPI.md#supplyconstruction) | **Post** /systems/{systemSymbol}/waypoints/{waypointSymbol}/construction/supply | Supply Construction Site


## Documentation For Models

 - [AcceptContract200Response](docs/AcceptContract200Response.md)
 - [AcceptContract200ResponseData](docs/AcceptContract200ResponseData.md)
 - [ActivityLevel](docs/ActivityLevel.md)
 - [Agent](docs/Agent.md)
 - [Chart](docs/Chart.md)
 - [Construction](docs/Construction.md)
 - [ConstructionMaterial](docs/ConstructionMaterial.md)
 - [Contract](docs/Contract.md)
 - [ContractDeliverGood](docs/ContractDeliverGood.md)
 - [ContractPayment](docs/ContractPayment.md)
 - [ContractTerms](docs/ContractTerms.md)
 - [Cooldown](docs/Cooldown.md)
 - [CreateChart201Response](docs/CreateChart201Response.md)
 - [CreateChart201ResponseData](docs/CreateChart201ResponseData.md)
 - [CreateShipShipScan201Response](docs/CreateShipShipScan201Response.md)
 - [CreateShipShipScan201ResponseData](docs/CreateShipShipScan201ResponseData.md)
 - [CreateShipSystemScan201Response](docs/CreateShipSystemScan201Response.md)
 - [CreateShipSystemScan201ResponseData](docs/CreateShipSystemScan201ResponseData.md)
 - [CreateShipWaypointScan201Response](docs/CreateShipWaypointScan201Response.md)
 - [CreateShipWaypointScan201ResponseData](docs/CreateShipWaypointScan201ResponseData.md)
 - [CreateSurvey201Response](docs/CreateSurvey201Response.md)
 - [CreateSurvey201ResponseData](docs/CreateSurvey201ResponseData.md)
 - [DeliverContract200Response](docs/DeliverContract200Response.md)
 - [DeliverContract200ResponseData](docs/DeliverContract200ResponseData.md)
 - [DeliverContractRequest](docs/DeliverContractRequest.md)
 - [DockShip200Response](docs/DockShip200Response.md)
 - [ExtractResources201Response](docs/ExtractResources201Response.md)
 - [ExtractResources201ResponseData](docs/ExtractResources201ResponseData.md)
 - [ExtractResourcesRequest](docs/ExtractResourcesRequest.md)
 - [Extraction](docs/Extraction.md)
 - [ExtractionYield](docs/ExtractionYield.md)
 - [Faction](docs/Faction.md)
 - [FactionSymbol](docs/FactionSymbol.md)
 - [FactionTrait](docs/FactionTrait.md)
 - [FactionTraitSymbol](docs/FactionTraitSymbol.md)
 - [FulfillContract200Response](docs/FulfillContract200Response.md)
 - [GetAgents200Response](docs/GetAgents200Response.md)
 - [GetConstruction200Response](docs/GetConstruction200Response.md)
 - [GetContract200Response](docs/GetContract200Response.md)
 - [GetContracts200Response](docs/GetContracts200Response.md)
 - [GetFaction200Response](docs/GetFaction200Response.md)
 - [GetFactions200Response](docs/GetFactions200Response.md)
 - [GetJumpGate200Response](docs/GetJumpGate200Response.md)
 - [GetMarket200Response](docs/GetMarket200Response.md)
 - [GetMounts200Response](docs/GetMounts200Response.md)
 - [GetMyAgent200Response](docs/GetMyAgent200Response.md)
 - [GetMyShip200Response](docs/GetMyShip200Response.md)
 - [GetMyShipCargo200Response](docs/GetMyShipCargo200Response.md)
 - [GetMyShips200Response](docs/GetMyShips200Response.md)
 - [GetShipCooldown200Response](docs/GetShipCooldown200Response.md)
 - [GetShipNav200Response](docs/GetShipNav200Response.md)
 - [GetShipyard200Response](docs/GetShipyard200Response.md)
 - [GetStatus200Response](docs/GetStatus200Response.md)
 - [GetStatus200ResponseAnnouncementsInner](docs/GetStatus200ResponseAnnouncementsInner.md)
 - [GetStatus200ResponseLeaderboards](docs/GetStatus200ResponseLeaderboards.md)
 - [GetStatus200ResponseLeaderboardsMostCreditsInner](docs/GetStatus200ResponseLeaderboardsMostCreditsInner.md)
 - [GetStatus200ResponseLeaderboardsMostSubmittedChartsInner](docs/GetStatus200ResponseLeaderboardsMostSubmittedChartsInner.md)
 - [GetStatus200ResponseLinksInner](docs/GetStatus200ResponseLinksInner.md)
 - [GetStatus200ResponseServerResets](docs/GetStatus200ResponseServerResets.md)
 - [GetStatus200ResponseStats](docs/GetStatus200ResponseStats.md)
 - [GetSystem200Response](docs/GetSystem200Response.md)
 - [GetSystemWaypoints200Response](docs/GetSystemWaypoints200Response.md)
 - [GetSystemWaypointsTraitsParameter](docs/GetSystemWaypointsTraitsParameter.md)
 - [GetSystems200Response](docs/GetSystems200Response.md)
 - [GetWaypoint200Response](docs/GetWaypoint200Response.md)
 - [InstallMount201Response](docs/InstallMount201Response.md)
 - [InstallMount201ResponseData](docs/InstallMount201ResponseData.md)
 - [InstallMountRequest](docs/InstallMountRequest.md)
 - [Jettison200Response](docs/Jettison200Response.md)
 - [Jettison200ResponseData](docs/Jettison200ResponseData.md)
 - [JettisonRequest](docs/JettisonRequest.md)
 - [JumpGate](docs/JumpGate.md)
 - [JumpShip200Response](docs/JumpShip200Response.md)
 - [JumpShip200ResponseData](docs/JumpShip200ResponseData.md)
 - [JumpShipRequest](docs/JumpShipRequest.md)
 - [Market](docs/Market.md)
 - [MarketTradeGood](docs/MarketTradeGood.md)
 - [MarketTransaction](docs/MarketTransaction.md)
 - [Meta](docs/Meta.md)
 - [NavigateShip200Response](docs/NavigateShip200Response.md)
 - [NavigateShip200ResponseData](docs/NavigateShip200ResponseData.md)
 - [NavigateShipRequest](docs/NavigateShipRequest.md)
 - [NegotiateContract200Response](docs/NegotiateContract200Response.md)
 - [NegotiateContract200ResponseData](docs/NegotiateContract200ResponseData.md)
 - [OrbitShip200Response](docs/OrbitShip200Response.md)
 - [OrbitShip200ResponseData](docs/OrbitShip200ResponseData.md)
 - [PatchShipNavRequest](docs/PatchShipNavRequest.md)
 - [PurchaseCargo201Response](docs/PurchaseCargo201Response.md)
 - [PurchaseCargoRequest](docs/PurchaseCargoRequest.md)
 - [PurchaseShip201Response](docs/PurchaseShip201Response.md)
 - [PurchaseShip201ResponseData](docs/PurchaseShip201ResponseData.md)
 - [PurchaseShipRequest](docs/PurchaseShipRequest.md)
 - [RefuelShip200Response](docs/RefuelShip200Response.md)
 - [RefuelShip200ResponseData](docs/RefuelShip200ResponseData.md)
 - [RefuelShipRequest](docs/RefuelShipRequest.md)
 - [Register201Response](docs/Register201Response.md)
 - [Register201ResponseData](docs/Register201ResponseData.md)
 - [RegisterRequest](docs/RegisterRequest.md)
 - [RemoveMount201Response](docs/RemoveMount201Response.md)
 - [RemoveMount201ResponseData](docs/RemoveMount201ResponseData.md)
 - [RemoveMountRequest](docs/RemoveMountRequest.md)
 - [ScannedShip](docs/ScannedShip.md)
 - [ScannedShipEngine](docs/ScannedShipEngine.md)
 - [ScannedShipFrame](docs/ScannedShipFrame.md)
 - [ScannedShipMountsInner](docs/ScannedShipMountsInner.md)
 - [ScannedShipReactor](docs/ScannedShipReactor.md)
 - [ScannedSystem](docs/ScannedSystem.md)
 - [ScannedWaypoint](docs/ScannedWaypoint.md)
 - [SellCargo201Response](docs/SellCargo201Response.md)
 - [SellCargo201ResponseData](docs/SellCargo201ResponseData.md)
 - [SellCargoRequest](docs/SellCargoRequest.md)
 - [Ship](docs/Ship.md)
 - [ShipCargo](docs/ShipCargo.md)
 - [ShipCargoItem](docs/ShipCargoItem.md)
 - [ShipCrew](docs/ShipCrew.md)
 - [ShipEngine](docs/ShipEngine.md)
 - [ShipFrame](docs/ShipFrame.md)
 - [ShipFuel](docs/ShipFuel.md)
 - [ShipFuelConsumed](docs/ShipFuelConsumed.md)
 - [ShipModificationTransaction](docs/ShipModificationTransaction.md)
 - [ShipModule](docs/ShipModule.md)
 - [ShipMount](docs/ShipMount.md)
 - [ShipNav](docs/ShipNav.md)
 - [ShipNavFlightMode](docs/ShipNavFlightMode.md)
 - [ShipNavRoute](docs/ShipNavRoute.md)
 - [ShipNavRouteWaypoint](docs/ShipNavRouteWaypoint.md)
 - [ShipNavStatus](docs/ShipNavStatus.md)
 - [ShipReactor](docs/ShipReactor.md)
 - [ShipRefine201Response](docs/ShipRefine201Response.md)
 - [ShipRefine201ResponseData](docs/ShipRefine201ResponseData.md)
 - [ShipRefine201ResponseDataProducedInner](docs/ShipRefine201ResponseDataProducedInner.md)
 - [ShipRefineRequest](docs/ShipRefineRequest.md)
 - [ShipRegistration](docs/ShipRegistration.md)
 - [ShipRequirements](docs/ShipRequirements.md)
 - [ShipRole](docs/ShipRole.md)
 - [ShipType](docs/ShipType.md)
 - [Shipyard](docs/Shipyard.md)
 - [ShipyardShip](docs/ShipyardShip.md)
 - [ShipyardShipCrew](docs/ShipyardShipCrew.md)
 - [ShipyardShipTypesInner](docs/ShipyardShipTypesInner.md)
 - [ShipyardTransaction](docs/ShipyardTransaction.md)
 - [Siphon](docs/Siphon.md)
 - [SiphonResources201Response](docs/SiphonResources201Response.md)
 - [SiphonResources201ResponseData](docs/SiphonResources201ResponseData.md)
 - [SiphonYield](docs/SiphonYield.md)
 - [SupplyConstruction201Response](docs/SupplyConstruction201Response.md)
 - [SupplyConstruction201ResponseData](docs/SupplyConstruction201ResponseData.md)
 - [SupplyConstructionRequest](docs/SupplyConstructionRequest.md)
 - [SupplyLevel](docs/SupplyLevel.md)
 - [Survey](docs/Survey.md)
 - [SurveyDeposit](docs/SurveyDeposit.md)
 - [System](docs/System.md)
 - [SystemFaction](docs/SystemFaction.md)
 - [SystemType](docs/SystemType.md)
 - [SystemWaypoint](docs/SystemWaypoint.md)
 - [TradeGood](docs/TradeGood.md)
 - [TradeSymbol](docs/TradeSymbol.md)
 - [TransferCargo200Response](docs/TransferCargo200Response.md)
 - [TransferCargoRequest](docs/TransferCargoRequest.md)
 - [Waypoint](docs/Waypoint.md)
 - [WaypointFaction](docs/WaypointFaction.md)
 - [WaypointModifier](docs/WaypointModifier.md)
 - [WaypointModifierSymbol](docs/WaypointModifierSymbol.md)
 - [WaypointOrbital](docs/WaypointOrbital.md)
 - [WaypointTrait](docs/WaypointTrait.md)
 - [WaypointTraitSymbol](docs/WaypointTraitSymbol.md)
 - [WaypointType](docs/WaypointType.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### AgentToken

- **Type**: HTTP Bearer token authentication

Example

```golang
auth := context.WithValue(context.Background(), stsdk.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

joel@spacetraders.io

