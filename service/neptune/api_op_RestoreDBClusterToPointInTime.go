// Code generated by smithy-go-codegen DO NOT EDIT.

package neptune

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/neptune/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Restores a DB cluster to an arbitrary point in time. Users can restore to any
// point in time before LatestRestorableTime for up to BackupRetentionPeriod days.
// The target DB cluster is created from the source DB cluster with the same
// configuration as the original DB cluster, except that the new DB cluster is
// created with the default DB security group. This action only restores the DB
// cluster, not the DB instances for that DB cluster. You must invoke the
// CreateDBInstance action to create DB instances for the restored DB cluster,
// specifying the identifier of the restored DB cluster in DBClusterIdentifier. You
// can create DB instances only after the RestoreDBClusterToPointInTime action has
// completed and the DB cluster is available.
func (c *Client) RestoreDBClusterToPointInTime(ctx context.Context, params *RestoreDBClusterToPointInTimeInput, optFns ...func(*Options)) (*RestoreDBClusterToPointInTimeOutput, error) {
	if params == nil {
		params = &RestoreDBClusterToPointInTimeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RestoreDBClusterToPointInTime", params, optFns, c.addOperationRestoreDBClusterToPointInTimeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RestoreDBClusterToPointInTimeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RestoreDBClusterToPointInTimeInput struct {

	// The name of the new DB cluster to be created. Constraints:
	//
	// * Must contain from
	// 1 to 63 letters, numbers, or hyphens
	//
	// * First character must be a letter
	//
	// *
	// Cannot end with a hyphen or contain two consecutive hyphens
	//
	// This member is required.
	DBClusterIdentifier *string

	// The identifier of the source DB cluster from which to restore. Constraints:
	//
	// *
	// Must match the identifier of an existing DBCluster.
	//
	// This member is required.
	SourceDBClusterIdentifier *string

	// The name of the DB cluster parameter group to associate with the new DB cluster.
	// Constraints:
	//
	// * If supplied, must match the name of an existing
	// DBClusterParameterGroup.
	DBClusterParameterGroupName *string

	// The DB subnet group name to use for the new DB cluster. Constraints: If
	// supplied, must match the name of an existing DBSubnetGroup. Example:
	// mySubnetgroup
	DBSubnetGroupName *string

	// A value that indicates whether the DB cluster has deletion protection enabled.
	// The database can't be deleted when deletion protection is enabled. By default,
	// deletion protection is disabled.
	DeletionProtection *bool

	// The list of logs that the restored DB cluster is to export to CloudWatch Logs.
	EnableCloudwatchLogsExports []string

	// True to enable mapping of Amazon Identity and Access Management (IAM) accounts
	// to database accounts, and otherwise false. Default: false
	EnableIAMDatabaseAuthentication *bool

	// The Amazon KMS key identifier to use when restoring an encrypted DB cluster from
	// an encrypted DB cluster. The KMS key identifier is the Amazon Resource Name
	// (ARN) for the KMS encryption key. If you are restoring a DB cluster with the
	// same Amazon account that owns the KMS encryption key used to encrypt the new DB
	// cluster, then you can use the KMS key alias instead of the ARN for the KMS
	// encryption key. You can restore to a new DB cluster and encrypt the new DB
	// cluster with a KMS key that is different than the KMS key used to encrypt the
	// source DB cluster. The new DB cluster is encrypted with the KMS key identified
	// by the KmsKeyId parameter. If you do not specify a value for the KmsKeyId
	// parameter, then the following will occur:
	//
	// * If the DB cluster is encrypted,
	// then the restored DB cluster is encrypted using the KMS key that was used to
	// encrypt the source DB cluster.
	//
	// * If the DB cluster is not encrypted, then the
	// restored DB cluster is not encrypted.
	//
	// If DBClusterIdentifier refers to a DB
	// cluster that is not encrypted, then the restore request is rejected.
	KmsKeyId *string

	// (Not supported by Neptune)
	OptionGroupName *string

	// The port number on which the new DB cluster accepts connections. Constraints:
	// Value must be 1150-65535 Default: The same port as the original DB cluster.
	Port *int32

	// The date and time to restore the DB cluster to. Valid Values: Value must be a
	// time in Universal Coordinated Time (UTC) format Constraints:
	//
	// * Must be before
	// the latest restorable time for the DB instance
	//
	// * Must be specified if
	// UseLatestRestorableTime parameter is not provided
	//
	// * Cannot be specified if
	// UseLatestRestorableTime parameter is true
	//
	// * Cannot be specified if RestoreType
	// parameter is copy-on-write
	//
	// Example: 2015-03-07T23:45:00Z
	RestoreToTime *time.Time

	// The type of restore to be performed. You can specify one of the following
	// values:
	//
	// * full-copy - The new DB cluster is restored as a full copy of the
	// source DB cluster.
	//
	// * copy-on-write - The new DB cluster is restored as a clone
	// of the source DB cluster.
	//
	// If you don't specify a RestoreType value, then the
	// new DB cluster is restored as a full copy of the source DB cluster.
	RestoreType *string

	// Contains the scaling configuration of a Neptune Serverless DB cluster. For more
	// information, see Using Amazon Neptune Serverless
	// (https://docs.aws.amazon.com/neptune/latest/userguide/neptune-serverless-using.html)
	// in the Amazon Neptune User Guide.
	ServerlessV2ScalingConfiguration *types.ServerlessV2ScalingConfiguration

	// The tags to be applied to the restored DB cluster.
	Tags []types.Tag

	// A value that is set to true to restore the DB cluster to the latest restorable
	// backup time, and false otherwise. Default: false Constraints: Cannot be
	// specified if RestoreToTime parameter is provided.
	UseLatestRestorableTime bool

	// A list of VPC security groups that the new DB cluster belongs to.
	VpcSecurityGroupIds []string

	noSmithyDocumentSerde
}

type RestoreDBClusterToPointInTimeOutput struct {

	// Contains the details of an Amazon Neptune DB cluster. This data type is used as
	// a response element in the DescribeDBClusters action.
	DBCluster *types.DBCluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRestoreDBClusterToPointInTimeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpRestoreDBClusterToPointInTime{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpRestoreDBClusterToPointInTime{}, middleware.After)
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
	if err = addOpRestoreDBClusterToPointInTimeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRestoreDBClusterToPointInTime(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRestoreDBClusterToPointInTime(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "RestoreDBClusterToPointInTime",
	}
}
