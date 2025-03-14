// Code generated by smithy-go-codegen DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/databasemigrationservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists categories for all event source types, or, if specified, for a specified
// source type. You can see a list of the event categories and source types in
// Working with Events and Notifications
// (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Events.html) in the
// Database Migration Service User Guide.
func (c *Client) DescribeEventCategories(ctx context.Context, params *DescribeEventCategoriesInput, optFns ...func(*Options)) (*DescribeEventCategoriesOutput, error) {
	if params == nil {
		params = &DescribeEventCategoriesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEventCategories", params, optFns, c.addOperationDescribeEventCategoriesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEventCategoriesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeEventCategoriesInput struct {

	// Filters applied to the event categories.
	Filters []types.Filter

	// The type of DMS resource that generates events. Valid values:
	// replication-instance | replication-task
	SourceType *string

	noSmithyDocumentSerde
}

type DescribeEventCategoriesOutput struct {

	// A list of event categories.
	EventCategoryGroupList []types.EventCategoryGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeEventCategoriesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeEventCategories{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeEventCategories{}, middleware.After)
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
	if err = addOpDescribeEventCategoriesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEventCategories(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeEventCategories(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dms",
		OperationName: "DescribeEventCategories",
	}
}
