// Code generated by smithy-go-codegen DO NOT EDIT.

package neptune

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/neptune/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a global database. The primary and all secondary clusters must already
// be detached or deleted first.
func (c *Client) DeleteGlobalCluster(ctx context.Context, params *DeleteGlobalClusterInput, optFns ...func(*Options)) (*DeleteGlobalClusterOutput, error) {
	if params == nil {
		params = &DeleteGlobalClusterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteGlobalCluster", params, optFns, c.addOperationDeleteGlobalClusterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteGlobalClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteGlobalClusterInput struct {

	// The cluster identifier of the global database cluster being deleted.
	//
	// This member is required.
	GlobalClusterIdentifier *string

	noSmithyDocumentSerde
}

type DeleteGlobalClusterOutput struct {

	// Contains the details of an Amazon Neptune global database. This data type is
	// used as a response element for the CreateGlobalCluster, DescribeGlobalClusters,
	// ModifyGlobalCluster, DeleteGlobalCluster, FailoverGlobalCluster, and
	// RemoveFromGlobalCluster actions.
	GlobalCluster *types.GlobalCluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteGlobalClusterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeleteGlobalCluster{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteGlobalCluster{}, middleware.After)
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
	if err = addOpDeleteGlobalClusterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteGlobalCluster(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteGlobalCluster(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DeleteGlobalCluster",
	}
}
