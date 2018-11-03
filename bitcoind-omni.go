package bitcoind

import (
	"encoding/json"
	"strings"
)

type OmniBalance struct {
	Balance string `json:"balance"`
	Reserved string `json:"reserved"`
	Frozen string `json:"frozen"`
}

// OmniGetTrade return transaction info for given transaction id.
func (b *Bitcoind) OmniGetTrade(txId string) (rawTx interface{}, err error){
	r, err := b.client.call("omni_gettrade", []interface{}{txId})
	if err = handleError(err, &r); err != nil {
		return
	}
	var t OmniTransaction
	err = json.Unmarshal(r.Result, &t)
	rawTx = t
	return
}

// OmniSend return transaction hash for send
func (b *Bitcoind) OmniSend(fromAddress, toAddress string, propertyId int, amount, redeemaddress, referenceamount string) (txID string, err error){
	params := []interface{}{fromAddress, toAddress, propertyId, amount}
	if strings.Compare(redeemaddress, "") != 0 {
		params = append(params, redeemaddress)
		if strings.Compare(referenceamount, "") != 0 {
			params = append(params, referenceamount)
		}
	}
	r, err := b.client.call("omni_send", params)
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &txID)
	return
}

// Creates and sends a funded simple send transaction
func (b *Bitcoind) OmniFundedSend(fromAddress, toAddress string, propertyId int, amount, feeaddress string)  (txID string, err error){
	r, err := b.client.call("omni_funded_send", []interface{}{fromAddress, toAddress, propertyId, amount, feeaddress})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &txID)
	return
}

func (b *Bitcoind) OmniGetBalance(address string, propertyId int)  (balance *OmniBalance, err error) {
	r, err := b.client.call("omni_getbalance", []interface{}{address, propertyId})
	if err = handleError(err, &r); err != nil {
		return nil, err
	}
	err = json.Unmarshal(r.Result, &balance)
	return
}

func (b *Bitcoind) OmniListBlockTransactions(index int64) (txId []string, err error) {
	r, err := b.client.call("omni_listblocktransactions", []interface{}{index})
	if err = handleError(err, &r); err != nil {
		return
	}
	err = json.Unmarshal(r.Result, &txId)
	return
}
