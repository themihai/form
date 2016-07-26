package form

import (
	"errors"
	"net/url"
	"testing"
	"time"

	. "gopkg.in/go-playground/assert.v1"
)

// NOTES:
// - Run "go test" to run tests
// - Run "gocov test | gocov report" to report on test converage by file
// - Run "gocov test | gocov annotate -" to report on all code and functions, those ,marked with "MISS" were never called
//
// or
//
// -- may be a good idea to change to output path to somewherelike /tmp
// go test -coverprofile cover.out && go tool cover -html=cover.out -o cover.html
//
//
// go test -cpuprofile cpu.out
// ./validator.test -test.bench=. -test.cpuprofile=cpu.prof
// go tool pprof validator.test cpu.prof
//
//
// go test -memprofile mem.out

func TestDecoderInt(t *testing.T) {

	type TestInt struct {
		Int              int
		Int8             int8
		Int16            int16
		Int32            int32
		Int64            int64
		IntPtr           *int
		Int8Ptr          *int8
		Int16Ptr         *int16
		Int32Ptr         *int32
		Int64Ptr         *int64
		IntArray         []int
		IntPtrArray      []*int
		IntArrayArray    [][]int
		IntPtrArrayArray [][]*int
		IntMap           map[int]int
		IntPtrMap        map[*int]*int
		NoURLValue       int
		IntNoValues      int
		Int8NoValues     int8
		Int16NoValues    int16
		Int32NoValues    int32
		Int64NoValues    int64
	}

	values := url.Values{
		"Int":                    []string{"3"},
		"Int8":                   []string{"3"},
		"Int16":                  []string{"3"},
		"Int32":                  []string{"3"},
		"Int64":                  []string{"3"},
		"IntPtr":                 []string{"3"},
		"Int8Ptr":                []string{"3"},
		"Int16Ptr":               []string{"3"},
		"Int32Ptr":               []string{"3"},
		"Int64Ptr":               []string{"3"},
		"IntArray":               []string{"1", "2", "3"},
		"IntPtrArray[0]":         []string{"1"},
		"IntPtrArray[2]":         []string{"3"},
		"IntArrayArray[0][0]":    []string{"1"},
		"IntArrayArray[0][2]":    []string{"3"},
		"IntArrayArray[2][0]":    []string{"1"},
		"IntPtrArrayArray[0][0]": []string{"1"},
		"IntPtrArrayArray[0][2]": []string{"3"},
		"IntPtrArrayArray[2][0]": []string{"1"},
		"IntMap[1]":              []string{"3"},
		"IntPtrMap[1]":           []string{"3"},
	}

	var test TestInt

	test.IntArray = make([]int, 4)

	decoder := NewDecoder()
	errs := decoder.Decode(&test, values)
	Equal(t, errs, nil)

	Equal(t, test.Int, int(3))
	Equal(t, test.Int8, int8(3))
	Equal(t, test.Int16, int16(3))
	Equal(t, test.Int32, int32(3))
	Equal(t, test.Int64, int64(3))

	Equal(t, *test.IntPtr, int(3))
	Equal(t, *test.Int8Ptr, int8(3))
	Equal(t, *test.Int16Ptr, int16(3))
	Equal(t, *test.Int32Ptr, int32(3))
	Equal(t, *test.Int64Ptr, int64(3))

	Equal(t, len(test.IntArray), 4)
	Equal(t, test.IntArray[0], int(1))
	Equal(t, test.IntArray[1], int(2))
	Equal(t, test.IntArray[2], int(3))
	Equal(t, test.IntArray[3], int(0))

	Equal(t, len(test.IntPtrArray), 3)
	Equal(t, *test.IntPtrArray[0], int(1))
	Equal(t, test.IntPtrArray[1], nil)
	Equal(t, *test.IntPtrArray[2], int(3))

	Equal(t, len(test.IntArrayArray), 3)
	Equal(t, len(test.IntArrayArray[0]), 3)
	Equal(t, len(test.IntArrayArray[1]), 0)
	Equal(t, len(test.IntArrayArray[2]), 1)
	Equal(t, test.IntArrayArray[0][0], int(1))
	Equal(t, test.IntArrayArray[0][1], int(0))
	Equal(t, test.IntArrayArray[0][2], int(3))
	Equal(t, test.IntArrayArray[2][0], int(1))

	Equal(t, len(test.IntPtrArrayArray), 3)
	Equal(t, len(test.IntPtrArrayArray[0]), 3)
	Equal(t, len(test.IntPtrArrayArray[1]), 0)
	Equal(t, len(test.IntPtrArrayArray[2]), 1)
	Equal(t, *test.IntPtrArrayArray[0][0], int(1))
	Equal(t, test.IntPtrArrayArray[0][1], nil)
	Equal(t, *test.IntPtrArrayArray[0][2], int(3))
	Equal(t, *test.IntPtrArrayArray[2][0], int(1))

	Equal(t, len(test.IntMap), 1)
	Equal(t, len(test.IntPtrMap), 1)

	v, ok := test.IntMap[1]
	Equal(t, ok, true)
	Equal(t, v, int(3))

	Equal(t, test.NoURLValue, int(0))

	Equal(t, test.IntNoValues, int(0))
	Equal(t, test.Int8NoValues, int8(0))
	Equal(t, test.Int16NoValues, int16(0))
	Equal(t, test.Int32NoValues, int32(0))
	Equal(t, test.Int64NoValues, int64(0))
}

