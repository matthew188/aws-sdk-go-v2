// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// A structure that contains the configuration settings for an Amazon Transcribe
// call analytics processor.
type AmazonTranscribeCallAnalyticsProcessorConfiguration struct {

	// The language code in the configuration.
	//
	// This member is required.
	LanguageCode CallAnalyticsLanguageCode

	// By default, all CategoryEvents will be sent to the insights target. If this
	// parameter is specified, only included categories will be sent to the insights
	// target.
	CallAnalyticsStreamCategories []string

	// Labels all personally identifiable information (PII) identified in your
	// transcript. Content identification is performed at the segment level; PII
	// specified in PiiEntityTypes is flagged upon complete transcription of an audio
	// segment. You can’t set ContentIdentificationType and ContentRedactionType in the
	// same request. If you do, your request returns a BadRequestException. For more
	// information, see Redacting or identifying personally identifiable information
	// (https://docs.aws.amazon.com/transcribe/latest/dg/pii-redaction.html) in the
	// Amazon Transcribe Developer Guide.
	ContentIdentificationType ContentType

	// Redacts all personally identifiable information (PII) identified in your
	// transcript. Content redaction is performed at the segment level; PII specified
	// in PiiEntityTypes is redacted upon complete transcription of an audio segment.
	// You can’t set ContentRedactionType and ContentIdentificationType in the same
	// request. If you do, your request returns a BadRequestException. For more
	// information, see Redacting or identifying personally identifiable information
	// (https://docs.aws.amazon.com/transcribe/latest/dg/pii-redaction.html) in the
	// Amazon Transcribe Developer Guide.
	ContentRedactionType ContentType

	// Enables partial result stabilization for your transcription. Partial result
	// stabilization can reduce latency in your output, but may impact accuracy. For
	// more information, see Partial-result stabilization
	// (https://docs.aws.amazon.com/transcribe/latest/dg/streaming.html#streaming-partial-result-stabilization)
	// in the Amazon Transcribe Developer Guide.
	EnablePartialResultsStabilization bool

	// If true, UtteranceEvents with IsPartial: true are filtered out of the insights
	// target.
	FilterPartialResults bool

	// Specifies the name of the custom language model to use when processing a
	// transcription. Note that language model names are case sensitive. The language
	// of the specified language model must match the language code specified in the
	// transcription request. If the languages don't match, the custom language model
	// isn't applied. Language mismatches don't generate errors or warnings. For more
	// information, see Custom language models
	// (https://docs.aws.amazon.com/transcribe/latest/dg/custom-language-models.html)
	// in the Amazon Transcribe Developer Guide.
	LanguageModelName *string

	// Specifies the level of stability to use when you enable partial results
	// stabilization (EnablePartialResultsStabilization). Low stability provides the
	// highest accuracy. High stability transcribes faster, but with slightly lower
	// accuracy. For more information, see Partial-result stabilization
	// (https://docs.aws.amazon.com/transcribe/latest/dg/streaming.html#streaming-partial-result-stabilization)
	// in the Amazon Transcribe Developer Guide.
	PartialResultsStability PartialResultsStability

	// Specifies the types of personally identifiable information (PII) to redact from
	// a transcript. You can include as many types as you'd like, or you can select
	// ALL. To include PiiEntityTypes in your Call Analytics request, you must also
	// include ContentIdentificationType or ContentRedactionType, but you can't include
	// both. Values must be comma-separated and can include: ADDRESS,
	// BANK_ACCOUNT_NUMBER, BANK_ROUTING, CREDIT_DEBIT_CVV, CREDIT_DEBIT_EXPIRY,
	// CREDIT_DEBIT_NUMBER, EMAIL, NAME, PHONE, PIN, SSN, or ALL. Length Constraints:
	// Minimum length of 1. Maximum length of 300.
	PiiEntityTypes *string

	// The settings for a post-call analysis task in an analytics configuration.
	PostCallAnalyticsSettings *PostCallAnalyticsSettings

	// Specifies how to apply a vocabulary filter to a transcript. To replace words
	// with ***, choose mask. To delete words, choose remove. To flag words without
	// changing them, choose tag.
	VocabularyFilterMethod VocabularyFilterMethod

	// Specifies the name of the custom vocabulary filter to use when processing a
	// transcription. Note that vocabulary filter names are case sensitive. If the
	// language of the specified custom vocabulary filter doesn't match the language
	// identified in your media, the vocabulary filter is not applied to your
	// transcription. For more information, see Using vocabulary filtering with
	// unwanted words
	// (https://docs.aws.amazon.com/transcribe/latest/dg/vocabulary-filtering.html) in
	// the Amazon Transcribe Developer Guide. Length Constraints: Minimum length of 1.
	// Maximum length of 200.
	VocabularyFilterName *string

	// Specifies the name of the custom vocabulary to use when processing a
	// transcription. Note that vocabulary names are case sensitive. If the language of
	// the specified custom vocabulary doesn't match the language identified in your
	// media, the custom vocabulary is not applied to your transcription. For more
	// information, see Custom vocabularies
	// (https://docs.aws.amazon.com/transcribe/latest/dg/custom-vocabulary.html) in the
	// Amazon Transcribe Developer Guide. Length Constraints: Minimum length of 1.
	// Maximum length of 200.
	VocabularyName *string

	noSmithyDocumentSerde
}

