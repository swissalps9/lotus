package reward

import (
	"github.com/filecoin-project/go-state-types/abi"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/go-state-types/cbor"
	builtin0 "github.com/filecoin-project/specs-actors/actors/builtin"

	"github.com/filecoin-project/lotus/chain/actors/adt"
	"github.com/filecoin-project/lotus/chain/actors/builtin"
	"github.com/filecoin-project/lotus/chain/types"
)

var Address = builtin0.RewardActorAddr

func Load(store adt.Store, act *types.Actor) (st State, err error) {
	switch act.Code {
	case builtin0.RewardActorCodeID:
		out := state0{store: store}
		err := store.Get(store.Context(), act.Head, &out)
		if err != nil {
			return nil, err
		}
		return &out, nil
	}
	return nil, xerrors.Errorf("unknown actor code %s", act.Code)
}

type State interface {
	cbor.Er

	ThisEpochBaselinePower() (abi.StoragePower, error)
	ThisEpochReward() (abi.StoragePower, error)
	ThisEpochRewardSmoothed() (builtin.FilterEstimate, error)

	EffectiveBaselinePower() (abi.StoragePower, error)
	EffectiveNetworkTime() (abi.ChainEpoch, error)

	TotalStoragePowerReward() (abi.TokenAmount, error)

	CumsumBaseline() (abi.StoragePower, error)
	CumsumRealized() (abi.StoragePower, error)
}