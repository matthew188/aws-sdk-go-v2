// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudcontrol

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/cloudcontrol/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates the specified resource. For more information, see Creating a resource
// (https://docs.aws.amazon.com/cloudcontrolapi/latest/userguide/resource-operations-create.html)
// in the Amazon Web Services Cloud Control API User Guide. After you have
// initiated a resource creation request, you can monitor the progress of your
// request by calling GetResourceRequestStatus
// (https://docs.aws.amazon.com/cloudcontrolapi/latest/APIReference/API_GetResourceRequestStatus.html)
// using the RequestToken of the ProgressEvent type returned by CreateResource.
func (c *Client) CreateResource(ctx context.Context, params *CreateResourceInput, optFns ...func(*Options)) (*CreateResourceOutput, error) {
	if params == nil {
		params = &CreateResourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateResource", params, optFns, c.addOperationCreateResourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateResourceInput struct {

	// Structured data format representing the desired state of the resource,
	// consisting of that resource's properties and their desired values. Cloud Control
	// API currently supports JSON as a structured data format. Specify the desired
	// state as one of the following:
	//
	// * A JSON blob
	//
	// * A local path containing the
	// desired state in JSON data format
	//
	// For more information, see Composing the
	// desired state of the resource
	// (https://docs.aws.amazon.com/cloudcontrolapi/latest/userguide/resource-operations-create.html#resource-operations-create-desiredstate)
	// in the Amazon Web Services Cloud Control API User Guide. For more information
	// about the properties of a specific resource, refer to the related topic for the
	// resource in the Resource and property types reference
	// (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-template-resource-type-ref.html)
	// in the CloudFormation Users Guide.
	//
	// This member is required.
	DesiredState *string

	// The name of the resource type.
	//
	// This member is required.
	TypeName *string

	// A unique identifier to ensure the idempotency of the resource request. As a best
	// practice, specify this token to ensure idempotency, so that Amazon Web Services
	// Cloud Control API can accurately distinguish between request retries and new
	// resource requests. You might retry a resource request to ensure that it was
	// successfully received. A client token is valid for 36 hours once used. After
	// that, a resource request with the same client token is treated as a new request.
	// If you do not specify a client token, one is generated for inclusion in the
	// request. For more information, see Ensuring resource operation requests are
	// unique
	// (https://docs.aws.amazon.com/cloudcontrolapi/latest/userguide/resource-operations.html#resource-operations-idempotency)
	// in the Amazon Web Services Cloud Control API User Guide.
	ClientToken *string

	// The Amazon Resource Name (ARN) of the Identity and Access Management (IAM) role
	// for Cloud Control API to use when performing this resource operation. The role
	// specified must have the permissions required for this operation. The necessary
	// permissions for each event handler are defined in the handlers
	// (https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/resource-type-schema.html#schema-properties-handlers)
	// section of the resource type definition schema
	// (https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/resource-type-schema.html).
	// If you do not specify a role, Cloud Control API uses a temporary session created
	// using your Amazon Web Services user credentials. For more information, see
	// Specifying credentials
	// (https://docs.aws.amazon.com/cloudcontrolapi/latest/userguide/resource-operations.html#resource-operations-permissions)
	// in the Amazon Web Services Cloud Control API User Guide.
	RoleArn *string

	// For private resource types, the type version to use in this resource operation.
	// If you do not specify a resource version, CloudFormation uses the default
	// version.
	TypeVersionId *string

	noSmithyDocumentSerde
}

type CreateResourceOutput struct {

	// Represents the current status of the resource creation request. After you have
	// initiated a resource creation request, you can monitor the progress of your
	// request by calling GetResourceRequestStatus
	// (https://docs.aws.amazon.com/cloudcontrolapi/latest/APIReference/API_GetResourceRequestStatus.html)
	// using the RequestToken of the ProgressEvent returned by CreateResource.
	ProgressEvent *types.ProgressEvent

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateResourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCreateResource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCreateResource{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateResourceMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateResourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateResource(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateResource struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateResource) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateResourceInput ")
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
func addIdempotencyToken_opCreateResourceMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateResource{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateResource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudcontrolapi",
		OperationName: "CreateResource",
	}
}
