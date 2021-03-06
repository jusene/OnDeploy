// Copyright (c) 2012-2020 Ugorji Nwoke. All rights reserved.
// Use of this source code is governed by a MIT license found in the LICENSE file.

// Code generated from mammoth-test.go.tmpl - DO NOT EDIT.

package codec

import "testing"
import "fmt"

// TestMammoth has all the different paths optimized in fast-path
// It has all the primitives, slices and maps.
//
// For each of those types, it has a pointer and a non-pointer field.

func init() { _ = fmt.Printf } // so we can include fmt as needed

type TestMammoth struct {
	FIntf       interface{}
	FptrIntf    *interface{}
	FString     string
	FptrString  *string
	FBytes      []byte
	FptrBytes   *[]byte
	FFloat32    float32
	FptrFloat32 *float32
	FFloat64    float64
	FptrFloat64 *float64
	FUint       uint
	FptrUint    *uint
	FUint8      uint8
	FptrUint8   *uint8
	FUint16     uint16
	FptrUint16  *uint16
	FUint32     uint32
	FptrUint32  *uint32
	FUint64     uint64
	FptrUint64  *uint64
	FUintptr    uintptr
	FptrUintptr *uintptr
	FInt        int
	FptrInt     *int
	FInt8       int8
	FptrInt8    *int8
	FInt16      int16
	FptrInt16   *int16
	FInt32      int32
	FptrInt32   *int32
	FInt64      int64
	FptrInt64   *int64
	FBool       bool
	FptrBool    *bool

	FSliceIntf       []interface{}
	FptrSliceIntf    *[]interface{}
	FSliceString     []string
	FptrSliceString  *[]string
	FSliceBytes      [][]byte
	FptrSliceBytes   *[][]byte
	FSliceFloat64    []float64
	FptrSliceFloat64 *[]float64
	FSliceUint64     []uint64
	FptrSliceUint64  *[]uint64
	FSliceInt        []int
	FptrSliceInt     *[]int
	FSliceInt32      []int32
	FptrSliceInt32   *[]int32
	FSliceInt64      []int64
	FptrSliceInt64   *[]int64
	FSliceBool       []bool
	FptrSliceBool    *[]bool

	FMapStringIntf       map[string]interface{}
	FptrMapStringIntf    *map[string]interface{}
	FMapStringString     map[string]string
	FptrMapStringString  *map[string]string
	FMapStringBytes      map[string][]byte
	FptrMapStringBytes   *map[string][]byte
	FMapStringUint8      map[string]uint8
	FptrMapStringUint8   *map[string]uint8
	FMapStringUint64     map[string]uint64
	FptrMapStringUint64  *map[string]uint64
	FMapStringInt        map[string]int
	FptrMapStringInt     *map[string]int
	FMapStringInt64      map[string]int64
	FptrMapStringInt64   *map[string]int64
	FMapStringFloat64    map[string]float64
	FptrMapStringFloat64 *map[string]float64
	FMapStringBool       map[string]bool
	FptrMapStringBool    *map[string]bool
	FMapUint8Intf        map[uint8]interface{}
	FptrMapUint8Intf     *map[uint8]interface{}
	FMapUint8String      map[uint8]string
	FptrMapUint8String   *map[uint8]string
	FMapUint8Bytes       map[uint8][]byte
	FptrMapUint8Bytes    *map[uint8][]byte
	FMapUint8Uint8       map[uint8]uint8
	FptrMapUint8Uint8    *map[uint8]uint8
	FMapUint8Uint64      map[uint8]uint64
	FptrMapUint8Uint64   *map[uint8]uint64
	FMapUint8Int         map[uint8]int
	FptrMapUint8Int      *map[uint8]int
	FMapUint8Int64       map[uint8]int64
	FptrMapUint8Int64    *map[uint8]int64
	FMapUint8Float64     map[uint8]float64
	FptrMapUint8Float64  *map[uint8]float64
	FMapUint8Bool        map[uint8]bool
	FptrMapUint8Bool     *map[uint8]bool
	FMapUint64Intf       map[uint64]interface{}
	FptrMapUint64Intf    *map[uint64]interface{}
	FMapUint64String     map[uint64]string
	FptrMapUint64String  *map[uint64]string
	FMapUint64Bytes      map[uint64][]byte
	FptrMapUint64Bytes   *map[uint64][]byte
	FMapUint64Uint8      map[uint64]uint8
	FptrMapUint64Uint8   *map[uint64]uint8
	FMapUint64Uint64     map[uint64]uint64
	FptrMapUint64Uint64  *map[uint64]uint64
	FMapUint64Int        map[uint64]int
	FptrMapUint64Int     *map[uint64]int
	FMapUint64Int64      map[uint64]int64
	FptrMapUint64Int64   *map[uint64]int64
	FMapUint64Float64    map[uint64]float64
	FptrMapUint64Float64 *map[uint64]float64
	FMapUint64Bool       map[uint64]bool
	FptrMapUint64Bool    *map[uint64]bool
	FMapIntIntf          map[int]interface{}
	FptrMapIntIntf       *map[int]interface{}
	FMapIntString        map[int]string
	FptrMapIntString     *map[int]string
	FMapIntBytes         map[int][]byte
	FptrMapIntBytes      *map[int][]byte
	FMapIntUint8         map[int]uint8
	FptrMapIntUint8      *map[int]uint8
	FMapIntUint64        map[int]uint64
	FptrMapIntUint64     *map[int]uint64
	FMapIntInt           map[int]int
	FptrMapIntInt        *map[int]int
	FMapIntInt64         map[int]int64
	FptrMapIntInt64      *map[int]int64
	FMapIntFloat64       map[int]float64
	FptrMapIntFloat64    *map[int]float64
	FMapIntBool          map[int]bool
	FptrMapIntBool       *map[int]bool
	FMapInt64Intf        map[int64]interface{}
	FptrMapInt64Intf     *map[int64]interface{}
	FMapInt64String      map[int64]string
	FptrMapInt64String   *map[int64]string
	FMapInt64Bytes       map[int64][]byte
	FptrMapInt64Bytes    *map[int64][]byte
	FMapInt64Uint8       map[int64]uint8
	FptrMapInt64Uint8    *map[int64]uint8
	FMapInt64Uint64      map[int64]uint64
	FptrMapInt64Uint64   *map[int64]uint64
	FMapInt64Int         map[int64]int
	FptrMapInt64Int      *map[int64]int
	FMapInt64Int64       map[int64]int64
	FptrMapInt64Int64    *map[int64]int64
	FMapInt64Float64     map[int64]float64
	FptrMapInt64Float64  *map[int64]float64
	FMapInt64Bool        map[int64]bool
	FptrMapInt64Bool     *map[int64]bool
}

type typMbsSliceIntf []interface{}

func (_ typMbsSliceIntf) MapBySlice() {}

type typMbsSliceString []string

func (_ typMbsSliceString) MapBySlice() {}

type typMbsSliceBytes [][]byte

func (_ typMbsSliceBytes) MapBySlice() {}

type typMbsSliceFloat64 []float64

func (_ typMbsSliceFloat64) MapBySlice() {}

type typMbsSliceUint64 []uint64

func (_ typMbsSliceUint64) MapBySlice() {}

type typMbsSliceInt []int

func (_ typMbsSliceInt) MapBySlice() {}

type typMbsSliceInt32 []int32

func (_ typMbsSliceInt32) MapBySlice() {}

type typMbsSliceInt64 []int64

func (_ typMbsSliceInt64) MapBySlice() {}

type typMbsSliceBool []bool

func (_ typMbsSliceBool) MapBySlice() {}

type typMapMapStringIntf map[string]interface{}
type typMapMapStringString map[string]string
type typMapMapStringBytes map[string][]byte
type typMapMapStringUint8 map[string]uint8
type typMapMapStringUint64 map[string]uint64
type typMapMapStringInt map[string]int
type typMapMapStringInt64 map[string]int64
type typMapMapStringFloat64 map[string]float64
type typMapMapStringBool map[string]bool
type typMapMapUint8Intf map[uint8]interface{}
type typMapMapUint8String map[uint8]string
type typMapMapUint8Bytes map[uint8][]byte
type typMapMapUint8Uint8 map[uint8]uint8
type typMapMapUint8Uint64 map[uint8]uint64
type typMapMapUint8Int map[uint8]int
type typMapMapUint8Int64 map[uint8]int64
type typMapMapUint8Float64 map[uint8]float64
type typMapMapUint8Bool map[uint8]bool
type typMapMapUint64Intf map[uint64]interface{}
type typMapMapUint64String map[uint64]string
type typMapMapUint64Bytes map[uint64][]byte
type typMapMapUint64Uint8 map[uint64]uint8
type typMapMapUint64Uint64 map[uint64]uint64
type typMapMapUint64Int map[uint64]int
type typMapMapUint64Int64 map[uint64]int64
type typMapMapUint64Float64 map[uint64]float64
type typMapMapUint64Bool map[uint64]bool
type typMapMapIntIntf map[int]interface{}
type typMapMapIntString map[int]string
type typMapMapIntBytes map[int][]byte
type typMapMapIntUint8 map[int]uint8
type typMapMapIntUint64 map[int]uint64
type typMapMapIntInt map[int]int
type typMapMapIntInt64 map[int]int64
type typMapMapIntFloat64 map[int]float64
type typMapMapIntBool map[int]bool
type typMapMapInt64Intf map[int64]interface{}
type typMapMapInt64String map[int64]string
type typMapMapInt64Bytes map[int64][]byte
type typMapMapInt64Uint8 map[int64]uint8
type typMapMapInt64Uint64 map[int64]uint64
type typMapMapInt64Int map[int64]int
type typMapMapInt64Int64 map[int64]int64
type typMapMapInt64Float64 map[int64]float64
type typMapMapInt64Bool map[int64]bool