// A structure that contains the configuration settings for an Amazon Transcribe
// processor.
type AmazonTranscribeProcessorConfiguration struct {

	// The language code that represents the language spoken in your audio. If you're
	// unsure of the language spoken in your audio, consider using IdentifyLanguage to
	// enable automatic language identification. For a list of languages that real-time
	// Call Analytics supports, see the Supported languages table
	// (https://docs.aws.amazon.com/transcribe/latest/dg/supported-languages.html) in
	// the Amazon Transcribe Developer Guide.
	//
	// This member is required.
	LanguageCode CallAnalyticsLanguageCode

	// Labels all personally identifiable information (PII) identified in your
	// transcript. Content identification is performed at the segment level; PII
	// specified in PiiEntityTypes is flagged upon complete transcription of an audio
	// segment. You can’t set ContentIdentificationType and ContentRedactionType in the
	// same request. If you set both, your request returns a BadRequestException. For
	// more information, see Redacting or identifying personally identifiable
	// information
	// (https://docs.aws.amazon.com/transcribe/latest/dg/pii-redaction.html) in the
	// Amazon Transcribe Developer Guide.
	ContentIdentificationType ContentType

	// Redacts all personally identifiable information (PII) identified in your
	// transcript. Content redaction is performed at the segment level; PII specified
	// in PiiEntityTypes is redacted upon complete transcription of an audio segment.
	// You can’t set ContentRedactionType and ContentIdentificationType in the same
	// request. If you set both, your request returns a BadRequestException. For more
	// information, see Redacting or identifying personally identifiable information
	// (https://docs.aws.amazon.com/transcribe/latest/dg/pii-redaction.html) in the
	// Amazon Transcribe Developer Guide.
	ContentRedactionType ContentType

	// Enables partial result stabilization for your transcription. Partial result
	// stabilization can reduce latency in your output, but may impact accuracy. For
	// more information, see Partial-result stabilization
	// (https://docs.aws.amazon.com/transcribe/latest/dg/streaming.html#streaming-partial-result-stabilization)
	// in the Amazon Transcribe Developer Guide.
	EnablePartialResultsStabilization bool

	// If true, TranscriptEvents with IsPartial: true are filtered out of the insights
	// target.
	FilterPartialResults bool

	// The name of the custom language model that you want to use when processing your
	// transcription. Note that language model names are case sensitive. The language
	// of the specified language model must match the language code you specify in your
	// transcription request. If the languages don't match, the custom language model
	// isn't applied. There are no errors or warnings associated with a language
	// mismatch. For more information, see Custom language models
	// (https://docs.aws.amazon.com/transcribe/latest/dg/custom-language-models.html)
	// in the Amazon Transcribe Developer Guide.
	LanguageModelName *string

	// The level of stability to use when you enable partial results stabilization
	// (EnablePartialResultsStabilization). Low stability provides the highest
	// accuracy. High stability transcribes faster, but with slightly lower accuracy.
	// For more information, see Partial-result stabilization
	// (https://docs.aws.amazon.com/transcribe/latest/dg/streaming.html#streaming-partial-result-stabilization)
	// in the Amazon Transcribe Developer Guide.
	PartialResultsStability PartialResultsStability

	// The types of personally identifiable information (PII) to redact from a
	// transcript. You can include as many types as you'd like, or you can select ALL.
	// To include PiiEntityTypes in your Call Analytics request, you must also include
	// ContentIdentificationType or ContentRedactionType, but you can't include both.
	// Values must be comma-separated and can include: ADDRESS, BANK_ACCOUNT_NUMBER,
	// BANK_ROUTING, CREDIT_DEBIT_CVV, CREDIT_DEBIT_EXPIRY, CREDIT_DEBIT_NUMBER, EMAIL,
	// NAME, PHONE, PIN, SSN, or ALL. Length Constraints: Minimum length of 1. Maximum
	// length of 300.
	PiiEntityTypes *string

	// Enables speaker partitioning (diarization) in your transcription output. Speaker
	// partitioning labels the speech from individual speakers in your media file. For
	// more information, see Partitioning speakers (diarization)
	// (https://docs.aws.amazon.com/transcribe/latest/dg/diarization.html) in the
	// Amazon Transcribe Developer Guide.
	ShowSpeakerLabel bool

	// The vocabulary filtering method used in your Call Analytics transcription.
	VocabularyFilterMethod VocabularyFilterMethod

	// The name of the custom vocabulary filter that you specified in your Call
	// Analytics request. Length Constraints: Minimum length of 1. Maximum length of
	// 200.
	VocabularyFilterName *string

	// The name of the custom vocabulary that you specified in your Call Analytics
	// request. Length Constraints: Minimum length of 1. Maximum length of 200.
	VocabularyName *string

	noSmithyDocumentSerde
}

