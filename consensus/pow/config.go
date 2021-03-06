package pow

import (
	"tinychain/common"
)

type Config struct {
	Miner     bool
	PrivKey    []byte
	GasLimit   uint64
	Difficulty uint64
}

func newConfig(config *common.Config) *Config {

	return &Config{
		Miner:     config.GetBool("consensus.mining"),
		PrivKey:    []byte(config.GetString("consensus.private_key")),
		GasLimit:   uint64(config.GetInt64("consensus.gas_limit")),
		Difficulty: uint64(config.GetInt64("consensus.difficulty")),
	}
}
