// Code generated by smithy-go-codegen DO NOT EDIT.

package appstream

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/appstream/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a list that describes one or more application fleet associations.
// Either ApplicationArn or FleetName must be specified.
func (c *Client) DescribeApplicationFleetAssociations(ctx context.Context, params *DescribeApplicationFleetAssociationsInput, optFns ...func(*Options)) (*DescribeApplicationFleetAssociationsOutput, error) {
	if params == nil {
		params = &DescribeApplicationFleetAssociationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeApplicationFleetAssociations", params, optFns, c.addOperationDescribeApplicationFleetAssociationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeApplicationFleetAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeApplicationFleetAssociationsInput struct {

	// The ARN of the application.
	ApplicationArn *string

	// The name of the fleet.
	FleetName *string

	// The maximum size of each page of results.
	MaxResults *int32

	// The pagination token used to retrieve the next page of results for this
	// operation.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeApplicationFleetAssociationsOutput struct {

	// The application fleet associations in the list.
	ApplicationFleetAssociations []types.ApplicationFleetAssociation

	// The pagination token used to retrieve the next page of results for this
	// operation.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeApplicationFleetAssociationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeApplicationFleetAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeApplicationFleetAssociations{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeApplicationFleetAssociations(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeApplicationFleetAssociations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appstream",
		OperationName: "DescribeApplicationFleetAssociations",
	}
}
