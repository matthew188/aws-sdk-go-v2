// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describe a thing group. Requires permission to access the DescribeThingGroup
// (https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions)
// action.
func (c *Client) DescribeThingGroup(ctx context.Context, params *DescribeThingGroupInput, optFns ...func(*Options)) (*DescribeThingGroupOutput, error) {
	if params == nil {
		params = &DescribeThingGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeThingGroup", params, optFns, c.addOperationDescribeThingGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeThingGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeThingGroupInput struct {

	// The name of the thing group.
	//
	// This member is required.
	ThingGroupName *string

	noSmithyDocumentSerde
}

type DescribeThingGroupOutput struct {

	// The dynamic thing group index name.
	IndexName *string

	// The dynamic thing group search query string.
	QueryString *string

	// The dynamic thing group query version.
	QueryVersion *string

	// The dynamic thing group status.
	Status types.DynamicGroupStatus

	// The thing group ARN.
	ThingGroupArn *string

	// The thing group ID.
	ThingGroupId *string

	// Thing group metadata.
	ThingGroupMetadata *types.ThingGroupMetadata

	// The name of the thing group.
	ThingGroupName *string

	// The thing group properties.
	ThingGroupProperties *types.ThingGroupProperties

	// The version of the thing group.
	Version int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeThingGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeThingGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeThingGroup{}, middleware.After)
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
	if err = addOpDescribeThingGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeThingGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeThingGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "DescribeThingGroup",
	}
}
