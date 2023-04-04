// Code generated by smithy-go-codegen DO NOT EDIT.

package transcribe

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/transcribe/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Transcribes the audio from a media file and applies any additional Request
// Parameters you choose to include in your request. To make a
// StartTranscriptionJob request, you must first upload your media file into an
// Amazon S3 bucket; you can then specify the Amazon S3 location of the file using
// the Media parameter. You must include the following parameters in your
// StartTranscriptionJob request:
//
// * region: The Amazon Web Services Region where
// you are making your request. For a list of Amazon Web Services Regions supported
// with Amazon Transcribe, refer to Amazon Transcribe endpoints and quotas
// (https://docs.aws.amazon.com/general/latest/gr/transcribe.html).
//
// *
// TranscriptionJobName: A custom name you create for your transcription job that
// is unique within your Amazon Web Services account.
//
// * Media (MediaFileUri): The
// Amazon S3 location of your media file.
//
// * One of LanguageCode, IdentifyLanguage,
// or IdentifyMultipleLanguages: If you know the language of your media file,
// specify it using the LanguageCode parameter; you can find all valid language
// codes in the Supported languages
// (https://docs.aws.amazon.com/transcribe/latest/dg/supported-languages.html)
// table. If you don't know the languages spoken in your media, use either
// IdentifyLanguage or IdentifyMultipleLanguages and let Amazon Transcribe identify
// the languages for you.
func (c *Client) StartTranscriptionJob(ctx context.Context, params *StartTranscriptionJobInput, optFns ...func(*Options)) (*StartTranscriptionJobOutput, error) {
	if params == nil {
		params = &StartTranscriptionJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartTranscriptionJob", params, optFns, c.addOperationStartTranscriptionJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartTranscriptionJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartTranscriptionJobInput struct {

	// Describes the Amazon S3 location of the media file you want to use in your
	// request.
	//
	// This member is required.
	Media *types.Media

	// A unique name, chosen by you, for your transcription job. The name that you
	// specify is also used as the default name of your transcription output file. If
	// you want to specify a different name for your transcription output, use the
	// OutputKey parameter. This name is case sensitive, cannot contain spaces, and
	// must be unique within an Amazon Web Services account. If you try to create a new
	// job with the same name as an existing job, you get a ConflictException error.
	//
	// This member is required.
	TranscriptionJobName *string

	// Makes it possible to redact or flag specified personally identifiable
	// information (PII) in your transcript. If you use ContentRedaction, you must also
	// include the sub-parameters: PiiEntityTypes, RedactionOutput, and RedactionType.
	ContentRedaction *types.ContentRedaction

	// Enables automatic language identification in your transcription job request. Use
	// this parameter if your media file contains only one language. If your media
	// contains multiple languages, use IdentifyMultipleLanguages instead. If you
	// include IdentifyLanguage, you can optionally include a list of language codes,
	// using LanguageOptions, that you think may be present in your media file.
	// Including LanguageOptions restricts IdentifyLanguage to only the language
	// options that you specify, which can improve transcription accuracy. If you want
	// to apply a custom language model, a custom vocabulary, or a custom vocabulary
	// filter to your automatic language identification request, include
	// LanguageIdSettings with the relevant sub-parameters (VocabularyName,
	// LanguageModelName, and VocabularyFilterName). If you include LanguageIdSettings,
	// also include LanguageOptions. Note that you must include one of LanguageCode,
	// IdentifyLanguage, or IdentifyMultipleLanguages in your request. If you include
	// more than one of these parameters, your transcription job fails.
	IdentifyLanguage *bool

	// Enables automatic multi-language identification in your transcription job
	// request. Use this parameter if your media file contains more than one language.
	// If your media contains only one language, use IdentifyLanguage instead. If you
	// include IdentifyMultipleLanguages, you can optionally include a list of language
	// codes, using LanguageOptions, that you think may be present in your media file.
	// Including LanguageOptions restricts IdentifyLanguage to only the language
	// options that you specify, which can improve transcription accuracy. If you want
	// to apply a custom vocabulary or a custom vocabulary filter to your automatic
	// language identification request, include LanguageIdSettings with the relevant
	// sub-parameters (VocabularyName and VocabularyFilterName). If you include
	// LanguageIdSettings, also include LanguageOptions. Note that you must include one
	// of LanguageCode, IdentifyLanguage, or IdentifyMultipleLanguages in your request.
	// If you include more than one of these parameters, your transcription job fails.
	IdentifyMultipleLanguages *bool

	// Makes it possible to control how your transcription job is processed. Currently,
	// the only JobExecutionSettings modification you can choose is enabling job
	// queueing using the AllowDeferredExecution sub-parameter. If you include
	// JobExecutionSettings in your request, you must also include the sub-parameters:
	// AllowDeferredExecution and DataAccessRoleArn.
	JobExecutionSettings *types.JobExecutionSettings

	// A map of plain text, non-secret key:value pairs, known as encryption context
	// pairs, that provide an added layer of security for your data. For more
	// information, see KMS encryption context
	// (https://docs.aws.amazon.com/transcribe/latest/dg/key-management.html#kms-context)
	// and Asymmetric keys in KMS
	// (https://docs.aws.amazon.com/transcribe/latest/dg/symmetric-asymmetric.html).
	KMSEncryptionContext map[string]string

	// The language code that represents the language spoken in the input media file.
	// If you're unsure of the language spoken in your media file, consider using
	// IdentifyLanguage or IdentifyMultipleLanguages to enable automatic language
	// identification. Note that you must include one of LanguageCode,
	// IdentifyLanguage, or IdentifyMultipleLanguages in your request. If you include
	// more than one of these parameters, your transcription job fails. For a list of
	// supported languages and their associated language codes, refer to the Supported
	// languages
	// (https://docs.aws.amazon.com/transcribe/latest/dg/supported-languages.html)
	// table. To transcribe speech in Modern Standard Arabic (ar-SA), your media file
	// must be encoded at a sample rate of 16,000 Hz or higher.
	LanguageCode types.LanguageCode

	// If using automatic language identification in your request and you want to apply
	// a custom language model, a custom vocabulary, or a custom vocabulary filter,
	// include LanguageIdSettings with the relevant sub-parameters (VocabularyName,
	// LanguageModelName, and VocabularyFilterName). Note that multi-language
	// identification (IdentifyMultipleLanguages) doesn't support custom language
	// models. LanguageIdSettings supports two to five language codes. Each language
	// code you include can have an associated custom language model, custom
	// vocabulary, and custom vocabulary filter. The language codes that you specify
	// must match the languages of the associated custom language models, custom
	// vocabularies, and custom vocabulary filters. It's recommended that you include
	// LanguageOptions when using LanguageIdSettings to ensure that the correct
	// language dialect is identified. For example, if you specify a custom vocabulary
	// that is in en-US but Amazon Transcribe determines that the language spoken in
	// your media is en-AU, your custom vocabulary is not applied to your
	// transcription. If you include LanguageOptions and include en-US as the only
	// English language dialect, your custom vocabulary is applied to your
	// transcription. If you want to include a custom language model with your request
	// but do not want to use automatic language identification, use instead the
	// parameter with the LanguageModelName sub-parameter. If you want to include a
	// custom vocabulary or a custom vocabulary filter (or both) with your request but
	// do not want to use automatic language identification, use instead the  parameter
	// with the VocabularyName or VocabularyFilterName (or both) sub-parameter.
	LanguageIdSettings map[string]types.LanguageIdSettings

	// You can specify two or more language codes that represent the languages you
	// think may be present in your media. Including more than five is not recommended.
	// If you're unsure what languages are present, do not include this parameter. If
	// you include LanguageOptions in your request, you must also include
	// IdentifyLanguage. For more information, refer to Supported languages
	// (https://docs.aws.amazon.com/transcribe/latest/dg/supported-languages.html). To
	// transcribe speech in Modern Standard Arabic (ar-SA), your media file must be
	// encoded at a sample rate of 16,000 Hz or higher.
	LanguageOptions []types.LanguageCode

	// Specify the format of your input media file.
	MediaFormat types.MediaFormat

	// The sample rate, in hertz, of the audio track in your input media file. If you
	// don't specify the media sample rate, Amazon Transcribe determines it for you. If
	// you specify the sample rate, it must match the rate detected by Amazon
	// Transcribe. If there's a mismatch between the value that you specify and the
	// value detected, your job fails. In most cases, you can omit MediaSampleRateHertz
	// and let Amazon Transcribe determine the sample rate.
	MediaSampleRateHertz *int32

	// Specify the custom language model you want to include with your transcription
	// job. If you include ModelSettings in your request, you must include the
	// LanguageModelName sub-parameter. For more information, see Custom language
	// models
	// (https://docs.aws.amazon.com/transcribe/latest/dg/custom-language-models.html).
	ModelSettings *types.ModelSettings

	// The name of the Amazon S3 bucket where you want your transcription output
	// stored. Do not include the S3:// prefix of the specified bucket. If you want
	// your output to go to a sub-folder of this bucket, specify it using the OutputKey
	// parameter; OutputBucketName only accepts the name of a bucket. For example, if
	// you want your output stored in S3://DOC-EXAMPLE-BUCKET, set OutputBucketName to
	// DOC-EXAMPLE-BUCKET. However, if you want your output stored in
	// S3://DOC-EXAMPLE-BUCKET/test-files/, set OutputBucketName to DOC-EXAMPLE-BUCKET
	// and OutputKey to test-files/. Note that Amazon Transcribe must have permission
	// to use the specified location. You can change Amazon S3 permissions using the
	// Amazon Web Services Management Console (https://console.aws.amazon.com/s3). See
	// also Permissions Required for IAM User Roles
	// (https://docs.aws.amazon.com/transcribe/latest/dg/security_iam_id-based-policy-examples.html#auth-role-iam-user).
	// If you don't specify OutputBucketName, your transcript is placed in a
	// service-managed Amazon S3 bucket and you are provided with a URI to access your
	// transcript.
	OutputBucketName *string

	// The KMS key you want to use to encrypt your transcription output. If using a key
	// located in the current Amazon Web Services account, you can specify your KMS key
	// in one of four ways:
	//
	// * Use the KMS key ID itself. For example,
	// 1234abcd-12ab-34cd-56ef-1234567890ab.
	//
	// * Use an alias for the KMS key ID. For
	// example, alias/ExampleAlias.
	//
	// * Use the Amazon Resource Name (ARN) for the KMS
	// key ID. For example,
	// arn:aws:kms:region:account-ID:key/1234abcd-12ab-34cd-56ef-1234567890ab.
	//
	// * Use
	// the ARN for the KMS key alias. For example,
	// arn:aws:kms:region:account-ID:alias/ExampleAlias.
	//
	// If using a key located in a
	// different Amazon Web Services account than the current Amazon Web Services
	// account, you can specify your KMS key in one of two ways:
	//
	// * Use the ARN for the
	// KMS key ID. For example,
	// arn:aws:kms:region:account-ID:key/1234abcd-12ab-34cd-56ef-1234567890ab.
	//
	// * Use
	// the ARN for the KMS key alias. For example,
	// arn:aws:kms:region:account-ID:alias/ExampleAlias.
	//
	// If you don't specify an
	// encryption key, your output is encrypted with the default Amazon S3 key
	// (SSE-S3). If you specify a KMS key to encrypt your output, you must also specify
	// an output location using the OutputLocation parameter. Note that the role making
	// the request must have permission to use the specified KMS key.
	OutputEncryptionKMSKeyId *string

	// Use in combination with OutputBucketName to specify the output location of your
	// transcript and, optionally, a unique name for your output file. The default name
	// for your transcription output is the same as the name you specified for your
	// transcription job (TranscriptionJobName). Here are some examples of how you can
	// use OutputKey:
	//
	// * If you specify 'DOC-EXAMPLE-BUCKET' as the OutputBucketName
	// and 'my-transcript.json' as the OutputKey, your transcription output path is
	// s3://DOC-EXAMPLE-BUCKET/my-transcript.json.
	//
	// * If you specify
	// 'my-first-transcription' as the TranscriptionJobName, 'DOC-EXAMPLE-BUCKET' as
	// the OutputBucketName, and 'my-transcript' as the OutputKey, your transcription
	// output path is
	// s3://DOC-EXAMPLE-BUCKET/my-transcript/my-first-transcription.json.
	//
	// * If you
	// specify 'DOC-EXAMPLE-BUCKET' as the OutputBucketName and
	// 'test-files/my-transcript.json' as the OutputKey, your transcription output path
	// is s3://DOC-EXAMPLE-BUCKET/test-files/my-transcript.json.
	//
	// * If you specify
	// 'my-first-transcription' as the TranscriptionJobName, 'DOC-EXAMPLE-BUCKET' as
	// the OutputBucketName, and 'test-files/my-transcript' as the OutputKey, your
	// transcription output path is
	// s3://DOC-EXAMPLE-BUCKET/test-files/my-transcript/my-first-transcription.json.
	//
	// If
	// you specify the name of an Amazon S3 bucket sub-folder that doesn't exist, one
	// is created for you.
	OutputKey *string

	// Specify additional optional settings in your request, including channel
	// identification, alternative transcriptions, speaker partitioning. You can use
	// that to apply custom vocabularies and vocabulary filters. If you want to include
	// a custom vocabulary or a custom vocabulary filter (or both) with your request
	// but do not want to use automatic language identification, use Settings with the
	// VocabularyName or VocabularyFilterName (or both) sub-parameter. If you're using
	// automatic language identification with your request and want to include a custom
	// language model, a custom vocabulary, or a custom vocabulary filter, use instead
	// the  parameter with the LanguageModelName, VocabularyName or
	// VocabularyFilterName sub-parameters.
	Settings *types.Settings

	// Produces subtitle files for your input media. You can specify WebVTT (*.vtt) and
	// SubRip (*.srt) formats.
	Subtitles *types.Subtitles

	// Adds one or more custom tags, each in the form of a key:value pair, to a new
	// transcription job at the time you start this new job. To learn more about using
	// tags with Amazon Transcribe, refer to Tagging resources
	// (https://docs.aws.amazon.com/transcribe/latest/dg/tagging.html).
	Tags []types.Tag

	noSmithyDocumentSerde
}

type StartTranscriptionJobOutput struct {

	// Provides detailed information about the current transcription job, including job
	// status and, if applicable, failure reason.
	TranscriptionJob *types.TranscriptionJob

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartTranscriptionJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartTranscriptionJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartTranscriptionJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpStartTranscriptionJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartTranscriptionJob(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opStartTranscriptionJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transcribe",
		OperationName: "StartTranscriptionJob",
	}
}
