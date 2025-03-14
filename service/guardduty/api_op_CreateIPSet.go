// Code generated by smithy-go-codegen DO NOT EDIT.

package guardduty

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/guardduty/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new IPSet, which is called a trusted IP list in the console user
// interface. An IPSet is a list of IP addresses that are trusted for secure
// communication with Amazon Web Services infrastructure and applications.
// GuardDuty doesn't generate findings for IP addresses that are included in
// IPSets. Only users from the administrator account can use this operation.
func (c *Client) CreateIPSet(ctx context.Context, params *CreateIPSetInput, optFns ...func(*Options)) (*CreateIPSetOutput, error) {
	if params == nil {
		params = &CreateIPSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateIPSet", params, optFns, c.addOperationCreateIPSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateIPSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateIPSetInput struct {

	// A Boolean value that indicates whether GuardDuty is to start using the uploaded
	// IPSet.
	//
	// This member is required.
	Activate bool

	// The unique ID of the detector of the GuardDuty account that you want to create
	// an IPSet for.
	//
	// This member is required.
	DetectorId *string

	// The format of the file that contains the IPSet.
	//
	// This member is required.
	Format types.IpSetFormat

	// The URI of the file that contains the IPSet.
	//
	// This member is required.
	Location *string

	// The user-friendly name to identify the IPSet. Allowed characters are
	// alphanumeric, whitespace, dash (-), and underscores (_).
	//
	// This member is required.
	Name *string

	// The idempotency token for the create request.
	ClientToken *string

	// The tags to be added to a new IP set resource.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateIPSetOutput struct {

	// The ID of the IPSet resource.
	//
	// This member is required.
	IpSetId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateIPSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateIPSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateIPSet{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateIPSetMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateIPSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateIPSet(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateIPSet struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateIPSet) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateIPSet) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateIPSetInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateIPSetInput ")
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
func addIdempotencyToken_opCreateIPSetMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateIPSet{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateIPSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "guardduty",
		OperationName: "CreateIPSet",
	}
}
