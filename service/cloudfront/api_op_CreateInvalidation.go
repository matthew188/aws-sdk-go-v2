// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Create a new invalidation.
func (c *Client) CreateInvalidation(ctx context.Context, params *CreateInvalidationInput, optFns ...func(*Options)) (*CreateInvalidationOutput, error) {
	if params == nil {
		params = &CreateInvalidationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateInvalidation", params, optFns, c.addOperationCreateInvalidationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateInvalidationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request to create an invalidation.
type CreateInvalidationInput struct {

	// The distribution's id.
	//
	// This member is required.
	DistributionId *string

	// The batch information for the invalidation.
	//
	// This member is required.
	InvalidationBatch *types.InvalidationBatch

	noSmithyDocumentSerde
}

// The returned result of the corresponding request.
type CreateInvalidationOutput struct {

	// The invalidation's information.
	Invalidation *types.Invalidation

	// The fully qualified URI of the distribution and invalidation batch request,
	// including the Invalidation ID.
	Location *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateInvalidationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpCreateInvalidation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpCreateInvalidation{}, middleware.After)
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
	if err = addOpCreateInvalidationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateInvalidation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateInvalidation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "CreateInvalidation",
	}
}
