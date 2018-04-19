// This is package for work Waves platform
package GoWaves

import (
	"github.com/btcsuite/btcutil/base58"
)

// node struct
type node struct {
	addr    string
	chainID string
}

// for private parse json response of node version
type nodeVersion struct {
	Version string `json:"version" xml:"version"`
}

// for private parse json response of last blocks
type nodeHeight struct {
	Height int `json:"height" xml:"height"`
}

// chainids
const (
	MAINNET = "W" // chainId 'W' for mainnet
	TESTNET = "T" // chainId 'T' for testnet
)

// Return Node
func Node(addr string, chainID string) *node {
	nodeAPI = api{
		addr: addr,
	}
	return &node{
		addr:    addr,
		chainID: chainID,
	}
}

// Return account
func (n *node) Account() account {
	seed := generateSeed()
	privateKey := generatePrivateKey(seed, 0)
	publicKey := generatePublicKey(privateKey)
	address := generateAddress(publicKey, n.chainID)
	return account{
		publicKey:  publicKey,
		privateKey: privateKey[:],
		address:    address,
		seed:       seed,
	}
}

// Return account by base58 address
func (n *node) AccountByAddress(address58 string) account {
	return account{
		address: base58.Decode(address58),
	}
}

// Return account by base58 public key
func (n *node) AccountByPublicKey(publicKey58 string) account {
	publicKey := base58.Decode(publicKey58)
	address := generateAddress(publicKey, n.chainID)
	return account{
		publicKey: publicKey,
		address:   address,
	}
}

// Return account by base58 private key
func (n *node) AccountByPrivateKey(privateKey58 string) account {
	privateKey := base58.Decode(privateKey58)

	var b [32]byte
	copy(b[:], privateKey)

	publicKey := generatePublicKey(b)
	address := generateAddress(publicKey, n.chainID)

	return account{
		publicKey:  publicKey,
		privateKey: privateKey,
		address:    address,
	}
}

// Return account by seed
func (n *node) AccountBySeed(seed string) account {
	privateKey := generatePrivateKey(seed, 0)
	publicKey := generatePublicKey(privateKey)
	address := generateAddress(publicKey, n.chainID)
	return account{
		publicKey:  publicKey,
		privateKey: privateKey[:],
		address:    address,
		seed:       seed,
	}
}

// Return account by alias
func (n *node) AccountByAlias(alias string) account {
	return account{}
}

// Return node version
func (n *node) GetVersion() string {
	node := nodeVersion{}
	nodeAPI.get("/node/version", &node)
	return node.Version
}

// Return node version
func (n *node) GetHeight() int {
	node := nodeHeight{}
	nodeAPI.get("/blocks/height", &node)
	return node.Height
}
