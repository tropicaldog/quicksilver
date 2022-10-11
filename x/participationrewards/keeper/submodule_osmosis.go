package keeper

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"

	osmosistypes "github.com/ingenuity-build/quicksilver/osmosis-types"
	osmolockup "github.com/ingenuity-build/quicksilver/osmosis-types/lockup"
	"github.com/ingenuity-build/quicksilver/x/participationrewards/types"
)

type OsmosisModule struct{}

var _ Submodule = &OsmosisModule{}

func (m *OsmosisModule) Hooks(ctx sdk.Context, k Keeper) {
	// osmosis params
	params, found := k.GetProtocolData(ctx, "osmosis/params")
	if !found {
		k.Logger(ctx).Error("unable to query osmosis/params in OsmosisModule hook")
		return
	}

	paramsData := types.OsmosisParamsProtocolData{}
	if err := json.Unmarshal(params.Data, &paramsData); err != nil {
		k.Logger(ctx).Error("unable to unmarshal osmosis/params in OsmosisModule hook", "error", err)
		return
	}

	data, found := k.GetProtocolData(ctx, fmt.Sprintf("connection/%s", paramsData.ChainID))
	if !found {
		k.Logger(ctx).Error(fmt.Sprintf("unable to query connection/%s in OsmosisModule hook", paramsData.ChainID))
		return
	}

	connectionData := types.ConnectionProtocolData{}
	if err := json.Unmarshal(data.Data, &connectionData); err != nil {
		k.Logger(ctx).Error(fmt.Sprintf("unable to unmarshal connection/%s in OsmosisModule hook", paramsData.ChainID))
		return
	}

	k.IteratePrefixedProtocolDatas(ctx, osmosistypes.PoolsPrefix, func(idx int64, data types.ProtocolData) bool {
		ipool, err := types.UnmarshalProtocolData(types.ProtocolDataOsmosisPool, data.Data)
		if err != nil {
			return false
		}
		pool, _ := ipool.(types.OsmosisPoolProtocolData)

		// update pool datas
		k.IcqKeeper.MakeRequest(ctx, connectionData.ConnectionID, connectionData.ChainID, "store/gamm/key", m.GetKeyPrefixPools(pool.PoolID), sdk.NewInt(-1), types.ModuleName, "osmosispoolupdate", 0) // query pool data
		return false
	})
}

func (m *OsmosisModule) IsActive() bool {
	return true
}

func (m *OsmosisModule) IsReady() bool {
	return true
}

func (m *OsmosisModule) ValidateClaim(ctx sdk.Context, k *Keeper, msg *types.MsgSubmitClaim) (uint64, error) {
	var amount uint64
	for _, proof := range msg.Proofs {
		lockupResponse := osmolockup.LockedResponse{}
		err := k.cdc.Unmarshal(proof.Data, &lockupResponse)
		if err != nil {
			return 0, err
		}

		_, lockupOwner, err := bech32.DecodeAndConvert(lockupResponse.Lock.Owner)
		if err != nil {
			return 0, err
		}

		if sdk.AccAddress(lockupOwner).String() != msg.UserAddress {
			return 0, fmt.Errorf("not a valid proof for submitting user")
		}

		sdkAmount, err := osmosistypes.DetermineApplicableTokensInPool(ctx, k, lockupResponse, msg.Zone)
		if err != nil {
			return 0, err
		}

		if sdkAmount.IsNegative() {
			return 0, fmt.Errorf("unexpected negative amount")
		}
		amount += sdkAmount.Uint64()
	}
	return amount, nil
}

func (m *OsmosisModule) GetKeyPrefixPools(poolID uint64) []byte {
	return append([]byte{0x02}, sdk.Uint64ToBigEndian(poolID)...)
}
