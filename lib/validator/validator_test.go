package validator_test

import (
	"projects/webapi/lib/validator"
	"testing"

	"github.com/stretchr/testify/assert"
)

type struct1 struct {
	Name  string
	Email string
	Age   int
}

type struct2 struct {
	Name  string `validates:"presence"`
	Email string `validates:"presence,format=/^[A-Z0-9._%a-z\-]+@(?:[A-Z0-9a-z\-]+\.)+[A-Za-z]{2,4}$/i"`
	Age   int    `validates:"min=18,max=100"`
}

func Test_Struct_Without_Tags(t *testing.T) {
	a := assert.New(t)

	strc := struct1{
		"Hugo",
		"test@test.com",
		10,
	}

	valid := validator.Valid(strc)

	a.True(valid)
	a.Equal(0, len(validator.ErrosListHolder))

}

func Test_Struct_Being_Valid(t *testing.T) {
	a := assert.New(t)

	strc := struct2{
		"Hugo",
		"test@test.com",
		30,
	}

	valid := validator.Valid(strc)

	a.True(valid)
	a.Equal(0, len(validator.ErrosListHolder))

}

func Test_Struct_Not_Valid_Min(t *testing.T) {
	a := assert.New(t)

	strc := struct2{
		"Hugo",
		"test@test.com",
		10,
	}

	valid := validator.Valid(strc)

	a.False(valid)
	a.Equal(1, len(validator.ErrosListHolder))
}

func Test_Struct_Not_Valid_Max(t *testing.T) {
	a := assert.New(t)

	strc := struct2{
		"Hugo",
		"test@test.com",
		200,
	}

	valid := validator.Valid(strc)

	a.False(valid)
	a.Equal(1, len(validator.ErrosListHolder))
}

func Test_Struct_Not_Valid_Format(t *testing.T) {
	a := assert.New(t)

	strc := struct2{
		"Hugo",
		"test",
		30,
	}

	valid := validator.Valid(strc)

	a.False(valid)
	a.Equal(1, len(validator.ErrosListHolder))
}
