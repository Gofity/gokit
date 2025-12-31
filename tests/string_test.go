package tests

import (
	"testing"

	"github.com/Gofity/gokit"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type StringTestSuite struct {
	suite.Suite
}

func (x *StringTestSuite) TestAppend() {
	data := gokit.String("Hello ")
	data.Append("World")

	require.Equal(x.T(), "Hello World", string(data))
}

func (x *StringTestSuite) TestPrepend() {
	data := gokit.String(" World")
	data.Prepend("Hello")

	require.Equal(x.T(), "Hello World", string(data))
}

func (x *StringTestSuite) TestSplit() {
	data := gokit.String("Hello World")
	chunks := data.Split(" ")

	require.Equal(x.T(), true, chunks.Equal([]gokit.String{
		"Hello",
		"World",
	}))
}

// =======================

func TestString(t *testing.T) {
	suite.Run(t, new(StringTestSuite))
}