func TestDecoderUint(t *testing.T) {

	type TestUint struct {
		Uint              uint
		Uint8             uint8
		Uint16            uint16
		Uint32            uint32
		Uint64            uint64
		UintPtr           *uint
		Uint8Ptr          *uint8
		Uint16Ptr         *uint16
		Uint32Ptr         *uint32
		Uint64Ptr         *uint64
		UintArray         []uint
		UintPtrArray      []*uint
		UintArrayArray    [][]uint
		UintPtrArrayArray [][]*uint
		UintMap           map[uint]uint
		UintPtrMap        map[*uint]*uint
		NoURLValue        uint
		UintNoValues      uint
		Uint8NoValues     uint8
		Uint16NoValues    uint16
		Uint32NoValues    uint32
		Uint64NoValues    uint64
	}

	values := url.Values{
		"Uint":                    []string{"3"},
		"Uint8":                   []string{"3"},
		"Uint16":                  []string{"3"},
		"Uint32":                  []string{"3"},
		"Uint64":                  []string{"3"},
		"UintPtr":                 []string{"3"},
		"Uint8Ptr":                []string{"3"},
		"Uint16Ptr":               []string{"3"},
		"Uint32Ptr":               []string{"3"},
		"Uint64Ptr":               []string{"3"},
		"UintArray":               []string{"1", "2", "3"},
		"UintPtrArray[0]":         []string{"1"},
		"UintPtrArray[2]":         []string{"3"},
		"UintArrayArray[0][0]":    []string{"1"},
		"UintArrayArray[0][2]":    []string{"3"},
		"UintArrayArray[2][0]":    []string{"1"},
		"UintPtrArrayArray[0][0]": []string{"1"},
		"UintPtrArrayArray[0][2]": []string{"3"},
		"UintPtrArrayArray[2][0]": []string{"1"},
		"UintMap[1]":              []string{"3"},
		"UintPtrMap[1]":           []string{"3"},
	}

	var test TestUint

	test.UintArray = make([]uint, 4)

	decoder := NewDecoder()
	errs := decoder.Decode(&test, values)
	Equal(t, errs, nil)

	Equal(t, test.Uint, uint(3))
	Equal(t, test.Uint8, uint8(3))
	Equal(t, test.Uint16, uint16(3))
	Equal(t, test.Uint32, uint32(3))
	Equal(t, test.Uint64, uint64(3))

	Equal(t, *test.UintPtr, uint(3))
	Equal(t, *test.Uint8Ptr, uint8(3))
	Equal(t, *test.Uint16Ptr, uint16(3))
	Equal(t, *test.Uint32Ptr, uint32(3))
	Equal(t, *test.Uint64Ptr, uint64(3))

	Equal(t, len(test.UintArray), 4)
	Equal(t, test.UintArray[0], uint(1))
	Equal(t, test.UintArray[1], uint(2))
	Equal(t, test.UintArray[2], uint(3))
	Equal(t, test.UintArray[3], uint(0))

	Equal(t, len(test.UintPtrArray), 3)
	Equal(t, *test.UintPtrArray[0], uint(1))
	Equal(t, test.UintPtrArray[1], nil)
	Equal(t, *test.UintPtrArray[2], uint(3))

	Equal(t, len(test.UintArrayArray), 3)
	Equal(t, len(test.UintArrayArray[0]), 3)
	Equal(t, len(test.UintArrayArray[1]), 0)
	Equal(t, len(test.UintArrayArray[2]), 1)
	Equal(t, test.UintArrayArray[0][0], uint(1))
	Equal(t, test.UintArrayArray[0][1], uint(0))
	Equal(t, test.UintArrayArray[0][2], uint(3))
	Equal(t, test.UintArrayArray[2][0], uint(1))

	Equal(t, len(test.UintPtrArrayArray), 3)
	Equal(t, len(test.UintPtrArrayArray[0]), 3)
	Equal(t, len(test.UintPtrArrayArray[1]), 0)
	Equal(t, len(test.UintPtrArrayArray[2]), 1)
	Equal(t, *test.UintPtrArrayArray[0][0], uint(1))
	Equal(t, test.UintPtrArrayArray[0][1], nil)
	Equal(t, *test.UintPtrArrayArray[0][2], uint(3))
	Equal(t, *test.UintPtrArrayArray[2][0], uint(1))

	Equal(t, len(test.UintMap), 1)
	Equal(t, len(test.UintPtrMap), 1)

	v, ok := test.UintMap[1]
	Equal(t, ok, true)
	Equal(t, v, uint(3))

	Equal(t, test.NoURLValue, uint(0))

	Equal(t, test.UintNoValues, uint(0))
	Equal(t, test.Uint8NoValues, uint8(0))
	Equal(t, test.Uint16NoValues, uint16(0))
	Equal(t, test.Uint32NoValues, uint32(0))
	Equal(t, test.Uint64NoValues, uint64(0))
}

func TestDecoderString(t *testing.T) {

	type TestString struct {
		String              string
		StringPtr           *string
		StringArray         []string
		StringPtrArray      []*string
		StringArrayArray    [][]string
		StringPtrArrayArray [][]*string
		StringMap           map[string]string
		StringPtrMap        map[*string]*string
		NoURLValue          string
	}

	values := url.Values{
		"String":                    []string{"3"},
		"StringPtr":                 []string{"3"},
		"StringArray":               []string{"1", "2", "3"},
		"StringPtrArray[0]":         []string{"1"},
		"StringPtrArray[2]":         []string{"3"},
		"StringArrayArray[0][0]":    []string{"1"},
		"StringArrayArray[0][2]":    []string{"3"},
		"StringArrayArray[2][0]":    []string{"1"},
		"StringPtrArrayArray[0][0]": []string{"1"},
		"StringPtrArrayArray[0][2]": []string{"3"},
		"StringPtrArrayArray[2][0]": []string{"1"},
		"StringMap[1]":              []string{"3"},
		"StringPtrMap[1]":           []string{"3"},
	}

	var test TestString

	test.StringArray = make([]string, 4)

	decoder := NewDecoder()
	errs := decoder.Decode(&test, values)
	Equal(t, errs, nil)

	Equal(t, test.String, "3")

	Equal(t, *test.StringPtr, "3")

	Equal(t, len(test.StringArray), 4)
	Equal(t, test.StringArray[0], "1")
	Equal(t, test.StringArray[1], "2")
	Equal(t, test.StringArray[2], "3")
	Equal(t, test.StringArray[3], "")

	Equal(t, len(test.StringPtrArray), 3)
	Equal(t, *test.StringPtrArray[0], "1")
	Equal(t, test.StringPtrArray[1], nil)
	Equal(t, *test.StringPtrArray[2], "3")

	Equal(t, len(test.StringArrayArray), 3)
	Equal(t, len(test.StringArrayArray[0]), 3)
	Equal(t, len(test.StringArrayArray[1]), 0)
	Equal(t, len(test.StringArrayArray[2]), 1)
	Equal(t, test.StringArrayArray[0][0], "1")
	Equal(t, test.StringArrayArray[0][1], "")
	Equal(t, test.StringArrayArray[0][2], "3")
	Equal(t, test.StringArrayArray[2][0], "1")

	Equal(t, len(test.StringPtrArrayArray), 3)
	Equal(t, len(test.StringPtrArrayArray[0]), 3)
	Equal(t, len(test.StringPtrArrayArray[1]), 0)
	Equal(t, len(test.StringPtrArrayArray[2]), 1)
	Equal(t, *test.StringPtrArrayArray[0][0], "1")
	Equal(t, test.StringPtrArrayArray[0][1], nil)
	Equal(t, *test.StringPtrArrayArray[0][2], "3")
	Equal(t, *test.StringPtrArrayArray[2][0], "1")

	Equal(t, len(test.StringMap), 1)
	Equal(t, len(test.StringPtrMap), 1)

	v, ok := test.StringMap["1"]
	Equal(t, ok, true)
	Equal(t, v, "3")

	Equal(t, test.NoURLValue, "")
}

