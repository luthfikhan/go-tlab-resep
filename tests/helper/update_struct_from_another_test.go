package helper_test

import (
	"testing"

	"github.com/luthfikhan/go-tlab-resep/helper"
	"github.com/stretchr/testify/assert"
)

func TestUpdateStructFromAnother(t *testing.T) {
	type aStruct struct {
		Key1 string `json:"key1"`
		Key2 string `json:"key2"`
	}

	type bStruct struct {
		Key1 string `json:"key1"`
		Key2 string `json:"key2"`
	}
	toUpdate := aStruct{
		Key1: "1",
		Key2: "2",
	}
	updateFrom := bStruct{
		Key1: "1_v2",
		Key2: "2_v2",
	}

	toUpdate = *helper.UpdateStructFromAnother(toUpdate, updateFrom)
	assert.Equal(t, toUpdate.Key1, updateFrom.Key1)
	assert.Equal(t, toUpdate.Key2, updateFrom.Key2)
}
