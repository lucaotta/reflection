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

func TestMarshalSlice(t *testing.T) {
	res, err := Marshal([]int{})
	assert.Nil(t, err)
	assert.EqualValues(t, []byte{'[', ']'}, res)
}

func TestMarshalSlice2(t *testing.T) {
	res, err := Marshal([]int{1, -2})
	assert.Nil(t, err)
	assert.EqualValues(t, []byte{'[', '1', ',', '-', '2', ']'}, res)
}

func TestMarshalArray(t *testing.T) {
	res, err := Marshal([0]int{})
	assert.Nil(t, err)
	assert.EqualValues(t, []byte{'[', ']'}, res)
}

func TestMarshalArray2(t *testing.T) {
	res, err := Marshal([3]uint16{2, 3, 4})
	assert.Nil(t, err)
	assert.EqualValues(t, []byte{'[', '2', ',', '3', ',', '4', ']'}, res)
}
