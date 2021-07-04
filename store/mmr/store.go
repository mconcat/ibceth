package mmr

import (
	"encoding/binary"
	"math/bits"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/cosmos/cosmos-sdk/store/types"

	dbm "github.com/tendermint/tm-db"
)

var _ types.CommitStore = &Store{}

type Store struct {
	db dbm.DB
}

// index should be 1-indexed
func height(index uint64) uint64 {
	for {
		// get bit size of index
		bitlen := bits.Len64(index)

		// stop if all bits are 1
		if bitlen == bits.OnesCount64(index) {
			return uint64(bits.Len64(index)) - 1
		}

		// subtract largest 2^n-1 from index where 2^n < index
		index -= (1 << (bitlen - 1)) - 1
	}
}

func isLeft(index uint64) bool {
	for {
		if index%2 == 0 {
			bitlen := bits.Len64(index / 2)
			if bitlen == bits.OnesCount64(index/2) {
				return false
			}
		} else {
			bitlen := bits.Len64(index)
			if bitlen == bits.OnesCount64(index) {
				return true
			}
		}

		bitlen := bits.Len64(index)
		index -= (1 << (bitlen - 1)) - 1
	}
}

func isLeaf(index uint64) bool {
	return height(index) == 0
}

func leftSibling(index uint64) uint64 {
	height := height(index)
	return index - ((1 << (height + 1)) - 1)
}

func rightSibling(index uint64) uint64 {
	height := height(index)
	return index + (1 << (height + 1)) - 1
}

func leftChild(index uint64) uint64 {
	height := height(index)
	return index - (1 << height)
}

func rightChild(index uint64) uint64 {
	return index - 1
}

// https://github.com/mimblewimble/grin/blob/master/doc/mmr.md
func hashLeaf(index uint64, value []byte) []byte {
	// keccak256(abi.encodePacked(uint64(node.index)), node.value)
	return crypto.Keccak256(getKey(index), value)
}

func hashNode(index uint64, left, right []byte) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, index)
	// keccak256(abi.encodePacked(uint64(node.index)), left, right)
	return crypto.Keccak256(getKey(index), left, right)
}

func (store *Store) Commit() types.CommitID {

}

func (store *Store) Append(value []byte) uint64 {
	size := store.Size()
	index := size + 1

	for {
		store.Set(getKey(index), value)

		if isLeft(index) {
			return size + 1
		}

		index++
	}
}

func (store *Store) Delete(index uint64) bool {

}

func (store *Store) Get(index uint64) []byte {

}

func (store *Store) Has(index uint64) bool {

}

func (store *Store) Size() uint64 {
	iter, err := store.db.ReverseIterator(nil, nil)
	if err != nil {
		panic(err)
	}
	defer iter.Close()

	if !iter.Valid() {
		return 0
	}

	return binary.BigEndian.Uint64(iter.Key())
}

func getKey(index uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, index)
	return bz
}
