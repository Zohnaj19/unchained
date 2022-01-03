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

// TxResponseResult struct for TxResponseResult
type TxResponseResult struct {
	Hash string `json:"hash"`
	Height string `json:"height"`
	Index int32 `json:"index"`
	TxResult TxResponseResultTxResult `json:"tx_result"`
	Tx string `json:"tx"`
}

// NewTxResponseResult instantiates a new TxResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTxResponseResult(hash string, height string, index int32, txResult TxResponseResultTxResult, tx string) *TxResponseResult {
	this := TxResponseResult{}
	this.Hash = hash
	this.Height = height
	this.Index = index
	this.TxResult = txResult
	this.Tx = tx
	return &this
}

// NewTxResponseResultWithDefaults instantiates a new TxResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTxResponseResultWithDefaults() *TxResponseResult {
	this := TxResponseResult{}
	return &this
}

// GetHash returns the Hash field value
func (o *TxResponseResult) GetHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hash
}

// GetHashOk returns a tuple with the Hash field value
// and a boolean to check if the value has been set.
func (o *TxResponseResult) GetHashOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Hash, true
}

// SetHash sets field value
func (o *TxResponseResult) SetHash(v string) {
	o.Hash = v
}

// GetHeight returns the Height field value
func (o *TxResponseResult) GetHeight() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Height
}

// GetHeightOk returns a tuple with the Height field value
// and a boolean to check if the value has been set.
func (o *TxResponseResult) GetHeightOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Height, true
}

// SetHeight sets field value
func (o *TxResponseResult) SetHeight(v string) {
	o.Height = v
}

// GetIndex returns the Index field value
func (o *TxResponseResult) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *TxResponseResult) GetIndexOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *TxResponseResult) SetIndex(v int32) {
	o.Index = v
}

// GetTxResult returns the TxResult field value
func (o *TxResponseResult) GetTxResult() TxResponseResultTxResult {
	if o == nil {
		var ret TxResponseResultTxResult
		return ret
	}

	return o.TxResult
}

// GetTxResultOk returns a tuple with the TxResult field value
// and a boolean to check if the value has been set.
func (o *TxResponseResult) GetTxResultOk() (*TxResponseResultTxResult, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TxResult, true
}

// SetTxResult sets field value
func (o *TxResponseResult) SetTxResult(v TxResponseResultTxResult) {
	o.TxResult = v
}

// GetTx returns the Tx field value
func (o *TxResponseResult) GetTx() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tx
}

// GetTxOk returns a tuple with the Tx field value
// and a boolean to check if the value has been set.
func (o *TxResponseResult) GetTxOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Tx, true
}

// SetTx sets field value
func (o *TxResponseResult) SetTx(v string) {
	o.Tx = v
}

func (o TxResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["hash"] = o.Hash
	}
	if true {
		toSerialize["height"] = o.Height
	}
	if true {
		toSerialize["index"] = o.Index
	}
	if true {
		toSerialize["tx_result"] = o.TxResult
	}
	if true {
		toSerialize["tx"] = o.Tx
	}
	return json.Marshal(toSerialize)
}

type NullableTxResponseResult struct {
	value *TxResponseResult
	isSet bool
}

func (v NullableTxResponseResult) Get() *TxResponseResult {
	return v.value
}

func (v *NullableTxResponseResult) Set(val *TxResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableTxResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableTxResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTxResponseResult(val *TxResponseResult) *NullableTxResponseResult {
	return &NullableTxResponseResult{value: val, isSet: true}
}

func (v NullableTxResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTxResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


