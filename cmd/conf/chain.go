// Copyright 2021-2022, Offchain Labs, Inc.
// For license information, see https://github.com/nitro/blob/master/LICENSE

package conf

import (
	"github.com/offchainlabs/nitro/cmd/genericconf"
	flag "github.com/spf13/pflag"
)

type L1Config struct {
	ChainID            uint64                   `koanf:"chain-id"`
	URL                string                   `koanf:"url"`
	ConnectionAttempts int                      `koanf:"connection-attempts"`
	Wallet             genericconf.WalletConfig `koanf:"wallet"`
}

var L1ConfigDefault = L1Config{
	ChainID:            0,
	URL:                "",
	ConnectionAttempts: 15,
	Wallet:             genericconf.WalletConfigDefault,
}

func L1ConfigAddOptions(prefix string, f *flag.FlagSet) {
	f.Uint64(prefix+".chain-id", L1ConfigDefault.ChainID, "if set other than 0, will be used to validate database and L1 connection")
	f.String(prefix+".url", L1ConfigDefault.URL, "layer 1 ethereum node RPC URL")
	f.Int(prefix+".connection-attempts", L1ConfigDefault.ConnectionAttempts, "layer 1 RPC connection attempts (spaced out at least 1 second per attempt, 0 to retry infinitely)")
	genericconf.WalletConfigAddOptions(prefix+".wallet", f, "wallet")
}

func (c *L1Config) ResolveDirectoryNames(chain string) {
	c.Wallet.ResolveDirectoryNames(chain)
}

type L2Config struct {
	ChainID        uint64                   `koanf:"chain-id"`
	ChainInfoFiles []string                 `koanf:"chain-info-files"`
	DevWallet      genericconf.WalletConfig `koanf:"dev-wallet"`
}

var L2ConfigDefault = L2Config{
	ChainID:        0,
	ChainInfoFiles: []string{}, // Default file used is chaininfo/arbitrum_chain_info.json, stored in DefaultChainInfo in chain_info.go
	DevWallet:      genericconf.WalletConfigDefault,
}

func L2ConfigAddOptions(prefix string, f *flag.FlagSet) {
	f.Uint64(prefix+".chain-id", L2ConfigDefault.ChainID, "L2 chain ID (determines Arbitrum network)")
	f.StringSlice(prefix+".chain-info-files", L2ConfigDefault.ChainInfoFiles, "L2 chain info json files")

	// Dev wallet does not exist unless specified
	genericconf.WalletConfigAddOptions(prefix+".dev-wallet", f, "")
}

func (c *L2Config) ResolveDirectoryNames(chain string) {
	c.DevWallet.ResolveDirectoryNames(chain)
}