func TestDecoderFloat(t *testing.T) {

	type TestFloat struct {
		Float32              float32
		Float32Ptr           *float32
		Float64              float64
		Float64Ptr           *float64
		Float32Array         []float32
		Float32PtrArray      []*float32
		Float32ArrayArray    [][]float32
		Float32PtrArrayArray [][]*float32
		Float32Map           map[float32]float32
		Float32PtrMap        map[*float32]*float32
		Float32NoValue       float32
		Float64NoValue       float64
	}

	values := url.Values{
		"Float32":                    []string{"3.3"},
		"Float32Ptr":                 []string{"3.3"},
		"Float64":                    []string{"3.3"},
		"Float64Ptr":                 []string{"3.3"},
		"Float32Array":               []string{"1.1", "2.2", "3.3"},
		"Float32PtrArray[0]":         []string{"1.1"},
		"Float32PtrArray[2]":         []string{"3.3"},
		"Float32ArrayArray[0][0]":    []string{"1.1"},
		"Float32ArrayArray[0][2]":    []string{"3.3"},
		"Float32ArrayArray[2][0]":    []string{"1.1"},
		"Float32PtrArrayArray[0][0]": []string{"1.1"},
		"Float32PtrArrayArray[0][2]": []string{"3.3"},
		"Float32PtrArrayArray[2][0]": []string{"1.1"},
		"Float32Map[1.1]":            []string{"3.3"},
		"Float32PtrMap[1.1]":         []string{"3.3"},
	}

	var test TestFloat

	test.Float32Array = make([]float32, 4)

	decoder := NewDecoder()
	errs := decoder.Decode(&test, values)
	Equal(t, errs, nil)

	Equal(t, test.Float32, float32(3.3))
	Equal(t, test.Float64, float64(3.3))

	Equal(t, *test.Float32Ptr, float32(3.3))
	Equal(t, *test.Float64Ptr, float64(3.3))

	Equal(t, len(test.Float32Array), 4)
	Equal(t, test.Float32Array[0], float32(1.1))
	Equal(t, test.Float32Array[1], float32(2.2))
	Equal(t, test.Float32Array[2], float32(3.3))
	Equal(t, test.Float32Array[3], float32(0.0))

	Equal(t, len(test.Float32PtrArray), 3)
	Equal(t, *test.Float32PtrArray[0], float32(1.1))
	Equal(t, test.Float32PtrArray[1], nil)
	Equal(t, *test.Float32PtrArray[2], float32(3.3))

	Equal(t, len(test.Float32ArrayArray), 3)
	Equal(t, len(test.Float32ArrayArray[0]), 3)
	Equal(t, len(test.Float32ArrayArray[1]), 0)
	Equal(t, len(test.Float32ArrayArray[2]), 1)
	Equal(t, test.Float32ArrayArray[0][0], float32(1.1))
	Equal(t, test.Float32ArrayArray[0][1], float32(0.0))
	Equal(t, test.Float32ArrayArray[0][2], float32(3.3))
	Equal(t, test.Float32ArrayArray[2][0], float32(1.1))

	Equal(t, len(test.Float32PtrArrayArray), 3)
	Equal(t, len(test.Float32PtrArrayArray[0]), 3)
	Equal(t, len(test.Float32PtrArrayArray[1]), 0)
	Equal(t, len(test.Float32PtrArrayArray[2]), 1)
	Equal(t, *test.Float32PtrArrayArray[0][0], float32(1.1))
	Equal(t, test.Float32PtrArrayArray[0][1], nil)
	Equal(t, *test.Float32PtrArrayArray[0][2], float32(3.3))
	Equal(t, *test.Float32PtrArrayArray[2][0], float32(1.1))

	Equal(t, len(test.Float32Map), 1)
	Equal(t, len(test.Float32PtrMap), 1)

	v, ok := test.Float32Map[float32(1.1)]
	Equal(t, ok, true)
	Equal(t, v, float32(3.3))

	Equal(t, test.Float32NoValue, float32(0.0))
	Equal(t, test.Float64NoValue, float64(0.0))
}

func TestDecoderBool(t *testing.T) {

	type TestBool struct {
		Bool              bool
		BoolPtr           *bool
		BoolArray         []bool
		BoolPtrArray      []*bool
		BoolArrayArray    [][]bool
		BoolPtrArrayArray [][]*bool
		BoolMap           map[bool]bool
		BoolPtrMap        map[*bool]*bool
		NoURLValue        bool
	}

	values := url.Values{
		"Bool":                    []string{"true"},
		"BoolPtr":                 []string{"true"},
		"BoolArray":               []string{"off", "t", "on"},
		"BoolPtrArray[0]":         []string{"true"},
		"BoolPtrArray[2]":         []string{"T"},
		"BoolArrayArray[0][0]":    []string{"TRUE"},
		"BoolArrayArray[0][2]":    []string{"True"},
		"BoolArrayArray[2][0]":    []string{"true"},
		"BoolPtrArrayArray[0][0]": []string{"true"},
		"BoolPtrArrayArray[0][2]": []string{"t"},
		"BoolPtrArrayArray[2][0]": []string{"1"},
		"BoolMap[true]":           []string{"true"},
		"BoolPtrMap[t]":           []string{"true"},
	}

	var test TestBool

	test.BoolArray = make([]bool, 4)

	decoder := NewDecoder()
	errs := decoder.Decode(&test, values)
	Equal(t, errs, nil)

	Equal(t, test.Bool, true)

	Equal(t, *test.BoolPtr, true)

	Equal(t, len(test.BoolArray), 4)
	Equal(t, test.BoolArray[0], false)
	Equal(t, test.BoolArray[1], true)
	Equal(t, test.BoolArray[2], true)
	Equal(t, test.BoolArray[3], false)

	Equal(t, len(test.BoolPtrArray), 3)
	Equal(t, *test.BoolPtrArray[0], true)
	Equal(t, test.BoolPtrArray[1], nil)
	Equal(t, *test.BoolPtrArray[2], true)

	Equal(t, len(test.BoolArrayArray), 3)
	Equal(t, len(test.BoolArrayArray[0]), 3)
	Equal(t, len(test.BoolArrayArray[1]), 0)
	Equal(t, len(test.BoolArrayArray[2]), 1)
	Equal(t, test.BoolArrayArray[0][0], true)
	Equal(t, test.BoolArrayArray[0][1], false)
	Equal(t, test.BoolArrayArray[0][2], true)
	Equal(t, test.BoolArrayArray[2][0], true)

	Equal(t, len(test.BoolPtrArrayArray), 3)
	Equal(t, len(test.BoolPtrArrayArray[0]), 3)
	Equal(t, len(test.BoolPtrArrayArray[1]), 0)
	Equal(t, len(test.BoolPtrArrayArray[2]), 1)
	Equal(t, *test.BoolPtrArrayArray[0][0], true)
	Equal(t, test.BoolPtrArrayArray[0][1], nil)
	Equal(t, *test.BoolPtrArrayArray[0][2], true)
	Equal(t, *test.BoolPtrArrayArray[2][0], true)

	Equal(t, len(test.BoolMap), 1)
	Equal(t, len(test.BoolPtrMap), 1)

	v, ok := test.BoolMap[true]
	Equal(t, ok, true)
	Equal(t, v, true)

	Equal(t, test.NoURLValue, false)
}

