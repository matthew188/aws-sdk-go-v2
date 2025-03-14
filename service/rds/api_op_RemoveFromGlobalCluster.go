// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Detaches an Aurora secondary cluster from an Aurora global database cluster. The
// cluster becomes a standalone cluster with read-write capability instead of being
// read-only and receiving data from a primary cluster in a different Region. This
// action only applies to Aurora DB clusters.
func (c *Client) RemoveFromGlobalCluster(ctx context.Context, params *RemoveFromGlobalClusterInput, optFns ...func(*Options)) (*RemoveFromGlobalClusterOutput, error) {
	if params == nil {
		params = &RemoveFromGlobalClusterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RemoveFromGlobalCluster", params, optFns, c.addOperationRemoveFromGlobalClusterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RemoveFromGlobalClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RemoveFromGlobalClusterInput struct {

	// The Amazon Resource Name (ARN) identifying the cluster that was detached from
	// the Aurora global database cluster.
	DbClusterIdentifier *string

	// The cluster identifier to detach from the Aurora global database cluster.
	GlobalClusterIdentifier *string

	noSmithyDocumentSerde
}

type RemoveFromGlobalClusterOutput struct {

	// A data type representing an Aurora global database.
	GlobalCluster *types.GlobalCluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRemoveFromGlobalClusterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpRemoveFromGlobalCluster{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpRemoveFromGlobalCluster{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRemoveFromGlobalCluster(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRemoveFromGlobalCluster(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "RemoveFromGlobalCluster",
	}
}
