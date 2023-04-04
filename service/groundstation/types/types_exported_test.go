// Code generated by smithy-go-codegen DO NOT EDIT.

package types_test

import (
	"fmt"
	"github.com/matthew188/aws-sdk-go-v2/service/groundstation/types"
)

func ExampleConfigDetails_outputUsage() {
	var union types.ConfigDetails
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ConfigDetailsMemberAntennaDemodDecodeDetails:
		_ = v.Value // Value is types.AntennaDemodDecodeDetails

	case *types.ConfigDetailsMemberEndpointDetails:
		_ = v.Value // Value is types.EndpointDetails

	case *types.ConfigDetailsMemberS3RecordingDetails:
		_ = v.Value // Value is types.S3RecordingDetails

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.S3RecordingDetails
var _ *types.AntennaDemodDecodeDetails
var _ *types.EndpointDetails

func ExampleConfigTypeData_outputUsage() {
	var union types.ConfigTypeData
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ConfigTypeDataMemberAntennaDownlinkConfig:
		_ = v.Value // Value is types.AntennaDownlinkConfig

	case *types.ConfigTypeDataMemberAntennaDownlinkDemodDecodeConfig:
		_ = v.Value // Value is types.AntennaDownlinkDemodDecodeConfig

	case *types.ConfigTypeDataMemberAntennaUplinkConfig:
		_ = v.Value // Value is types.AntennaUplinkConfig

	case *types.ConfigTypeDataMemberDataflowEndpointConfig:
		_ = v.Value // Value is types.DataflowEndpointConfig

	case *types.ConfigTypeDataMemberS3RecordingConfig:
		_ = v.Value // Value is types.S3RecordingConfig

	case *types.ConfigTypeDataMemberTrackingConfig:
		_ = v.Value // Value is types.TrackingConfig

	case *types.ConfigTypeDataMemberUplinkEchoConfig:
		_ = v.Value // Value is types.UplinkEchoConfig

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.TrackingConfig
var _ *types.UplinkEchoConfig
var _ *types.AntennaDownlinkDemodDecodeConfig
var _ *types.S3RecordingConfig
var _ *types.AntennaUplinkConfig
var _ *types.AntennaDownlinkConfig
var _ *types.DataflowEndpointConfig

func ExampleEphemerisData_outputUsage() {
	var union types.EphemerisData
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.EphemerisDataMemberOem:
		_ = v.Value // Value is types.OEMEphemeris

	case *types.EphemerisDataMemberTle:
		_ = v.Value // Value is types.TLEEphemeris

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.TLEEphemeris
var _ *types.OEMEphemeris

func ExampleEphemerisTypeDescription_outputUsage() {
	var union types.EphemerisTypeDescription
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.EphemerisTypeDescriptionMemberOem:
		_ = v.Value // Value is types.EphemerisDescription

	case *types.EphemerisTypeDescriptionMemberTle:
		_ = v.Value // Value is types.EphemerisDescription

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.EphemerisDescription

func ExampleKmsKey_outputUsage() {
	var union types.KmsKey
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.KmsKeyMemberKmsAliasArn:
		_ = v.Value // Value is string

	case *types.KmsKeyMemberKmsKeyArn:
		_ = v.Value // Value is string

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *string
var _ *string
