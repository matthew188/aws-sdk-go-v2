// Code generated by smithy-go-codegen DO NOT EDIT.

package macie2

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/macie2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the sensitivity scoring settings for an S3 bucket.
func (c *Client) UpdateResourceProfileDetections(ctx context.Context, params *UpdateResourceProfileDetectionsInput, optFns ...func(*Options)) (*UpdateResourceProfileDetectionsOutput, error) {
	if params == nil {
		params = &UpdateResourceProfileDetectionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateResourceProfileDetections", params, optFns, c.addOperationUpdateResourceProfileDetectionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateResourceProfileDetectionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateResourceProfileDetectionsInput struct {

	// The Amazon Resource Name (ARN) of the S3 bucket that the request applies to.
	//
	// This member is required.
	ResourceArn *string

	// An array of objects, one for each custom data identifier or managed data
	// identifier that detected the type of sensitive data to start excluding or
	// including in the bucket's score. To start including all sensitive data types in
	// the score, don't specify any values for this array.
	SuppressDataIdentifiers []types.SuppressDataIdentifier

	noSmithyDocumentSerde
}

type UpdateResourceProfileDetectionsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateResourceProfileDetectionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateResourceProfileDetections{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateResourceProfileDetections{}, middleware.After)
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
	if err = addOpUpdateResourceProfileDetectionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateResourceProfileDetections(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateResourceProfileDetections(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "macie2",
		OperationName: "UpdateResourceProfileDetections",
	}
}
