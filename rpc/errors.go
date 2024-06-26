// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package rpc

import "errors"

var (
	ErrTxNotFound            = errors.New("tx not found")
	ErrAssetNotFound         = errors.New("asset not found")
	ErrOrderNotFound         = errors.New("order not found")
	ErrProjectNotFound       = errors.New("project not found")
	ErrUpdateNotFound        = errors.New("update not found")
	ErrMachineCIDNotFound    = errors.New("machine cid not found")
	ErrAttestMachineNotFound = errors.New("attested Machine not found")
	ErrNotarizedDataNotFound = errors.New("Invalid Notarized Data")
)
