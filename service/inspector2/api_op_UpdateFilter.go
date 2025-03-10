// Code generated by smithy-go-codegen DO NOT EDIT.

package inspector2

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/inspector2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Specifies the action that is to be applied to the findings that match the
// filter.
func (c *Client) UpdateFilter(ctx context.Context, params *UpdateFilterInput, optFns ...func(*Options)) (*UpdateFilterOutput, error) {
	if params == nil {
		params = &UpdateFilterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateFilter", params, optFns, c.addOperationUpdateFilterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateFilterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateFilterInput struct {

	// The Amazon Resource Number (ARN) of the filter to update.
	//
	// This member is required.
	FilterArn *string

	// Specifies the action that is to be applied to the findings that match the
	// filter.
	Action types.FilterAction

	// A description of the filter.
	Description *string

	// Defines the criteria to be update in the filter.
	FilterCriteria *types.FilterCriteria

	// The name of the filter.
	Name *string

	// The reason the filter was updated.
	Reason *string

	noSmithyDocumentSerde
}

type UpdateFilterOutput struct {

	// The Amazon Resource Number (ARN) of the successfully updated filter.
	//
	// This member is required.
	Arn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateFilterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateFilter{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateFilter{}, middleware.After)
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
	if err = addOpUpdateFilterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateFilter(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateFilter(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "inspector2",
		OperationName: "UpdateFilter",
	}
}
