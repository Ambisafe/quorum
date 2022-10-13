package istanbulcommon

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
)

// GetSignatureAddress gets the signer address from the signature
func GetSignatureAddress(data []byte, sig []byte) (common.Address, error) {
	// 1. Keccak data
	hashData := crypto.Keccak256(data)
	// 2. Recover public key
	pubkey, err := crypto.SigToPub(hashData, sig)
	if err != nil {
		log.Error("istanbulcommon.Utils GetSignatureAddress crypto.SigToPub(hashData, sig)", "err", err, "data", hexutil.Encode(data), "sig", hexutil.Encode(sig))
		return common.Address{}, err
	}
	return crypto.PubkeyToAddress(*pubkey), nil
}
