// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehendmedical

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/comprehendmedical/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts an asynchronous medical entity detection job for a collection of
// documents. Use the DescribeEntitiesDetectionV2Job operation to track the status
// of a job.
func (c *Client) StartEntitiesDetectionV2Job(ctx context.Context, params *StartEntitiesDetectionV2JobInput, optFns ...func(*Options)) (*StartEntitiesDetectionV2JobOutput, error) {
	if params == nil {
		params = &StartEntitiesDetectionV2JobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartEntitiesDetectionV2Job", params, optFns, c.addOperationStartEntitiesDetectionV2JobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartEntitiesDetectionV2JobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartEntitiesDetectionV2JobInput struct {

	// The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM)
	// role that grants Comprehend Medical; read access to your input data. For more
	// information, see  Role-Based Permissions Required for Asynchronous Operations
	// (https://docs.aws.amazon.com/comprehend/latest/dg/access-control-managing-permissions-med.html#auth-role-permissions-med).
	//
	// This member is required.
	DataAccessRoleArn *string

	// The input configuration that specifies the format and location of the input data
	// for the job.
	//
	// This member is required.
	InputDataConfig *types.InputDataConfig

	// The language of the input documents. All documents must be in the same language.
	// Comprehend Medical; processes files in US English (en).
	//
	// This member is required.
	LanguageCode types.LanguageCode

	// The output configuration that specifies where to send the output files.
	//
	// This member is required.
	OutputDataConfig *types.OutputDataConfig

	// A unique identifier for the request. If you don't set the client request token,
	// Comprehend Medical; generates one for you.
	ClientRequestToken *string

	// The identifier of the job.
	JobName *string

	// An AWS Key Management Service key to encrypt your output files. If you do not
	// specify a key, the files are written in plain text.
	KMSKey *string

	noSmithyDocumentSerde
}

type StartEntitiesDetectionV2JobOutput struct {

	// The identifier generated for the job. To get the status of a job, use this
	// identifier with the DescribeEntitiesDetectionV2Job operation.
	JobId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartEntitiesDetectionV2JobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartEntitiesDetectionV2Job{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartEntitiesDetectionV2Job{}, middleware.After)
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
	if err = addIdempotencyToken_opStartEntitiesDetectionV2JobMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpStartEntitiesDetectionV2JobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartEntitiesDetectionV2Job(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpStartEntitiesDetectionV2Job struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpStartEntitiesDetectionV2Job) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpStartEntitiesDetectionV2Job) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*StartEntitiesDetectionV2JobInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *StartEntitiesDetectionV2JobInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opStartEntitiesDetectionV2JobMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpStartEntitiesDetectionV2Job{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opStartEntitiesDetectionV2Job(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehendmedical",
		OperationName: "StartEntitiesDetectionV2Job",
	}
}
