// Code generated by smithy-go-codegen DO NOT EDIT.

package types_test

import (
	"fmt"
	"github.com/matthew188/aws-sdk-go-v2/service/kinesis/types"
)

func ExampleSubscribeToShardEventStream_outputUsage() {
	var union types.SubscribeToShardEventStream
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.SubscribeToShardEventStreamMemberSubscribeToShardEvent:
		_ = v.Value // Value is types.SubscribeToShardEvent

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.SubscribeToShardEvent
