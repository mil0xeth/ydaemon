package env

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/yearn/ydaemon/internal/models"
)

type TExtraStakingContracts struct {
	VaultAddress   common.Address
	StakingAddress common.Address
	Tag            string
}
type TChainCurve struct {
	RegistryAddress common.Address
	FactoryAddress  common.Address
	PoolsURIs       []string
	GaugesURI       string
}
type TChainExtraURI struct {
	GammaMerklURI      string
	GammaHypervisorURI []string
	PendleCoreURI      string
}
type TChain struct {
	ID                    uint64
	RpcURI                string
	SubgraphURI           string
	EtherscanURI          string
	MaxBlockRange         uint64
	MaxBatchSize          uint64
	AvgBlocksPerDay       int
	CanUseWebsocket       bool
	LensContract          TContractData
	MulticallContract     TContractData
	YBribeV3Contract      TContractData
	PartnerContract       TContractData
	APROracleContract     TContractData
	Coin                  models.TERC20Token
	StakingRewardRegistry []TContractData
	Registries            []TContractData
	ExtraStakingContracts []TExtraStakingContracts
	ExtraVaults           []models.TVaultsFromRegistry
	BlacklistedVaults     []common.Address
	ExtraTokens           []common.Address
	IgnoredTokens         []common.Address
	Curve                 TChainCurve
	ExtraURI              TChainExtraURI
}

var DEFAULT_COIN_ADDRESS = common.HexToAddress(`0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE`)

// CHAINS is the list of supported chains
var CHAINS = map[uint64]TChain{
	1:     ETHEREUM,
	10:    OPTIMISM,
	100:   GNOSIS,
	137:   POLYGON,
	250:   FANTOM,
	8453:  BASE,
	42161: ARBITRUM,
}

var SUPPORTED_CHAIN_IDS = []uint64{}

func init() {
	for k := range CHAINS {
		SUPPORTED_CHAIN_IDS = append(SUPPORTED_CHAIN_IDS, k)
	}
}
