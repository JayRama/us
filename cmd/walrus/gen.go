package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/pkg/errors"
	"github.com/JayRama/us/cmd/walrus/api"
	"github.com/JayRama/us/wallet"
)

func gen(seed wallet.Seed, indexStr string) error {
	index, err := strconv.ParseUint(indexStr, 10, 64)
	if err != nil {
		return errors.Wrap(err, "invalid key index")
	}
	info := wallet.SeedAddressInfo{
		UnlockConditions: wallet.StandardUnlockConditions(seed.PublicKey(index)),
		KeyIndex:         index,
	}
	js, _ := json.MarshalIndent(api.ResponseAddressesAddr(info), "", "\t")
	fmt.Println(info.UnlockConditions.UnlockHash())
	fmt.Println(string(js))
	return nil
}
