// Code generated by smithy-go-codegen DO NOT EDIT.

package memorydb

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/memorydb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Apply the service update to a list of clusters supplied. For more information on
// service updates and applying them, see Applying the service updates
// (https://docs.aws.amazon.com/MemoryDB/latest/devguide/managing-updates.html#applying-updates).
func (c *Client) BatchUpdateCluster(ctx context.Context, params *BatchUpdateClusterInput, optFns ...func(*Options)) (*BatchUpdateClusterOutput, error) {
	if params == nil {
		params = &BatchUpdateClusterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchUpdateCluster", params, optFns, c.addOperationBatchUpdateClusterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchUpdateClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchUpdateClusterInput struct {

	// The cluster names to apply the updates.
	//
	// This member is required.
	ClusterNames []string

	// The unique ID of the service update
	ServiceUpdate *types.ServiceUpdateRequest

	noSmithyDocumentSerde
}

type BatchUpdateClusterOutput struct {

	// The list of clusters that have been updated.
	ProcessedClusters []types.Cluster

	// The list of clusters where updates have not been applied.
	UnprocessedClusters []types.UnprocessedCluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchUpdateClusterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpBatchUpdateCluster{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchUpdateCluster{}, middleware.After)
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
	if err = addOpBatchUpdateClusterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchUpdateCluster(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchUpdateCluster(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "memorydb",
		OperationName: "BatchUpdateCluster",
	}
}
