// Code generated by smithy-go-codegen DO NOT EDIT.

package directconnect

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/directconnect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the attachments between your Direct Connect gateways and virtual
// interfaces. You must specify a Direct Connect gateway, a virtual interface, or
// both. If you specify a Direct Connect gateway, the response contains all virtual
// interfaces attached to the Direct Connect gateway. If you specify a virtual
// interface, the response contains all Direct Connect gateways attached to the
// virtual interface. If you specify both, the response contains the attachment
// between the Direct Connect gateway and the virtual interface.
func (c *Client) DescribeDirectConnectGatewayAttachments(ctx context.Context, params *DescribeDirectConnectGatewayAttachmentsInput, optFns ...func(*Options)) (*DescribeDirectConnectGatewayAttachmentsOutput, error) {
	if params == nil {
		params = &DescribeDirectConnectGatewayAttachmentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDirectConnectGatewayAttachments", params, optFns, c.addOperationDescribeDirectConnectGatewayAttachmentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDirectConnectGatewayAttachmentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDirectConnectGatewayAttachmentsInput struct {

	// The ID of the Direct Connect gateway.
	DirectConnectGatewayId *string

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value. If
	// MaxResults is given a value larger than 100, only 100 results are returned.
	MaxResults *int32

	// The token provided in the previous call to retrieve the next page.
	NextToken *string

	// The ID of the virtual interface.
	VirtualInterfaceId *string

	noSmithyDocumentSerde
}

type DescribeDirectConnectGatewayAttachmentsOutput struct {

	// The attachments.
	DirectConnectGatewayAttachments []types.DirectConnectGatewayAttachment

	// The token to retrieve the next page.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeDirectConnectGatewayAttachmentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeDirectConnectGatewayAttachments{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeDirectConnectGatewayAttachments{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDirectConnectGatewayAttachments(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeDirectConnectGatewayAttachments(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "directconnect",
		OperationName: "DescribeDirectConnectGatewayAttachments",
	}
}
