// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package genesis

import (
	"time"

	"github.com/flare-foundation/flare/utils/units"
)

var (
	localGenesisConfigJSON = `{
		"networkID": 12345,
		"allocations": [],
		"startTime": 1630987200,
		"initialStakeDuration": 31536000,
		"initialStakeDurationOffset": 5400,
		"initialStakedFunds": [],
		"initialStakers": [],
		"cChainGenesis": "",
		"message": "purgatory"
	}`

	// localCChainGenesis is the C-Chain genesis block used for the local
	// network.
	localCChainGenesis = `{
		"config": {
			"chainId": 20210406,
			"homesteadBlock": 0,
			"daoForkBlock": 0,
			"daoForkSupport": true,
			"eip150Block": 0,
			"eip150Hash": "0x2086799aeebeae135c246c65021c82b4e15a2c451340993aacfd2751886514f0",
			"eip155Block": 0,
			"eip158Block": 0,
			"byzantiumBlock": 0,
			"constantinopleBlock": 0,
			"petersburgBlock": 0,
			"istanbulBlock": 0,
			"muirGlacierBlock": 0,
			"apricotPhase1BlockTimestamp": 0,
			"apricotPhase2BlockTimestamp": 0,
			"apricotPhase3BlockTimestamp": 0,
			"apricotPhase4BlockTimestamp": 0
		},
		"nonce": "0x0",
		"timestamp": "0x0",
		"extraData": "0x00",
		"gasLimit": "0x5f5e100",
		"difficulty": "0x0",
		"mixHash": "0x0000000000000000000000000000000000000000000000000000000000000000",
		"coinbase": "0x0000000000000000000000000000000000000000",
		"alloc": {
			"0x96216849c49358B10257cb55b28eA603c874b05E": {
				"balance": "0x4EE2D6D415B85ACEF8100000000"
			}
		},
		"number": "0x0",
		"gasUsed": "0x0",
		"parentHash": "0x0000000000000000000000000000000000000000000000000000000000000000"
	}`

	// LocalParams are the params used for local networks
	LocalParams = Params{
		TxFeeConfig: TxFeeConfig{
			TxFee:                 units.MilliAvax,
			CreateAssetTxFee:      units.MilliAvax,
			CreateSubnetTxFee:     100 * units.MilliAvax,
			CreateBlockchainTxFee: 100 * units.MilliAvax,
		},
		StakingConfig: StakingConfig{
			UptimeRequirement:  .8, // 80%
			MinValidatorStake:  2 * units.KiloAvax,
			MaxValidatorStake:  3 * units.MegaAvax,
			MinDelegatorStake:  25 * units.Avax,
			MinDelegationFee:   20000, // 2%
			MinStakeDuration:   24 * time.Hour,
			MaxStakeDuration:   365 * 24 * time.Hour,
			StakeMintingPeriod: 365 * 24 * time.Hour,
		},
		EpochConfig: EpochConfig{
			EpochFirstTransition: time.Unix(1607626800, 0),
			EpochDuration:        6 * time.Hour,
		},
	}
)
