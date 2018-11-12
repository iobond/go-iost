package cverifier

import (
	"bytes"
	"errors"
	"time"

	"github.com/iost-official/go-iost/core/block"
)

var (
	errFutureBlk  = errors.New("block from future")
	errOldBlk     = errors.New("block too old")
	errParentHash = errors.New("wrong parent hash")
	errNumber     = errors.New("wrong number")
	errTxHash     = errors.New("wrong txs hash")
	errMerkleHash = errors.New("wrong tx receipt merkle hash")

	// TxExecTimeLimit the maximum verify execution time of a transaction
	TxExecTimeLimit = 400 * time.Millisecond
)

// VerifyBlockHead verifies the block head.
func VerifyBlockHead(blk *block.Block, parentBlock *block.Block, lib *block.Block) error {
	bh := blk.Head
	if bh.Time > time.Now().UnixNano() {
		return errFutureBlk
	}
	if bh.Time < lib.Head.Time {
		return errOldBlk
	}
	if !bytes.Equal(bh.ParentHash, parentBlock.HeadHash()) {
		return errParentHash
	}
	if bh.Number != parentBlock.Head.Number+1 {
		return errNumber
	}
	if !bytes.Equal(blk.CalculateTxsHash(), bh.TxsHash) {
		return errTxHash
	}
	if !bytes.Equal(blk.CalculateMerkleHash(), bh.MerkleHash) {
		return errMerkleHash
	}
	return nil
}
