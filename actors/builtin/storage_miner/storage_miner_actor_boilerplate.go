package storage_miner

import (
	"github.com/filecoin-project/specs-actors/actors/abi"
	vmr "github.com/filecoin-project/specs-actors/actors/runtime"
	autil "github.com/filecoin-project/specs-actors/actors/util"
)

type SectorStorageWeightDesc = autil.SectorStorageWeightDesc
type SectorTerminationType = autil.SectorTermination

var RT_ConfirmFundsReceiptOrAbort_RefundRemainder = vmr.RT_ConfirmFundsReceiptOrAbort_RefundRemainder

////////////////////////////////////////////////////////////////////////////////
// Boilerplate
//
// This boilerplate should be essentially identical for all actors, and
// conceptually belongs in the runtime/VM. It is only duplicated here as a
// workaround due to the lack of generics support in Go.
////////////////////////////////////////////////////////////////////////////////

type InvocOutput = vmr.InvocOutput
type Runtime = vmr.Runtime

var Assert = autil.Assert
var IMPL_FINISH = autil.IMPL_FINISH
var IMPL_TODO = autil.IMPL_TODO
var TODO = autil.TODO

func Release(rt Runtime, h vmr.ActorStateHandle, st StorageMinerActorState) {
	checkCID := abi.ActorSubstateCID(rt.IpldPut(&st))
	h.Release(checkCID)
}

func UpdateRelease(rt Runtime, h vmr.ActorStateHandle, st StorageMinerActorState) {
	newCID := abi.ActorSubstateCID(rt.IpldPut(&st))
	h.UpdateRelease(newCID)
}