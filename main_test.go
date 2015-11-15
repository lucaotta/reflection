package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalInt(t *testing.T) {
	res, err := Marshal(-4)
	assert.Nil(t, err)
	assert.EqualValues(t, []byte{'-', '4'}, res)
}

func TestMarshalIntMoreDigits(t *testing.T) {
	res, err := Marshal(-124)
	assert.Nil(t, err)
	assert.EqualValues(t, []byte{'-', '1', '2', '4'}, res)
}

func TestMarshalUint(t *testing.T) {
	res, err := Marshal(4)
	assert.Nil(t, err)
	assert.EqualValues(t, []byte{'4'}, res)
}

func TestMarshalUint2(t *testing.T) {
	res, err := Marshal(124)
	assert.Nil(t, err)
	assert.EqualValues(t, []byte{'1', '2', '4'}, res)
}
