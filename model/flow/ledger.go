// (c) 2019 Dapper Labs - ALL RIGHTS RESERVED

package flow

import (
	"github.com/dapperlabs/flow-go/crypto"
)

// Ledger is a map of register values.
type Ledger map[string][]byte

// MergeWith inserts all key/value pairs from another ledger into this one.
func (l Ledger) MergeWith(ledger Ledger) {
	for key, value := range ledger {
		if ledger[key] == nil {
			delete(l, key)
		} else {
			l[key] = value
		}
	}
}

// NewView returns a new read-only view onto this ledger.
func (l Ledger) NewView() *LedgerView {
	return &LedgerView{
		new: make(Ledger),
		old: l,
	}
}

// LedgerView provides a read-only view into an existing ledger set.
//
// Values are written to a temporary register cache that can later be
// committed to the world state.
type LedgerView struct {
	new Ledger
	old Ledger
}

// Updated returns the set of registers that were written to this view.
func (r *LedgerView) Updated() Ledger {
	return r.new
}

// Get gets a register from this view.
func (r *LedgerView) Get(key string) (value []byte, exists bool) {
	value = r.new[key]
	if value != nil {
		return value, true
	}

	value = r.old[key]
	if value == nil {
		return nil, false
	}

	return value, true
}

// Set sets a register in this view.
func (r *LedgerView) Set(key string, value []byte) {
	r.new[key] = value
}

// Delete deletes a register in this view.
func (r *LedgerView) Delete(key string) {
	r.new[key] = nil
}

type IntermediateRegisters struct {
	TransactionHash crypto.Hash
	Registers       Ledger
	ComputeUsed     uint64
}
