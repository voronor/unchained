/*
Tendermint RPC

Tendermint supports the following RPC protocols:  * URI over HTTP * JSONRPC over HTTP * JSONRPC over websockets  ## Configuration  RPC can be configured by tuning parameters under `[rpc]` table in the `$TMHOME/config/config.toml` file or by using the `--rpc.X` command-line flags.  Default rpc listen address is `tcp://0.0.0.0:26657`. To set another address, set the `laddr` config parameter to desired value. CORS (Cross-Origin Resource Sharing) can be enabled by setting `cors_allowed_origins`, `cors_allowed_methods`, `cors_allowed_headers` config parameters.  ## Arguments  Arguments which expect strings or byte arrays may be passed as quoted strings, like `\"abc\"` or as `0x`-prefixed strings, like `0x616263`.  ## URI/HTTP  A REST like interface.      curl localhost:26657/block?height=5  ## JSONRPC/HTTP  JSONRPC requests can be POST'd to the root RPC endpoint via HTTP.      curl --header \"Content-Type: application/json\" --request POST --data '{\"method\": \"block\", \"params\": [\"5\"], \"id\": 1}' localhost:26657  ## JSONRPC/websockets  JSONRPC requests can be also made via websocket. The websocket endpoint is at `/websocket`, e.g. `localhost:26657/websocket`. Asynchronous RPC functions like event `subscribe` and `unsubscribe` are only available via websockets.  Example using https://github.com/hashrocket/ws:      ws ws://localhost:26657/websocket     > { \"jsonrpc\": \"2.0\", \"method\": \"subscribe\", \"params\": [\"tm.event='NewBlock'\"], \"id\": 1 } 

API version: Master
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// DumpConsensusResponseResult struct for DumpConsensusResponseResult
type DumpConsensusResponseResult struct {
	RoundState DumpConsensusResponseResultRoundState `json:"round_state"`
	Peers []DumpConsensusResponseResultPeers `json:"peers"`
}

// NewDumpConsensusResponseResult instantiates a new DumpConsensusResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDumpConsensusResponseResult(roundState DumpConsensusResponseResultRoundState, peers []DumpConsensusResponseResultPeers) *DumpConsensusResponseResult {
	this := DumpConsensusResponseResult{}
	this.RoundState = roundState
	this.Peers = peers
	return &this
}

// NewDumpConsensusResponseResultWithDefaults instantiates a new DumpConsensusResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDumpConsensusResponseResultWithDefaults() *DumpConsensusResponseResult {
	this := DumpConsensusResponseResult{}
	return &this
}

// GetRoundState returns the RoundState field value
func (o *DumpConsensusResponseResult) GetRoundState() DumpConsensusResponseResultRoundState {
	if o == nil {
		var ret DumpConsensusResponseResultRoundState
		return ret
	}

	return o.RoundState
}

// GetRoundStateOk returns a tuple with the RoundState field value
// and a boolean to check if the value has been set.
func (o *DumpConsensusResponseResult) GetRoundStateOk() (*DumpConsensusResponseResultRoundState, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RoundState, true
}

// SetRoundState sets field value
func (o *DumpConsensusResponseResult) SetRoundState(v DumpConsensusResponseResultRoundState) {
	o.RoundState = v
}

// GetPeers returns the Peers field value
func (o *DumpConsensusResponseResult) GetPeers() []DumpConsensusResponseResultPeers {
	if o == nil {
		var ret []DumpConsensusResponseResultPeers
		return ret
	}

	return o.Peers
}

// GetPeersOk returns a tuple with the Peers field value
// and a boolean to check if the value has been set.
func (o *DumpConsensusResponseResult) GetPeersOk() (*[]DumpConsensusResponseResultPeers, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Peers, true
}

// SetPeers sets field value
func (o *DumpConsensusResponseResult) SetPeers(v []DumpConsensusResponseResultPeers) {
	o.Peers = v
}

func (o DumpConsensusResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["round_state"] = o.RoundState
	}
	if true {
		toSerialize["peers"] = o.Peers
	}
	return json.Marshal(toSerialize)
}

type NullableDumpConsensusResponseResult struct {
	value *DumpConsensusResponseResult
	isSet bool
}

func (v NullableDumpConsensusResponseResult) Get() *DumpConsensusResponseResult {
	return v.value
}

func (v *NullableDumpConsensusResponseResult) Set(val *DumpConsensusResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableDumpConsensusResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableDumpConsensusResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDumpConsensusResponseResult(val *DumpConsensusResponseResult) *NullableDumpConsensusResponseResult {
	return &NullableDumpConsensusResponseResult{value: val, isSet: true}
}

func (v NullableDumpConsensusResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDumpConsensusResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

