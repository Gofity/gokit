package tests

import (
	"testing"

	"github.com/Gofity/gokit"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type MapTestSuite struct {
	suite.Suite
}

func (x *MapTestSuite) TestMerge() {
	data := gokit.Map[string, int]{"students": 1}
	data = data.Merge(map[string]int{"pets": 3})

	students, ok := data["students"]
	require.Equal(x.T(), true, ok)
	require.Equal(x.T(), 1, students)

	pets, ok := data["pets"]
	require.Equal(x.T(), true, ok)
	require.Equal(x.T(), 3, pets)

	data = data.Merge(map[string]int{"pets": 5})

	pets, ok = data["pets"]
	require.Equal(x.T(), true, ok)
	require.Equal(x.T(), 5, pets)
}

func (x *MapTestSuite) TestKeys() {
	data := gokit.Map[string, int]{
		"students": 1,
		"pets":     3,
	}

	keys := data.Keys()
	require.Equal(x.T(), 2, keys.Size())
}

func (x *MapTestSuite) TestValues() {
	data := gokit.Map[string, int]{
		"students": 1,
		"pets":     3,
	}

	var keys gokit.Array[int] = data.Values()
	require.Equal(x.T(), 2, keys.Size())
}

// =======================

func TestMap(t *testing.T) {
	suite.Run(t, new(MapTestSuite))
}
