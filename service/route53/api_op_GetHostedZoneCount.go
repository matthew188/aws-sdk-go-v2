// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the number of hosted zones that are associated with the current Amazon
// Web Services account.
func (c *Client) GetHostedZoneCount(ctx context.Context, params *GetHostedZoneCountInput, optFns ...func(*Options)) (*GetHostedZoneCountOutput, error) {
	if params == nil {
		params = &GetHostedZoneCountInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetHostedZoneCount", params, optFns, c.addOperationGetHostedZoneCountMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetHostedZoneCountOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to retrieve a count of all the hosted zones that are associated with
// the current Amazon Web Services account.
type GetHostedZoneCountInput struct {
	noSmithyDocumentSerde
}

// A complex type that contains the response to a GetHostedZoneCount request.
type GetHostedZoneCountOutput struct {

	// The total number of public and private hosted zones that are associated with the
	// current Amazon Web Services account.
	//
	// This member is required.
	HostedZoneCount *int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetHostedZoneCountMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetHostedZoneCount{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetHostedZoneCount{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetHostedZoneCount(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetHostedZoneCount(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "GetHostedZoneCount",
	}
}