// The configuration for the artifacts concatenation.
type ArtifactsConcatenationConfiguration struct {

	// The configuration for the audio artifacts concatenation.
	//
	// This member is required.
	Audio *AudioConcatenationConfiguration

	// The configuration for the composited video artifacts concatenation.
	//
	// This member is required.
	CompositedVideo *CompositedVideoConcatenationConfiguration

	// The configuration for the content artifacts concatenation.
	//
	// This member is required.
	Content *ContentConcatenationConfiguration

	// The configuration for the data channel artifacts concatenation.
	//
	// This member is required.
	DataChannel *DataChannelConcatenationConfiguration

	// The configuration for the meeting events artifacts concatenation.
	//
	// This member is required.
	MeetingEvents *MeetingEventsConcatenationConfiguration

	// The configuration for the transcription messages artifacts concatenation.
	//
	// This member is required.
	TranscriptionMessages *TranscriptionMessagesConcatenationConfiguration

	// The configuration for the video artifacts concatenation.
	//
	// This member is required.
	Video *VideoConcatenationConfiguration

	noSmithyDocumentSerde
}

// The configuration for the artifacts.
type ArtifactsConfiguration struct {

	// The configuration for the audio artifacts.
	//
	// This member is required.
	Audio *AudioArtifactsConfiguration

	// The configuration for the content artifacts.
	//
	// This member is required.
	Content *ContentArtifactsConfiguration

	// The configuration for the video artifacts.
	//
	// This member is required.
	Video *VideoArtifactsConfiguration

	// Enables video compositing.
	CompositedVideo *CompositedVideoArtifactsConfiguration

	noSmithyDocumentSerde
}

// The audio artifact configuration object.
type AudioArtifactsConfiguration struct {

	// The MUX type of the audio artifact configuration object.
	//
	// This member is required.
	MuxType AudioMuxType

	noSmithyDocumentSerde
}

// The audio artifact concatenation configuration object.
type AudioConcatenationConfiguration struct {

	// Enables or disables the configuration object.
	//
	// This member is required.
	State AudioArtifactsConcatenationState

	noSmithyDocumentSerde
}

// Defines an audio channel in a Kinesis video stream.
type ChannelDefinition struct {

	// The channel ID.
	//
	// This member is required.
	ChannelId int32

	// Specifies whether the audio in a channel belongs to the AGENT or CUSTOMER.
	ParticipantRole ParticipantRole

	noSmithyDocumentSerde
}

// The configuration object of the Amazon Chime SDK meeting concatenation for a
// specified media pipeline.
type ChimeSdkMeetingConcatenationConfiguration struct {

	// The configuration for the artifacts in an Amazon Chime SDK meeting
	// concatenation.
	//
	// This member is required.
	ArtifactsConfiguration *ArtifactsConcatenationConfiguration

	noSmithyDocumentSerde
}

// The configuration object of the Amazon Chime SDK meeting for a specified media
// pipeline. SourceType must be ChimeSdkMeeting.
type ChimeSdkMeetingConfiguration struct {

	// The configuration for the artifacts in an Amazon Chime SDK meeting.
	ArtifactsConfiguration *ArtifactsConfiguration

	// The source configuration for a specified media pipeline.
	SourceConfiguration *SourceConfiguration

	noSmithyDocumentSerde
}

// The media pipeline's configuration object.
type ChimeSdkMeetingLiveConnectorConfiguration struct {

	// The configuration object's Chime SDK meeting ARN.
	//
	// This member is required.
	Arn *string

	// The configuration object's multiplex type.
	//
	// This member is required.
	MuxType LiveConnectorMuxType

	// The media pipeline's composited video.
	CompositedVideo *CompositedVideoArtifactsConfiguration

	// The source configuration settings of the media pipeline's configuration object.
	SourceConfiguration *SourceConfiguration

	noSmithyDocumentSerde
}

