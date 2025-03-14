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

// Creates a new ThreatIntelSet. ThreatIntelSets consist of known malicious IP
// addresses. GuardDuty generates findings based on ThreatIntelSets. Only users of
// the administrator account can use this operation.
func (c *Client) CreateThreatIntelSet(ctx context.Context, params *CreateThreatIntelSetInput, optFns ...func(*Options)) (*CreateThreatIntelSetOutput, error) {
	if params == nil {
		params = &CreateThreatIntelSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateThreatIntelSet", params, optFns, c.addOperationCreateThreatIntelSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateThreatIntelSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateThreatIntelSetInput struct {

	// A Boolean value that indicates whether GuardDuty is to start using the uploaded
	// ThreatIntelSet.
	//
	// This member is required.
	Activate bool

	// The unique ID of the detector of the GuardDuty account that you want to create a
	// threatIntelSet for.
	//
	// This member is required.
	DetectorId *string

	// The format of the file that contains the ThreatIntelSet.
	//
	// This member is required.
	Format types.ThreatIntelSetFormat

	// The URI of the file that contains the ThreatIntelSet.
	//
	// This member is required.
	Location *string

	// A user-friendly ThreatIntelSet name displayed in all findings that are generated
	// by activity that involves IP addresses included in this ThreatIntelSet.
	//
	// This member is required.
	Name *string

	// The idempotency token for the create request.
	ClientToken *string

	// The tags to be added to a new threat list resource.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateThreatIntelSetOutput struct {

	// The ID of the ThreatIntelSet resource.
	//
	// This member is required.
	ThreatIntelSetId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateThreatIntelSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateThreatIntelSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateThreatIntelSet{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateThreatIntelSetMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateThreatIntelSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateThreatIntelSet(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateThreatIntelSet struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateThreatIntelSet) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateThreatIntelSet) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateThreatIntelSetInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateThreatIntelSetInput ")
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
func addIdempotencyToken_opCreateThreatIntelSetMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateThreatIntelSet{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateThreatIntelSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "guardduty",
		OperationName: "CreateThreatIntelSet",
	}
}
