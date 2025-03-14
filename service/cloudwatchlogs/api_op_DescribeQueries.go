// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchlogs

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of CloudWatch Logs Insights queries that are scheduled, running,
// or have been run recently in this account. You can request all queries or limit
// it to queries of a specific log group or queries with a certain status.
func (c *Client) DescribeQueries(ctx context.Context, params *DescribeQueriesInput, optFns ...func(*Options)) (*DescribeQueriesOutput, error) {
	if params == nil {
		params = &DescribeQueriesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeQueries", params, optFns, c.addOperationDescribeQueriesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeQueriesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeQueriesInput struct {

	// Limits the returned queries to only those for the specified log group.
	LogGroupName *string

	// Limits the number of returned queries to the specified number.
	MaxResults *int32

	// The token for the next set of items to return. The token expires after 24 hours.
	NextToken *string

	// Limits the returned queries to only those that have the specified status. Valid
	// values are Cancelled, Complete, Failed, Running, and Scheduled.
	Status types.QueryStatus

	noSmithyDocumentSerde
}

type DescribeQueriesOutput struct {

	// The token for the next set of items to return. The token expires after 24 hours.
	NextToken *string

	// The list of queries that match the request.
	Queries []types.QueryInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeQueriesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeQueries{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeQueries{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeQueries(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeQueries(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "logs",
		OperationName: "DescribeQueries",
	}
}
