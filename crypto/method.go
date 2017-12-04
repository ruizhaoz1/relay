/*

  Copyright 2017 Loopring Project Ltd (Loopring Foundation).

  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.

*/

package crypto

import (
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"fmt"
)

func ValidateSignatureValues(v byte, r, s []byte) bool {
	return crypto.ValidateSignatureValues(v, r, s)
}

func GenerateHash(data ...[]byte) []byte {
	fmt.Print(crypto)
	return crypto.GenerateHash(data...)
}

func Sign(hash []byte, signerAddr common.Address) ([]byte, error) {
	signer := accounts.Account{Address: signerAddr}
	return crypto.Sign(hash, signer)
}

func SigToAddress(hash, sig []byte) ([]byte, error) {
	return crypto.SigToAddress(hash, sig)
}

func VRSToSig(v byte, r, s []byte) ([]byte, error) {
	return crypto.VRSToSig(v, r, s)
}

func SigToVRS(sig []byte) (v byte, r []byte, s []byte) {
	return crypto.SigToVRS(sig)
}

func UnlockAccount(acc accounts.Account, passphrase string) error {
	return crypto.UnlockAccount(acc, passphrase)
}

func SignTx(a accounts.Account, tx *types.Transaction, chainID *big.Int) (*types.Transaction, error) {
	return crypto.SignTx(a, tx, chainID)
}
func Initialize(c Crypto) {
	fmt.Print("cccccc")
	fmt.Print(c)
	crypto = c
}