func TestDecoderStruct(t *testing.T) {

	type Phone struct {
		Number string
	}

	type TestStruct struct {
		Name      string `form:"name"`
		Phone     []Phone
		PhonePtr  []*Phone
		NonNilPtr *Phone
		Ignore    string `form:"-"`
		Anonymous struct {
			Value     string
			Ignore    string `form:"-"`
			unexposed string
		}
		Time                       time.Time
		TimePtr                    *time.Time
		unexposed                  string
		Invalid                    interface{}
		ExistingMap                map[string]string `form:"mp"`
		MapNoValue                 map[int]int
		NilArray                   []string
		TooSmallArray              []string
		TooSmallCapOKArray         []string
		ZeroLengthArray            []string
		TooSmallNumberedArray      []string
		TooSmallCapOKNumberedArray []string
		BigEnoughNumberedArray     []string
		IfaceNonNil                interface{}
		IfaceInvalid               interface{}
		TimeMapKey                 map[time.Time]string
	}

	values := url.Values{
		"name":                          []string{"joeybloggs"},
		"Ignore":                        []string{"ignore"},
		"Phone[0].Number":               []string{"1(111)111-1111"},
		"Phone[1].Number":               []string{"9(999)999-9999"},
		"PhonePtr[0].Number":            []string{"1(111)111-1111"},
		"PhonePtr[1].Number":            []string{"9(999)999-9999"},
		"NonNilPtr.Number":              []string{"9(999)999-9999"},
		"Anonymous.Value":               []string{"Anon"},
		"Time":                          []string{"2016-01-02"},
		"TimePtr":                       []string{"2016-01-02"},
		"mp[key]":                       []string{"value"},
		"NilArray":                      []string{"1", "2"},
		"TooSmallArray":                 []string{"1", "2"},
		"TooSmallCapOKArray":            []string{"1", "2"},
		"ZeroLengthArray":               []string{},
		"TooSmallNumberedArray[2]":      []string{"2"},
		"TooSmallCapOKNumberedArray[2]": []string{"2"},
		"BigEnoughNumberedArray[2]":     []string{"1"},
		"TimeMapKey[2016-01-02]":        []string{"time"},
	}

	var test TestStruct
	test.ExistingMap = map[string]string{"existingkey": "existingvalue"}
	test.NonNilPtr = new(Phone)
	test.IfaceNonNil = new(Phone)
	test.IfaceInvalid = nil
	test.TooSmallArray = []string{"0"}
	test.TooSmallCapOKArray = make([]string, 0, 10)
	test.TooSmallNumberedArray = []string{"0"}
	test.TooSmallCapOKNumberedArray = make([]string, 0, 10)
	test.BigEnoughNumberedArray = make([]string, 3, 10)

	decoder := NewDecoder()
	decoder.SetTagName("form")
	decoder.RegisterCustomTypeFunc(func(vals []string) (interface{}, error) {
		return time.Parse("2006-01-02", vals[0])
	}, time.Time{})

	errs := decoder.Decode(&test, values)
	Equal(t, errs, nil)

	Equal(t, test.Name, "joeybloggs")
	Equal(t, test.Ignore, "")
	Equal(t, len(test.Phone), 2)
	Equal(t, test.Phone[0].Number, "1(111)111-1111")
	Equal(t, test.Phone[1].Number, "9(999)999-9999")
	Equal(t, len(test.PhonePtr), 2)
	Equal(t, (*test.PhonePtr[0]).Number, "1(111)111-1111")
	Equal(t, (*test.PhonePtr[1]).Number, "9(999)999-9999")
	Equal(t, test.NonNilPtr.Number, "9(999)999-9999")
	Equal(t, test.Anonymous.Value, "Anon")
	Equal(t, len(test.ExistingMap), 2)
	Equal(t, test.ExistingMap["existingkey"], "existingvalue")
	Equal(t, test.ExistingMap["key"], "value")
	Equal(t, len(test.NilArray), 2)
	Equal(t, test.NilArray[0], "1")
	Equal(t, test.NilArray[1], "2")
	Equal(t, len(test.TooSmallArray), 2)
	Equal(t, test.TooSmallArray[0], "1")
	Equal(t, test.TooSmallArray[1], "2")
	Equal(t, len(test.ZeroLengthArray), 0)
	Equal(t, len(test.TooSmallNumberedArray), 3)
	Equal(t, test.TooSmallNumberedArray[0], "0")
	Equal(t, test.TooSmallNumberedArray[1], "")
	Equal(t, test.TooSmallNumberedArray[2], "2")
	Equal(t, len(test.BigEnoughNumberedArray), 3)
	Equal(t, cap(test.BigEnoughNumberedArray), 10)
	Equal(t, test.BigEnoughNumberedArray[0], "")
	Equal(t, test.BigEnoughNumberedArray[1], "")
	Equal(t, test.BigEnoughNumberedArray[2], "1")
	Equal(t, len(test.TooSmallCapOKArray), 2)
	Equal(t, cap(test.TooSmallCapOKArray), 10)
	Equal(t, test.TooSmallCapOKArray[0], "1")
	Equal(t, test.TooSmallCapOKArray[1], "2")
	Equal(t, len(test.TooSmallCapOKNumberedArray), 3)
	Equal(t, cap(test.TooSmallCapOKNumberedArray), 10)
	Equal(t, test.TooSmallCapOKNumberedArray[0], "")
	Equal(t, test.TooSmallCapOKNumberedArray[1], "")
	Equal(t, test.TooSmallCapOKNumberedArray[2], "2")

	tm, _ := time.Parse("2006-01-02", "2016-01-02")
	Equal(t, test.Time.Equal(tm), true)
	Equal(t, (*test.TimePtr).Equal(tm), true)

	NotEqual(t, test.TimeMapKey, nil)
	Equal(t, len(test.TimeMapKey), 1)

	_, ok := test.TimeMapKey[tm]
	Equal(t, ok, true)

	s := struct {
		Value     string
		Ignore    string `form:"-"`
		unexposed string
	}{}

	errs = decoder.Decode(&s, values)
	Equal(t, errs, nil)
	Equal(t, s.Value, "")
	Equal(t, s.Ignore, "")
	Equal(t, s.unexposed, "")

}

