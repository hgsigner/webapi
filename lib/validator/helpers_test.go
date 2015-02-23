package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_filterValidationTag(t *testing.T) {
	a := assert.New(t)

	tests := []struct {
		tagType, value, out string
	}{
		{"tname", "presence", "presence"},
		{"tname", "max=0", "max"},
		{"tname", "min=200", "min"},
		{"tvalue", "presence", ""},
		{"tvalue", "max=0", "0"},
		{"tvalue", "min=200", "200"},
	}

	for _, v := range tests {
		r := filterValidationTag(v.tagType, v.value)
		a.Equal(v.out, r)
	}

}
