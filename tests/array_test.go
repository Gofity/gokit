package tests

import (
	"testing"

	"github.com/Gofity/gokit"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type ArrayTestSuite struct {
	suite.Suite
}

func (x *ArrayTestSuite) TestAt() {
	data := gokit.Array[int]{1, 2, 3, 4}

	item, ok := data.At(2)
	require.Equal(x.T(), true, ok)
	require.Equal(x.T(), 3, item)

	_, ok = data.At(5)
	require.Equal(x.T(), false, ok)
}

func (x *ArrayTestSuite) TestLastIndex() {
	data := gokit.Array[int]{1, 2, 3, 4}
	require.Equal(x.T(), 3, data.LastIndex())
}

func (x *ArrayTestSuite) TestConcat() {
	data := gokit.Array[int]{1, 2, 3, 4}
	data = data.Concat([]int{5, 6})

	expected := gokit.Array[int]{1, 2, 3, 4, 5, 6}
	require.Equal(x.T(), true, data.Equal(expected))
}

func (x *ArrayTestSuite) TestAppend() {
	data := gokit.Array[int]{1, 2, 3, 4, 5}
	data.Append(6, 7, 8, 9, 10)

	expected := gokit.Array[int]{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	require.Equal(x.T(), true, data.Equal(expected))
}

func (x *ArrayTestSuite) TestPrepend() {
	data := gokit.Array[int]{1, 2, 3, 4, 5}
	data.Prepend(6, 7, 8, 9, 10)

	expected := gokit.Array[int]{6, 7, 8, 9, 10, 1, 2, 3, 4, 5}
	require.Equal(x.T(), true, data.Equal(expected))
}

func (x *ArrayTestSuite) TestFilter() {
	data := gokit.Array[int]{1, 2, 3, 4, 5}

	data = data.Filter(func(v int) bool {
		return (v % 2) != 0
	})

	expected := gokit.Array[int]{1, 3, 5}
	require.Equal(x.T(), true, data.Equal(expected))
}

func (x *ArrayTestSuite) TestFind() {
	data := gokit.Array[int]{2, 4, 6, 8, 10}

	item, ok := data.Find(func(v int) bool {
		return v >= 3
	})

	require.Equal(x.T(), true, ok)
	require.Equal(x.T(), 4, item)
}

func (x *ArrayTestSuite) TestFindLast() {
	data := gokit.Array[int]{2, 4, 6, 8, 10}

	item, ok := data.FindLast(func(v int) bool {
		return v >= 3
	})

	require.Equal(x.T(), true, ok)
	require.Equal(x.T(), 10, item)
}

func (x *ArrayTestSuite) TestFindIndex() {
	data := gokit.Array[int]{2, 4, 6, 8, 10}

	require.Equal(x.T(), 1, data.FindIndex(func(v int) bool {
		return v >= 3
	}))
}

func (x *ArrayTestSuite) TestFindLastIndex() {
	data := gokit.Array[int]{2, 4, 6, 8, 10}

	require.Equal(x.T(), 4, data.FindLastIndex(func(v int) bool {
		return v >= 3
	}))
}

func (x *ArrayTestSuite) TestReduce() {
	data := gokit.Array[int]{2, 4, 6, 8, 10}

	require.Equal(x.T(), 30, data.Reduce(func(accumulator, v int) int {
		return accumulator + v
	}))
}

func (x *ArrayTestSuite) TestSub() {
	data := gokit.Array[int]{2, 4, 6, 8, 10}

	sub := data.Sub(2)
	require.Equal(x.T(), true, sub.Equal([]int{6, 8, 10}))

	sub = data.Sub(2, 2)
	require.Equal(x.T(), true, sub.Equal([]int{6, 8}))
}

func (x *ArrayTestSuite) TestSlice() {
	data := gokit.Array[int]{2, 4, 6, 8, 10}

	sub := data.Slice(2)
	require.Equal(x.T(), true, sub.Equal([]int{6, 8, 10}))

	sub = data.Slice(2, 4)
	require.Equal(x.T(), true, sub.Equal([]int{6, 8}))
}

func (x *ArrayTestSuite) TestSplice() {
	data := gokit.Array[int]{2, 4, 6, 8, 10}

	value := data.Splice(2, 0, 200)
	require.Equal(x.T(), true, value.Equal([]int{2, 4, 200, 6, 8, 10}))

	value = data.Splice(1, 1, 101, 102)
	require.Equal(x.T(), true, value.Equal([]int{2, 101, 102, 6, 8, 10}))
}

func (x *ArrayTestSuite) TestMap() {
	data := gokit.Array[int]{2, 4, 6, 8, 10}

	value := data.Map(func(v int) int {
		return v * 2
	})

	require.Equal(x.T(), true, value.Equal([]int{4, 8, 12, 16, 20}))

	value = data.Map(func(v int) int {
		return v / 2
	})

	require.Equal(x.T(), true, value.Equal([]int{1, 2, 3, 4, 5}))
}

func (x *ArrayTestSuite) TestReverse() {
	data := gokit.Array[int]{1, 2, 3}
	data = data.Reverse()

	require.Equal(x.T(), true, data.Equal([]int{3, 2, 1}))
}

func (x *ArrayTestSuite) TestIndexOf() {
	data := gokit.Array[int]{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	require.Equal(x.T(), 2, data.IndexOf(3))
}

func (x *ArrayTestSuite) TestLastIndexOf() {
	data := gokit.Array[int]{1, 2, 3, 4, 5, 6, 7, 8, 5, 10}
	require.Equal(x.T(), 8, data.LastIndexOf(5))
}

func (x *ArrayTestSuite) TestPop() {
	data := gokit.Array[int]{1, 2, 3, 4, 5}

	item, ok := data.Pop()

	require.Equal(x.T(), true, ok)
	require.Equal(x.T(), 5, item)
	require.Equal(x.T(), true, data.Equal([]int{1, 2, 3, 4}))
}

func (x *ArrayTestSuite) TestShift() {
	data := gokit.Array[int]{1, 2, 3, 4, 5}

	item, ok := data.Shift()

	require.Equal(x.T(), true, ok)
	require.Equal(x.T(), 1, item)
	require.Equal(x.T(), true, data.Equal([]int{2, 3, 4, 5}))
}

func (x *ArrayTestSuite) TestJoin() {
	data := gokit.Array[int]{1, 2, 3, 4, 5}

	expected := "1-2-3-4-5"
	require.Equal(x.T(), expected, data.Join("-"))
}

// =======================

func TestArray(t *testing.T) {
	suite.Run(t, new(ArrayTestSuite))
}