func TestStructSlice(t *testing.T) {
	type Phone struct {
		Number string
	}

	type TestStruct struct {
		Name      string `form:"name"`
		Phone     []Phone
		PhonePtr  []*Phone
		NonNilPtr *Phone
		Ignore    string `form:"-"`
		Anonymous struct {
			Value     string
			Ignore    string `form:"-"`
			unexposed string
		}
		Time                       time.Time
		TimePtr                    *time.Time
		unexposed                  string
		Invalid                    interface{}
		ExistingMap                map[string]string `form:"mp"`
		MapNoValue                 map[int]int
		NilArray                   []string
		TooSmallArray              []string
		TooSmallCapOKArray         []string
		ZeroLengthArray            []string
		TooSmallNumberedArray      []string
		TooSmallCapOKNumberedArray []string
		BigEnoughNumberedArray     []string
		IfaceNonNil                interface{}
		IfaceInvalid               interface{}
		TimeMapKey                 map[time.Time]string
	}

	values := url.Values{
		"[0].name":                          []string{"joeybloggs"},
		"[0].Ignore":                        []string{"ignore"},
		"[0].Phone[0].Number":               []string{"1(111)111-1111"},
		"[0].Phone[1].Number":               []string{"9(999)999-9999"},
		"[0].PhonePtr[0].Number":            []string{"1(111)111-1111"},
		"[0].PhonePtr[1].Number":            []string{"9(999)999-9999"},
		"[0].NonNilPtr.Number":              []string{"9(999)999-9999"},
		"[0].Anonymous.Value":               []string{"Anon"},
		"[0].Time":                          []string{"2016-01-02"},
		"[0].TimePtr":                       []string{"2016-01-02"},
		"[0].mp[key]":                       []string{"value"},
		"[0].NilArray":                      []string{"1", "2"},
		"[0].TooSmallArray":                 []string{"1", "2"},
		"[0].TooSmallCapOKArray":            []string{"1", "2"},
		"[0].ZeroLengthArray":               []string{},
		"[0].TooSmallNumberedArray[2]":      []string{"2"},
		"[0].TooSmallCapOKNumberedArray[2]": []string{"2"},
		"[0].BigEnoughNumberedArray[2]":     []string{"1"},
		"[0].TimeMapKey[2016-01-02]":        []string{"time"},

		"[1].name":                          []string{"joeybloggs1"},
		"[1].Ignore":                        []string{"ignore1"},
		"[1].Phone[0].Number":               []string{"2(222)222-2222"},
		"[1].Phone[1].Number":               []string{"3(333)333-3333"},
		"[1].PhonePtr[0].Number":            []string{"2(222)222-2222"},
		"[1].PhonePtr[1].Number":            []string{"3(333)333-3333"},
		"[1].NonNilPtr.Number":              []string{"3(333)333-3333"},
		"[1].Anonymous.Value":               []string{"Anon1"},
		"[1].Time":                          []string{"2016-01-03"},
		"[1].TimePtr":                       []string{"2016-01-03"},
		"[1].mp[key1]":                      []string{"value1"},
		"[1].NilArray":                      []string{"3", "4"},
		"[1].TooSmallArray":                 []string{"3", "4"},
		"[1].TooSmallCapOKArray":            []string{"3", "4"},
		"[1].ZeroLengthArray":               []string{},
		"[1].TooSmallNumberedArray[3]":      []string{"3"},
		"[1].TooSmallCapOKNumberedArray[3]": []string{"3"},
		"[1].BigEnoughNumberedArray[2]":     []string{"2"},
		"[1].TimeMapKey[2016-01-03]":        []string{"time"},
	}

	test := []TestStruct{
		TestStruct{
			ExistingMap:                map[string]string{"existingkey": "existingvalue"},
			NonNilPtr:                  new(Phone),
			IfaceNonNil:                new(Phone),
			IfaceInvalid:               nil,
			TooSmallArray:              []string{"0"},
			TooSmallCapOKArray:         make([]string, 0, 10),
			TooSmallNumberedArray:      []string{"0"},
			TooSmallCapOKNumberedArray: make([]string, 0, 10),
			BigEnoughNumberedArray:     make([]string, 3, 10),
		},
		TestStruct{
			ExistingMap:                map[string]string{"existingkey1": "existingvalue1"},
			NonNilPtr:                  new(Phone),
			IfaceNonNil:                new(Phone),
			IfaceInvalid:               nil,
			TooSmallArray:              []string{"0"},
			TooSmallCapOKArray:         make([]string, 0, 10),
			TooSmallNumberedArray:      []string{"0"},
			TooSmallCapOKNumberedArray: make([]string, 0, 10),
			BigEnoughNumberedArray:     make([]string, 3, 10),
		},
	}
	decoder := NewDecoder()
	decoder.SetTagName("form")
	decoder.RegisterCustomTypeFunc(func(vals []string) (interface{}, error) {
		return time.Parse("2006-01-02", vals[0])
	}, time.Time{})

	errs := decoder.Decode(&test, values)
	Equal(t, errs, nil)

	Equal(t, test[0].Name, "joeybloggs")
	Equal(t, test[0].Ignore, "")
	Equal(t, len(test[0].Phone), 2)
	Equal(t, test[0].Phone[0].Number, "1(111)111-1111")
	Equal(t, test[0].Phone[1].Number, "9(999)999-9999")
	Equal(t, len(test[0].PhonePtr), 2)
	Equal(t, (*test[0].PhonePtr[0]).Number, "1(111)111-1111")
	Equal(t, (*test[0].PhonePtr[1]).Number, "9(999)999-9999")
	Equal(t, test[0].NonNilPtr.Number, "9(999)999-9999")
	Equal(t, test[0].Anonymous.Value, "Anon")
	Equal(t, len(test[0].ExistingMap), 1)
	//	Equal(t, test[0].ExistingMap["existingkey"], "existingvalue")
	Equal(t, test[0].ExistingMap["key"], "value")
	Equal(t, len(test[0].NilArray), 2)
	Equal(t, test[0].NilArray[0], "1")
	Equal(t, test[0].NilArray[1], "2")
	Equal(t, len(test[0].TooSmallArray), 2)
	Equal(t, test[0].TooSmallArray[0], "1")
	Equal(t, test[0].TooSmallArray[1], "2")
	Equal(t, len(test[0].ZeroLengthArray), 0)
	Equal(t, len(test[0].TooSmallNumberedArray), 3)
	//	Equal(t, test[0].TooSmallNumberedArray[0], "0")
	Equal(t, test[0].TooSmallNumberedArray[1], "")
	Equal(t, test[0].TooSmallNumberedArray[2], "2")
	Equal(t, len(test[0].BigEnoughNumberedArray), 3)
	//	Equal(t, cap(test[0].BigEnoughNumberedArray), 10)
	Equal(t, test[0].BigEnoughNumberedArray[0], "")
	Equal(t, test[0].BigEnoughNumberedArray[1], "")
	Equal(t, test[0].BigEnoughNumberedArray[2], "1")
	Equal(t, len(test[0].TooSmallCapOKArray), 2)
	//	Equal(t, cap(test[0].TooSmallCapOKArray), 10)
	Equal(t, test[0].TooSmallCapOKArray[0], "1")
	Equal(t, test[0].TooSmallCapOKArray[1], "2")
	Equal(t, len(test[0].TooSmallCapOKNumberedArray), 3)
	//	Equal(t, cap(test[0].TooSmallCapOKNumberedArray), 10)
	Equal(t, test[0].TooSmallCapOKNumberedArray[0], "")
	Equal(t, test[0].TooSmallCapOKNumberedArray[1], "")
	Equal(t, test[0].TooSmallCapOKNumberedArray[2], "2")

	tm, _ := time.Parse("2006-01-02", "2016-01-02")
	Equal(t, test[0].Time.Equal(tm), true)
	Equal(t, (*test[0].TimePtr).Equal(tm), true)

	NotEqual(t, test[0].TimeMapKey, nil)
	Equal(t, len(test[0].TimeMapKey), 1)

	_, ok := test[0].TimeMapKey[tm]
	Equal(t, ok, true)

	s := struct {
		Value     string
		Ignore    string `form:"-"`
		unexposed string
	}{}

	errs = decoder.Decode(&s, values)
	Equal(t, errs, nil)
	Equal(t, s.Value, "")
	Equal(t, s.Ignore, "")
	Equal(t, s.unexposed, "")

	Equal(t, test[1].Name, "joeybloggs1")
	Equal(t, test[1].Ignore, "")
	Equal(t, len(test[1].Phone), 2)
	Equal(t, test[1].Phone[0].Number, "2(222)222-2222")
	Equal(t, test[1].Phone[1].Number, "3(333)333-3333")
	Equal(t, len(test[1].PhonePtr), 2)
	Equal(t, (*test[1].PhonePtr[0]).Number, "2(222)222-2222")
	Equal(t, (*test[1].PhonePtr[1]).Number, "3(333)333-3333")
	Equal(t, test[1].NonNilPtr.Number, "3(333)333-3333")
	Equal(t, test[1].Anonymous.Value, "Anon1")
	Equal(t, len(test[1].ExistingMap), 1)
	//	Equal(t, test[1].ExistingMap["existingkey1"], "existingvalue1")
	Equal(t, test[1].ExistingMap["key1"], "value1")
	Equal(t, len(test[1].NilArray), 2)
	Equal(t, test[1].NilArray[0], "3")
	Equal(t, test[1].NilArray[1], "4")
	Equal(t, len(test[1].TooSmallArray), 2)
	Equal(t, test[1].TooSmallArray[0], "3")
	Equal(t, test[1].TooSmallArray[1], "4")
	Equal(t, len(test[1].ZeroLengthArray), 0)
	Equal(t, len(test[1].TooSmallNumberedArray), 3)
	//	Equal(t, test[1].TooSmallNumberedArray[0], "0")
	Equal(t, test[1].TooSmallNumberedArray[1], "")
	Equal(t, test[1].TooSmallNumberedArray[2], "3")
	Equal(t, len(test[1].BigEnoughNumberedArray), 3)
	//	Equal(t, cap(test[1].BigEnoughNumberedArray), 10)
	Equal(t, test[1].BigEnoughNumberedArray[0], "")
	Equal(t, test[1].BigEnoughNumberedArray[1], "")
	Equal(t, test[1].BigEnoughNumberedArray[2], "2")
	Equal(t, len(test[1].TooSmallCapOKArray), 2)
	//	Equal(t, cap(test[1].TooSmallCapOKArray), 10)
	Equal(t, test[1].TooSmallCapOKArray[0], "3")
	Equal(t, test[1].TooSmallCapOKArray[1], "4")
	Equal(t, len(test[1].TooSmallCapOKNumberedArray), 3)
	//	Equal(t, cap(test[1].TooSmallCapOKNumberedArray), 10)
	Equal(t, test[1].TooSmallCapOKNumberedArray[0], "")
	Equal(t, test[1].TooSmallCapOKNumberedArray[1], "")
	Equal(t, test[1].TooSmallCapOKNumberedArray[2], "3")

	tm, _ = time.Parse("2006-01-02", "2016-01-03")
	Equal(t, test[1].Time.Equal(tm), true)
	Equal(t, (*test[1].TimePtr).Equal(tm), true)

	NotEqual(t, test[1].TimeMapKey, nil)
	Equal(t, len(test[1].TimeMapKey), 1)

	_, ok = test[1].TimeMapKey[tm]
	Equal(t, ok, true)

}