// Specifies the configuration for compositing video artifacts.
type CompositedVideoArtifactsConfiguration struct {

	// The GridView configuration setting.
	//
	// This member is required.
	GridViewConfiguration *GridViewConfiguration

	// The layout setting, such as GridView in the configuration object.
	Layout LayoutOption

	// The video resolution setting in the configuration object. Default: HD at 1280 x
	// 720. FHD resolution: 1920 x 1080.
	Resolution ResolutionOption

	noSmithyDocumentSerde
}

// The composited video configuration object for a specified media pipeline.
// SourceType must be ChimeSdkMeeting.
type CompositedVideoConcatenationConfiguration struct {

	// Enables or disables the configuration object.
	//
	// This member is required.
	State ArtifactsConcatenationState

	noSmithyDocumentSerde
}

// The data sink of the configuration object.
type ConcatenationSink struct {

	// The configuration settings for an Amazon S3 bucket sink.
	//
	// This member is required.
	S3BucketSinkConfiguration *S3BucketSinkConfiguration

	// The type of data sink in the configuration object.
	//
	// This member is required.
	Type ConcatenationSinkType

	noSmithyDocumentSerde
}

// The source type and media pipeline configuration settings in a configuration
// object.
type ConcatenationSource struct {

	// The concatenation settings for the media pipeline in a configuration object.
	//
	// This member is required.
	MediaCapturePipelineSourceConfiguration *MediaCapturePipelineSourceConfiguration

	// The type of concatenation source in a configuration object.
	//
	// This member is required.
	Type ConcatenationSourceType

	noSmithyDocumentSerde
}

// The content artifact object.
type ContentArtifactsConfiguration struct {

	// Indicates whether the content artifact is enabled or disabled.
	//
	// This member is required.
	State ArtifactsState

	// The MUX type of the artifact configuration.
	MuxType ContentMuxType

	noSmithyDocumentSerde
}

// The composited content configuration object for a specified media pipeline.
type ContentConcatenationConfiguration struct {

	// Enables or disables the configuration object.
	//
	// This member is required.
	State ArtifactsConcatenationState

	noSmithyDocumentSerde
}

// The content configuration object's data channel.
type DataChannelConcatenationConfiguration struct {

	// Enables or disables the configuration object.
	//
	// This member is required.
	State ArtifactsConcatenationState

	noSmithyDocumentSerde
}

// Describes the timestamp range and timestamp origin of a range of fragments. Only
// fragments with a start timestamp greater than or equal to the given start time
// and less than or equal to the end time are returned. For example, say a stream
// contains fragments with the following start timestamps:
//
// * 00:00:00
//
// *
// 00:00:02
//
// * 00:00:04
//
// * 00:00:06
//
// A fragment selector range with a start time of
// 00:00:01 and end time of 00:00:04 would return the fragments with start times of
// 00:00:02 and 00:00:04.
type FragmentSelector struct {

	// The origin of the timestamps to use, Server or Producer. For more information,
	// see StartSelectorType in the Amazon Kinesis Video Streams Developer Guide.
	//
	// This member is required.
	FragmentSelectorType FragmentSelectorType

	// The range of timestamps to return.
	//
	// This member is required.
	TimestampRange *TimestampRange

	noSmithyDocumentSerde
}

// Specifies the type of grid layout.
type GridViewConfiguration struct {

	// Defines the layout of the video tiles when content sharing is enabled.
	//
	// This member is required.
	ContentShareLayout ContentShareLayoutOption

	// Defines the configuration options for a presenter only video tile.
	PresenterOnlyConfiguration *PresenterOnlyConfiguration

	noSmithyDocumentSerde
}

// A structure that contains the configuration settings for an issue detection
// task.
type IssueDetectionConfiguration struct {

	// The name of the issue detection rule.
	//
	// This member is required.
	RuleName *string

	noSmithyDocumentSerde
}

// A structure that contains the settings for a keyword match task.
type KeywordMatchConfiguration struct {

	// The keywords or phrases that you want to match.
	//
	// This member is required.
	Keywords []string

	// The name of the keyword match rule.
	//
	// This member is required.
	RuleName *string

	// Matches keywords or phrases on their presence or absence. If set to TRUE, the
	// rule matches when all the specified keywords or phrases are absent. Default:
	// FALSE.
	Negate bool

	noSmithyDocumentSerde
}

// A structure that contains the configuration settings for a Kinesis Data Stream
// sink.
type KinesisDataStreamSinkConfiguration struct {

	// The URL of the sink, https://aws.amazon.com/kinesis/data-streams/
	// (https://aws.amazon.com/kinesis/data-streams/).
	InsightsTarget *string

	noSmithyDocumentSerde
}

// A structure that contains the runtime settings for recording a Kinesis video
// stream.
type KinesisVideoStreamRecordingSourceRuntimeConfiguration struct {

	// Describes the timestamp range and timestamp origin of a range of fragments in
	// the Kinesis video stream.
	//
	// This member is required.
	FragmentSelector *FragmentSelector

	// The stream or streams to be recorded.
	//
	// This member is required.
	Streams []RecordingStreamConfiguration

	noSmithyDocumentSerde
}

