// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Disassociates an Amazon Web Services Identity and Access Management (IAM) role
// from a DB instance.
func (c *Client) RemoveRoleFromDBInstance(ctx context.Context, params *RemoveRoleFromDBInstanceInput, optFns ...func(*Options)) (*RemoveRoleFromDBInstanceOutput, error) {
	if params == nil {
		params = &RemoveRoleFromDBInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RemoveRoleFromDBInstance", params, optFns, c.addOperationRemoveRoleFromDBInstanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RemoveRoleFromDBInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RemoveRoleFromDBInstanceInput struct {

	// The name of the DB instance to disassociate the IAM role from.
	//
	// This member is required.
	DBInstanceIdentifier *string

	// The name of the feature for the DB instance that the IAM role is to be
	// disassociated from. For information about supported feature names, see
	// DBEngineVersion.
	//
	// This member is required.
	FeatureName *string

	// The Amazon Resource Name (ARN) of the IAM role to disassociate from the DB
	// instance, for example, arn:aws:iam::123456789012:role/AccessRole.
	//
	// This member is required.
	RoleArn *string

	noSmithyDocumentSerde
}

type RemoveRoleFromDBInstanceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRemoveRoleFromDBInstanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpRemoveRoleFromDBInstance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpRemoveRoleFromDBInstance{}, middleware.After)
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
	if err = addOpRemoveRoleFromDBInstanceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRemoveRoleFromDBInstance(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRemoveRoleFromDBInstance(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "RemoveRoleFromDBInstance",
	}
}
