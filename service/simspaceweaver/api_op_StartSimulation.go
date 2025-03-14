// Code generated by smithy-go-codegen DO NOT EDIT.

package simspaceweaver

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/simspaceweaver/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Starts a simulation with the given name and schema.
func (c *Client) StartSimulation(ctx context.Context, params *StartSimulationInput, optFns ...func(*Options)) (*StartSimulationOutput, error) {
	if params == nil {
		params = &StartSimulationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartSimulation", params, optFns, c.addOperationStartSimulationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartSimulationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartSimulationInput struct {

	// The name of the simulation.
	//
	// This member is required.
	Name *string

	// The Amazon Resource Name (ARN) of the Identity and Access Management (IAM) role
	// that the simulation assumes to perform actions. For more information about ARNs,
	// see Amazon Resource Names (ARNs)
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in
	// the Amazon Web Services General Reference. For more information about IAM roles,
	// see IAM roles (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles.html)
	// in the Identity and Access Management User Guide.
	//
	// This member is required.
	RoleArn *string

	// The location of the simulation schema in Amazon Simple Storage Service (Amazon
	// S3). For more information about Amazon S3, see the  Amazon Simple Storage
	// Service User Guide
	// (https://docs.aws.amazon.com/AmazonS3/latest/userguide/Welcome.html).
	//
	// This member is required.
	SchemaS3Location *types.S3Location

	// A value that you provide to ensure that repeated calls to this API operation
	// using the same parameters complete only once. A ClientToken is also known as an
	// idempotency token. A ClientToken expires after 24 hours.
	ClientToken *string

	// The description of the simulation.
	Description *string

	// The maximum running time of the simulation, specified as a number of months (m
	// or M), hours (h or H), or days (d or D). The simulation stops when it reaches
	// this limit.
	MaximumDuration *string

	// A list of tags for the simulation. For more information about tags, see Tagging
	// Amazon Web Services resources
	// (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in the Amazon
	// Web Services General Reference.
	Tags map[string]string

	noSmithyDocumentSerde
}

type StartSimulationOutput struct {

	// The Amazon Resource Name (ARN) of the simulation. For more information about
	// ARNs, see Amazon Resource Names (ARNs)
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in
	// the Amazon Web Services General Reference.
	Arn *string

	// The time when the simulation was created, expressed as the number of seconds and
	// milliseconds in UTC since the Unix epoch (0:0:0.000, January 1, 1970).
	CreationTime *time.Time

	// A universally unique identifier (UUID) for this simulation.
	ExecutionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartSimulationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartSimulation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartSimulation{}, middleware.After)
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
	if err = addIdempotencyToken_opStartSimulationMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpStartSimulationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartSimulation(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpStartSimulation struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpStartSimulation) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpStartSimulation) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*StartSimulationInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *StartSimulationInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opStartSimulationMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpStartSimulation{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opStartSimulation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "simspaceweaver",
		OperationName: "StartSimulation",
	}
}