// The runtime configuration settings for the Kinesis video stream source.
type KinesisVideoStreamSourceRuntimeConfiguration struct {

	// Specifies the encoding of your input audio. Supported format: PCM (only signed
	// 16-bit little-endian audio formats, which does not include WAV) For more
	// information, see Media formats
	// (https://docs.aws.amazon.com/transcribe/latest/dg/how-input.html#how-input-audio)
	// in the Amazon Transcribe Developer Guide.
	//
	// This member is required.
	MediaEncoding MediaEncoding

	// The sample rate of the input audio (in hertz). Low-quality audio, such as
	// telephone audio, is typically around 8,000 Hz. High-quality audio typically
	// ranges from 16,000 Hz to 48,000 Hz. Note that the sample rate you specify must
	// match that of your audio. Valid Range: Minimum value of 8000. Maximum value of
	// 48000.
	//
	// This member is required.
	MediaSampleRate *int32

	// The streams in the source runtime configuration of a Kinesis video stream.
	//
	// This member is required.
	Streams []StreamConfiguration

	noSmithyDocumentSerde
}

// A structure that contains the configuration settings for an AWS Lambda
// function's data sink.
type LambdaFunctionSinkConfiguration struct {

	// The URL of the sink, https://aws.amazon.com/kinesis/data-streams/
	// (https://aws.amazon.com/kinesis/data-streams/).
	InsightsTarget *string

	noSmithyDocumentSerde
}

// The media pipeline's RTMP configuration object.
type LiveConnectorRTMPConfiguration struct {

	// The URL of the RTMP configuration.
	//
	// This member is required.
	Url *string

	// The audio channels set for the RTMP configuration
	AudioChannels AudioChannelsOption

	// The audio sample rate set for the RTMP configuration. Default: 48000.
	AudioSampleRate *string

	noSmithyDocumentSerde
}

// The media pipeline's sink configuration settings.
type LiveConnectorSinkConfiguration struct {

	// The sink configuration's RTMP configuration settings.
	//
	// This member is required.
	RTMPConfiguration *LiveConnectorRTMPConfiguration

	// The sink configuration's sink type.
	//
	// This member is required.
	SinkType LiveConnectorSinkType

	noSmithyDocumentSerde
}

// The data source configuration object of a streaming media pipeline.
type LiveConnectorSourceConfiguration struct {

	// The configuration settings of the connector pipeline.
	//
	// This member is required.
	ChimeSdkMeetingLiveConnectorConfiguration *ChimeSdkMeetingLiveConnectorConfiguration

	// The source configuration's media source type.
	//
	// This member is required.
	SourceType LiveConnectorSourceType

	noSmithyDocumentSerde
}

// A media pipeline object consisting of an ID, source type, source ARN, a sink
// type, a sink ARN, and a configuration object.
type MediaCapturePipeline struct {

	// The configuration for a specified media pipeline. SourceType must be
	// ChimeSdkMeeting.
	ChimeSdkMeetingConfiguration *ChimeSdkMeetingConfiguration

	// The time at which the pipeline was created, in ISO 8601 format.
	CreatedTimestamp *time.Time

	// The ARN of the media capture pipeline
	MediaPipelineArn *string

	// The ID of a media pipeline.
	MediaPipelineId *string

	// ARN of the destination to which the media artifacts are saved.
	SinkArn *string

	// Destination type to which the media artifacts are saved. You must use an S3
	// Bucket.
	SinkType MediaPipelineSinkType

	// ARN of the source from which the media artifacts are saved.
	SourceArn *string

	// Source type from which media artifacts are saved. You must use ChimeMeeting.
	SourceType MediaPipelineSourceType

	// The status of the media pipeline.
	Status MediaPipelineStatus

	// The time at which the pipeline was updated, in ISO 8601 format.
	UpdatedTimestamp *time.Time

	noSmithyDocumentSerde
}

// The source configuration object of a media capture pipeline.
type MediaCapturePipelineSourceConfiguration struct {

	// The meeting configuration settings in a media capture pipeline configuration
	// object.
	//
	// This member is required.
	ChimeSdkMeetingConfiguration *ChimeSdkMeetingConcatenationConfiguration

	// The media pipeline ARN in the configuration object of a media capture pipeline.
	//
	// This member is required.
	MediaPipelineArn *string

	noSmithyDocumentSerde
}

// The summary data of a media capture pipeline.
type MediaCapturePipelineSummary struct {

	// The ARN of the media pipeline in the summary.
	MediaPipelineArn *string

	// The ID of the media pipeline in the summary.
	MediaPipelineId *string

	noSmithyDocumentSerde
}

