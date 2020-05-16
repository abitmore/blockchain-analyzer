package processor

import (
	"testing"

	"github.com/danhper/blockchain-data-fetcher/core"
	"github.com/danhper/blockchain-data-fetcher/xrp"
	"github.com/stretchr/testify/assert"
)

func TestComputeBlockNumbers(t *testing.T) {
	reader := core.GetXRPLedgersReader(core.RealLedgersFilename)
	blockchain := xrp.New()
	blocks, err := ComputeBlockNumbers(reader, blockchain)
	assert.Nil(t, err)
	assert.Len(t, blocks, 100)
	assert.Contains(t, blocks, uint64(54387329))
}

func TestGetMissingBlockNumbersValid(t *testing.T) {
	reader := core.GetXRPLedgersReader(core.RealLedgersFilename)
	blockchain := xrp.New()
	missing, err := GetMissingBlockNumbers(reader, blockchain)
	assert.Nil(t, err)
	assert.Len(t, missing, 0)
}

func TestGetMissingBlockNumbersInvalid(t *testing.T) {
	reader := core.GetXRPLedgersReader(core.MissingLedgersFilename)
	blockchain := xrp.New()
	missing, err := GetMissingBlockNumbers(reader, blockchain)
	assert.Nil(t, err)
	assert.Len(t, missing, 1)
	assert.Equal(t, missing[0], uint64(124))
}