func doTestMammothSlices(t *testing.T, h Handle) {
	var v17va [8]interface{}
	for _, v := range [][]interface{}{nil, {}, {"string-is-an-interface-2", nil, nil, "string-is-an-interface-3"}} {
		var v17v1, v17v2 []interface{}
		var bs17 []byte
		v17v1 = v
		bs17 = testMarshalErr(v17v1, h, t, "enc-slice-v17")
		if v != nil {
			if v == nil {
				v17v2 = nil
			} else {
				v17v2 = make([]interface{}, len(v))
			}
			testUnmarshalErr(v17v2, bs17, h, t, "dec-slice-v17")
			testDeepEqualErr(v17v1, v17v2, t, "equal-slice-v17")
			if v == nil {
				v17v2 = nil
			} else {
				v17v2 = make([]interface{}, len(v))
			}
			testUnmarshalErr(rv4i(v17v2), bs17, h, t, "dec-slice-v17-noaddr") // non-addressable value
			testDeepEqualErr(v17v1, v17v2, t, "equal-slice-v17-noaddr")
		}
		// ...
		bs17 = testMarshalErr(&v17v1, h, t, "enc-slice-v17-p")
		v17v2 = nil
		testUnmarshalErr(&v17v2, bs17, h, t, "dec-slice-v17-p")
		testDeepEqualErr(v17v1, v17v2, t, "equal-slice-v17-p")
		v17va = [8]interface{}{} // clear the array
		v17v2 = v17va[:1:1]
		testUnmarshalErr(&v17v2, bs17, h, t, "dec-slice-v17-p-1")
		testDeepEqualErr(v17v1, v17v2, t, "equal-slice-v17-p-1")
		v17va = [8]interface{}{} // clear the array
		v17v2 = v17va[:len(v17v1):len(v17v1)]
		testUnmarshalErr(&v17v2, bs17, h, t, "dec-slice-v17-p-len")
		testDeepEqualErr(v17v1, v17v2, t, "equal-slice-v17-p-len")
		v17va = [8]interface{}{} // clear the array
		v17v2 = v17va[:]
		testUnmarshalErr(&v17v2, bs17, h, t, "dec-slice-v17-p-cap")
		testDeepEqualErr(v17v1, v17v2, t, "equal-slice-v17-p-cap")
		if len(v17v1) > 1 {
			v17va = [8]interface{}{} // clear the array
			testUnmarshalErr((&v17va)[:len(v17v1)], bs17, h, t, "dec-slice-v17-p-len-noaddr")
			testDeepEqualErr(v17v1, v17va[:len(v17v1)], t, "equal-slice-v17-p-len-noaddr")
			v17va = [8]interface{}{} // clear the array
			testUnmarshalErr((&v17va)[:], bs17, h, t, "dec-slice-v17-p-cap-noaddr")
			testDeepEqualErr(v17v1, v17va[:len(v17v1)], t, "equal-slice-v17-p-cap-noaddr")
		}
		// ...
		var v17v3, v17v4 typMbsSliceIntf
		v17v2 = nil
		if v != nil {
			v17v2 = make([]interface{}, len(v))
		}
		v17v3 = typMbsSliceIntf(v17v1)
		v17v4 = typMbsSliceIntf(v17v2)
		if v != nil {
			bs17 = testMarshalErr(v17v3, h, t, "enc-slice-v17-custom")
			testUnmarshalErr(v17v4, bs17, h, t, "dec-slice-v17-custom")
			testDeepEqualErr(v17v3, v17v4, t, "equal-slice-v17-custom")
		}
		bs17 = testMarshalErr(&v17v3, h, t, "enc-slice-v17-custom-p")
		v17v2 = nil
		v17v4 = typMbsSliceIntf(v17v2)
		testUnmarshalErr(&v17v4, bs17, h, t, "dec-slice-v17-custom-p")
		testDeepEqualErr(v17v3, v17v4, t, "equal-slice-v17-custom-p")
	}
	var v18va [8]string
	for _, v := range [][]string{nil, {}, {"some-string-2", "", "", "some-string-3"}} {
		var v18v1, v18v2 []string
		var bs18 []byte
		v18v1 = v
		bs18 = testMarshalErr(v18v1, h, t, "enc-slice-v18")
		if v != nil {
			if v == nil {
				v18v2 = nil
			} else {
				v18v2 = make([]string, len(v))
			}
			testUnmarshalErr(v18v2, bs18, h, t, "dec-slice-v18")
			testDeepEqualErr(v18v1, v18v2, t, "equal-slice-v18")
			if v == nil {
				v18v2 = nil
			} else {
				v18v2 = make([]string, len(v))
			}
			testUnmarshalErr(rv4i(v18v2), bs18, h, t, "dec-slice-v18-noaddr") // non-addressable value
			testDeepEqualErr(v18v1, v18v2, t, "equal-slice-v18-noaddr")
		}
		// ...
		bs18 = testMarshalErr(&v18v1, h, t, "enc-slice-v18-p")
		v18v2 = nil
		testUnmarshalErr(&v18v2, bs18, h, t, "dec-slice-v18-p")
		testDeepEqualErr(v18v1, v18v2, t, "equal-slice-v18-p")
		v18va = [8]string{} // clear the array
		v18v2 = v18va[:1:1]
		testUnmarshalErr(&v18v2, bs18, h, t, "dec-slice-v18-p-1")
		testDeepEqualErr(v18v1, v18v2, t, "equal-slice-v18-p-1")
		v18va = [8]string{} // clear the array
		v18v2 = v18va[:len(v18v1):len(v18v1)]
		testUnmarshalErr(&v18v2, bs18, h, t, "dec-slice-v18-p-len")
		testDeepEqualErr(v18v1, v18v2, t, "equal-slice-v18-p-len")
		v18va = [8]string{} // clear the array
		v18v2 = v18va[:]
		testUnmarshalErr(&v18v2, bs18, h, t, "dec-slice-v18-p-cap")
		testDeepEqualErr(v18v1, v18v2, t, "equal-slice-v18-p-cap")
		if len(v18v1) > 1 {
			v18va = [8]string{} // clear the array
			testUnmarshalErr((&v18va)[:len(v18v1)], bs18, h, t, "dec-slice-v18-p-len-noaddr")
			testDeepEqualErr(v18v1, v18va[:len(v18v1)], t, "equal-slice-v18-p-len-noaddr")
			v18va = [8]string{} // clear the array
			testUnmarshalErr((&v18va)[:], bs18, h, t, "dec-slice-v18-p-cap-noaddr")
			testDeepEqualErr(v18v1, v18va[:len(v18v1)], t, "equal-slice-v18-p-cap-noaddr")
		}
		// ...
		var v18v3, v18v4 typMbsSliceString
		v18v2 = nil
		if v != nil {
			v18v2 = make([]string, len(v))
		}
		v18v3 = typMbsSliceString(v18v1)
		v18v4 = typMbsSliceString(v18v2)
		if v != nil {
			bs18 = testMarshalErr(v18v3, h, t, "enc-slice-v18-custom")
			testUnmarshalErr(v18v4, bs18, h, t, "dec-slice-v18-custom")
			testDeepEqualErr(v18v3, v18v4, t, "equal-slice-v18-custom")
		}
		bs18 = testMarshalErr(&v18v3, h, t, "enc-slice-v18-custom-p")
		v18v2 = nil
		v18v4 = typMbsSliceString(v18v2)
		testUnmarshalErr(&v18v4, bs18, h, t, "dec-slice-v18-custom-p")
		testDeepEqualErr(v18v3, v18v4, t, "equal-slice-v18-custom-p")
	}
	var v19va [8][]byte
	for _, v := range [][][]byte{nil, {}, {[]byte("some-string-2"), nil, nil, []byte("some-string-3")}} {
		var v19v1, v19v2 [][]byte
		var bs19 []byte
		v19v1 = v
		bs19 = testMarshalErr(v19v1, h, t, "enc-slice-v19")
		if v != nil {
			if v == nil {
				v19v2 = nil
			} else {
				v19v2 = make([][]byte, len(v))
			}
			testUnmarshalErr(v19v2, bs19, h, t, "dec-slice-v19")
			testDeepEqualErr(v19v1, v19v2, t, "equal-slice-v19")
			if v == nil {
				v19v2 = nil
			} else {
				v19v2 = make([][]byte, len(v))
			}
			testUnmarshalErr(rv4i(v19v2), bs19, h, t, "dec-slice-v19-noaddr") // non-addressable value
			testDeepEqualErr(v19v1, v19v2, t, "equal-slice-v19-noaddr")
		}
		// ...
		bs19 = testMarshalErr(&v19v1, h, t, "enc-slice-v19-p")
		v19v2 = nil
		testUnmarshalErr(&v19v2, bs19, h, t, "dec-slice-v19-p")
		testDeepEqualErr(v19v1, v19v2, t, "equal-slice-v19-p")
		v19va = [8][]byte{} // clear the array
		v19v2 = v19va[:1:1]
		testUnmarshalErr(&v19v2, bs19, h, t, "dec-slice-v19-p-1")
		testDeepEqualErr(v19v1, v19v2, t, "equal-slice-v19-p-1")
		v19va = [8][]byte{} // clear the array
		v19v2 = v19va[:len(v19v1):len(v19v1)]
		testUnmarshalErr(&v19v2, bs19, h, t, "dec-slice-v19-p-len")
		testDeepEqualErr(v19v1, v19v2, t, "equal-slice-v19-p-len")
		v19va = [8][]byte{} // clear the array
		v19v2 = v19va[:]
		testUnmarshalErr(&v19v2, bs19, h, t, "dec-slice-v19-p-cap")
		testDeepEqualErr(v19v1, v19v2, t, "equal-slice-v19-p-cap")
		if len(v19v1) > 1 {
			v19va = [8][]byte{} // clear the array
			testUnmarshalErr((&v19va)[:len(v19v1)], bs19, h, t, "dec-slice-v19-p-len-noaddr")
			testDeepEqualErr(v19v1, v19va[:len(v19v1)], t, "equal-slice-v19-p-len-noaddr")
			v19va = [8][]byte{} // clear the array
			testUnmarshalErr((&v19va)[:], bs19, h, t, "dec-slice-v19-p-cap-noaddr")
			testDeepEqualErr(v19v1, v19va[:len(v19v1)], t, "equal-slice-v19-p-cap-noaddr")
		}
		// ...
		var v19v3, v19v4 typMbsSliceBytes
		v19v2 = nil
		if v != nil {
			v19v2 = make([][]byte, len(v))
		}
		v19v3 = typMbsSliceBytes(v19v1)
		v19v4 = typMbsSliceBytes(v19v2)
		if v != nil {
			bs19 = testMarshalErr(v19v3, h, t, "enc-slice-v19-custom")
			testUnmarshalErr(v19v4, bs19, h, t, "dec-slice-v19-custom")
			testDeepEqualErr(v19v3, v19v4, t, "equal-slice-v19-custom")
		}
		bs19 = testMarshalErr(&v19v3, h, t, "enc-slice-v19-custom-p")
		v19v2 = nil
		v19v4 = typMbsSliceBytes(v19v2)
		testUnmarshalErr(&v19v4, bs19, h, t, "dec-slice-v19-custom-p")
		testDeepEqualErr(v19v3, v19v4, t, "equal-slice-v19-custom-p")
	}
	var v20va [8]float64
	for _, v := range [][]float64{nil, {}, {22.2, 0, 0, 33.3e3}} {
		var v20v1, v20v2 []float64
		var bs20 []byte
		v20v1 = v
		bs20 = testMarshalErr(v20v1, h, t, "enc-slice-v20")
		if v != nil {
			if v == nil {
				v20v2 = nil
			} else {
				v20v2 = make([]float64, len(v))
			}
			testUnmarshalErr(v20v2, bs20, h, t, "dec-slice-v20")
			testDeepEqualErr(v20v1, v20v2, t, "equal-slice-v20")
			if v == nil {
				v20v2 = nil
			} else {
				v20v2 = make([]float64, len(v))
			}
			testUnmarshalErr(rv4i(v20v2), bs20, h, t, "dec-slice-v20-noaddr") // non-addressable value
			testDeepEqualErr(v20v1, v20v2, t, "equal-slice-v20-noaddr")
		}
		// ...
		bs20 = testMarshalErr(&v20v1, h, t, "enc-slice-v20-p")
		v20v2 = nil
		testUnmarshalErr(&v20v2, bs20, h, t, "dec-slice-v20-p")
		testDeepEqualErr(v20v1, v20v2, t, "equal-slice-v20-p")
		v20va = [8]float64{} // clear the array
		v20v2 = v20va[:1:1]
		testUnmarshalErr(&v20v2, bs20, h, t, "dec-slice-v20-p-1")
		testDeepEqualErr(v20v1, v20v2, t, "equal-slice-v20-p-1")
		v20va = [8]float64{} // clear the array
		v20v2 = v20va[:len(v20v1):len(v20v1)]
		testUnmarshalErr(&v20v2, bs20, h, t, "dec-slice-v20-p-len")
		testDeepEqualErr(v20v1, v20v2, t, "equal-slice-v20-p-len")
		v20va = [8]float64{} // clear the array
		v20v2 = v20va[:]
		testUnmarshalErr(&v20v2, bs20, h, t, "dec-slice-v20-p-cap")
		testDeepEqualErr(v20v1, v20v2, t, "equal-slice-v20-p-cap")
		if len(v20v1) > 1 {
			v20va = [8]float64{} // clear the array
			testUnmarshalErr((&v20va)[:len(v20v1)], bs20, h, t, "dec-slice-v20-p-len-noaddr")
			testDeepEqualErr(v20v1, v20va[:len(v20v1)], t, "equal-slice-v20-p-len-noaddr")
			v20va = [8]float64{} // clear the array
			testUnmarshalErr((&v20va)[:], bs20, h, t, "dec-slice-v20-p-cap-noaddr")
			testDeepEqualErr(v20v1, v20va[:len(v20v1)], t, "equal-slice-v20-p-cap-noaddr")
		}
		// ...
		var v20v3, v20v4 typMbsSliceFloat64
		v20v2 = nil
		if v != nil {
			v20v2 = make([]float64, len(v))
		}
		v20v3 = typMbsSliceFloat64(v20v1)
		v20v4 = typMbsSliceFloat64(v20v2)
		if v != nil {
			bs20 = testMarshalErr(v20v3, h, t, "enc-slice-v20-custom")
			testUnmarshalErr(v20v4, bs20, h, t, "dec-slice-v20-custom")
			testDeepEqualErr(v20v3, v20v4, t, "equal-slice-v20-custom")
		}
		bs20 = testMarshalErr(&v20v3, h, t, "enc-slice-v20-custom-p")
		v20v2 = nil
		v20v4 = typMbsSliceFloat64(v20v2)
		testUnmarshalErr(&v20v4, bs20, h, t, "dec-slice-v20-custom-p")
		testDeepEqualErr(v20v3, v20v4, t, "equal-slice-v20-custom-p")
	}
	var v21va [8]uint64
	for _, v := range [][]uint64{nil, {}, {77, 0, 0, 127}} {
		var v21v1, v21v2 []uint64
		var bs21 []byte
		v21v1 = v
		bs21 = testMarshalErr(v21v1, h, t, "enc-slice-v21")
		if v != nil {
			if v == nil {
				v21v2 = nil
			} else {
				v21v2 = make([]uint64, len(v))
			}
			testUnmarshalErr(v21v2, bs21, h, t, "dec-slice-v21")
			testDeepEqualErr(v21v1, v21v2, t, "equal-slice-v21")
			if v == nil {
				v21v2 = nil
			} else {
				v21v2 = make([]uint64, len(v))
			}
			testUnmarshalErr(rv4i(v21v2), bs21, h, t, "dec-slice-v21-noaddr") // non-addressable value
			testDeepEqualErr(v21v1, v21v2, t, "equal-slice-v21-noaddr")
		}
		// ...
		bs21 = testMarshalErr(&v21v1, h, t, "enc-slice-v21-p")
		v21v2 = nil
		testUnmarshalErr(&v21v2, bs21, h, t, "dec-slice-v21-p")
		testDeepEqualErr(v21v1, v21v2, t, "equal-slice-v21-p")
		v21va = [8]uint64{} // clear the array
		v21v2 = v21va[:1:1]
		testUnmarshalErr(&v21v2, bs21, h, t, "dec-slice-v21-p-1")
		testDeepEqualErr(v21v1, v21v2, t, "equal-slice-v21-p-1")
		v21va = [8]uint64{} // clear the array
		v21v2 = v21va[:len(v21v1):len(v21v1)]
		testUnmarshalErr(&v21v2, bs21, h, t, "dec-slice-v21-p-len")
		testDeepEqualErr(v21v1, v21v2, t, "equal-slice-v21-p-len")
		v21va = [8]uint64{} // clear the array
		v21v2 = v21va[:]
		testUnmarshalErr(&v21v2, bs21, h, t, "dec-slice-v21-p-cap")
		testDeepEqualErr(v21v1, v21v2, t, "equal-slice-v21-p-cap")
		if len(v21v1) > 1 {
			v21va = [8]uint64{} // clear the array
			testUnmarshalErr((&v21va)[:len(v21v1)], bs21, h, t, "dec-slice-v21-p-len-noaddr")
			testDeepEqualErr(v21v1, v21va[:len(v21v1)], t, "equal-slice-v21-p-len-noaddr")
			v21va = [8]uint64{} // clear the array
			testUnmarshalErr((&v21va)[:], bs21, h, t, "dec-slice-v21-p-cap-noaddr")
			testDeepEqualErr(v21v1, v21va[:len(v21v1)], t, "equal-slice-v21-p-cap-noaddr")
		}
		// ...
		var v21v3, v21v4 typMbsSliceUint64
		v21v2 = nil
		if v != nil {
			v21v2 = make([]uint64, len(v))
		}
		v21v3 = typMbsSliceUint64(v21v1)
		v21v4 = typMbsSliceUint64(v21v2)
		if v != nil {
			bs21 = testMarshalErr(v21v3, h, t, "enc-slice-v21-custom")
			testUnmarshalErr(v21v4, bs21, h, t, "dec-slice-v21-custom")
			testDeepEqualErr(v21v3, v21v4, t, "equal-slice-v21-custom")
		}
		bs21 = testMarshalErr(&v21v3, h, t, "enc-slice-v21-custom-p")
		v21v2 = nil
		v21v4 = typMbsSliceUint64(v21v2)
		testUnmarshalErr(&v21v4, bs21, h, t, "dec-slice-v21-custom-p")
		testDeepEqualErr(v21v3, v21v4, t, "equal-slice-v21-custom-p")
	}
	var v22va [8]int
	for _, v := range [][]int{nil, {}, {111, 0, 0, 77}} {
		var v22v1, v22v2 []int
		var bs22 []byte
		v22v1 = v
		bs22 = testMarshalErr(v22v1, h, t, "enc-slice-v22")
		if v != nil {
			if v == nil {
				v22v2 = nil
			} else {
				v22v2 = make([]int, len(v))
			}
			testUnmarshalErr(v22v2, bs22, h, t, "dec-slice-v22")
			testDeepEqualErr(v22v1, v22v2, t, "equal-slice-v22")
			if v == nil {
				v22v2 = nil
			} else {
				v22v2 = make([]int, len(v))
			}
			testUnmarshalErr(rv4i(v22v2), bs22, h, t, "dec-slice-v22-noaddr") // non-addressable value
			testDeepEqualErr(v22v1, v22v2, t, "equal-slice-v22-noaddr")
		}
		// ...
		bs22 = testMarshalErr(&v22v1, h, t, "enc-slice-v22-p")
		v22v2 = nil
		testUnmarshalErr(&v22v2, bs22, h, t, "dec-slice-v22-p")
		testDeepEqualErr(v22v1, v22v2, t, "equal-slice-v22-p")
		v22va = [8]int{} // clear the array
		v22v2 = v22va[:1:1]
		testUnmarshalErr(&v22v2, bs22, h, t, "dec-slice-v22-p-1")
		testDeepEqualErr(v22v1, v22v2, t, "equal-slice-v22-p-1")
		v22va = [8]int{} // clear the array
		v22v2 = v22va[:len(v22v1):len(v22v1)]
		testUnmarshalErr(&v22v2, bs22, h, t, "dec-slice-v22-p-len")
		testDeepEqualErr(v22v1, v22v2, t, "equal-slice-v22-p-len")
		v22va = [8]int{} // clear the array
		v22v2 = v22va[:]
		testUnmarshalErr(&v22v2, bs22, h, t, "dec-slice-v22-p-cap")
		testDeepEqualErr(v22v1, v22v2, t, "equal-slice-v22-p-cap")
		if len(v22v1) > 1 {
			v22va = [8]int{} // clear the array
			testUnmarshalErr((&v22va)[:len(v22v1)], bs22, h, t, "dec-slice-v22-p-len-noaddr")
			testDeepEqualErr(v22v1, v22va[:len(v22v1)], t, "equal-slice-v22-p-len-noaddr")
			v22va = [8]int{} // clear the array
			testUnmarshalErr((&v22va)[:], bs22, h, t, "dec-slice-v22-p-cap-noaddr")
			testDeepEqualErr(v22v1, v22va[:len(v22v1)], t, "equal-slice-v22-p-cap-noaddr")
		}
		// ...
		var v22v3, v22v4 typMbsSliceInt
		v22v2 = nil
		if v != nil {
			v22v2 = make([]int, len(v))
		}
		v22v3 = typMbsSliceInt(v22v1)
		v22v4 = typMbsSliceInt(v22v2)
		if v != nil {
			bs22 = testMarshalErr(v22v3, h, t, "enc-slice-v22-custom")
			testUnmarshalErr(v22v4, bs22, h, t, "dec-slice-v22-custom")
			testDeepEqualErr(v22v3, v22v4, t, "equal-slice-v22-custom")
		}
		bs22 = testMarshalErr(&v22v3, h, t, "enc-slice-v22-custom-p")
		v22v2 = nil
		v22v4 = typMbsSliceInt(v22v2)
		testUnmarshalErr(&v22v4, bs22, h, t, "dec-slice-v22-custom-p")
		testDeepEqualErr(v22v3, v22v4, t, "equal-slice-v22-custom-p")
	}
	var v23va [8]int32
	for _, v := range [][]int32{nil, {}, {127, 0, 0, 111}} {
		var v23v1, v23v2 []int32
		var bs23 []byte
		v23v1 = v
		bs23 = testMarshalErr(v23v1, h, t, "enc-slice-v23")
		if v != nil {
			if v == nil {
				v23v2 = nil
			} else {
				v23v2 = make([]int32, len(v))
			}
			testUnmarshalErr(v23v2, bs23, h, t, "dec-slice-v23")
			testDeepEqualErr(v23v1, v23v2, t, "equal-slice-v23")
			if v == nil {
				v23v2 = nil
			} else {
				v23v2 = make([]int32, len(v))
			}
			testUnmarshalErr(rv4i(v23v2), bs23, h, t, "dec-slice-v23-noaddr") // non-addressable value
			testDeepEqualErr(v23v1, v23v2, t, "equal-slice-v23-noaddr")
		}
		// ...
		bs23 = testMarshalErr(&v23v1, h, t, "enc-slice-v23-p")
		v23v2 = nil
		testUnmarshalErr(&v23v2, bs23, h, t, "dec-slice-v23-p")
		testDeepEqualErr(v23v1, v23v2, t, "equal-slice-v23-p")
		v23va = [8]int32{} // clear the array
		v23v2 = v23va[:1:1]
		testUnmarshalErr(&v23v2, bs23, h, t, "dec-slice-v23-p-1")
		testDeepEqualErr(v23v1, v23v2, t, "equal-slice-v23-p-1")
		v23va = [8]int32{} // clear the array
		v23v2 = v23va[:len(v23v1):len(v23v1)]
		testUnmarshalErr(&v23v2, bs23, h, t, "dec-slice-v23-p-len")
		testDeepEqualErr(v23v1, v23v2, t, "equal-slice-v23-p-len")
		v23va = [8]int32{} // clear the array
		v23v2 = v23va[:]
		testUnmarshalErr(&v23v2, bs23, h, t, "dec-slice-v23-p-cap")
		testDeepEqualErr(v23v1, v23v2, t, "equal-slice-v23-p-cap")
		if len(v23v1) > 1 {
			v23va = [8]int32{} // clear the array
			testUnmarshalErr((&v23va)[:len(v23v1)], bs23, h, t, "dec-slice-v23-p-len-noaddr")
			testDeepEqualErr(v23v1, v23va[:len(v23v1)], t, "equal-slice-v23-p-len-noaddr")
			v23va = [8]int32{} // clear the array
			testUnmarshalErr((&v23va)[:], bs23, h, t, "dec-slice-v23-p-cap-noaddr")
			testDeepEqualErr(v23v1, v23va[:len(v23v1)], t, "equal-slice-v23-p-cap-noaddr")
		}
		// ...
		var v23v3, v23v4 typMbsSliceInt32
		v23v2 = nil
		if v != nil {
			v23v2 = make([]int32, len(v))
		}
		v23v3 = typMbsSliceInt32(v23v1)
		v23v4 = typMbsSliceInt32(v23v2)
		if v != nil {
			bs23 = testMarshalErr(v23v3, h, t, "enc-slice-v23-custom")
			testUnmarshalErr(v23v4, bs23, h, t, "dec-slice-v23-custom")
			testDeepEqualErr(v23v3, v23v4, t, "equal-slice-v23-custom")
		}
		bs23 = testMarshalErr(&v23v3, h, t, "enc-slice-v23-custom-p")
		v23v2 = nil
		v23v4 = typMbsSliceInt32(v23v2)
		testUnmarshalErr(&v23v4, bs23, h, t, "dec-slice-v23-custom-p")
		testDeepEqualErr(v23v3, v23v4, t, "equal-slice-v23-custom-p")
	}
	var v24va [8]int64
	for _, v := range [][]int64{nil, {}, {77, 0, 0, 127}} {
		var v24v1, v24v2 []int64
		var bs24 []byte
		v24v1 = v
		bs24 = testMarshalErr(v24v1, h, t, "enc-slice-v24")
		if v != nil {
			if v == nil {
				v24v2 = nil
			} else {
				v24v2 = make([]int64, len(v))
			}
			testUnmarshalErr(v24v2, bs24, h, t, "dec-slice-v24")
			testDeepEqualErr(v24v1, v24v2, t, "equal-slice-v24")
			if v == nil {
				v24v2 = nil
			} else {
				v24v2 = make([]int64, len(v))
			}
			testUnmarshalErr(rv4i(v24v2), bs24, h, t, "dec-slice-v24-noaddr") // non-addressable value
			testDeepEqualErr(v24v1, v24v2, t, "equal-slice-v24-noaddr")
		}
		// ...
		bs24 = testMarshalErr(&v24v1, h, t, "enc-slice-v24-p")
		v24v2 = nil
		testUnmarshalErr(&v24v2, bs24, h, t, "dec-slice-v24-p")
		testDeepEqualErr(v24v1, v24v2, t, "equal-slice-v24-p")
		v24va = [8]int64{} // clear the array
		v24v2 = v24va[:1:1]
		testUnmarshalErr(&v24v2, bs24, h, t, "dec-slice-v24-p-1")
		testDeepEqualErr(v24v1, v24v2, t, "equal-slice-v24-p-1")
		v24va = [8]int64{} // clear the array
		v24v2 = v24va[:len(v24v1):len(v24v1)]
		testUnmarshalErr(&v24v2, bs24, h, t, "dec-slice-v24-p-len")
		testDeepEqualErr(v24v1, v24v2, t, "equal-slice-v24-p-len")
		v24va = [8]int64{} // clear the array
		v24v2 = v24va[:]
		testUnmarshalErr(&v24v2, bs24, h, t, "dec-slice-v24-p-cap")
		testDeepEqualErr(v24v1, v24v2, t, "equal-slice-v24-p-cap")
		if len(v24v1) > 1 {
			v24va = [8]int64{} // clear the array
			testUnmarshalErr((&v24va)[:len(v24v1)], bs24, h, t, "dec-slice-v24-p-len-noaddr")
			testDeepEqualErr(v24v1, v24va[:len(v24v1)], t, "equal-slice-v24-p-len-noaddr")
			v24va = [8]int64{} // clear the array
			testUnmarshalErr((&v24va)[:], bs24, h, t, "dec-slice-v24-p-cap-noaddr")
			testDeepEqualErr(v24v1, v24va[:len(v24v1)], t, "equal-slice-v24-p-cap-noaddr")
		}
		// ...
		var v24v3, v24v4 typMbsSliceInt64
		v24v2 = nil
		if v != nil {
			v24v2 = make([]int64, len(v))
		}
		v24v3 = typMbsSliceInt64(v24v1)
		v24v4 = typMbsSliceInt64(v24v2)
		if v != nil {
			bs24 = testMarshalErr(v24v3, h, t, "enc-slice-v24-custom")
			testUnmarshalErr(v24v4, bs24, h, t, "dec-slice-v24-custom")
			testDeepEqualErr(v24v3, v24v4, t, "equal-slice-v24-custom")
		}
		bs24 = testMarshalErr(&v24v3, h, t, "enc-slice-v24-custom-p")
		v24v2 = nil
		v24v4 = typMbsSliceInt64(v24v2)
		testUnmarshalErr(&v24v4, bs24, h, t, "dec-slice-v24-custom-p")
		testDeepEqualErr(v24v3, v24v4, t, "equal-slice-v24-custom-p")
	}
	var v25va [8]bool
	for _, v := range [][]bool{nil, {}, {false, false, false, true}} {
		var v25v1, v25v2 []bool
		var bs25 []byte
		v25v1 = v
		bs25 = testMarshalErr(v25v1, h, t, "enc-slice-v25")
		if v != nil {
			if v == nil {
				v25v2 = nil
			} else {
				v25v2 = make([]bool, len(v))
			}
			testUnmarshalErr(v25v2, bs25, h, t, "dec-slice-v25")
			testDeepEqualErr(v25v1, v25v2, t, "equal-slice-v25")
			if v == nil {
				v25v2 = nil
			} else {
				v25v2 = make([]bool, len(v))
			}
			testUnmarshalErr(rv4i(v25v2), bs25, h, t, "dec-slice-v25-noaddr") // non-addressable value
			testDeepEqualErr(v25v1, v25v2, t, "equal-slice-v25-noaddr")
		}
		// ...
		bs25 = testMarshalErr(&v25v1, h, t, "enc-slice-v25-p")
		v25v2 = nil
		testUnmarshalErr(&v25v2, bs25, h, t, "dec-slice-v25-p")
		testDeepEqualErr(v25v1, v25v2, t, "equal-slice-v25-p")
		v25va = [8]bool{} // clear the array
		v25v2 = v25va[:1:1]
		testUnmarshalErr(&v25v2, bs25, h, t, "dec-slice-v25-p-1")
		testDeepEqualErr(v25v1, v25v2, t, "equal-slice-v25-p-1")
		v25va = [8]bool{} // clear the array
		v25v2 = v25va[:len(v25v1):len(v25v1)]
		testUnmarshalErr(&v25v2, bs25, h, t, "dec-slice-v25-p-len")
		testDeepEqualErr(v25v1, v25v2, t, "equal-slice-v25-p-len")
		v25va = [8]bool{} // clear the array
		v25v2 = v25va[:]
		testUnmarshalErr(&v25v2, bs25, h, t, "dec-slice-v25-p-cap")
		testDeepEqualErr(v25v1, v25v2, t, "equal-slice-v25-p-cap")
		if len(v25v1) > 1 {
			v25va = [8]bool{} // clear the array
			testUnmarshalErr((&v25va)[:len(v25v1)], bs25, h, t, "dec-slice-v25-p-len-noaddr")
			testDeepEqualErr(v25v1, v25va[:len(v25v1)], t, "equal-slice-v25-p-len-noaddr")
			v25va = [8]bool{} // clear the array
			testUnmarshalErr((&v25va)[:], bs25, h, t, "dec-slice-v25-p-cap-noaddr")
			testDeepEqualErr(v25v1, v25va[:len(v25v1)], t, "equal-slice-v25-p-cap-noaddr")
		}
		// ...
		var v25v3, v25v4 typMbsSliceBool
		v25v2 = nil
		if v != nil {
			v25v2 = make([]bool, len(v))
		}
		v25v3 = typMbsSliceBool(v25v1)
		v25v4 = typMbsSliceBool(v25v2)
		if v != nil {
			bs25 = testMarshalErr(v25v3, h, t, "enc-slice-v25-custom")
			testUnmarshalErr(v25v4, bs25, h, t, "dec-slice-v25-custom")
			testDeepEqualErr(v25v3, v25v4, t, "equal-slice-v25-custom")
		}
		bs25 = testMarshalErr(&v25v3, h, t, "enc-slice-v25-custom-p")
		v25v2 = nil
		v25v4 = typMbsSliceBool(v25v2)
		testUnmarshalErr(&v25v4, bs25, h, t, "dec-slice-v25-custom-p")
		testDeepEqualErr(v25v3, v25v4, t, "equal-slice-v25-custom-p")
	}

}