// Concatenates audio and video data from one or more data streams.
type MediaConcatenationPipeline struct {

	// The time at which the concatenation pipeline was created.
	CreatedTimestamp *time.Time

	// The ARN of the media pipeline that you specify in the SourceConfiguration
	// object.
	MediaPipelineArn *string

	// The ID of the media pipeline being concatenated.
	MediaPipelineId *string

	// The data sinks of the concatenation pipeline.
	Sinks []ConcatenationSink

	// The data sources being concatenated.
	Sources []ConcatenationSource

	// The status of the concatenation pipeline.
	Status MediaPipelineStatus

	// The time at which the concatenation pipeline was last updated.
	UpdatedTimestamp *time.Time

	noSmithyDocumentSerde
}

// A media pipeline that streams call analytics data.
type MediaInsightsPipeline struct {

	// The time at which the media insights pipeline was created.
	CreatedTimestamp *time.Time

	// The runtime configuration settings for a Kinesis recording video stream in a
	// media insights pipeline.
	KinesisVideoStreamRecordingSourceRuntimeConfiguration *KinesisVideoStreamRecordingSourceRuntimeConfiguration

	// The configuration settings for a Kinesis runtime video stream in a media
	// insights pipeline.
	KinesisVideoStreamSourceRuntimeConfiguration *KinesisVideoStreamSourceRuntimeConfiguration

	// The ARN of a media insight pipeline's configuration settings.
	MediaInsightsPipelineConfigurationArn *string

	// The runtime metadata of a media insights pipeline.
	MediaInsightsRuntimeMetadata map[string]string

	// The ARN of a media insights pipeline.
	MediaPipelineArn *string

	// The ID of a media insights pipeline.
	MediaPipelineId *string

	// The runtime configuration of the Amazon S3 bucket that stores recordings in a
	// media insights pipeline.
	S3RecordingSinkRuntimeConfiguration *S3RecordingSinkRuntimeConfiguration

	// The status of a media insights pipeline.
	Status MediaPipelineStatus

	noSmithyDocumentSerde
}

// A structure that contains the configuration settings for a media insights
// pipeline.
type MediaInsightsPipelineConfiguration struct {

	// The time at which the configuration was created.
	CreatedTimestamp *time.Time

	// The elements in the configuration.
	Elements []MediaInsightsPipelineConfigurationElement

	// The ARN of the configuration.
	MediaInsightsPipelineConfigurationArn *string

	// The ID of the configuration.
	MediaInsightsPipelineConfigurationId *string

	// The name of the configuration.
	MediaInsightsPipelineConfigurationName *string

	// Lists the rules that trigger a real-time alert.
	RealTimeAlertConfiguration *RealTimeAlertConfiguration

	// The ARN of the role used by the service to access Amazon Web Services resources.
	ResourceAccessRoleArn *string

	// The time at which the configuration was last updated.
	UpdatedTimestamp *time.Time

	noSmithyDocumentSerde
}

// An element in a media insights pipeline configuration.
type MediaInsightsPipelineConfigurationElement struct {

	// The element type.
	//
	// This member is required.
	Type MediaInsightsPipelineConfigurationElementType

	// The analytics configuration settings for transcribing audio in a media insights
	// pipeline configuration element.
	AmazonTranscribeCallAnalyticsProcessorConfiguration *AmazonTranscribeCallAnalyticsProcessorConfiguration

	// The transcription processor configuration settings in a media insights pipeline
	// configuration element.
	AmazonTranscribeProcessorConfiguration *AmazonTranscribeProcessorConfiguration

	// The configuration settings for the Kinesis Data Stream Sink in a media insights
	// pipeline configuration element.
	KinesisDataStreamSinkConfiguration *KinesisDataStreamSinkConfiguration

	// The configuration settings for the Amazon Web Services Lambda sink in a media
	// insights pipeline configuration element.
	LambdaFunctionSinkConfiguration *LambdaFunctionSinkConfiguration

	// The configuration settings for the Amazon S3 recording bucket in a media
	// insights pipeline configuration element.
	S3RecordingSinkConfiguration *S3RecordingSinkConfiguration

	// The configuration settings for an SNS topic sink in a media insights pipeline
	// configuration element.
	SnsTopicSinkConfiguration *SnsTopicSinkConfiguration

	// The configuration settings for an SQS queue sink in a media insights pipeline
	// configuration element.
	SqsQueueSinkConfiguration *SqsQueueSinkConfiguration

	// The voice analytics configuration settings in a media insights pipeline
	// configuration element.
	VoiceAnalyticsProcessorConfiguration *VoiceAnalyticsProcessorConfiguration

	noSmithyDocumentSerde
}

