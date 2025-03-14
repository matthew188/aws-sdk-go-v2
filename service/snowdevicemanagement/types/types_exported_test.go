// Code generated by smithy-go-codegen DO NOT EDIT.

package types_test

import (
	"fmt"
	"github.com/matthew188/aws-sdk-go-v2/service/snowdevicemanagement/types"
)

func ExampleCommand_outputUsage() {
	var union types.Command
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.CommandMemberReboot:
		_ = v.Value // Value is types.Reboot

	case *types.CommandMemberUnlock:
		_ = v.Value // Value is types.Unlock

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.Reboot
var _ *types.Unlock
