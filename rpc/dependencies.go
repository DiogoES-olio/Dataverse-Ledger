// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package rpc

import (
	"context"

	"dataverse/genesis"
	"dataverse/orderbook"
	"dataverse/storage"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/trace"
	"github.com/ava-labs/hypersdk/chain"
	"github.com/ava-labs/hypersdk/codec"
)

type Controller interface {
	Genesis() *genesis.Genesis
	Tracer() trace.Tracer
	GetTransaction(context.Context, ids.ID) (bool, int64, bool, chain.Dimensions, uint64, error)
	GetAssetFromState(context.Context, ids.ID) (bool, []byte, uint8, []byte, uint64, codec.Address, bool, error)
	GetBalanceFromState(context.Context, codec.Address, ids.ID) (uint64, error)
	Orders(pair string, limit int) []*orderbook.Order
	GetOrderFromState(context.Context, ids.ID) (
		bool, // exists
		ids.ID, // in
		uint64, // inTick
		ids.ID, // out
		uint64, // outTick
		uint64, // remaining
		codec.Address, // owner
		error,
	)
	GetLoanFromState(context.Context, ids.ID, ids.ID) (uint64, error)
	GetProjectFromState(context.Context, ids.ID) (bool, storage.ProjectData, error)
	GetUpdateFromState(context.Context, ids.ID) (bool, storage.UpdateData, error)
	GetMachineCID(context.Context, ids.ID) (bool, storage.RegisterMachineCIDData, error)
	GetAttestMachine(context.Context, ids.ID) (bool, storage.AttestMachineData, error)
	GetNotarizeData(context.Context, ids.ID) (bool, storage.NotarizeDataData, error)
}