// A summary of the media insights pipeline configuration.
type MediaInsightsPipelineConfigurationSummary struct {

	// The ARN of the media insights pipeline configuration.
	MediaInsightsPipelineConfigurationArn *string

	// The ID of the media insights pipeline configuration.
	MediaInsightsPipelineConfigurationId *string

	// The name of the media insights pipeline configuration.
	MediaInsightsPipelineConfigurationName *string

	noSmithyDocumentSerde
}

// The connector pipeline.
type MediaLiveConnectorPipeline struct {

	// The time at which the connector pipeline was created.
	CreatedTimestamp *time.Time

	// The connector pipeline's ARN.
	MediaPipelineArn *string

	// The connector pipeline's ID.
	MediaPipelineId *string

	// The connector pipeline's data sinks.
	Sinks []LiveConnectorSinkConfiguration

	// The connector pipeline's data sources.
	Sources []LiveConnectorSourceConfiguration

	// The connector pipeline's status.
	Status MediaPipelineStatus

	// The time at which the connector pipeline was last updated.
	UpdatedTimestamp *time.Time

	noSmithyDocumentSerde
}

// A pipeline consisting of a media capture, media concatenation, or live-streaming
// pipeline.
type MediaPipeline struct {

	// A pipeline that enables users to capture audio and video.
	MediaCapturePipeline *MediaCapturePipeline

	// The media concatenation pipeline in a media pipeline.
	MediaConcatenationPipeline *MediaConcatenationPipeline

	// The media insights pipeline of a media pipeline.
	MediaInsightsPipeline *MediaInsightsPipeline

	// The connector pipeline of the media pipeline.
	MediaLiveConnectorPipeline *MediaLiveConnectorPipeline

	noSmithyDocumentSerde
}

// The summary of the media pipeline.
type MediaPipelineSummary struct {

	// The ARN of the media pipeline in the summary.
	MediaPipelineArn *string

	// The ID of the media pipeline in the summary.
	MediaPipelineId *string

	noSmithyDocumentSerde
}

// The configuration object for an event concatenation pipeline.
type MeetingEventsConcatenationConfiguration struct {

	// Enables or disables the configuration object.
	//
	// This member is required.
	State ArtifactsConcatenationState

	noSmithyDocumentSerde
}

// The settings for a post-call voice analytics task.
type PostCallAnalyticsSettings struct {

	// The ARN of the role used by Amazon Web Services Transcribe to upload your post
	// call analysis. For more information, see Post-call analytics with real-time
	// transcriptions
	// (https://docs.aws.amazon.com/transcribe/latest/dg/tca-post-call.html) in the
	// Amazon Transcribe Developer Guide.
	//
	// This member is required.
	DataAccessRoleArn *string

	// The URL of the Amazon S3 bucket that contains the post-call data.
	//
	// This member is required.
	OutputLocation *string

	// The content redaction output settings for a post-call analysis task.
	ContentRedactionOutput ContentRedactionOutput

	// The ID of the KMS (Key Management System) key used to encrypt the output.
	OutputEncryptionKMSKeyId *string

	noSmithyDocumentSerde
}

// Defines the configuration for a presenter only video tile.
type PresenterOnlyConfiguration struct {

	// Defines the position of the presenter video tile. Default: TopRight.
	PresenterPosition PresenterPosition

	noSmithyDocumentSerde
}

// A structure that contains the configuration settings for real-time alerts.
type RealTimeAlertConfiguration struct {

	// Turns off real-time alerts.
	Disabled bool

	// The rules in the alert. Rules specify the words or phrases that you want to be
	// notified about.
	Rules []RealTimeAlertRule

	noSmithyDocumentSerde
}

// Specifies the words or phrases that trigger an alert.
type RealTimeAlertRule struct {

	// The type of alert rule.
	//
	// This member is required.
	Type RealTimeAlertRuleType

	// Specifies the issue detection settings for a real-time alert rule.
	IssueDetectionConfiguration *IssueDetectionConfiguration

	// Specifies the settings for matching the keywords in a real-time alert rule.
	KeywordMatchConfiguration *KeywordMatchConfiguration

	// Specifies the settings for predicting sentiment in a real-time alert rule.
	SentimentConfiguration *SentimentConfiguration

	noSmithyDocumentSerde
}

// A structure the holds the settings for recording audio and video.
type RecordingStreamConfiguration struct {

	// The ARN of the recording stream.
	StreamArn *string

	noSmithyDocumentSerde
}

// The configuration settings for the S3 bucket.
type S3BucketSinkConfiguration struct {

	// The destination URL of the S3 bucket.
	//
	// This member is required.
	Destination *string

	noSmithyDocumentSerde
}

