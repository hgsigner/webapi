package validator

import (
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
	Email string `validates:"presence,format=email"`
	Age   int    `validates:"min=18,max=100"`
}

func Test_Struct_Without_Tags(t *testing.T) {
	a := assert.New(t)

	strc := struct1{
		"Hugo",
		"test@test.com",
		10,
	}

	valid := Valid(strc)

	a.True(valid)
	a.Equal(0, len(ErrosListHolder))

}

func Test_Struct_Being_Valid(t *testing.T) {
	a := assert.New(t)

	strc := struct2{
		"Hugo",
		"test@test.com",
		30,
	}

	valid := Valid(strc)

	a.True(valid)
	a.Equal(0, len(ErrosListHolder))

}

func Test_Struct_Not_Valid_Min(t *testing.T) {
	a := assert.New(t)

	strc := struct2{
		"Hugo",
		"test@test.com",
		10,
	}

	valid := Valid(strc)

	a.False(valid)
	a.Equal(1, len(ErrosListHolder))
	a.Equal("age", ErrosListHolder[0].Field)
	a.Equal("Must be greater than 18", ErrosListHolder[0].Message)
}

func Test_Struct_Not_Valid_Max(t *testing.T) {
	a := assert.New(t)

	strc := struct2{
		"Hugo",
		"test@test.com",
		200,
	}

	valid := Valid(strc)

	a.False(valid)
	a.Equal(1, len(ErrosListHolder))
	a.Equal("age", ErrosListHolder[0].Field)
	a.Equal("Must be less than 100", ErrosListHolder[0].Message)
}

func Test_Struct_Not_Valid_Format(t *testing.T) {
	a := assert.New(t)

	strc := struct2{
		"Hugo",
		"test",
		30,
	}

	valid := Valid(strc)

	a.False(valid)
	a.Equal(1, len(ErrosListHolder))
	a.Equal("email", ErrosListHolder[0].Field)
	a.Equal("Format must match", ErrosListHolder[0].Message)
}