func TestDecoderNativeTime(t *testing.T) {

	type TestError struct {
		Time        time.Time
		TimeNoValue time.Time
		TimePtr     *time.Time
	}

	values := url.Values{
		"Time":        []string{"2006-01-02T15:04:05Z"},
		"TimeNoValue": []string{""},
		"TimePtr":     []string{"2006-01-02T15:04:05Z"},
	}

	var test TestError

	decoder := NewDecoder()

	errs := decoder.Decode(&test, values)
	Equal(t, errs, nil)

	tm, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	Equal(t, test.Time.Equal(tm), true)
	Equal(t, test.TimeNoValue.Equal(tm), false)

	NotEqual(t, test.TimePtr, nil)
	Equal(t, (*test.TimePtr).Equal(tm), true)
}

func TestDecoderErrors(t *testing.T) {

	type TestError struct {
		Bool                  bool `form:"bool"`
		Int                   int
		Int8                  int8
		Int16                 int16
		Int32                 int32
		Uint                  uint
		Uint8                 uint8
		Uint16                uint16
		Uint32                uint32
		Float32               float32
		Float64               float64
		String                string
		Time                  time.Time
		MapBadIntKey          map[int]int
		MapBadInt8Key         map[int8]int8
		MapBadInt16Key        map[int16]int16
		MapBadInt32Key        map[int32]int32
		MapBadUintKey         map[uint]uint
		MapBadUint8Key        map[uint8]uint8
		MapBadUint16Key       map[uint16]uint16
		MapBadUint32Key       map[uint32]uint32
		MapBadFloat32Key      map[float32]float32
		MapBadFloat64Key      map[float64]float64
		MapBadBoolKey         map[bool]bool
		MapBadKeyType         map[complex64]int
		BadArrayValue         []int
		BadMapKey             map[time.Time]string
		OverflowNilArray      []int
		OverFlowExistingArray []int
		BadArrayIndex         []int
	}

	values := url.Values{
		"bool":                       []string{"uh-huh"},
		"Int":                        []string{"bad"},
		"Int8":                       []string{"bad"},
		"Int16":                      []string{"bad"},
		"Int32":                      []string{"bad"},
		"Uint":                       []string{"bad"},
		"Uint8":                      []string{"bad"},
		"Uint16":                     []string{"bad"},
		"Uint32":                     []string{"bad"},
		"Float32":                    []string{"bad"},
		"Float64":                    []string{"bad"},
		"String":                     []string{"str bad return val"},
		"Time":                       []string{"bad"},
		"MapBadIntKey[key]":          []string{"1"},
		"MapBadInt8Key[key]":         []string{"1"},
		"MapBadInt16Key[key]":        []string{"1"},
		"MapBadInt32Key[key]":        []string{"1"},
		"MapBadUintKey[key]":         []string{"1"},
		"MapBadUint8Key[key]":        []string{"1"},
		"MapBadUint16Key[key]":       []string{"1"},
		"MapBadUint32Key[key]":       []string{"1"},
		"MapBadFloat32Key[key]":      []string{"1.1"},
		"MapBadFloat64Key[key]":      []string{"1.1"},
		"MapBadBoolKey[uh-huh]":      []string{"true"},
		"MapBadKeyType[1.4]":         []string{"5"},
		"BadArrayValue[0]":           []string{"badintval"},
		"BadMapKey[badtime]":         []string{"badtime"},
		"OverflowNilArray[999]":      []string{"idx 1000"},
		"OverFlowExistingArray[999]": []string{"idx 1000"},
		"BadArrayIndex[bad index]":   []string{"bad idx"},
	}

	test := TestError{
		OverFlowExistingArray: make([]int, 2, 2),
	}

	decoder := NewDecoder()
	decoder.SetMaxArraySize(4)
	decoder.RegisterCustomTypeFunc(func(vals []string) (interface{}, error) {
		return nil, errors.New("Bad Type Conversion")
	}, "")

	errs := decoder.Decode(&test, values)
	NotEqual(t, errs, nil)

	e := errs.Error()
	NotEqual(t, e, "")

	err := errs.(DecodeErrors)
	Equal(t, len(err), 30)

	k := err["bool"]
	Equal(t, k.Error(), "Invalid Boolean Value 'uh-huh' Type 'bool' Namespace 'bool'")

	k = err["Int"]
	Equal(t, k.Error(), "Invalid Integer Value 'bad' Type 'int' Namespace 'Int'")

	k = err["Int8"]
	Equal(t, k.Error(), "Invalid Integer Value 'bad' Type 'int8' Namespace 'Int8'")

	k = err["Int16"]
	Equal(t, k.Error(), "Invalid Integer Value 'bad' Type 'int16' Namespace 'Int16'")

	k = err["Int32"]
	Equal(t, k.Error(), "Invalid Integer Value 'bad' Type 'int32' Namespace 'Int32'")

	k = err["Uint"]
	Equal(t, k.Error(), "Invalid Unsigned Integer Value 'bad' Type 'uint' Namespace 'Uint'")

	k = err["Uint8"]
	Equal(t, k.Error(), "Invalid Unsigned Integer Value 'bad' Type 'uint8' Namespace 'Uint8'")

	k = err["Uint16"]
	Equal(t, k.Error(), "Invalid Unsigned Integer Value 'bad' Type 'uint16' Namespace 'Uint16'")

	k = err["Uint32"]
	Equal(t, k.Error(), "Invalid Unsigned Integer Value 'bad' Type 'uint32' Namespace 'Uint32'")

	k = err["Float32"]
	Equal(t, k.Error(), "Invalid Float Value 'bad' Type 'float32' Namespace 'Float32'")

	k = err["Float64"]
	Equal(t, k.Error(), "Invalid Float Value 'bad' Type 'float64' Namespace 'Float64'")

	k = err["String"]
	Equal(t, k.Error(), "Bad Type Conversion")

	k = err["Time"]
	Equal(t, k.Error(), "parsing time \"bad\" as \"2006-01-02T15:04:05Z07:00\": cannot parse \"bad\" as \"2006\"")

	k = err["MapBadIntKey"]
	Equal(t, k.Error(), "Invalid Integer Value 'key' Type 'int' Namespace 'MapBadIntKey'")

	k = err["MapBadInt8Key"]
	Equal(t, k.Error(), "Invalid Integer Value 'key' Type 'int8' Namespace 'MapBadInt8Key'")

	k = err["MapBadInt16Key"]
	Equal(t, k.Error(), "Invalid Integer Value 'key' Type 'int16' Namespace 'MapBadInt16Key'")

	k = err["MapBadInt32Key"]
	Equal(t, k.Error(), "Invalid Integer Value 'key' Type 'int32' Namespace 'MapBadInt32Key'")

	k = err["MapBadUintKey"]
	Equal(t, k.Error(), "Invalid Unsigned Integer Value 'key' Type 'uint' Namespace 'MapBadUintKey'")

	k = err["MapBadUint8Key"]
	Equal(t, k.Error(), "Invalid Unsigned Integer Value 'key' Type 'uint8' Namespace 'MapBadUint8Key'")

	k = err["MapBadUint16Key"]
	Equal(t, k.Error(), "Invalid Unsigned Integer Value 'key' Type 'uint16' Namespace 'MapBadUint16Key'")

	k = err["MapBadUint32Key"]
	Equal(t, k.Error(), "Invalid Unsigned Integer Value 'key' Type 'uint32' Namespace 'MapBadUint32Key'")

	k = err["MapBadFloat32Key"]
	Equal(t, k.Error(), "Invalid Float Value 'key' Type 'float32' Namespace 'MapBadFloat32Key'")

	k = err["MapBadFloat64Key"]
	Equal(t, k.Error(), "Invalid Float Value 'key' Type 'float64' Namespace 'MapBadFloat64Key'")

	k = err["MapBadBoolKey"]
	Equal(t, k.Error(), "Invalid Boolean Value 'uh-huh' Type 'bool' Namespace 'MapBadBoolKey'")

	k = err["MapBadKeyType"]
	Equal(t, k.Error(), "Unsupported Map Key '1.4', Type 'complex64' Namespace 'MapBadKeyType'")

	k = err["BadArrayValue[0]"]
	Equal(t, k.Error(), "Invalid Integer Value 'badintval' Type 'int' Namespace 'BadArrayValue[0]'")

	k = err["OverflowNilArray"]
	Equal(t, k.Error(), "Array size of '1000' is larger than the maximum currently set on the decoder of '4'. To increase this limit please see, SetMaxArraySize(size uint)")

	k = err["OverFlowExistingArray"]
	Equal(t, k.Error(), "Array size of '1000' is larger than the maximum currently set on the decoder of '4'. To increase this limit please see, SetMaxArraySize(size uint)")

	k = err["BadArrayIndex"]
	Equal(t, k.Error(), "Invalid Array index 'bad index'")

	type TestError2 struct {
		BadMapKey map[time.Time]string
	}

	values2 := url.Values{
		"BadMapKey[badtime]": []string{"badtime"},
	}

	var test2 TestError2
	decoder2 := NewDecoder()
	decoder2.RegisterCustomTypeFunc(func(vals []string) (interface{}, error) {
		return time.Parse("2006-01-02", vals[0])
	}, time.Time{})

	errs = decoder2.Decode(&test2, values2)
	NotEqual(t, errs, nil)

	e = errs.Error()
	NotEqual(t, e, "")

	k = err["BadMapKey"]
	Equal(t, k.Error(), "Unsupported Map Key 'badtime', Type 'time.Time' Namespace 'BadMapKey'")
}

