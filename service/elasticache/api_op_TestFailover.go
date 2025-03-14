// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/elasticache/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Represents the input of a TestFailover operation which test automatic failover
// on a specified node group (called shard in the console) in a replication group
// (called cluster in the console). This API is designed for testing the behavior
// of your application in case of ElastiCache failover. It is not designed to be an
// operational tool for initiating a failover to overcome a problem you may have
// with the cluster. Moreover, in certain conditions such as large-scale
// operational events, Amazon may block this API. Note the following
//
// * A customer
// can use this operation to test automatic failover on up to 5 shards (called node
// groups in the ElastiCache API and Amazon CLI) in any rolling 24-hour period.
//
// *
// If calling this operation on shards in different clusters (called replication
// groups in the API and CLI), the calls can be made concurrently.
//
// * If calling
// this operation multiple times on different shards in the same Redis (cluster
// mode enabled) replication group, the first node replacement must complete before
// a subsequent call can be made.
//
// * To determine whether the node replacement is
// complete you can check Events using the Amazon ElastiCache console, the Amazon
// CLI, or the ElastiCache API. Look for the following automatic failover related
// events, listed here in order of occurrance:
//
// * Replication group message: Test
// Failover API called for node group
//
// * Cache cluster message: Failover from
// primary node to replica node completed
//
// * Replication group message: Failover
// from primary node to replica node completed
//
// * Cache cluster message: Recovering
// cache nodes
//
// * Cache cluster message: Finished recovery for cache nodes
//
// For
// more information see:
//
// * Viewing ElastiCache Events
// (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/ECEvents.Viewing.html)
// in the ElastiCache User Guide
//
// * DescribeEvents
// (https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_DescribeEvents.html)
// in the ElastiCache API Reference
//
// Also see, Testing Multi-AZ
// (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/AutoFailover.html#auto-failover-test)
// in the ElastiCache User Guide.
func (c *Client) TestFailover(ctx context.Context, params *TestFailoverInput, optFns ...func(*Options)) (*TestFailoverOutput, error) {
	if params == nil {
		params = &TestFailoverInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "TestFailover", params, optFns, c.addOperationTestFailoverMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*TestFailoverOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TestFailoverInput struct {

	// The name of the node group (called shard in the console) in this replication
	// group on which automatic failover is to be tested. You may test automatic
	// failover on up to 5 node groups in any rolling 24-hour period.
	//
	// This member is required.
	NodeGroupId *string

	// The name of the replication group (console: cluster) whose automatic failover is
	// being tested by this operation.
	//
	// This member is required.
	ReplicationGroupId *string

	noSmithyDocumentSerde
}

type TestFailoverOutput struct {

	// Contains all of the attributes of a specific Redis replication group.
	ReplicationGroup *types.ReplicationGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationTestFailoverMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpTestFailover{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpTestFailover{}, middleware.After)
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
	if err = addOpTestFailoverValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opTestFailover(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opTestFailover(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "TestFailover",
	}
}