func doTestMammothMaps(t *testing.T, h Handle) {
	for _, v := range []map[string]interface{}{nil, {}, {"some-string-1": nil, "some-string-2": "string-is-an-interface-1"}} {
		// fmt.Printf(">>>> running mammoth map v26: %v\n", v)
		var v26v1, v26v2 map[string]interface{}
		var bs26 []byte
		v26v1 = v
		bs26 = testMarshalErr(v26v1, h, t, "enc-map-v26")
		if v != nil {
			if v == nil {
				v26v2 = nil
			} else {
				v26v2 = make(map[string]interface{}, len(v))
			} // reset map
			testUnmarshalErr(v26v2, bs26, h, t, "dec-map-v26")
			testDeepEqualErr(v26v1, v26v2, t, "equal-map-v26")
			if v == nil {
				v26v2 = nil
			} else {
				v26v2 = make(map[string]interface{}, len(v))
			} // reset map
			testUnmarshalErr(rv4i(v26v2), bs26, h, t, "dec-map-v26-noaddr") // decode into non-addressable map value
			testDeepEqualErr(v26v1, v26v2, t, "equal-map-v26-noaddr")
		}
		if v == nil {
			v26v2 = nil
		} else {
			v26v2 = make(map[string]interface{}, len(v))
		} // reset map
		testUnmarshalErr(&v26v2, bs26, h, t, "dec-map-v26-p-len")
		testDeepEqualErr(v26v1, v26v2, t, "equal-map-v26-p-len")
		bs26 = testMarshalErr(&v26v1, h, t, "enc-map-v26-p")
		v26v2 = nil
		testUnmarshalErr(&v26v2, bs26, h, t, "dec-map-v26-p-nil")
		testDeepEqualErr(v26v1, v26v2, t, "equal-map-v26-p-nil")
		// ...
		if v == nil {
			v26v2 = nil
		} else {
			v26v2 = make(map[string]interface{}, len(v))
		} // reset map
		var v26v3, v26v4 typMapMapStringIntf
		v26v3 = typMapMapStringIntf(v26v1)
		v26v4 = typMapMapStringIntf(v26v2)
		if v != nil {
			bs26 = testMarshalErr(v26v3, h, t, "enc-map-v26-custom")
			testUnmarshalErr(v26v4, bs26, h, t, "dec-map-v26-p-len")
			testDeepEqualErr(v26v3, v26v4, t, "equal-map-v26-p-len")
		}
	}
	for _, v := range []map[string]string{nil, {}, {"some-string-3": "", "some-string-1": "some-string-2"}} {
		// fmt.Printf(">>>> running mammoth map v27: %v\n", v)
		var v27v1, v27v2 map[string]string
		var bs27 []byte
		v27v1 = v
		bs27 = testMarshalErr(v27v1, h, t, "enc-map-v27")
		if v != nil {
			if v == nil {
				v27v2 = nil
			} else {
				v27v2 = make(map[string]string, len(v))
			} // reset map
			testUnmarshalErr(v27v2, bs27, h, t, "dec-map-v27")
			testDeepEqualErr(v27v1, v27v2, t, "equal-map-v27")
			if v == nil {
				v27v2 = nil
			} else {
				v27v2 = make(map[string]string, len(v))
			} // reset map
			testUnmarshalErr(rv4i(v27v2), bs27, h, t, "dec-map-v27-noaddr") // decode into non-addressable map value
			testDeepEqualErr(v27v1, v27v2, t, "equal-map-v27-noaddr")
		}
		if v == nil {
			v27v2 = nil
		} else {
			v27v2 = make(map[string]string, len(v))
		} // reset map
		testUnmarshalErr(&v27v2, bs27, h, t, "dec-map-v27-p-len")
		testDeepEqualErr(v27v1, v27v2, t, "equal-map-v27-p-len")
		bs27 = testMarshalErr(&v27v1, h, t, "enc-map-v27-p")
		v27v2 = nil
		testUnmarshalErr(&v27v2, bs27, h, t, "dec-map-v27-p-nil")
		testDeepEqualErr(v27v1, v27v2, t, "equal-map-v27-p-nil")
		// ...
		if v == nil {
			v27v2 = nil
		} else {
			v27v2 = make(map[string]string, len(v))
		} // reset map
		var v27v3, v27v4 typMapMapStringString
		v27v3 = typMapMapStringString(v27v1)
		v27v4 = typMapMapStringString(v27v2)
		if v != nil {
			bs27 = testMarshalErr(v27v3, h, t, "enc-map-v27-custom")
			testUnmarshalErr(v27v4, bs27, h, t, "dec-map-v27-p-len")
			testDeepEqualErr(v27v3, v27v4, t, "equal-map-v27-p-len")
		}
	}
	for _, v := range []map[string][]byte{nil, {}, {"some-string-3": nil, "some-string-1": []byte("some-string-1")}} {
		// fmt.Printf(">>>> running mammoth map v28: %v\n", v)
		var v28v1, v28v2 map[string][]byte
		var bs28 []byte
		v28v1 = v
		bs28 = testMarshalErr(v28v1, h, t, "enc-map-v28")
		if v != nil {
			if v == nil {
				v28v2 = nil
			} else {
				v28v2 = make(map[string][]byte, len(v))
			} // reset map
			testUnmarshalErr(v28v2, bs28, h, t, "dec-map-v28")
			testDeepEqualErr(v28v1, v28v2, t, "equal-map-v28")
			if v == nil {
				v28v2 = nil
			} else {
				v28v2 = make(map[string][]byte, len(v))
			} // reset map
			testUnmarshalErr(rv4i(v28v2), bs28, h, t, "dec-map-v28-noaddr") // decode into non-addressable map value
			testDeepEqualErr(v28v1, v28v2, t, "equal-map-v28-noaddr")
		}
		if v == nil {
			v28v2 = nil
		} else {
			v28v2 = make(map[string][]byte, len(v))
		} // reset map
		testUnmarshalErr(&v28v2, bs28, h, t, "dec-map-v28-p-len")
		testDeepEqualErr(v28v1, v28v2, t, "equal-map-v28-p-len")
		bs28 = testMarshalErr(&v28v1, h, t, "enc-map-v28-p")
		v28v2 = nil
		testUnmarshalErr(&v28v2, bs28, h, t, "dec-map-v28-p-nil")
		testDeepEqualErr(v28v1, v28v2, t, "equal-map-v28-p-nil")
		// ...
		if v == nil {
			v28v2 = nil
		} else {
			v28v2 = make(map[string][]byte, len(v))
		} // reset map
		var v28v3, v28v4 typMapMapStringBytes
		v28v3 = typMapMapStringBytes(v28v1)
		v28v4 = typMapMapStringBytes(v28v2)
		if v != nil {
			bs28 = testMarshalErr(v28v3, h, t, "enc-map-v28-custom")
			testUnmarshalErr(v28v4, bs28, h, t, "dec-map-v28-p-len")
			testDeepEqualErr(v28v3, v28v4, t, "equal-map-v28-p-len")
		}
	}
	for _, v := range []map[string]uint8{nil, {}, {"some-string-2": 0, "some-string-3": 111}} {
		// fmt.Printf(">>>> running mammoth map v29: %v\n", v)
		var v29v1, v29v2 map[string]uint8
		var bs29 []byte
		v29v1 = v
		bs29 = testMarshalErr(v29v1, h, t, "enc-map-v29")
		if v != nil {
			if v == nil {
				v29v2 = nil
			} else {
				v29v2 = make(map[string]uint8, len(v))
			} // reset map
			testUnmarshalErr(v29v2, bs29, h, t, "dec-map-v29")
			testDeepEqualErr(v29v1, v29v2, t, "equal-map-v29")
			if v == nil {
				v29v2 = nil
			} else {
				v29v2 = make(map[string]uint8, len(v))
			} // reset map
			testUnmarshalErr(rv4i(v29v2), bs29, h, t, "dec-map-v29-noaddr") // decode into non-addressable map value
			testDeepEqualErr(v29v1, v29v2, t, "equal-map-v29-noaddr")
		}
		if v == nil {
			v29v2 = nil
		} else {
			v29v2 = make(map[string]uint8, len(v))
		} // reset map
		testUnmarshalErr(&v29v2, bs29, h, t, "dec-map-v29-p-len")
		testDeepEqualErr(v29v1, v29v2, t, "equal-map-v29-p-len")
		bs29 = testMarshalErr(&v29v1, h, t, "enc-map-v29-p")
		v29v2 = nil
		testUnmarshalErr(&v29v2, bs29, h, t, "dec-map-v29-p-nil")
		testDeepEqualErr(v29v1, v29v2, t, "equal-map-v29-p-nil")
		// ...
		if v == nil {
			v29v2 = nil
		} else {
			v29v2 = make(map[string]uint8, len(v))
		} // reset map
		var v29v3, v29v4 typMapMapStringUint8
		v29v3 = typMapMapStringUint8(v29v1)
		v29v4 = typMapMapStringUint8(v29v2)
		if v != nil {
			bs29 = testMarshalErr(v29v3, h, t, "enc-map-v29-custom")
			testUnmarshalErr(v29v4, bs29, h, t, "dec-map-v29-p-len")
			testDeepEqualErr(v29v3, v29v4, t, "equal-map-v29-p-len")
		}
	}
	for _, v := range []map[string]uint64{nil, {}, {"some-string-1": 0, "some-string-2": 77}} {
		// fmt.Printf(">>>> running mammoth map v30: %v\n", v)
		var v30v1, v30v2 map[string]uint64
		var bs30 []byte
		v30v1 = v
		bs30 = testMarshalErr(v30v1, h, t, "enc-map-v30")
		if v != nil {
			if v == nil {
				v30v2 = nil
			} else {
				v30v2 = make(map[string]uint64, len(v))
			} // reset map
			testUnmarshalErr(v30v2, bs30, h, t, "dec-map-v30")
			testDeepEqualErr(v30v1, v30v2, t, "equal-map-v30")
			if v == nil {
				v30v2 = nil
			} else {
				v30v2 = make(map[string]uint64, len(v))
			} // reset map
			testUnmarshalErr(rv4i(v30v2), bs30, h, t, "dec-map-v30-noaddr") // decode into non-addressable map value
			testDeepEqualErr(v30v1, v30v2, t, "equal-map-v30-noaddr")
		}
		if v == nil {
			v30v2 = nil
		} else {
			v30v2 = make(map[string]uint64, len(v))
		} // reset map
		testUnmarshalErr(&v30v2, bs30, h, t, "dec-map-v30-p-len")
		testDeepEqualErr(v30v1, v30v2, t, "equal-map-v30-p-len")
		bs30 = testMarshalErr(&v30v1, h, t, "enc-map-v30-p")
		v30v2 = nil
		testUnmarshalErr(&v30v2, bs30, h, t, "dec-map-v30-p-nil")
		testDeepEqualErr(v30v1, v30v2, t, "equal-map-v30-p-nil")
		// ...
		if v == nil {
			v30v2 = nil
		} else {
			v30v2 = make(map[string]uint64, len(v))
		} // reset map
		var v30v3, v30v4 typMapMapStringUint64
		v30v3 = typMapMapStringUint64(v30v1)
		v30v4 = typMapMapStringUint64(v30v2)
		if v != nil {
			bs30 = testMarshalErr(v30v3, h, t, "enc-map-v30-custom")
			testUnmarshalErr(v30v4, bs30, h, t, "dec-map-v30-p-len")
			testDeepEqualErr(v30v3, v30v4, t, "equal-map-v30-p-len")
		}
	}
	for _, v := range []map[string]int{nil, {}, {"some-string-3": 0, "some-string-1": 127}} {
		// fmt.Printf(">>>> running mammoth map v31: %v\n", v)
		var v31v1, v31v2 map[string]int
		var bs31 []byte
		v31v1 = v
		bs31 = testMarshalErr(v31v1, h, t, "enc-map-v31")
		if v != nil {
			if v == nil {
				v31v2 = nil
			} else {
				v31v2 = make(map[string]int, len(v))
			} // reset map
			testUnmarshalErr(v31v2, bs31, h, t, "dec-map-v31")
			testDeepEqualErr(v31v1, v31v2, t, "equal-map-v31")
			if v == nil {
				v31v2 = nil
			} else {
				v31v2 = make(map[string]int, len(v))
			} // reset map
			testUnmarshalErr(rv4i(v31v2), bs31, h, t, "dec-map-v31-noaddr") // decode into non-addressable map value
			testDeepEqualErr(v31v1, v31v2, t, "equal-map-v31-noaddr")
		}
		if v == nil {
			v31v2 = nil
		} else {
			v31v2 = make(map[string]int, len(v))
		} // reset map
		testUnmarshalErr(&v31v2, bs31, h, t, "dec-map-v31-p-len")
		testDeepEqualErr(v31v1, v31v2, t, "equal-map-v31-p-len")
		bs31 = testMarshalErr(&v31v1, h, t, "enc-map-v31-p")
		v31v2 = nil
		testUnmarshalErr(&v31v2, bs31, h, t, "dec-map-v31-p-nil")
		testDeepEqualErr(v31v1, v31v2, t, "equal-map-v31-p-nil")
		// ...
		if v == nil {
			v31v2 = nil
		} else {
			v31v2 = make(map[string]int, len(v))
		} // reset map
		var v31v3, v31v4 typMapMapStringInt
		v31v3 = typMapMapStringInt(v31v1)
		v31v4 = typMapMapStringInt(v31v2)
		if v != nil {
			bs31 = testMarshalErr(v31v3, h, t, "enc-map-v31-custom")
			testUnmarshalErr(v31v4, bs31, h, t, "dec-map-v31-p-len")
			testDeepEqualErr(v31v3, v31v4, t, "equal-map-v31-p-len")
		}
	}
	for _, v := range []map[string]int64{nil, {}, {"some-string-2": 0, "some-string-3": 111}} {
		// fmt.Printf(">>>> running mammoth map v32: %v\n", v)
		var v32v1, v32v2 map[string]int64
		var bs32 []byte
		v32v1 = v
		bs32 = testMarshalErr(v32v1, h, t, "enc-map-v32")
		if v != nil {
			if v == nil {
				v32v2 = nil
			} else {
				v32v2 = make(map[string]int64, len(v))
			} // reset map
			testUnmarshalErr(v32v2, bs32, h, t, "dec-map-v32")
			testDeepEqualErr(v32v1, v32v2, t, "equal-map-v32")
			if v == nil {
				v32v2 = nil
			} else {
				v32v2 = make(map[string]int64, len(v))
			} // reset map
			testUnmarshalErr(rv4i(v32v2), bs32, h, t, "dec-map-v32-noaddr") // decode into non-addressable map value
			testDeepEqualErr(v32v1, v32v2, t, "equal-map-v32-noaddr")
		}
		if v == nil {
			v32v2 = nil
		} else {
			v32v2 = make(map[string]int64, len(v))
		} // reset map
		testUnmarshalErr(&v32v2, bs32, h, t, "dec-map-v32-p-len")
		testDeepEqualErr(v32v1, v32v2, t, "equal-map-v32-p-len")
		bs32 = testMarshalErr(&v32v1, h, t, "enc-map-v32-p")
		v32v2 = nil
		testUnmarshalErr(&v32v2, bs32, h, t, "dec-map-v32-p-nil")
		testDeepEqualErr(v32v1, v32v2, t, "equal-map-v32-p-nil")
		// ...
		if v == nil {
			v32v2 = nil
		} else {
			v32v2 = make(map[string]int64, len(v))
		} // reset map
		var v32v3, v32v4 typMapMapStringInt64
		v32v3 = typMapMapStringInt64(v32v1)
		v32v4 = typMapMapStringInt64(v32v2)
		if v != nil {
			bs32 = testMarshalErr(v32v3, h, t, "enc-map-v32-custom")
			testUnmarshalErr(v32v4, bs32, h, t, "dec-map-v32-p-len")
			testDeepEqualErr(v32v3, v32v4, t, "equal-map-v32-p-len")
		}
	}
	for _, v := range []map[string]float64{nil, {}, {"some-string-1": 0, "some-string-2": 11.1}} {
		// fmt.Printf(">>>> running mammoth map v33: %v\n", v)
		var v33v1, v33v2 map[string]float64
		var bs33 []byte
		v33v1 = v
		bs33 = testMarshalErr(v33v1, h, t, "enc-map-v33")
		if v != nil {
			if v == nil {
				v33v2 = nil
			} else {
				v33v2 = make(map[string]float64, len(v))
			} // reset map
			testUnmarshalErr(v33v2, bs33, h, t, "dec-map-v33")
			testDeepEqualErr(v33v1, v33v2, t, "equal-map-v33")
			if v == nil {
				v33v2 = nil
			} else {
				v33v2 = make(map[string]float64, len(v))
			} // reset map
			testUnmarshalErr(rv4i(v33v2), bs33, h, t, "dec-map-v33-noaddr") // decode into non-addressable map value
			testDeepEqualErr(v33v1, v33v2, t, "equal-map-v33-noaddr")
		}
		if v == nil {
			v33v2 = nil
		} else {
			v33v2 = make(map[string]float64, len(v))
		} // reset map
		testUnmarshalErr(&v33v2, bs33, h, t, "dec-map-v33-p-len")
		testDeepEqualErr(v33v1, v33v2, t, "equal-map-v33-p-len")
		bs33 = testMarshalErr(&v33v1, h, t, "enc-map-v33-p")
		v33v2 = nil
		testUnmarshalErr(&v33v2, bs33, h, t, "dec-map-v33-p-nil")
		testDeepEqualErr(v33v1, v33v2, t, "equal-map-v33-p-nil")
		// ...
		if v == nil {
			v33v2 = nil
		} else {
			v33v2 = make(map[string]float64, len(v))
		} // reset map
		var v33v3, v33v4 typMapMapStringFloat64
		v33v3 = typMapMapStringFloat64(v33v1)
		v33v4 = typMapMapStringFloat64(v33v2)
		if v != nil {
			bs33 = testMarshalErr(v33v3, h, t, "enc-map-v33-custom")
			testUnmarshalErr(v33v4, bs33, h, t, "dec-map-v33-p-len")
			testDeepEqualErr(v33v3, v33v4, t, "equal-map-v33-p-len")
		}
	}
	for _, v := range []map[string]bool{nil, {}, {"some-string-3": false, "some-string-1": true}} {
		// fmt.Printf(">>>> running mammoth map v34: %v\n", v)
		var v34v1, v34v2 map[string]bool
		var bs34 []byte
		v34v1 = v
		bs34 = testMarshalErr(v34v1, h, t, "enc-map-v34")
		if v != nil {
			if v == nil {
				v34v2 = nil
			} else {
				v34v2 = make(map[string]bool, len(v))
			} // reset map
			testUnmarshalErr(v34v2, bs34, h, t, "dec-map-v34")
			testDeepEqualErr(v34v1, v34v2, t, "equal-map-v34")
			if v == nil {
				v34v2 = nil
			} else {
				v34v2 = make(map[string]bool, len(v))
			} // reset map
			testUnmarshalErr(rv4i(v34v2), bs34, h, t, "dec-map-v34-noaddr") // decode into non-addressable map value
			testDeepEqualErr(v34v1, v34v2, t, "equal-map-v34-noaddr")
		}
		if v == nil {
			v34v2 = nil
		} else {
			v34v2 = make(map[string]bool, len(v))
		} // reset map
		testUnmarshalErr(&v34v2, bs34, h, t, "dec-map-v34-p-len")
		testDeepEqualErr(v34v1, v34v2, t, "equal-map-v34-p-len")
		bs34 = testMarshalErr(&v34v1, h, t, "enc-map-v34-p")
		v34v2 = nil
		testUnmarshalErr(&v34v2, bs34, h, t, "dec-map-v34-p-nil")
		testDeepEqualErr(v34v1, v34v2, t, "equal-map-v34-p-nil")
		// ...
		if v == nil {
			v34v2 = nil
		} else {
			v34v2 = make(map[string]bool, len(v))
		} // reset map
		var v34v3, v34v4 typMapMapStringBool
		v34v3 = typMapMapStringBool(v34v1)
		v34v4 = typMapMapStringBool(v34v2)
		if v != nil {
			bs34 = testMarshalErr(v34v3, h, t, "enc-map-v34-custom")
			testUnmarshalErr(v34v4, bs34, h, t, "dec-map-v34-p-len")
			testDeepEqualErr(v34v3, v34v4, t, "equal-map-v34-p-len")
		}
	}
	for _, v := range []map[uint8]interface{}{nil, {}, {77: nil, 127: "string-is-an-interface-2"}} {
		// fmt.Printf(">>>> running mammoth map v35: %v\n", v)
		var v35v1, v35v2 map[uint8]interface{}
		var bs35 []byte
		v35v1 = v
		bs35 = testMarshalErr(v35v1, h, t, "enc-map-v35")
		if v != nil {
			if v == nil {
				v35v2 = nil
			} else {
				v35v2 = make(map[uint8]interface{}, len(v))
			} // reset map
			testUnmarshalErr(v35v2, bs35, h, t, "dec-map-v35")
			testDeepEqualErr(v35v1, v35v2, t, "equal-map-v35")
			if v == nil {
				v35v2 = nil
			} else {
				v35v2 = make(map[uint8]interface{}, len(v))
			} // reset map
			testUnmarshalErr(rv4i(v35v2), bs35, h, t, "dec-map-v35-noaddr") // decode into non-addressable map value
			testDeepEqualErr(v35v1, v35v2, t, "equal-map-v35-noaddr")
		}
		if v == nil {
			v35v2 = nil
		} else {
			v35v2 = make(map[uint8]interface{}, len(v))
		} // reset map
		testUnmarshalErr(&v35v2, bs35, h, t, "dec-map-v35-p-len")
		testDeepEqualErr(v35v1, v35v2, t, "equal-map-v35-p-len")
		bs35 = testMarshalErr(&v35v1, h, t, "enc-map-v35-p")
		v35v2 = nil
		testUnmarshalErr(&v35v2, bs35, h, t, "dec-map-v35-p-nil")
		testDeepEqualErr(v35v1, v35v2, t, "equal-map-v35-p-nil")
		// ...
		if v == nil {
			v35v2 = nil
		} else {
			v35v2 = make(map[uint8]interface{}, len(v))
		} // reset map
		var v35v3, v35v4 typMapMapUint8Intf
		v35v3 = typMapMapUint8Intf(v35v1)
		v35v4 = typMapMapUint8Intf(v35v2)
		if v != nil {
			bs35 = testMarshalErr(v35v3, h, t, "enc-map-v35-custom")
			testUnmarshalErr(v35v4, bs35, h, t, "dec-map-v35-p-len")
			testDeepEqualErr(v35v3, v35v4, t, "equal-map-v35-p-len")
		}
	}
	for _, v := range []map[uint8]string{nil, {}, {111: "", 77: "some-string-2"}} {
		// fmt.Printf(">>>> running mammoth map v36: %v\n", v)
		var v36v1, v36v2 map[uint8]string
		var bs36 []byte
		v36v1 = v
		bs36 = testMarshalErr(v36v1, h, t, "enc-map-v36")
		if v != nil {
			if v == nil {
				v36v2 = nil
			} else {
				v36v2 = make(map[uint8]string, len(v))
			} // reset map
			testUnmarshalErr(v36v2, bs36, h, t, "dec-map-v36")
			testDeepEqualErr(v36v1, v36v2, t, "equal-map-v36")
			if v == nil {
				v36v2 = nil
			} else {
				v36v2 = make(map[uint8]string, len(v))
			} // reset map
			testUnmarshalErr(rv4i(v36v2), bs36, h, t, "dec-map-v36-noaddr") // decode into non-addressable map value
			testDeepEqualErr(v36v1, v36v2, t, "equal-map-v36-noaddr")
		}
		if v == nil {
			v36v2 = nil
		} else {
			v36v2 = make(map[uint8]string, len(v))
		} // reset map
		testUnmarshalErr(&v36v2, bs36, h, t, "dec-map-v36-p-len")
		testDeepEqualErr(v36v1, v36v2, t, "equal-map-v36-p-len")
		bs36 = testMarshalErr(&v36v1, h, t, "enc-map-v36-p")
		v36v2 = nil
		testUnmarshalErr(&v36v2, bs36, h, t, "dec-map-v36-p-nil")
		testDeepEqualErr(v36v1, v36v2, t, "equal-map-v36-p-nil")
		// ...
		if v == nil {
			v36v2 = nil
		} else {
			v36v2 = make(map[uint8]string, len(v))
		} // reset map
		var v36v3, v36v4 typMapMapUint8String
		v36v3 = typMapMapUint8String(v36v1)
		v36v4 = typMapMapUint8String(v36v2)
		if v != nil {
			bs36 = testMarshalErr(v36v3, h, t, "enc-map-v36-custom")
			testUnmarshalErr(v36v4, bs36, h, t, "dec-map-v36-p-len")
			testDeepEqualErr(v36v3, v36v4, t, "equal-map-v36-p-len")
		}
	}
	for _, v := range []map[uint8][]byte{nil, {}, {127: nil, 111: []byte("some-string-2")}} {
		// fmt.Printf(">>>> running mammoth map v37: %v\n", v)
		var v37v1, v37v2 map[uint8][]byte
		var bs37 []byte
		v37v1 = v
		bs37 = testMarshalErr(v37v1, h, t, "enc-map-v37")
		if v != nil {
			if v == nil {
				v37v2 = nil
			} else {
				v37v2 = make(map[uint8][]byte, len(v))
			} // reset map
			testUnmarshalErr(v37v2, bs37, h, t, "dec-map-v37")
			testDeepEqualErr(v37v1, v37v2, t, "equal-map-v37")
			if v == nil {
				v37v2 = nil
			} else {
				v37v2 = make(map[uint8][]byte, len(v))
			} // reset map
			testUnmarshalErr(rv4i(v37v2), bs37, h, t, "dec-map-v37-noaddr") // decode into non-addressable map value
			testDeepEqualErr(v37v1, v37v2, t, "equal-map-v37-noaddr")
		}
		if v == nil {
			v37v2 = nil
		} else {
			v37v2 = make(map[uint8][]byte, len(v))
		} // reset map
		testUnmarshalErr(&v37v2, bs37, h, t, "dec-map-v37-p-len")
		testDeepEqualErr(v37v1, v37v2, t, "equal-map-v37-p-len")
		bs37 = testMarshalErr(&v37v1, h, t, "enc-map-v37-p")
		v37v2 = nil
		testUnmarshalErr(&v37v2, bs37, h, t, "dec-map-v37-p-nil")
		testDeepEqualErr(v37v1, v37v2, t, "equal-map-v37-p-nil")
		// ...
		if v == nil {
			v37v2 = nil
		} else {
			v37v2 = make(map[uint8][]byte, len(v))
		} // reset map
		var v37v3, v37v4 typMapMapUint8Bytes
		v37v3 = typMapMapUint8Bytes(v37v1)
		v37v4 = typMapMapUint8Bytes(v37v2)
		if v != nil {
			bs37 = testMarshalErr(v37v3, h, t, "enc-map-v37-custom")
			testUnmarshalErr(v37v4, bs37, h, t, "dec-map-v37-p-len")
			testDeepEqualErr(v37v3, v37v4, t, "equal-map-v37-p-len")
		}
	}
	for _, v := range []map[uint8]uint8{nil, {}, {77: 0, 127: 111}} {
		// fmt.Printf(">>>> running mammoth map v38: %v\n", v)
		var v38v1, v38v2 map[uint8]uint8
		var bs38 []byte
		v38v1 = v
		bs38 = testMarshalErr(v38v1, h, t, "enc-map-v38")
		if v != nil {
			if v == nil {
				v38v2 = nil
			} else {
				v38v2 = make(map[uint8]uint8, len(v))
			} // reset map
			testUnmarshalErr(v38v2, bs38, h, t, "dec-map-v38")
			testDeepEqualErr(v38v1, v38v2, t, "equal-map-v38")
			if v == nil {
				v38v2 = nil
			} else {
				v38v2 = make(map[uint8]uint8, len(v))
			} // reset map
			testUnmarshalErr(rv4i(v38v2), bs38, h, t, "dec-map-v38-noaddr") // decode into non-addressable map value
			testDeepEqualErr(v38v1, v38v2, t, "equal-map-v38-noaddr")
		}
		if v == nil {
			v38v2 = nil
		} else {
			v38v2 = make(map[uint8]uint8, len(v))
		} // reset map
		testUnmarshalErr(&v38v2, bs38, h, t, "dec-map-v38-p-len")
		testDeepEqualErr(v38v1, v38v2, t, "equal-map-v38-p-len")
		bs38 = testMarshalErr(&v38v1, h, t, "enc-map-v38-p")
		v38v2 = nil
		testUnmarshalErr(&v38v2, bs38, h, t, "dec-map-v38-p-nil")
		testDeepEqualErr(v38v1, v38v2, t, "equal-map-v38-p-nil")
		// ...
		if v == nil {
			v38v2 = nil
		} else {
			v38v2 = make(map[uint8]uint8, len(v))
		} // reset map
		var v38v3, v38v4 typMapMapUint8Uint8
		v38v3 = typMapMapUint8Uint8(v38v1)
		v38v4 = typMapMapUint8Uint8(v38v2)
		if v != nil {
			bs38 = testMarshalErr(v38v3, h, t, "enc-map-v38-custom")
			testUnmarshalErr(v38v4, bs38, h, t, "dec-map-v38-p-len")
			testDeepEqualErr(v38v3, v38v4, t, "equal-map-v38-p-len")
		}
	}
	for _, v := range []map[uint8]uint64{nil, {}, {77: 0, 127: 111}} {
		// fmt.Printf(">>>> running mammoth map v39: %v\n", v)
		var v39v1, v39v2 map[uint8]uint64
		var bs39 []byte
		v39v1 = v
		bs39 = testMarshalErr(v39v1, h, t, "enc-map-v39")
		if v != nil {
			if v == nil {
				v39v2 = nil
			} else {
				v39v2 = make(map[uint8]uint64, len(v))
			} // reset map
			testUnmarshalErr(v39v2, bs39, h, t, "dec-map-v39")
			testDeepEqualErr(v39v1, v39v2, t, "equal-map-v39")
			if v == nil {
				v39v2 = nil
			} else {
				v39v2 = make(map[uint8]uint64, len(v))
			} // reset map
			testUnmarshalErr(rv4i(v39v2), bs39, h, t, "dec-map-v39-noaddr") // decode into non-addressable map value
			testDeepEqualErr(v39v1, v39v2, t, "equal-map-v39-noaddr")
		}
		if v == nil {
			v39v2 = nil
		} else {
			v39v2 = make(map[uint8]uint64, len(v))
		} // reset map
		testUnmarshalErr(&v39v2, bs39, h, t, "dec-map-v39-p-len")
		testDeepEqualErr(v39v1, v39v2, t, "equal-map-v39-p-len")
		bs39 = testMarshalErr(&v39v1, h, t, "enc-map-v39-p")
		v39v2 = nil
		testUnmarshalErr(&v39v2, bs39, h, t, "dec-map-v39-p-nil")
		testDeepEqualErr(v39v1, v39v2, t, "equal-map-v39-p-nil")
		// ...
		if v == nil {
			v39v2 = nil
		} else {
			v39v2 = make(map[uint8]uint64, len(v))
		} // reset map
		var v39v3, v39v4 typMapMapUint8Uint64
		v39v3 = typMapMapUint8Uint64(v39v1)
		v39v4 = typMapMapUint8Uint64(v39v2)
		if v != nil {
			bs39 = testMarshalErr(v39v3, h, t, "enc-map-v39-custom")
			testUnmarshalErr(v39v4, bs39, h, t, "dec-map-v39-p-len")
			testDeepEqualErr(v39v3, v39v4, t, "equal-map-v39-p-len")
		}
	}
	for _, v := range []map[uint8]int{nil, {}, {77: 0, 127: 111}} {
		// fmt.Printf(">>>> running mammoth map v40: %v\n", v)
		var v40v1, v40v2 map[uint8]int
		var bs40 []byte
		v40v1 = v
		bs40 = testMarshalErr(v40v1, h, t, "enc-map-v40")
		if v != nil {
			if v == nil {
				v40v2 = nil
			} else {
				v40v2 = make(map[uint8]int, len(v))
			} // reset map
			testUnmarshalErr(v40v2, bs40, h, t, "dec-map-v40")
			testDeepEqualErr(v40v1, v40v2, t, "equal-map-v40")
			if v == nil {
				v40v2 = nil
			} else {
				v40v2 = make(map[uint8]int, len(v))
			} // reset map
			testUnmarshalErr(rv4i(v40v2), bs40, h, t, "dec-map-v40-noaddr") // decode into non-addressable map value
			testDeepEqualErr(v40v1, v40v2, t, "equal-map-v40-noaddr")
		}
		if v == nil {
			v40v2 = nil
		} else {
			v40v2 = make(map[uint8]int, len(v))
		} // reset map
		testUnmarshalErr(&v40v2, bs40, h, t, "dec-map-v40-p-len")
		testDeepEqualErr(v40v1, v40v2, t, "equal-map-v40-p-len")
		bs40 = testMarshalErr(&v40v1, h, t, "enc-map-v40-p")
		v40v2 = nil
		testUnmarshalErr(&v40v2, bs40, h, t, "dec-map-v40-p-nil")
		testDeepEqualErr(v40v1, v40v2, t, "equal-map-v40-p-nil")
		// ...
		if v == nil {
			v40v2 = nil
		} else {
			v40v2 = make(map[uint8]int, len(v))
		} // reset map
		var v40v3, v40v4 typMapMapUint8Int
		v40v3 = typMapMapUint8Int(v40v1)
		v40v4 = typMapMapUint8Int(v40v2)
		if v != nil {
			bs40 = testMarshalErr(v40v3, h, t, "enc-map-v40-custom")
			testUnmarshalErr(v40v4, bs40, h, t, "dec-map-v40-p-len")
			testDeepEqualErr(v40v3, v40v4, t, "equal-map-v40-p-len")
		}
	}
	for _, v := range []map[uint8]int64{nil, {}, {77: 0, 127: 111}} {
		// fmt.Printf(">>>> running mammoth map v41: %v\n", v)
		var v41v1, v41v2 map[uint8]int64
		var bs41 []byte
		v41v1 = v
		bs41 = testMarshalErr(v41v1, h, t, "enc-map-v41")
		if v != nil {
			if v == nil {
				v41v2 = nil
			} else {
				v41v2 = make(map[uint8]int64, len(v))
			} // reset map
			testUnmarshalErr(v41v2, bs41, h, t, "dec-map-v41")
			testDeepEqualErr(v41v1, v41v2, t, "equal-map-v41")
			if v == nil {
				v41v2 = nil
			} else {
				v41v2 = make(map[uint8]int64, len(v))
			} // reset map
			testUnmarshalErr(rv4i(v41v2), bs41, h, t, "dec-map-v41-noaddr") // decode into non-addressable map value
			testDeepEqualErr(v41v1, v41v2, t, "equal-map-v41-noaddr")
		}
		if v == nil {
			v41v2 = nil
		} else {
			v41v2 = make(map[uint8]int64, len(v))
		} // reset map
		testUnmarshalErr(&v41v2, bs41, h, t, "dec-map-v41-p-len")
		testDeepEqualErr(v41v1, v41v2, t, "equal-map-v41-p-len")
		bs41 = testMarshalErr(&v41v1, h, t, "enc-map-v41-p")
		v41v2 = nil
		testUnmarshalErr(&v41v2, bs41, h, t, "dec-map-v41-p-nil")
		testDeepEqualErr(v41v1, v41v2, t, "equal-map-v41-p-nil")
		// ...
		if v == nil {
			v41v2 = nil
		} else {
			v41v2 = make(map[uint8]int64, len(v))
		} // reset map
		var v41v3, v41v4 typMapMapUint8Int64
		v41v3 = typMapMapUint8Int64(v41v1)
		v41v4 = typMapMapUint8Int64(v41v2)
		if v != nil {
			bs41 = testMarshalErr(v41v3, h, t, "enc-map-v41-custom")
			testUnmarshalErr(v41v4, bs41, h, t, "dec-map-v41-p-len")
			testDeepEqualErr(v41v3, v41v4, t, "equal-map-v41-p-len")
		}
	}
	for _, v := range []map[uint8]float64{nil, {}, {77: 0, 127: 22.2}} {
		// fmt.Printf(">>>> running mammoth map v42: %v\n", v)
		var v42v1, v42v2 map[uint8]float64
		var bs42 []byte
		v42v1 = v
		bs42 = testMarshalErr(v42v1, h, t, "enc-map-v42")
		if v != nil {
			if v == nil {
				v42v2 = nil
			} else {
				v42v2 = make(map[uint8]float64, len(v))
			} // reset map
			testUnmarshalErr(v42v2, bs42, h, t, "dec-map-v42")
			testDeepEqualErr(v42v1, v42v2, t, "equal-map-v42")
			if v == nil {
				v42v2 = nil
			} else {
				v42v2 = make(map[uint8]float64, len(v))
			} // reset map
			testUnmarshalErr(rv4i(v42v2), bs42, h, t, "dec-map-v42-noaddr") // decode into non-addressable map value
			testDeepEqualErr(v42v1, v42v2, t, "equal-map-v42-noaddr")
		}
		if v == nil {
			v42v2 = nil
		} else {
			v42v2 = make(map[uint8]float64, len(v))
		} // reset map
		testUnmarshalErr(&v42v2, bs42, h, t, "dec-map-v42-p-len")
		testDeepEqualErr(v42v1, v42v2, t, "equal-map-v42-p-len")
		bs42 = testMarshalErr(&v42v1, h, t, "enc-map-v42-p")
		v42v2 = nil
		testUnmarshalErr(&v42v2, bs42, h, t, "dec-map-v42-p-nil")
		testDeepEqualErr(v42v1, v42v2, t, "equal-map-v42-p-nil")
		// ...
		if v == nil {
			v42v2 = nil
		} else {
			v42v2 = make(map[uint8]float64, len(v))
		} // reset map
		var v42v3, v42v4 typMapMapUint8Float64
		v42v3 = typMapMapUint8Float64(v42v1)
		v42v4 = typMapMapUint8Float64(v42v2)
		if v != nil {
			bs42 = testMarshalErr(v42v3, h, t, "enc-map-v42-custom")
			testUnmarshalErr(v42v4, bs42, h, t, "dec-map-v42-p-len")
			testDeepEqualErr(v42v3, v42v4, t, "equal-map-v42-p-len")
		}
	}
	for _, v := range []map[uint8]bool{nil, {}, {111: false, 77: false}} {
		// fmt.Printf(">>>> running mammoth map v43: %v\n", v)
		var v43v1, v43v2 map[uint8]bool
		var bs43 []byte
		v43v1 = v
		bs43 = testMarshalErr(v43v1, h, t, "enc-map-v43")
		if v != nil {
			if v == nil {
				v43v2 = nil
			} else {
				v43v2 = make(map[uint8]bool, len(v))
			} // reset map
			testUnmarshalErr(v43v2, bs43, h, t, "dec-map-v43")
			testDeepEqualErr(v43v1, v43v2, t, "equal-map-v43")
			if v == nil {
				v43v2 = nil
			} else {
				v43v2 = make(map[uint8]bool, len(v))
			} // reset map
			testUnmarshalErr(rv4i(v43v2), bs43, h, t, "dec-map-v43-noaddr") // decode into non-addressable map value
			testDeepEqualErr(v43v1, v43v2, t, "equal-map-v43-noaddr")
		}
		if v == nil {
			v43v2 = nil
		} else {
			v43v2 = make(map[uint8]bool, len(v))
		} // reset map
		testUnmarshalErr(&v43v2, bs43, h, t, "dec-map-v43-p-len")
		testDeepEqualErr(v43v1, v43v2, t, "equal-map-v43-p-len")
		bs43 = testMarshalErr(&v43v1, h, t, "enc-map-v43-p")
		v43v2 = nil
		testUnmarshalErr(&v43v2, bs43, h, t, "dec-map-v43-p-nil")
		testDeepEqualErr(v43v1, v43v2, t, "equal-map-v43-p-nil")
		// ...
		if v == nil {
			v43v2 = nil
		} else {
			v43v2 = make(map[uint8]bool, len(v))
		} // reset map
		var v43v3, v43v4 typMapMapUint8Bool
		v43v3 = typMapMapUint8Bool(v43v1)
		v43v4 = typMapMapUint8Bool(v43v2)
		if v != nil {
			bs43 = testMarshalErr(v43v3, h, t, "enc-map-v43-custom")
			testUnmarshalErr(v43v4, bs43, h, t, "dec-map-v43-p-len")
			testDeepEqualErr(v43v3, v43v4, t, "equal-map-v43-p-len")
		}
	}
	for _, v := range []map[uint64]interface{}{nil, {}, {127: nil, 111: "string-is-an-interface-3"}} {
		// fmt.Printf(">>>> running mammoth map v44: %v\n", v)
		var v44v1, v44v2 map[uint64]interface{}
		var bs44 []byte
		v44v1 = v
		bs44 = testMarshalErr(v44v1, h, t, "enc-map-v44")
		if v != nil {
			if v == nil {
				v44v2 = nil
			} else {
				v44v2 = make(map[uint64]interface{}, len(v))
			} // reset map
			testUnmarshalErr(v44v2, bs44, h, t, "dec-map-v44")
			testDeepEqualErr(v44v1, v44v2, t, "equal-map-v44")
			if v == nil {
				v44v2 = nil
			} else {
				v44v2 = make(map[uint64]interface{}, len(v))
			} // reset map
			testUnmarshalErr(rv4i(v44v2), bs44, h, t, "dec-map-v44-noaddr") // decode into non-addressable map value
			testDeepEqualErr(v44v1, v44v2, t, "equal-map-v44-noaddr")
		}
		if v == nil {
			v44v2 = nil
		} else {
			v44v2 = make(map[uint64]interface{}, len(v))
		} // reset map
		testUnmarshalErr(&v44v2, bs44, h, t, "dec-map-v44-p-len")
		testDeepEqualErr(v44v1, v44v2, t, "equal-map-v44-p-len")
		bs44 = testMarshalErr(&v44v1, h, t, "enc-map-v44-p")
		v44v2 = nil
		testUnmarshalErr(&v44v2, bs44, h, t, "dec-map-v44-p-nil")
		testDeepEqualErr(v44v1, v44v2, t, "equal-map-v44-p-nil")
		// ...
		if v == nil {
			v44v2 = nil
		} else {
			v44v2 = make(map[uint64]interface{}, len(v))
		} // reset map
		var v44v3, v44v4 typMapMapUint64Intf
		v44v3 = typMapMapUint64Intf(v44v1)
		v44v4 = typMapMapUint64Intf(v44v2)
		if v != nil {
			bs44 = testMarshalErr(v44v3, h, t, "enc-map-v44-custom")
			testUnmarshalErr(v44v4, bs44, h, t, "dec-map-v44-p-len")
			testDeepEqualErr(v44v3, v44v4, t, "equal-map-v44-p-len")
		}
	}
	for _, v := range []map[uint64]string{nil, {}, {77: "", 127: "some-string-3"}} {
		// fmt.Printf(">>>> running mammoth map v45: %v\n", v)
		var v45v1, v45v2 map[uint64]string
		var bs45 []byte
		v45v1 = v
		bs45 = testMarshalErr(v45v1, h, t, "enc-map-v45")
		if v != nil {
			if v == nil {
				v45v2 = nil
			} else {
				v45v2 = make(map[uint64]string, len(v))
			} // reset map
			testUnmarshalErr(v45v2, bs45, h, t, "dec-map-v45")
			testDeepEqualErr(v45v1, v45v2, t, "equal-map-v45")
			if v == nil {
				v45v2 = nil
			} else {
				v45v2 = make(map[uint64]string, len(v))
			} // reset map
			testUnmarshalErr(rv4i(v45v2), bs45, h, t, "dec-map-v45-noaddr") // decode into non-addressable map value
			testDeepEqualErr(v45v1, v45v2, t, "equal-map-v45-noaddr")
		}
		if v == nil {
			v45v2 = nil
		} else {
			v45v2 = make(map[uint64]string, len(v))
		} // reset map
		testUnmarshalErr(&v45v2, bs45, h, t, "dec-map-v45-p-len")
		testDeepEqualErr(v45v1, v45v2, t, "equal-map-v45-p-len")
		bs45 = testMarshalErr(&v45v1, h, t, "enc-map-v45-p")
		v45v2 = nil
		testUnmarshalErr(&v45v2, bs45, h, t, "dec-map-v45-p-nil")
		testDeepEqualErr(v45v1, v45v2, t, "equal-map-v45-p-nil")
		// ...
		if v == nil {
			v45v2 = nil
		} else {
			v45v2 = make(map[uint64]string, len(v))
		} // reset map
		var v45v3, v45v4 typMapMapUint64String
		v45v3 = typMapMapUint64String(v45v1)
		v45v4 = typMapMapUint64String(v45v2)
		if v != nil {
			bs45 = testMarshalErr(v45v3, h, t, "enc-map-v45-custom")
			testUnmarshalErr(v45v4, bs45, h, t, "dec-map-v45-p-len")
			testDeepEqualErr(v45v3, v45v4, t, "equal-map-v45-p-len")
		}
	}
	for _, v := range []map[uint64][]byte{nil, {}, {111: nil, 77: []byte("some-string-3")}} {
		// fmt.Printf(">>>> running mammoth map v46: %v\n", v)
		var v46v1, v46v2 map[uint64][]byte
		var bs46 []byte
		v46v1 = v
		bs46 = testMarshalErr(v46v1, h, t, "enc-map-v46")
		if v != nil {
			if v == nil {
				v46v2 = nil
			} else {
				v46v2 = make(map[uint64][]byte, len(v))
			} // reset map
			testUnmarshalErr(v46v2, bs46, h, t, "dec-map-v46")
			testDeepEqualErr(v46v1, v46v2, t, "equal-map-v46")
			if v == nil {
				v46v2 = nil
			} else {
				v46v2 = make(map[uint64][]byte, len(v))
			} // reset map
			testUnmarshalErr(rv4i(v46v2), bs46, h, t, "dec-map-v46-noaddr") // decode into non-addressable map value
			testDeepEqualErr(v46v1, v46v2, t, "equal-map-v46-noaddr")
		}
		if v == nil {
			v46v2 = nil
		} else {
			v46v2 = make(map[uint64][]byte, len(v))
		} // reset map
		testUnmarshalErr(&v46v2, bs46, h, t, "dec-map-v46-p-len")
		testDeepEqualErr(v46v1, v46v2, t, "equal-map-v46-p-len")
		bs46 = testMarshalErr(&v46v1, h, t, "enc-map-v46-p")
		v46v2 = nil
		testUnmarshalErr(&v46v2, bs46, h, t, "dec-map-v46-p-nil")
		testDeepEqualErr(v46v1, v46v2, t, "equal-map-v46-p-nil")
		// ...
		if v == nil {
			v46v2 = nil
		} else {
			v46v2 = make(map[uint64][]byte, len(v))
		} // reset map
		var v46v3, v46v4 typMapMapUint64Bytes
		v46v3 = typMapMapUint64Bytes(v46v1)
		v46v4 = typMapMapUint64Bytes(v46v2)
		if v != nil {
			bs46 = testMarshalErr(v46v3, h, t, "enc-map-v46-custom")
			testUnmarshalErr(v46v4, bs46, h, t, "dec-map-v46-p-len")
			testDeepEqualErr(v46v3, v46v4, t, "equal-map-v46-p-len")
		}
	}
	for _, v := range []map[uint64]uint8{nil, {}, {127: 0, 111: 77}} {
		// fmt.Printf(">>>> running mammoth map v47: %v\n", v)
		var v47v1, v47v2 map[uint64]uint8
		var bs47 []byte
		v47v1 = v
		bs47 = testMarshalErr(v47v1, h, t, "enc-map-v47")
		if v != nil {
			if v == nil {
				v47v2 = nil
			} else {
				v47v2 = make(map[uint64]uint8, len(v))
			} // reset map
			testUnmarshalErr(v47v2, bs47, h, t, "dec-map-v47")
			testDeepEqualErr(v47v1, v47v2, t, "equal-map-v47")
			if v == nil {
				v47v2 = nil
			} else {
				v47v2 = make(map[uint64]uint8, len(v))
			} // reset map
			testUnmarshalErr(rv4i(v47v2), bs47, h, t, "dec-map-v47-noaddr") // decode into non-addressable map value
			testDeepEqualErr(v47v1, v47v2, t, "equal-map-v47-noaddr")
		}
		if v == nil {
			v47v2 = nil
		} else {
			v47v2 = make(map[uint64]uint8, len(v))
		} // reset map
		testUnmarshalErr(&v47v2, bs47, h, t, "dec-map-v47-p-len")
		testDeepEqualErr(v47v1, v47v2, t, "equal-map-v47-p-len")
		bs47 = testMarshalErr(&v47v1, h, t, "enc-map-v47-p")
		v47v2 = nil
		testUnmarshalErr(&v47v2, bs47, h, t, "dec-map-v47-p-nil")
		testDeepEqualErr(v47v1, v47v2, t, "equal-map-v47-p-nil")
		// ...
		if v == nil {
			v47v2 = nil
		} else {
			v47v2 = make(map[uint64]uint8, len(v))
		} // reset map
		var v47v3, v47v4 typMapMapUint64Uint8
		v47v3 = typMapMapUint64Uint8(v47v1)
		v47v4 = typMapMapUint64Uint8(v47v2)
		if v != nil {
			bs47 = testMarshalErr(v47v3, h, t, "enc-map-v47-custom")
			testUnmarshalErr(v47v4, bs47, h, t, "dec-map-v47-p-len")
			testDeepEqualErr(v47v3, v47v4, t, "equal-map-v47-p-len")
		}
	}
	for _, v := range []map[uint64]uint64{nil, {}, {127: 0, 111: 77}} {
		// fmt.Printf(">>>> running mammoth map v48: %v\n", v)
		var v48v1, v48v2 map[uint64]uint64
		var bs48 []byte
		v48v1 = v
		bs48 = testMarshalErr(v48v1, h, t, "enc-map-v48")
		if v != nil {
			if v == nil {
				v48v2 = nil
			} else {
				v48v2 = make(map[uint64]uint64, len(v))
			} // reset map
			testUnmarshalErr(v48v2, bs48, h, t, "dec-map-v48")
			testDeepEqualErr(v48v1, v48v2, t, "equal-map-v48")
			if v == nil {
				v48v2 = nil
			} else {
				v48v2 = make(map[uint64]uint64, len(v))
			} // reset map
			testUnmarshalErr(rv4i(v48v2), bs48, h, t, "dec-map-v48-noaddr") // decode into non-addressable map value
			testDeepEqualErr(v48v1, v48v2, t, "equal-map-v48-noaddr")
		}
		if v == nil {
			v48v2 = nil
		} else {
			v48v2 = make(map[uint64]uint64, len(v))
		} // reset map
		testUnmarshalErr(&v48v2, bs48, h, t, "dec-map-v48-p-len")
		testDeepEqualErr(v48v1, v48v2, t, "equal-map-v48-p-len")
		bs48 = testMarshalErr(&v48v1, h, t, "enc-map-v48-p")
		v48v2 = nil
		testUnmarshalErr(&v48v2, bs48, h, t, "dec-map-v48-p-nil")
		testDeepEqualErr(v48v1, v48v2, t, "equal-map-v48-p-nil")
		// ...
		if v == nil {
			v48v2 = nil
		} else {
			v48v2 = make(map[uint64]uint64, len(v))
		} // reset map
		var v48v3, v48v4 typMapMapUint64Uint64
		v48v3 = typMapMapUint64Uint64(v48v1)
		v48v4 = typMapMapUint64Uint64(v48v2)
		if v != nil {
			bs48 = testMarshalErr(v48v3, h, t, "enc-map-v48-custom")
			testUnmarshalErr(v48v4, bs48, h, t, "dec-map-v48-p-len")
			testDeepEqualErr(v48v3, v48v4, t, "equal-map-v48-p-len")
		}
	}
	for _, v := range []map[uint64]int{nil, {}, {127: 0, 111: 77}} {
		// fmt.Printf(">>>> running mammoth map v49: %v\n", v)
		var v49v1, v49v2 map[uint64]int
		var bs49 []byte
		v49v1 = v
		bs49 = testMarshalErr(v49v1, h, t, "enc-map-v49")
		if v != nil {
			if v == nil {
				v49v2 = nil
			} else {
				v49v2 = make(map[uint64]int, len(v))
			} // reset map
			testUnmarshalErr(v49v2, bs49, h, t, "dec-map-v49")
			testDeepEqualErr(v49v1, v49v2, t, "equal-map-v49")
			if v == nil {
				v49v2 = nil
			} else {
				v49v2 = make(map[uint64]int, len(v))
			} // reset map
			testUnmarshalErr(rv4i(v49v2), bs49, h, t, "dec-map-v49-noaddr") // decode into non-addressable map value
			testDeepEqualErr(v49v1, v49v2, t, "equal-map-v49-noaddr")
		}
		if v == nil {
			v49v2 = nil
		} else {
			v49v2 = make(map[uint64]int, len(v))
		} // reset map
		testUnmarshalErr(&v49v2, bs49, h, t, "dec-map-v49-p-len")
		testDeepEqualErr(v49v1, v49v2, t, "equal-map-v49-p-len")
		bs49 = testMarshalErr(&v49v1, h, t, "enc-map-v49-p")
		v49v2 = nil
		testUnmarshalErr(&v49v2, bs49, h, t, "dec-map-v49-p-nil")
		testDeepEqualErr(v49v1, v49v2, t, "equal-map-v49-p-nil")
		// ...
		if v == nil {
			v49v2 = nil
		} else {
			v49v2 = make(map[uint64]int, len(v))
		} // reset map
		var v49v3, v49v4 typMapMapUint64Int
		v49v3 = typMapMapUint64Int(v49v1)
		v49v4 = typMapMapUint64Int(v49v2)
		if v != nil {
			bs49 = testMarshalErr(v49v3, h, t, "enc-map-v49-custom")
			testUnmarshalErr(v49v4, bs49, h, t, "dec-map-v49-p-len")
			testDeepEqualErr(v49v3, v49v4, t, "equal-map-v49-p-len")
		}
	}
	for _, v := range []map[uint64]int64{nil, {}, {127: 0, 111: 77}} {
		// fmt.Printf(">>>> running mammoth map v50: %v\n", v)
		var v50v1, v50v2 map[uint64]int64
		var bs50 []byte
		v50v1 = v
		bs50 = testMarshalErr(v50v1, h, t, "enc-map-v50")
		if v != nil {
			if v == nil {
				v50v2 = nil
			} else {
				v50v2 = make(map[uint64]int64, len(v))
			} // reset map
			testUnmarshalErr(v50v2, bs50, h, t, "dec-map-v50")
			testDeepEqualErr(v50v1, v50v2, t, "equal-map-v50")
			if v == nil {
				v50v2 = nil
			} else {
				v50v2 = make(map[uint64]int64, len(v))
			} // reset map
			testUnmarshalErr(rv4i(v50v2), bs50, h, t, "dec-map-v50-noaddr") // decode into non-addressable map value
			testDeepEqualErr(v50v1, v50v2, t, "equal-map-v50-noaddr")
		}
		if v == nil {
			v50v2 = nil
		} else {
			v50v2 = make(map[uint64]int64, len(v))
		} // reset map
		testUnmarshalErr(&v50v2, bs50, h, t, "dec-map-v50-p-len")
		testDeepEqualErr(v50v1, v50v2, t, "equal-map-v50-p-len")
		bs50 = testMarshalErr(&v50v1, h, t, "enc-map-v50-p")
		v50v2 = nil
		testUnmarshalErr(&v50v2, bs50, h, t, "dec-map-v50-p-nil")
		testDeepEqualErr(v50v1, v50v2, t, "equal-map-v50-p-nil")
		// ...
		if v == nil {
			v50v2 = nil
		} else {
			v50v2 = make(map[uint64]int64, len(v))
		} // reset map
		var v50v3, v50v4 typMapMapUint64Int64
		v50v3 = typMapMapUint64Int64(v50v1)
		v50v4 = typMapMapUint64Int64(v50v2)
		if v != nil {
			bs50 = testMarshalErr(v50v3, h, t, "enc-map-v50-custom")
			testUnmarshalErr(v50v4, bs50, h, t, "dec-map-v50-p-len")
			testDeepEqualErr(v50v3, v50v4, t, "equal-map-v50-p-len")
		}
	}
	for _, v := range []map[uint64]float64{nil, {}, {127: 0, 111: 33.3e3}} {
		// fmt.Printf(">>>> running mammoth map v51: %v\n", v)
		var v51v1, v51v2 map[uint64]float64
		var bs51 []byte
		v51v1 = v
		bs51 = testMarshalErr(v51v1, h, t, "enc-map-v51")
		if v != nil {
			if v == nil {
				v51v2 = nil
			} else {
				v51v2 = make(map[uint64]float64, len(v))
			} // reset map
			testUnmarshalErr(v51v2, bs51, h, t, "dec-map-v51")
			testDeepEqualErr(v51v1, v51v2, t, "equal-map-v51")
			if v == nil {
				v51v2 = nil
			} else {
				v51v2 = make(map[uint64]float64, len(v))
			} // reset map
			testUnmarshalErr(rv4i(v51v2), bs51, h, t, "dec-map-v51-noaddr") // decode into non-addressable map value
			testDeepEqualErr(v51v1, v51v2, t, "equal-map-v51-noaddr")
		}
		if v == nil {
			v51v2 = nil
		} else {
			v51v2 = make(map[uint64]float64, len(v))
		} // reset map
		testUnmarshalErr(&v51v2, bs51, h, t, "dec-map-v51-p-len")
		testDeepEqualErr(v51v1, v51v2, t, "equal-map-v51-p-len")
		bs51 = testMarshalErr(&v51v1, h, t, "enc-map-v51-p")
		v51v2 = nil
		testUnmarshalErr(&v51v2, bs51, h, t, "dec-map-v51-p-nil")
		testDeepEqualErr(v51v1, v51v2, t, "equal-map-v51-p-nil")
		// ...
		if v == nil {
			v51v2 = nil
		} else {
			v51v2 = make(map[uint64]float64, len(v))
		} // reset map
		var v51v3, v51v4 typMapMapUint64Float64
		v51v3 = typMapMapUint64Float64(v51v1)
		v51v4 = typMapMapUint64Float64(v51v2)
		if v != nil {
			bs51 = testMarshalErr(v51v3, h, t, "enc-map-v51-custom")
			testUnmarshalErr(v51v4, bs51, h, t, "dec-map-v51-p-len")
			testDeepEqualErr(v51v3, v51v4, t, "equal-map-v51-p-len")
		}
	}
	for _, v := range []map[uint64]bool{nil, {}, {77: false, 127: true}} {
		// fmt.Printf(">>>> running mammoth map v52: %v\n", v)
		var v52v1, v52v2 map[uint64]bool
		var bs52 []byte
		v52v1 = v
		bs52 = testMarshalErr(v52v1, h, t, "enc-map-v52")
		if v != nil {
			if v == nil {
				v52v2 = nil
			} else {
				v52v2 = make(map[uint64]bool, len(v))
			} // reset map
			testUnmarshalErr(v52v2, bs52, h, t, "dec-map-v52")
			testDeepEqualErr(v52v1, v52v2, t, "equal-map-v52")
			if v == nil {
				v52v2 = nil
			} else {
				v52v2 = make(map[uint64]bool, len(v))
			} // reset map
			testUnmarshalErr(rv4i(v52v2), bs52, h, t, "dec-map-v52-noaddr") // decode into non-addressable map value
			testDeepEqualErr(v52v1, v52v2, t, "equal-map-v52-noaddr")
		}
		if v == nil {
			v52v2 = nil
		} else {
			v52v2 = make(map[uint64]bool, len(v))
		} // reset map
		testUnmarshalErr(&v52v2, bs52, h, t, "dec-map-v52-p-len")
		testDeepEqualErr(v52v1, v52v2, t, "equal-map-v52-p-len")
		bs52 = testMarshalErr(&v52v1, h, t, "enc-map-v52-p")
		v52v2 = nil
		testUnmarshalErr(&v52v2, bs52, h, t, "dec-map-v52-p-nil")
		testDeepEqualErr(v52v1, v52v2, t, "equal-map-v52-p-nil")
		// ...
		if v == nil {
			v52v2 = nil
		} else {
			v52v2 = make(map[uint64]bool, len(v))
		} // reset map
		var v52v3, v52v4 typMapMapUint64Bool
		v52v3 = typMapMapUint64Bool(v52v1)
		v52v4 = typMapMapUint64Bool(v52v2)
		if v != nil {
			bs52 = testMarshalErr(v52v3, h, t, "enc-map-v52-custom")
			testUnmarshalErr(v52v4, bs52, h, t, "dec-map-v52-p-len")
			testDeepEqualErr(v52v3, v52v4, t, "equal-map-v52-p-len")
		}
	}
	for _, v := range []map[int]interface{}{nil, {}, {111: nil, 77: "string-is-an-interface-1"}} {
		// fmt.Printf(">>>> running mammoth map v53: %v\n", v)
		var v53v1, v53v2 map[int]interface{}
		var bs53 []byte
		v53v1 = v
		bs53 = testMarshalErr(v53v1, h, t, "enc-map-v53")
		if v != nil {
			if v == nil {
				v53v2 = nil
			} else {
				v53v2 = make(map[int]interface{}, len(v))
			} // reset map
			testUnmarshalErr(v53v2, bs53, h, t, "dec-map-v53")
			testDeepEqualErr(v53v1, v53v2, t, "equal-map-v53")
			if v == nil {
				v53v2 = nil
			} else {
				v53v2 = make(map[int]interface{}, len(v))
			} // reset map
			testUnmarshalErr(rv4i(v53v2), bs53, h, t, "dec-map-v53-noaddr") // decode into non-addressable map value
			testDeepEqualErr(v53v1, v53v2, t, "equal-map-v53-noaddr")
		}
		if v == nil {
			v53v2 = nil
		} else {
			v53v2 = make(map[int]interface{}, len(v))
		} // reset map
		testUnmarshalErr(&v53v2, bs53, h, t, "dec-map-v53-p-len")
		testDeepEqualErr(v53v1, v53v2, t, "equal-map-v53-p-len")
		bs53 = testMarshalErr(&v53v1, h, t, "enc-map-v53-p")
		v53v2 = nil
		testUnmarshalErr(&v53v2, bs53, h, t, "dec-map-v53-p-nil")
		testDeepEqualErr(v53v1, v53v2, t, "equal-map-v53-p-nil")
		// ...
		if v == nil {
			v53v2 = nil
		} else {
			v53v2 = make(map[int]interface{}, len(v))
		} // reset map
		var v53v3, v53v4 typMapMapIntIntf
		v53v3 = typMapMapIntIntf(v53v1)
		v53v4 = typMapMapIntIntf(v53v2)
		if v != nil {
			bs53 = testMarshalErr(v53v3, h, t, "enc-map-v53-custom")
			testUnmarshalErr(v53v4, bs53, h, t, "dec-map-v53-p-len")
			testDeepEqualErr(v53v3, v53v4, t, "equal-map-v53-p-len")
		}
	}
	for _, v := range []map[int]string{nil, {}, {127: "", 111: "some-string-1"}} {
		// fmt.Printf(">>>> running mammoth map v54: %v\n", v)
		var v54v1, v54v2 map[int]string
		var bs54 []byte
		v54v1 = v
		bs54 = testMarshalErr(v54v1, h, t, "enc-map-v54")
		if v != nil {
			if v == nil {
				v54v2 = nil
			} else {
				v54v2 = make(map[int]string, len(v))
			} // reset map
			testUnmarshalErr(v54v2, bs54, h, t, "dec-map-v54")
			testDeepEqualErr(v54v1, v54v2, t, "equal-map-v54")
			if v == nil {
				v54v2 = nil
			} else {
				v54v2 = make(map[int]string, len(v))
			} // reset map
			testUnmarshalErr(rv4i(v54v2), bs54, h, t, "dec-map-v54-noaddr") // decode into non-addressable map value
			testDeepEqualErr(v54v1, v54v2, t, "equal-map-v54-noaddr")
		}
		if v == nil {
			v54v2 = nil
		} else {
			v54v2 = make(map[int]string, len(v))
		} // reset map
		testUnmarshalErr(&v54v2, bs54, h, t, "dec-map-v54-p-len")
		testDeepEqualErr(v54v1, v54v2, t, "equal-map-v54-p-len")
		bs54 = testMarshalErr(&v54v1, h, t, "enc-map-v54-p")
		v54v2 = nil
		testUnmarshalErr(&v54v2, bs54, h, t, "dec-map-v54-p-nil")
		testDeepEqualErr(v54v1, v54v2, t, "equal-map-v54-p-nil")
		// ...
		if v == nil {
			v54v2 = nil
		} else {
			v54v2 = make(map[int]string, len(v))
		} // reset map
		var v54v3, v54v4 typMapMapIntString
		v54v3 = typMapMapIntString(v54v1)
		v54v4 = typMapMapIntString(v54v2)
		if v != nil {
			bs54 = testMarshalErr(v54v3, h, t, "enc-map-v54-custom")
			testUnmarshalErr(v54v4, bs54, h, t, "dec-map-v54-p-len")
			testDeepEqualErr(v54v3, v54v4, t, "equal-map-v54-p-len")
		}
	}
	for _, v := range []map[int][]byte{nil, {}, {77: nil, 127: []byte("some-string-1")}} {
		// fmt.Printf(">>>> running mammoth map v55: %v\n", v)
		var v55v1, v55v2 map[int][]byte
		var bs55 []byte
		v55v1 = v
		bs55 = testMarshalErr(v55v1, h, t, "enc-map-v55")
		if v != nil {
			if v == nil {
				v55v2 = nil
			} else {
				v55v2 = make(map[int][]byte, len(v))
			} // reset map
			testUnmarshalErr(v55v2, bs55, h, t, "dec-map-v55")
			testDeepEqualErr(v55v1, v55v2, t, "equal-map-v55")
			if v == nil {
				v55v2 = nil
			} else {
				v55v2 = make(map[int][]byte, len(v))
			} // reset map
			testUnmarshalErr(rv4i(v55v2), bs55, h, t, "dec-map-v55-noaddr") // decode into non-addressable map value
			testDeepEqualErr(v55v1, v55v2, t, "equal-map-v55-noaddr")
		}
		if v == nil {
			v55v2 = nil
		} else {
			v55v2 = make(map[int][]byte, len(v))
		} // reset map
		testUnmarshalErr(&v55v2, bs55, h, t, "dec-map-v55-p-len")
		testDeepEqualErr(v55v1, v55v2, t, "equal-map-v55-p-len")
		bs55 = testMarshalErr(&v55v1, h, t, "enc-map-v55-p")
		v55v2 = nil
		testUnmarshalErr(&v55v2, bs55, h, t, "dec-map-v55-p-nil")
		testDeepEqualErr(v55v1, v55v2, t, "equal-map-v55-p-nil")
		// ...
		if v == nil {
			v55v2 = nil
		} else {
			v55v2 = make(map[int][]byte, len(v))
		} // reset map
		var v55v3, v55v4 typMapMapIntBytes
		v55v3 = typMapMapIntBytes(v55v1)
		v55v4 = typMapMapIntBytes(v55v2)
		if v != nil {
			bs55 = testMarshalErr(v55v3, h, t, "enc-map-v55-custom")
			testUnmarshalErr(v55v4, bs55, h, t, "dec-map-v55-p-len")
			testDeepEqualErr(v55v3, v55v4, t, "equal-map-v55-p-len")
		}
	}
	for _, v := range []map[int]uint8{nil, {}, {111: 0, 77: 127}} {
		// fmt.Printf(">>>> running mammoth map v56: %v\n", v)
		var v56v1, v56v2 map[int]uint8
		var bs56 []byte
		v56v1 = v
		bs56 = testMarshalErr(v56v1, h, t, "enc-map-v56")
		if v != nil {
			if v == nil {
				v56v2 = nil
			} else {
				v56v2 = make(map[int]uint8, len(v))
			} // reset map
			testUnmarshalErr(v56v2, bs56, h, t, "dec-map-v56")
			testDeepEqualErr(v56v1, v56v2, t, "equal-map-v56")
			if v == nil {
				v56v2 = nil
			} else {
				v56v2 = make(map[int]uint8, len(v))
			} // reset map
			testUnmarshalErr(rv4i(v56v2), bs56, h, t, "dec-map-v56-noaddr") // decode into non-addressable map value
			testDeepEqualErr(v56v1, v56v2, t, "equal-map-v56-noaddr")
		}
		if v == nil {
			v56v2 = nil
		} else {
			v56v2 = make(map[int]uint8, len(v))
		} // reset map
		testUnmarshalErr(&v56v2, bs56, h, t, "dec-map-v56-p-len")
		testDeepEqualErr(v56v1, v56v2, t, "equal-map-v56-p-len")
		bs56 = testMarshalErr(&v56v1, h, t, "enc-map-v56-p")
		v56v2 = nil
		testUnmarshalErr(&v56v2, bs56, h, t, "dec-map-v56-p-nil")
		testDeepEqualErr(v56v1, v56v2, t, "equal-map-v56-p-nil")
		// ...
		if v == nil {
			v56v2 = nil
		} else {
			v56v2 = make(map[int]uint8, len(v))
		} // reset map
		var v56v3, v56v4 typMapMapIntUint8
		v56v3 = typMapMapIntUint8(v56v1)
		v56v4 = typMapMapIntUint8(v56v2)
		if v != nil {
			bs56 = testMarshalErr(v56v3, h, t, "enc-map-v56-custom")
			testUnmarshalErr(v56v4, bs56, h, t, "dec-map-v56-p-len")
			testDeepEqualErr(v56v3, v56v4, t, "equal-map-v56-p-len")
		}
	}
	for _, v := range []map[int]uint64{nil, {}, {111: 0, 77: 127}} {
		// fmt.Printf(">>>> running mammoth map v57: %v\n", v)
		var v57v1, v57v2 map[int]uint64
		var bs57 []byte
		v57v1 = v
		bs57 = testMarshalErr(v57v1, h, t, "enc-map-v57")
		if v != nil {
			if v == nil {
				v57v2 = nil
			} else {
				v57v2 = make(map[int]uint64, len(v))
			} // reset map
			testUnmarshalErr(v57v2, bs57, h, t, "dec-map-v57")
			testDeepEqualErr(v57v1, v57v2, t, "equal-map-v57")
			if v == nil {
				v57v2 = nil
			} else {
				v57v2 = make(map[int]uint64, len(v))
			} // reset map
			testUnmarshalErr(rv4i(v57v2), bs57, h, t, "dec-map-v57-noaddr") // decode into non-addressable map value
			testDeepEqualErr(v57v1, v57v2, t, "equal-map-v57-noaddr")
		}
		if v == nil {
			v57v2 = nil
		} else {
			v57v2 = make(map[int]uint64, len(v))
		} // reset map
		testUnmarshalErr(&v57v2, bs57, h, t, "dec-map-v57-p-len")
		testDeepEqualErr(v57v1, v57v2, t, "equal-map-v57-p-len")
		bs57 = testMarshalErr(&v57v1, h, t, "enc-map-v57-p")
		v57v2 = nil
		testUnmarshalErr(&v57v2, bs57, h, t, "dec-map-v57-p-nil")
		testDeepEqualErr(v57v1, v57v2, t, "equal-map-v57-p-nil")
		// ...
		if v == nil {
			v57v2 = nil
		} else {
			v57v2 = make(map[int]uint64, len(v))
		} // reset map
		var v57v3, v57v4 typMapMapIntUint64
		v57v3 = typMapMapIntUint64(v57v1)
		v57v4 = typMapMapIntUint64(v57v2)
		if v != nil {
			bs57 = testMarshalErr(v57v3, h, t, "enc-map-v57-custom")
			testUnmarshalErr(v57v4, bs57, h, t, "dec-map-v57-p-len")
			testDeepEqualErr(v57v3, v57v4, t, "equal-map-v57-p-len")
		}
	}
	for _, v := range []map[int]int{nil, {}, {111: 0, 77: 127}} {
		// fmt.Printf(">>>> running mammoth map v58: %v\n", v)
		var v58v1, v58v2 map[int]int
		var bs58 []byte
		v58v1 = v
		bs58 = testMarshalErr(v58v1, h, t, "enc-map-v58")
		if v != nil {
			if v == nil {
				v58v2 = nil
			} else {
				v58v2 = make(map[int]int, len(v))
			} // reset map
			testUnmarshalErr(v58v2, bs58, h, t, "dec-map-v58")
			testDeepEqualErr(v58v1, v58v2, t, "equal-map-v58")
			if v == nil {
				v58v2 = nil
			} else {
				v58v2 = make(map[int]int, len(v))
			} // reset map
			testUnmarshalErr(rv4i(v58v2), bs58, h, t, "dec-map-v58-noaddr") // decode into non-addressable map value
			testDeepEqualErr(v58v1, v58v2, t, "equal-map-v58-noaddr")
		}
		if v == nil {
			v58v2 = nil
		} else {
			v58v2 = make(map[int]int, len(v))
		} // reset map
		testUnmarshalErr(&v58v2, bs58, h, t, "dec-map-v58-p-len")
		testDeepEqualErr(v58v1, v58v2, t, "equal-map-v58-p-len")
		bs58 = testMarshalErr(&v58v1, h, t, "enc-map-v58-p")
		v58v2 = nil
		testUnmarshalErr(&v58v2, bs58, h, t, "dec-map-v58-p-nil")
		testDeepEqualErr(v58v1, v58v2, t, "equal-map-v58-p-nil")
		// ...
		if v == nil {
			v58v2 = nil
		} else {
			v58v2 = make(map[int]int, len(v))
		} // reset map
		var v58v3, v58v4 typMapMapIntInt
		v58v3 = typMapMapIntInt(v58v1)
		v58v4 = typMapMapIntInt(v58v2)
		if v != nil {
			bs58 = testMarshalErr(v58v3, h, t, "enc-map-v58-custom")
			testUnmarshalErr(v58v4, bs58, h, t, "dec-map-v58-p-len")
			testDeepEqualErr(v58v3, v58v4, t, "equal-map-v58-p-len")
		}
	}
	for _, v := range []map[int]int64{nil, {}, {111: 0, 77: 127}} {
		// fmt.Printf(">>>> running mammoth map v59: %v\n", v)
		var v59v1, v59v2 map[int]int64
		var bs59 []byte
		v59v1 = v
		bs59 = testMarshalErr(v59v1, h, t, "enc-map-v59")
		if v != nil {
			if v == nil {
				v59v2 = nil
			} else {
				v59v2 = make(map[int]int64, len(v))
			} // reset map
			testUnmarshalErr(v59v2, bs59, h, t, "dec-map-v59")
			testDeepEqualErr(v59v1, v59v2, t, "equal-map-v59")
			if v == nil {
				v59v2 = nil
			} else {
				v59v2 = make(map[int]int64, len(v))
			} // reset map
			testUnmarshalErr(rv4i(v59v2), bs59, h, t, "dec-map-v59-noaddr") // decode into non-addressable map value
			testDeepEqualErr(v59v1, v59v2, t, "equal-map-v59-noaddr")
		}
		if v == nil {
			v59v2 = nil
		} else {
			v59v2 = make(map[int]int64, len(v))
		} // reset map
		testUnmarshalErr(&v59v2, bs59, h, t, "dec-map-v59-p-len")
		testDeepEqualErr(v59v1, v59v2, t, "equal-map-v59-p-len")
		bs59 = testMarshalErr(&v59v1, h, t, "enc-map-v59-p")
		v59v2 = nil
		testUnmarshalErr(&v59v2, bs59, h, t, "dec-map-v59-p-nil")
		testDeepEqualErr(v59v1, v59v2, t, "equal-map-v59-p-nil")
		// ...
		if v == nil {
			v59v2 = nil
		} else {
			v59v2 = make(map[int]int64, len(v))
		} // reset map
		var v59v3, v59v4 typMapMapIntInt64
		v59v3 = typMapMapIntInt64(v59v1)
		v59v4 = typMapMapIntInt64(v59v2)
		if v != nil {
			bs59 = testMarshalErr(v59v3, h, t, "enc-map-v59-custom")
			testUnmarshalErr(v59v4, bs59, h, t, "dec-map-v59-p-len")
			testDeepEqualErr(v59v3, v59v4, t, "equal-map-v59-p-len")
		}
	}
	for _, v := range []map[int]float64{nil, {}, {111: 0, 77: 11.1}} {
		// fmt.Printf(">>>> running mammoth map v60: %v\n", v)
		var v60v1, v60v2 map[int]float64
		var bs60 []byte
		v60v1 = v
		bs60 = testMarshalErr(v60v1, h, t, "enc-map-v60")
		if v != nil {
			if v == nil {
				v60v2 = nil
			} else {
				v60v2 = make(map[int]float64, len(v))
			} // reset map
			testUnmarshalErr(v60v2, bs60, h, t, "dec-map-v60")
			testDeepEqualErr(v60v1, v60v2, t, "equal-map-v60")
			if v == nil {
				v60v2 = nil
			} else {
				v60v2 = make(map[int]float64, len(v))
			} // reset map
			testUnmarshalErr(rv4i(v60v2), bs60, h, t, "dec-map-v60-noaddr") // decode into non-addressable map value
			testDeepEqualErr(v60v1, v60v2, t, "equal-map-v60-noaddr")
		}
		if v == nil {
			v60v2 = nil
		} else {
			v60v2 = make(map[int]float64, len(v))
		} // reset map
		testUnmarshalErr(&v60v2, bs60, h, t, "dec-map-v60-p-len")
		testDeepEqualErr(v60v1, v60v2, t, "equal-map-v60-p-len")
		bs60 = testMarshalErr(&v60v1, h, t, "enc-map-v60-p")
		v60v2 = nil
		testUnmarshalErr(&v60v2, bs60, h, t, "dec-map-v60-p-nil")
		testDeepEqualErr(v60v1, v60v2, t, "equal-map-v60-p-nil")
		// ...
		if v == nil {
			v60v2 = nil
		} else {
			v60v2 = make(map[int]float64, len(v))
		} // reset map
		var v60v3, v60v4 typMapMapIntFloat64
		v60v3 = typMapMapIntFloat64(v60v1)
		v60v4 = typMapMapIntFloat64(v60v2)
		if v != nil {
			bs60 = testMarshalErr(v60v3, h, t, "enc-map-v60-custom")
			testUnmarshalErr(v60v4, bs60, h, t, "dec-map-v60-p-len")
			testDeepEqualErr(v60v3, v60v4, t, "equal-map-v60-p-len")
		}
	}
	for _, v := range []map[int]bool{nil, {}, {127: false, 111: true}} {
		// fmt.Printf(">>>> running mammoth map v61: %v\n", v)
		var v61v1, v61v2 map[int]bool
		var bs61 []byte
		v61v1 = v
		bs61 = testMarshalErr(v61v1, h, t, "enc-map-v61")
		if v != nil {
			if v == nil {
				v61v2 = nil
			} else {
				v61v2 = make(map[int]bool, len(v))
			} // reset map
			testUnmarshalErr(v61v2, bs61, h, t, "dec-map-v61")
			testDeepEqualErr(v61v1, v61v2, t, "equal-map-v61")
			if v == nil {
				v61v2 = nil
			} else {
				v61v2 = make(map[int]bool, len(v))
			} // reset map
			testUnmarshalErr(rv4i(v61v2), bs61, h, t, "dec-map-v61-noaddr") // decode into non-addressable map value
			testDeepEqualErr(v61v1, v61v2, t, "equal-map-v61-noaddr")
		}
		if v == nil {
			v61v2 = nil
		} else {
			v61v2 = make(map[int]bool, len(v))
		} // reset map
		testUnmarshalErr(&v61v2, bs61, h, t, "dec-map-v61-p-len")
		testDeepEqualErr(v61v1, v61v2, t, "equal-map-v61-p-len")
		bs61 = testMarshalErr(&v61v1, h, t, "enc-map-v61-p")
		v61v2 = nil
		testUnmarshalErr(&v61v2, bs61, h, t, "dec-map-v61-p-nil")
		testDeepEqualErr(v61v1, v61v2, t, "equal-map-v61-p-nil")
		// ...
		if v == nil {
			v61v2 = nil
		} else {
			v61v2 = make(map[int]bool, len(v))
		} // reset map
		var v61v3, v61v4 typMapMapIntBool
		v61v3 = typMapMapIntBool(v61v1)
		v61v4 = typMapMapIntBool(v61v2)
		if v != nil {
			bs61 = testMarshalErr(v61v3, h, t, "enc-map-v61-custom")
			testUnmarshalErr(v61v4, bs61, h, t, "dec-map-v61-p-len")
			testDeepEqualErr(v61v3, v61v4, t, "equal-map-v61-p-len")
		}
	}
	for _, v := range []map[int64]interface{}{nil, {}, {77: nil, 127: "string-is-an-interface-2"}} {
		// fmt.Printf(">>>> running mammoth map v62: %v\n", v)
		var v62v1, v62v2 map[int64]interface{}
		var bs62 []byte
		v62v1 = v
		bs62 = testMarshalErr(v62v1, h, t, "enc-map-v62")
		if v != nil {
			if v == nil {
				v62v2 = nil
			} else {
				v62v2 = make(map[int64]interface{}, len(v))
			} // reset map
			testUnmarshalErr(v62v2, bs62, h, t, "dec-map-v62")
			testDeepEqualErr(v62v1, v62v2, t, "equal-map-v62")
			if v == nil {
				v62v2 = nil
			} else {
				v62v2 = make(map[int64]interface{}, len(v))
			} // reset map
			testUnmarshalErr(rv4i(v62v2), bs62, h, t, "dec-map-v62-noaddr") // decode into non-addressable map value
			testDeepEqualErr(v62v1, v62v2, t, "equal-map-v62-noaddr")
		}
		if v == nil {
			v62v2 = nil
		} else {
			v62v2 = make(map[int64]interface{}, len(v))
		} // reset map
		testUnmarshalErr(&v62v2, bs62, h, t, "dec-map-v62-p-len")
		testDeepEqualErr(v62v1, v62v2, t, "equal-map-v62-p-len")
		bs62 = testMarshalErr(&v62v1, h, t, "enc-map-v62-p")
		v62v2 = nil
		testUnmarshalErr(&v62v2, bs62, h, t, "dec-map-v62-p-nil")
		testDeepEqualErr(v62v1, v62v2, t, "equal-map-v62-p-nil")
		// ...
		if v == nil {
			v62v2 = nil
		} else {
			v62v2 = make(map[int64]interface{}, len(v))
		} // reset map
		var v62v3, v62v4 typMapMapInt64Intf
		v62v3 = typMapMapInt64Intf(v62v1)
		v62v4 = typMapMapInt64Intf(v62v2)
		if v != nil {
			bs62 = testMarshalErr(v62v3, h, t, "enc-map-v62-custom")
			testUnmarshalErr(v62v4, bs62, h, t, "dec-map-v62-p-len")
			testDeepEqualErr(v62v3, v62v4, t, "equal-map-v62-p-len")
		}
	}
	for _, v := range []map[int64]string{nil, {}, {111: "", 77: "some-string-2"}} {
		// fmt.Printf(">>>> running mammoth map v63: %v\n", v)
		var v63v1, v63v2 map[int64]string
		var bs63 []byte
		v63v1 = v
		bs63 = testMarshalErr(v63v1, h, t, "enc-map-v63")
		if v != nil {
			if v == nil {
				v63v2 = nil
			} else {
				v63v2 = make(map[int64]string, len(v))
			} // reset map
			testUnmarshalErr(v63v2, bs63, h, t, "dec-map-v63")
			testDeepEqualErr(v63v1, v63v2, t, "equal-map-v63")
			if v == nil {
				v63v2 = nil
			} else {
				v63v2 = make(map[int64]string, len(v))
			} // reset map
			testUnmarshalErr(rv4i(v63v2), bs63, h, t, "dec-map-v63-noaddr") // decode into non-addressable map value
			testDeepEqualErr(v63v1, v63v2, t, "equal-map-v63-noaddr")
		}
		if v == nil {
			v63v2 = nil
		} else {
			v63v2 = make(map[int64]string, len(v))
		} // reset map
		testUnmarshalErr(&v63v2, bs63, h, t, "dec-map-v63-p-len")
		testDeepEqualErr(v63v1, v63v2, t, "equal-map-v63-p-len")
		bs63 = testMarshalErr(&v63v1, h, t, "enc-map-v63-p")
		v63v2 = nil
		testUnmarshalErr(&v63v2, bs63, h, t, "dec-map-v63-p-nil")
		testDeepEqualErr(v63v1, v63v2, t, "equal-map-v63-p-nil")
		// ...
		if v == nil {
			v63v2 = nil
		} else {
			v63v2 = make(map[int64]string, len(v))
		} // reset map
		var v63v3, v63v4 typMapMapInt64String
		v63v3 = typMapMapInt64String(v63v1)
		v63v4 = typMapMapInt64String(v63v2)
		if v != nil {
			bs63 = testMarshalErr(v63v3, h, t, "enc-map-v63-custom")
			testUnmarshalErr(v63v4, bs63, h, t, "dec-map-v63-p-len")
			testDeepEqualErr(v63v3, v63v4, t, "equal-map-v63-p-len")
		}
	}
	for _, v := range []map[int64][]byte{nil, {}, {127: nil, 111: []byte("some-string-2")}} {
		// fmt.Printf(">>>> running mammoth map v64: %v\n", v)
		var v64v1, v64v2 map[int64][]byte
		var bs64 []byte
		v64v1 = v
		bs64 = testMarshalErr(v64v1, h, t, "enc-map-v64")
		if v != nil {
			if v == nil {
				v64v2 = nil
			} else {
				v64v2 = make(map[int64][]byte, len(v))
			} // reset map
			testUnmarshalErr(v64v2, bs64, h, t, "dec-map-v64")
			testDeepEqualErr(v64v1, v64v2, t, "equal-map-v64")
			if v == nil {
				v64v2 = nil
			} else {
				v64v2 = make(map[int64][]byte, len(v))
			} // reset map
			testUnmarshalErr(rv4i(v64v2), bs64, h, t, "dec-map-v64-noaddr") // decode into non-addressable map value
			testDeepEqualErr(v64v1, v64v2, t, "equal-map-v64-noaddr")
		}
		if v == nil {
			v64v2 = nil
		} else {
			v64v2 = make(map[int64][]byte, len(v))
		} // reset map
		testUnmarshalErr(&v64v2, bs64, h, t, "dec-map-v64-p-len")
		testDeepEqualErr(v64v1, v64v2, t, "equal-map-v64-p-len")
		bs64 = testMarshalErr(&v64v1, h, t, "enc-map-v64-p")
		v64v2 = nil
		testUnmarshalErr(&v64v2, bs64, h, t, "dec-map-v64-p-nil")
		testDeepEqualErr(v64v1, v64v2, t, "equal-map-v64-p-nil")
		// ...
		if v == nil {
			v64v2 = nil
		} else {
			v64v2 = make(map[int64][]byte, len(v))
		} // reset map
		var v64v3, v64v4 typMapMapInt64Bytes
		v64v3 = typMapMapInt64Bytes(v64v1)
		v64v4 = typMapMapInt64Bytes(v64v2)
		if v != nil {
			bs64 = testMarshalErr(v64v3, h, t, "enc-map-v64-custom")
			testUnmarshalErr(v64v4, bs64, h, t, "dec-map-v64-p-len")
			testDeepEqualErr(v64v3, v64v4, t, "equal-map-v64-p-len")
		}
	}
	for _, v := range []map[int64]uint8{nil, {}, {77: 0, 127: 111}} {
		// fmt.Printf(">>>> running mammoth map v65: %v\n", v)
		var v65v1, v65v2 map[int64]uint8
		var bs65 []byte
		v65v1 = v
		bs65 = testMarshalErr(v65v1, h, t, "enc-map-v65")
		if v != nil {
			if v == nil {
				v65v2 = nil
			} else {
				v65v2 = make(map[int64]uint8, len(v))
			} // reset map
			testUnmarshalErr(v65v2, bs65, h, t, "dec-map-v65")
			testDeepEqualErr(v65v1, v65v2, t, "equal-map-v65")
			if v == nil {
				v65v2 = nil
			} else {
				v65v2 = make(map[int64]uint8, len(v))
			} // reset map
			testUnmarshalErr(rv4i(v65v2), bs65, h, t, "dec-map-v65-noaddr") // decode into non-addressable map value
			testDeepEqualErr(v65v1, v65v2, t, "equal-map-v65-noaddr")
		}
		if v == nil {
			v65v2 = nil
		} else {
			v65v2 = make(map[int64]uint8, len(v))
		} // reset map
		testUnmarshalErr(&v65v2, bs65, h, t, "dec-map-v65-p-len")
		testDeepEqualErr(v65v1, v65v2, t, "equal-map-v65-p-len")
		bs65 = testMarshalErr(&v65v1, h, t, "enc-map-v65-p")
		v65v2 = nil
		testUnmarshalErr(&v65v2, bs65, h, t, "dec-map-v65-p-nil")
		testDeepEqualErr(v65v1, v65v2, t, "equal-map-v65-p-nil")
		// ...
		if v == nil {
			v65v2 = nil
		} else {
			v65v2 = make(map[int64]uint8, len(v))
		} // reset map
		var v65v3, v65v4 typMapMapInt64Uint8
		v65v3 = typMapMapInt64Uint8(v65v1)
		v65v4 = typMapMapInt64Uint8(v65v2)
		if v != nil {
			bs65 = testMarshalErr(v65v3, h, t, "enc-map-v65-custom")
			testUnmarshalErr(v65v4, bs65, h, t, "dec-map-v65-p-len")
			testDeepEqualErr(v65v3, v65v4, t, "equal-map-v65-p-len")
		}
	}
	for _, v := range []map[int64]uint64{nil, {}, {77: 0, 127: 111}} {
		// fmt.Printf(">>>> running mammoth map v66: %v\n", v)
		var v66v1, v66v2 map[int64]uint64
		var bs66 []byte
		v66v1 = v
		bs66 = testMarshalErr(v66v1, h, t, "enc-map-v66")
		if v != nil {
			if v == nil {
				v66v2 = nil
			} else {
				v66v2 = make(map[int64]uint64, len(v))
			} // reset map
			testUnmarshalErr(v66v2, bs66, h, t, "dec-map-v66")
			testDeepEqualErr(v66v1, v66v2, t, "equal-map-v66")
			if v == nil {
				v66v2 = nil
			} else {
				v66v2 = make(map[int64]uint64, len(v))
			} // reset map
			testUnmarshalErr(rv4i(v66v2), bs66, h, t, "dec-map-v66-noaddr") // decode into non-addressable map value
			testDeepEqualErr(v66v1, v66v2, t, "equal-map-v66-noaddr")
		}
		if v == nil {
			v66v2 = nil
		} else {
			v66v2 = make(map[int64]uint64, len(v))
		} // reset map
		testUnmarshalErr(&v66v2, bs66, h, t, "dec-map-v66-p-len")
		testDeepEqualErr(v66v1, v66v2, t, "equal-map-v66-p-len")
		bs66 = testMarshalErr(&v66v1, h, t, "enc-map-v66-p")
		v66v2 = nil
		testUnmarshalErr(&v66v2, bs66, h, t, "dec-map-v66-p-nil")
		testDeepEqualErr(v66v1, v66v2, t, "equal-map-v66-p-nil")
		// ...
		if v == nil {
			v66v2 = nil
		} else {
			v66v2 = make(map[int64]uint64, len(v))
		} // reset map
		var v66v3, v66v4 typMapMapInt64Uint64
		v66v3 = typMapMapInt64Uint64(v66v1)
		v66v4 = typMapMapInt64Uint64(v66v2)
		if v != nil {
			bs66 = testMarshalErr(v66v3, h, t, "enc-map-v66-custom")
			testUnmarshalErr(v66v4, bs66, h, t, "dec-map-v66-p-len")
			testDeepEqualErr(v66v3, v66v4, t, "equal-map-v66-p-len")
		}
	}
	for _, v := range []map[int64]int{nil, {}, {77: 0, 127: 111}} {
		// fmt.Printf(">>>> running mammoth map v67: %v\n", v)
		var v67v1, v67v2 map[int64]int
		var bs67 []byte
		v67v1 = v
		bs67 = testMarshalErr(v67v1, h, t, "enc-map-v67")
		if v != nil {
			if v == nil {
				v67v2 = nil
			} else {
				v67v2 = make(map[int64]int, len(v))
			} // reset map
			testUnmarshalErr(v67v2, bs67, h, t, "dec-map-v67")
			testDeepEqualErr(v67v1, v67v2, t, "equal-map-v67")
			if v == nil {
				v67v2 = nil
			} else {
				v67v2 = make(map[int64]int, len(v))
			} // reset map
			testUnmarshalErr(rv4i(v67v2), bs67, h, t, "dec-map-v67-noaddr") // decode into non-addressable map value
			testDeepEqualErr(v67v1, v67v2, t, "equal-map-v67-noaddr")
		}
		if v == nil {
			v67v2 = nil
		} else {
			v67v2 = make(map[int64]int, len(v))
		} // reset map
		testUnmarshalErr(&v67v2, bs67, h, t, "dec-map-v67-p-len")
		testDeepEqualErr(v67v1, v67v2, t, "equal-map-v67-p-len")
		bs67 = testMarshalErr(&v67v1, h, t, "enc-map-v67-p")
		v67v2 = nil
		testUnmarshalErr(&v67v2, bs67, h, t, "dec-map-v67-p-nil")
		testDeepEqualErr(v67v1, v67v2, t, "equal-map-v67-p-nil")
		// ...
		if v == nil {
			v67v2 = nil
		} else {
			v67v2 = make(map[int64]int, len(v))
		} // reset map
		var v67v3, v67v4 typMapMapInt64Int
		v67v3 = typMapMapInt64Int(v67v1)
		v67v4 = typMapMapInt64Int(v67v2)
		if v != nil {
			bs67 = testMarshalErr(v67v3, h, t, "enc-map-v67-custom")
			testUnmarshalErr(v67v4, bs67, h, t, "dec-map-v67-p-len")
			testDeepEqualErr(v67v3, v67v4, t, "equal-map-v67-p-len")
		}
	}
	for _, v := range []map[int64]int64{nil, {}, {77: 0, 127: 111}} {
		// fmt.Printf(">>>> running mammoth map v68: %v\n", v)
		var v68v1, v68v2 map[int64]int64
		var bs68 []byte
		v68v1 = v
		bs68 = testMarshalErr(v68v1, h, t, "enc-map-v68")
		if v != nil {
			if v == nil {
				v68v2 = nil
			} else {
				v68v2 = make(map[int64]int64, len(v))
			} // reset map
			testUnmarshalErr(v68v2, bs68, h, t, "dec-map-v68")
			testDeepEqualErr(v68v1, v68v2, t, "equal-map-v68")
			if v == nil {
				v68v2 = nil
			} else {
				v68v2 = make(map[int64]int64, len(v))
			} // reset map
			testUnmarshalErr(rv4i(v68v2), bs68, h, t, "dec-map-v68-noaddr") // decode into non-addressable map value
			testDeepEqualErr(v68v1, v68v2, t, "equal-map-v68-noaddr")
		}
		if v == nil {
			v68v2 = nil
		} else {
			v68v2 = make(map[int64]int64, len(v))
		} // reset map
		testUnmarshalErr(&v68v2, bs68, h, t, "dec-map-v68-p-len")
		testDeepEqualErr(v68v1, v68v2, t, "equal-map-v68-p-len")
		bs68 = testMarshalErr(&v68v1, h, t, "enc-map-v68-p")
		v68v2 = nil
		testUnmarshalErr(&v68v2, bs68, h, t, "dec-map-v68-p-nil")
		testDeepEqualErr(v68v1, v68v2, t, "equal-map-v68-p-nil")
		// ...
		if v == nil {
			v68v2 = nil
		} else {
			v68v2 = make(map[int64]int64, len(v))
		} // reset map
		var v68v3, v68v4 typMapMapInt64Int64
		v68v3 = typMapMapInt64Int64(v68v1)
		v68v4 = typMapMapInt64Int64(v68v2)
		if v != nil {
			bs68 = testMarshalErr(v68v3, h, t, "enc-map-v68-custom")
			testUnmarshalErr(v68v4, bs68, h, t, "dec-map-v68-p-len")
			testDeepEqualErr(v68v3, v68v4, t, "equal-map-v68-p-len")
		}
	}
	for _, v := range []map[int64]float64{nil, {}, {77: 0, 127: 22.2}} {
		// fmt.Printf(">>>> running mammoth map v69: %v\n", v)
		var v69v1, v69v2 map[int64]float64
		var bs69 []byte
		v69v1 = v
		bs69 = testMarshalErr(v69v1, h, t, "enc-map-v69")
		if v != nil {
			if v == nil {
				v69v2 = nil
			} else {
				v69v2 = make(map[int64]float64, len(v))
			} // reset map
			testUnmarshalErr(v69v2, bs69, h, t, "dec-map-v69")
			testDeepEqualErr(v69v1, v69v2, t, "equal-map-v69")
			if v == nil {
				v69v2 = nil
			} else {
				v69v2 = make(map[int64]float64, len(v))
			} // reset map
			testUnmarshalErr(rv4i(v69v2), bs69, h, t, "dec-map-v69-noaddr") // decode into non-addressable map value
			testDeepEqualErr(v69v1, v69v2, t, "equal-map-v69-noaddr")
		}
		if v == nil {
			v69v2 = nil
		} else {
			v69v2 = make(map[int64]float64, len(v))
		} // reset map
		testUnmarshalErr(&v69v2, bs69, h, t, "dec-map-v69-p-len")
		testDeepEqualErr(v69v1, v69v2, t, "equal-map-v69-p-len")
		bs69 = testMarshalErr(&v69v1, h, t, "enc-map-v69-p")
		v69v2 = nil
		testUnmarshalErr(&v69v2, bs69, h, t, "dec-map-v69-p-nil")
		testDeepEqualErr(v69v1, v69v2, t, "equal-map-v69-p-nil")
		// ...
		if v == nil {
			v69v2 = nil
		} else {
			v69v2 = make(map[int64]float64, len(v))
		} // reset map
		var v69v3, v69v4 typMapMapInt64Float64
		v69v3 = typMapMapInt64Float64(v69v1)
		v69v4 = typMapMapInt64Float64(v69v2)
		if v != nil {
			bs69 = testMarshalErr(v69v3, h, t, "enc-map-v69-custom")
			testUnmarshalErr(v69v4, bs69, h, t, "dec-map-v69-p-len")
			testDeepEqualErr(v69v3, v69v4, t, "equal-map-v69-p-len")
		}
	}
	for _, v := range []map[int64]bool{nil, {}, {111: false, 77: false}} {
		// fmt.Printf(">>>> running mammoth map v70: %v\n", v)
		var v70v1, v70v2 map[int64]bool
		var bs70 []byte
		v70v1 = v
		bs70 = testMarshalErr(v70v1, h, t, "enc-map-v70")
		if v != nil {
			if v == nil {
				v70v2 = nil
			} else {
				v70v2 = make(map[int64]bool, len(v))
			} // reset map
			testUnmarshalErr(v70v2, bs70, h, t, "dec-map-v70")
			testDeepEqualErr(v70v1, v70v2, t, "equal-map-v70")
			if v == nil {
				v70v2 = nil
			} else {
				v70v2 = make(map[int64]bool, len(v))
			} // reset map
			testUnmarshalErr(rv4i(v70v2), bs70, h, t, "dec-map-v70-noaddr") // decode into non-addressable map value
			testDeepEqualErr(v70v1, v70v2, t, "equal-map-v70-noaddr")
		}
		if v == nil {
			v70v2 = nil
		} else {
			v70v2 = make(map[int64]bool, len(v))
		} // reset map
		testUnmarshalErr(&v70v2, bs70, h, t, "dec-map-v70-p-len")
		testDeepEqualErr(v70v1, v70v2, t, "equal-map-v70-p-len")
		bs70 = testMarshalErr(&v70v1, h, t, "enc-map-v70-p")
		v70v2 = nil
		testUnmarshalErr(&v70v2, bs70, h, t, "dec-map-v70-p-nil")
		testDeepEqualErr(v70v1, v70v2, t, "equal-map-v70-p-nil")
		// ...
		if v == nil {
			v70v2 = nil
		} else {
			v70v2 = make(map[int64]bool, len(v))
		} // reset map
		var v70v3, v70v4 typMapMapInt64Bool
		v70v3 = typMapMapInt64Bool(v70v1)
		v70v4 = typMapMapInt64Bool(v70v2)
		if v != nil {
			bs70 = testMarshalErr(v70v3, h, t, "enc-map-v70-custom")
			testUnmarshalErr(v70v4, bs70, h, t, "dec-map-v70-p-len")
			testDeepEqualErr(v70v3, v70v4, t, "equal-map-v70-p-len")
		}
	}

}

func doTestMammothMapsAndSlices(t *testing.T, h Handle) {
	doTestMammothSlices(t, h)
	doTestMammothMaps(t, h)
}
