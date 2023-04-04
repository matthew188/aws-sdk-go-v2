// Code generated by smithy-go-codegen DO NOT EDIT.

package codebuild

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/codebuild/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a batch build.
func (c *Client) DeleteBuildBatch(ctx context.Context, params *DeleteBuildBatchInput, optFns ...func(*Options)) (*DeleteBuildBatchOutput, error) {
	if params == nil {
		params = &DeleteBuildBatchInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteBuildBatch", params, optFns, c.addOperationDeleteBuildBatchMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteBuildBatchOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteBuildBatchInput struct {

	// The identifier of the batch build to delete.
	//
	// This member is required.
	Id *string

	noSmithyDocumentSerde
}

type DeleteBuildBatchOutput struct {

	// An array of strings that contain the identifiers of the builds that were
	// deleted.
	BuildsDeleted []string

	// An array of BuildNotDeleted objects that specify the builds that could not be
	// deleted.
	BuildsNotDeleted []types.BuildNotDeleted

	// The status code.
	StatusCode *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteBuildBatchMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteBuildBatch{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteBuildBatch{}, middleware.After)
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
	if err = addOpDeleteBuildBatchValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteBuildBatch(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteBuildBatch(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codebuild",
		OperationName: "DeleteBuildBatch",
	}
}
