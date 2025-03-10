// Code generated by smithy-go-codegen DO NOT EDIT.

package types_test

import (
	"fmt"
	"github.com/matthew188/aws-sdk-go-v2/service/rdsdata/types"
)

func ExampleArrayValue_outputUsage() {
	var union types.ArrayValue
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ArrayValueMemberArrayValues:
		_ = v.Value // Value is []types.ArrayValue

	case *types.ArrayValueMemberBooleanValues:
		_ = v.Value // Value is []bool

	case *types.ArrayValueMemberDoubleValues:
		_ = v.Value // Value is []float64

	case *types.ArrayValueMemberLongValues:
		_ = v.Value // Value is []int64

	case *types.ArrayValueMemberStringValues:
		_ = v.Value // Value is []string

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ []int64
var _ []float64
var _ []string
var _ []types.ArrayValue
var _ []bool

func ExampleField_outputUsage() {
	var union types.Field
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.FieldMemberArrayValue:
		_ = v.Value // Value is types.ArrayValue

	case *types.FieldMemberBlobValue:
		_ = v.Value // Value is []byte

	case *types.FieldMemberBooleanValue:
		_ = v.Value // Value is bool

	case *types.FieldMemberDoubleValue:
		_ = v.Value // Value is float64

	case *types.FieldMemberIsNull:
		_ = v.Value // Value is bool

	case *types.FieldMemberLongValue:
		_ = v.Value // Value is int64

	case *types.FieldMemberStringValue:
		_ = v.Value // Value is string

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ types.ArrayValue
var _ *string
var _ *bool
var _ *int64
var _ *float64
var _ []byte

func ExampleValue_outputUsage() {
	var union types.Value
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ValueMemberArrayValues:
		_ = v.Value // Value is []types.Value

	case *types.ValueMemberBigIntValue:
		_ = v.Value // Value is int64

	case *types.ValueMemberBitValue:
		_ = v.Value // Value is bool

	case *types.ValueMemberBlobValue:
		_ = v.Value // Value is []byte

	case *types.ValueMemberDoubleValue:
		_ = v.Value // Value is float64

	case *types.ValueMemberIntValue:
		_ = v.Value // Value is int32

	case *types.ValueMemberIsNull:
		_ = v.Value // Value is bool

	case *types.ValueMemberRealValue:
		_ = v.Value // Value is float32

	case *types.ValueMemberStringValue:
		_ = v.Value // Value is string

	case *types.ValueMemberStructValue:
		_ = v.Value // Value is types.StructValue

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.StructValue
var _ *string
var _ *int32
var _ *int64
var _ *bool
var _ *float64
var _ []types.Value
var _ *float32
var _ []byte
