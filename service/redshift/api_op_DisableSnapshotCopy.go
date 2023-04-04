// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/redshift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Disables the automatic copying of snapshots from one region to another region
// for a specified cluster. If your cluster and its snapshots are encrypted using
// an encrypted symmetric key from Key Management Service, use
// DeleteSnapshotCopyGrant to delete the grant that grants Amazon Redshift
// permission to the key in the destination region.
func (c *Client) DisableSnapshotCopy(ctx context.Context, params *DisableSnapshotCopyInput, optFns ...func(*Options)) (*DisableSnapshotCopyOutput, error) {
	if params == nil {
		params = &DisableSnapshotCopyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisableSnapshotCopy", params, optFns, c.addOperationDisableSnapshotCopyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisableSnapshotCopyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisableSnapshotCopyInput struct {

	// The unique identifier of the source cluster that you want to disable copying of
	// snapshots to a destination region. Constraints: Must be the valid name of an
	// existing cluster that has cross-region snapshot copy enabled.
	//
	// This member is required.
	ClusterIdentifier *string

	noSmithyDocumentSerde
}

type DisableSnapshotCopyOutput struct {

	// Describes a cluster.
	Cluster *types.Cluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDisableSnapshotCopyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDisableSnapshotCopy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDisableSnapshotCopy{}, middleware.After)
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
	if err = addOpDisableSnapshotCopyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisableSnapshotCopy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisableSnapshotCopy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "DisableSnapshotCopy",
	}
}
