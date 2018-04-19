package GoWaves

import (
	"errors"
	"github.com/btcsuite/btcutil/base58"
)

type account struct {
	publicKey  []byte // public key base58 decoded
	privateKey []byte // private key base58 decoded
	address    []byte // address base58 decoded
	seed       string // seed string
}

type accountAlias struct {
	Alias []string // alias name address
}

type balance struct {
	Confirmations float32 `json:"confirmations" xml:"confirmations"`
	Balance       float32 `json:"balance" xml:"balance"`
}

/*func (a *account) Transfer(to *account, amount int, assetId string, fee int, feeAssetId string, attach string) {
	transaction := url.Values{
		"senderPublicKey": {base58.Encode(a.address)},
		"assetId":         {assetId},
		"recipient":       {base58.Encode(to.address)},
		"amount":          {string(amount)},
		"fee":             {string(fee)},
		"feeAssetId":      {string(feeAssetId)},
		"timestamp":       {"china"},
		"attachment":      {base58.Encode([]byte(attach))},
	}
	//nodeAPI.post("/assets/broadcast/transfer", transaction)
}*/

// Return array aliases
func (a *account) GetAlias() []string {
	ac := accountAlias{}
	nodeAPI.get("/alias/by-address/"+string(a.address), &ac)
	return ac.Alias
}

// Return account balance struct
func (a *account) GetBalance() balance {
	b := balance{}
	nodeAPI.get("/addresses/balance", &b)
	return b
}

// Return base58 string address
func (a *account) GetAddress() string {
	return base58.Encode(a.address)
}

// Return base58 string public key
func (a *account) GetPublicKey() (string, error) {
	if len(a.publicKey) > 0 {
		return base58.Encode(a.publicKey), nil
	}

	return "", errors.New("can't return public key, check method that creates the account")
}

// Return base58 string private key
func (a *account) GetPrivateKey() (string, error) {
	if len(a.privateKey) > 0 {
		return base58.Encode(a.privateKey), nil
	}
	return "", errors.New("can't return private key, check method that creates the account")
}

// Return seed string of account
func (a *account) GetSeed() (string, error) {
	if len(a.seed) > 0 {
		return a.seed, nil
	}
	return "", errors.New("can't return seed, check method that creates the account")
}
