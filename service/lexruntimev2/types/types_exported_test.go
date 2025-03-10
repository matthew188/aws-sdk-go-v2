// Code generated by smithy-go-codegen DO NOT EDIT.

package types_test

import (
	"fmt"
	"github.com/matthew188/aws-sdk-go-v2/service/lexruntimev2/types"
)

func ExampleStartConversationRequestEventStream_outputUsage() {
	var union types.StartConversationRequestEventStream
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.StartConversationRequestEventStreamMemberAudioInputEvent:
		_ = v.Value // Value is types.AudioInputEvent

	case *types.StartConversationRequestEventStreamMemberConfigurationEvent:
		_ = v.Value // Value is types.ConfigurationEvent

	case *types.StartConversationRequestEventStreamMemberDisconnectionEvent:
		_ = v.Value // Value is types.DisconnectionEvent

	case *types.StartConversationRequestEventStreamMemberDTMFInputEvent:
		_ = v.Value // Value is types.DTMFInputEvent

	case *types.StartConversationRequestEventStreamMemberPlaybackCompletionEvent:
		_ = v.Value // Value is types.PlaybackCompletionEvent

	case *types.StartConversationRequestEventStreamMemberTextInputEvent:
		_ = v.Value // Value is types.TextInputEvent

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.AudioInputEvent
var _ *types.TextInputEvent
var _ *types.ConfigurationEvent
var _ *types.DisconnectionEvent
var _ *types.DTMFInputEvent
var _ *types.PlaybackCompletionEvent

func ExampleStartConversationResponseEventStream_outputUsage() {
	var union types.StartConversationResponseEventStream
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.StartConversationResponseEventStreamMemberAudioResponseEvent:
		_ = v.Value // Value is types.AudioResponseEvent

	case *types.StartConversationResponseEventStreamMemberHeartbeatEvent:
		_ = v.Value // Value is types.HeartbeatEvent

	case *types.StartConversationResponseEventStreamMemberIntentResultEvent:
		_ = v.Value // Value is types.IntentResultEvent

	case *types.StartConversationResponseEventStreamMemberPlaybackInterruptionEvent:
		_ = v.Value // Value is types.PlaybackInterruptionEvent

	case *types.StartConversationResponseEventStreamMemberTextResponseEvent:
		_ = v.Value // Value is types.TextResponseEvent

	case *types.StartConversationResponseEventStreamMemberTranscriptEvent:
		_ = v.Value // Value is types.TranscriptEvent

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.IntentResultEvent
var _ *types.HeartbeatEvent
var _ *types.AudioResponseEvent
var _ *types.PlaybackInterruptionEvent
var _ *types.TranscriptEvent
var _ *types.TextResponseEvent