/*
func TestDecoderPanics(t *testing.T) {

	type Phone struct {
		Number string
	}

	type TestError struct {
		Phone  []Phone
		Phone2 []Phone
		Phone3 []Phone
	}

	values := url.Values{
		"Phone[0.Number": []string{"1(111)111-1111"},
	}

	var test TestError

	decoder := NewDecoder()

	PanicMatches(t, func() { decoder.Decode(&test, values) }, "Invalid formatting for key 'Phone[0.Number' missing ']' bracket")

//	i := 1
//	PanicMatches(t, func() { decoder.Decode(&i, values) }, "interface value cannot be addressed")

	values = url.Values{
		"Phone0].Number": []string{"1(111)111-1111"},
	}

	PanicMatches(t, func() { decoder.Decode(&test, values) }, "Invalid formatting for key 'Phone0].Number' missing '[' bracket")

	values = url.Values{
		"Phone[[0.Number": []string{"1(111)111-1111"},
	}

	PanicMatches(t, func() { decoder.Decode(&test, values) }, "Invalid formatting for key 'Phone[[0.Number' missing ']' bracket")

	values = url.Values{
		"Phone0]].Number": []string{"1(111)111-1111"},
	}

	PanicMatches(t, func() { decoder.Decode(&test, values) }, "Invalid formatting for key 'Phone0]].Number' missing '[' bracket")
}

*/
func TestDecoderMapKeys(t *testing.T) {

	type TestMapKeys struct {
		MapIfaceKey   map[interface{}]string
		MapFloat32Key map[float32]float32
		MapFloat64Key map[float64]float64
		MapNestedInt  map[int]map[int]int
		MapInt8       map[int8]int8
		MapInt16      map[int16]int16
		MapInt32      map[int32]int32
		MapUint8      map[uint8]uint8
		MapUint16     map[uint16]uint16
		MapUint32     map[uint32]uint32
	}

	values := url.Values{
		"MapIfaceKey[key]":   []string{"3"},
		"MapFloat32Key[0.0]": []string{"3.3"},
		"MapFloat64Key[0.0]": []string{"3.3"},
		"MapNestedInt[1][2]": []string{"3"},
		"MapInt8[0]":         []string{"3"},
		"MapInt16[0]":        []string{"3"},
		"MapInt32[0]":        []string{"3"},
		"MapUint8[0]":        []string{"3"},
		"MapUint16[0]":       []string{"3"},
		"MapUint32[0]":       []string{"3"},
	}

	var test TestMapKeys

	decoder := NewDecoder()
	errs := decoder.Decode(&test, values)
	Equal(t, errs, nil)
	Equal(t, test.MapIfaceKey["key"], "3")
	Equal(t, test.MapFloat32Key[float32(0.0)], float32(3.3))
	Equal(t, test.MapFloat64Key[float64(0.0)], float64(3.3))

	Equal(t, test.MapInt8[int8(0)], int8(3))
	Equal(t, test.MapInt16[int16(0)], int16(3))
	Equal(t, test.MapInt32[int32(0)], int32(3))

	Equal(t, test.MapUint8[uint8(0)], uint8(3))
	Equal(t, test.MapUint16[uint16(0)], uint16(3))
	Equal(t, test.MapUint32[uint32(0)], uint32(3))

	Equal(t, len(test.MapNestedInt), 1)
	Equal(t, len(test.MapNestedInt[1]), 1)
	Equal(t, test.MapNestedInt[1][2], 3)
}

