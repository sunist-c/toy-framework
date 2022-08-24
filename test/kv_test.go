package test

import (
	"fmt"
	"github.com/sunist-c/toy-framework/ds/kv"
	"sync"
	"testing"
)

type TestCase[Key, Value any] struct {
	Indexes        []int
	StoreKeys      map[int]Key
	StoreValues    map[int]Value
	LoadKeys       map[int]Key
	WantLoadValues map[int]Value
	WantLoaded     map[int]bool
	DeleteKeys     map[int]Key
	RangeFunc      func(Key, Value) bool
	Equals         func(v1, v2 Value) bool
}

type testStruct struct {
	val int
}

var (
	intIntCase TestCase[int, int] = TestCase[int, int]{
		Indexes: []int{1, 2, 3},
		StoreKeys: map[int]int{
			1: 114514,
			2: 1919810,
			3: 2147,
		},
		StoreValues: map[int]int{
			1: 1919810,
			2: 114514,
			3: 65535,
		},
		DeleteKeys: map[int]int{
			1: 2147,
			2: 2147,
			3: 2147,
		},
		LoadKeys: map[int]int{
			1: 2147,
			2: 114514,
			3: 1919810,
		},
		WantLoaded: map[int]bool{
			1: false,
			2: true,
			3: true,
		},
		WantLoadValues: map[int]int{
			1: 0,
			2: 1919810,
			3: 114514,
		},
		RangeFunc: func(a, b int) bool {
			fmt.Printf("key: %d, value: %d\n", a, b)
			return true
		},
		Equals: func(a, b int) bool {
			return a == b
		},
	}
	stringStringCase TestCase[string, string] = TestCase[string, string]{
		Indexes: []int{1, 2, 3},
		StoreKeys: map[int]string{
			1: "114514",
			2: "1919810",
			3: "2147",
		},
		StoreValues: map[int]string{
			1: "1919810",
			2: "114514",
			3: "65535",
		},
		DeleteKeys: map[int]string{
			1: "2147",
			2: "2147",
			3: "2147",
		},
		LoadKeys: map[int]string{
			1: "2147",
			2: "114514",
			3: "1919810",
		},
		WantLoaded: map[int]bool{
			1: false,
			2: true,
			3: true,
		},
		WantLoadValues: map[int]string{
			1: "",
			2: "1919810",
			3: "114514",
		},
		RangeFunc: func(a, b string) bool {
			fmt.Printf("key: %v, value: %v\n", a, b)
			return true
		},
		Equals: func(a, b string) bool {
			return a == b
		},
	}
	intStructCase TestCase[int, testStruct] = TestCase[int, testStruct]{
		Indexes: []int{1, 2, 3},
		StoreKeys: map[int]int{
			1: 114514,
			2: 1919810,
			3: 2147,
		},
		StoreValues: map[int]testStruct{
			1: {val: 1919810},
			2: {val: 114514},
			3: {val: 65535},
		},
		DeleteKeys: map[int]int{
			1: 2147,
			2: 2147,
			3: 2147,
		},
		LoadKeys: map[int]int{
			1: 2147,
			2: 114514,
			3: 1919810,
		},
		WantLoaded: map[int]bool{
			1: false,
			2: true,
			3: true,
		},
		WantLoadValues: map[int]testStruct{
			1: {},
			2: {val: 1919810},
			3: {val: 114514},
		},
		RangeFunc: func(a int, b testStruct) bool {
			fmt.Printf("key: %d, value: %#v\n", a, b)
			return true
		},
		Equals: func(a, b testStruct) bool {
			return a.val == b.val
		},
	}
	structStructCase TestCase[testStruct, testStruct] = TestCase[testStruct, testStruct]{
		Indexes: []int{1, 2, 3},
		StoreKeys: map[int]testStruct{
			1: {val: 114514},
			2: {val: 1919810},
			3: {val: 2147},
		},
		StoreValues: map[int]testStruct{
			1: {val: 1919810},
			2: {val: 114514},
			3: {val: 65535},
		},
		DeleteKeys: map[int]testStruct{
			1: {val: 2147},
			2: {val: 2147},
			3: {val: 2147},
		},
		LoadKeys: map[int]testStruct{
			1: {val: 2147},
			2: {val: 114514},
			3: {val: 1919810},
		},
		WantLoaded: map[int]bool{
			1: false,
			2: true,
			3: true,
		},
		WantLoadValues: map[int]testStruct{
			1: {},
			2: {val: 1919810},
			3: {val: 114514},
		},
		RangeFunc: func(a, b testStruct) bool {
			fmt.Printf("key: %#v, value: %#v\n", a, b)
			return true
		},
		Equals: func(a, b testStruct) bool {
			return a.val == b.val
		},
	}
)

func UnitTestFunc[Key, Value any](m kv.IMap[Key, Value], testCase *TestCase[Key, Value], t *testing.T) {
	for _, index := range testCase.Indexes {
		m.Store(testCase.StoreKeys[index], testCase.StoreValues[index])
	}

	for _, index := range testCase.Indexes {
		m.Delete(testCase.DeleteKeys[index])
	}

	for _, index := range testCase.Indexes {
		got, gotten := m.Load(testCase.LoadKeys[index])
		if gotten != testCase.WantLoaded[index] {
			t.Errorf("key %v want loaded: %v, but load: %v", testCase.LoadKeys[index], testCase.WantLoaded[index], gotten)
		}
		if !testCase.Equals(got, testCase.WantLoadValues[index]) {
			t.Errorf("key %v want got: %v, but got: %v", testCase.LoadKeys[index], testCase.WantLoadValues[index], got)
		}
	}

	m.Range(testCase.RangeFunc)
}

func TestKv(t *testing.T) {
	t.Run("test-intIntCase", func(t *testing.T) {
		m := kv.NewKv[int, int]()
		UnitTestFunc(m, &intIntCase, t)
	})
	t.Run("test-stringStringCase", func(t *testing.T) {
		m := kv.NewKv[string, string]()
		UnitTestFunc(m, &stringStringCase, t)
	})
	t.Run("test-structStructCase", func(t *testing.T) {
		m := kv.NewKv[testStruct, testStruct]()
		UnitTestFunc(m, &structStructCase, t)
	})
	t.Run("test-intStructCase", func(t *testing.T) {
		m := kv.NewKv[int, testStruct]()
		UnitTestFunc(m, &intStructCase, t)
	})
}

func BenchmarkTestKv(b *testing.B) {
	goroutines, operations := 100, 10000
	for i := 0; i < b.N; i++ {
		m := kv.NewKv[int, int](10007)
		wg := &sync.WaitGroup{}
		wg.Add(goroutines * operations)
		for j := 0; j < goroutines; j++ {
			go func(index int) {
				for k := 0; k < operations; k++ {
					m.Store(index*operations+k, k)
					m.Load(index*operations + k)
					wg.Done()
				}
			}(j)
		}
		wg.Wait()
	}
}
