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

// Forces a failover for a DB cluster. For an Aurora DB cluster, failover for a DB
// cluster promotes one of the Aurora Replicas (read-only instances) in the DB
// cluster to be the primary DB instance (the cluster writer). For a Multi-AZ DB
// cluster, failover for a DB cluster promotes one of the readable standby DB
// instances (read-only instances) in the DB cluster to be the primary DB instance
// (the cluster writer). An Amazon Aurora DB cluster automatically fails over to an
// Aurora Replica, if one exists, when the primary DB instance fails. A Multi-AZ DB
// cluster automatically fails over to a readable standby DB instance when the
// primary DB instance fails. To simulate a failure of a primary instance for
// testing, you can force a failover. Because each instance in a DB cluster has its
// own endpoint address, make sure to clean up and re-establish any existing
// connections that use those endpoint addresses when the failover is complete. For
// more information on Amazon Aurora DB clusters, see  What is Amazon Aurora?
// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html)
// in the Amazon Aurora User Guide. For more information on Multi-AZ DB clusters,
// see  Multi-AZ DB cluster deployments
// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html)
// in the Amazon RDS User Guide.
func (c *Client) FailoverDBCluster(ctx context.Context, params *FailoverDBClusterInput, optFns ...func(*Options)) (*FailoverDBClusterOutput, error) {
	if params == nil {
		params = &FailoverDBClusterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "FailoverDBCluster", params, optFns, c.addOperationFailoverDBClusterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*FailoverDBClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type FailoverDBClusterInput struct {

	// A DB cluster identifier to force a failover for. This parameter isn't
	// case-sensitive. Constraints:
	//
	// * Must match the identifier of an existing
	// DBCluster.
	//
	// This member is required.
	DBClusterIdentifier *string

	// The name of the DB instance to promote to the primary DB instance. Specify the
	// DB instance identifier for an Aurora Replica or a Multi-AZ readable standby in
	// the DB cluster, for example mydbcluster-replica1. This setting isn't supported
	// for RDS for MySQL Multi-AZ DB clusters.
	TargetDBInstanceIdentifier *string

	noSmithyDocumentSerde
}

type FailoverDBClusterOutput struct {

	// Contains the details of an Amazon Aurora DB cluster or Multi-AZ DB cluster. For
	// an Amazon Aurora DB cluster, this data type is used as a response element in the
	// operations CreateDBCluster, DeleteDBCluster, DescribeDBClusters,
	// FailoverDBCluster, ModifyDBCluster, PromoteReadReplicaDBCluster,
	// RestoreDBClusterFromS3, RestoreDBClusterFromSnapshot,
	// RestoreDBClusterToPointInTime, StartDBCluster, and StopDBCluster. For a Multi-AZ
	// DB cluster, this data type is used as a response element in the operations
	// CreateDBCluster, DeleteDBCluster, DescribeDBClusters, FailoverDBCluster,
	// ModifyDBCluster, RebootDBCluster, RestoreDBClusterFromSnapshot, and
	// RestoreDBClusterToPointInTime. For more information on Amazon Aurora DB
	// clusters, see  What is Amazon Aurora?
	// (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/CHAP_AuroraOverview.html)
	// in the Amazon Aurora User Guide. For more information on Multi-AZ DB clusters,
	// see  Multi-AZ deployments with two readable standby DB instances
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/multi-az-db-clusters-concepts.html)
	// in the Amazon RDS User Guide.
	DBCluster *types.DBCluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationFailoverDBClusterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpFailoverDBCluster{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpFailoverDBCluster{}, middleware.After)
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
	if err = addOpFailoverDBClusterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opFailoverDBCluster(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opFailoverDBCluster(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "FailoverDBCluster",
	}
}
