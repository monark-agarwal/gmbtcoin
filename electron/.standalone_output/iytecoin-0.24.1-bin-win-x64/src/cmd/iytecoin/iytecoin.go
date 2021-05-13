package main

import (
	_ "net/http/pprof"

	"github.com/skycoin/skycoin/src/skycoin"
	"github.com/skycoin/skycoin/src/util/logging"
	"github.com/skycoin/skycoin/src/visor"
)

var (
	// Version of the node. Can be set by -ldflags
	Version = "0.24.1"
	// Commit ID. Can be set by -ldflags
	Commit = ""
	// Branch name. Can be set by -ldflags
	Branch = ""
	// ConfigMode (possible values are "", "STANDALONE_CLIENT").
	// This is used to change the default configuration.
	// Can be set by -ldflags
	ConfigMode = ""

	logger = logging.MustGetLogger("main")

		GenesisSignatureStr = "b14e67510f69af65ce51c7dd8e53dd41ac340236d7449f0f7bbcb93e04c69aa16e11c87c59a530da2681243991c63843d62bf2c1d5bb07d437cb9f612ef80c3e00"
// GenesisAddressStr genesis address string
GenesisAddressStr = "Vd9yz17NDwX3fHxFVYAGYNYD8gnNXgPEtb"
// BlockchainPubkeyStr pubic key string
BlockchainPubkeyStr = "0278cf11fafc5b0950dc9e85aa0bb504018ebcb0a2adc061b111a65b0699cd3150"
// BlockchainSeckeyStr empty private key string
BlockchainSeckeyStr = ""
// GenesisTimestamp genesis block create unix time
GenesisTimestamp uint64 = 1614525836
// GenesisCoinVolume represents the coin capacity
GenesisCoinVolume uint64 = 100e12
// DefaultConnections the default trust node addresses
DefaultConnections = []string{
"143.110.237.134:40000",
            "143.110.237.134:40001",
            "143.110.237.134:40002",
			"138.68.82.4:40001",
			"138.68.82.4:40002",

	}
)
func main() {
	// get node config
	nodeConfig := skycoin.NewNodeConfig(ConfigMode, skycoin.NodeParameters{
		GenesisSignatureStr: GenesisSignatureStr,
		GenesisAddressStr:   GenesisAddressStr,
		GenesisCoinVolume:   GenesisCoinVolume,
		GenesisTimestamp:    GenesisTimestamp,
		BlockchainPubkeyStr: BlockchainPubkeyStr,
		BlockchainSeckeyStr: BlockchainSeckeyStr,
		DefaultConnections:  DefaultConnections,
		PeerListURL:         "https://iytecoin.org/peers.txt",
		Port:                40000,
		WebInterfacePort:    4220,
		DataDirectory:       "$HOME/.iytecoin",
		ProfileCPUFile:      "skycoin.prof",
	})

	// create a new fiber coin instance
	coin := skycoin.NewCoin(
		skycoin.Config{
			Node: *nodeConfig,
			Build: visor.BuildInfo{
				Version: Version,
				Commit:  Commit,
				Branch:  Branch,
			},
		},
		logger,
	)

	// parse config values
	coin.ParseConfig()

	// run fiber coin node
	coin.Run()
}