// The structure that holds the settings for transmitting audio and video to the
// Amazon S3 bucket.
type S3RecordingSinkConfiguration struct {

	// The URL of the Amazon S3 bucket used as the recording sink.
	Destination *string

	noSmithyDocumentSerde
}

// A structure that holds the settings for transmitting audio and video recordings
// to the runtime Amazon S3 bucket.
type S3RecordingSinkRuntimeConfiguration struct {

	// The URL of the S3 bucket used as the runtime sink.
	//
	// This member is required.
	Destination *string

	// The file formats for the audio and video files sent to the Amazon S3 bucket.
	//
	// This member is required.
	RecordingFileFormat RecordingFileFormat

	noSmithyDocumentSerde
}

// The video streams for a specified media pipeline. The total number of video
// streams can't exceed 25.
type SelectedVideoStreams struct {

	// The attendee IDs of the streams selected for a media pipeline.
	AttendeeIds []string

	// The external user IDs of the streams selected for a media pipeline.
	ExternalUserIds []string

	noSmithyDocumentSerde
}

// A structure that contains the configuration settings for a sentiment analysis
// task.
type SentimentConfiguration struct {

	// The name of the rule in the sentiment configuration.
	//
	// This member is required.
	RuleName *string

	// The type of sentiment, POSITIVE, NEGATIVE, or NEUTRAL.
	//
	// This member is required.
	SentimentType SentimentType

	// Specifies the analysis interval.
	//
	// This member is required.
	TimePeriod int32

	noSmithyDocumentSerde
}

// The configuration settings for the SNS topic sink.
type SnsTopicSinkConfiguration struct {

	// The URL of the SNS sink, https://aws.amazon.com/kinesis/data-streams/
	// (https://aws.amazon.com/kinesis/data-streams/).
	InsightsTarget *string

	noSmithyDocumentSerde
}

// Source configuration for a specified media pipeline.
type SourceConfiguration struct {

	// The selected video streams for a specified media pipeline. The number of video
	// streams can't exceed 25.
	SelectedVideoStreams *SelectedVideoStreams

	noSmithyDocumentSerde
}

// The URL of the SQS sink.
type SqsQueueSinkConfiguration struct {

	// The URL of the SQS sink, https://aws.amazon.com/kinesis/data-streams/
	// (https://aws.amazon.com/kinesis/data-streams/).
	InsightsTarget *string

	noSmithyDocumentSerde
}

// Defines a streaming channel.
type StreamChannelDefinition struct {

	// The number of channels in a streaming channel.
	//
	// This member is required.
	NumberOfChannels *int32

	// The definitions of the channels in a streaming channel.
	ChannelDefinitions []ChannelDefinition

	noSmithyDocumentSerde
}

// The configuration settings for a stream.
type StreamConfiguration struct {

	// The ARN of the stream.
	//
	// This member is required.
	StreamArn *string

	// The streaming channel definition in the stream configuration.
	//
	// This member is required.
	StreamChannelDefinition *StreamChannelDefinition

	// The unique identifier of the fragment to begin processing.
	FragmentNumber *string

	noSmithyDocumentSerde
}

// A key/value pair that grants users access to meeting resources.
type Tag struct {

	// The key half of a tag.
	//
	// This member is required.
	Key *string

	// The value half of a tag.
	//
	// This member is required.
	Value *string

	noSmithyDocumentSerde
}

// The range of timestamps to return.
type TimestampRange struct {

	// The ending timestamp for the specified range.
	//
	// This member is required.
	EndTimestamp *time.Time

	// The starting timestamp for the specified range.
	//
	// This member is required.
	StartTimestamp *time.Time

	noSmithyDocumentSerde
}

// The configuration object for concatenating transcription messages.
type TranscriptionMessagesConcatenationConfiguration struct {

	// Enables or disables the configuration object.
	//
	// This member is required.
	State ArtifactsConcatenationState

	noSmithyDocumentSerde
}

// The video artifact configuration object.
type VideoArtifactsConfiguration struct {

	// Indicates whether the video artifact is enabled or disabled.
	//
	// This member is required.
	State ArtifactsState

	// The MUX type of the video artifact configuration object.
	MuxType VideoMuxType

	noSmithyDocumentSerde
}

// The configuration object of a video concatenation pipeline.
type VideoConcatenationConfiguration struct {

	// Enables or disables the configuration object.
	//
	// This member is required.
	State ArtifactsConcatenationState

	noSmithyDocumentSerde
}

// The configuration settings for a voice analytics processor.
type VoiceAnalyticsProcessorConfiguration struct {

	// The status of the speaker search task.
	SpeakerSearchStatus VoiceAnalyticsConfigurationStatus

	// The status of the voice tone analysis task.
	VoiceToneAnalysisStatus VoiceAnalyticsConfigurationStatus

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
