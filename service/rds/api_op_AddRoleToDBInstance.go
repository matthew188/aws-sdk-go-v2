// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Associates an Amazon Web Services Identity and Access Management (IAM) role with
// a DB instance. To add a role to a DB instance, the status of the DB instance
// must be available. This command doesn't apply to RDS Custom.
func (c *Client) AddRoleToDBInstance(ctx context.Context, params *AddRoleToDBInstanceInput, optFns ...func(*Options)) (*AddRoleToDBInstanceOutput, error) {
	if params == nil {
		params = &AddRoleToDBInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AddRoleToDBInstance", params, optFns, c.addOperationAddRoleToDBInstanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AddRoleToDBInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AddRoleToDBInstanceInput struct {

	// The name of the DB instance to associate the IAM role with.
	//
	// This member is required.
	DBInstanceIdentifier *string

	// The name of the feature for the DB instance that the IAM role is to be
	// associated with. For information about supported feature names, see
	// DBEngineVersion.
	//
	// This member is required.
	FeatureName *string

	// The Amazon Resource Name (ARN) of the IAM role to associate with the DB
	// instance, for example arn:aws:iam::123456789012:role/AccessRole.
	//
	// This member is required.
	RoleArn *string

	noSmithyDocumentSerde
}

type AddRoleToDBInstanceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAddRoleToDBInstanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpAddRoleToDBInstance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpAddRoleToDBInstance{}, middleware.After)
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
	if err = addOpAddRoleToDBInstanceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAddRoleToDBInstance(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAddRoleToDBInstance(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "AddRoleToDBInstance",
	}
}
