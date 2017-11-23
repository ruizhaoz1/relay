// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package types

import (
	"encoding/json"
	"errors"
	"math/big"
)

var _ = (*orderMarshaling)(nil)

func (o Order) MarshalJSON() ([]byte, error) {
	type Order struct {
		Protocol              Address  `json:"protocol" gencodec:"required"`
		TokenS                Address  `json:"tokenS" gencodec:"required"`
		TokenB                Address  `json:"tokenB" gencodec:"required"`
		AmountS               *Big     `json:"amountS" gencodec:"required"`
		AmountB               *Big     `json:"amountB" gencodec:"required"`
		Timestamp             *Big     `json:"timestamp" gencodec:"required"`
		Ttl                   *Big     `json:"ttl" gencodec:"required"`
		Salt                  *Big     `json:"salt" gencodec:"required"`
		LrcFee                *Big     `json:"lrcFee" `
		BuyNoMoreThanAmountB  bool     `json:"buyNoMoreThanAmountB" gencodec:"required"`
		MarginSplitPercentage uint8    `json:"marginSplitPercentage" gencodec:"required"`
		V                     uint8    `json:"v" gencodec:"required"`
		R                     Sign     `json:"r" gencodec:"required"`
		S                     Sign     `json:"s" gencodec:"required"`
		Price                 *big.Rat `json:"price"`
		Owner                 Address  `json:"owner"`
		Hash                  Hash     `json:"hash"`
	}
	var enc Order
	enc.Protocol = o.Protocol
	enc.TokenS = o.TokenS
	enc.TokenB = o.TokenB
	enc.AmountS = (*Big)(o.AmountS)
	enc.AmountB = (*Big)(o.AmountB)
	enc.Timestamp = (*Big)(o.Timestamp)
	enc.Ttl = (*Big)(o.Ttl)
	enc.Salt = (*Big)(o.Salt)
	enc.LrcFee = (*Big)(o.LrcFee)
	enc.BuyNoMoreThanAmountB = o.BuyNoMoreThanAmountB
	enc.MarginSplitPercentage = o.MarginSplitPercentage
	enc.V = o.V
	enc.R = o.R
	enc.S = o.S
	enc.Price = o.Price
	enc.Owner = o.Owner
	enc.Hash = o.Hash
	return json.Marshal(&enc)
}

func (o *Order) UnmarshalJSON(input []byte) error {
	type Order struct {
		Protocol              *Address `json:"protocol" gencodec:"required"`
		TokenS                *Address `json:"tokenS" gencodec:"required"`
		TokenB                *Address `json:"tokenB" gencodec:"required"`
		AmountS               *Big     `json:"amountS" gencodec:"required"`
		AmountB               *Big     `json:"amountB" gencodec:"required"`
		Timestamp             *Big     `json:"timestamp" gencodec:"required"`
		Ttl                   *Big     `json:"ttl" gencodec:"required"`
		Salt                  *Big     `json:"salt" gencodec:"required"`
		LrcFee                *Big     `json:"lrcFee" `
		BuyNoMoreThanAmountB  *bool    `json:"buyNoMoreThanAmountB" gencodec:"required"`
		MarginSplitPercentage *uint8   `json:"marginSplitPercentage" gencodec:"required"`
		V                     *uint8   `json:"v" gencodec:"required"`
		R                     *Sign    `json:"r" gencodec:"required"`
		S                     *Sign    `json:"s" gencodec:"required"`
		Price                 *big.Rat `json:"price"`
		Owner                 *Address `json:"owner"`
		Hash                  *Hash    `json:"hash"`
	}
	var dec Order
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Protocol == nil {
		return errors.New("missing required field 'protocol' for Order")
	}
	o.Protocol = *dec.Protocol
	if dec.TokenS == nil {
		return errors.New("missing required field 'tokenS' for Order")
	}
	o.TokenS = *dec.TokenS
	if dec.TokenB == nil {
		return errors.New("missing required field 'tokenB' for Order")
	}
	o.TokenB = *dec.TokenB
	if dec.AmountS == nil {
		return errors.New("missing required field 'amountS' for Order")
	}
	o.AmountS = (*big.Int)(dec.AmountS)
	if dec.AmountB == nil {
		return errors.New("missing required field 'amountB' for Order")
	}
	o.AmountB = (*big.Int)(dec.AmountB)
	if dec.Timestamp == nil {
		return errors.New("missing required field 'timestamp' for Order")
	}
	o.Timestamp = (*big.Int)(dec.Timestamp)
	if dec.Ttl == nil {
		return errors.New("missing required field 'ttl' for Order")
	}
	o.Ttl = (*big.Int)(dec.Ttl)
	if dec.Salt == nil {
		return errors.New("missing required field 'salt' for Order")
	}
	o.Salt = (*big.Int)(dec.Salt)
	if dec.LrcFee != nil {
		o.LrcFee = (*big.Int)(dec.LrcFee)
	}
	if dec.BuyNoMoreThanAmountB == nil {
		return errors.New("missing required field 'buyNoMoreThanAmountB' for Order")
	}
	o.BuyNoMoreThanAmountB = *dec.BuyNoMoreThanAmountB
	if dec.MarginSplitPercentage == nil {
		return errors.New("missing required field 'marginSplitPercentage' for Order")
	}
	o.MarginSplitPercentage = *dec.MarginSplitPercentage
	if dec.V == nil {
		return errors.New("missing required field 'v' for Order")
	}
	o.V = *dec.V
	if dec.R == nil {
		return errors.New("missing required field 'r' for Order")
	}
	o.R = *dec.R
	if dec.S == nil {
		return errors.New("missing required field 's' for Order")
	}
	o.S = *dec.S
	if dec.Price != nil {
		o.Price = dec.Price
	}
	if dec.Owner != nil {
		o.Owner = *dec.Owner
	}
	if dec.Hash != nil {
		o.Hash = *dec.Hash
	}
	return nil
}
