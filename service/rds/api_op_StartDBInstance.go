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

// Starts an Amazon RDS DB instance that was stopped using the Amazon Web Services
// console, the stop-db-instance CLI command, or the StopDBInstance action. For
// more information, see  Starting an Amazon RDS DB instance That Was Previously
// Stopped
// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_StartInstance.html)
// in the Amazon RDS User Guide. This command doesn't apply to RDS Custom, Aurora
// MySQL, and Aurora PostgreSQL. For Aurora DB clusters, use StartDBCluster
// instead.
func (c *Client) StartDBInstance(ctx context.Context, params *StartDBInstanceInput, optFns ...func(*Options)) (*StartDBInstanceOutput, error) {
	if params == nil {
		params = &StartDBInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartDBInstance", params, optFns, c.addOperationStartDBInstanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartDBInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartDBInstanceInput struct {

	// The user-supplied instance identifier.
	//
	// This member is required.
	DBInstanceIdentifier *string

	noSmithyDocumentSerde
}

type StartDBInstanceOutput struct {

	// Contains the details of an Amazon RDS DB instance. This data type is used as a
	// response element in the operations CreateDBInstance,
	// CreateDBInstanceReadReplica, DeleteDBInstance, DescribeDBInstances,
	// ModifyDBInstance, PromoteReadReplica, RebootDBInstance,
	// RestoreDBInstanceFromDBSnapshot, RestoreDBInstanceFromS3,
	// RestoreDBInstanceToPointInTime, StartDBInstance, and StopDBInstance.
	DBInstance *types.DBInstance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartDBInstanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpStartDBInstance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpStartDBInstance{}, middleware.After)
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
	if err = addOpStartDBInstanceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartDBInstance(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartDBInstance(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "StartDBInstance",
	}
}
