package helper_test

import (
	"testing"

	"github.com/luthfikhan/go-tlab-resep/helper"
	"github.com/stretchr/testify/assert"
)

func TestTypeConverter(t *testing.T) {
	type aStruct struct {
		Key1 string `json:"key1"`
		Key2 string `json:"key2"`
	}

	type bStruct struct {
		Key1 string `json:"key1"`
		Key3 string `json:"key2"`
	}
	a := aStruct{
		Key1: "1",
		Key2: "2",
	}

	b, _ := helper.TypeConverter[bStruct](a)
	assert.Equal(t, a.Key1, b.Key1)
	assert.Equal(t, a.Key2, b.Key3)
}

func TestTypeConverterDiferenceDataType(t *testing.T) {
	type aStruct struct {
		Key1 int    `json:"key1"`
		Key2 string `json:"key2"`
	}

	type bStruct struct {
		Key1 string `json:"key1"`
	}
	a := aStruct{
		Key1: 1,
		Key2: "2",
	}

	_, err := helper.TypeConverter[bStruct](a)
	assert.NotNil(t, err)
}
