// Code generated by smithy-go-codegen DO NOT EDIT.

package vpclattice

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/vpclattice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves information about the specified service network.
func (c *Client) GetServiceNetwork(ctx context.Context, params *GetServiceNetworkInput, optFns ...func(*Options)) (*GetServiceNetworkOutput, error) {
	if params == nil {
		params = &GetServiceNetworkInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetServiceNetwork", params, optFns, c.addOperationGetServiceNetworkMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetServiceNetworkOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetServiceNetworkInput struct {

	// The ID or Amazon Resource Name (ARN) of the service network.
	//
	// This member is required.
	ServiceNetworkIdentifier *string

	noSmithyDocumentSerde
}

type GetServiceNetworkOutput struct {

	// The Amazon Resource Name (ARN) of the service network.
	Arn *string

	// The type of IAM policy.
	AuthType types.AuthType

	// The date and time that the service network was created, specified in ISO-8601
	// format.
	CreatedAt *time.Time

	// The ID of the service network.
	Id *string

	// The date and time of the last update, specified in ISO-8601 format.
	LastUpdatedAt *time.Time

	// The name of the service network.
	Name *string

	// The number of services associated with the service network.
	NumberOfAssociatedServices *int64

	// The number of VPCs associated with the service network.
	NumberOfAssociatedVPCs *int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetServiceNetworkMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetServiceNetwork{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetServiceNetwork{}, middleware.After)
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
	if err = addOpGetServiceNetworkValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetServiceNetwork(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetServiceNetwork(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "vpc-lattice",
		OperationName: "GetServiceNetwork",
	}
}