func TestDecoderStructRecursion(t *testing.T) {

	type Nested struct {
		Value  string
		Nested *Nested
	}

	type TestRecursive struct {
		Nested Nested
	}

	values := url.Values{
		"Value":        []string{"value"},
		"Nested.Value": []string{"value"},
	}

	var test TestRecursive

	decoder := NewDecoder()
	errs := decoder.Decode(&test, values)
	Equal(t, errs, nil)
	Equal(t, test.Nested.Value, "value")
	Equal(t, test.Nested.Nested, nil)
}

func TestDecoderFormDecode(t *testing.T) {

	type Struct2 struct {
		Foo string
		Bar string
	}

	type Struct2Wrapper struct {
		InnerSlice []Struct2
	}

	sliceValues := map[string][]string{
		"InnerSlice[0].Foo": {"foo-is-set"},
	}

	singleValues := map[string][]string{
		"Foo": {"foo-is-set"},
	}

	fd := NewDecoder()

	dst := Struct2Wrapper{}
	err := fd.Decode(&dst, sliceValues)
	Equal(t, err, nil)
	NotEqual(t, dst.InnerSlice, nil)
	Equal(t, dst.InnerSlice[0].Foo, "foo-is-set")

	dst2 := Struct2{}
	err = fd.Decode(&dst2, singleValues)
	Equal(t, err, nil)
	Equal(t, dst2.Foo, "foo-is-set")
}

func TestDecoderArrayKeysSort(t *testing.T) {

	type Struct struct {
		Array []int
	}

	values := map[string][]string{

		"Array[2]":  {"2"},
		"Array[10]": {"10"},
	}

	var test Struct

	d := NewDecoder()

	err := d.Decode(&test, values)
	Equal(t, err, nil)

	Equal(t, len(test.Array), 11)
	Equal(t, test.Array[2], int(2))
	Equal(t, test.Array[10], int(10))
}

func TestDecoderIncreasingKeys(t *testing.T) {

	type Struct struct {
		Array []int
	}

	values := map[string][]string{
		"Array[2]": {"2"},
	}

	var test Struct

	d := NewDecoder()

	err := d.Decode(&test, values)
	Equal(t, err, nil)

	Equal(t, len(test.Array), 3)
	Equal(t, test.Array[2], int(2))

	values["Array[10]"] = []string{"10"}

	var test2 Struct

	err = d.Decode(&test2, values)
	Equal(t, err, nil)

	Equal(t, len(test2.Array), 11)
	Equal(t, test2.Array[2], int(2))
	Equal(t, test2.Array[10], int(10))
}
