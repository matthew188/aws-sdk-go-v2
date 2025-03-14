// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/apigateway/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates a documentation version.
func (c *Client) UpdateDocumentationVersion(ctx context.Context, params *UpdateDocumentationVersionInput, optFns ...func(*Options)) (*UpdateDocumentationVersionOutput, error) {
	if params == nil {
		params = &UpdateDocumentationVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDocumentationVersion", params, optFns, c.addOperationUpdateDocumentationVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDocumentationVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Updates an existing documentation version of an API.
type UpdateDocumentationVersionInput struct {

	// The version identifier of the to-be-updated documentation version.
	//
	// This member is required.
	DocumentationVersion *string

	// The string identifier of the associated RestApi..
	//
	// This member is required.
	RestApiId *string

	// For more information about supported patch operations, see Patch Operations
	// (https://docs.aws.amazon.com/apigateway/latest/api/patch-operations.html).
	PatchOperations []types.PatchOperation

	noSmithyDocumentSerde
}

// A snapshot of the documentation of an API.
type UpdateDocumentationVersionOutput struct {

	// The date when the API documentation snapshot is created.
	CreatedDate *time.Time

	// The description of the API documentation snapshot.
	Description *string

	// The version identifier of the API documentation snapshot.
	Version *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateDocumentationVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateDocumentationVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateDocumentationVersion{}, middleware.After)
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
	if err = addOpUpdateDocumentationVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDocumentationVersion(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addAcceptHeader(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opUpdateDocumentationVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "UpdateDocumentationVersion",
	}
}
