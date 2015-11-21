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

func TestMarshalEmptyStruct(t *testing.T) {
	type Empty struct{}
	res, err := Marshal(Empty{})
	assert.Nil(t, err)
	assert.EqualValues(t, []byte{'{', '}'}, res)
}

func TestMarshalUnexportedFields(t *testing.T) {
	type Private struct{ foo int }
	res, err := Marshal(Private{})
	assert.Nil(t, err)
	assert.EqualValues(t, []byte{'{', '}'}, res)
}

func TestMarshalExportedFields(t *testing.T) {
	type Public struct{ Foo int }
	res, err := Marshal(Public{})
	assert.Nil(t, err)
	assert.EqualValues(t, []byte{'{', 'F', 'o', 'o', ':', '0', '}'}, res)
}

func TestMarshalExportedFields2(t *testing.T) {
	type Public struct{ Foo, Bar int }
	res, err := Marshal(Public{})
	assert.Nil(t, err)
	assert.EqualValues(t, []byte{'{', 'F', 'o', 'o', ':', '0', ',', 'B', 'a', 'r', ':', '0', '}'}, res)
}

func TestMarshalMixedFields(t *testing.T) {
	type Public struct{ Foo, bar, Baz int }
	res, err := Marshal(Public{})
	assert.Nil(t, err)
	assert.EqualValues(t, []byte{'{', 'F', 'o', 'o', ':', '0', ',', 'B', 'a', 'z', ':', '0', '}'}, res)
}

type PublicWithFunc struct{ Foo int }

func (p PublicWithFunc) F() {}

func TestMarshalStructWithFunc(t *testing.T) {
	res, err := Marshal(PublicWithFunc{})
	assert.Nil(t, err)
	assert.EqualValues(t, []byte{'{', 'F', 'o', 'o', ':', '0', '}'}, res)
}

type PublicWithTag struct {
	Foo int `json:"foo"`
}

func TestMarshalWithTag(t *testing.T) {
	res, err := Marshal(PublicWithTag{Foo: 1})
	assert.Nil(t, err)
	assert.EqualValues(t, []byte{'{', 'f', 'o', 'o', ':', '1', '}'}, res)
}

func TestMarshalInnerStruct(t *testing.T) {
	type I struct{ Foo int }
	type O struct{ Bar I }
	res, err := Marshal(O{})
	assert.Nil(t, err)
	assert.EqualValues(t, []byte{'{', 'B', 'a', 'r', ':', '{', 'F', 'o', 'o', ':', '0', '}', '}'}, res)
}

func TestMarshalInnerArray(t *testing.T) {
	type O struct{ Bar []int }
	res, err := Marshal(O{Bar: []int{1, 2}})
	assert.Nil(t, err)
	assert.EqualValues(t, []byte{'{', 'B', 'a', 'r', ':', '[', '1', ',', '2', ']', '}'}, res)
}
