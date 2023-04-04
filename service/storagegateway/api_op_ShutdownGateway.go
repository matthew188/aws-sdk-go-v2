// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Shuts down a gateway. To specify which gateway to shut down, use the Amazon
// Resource Name (ARN) of the gateway in the body of your request. The operation
// shuts down the gateway service component running in the gateway's virtual
// machine (VM) and not the host VM. If you want to shut down the VM, it is
// recommended that you first shut down the gateway component in the VM to avoid
// unpredictable conditions. After the gateway is shutdown, you cannot call any
// other API except StartGateway, DescribeGatewayInformation, and ListGateways. For
// more information, see ActivateGateway. Your applications cannot read from or
// write to the gateway's storage volumes, and there are no snapshots taken. When
// you make a shutdown request, you will get a 200 OK success response immediately.
// However, it might take some time for the gateway to shut down. You can call the
// DescribeGatewayInformation API to check the status. For more information, see
// ActivateGateway. If do not intend to use the gateway again, you must delete the
// gateway (using DeleteGateway) to no longer pay software charges associated with
// the gateway.
func (c *Client) ShutdownGateway(ctx context.Context, params *ShutdownGatewayInput, optFns ...func(*Options)) (*ShutdownGatewayOutput, error) {
	if params == nil {
		params = &ShutdownGatewayInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ShutdownGateway", params, optFns, c.addOperationShutdownGatewayMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ShutdownGatewayOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A JSON object containing the Amazon Resource Name (ARN) of the gateway to shut
// down.
type ShutdownGatewayInput struct {

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation to
	// return a list of gateways for your account and Amazon Web Services Region.
	//
	// This member is required.
	GatewayARN *string

	noSmithyDocumentSerde
}

// A JSON object containing the Amazon Resource Name (ARN) of the gateway that was
// shut down.
type ShutdownGatewayOutput struct {

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation to
	// return a list of gateways for your account and Amazon Web Services Region.
	GatewayARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationShutdownGatewayMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpShutdownGateway{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpShutdownGateway{}, middleware.After)
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
	if err = addOpShutdownGatewayValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opShutdownGateway(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opShutdownGateway(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "ShutdownGateway",
	}
}
