// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisanalyticsv2

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/kinesisanalyticsv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Reverts the application to the previous running version. You can roll back an
// application if you suspect it is stuck in a transient status. You can roll back
// an application only if it is in the UPDATING or AUTOSCALING status. When you
// rollback an application, it loads state data from the last successful snapshot.
// If the application has no snapshots, Kinesis Data Analytics rejects the rollback
// request. This action is not supported for Kinesis Data Analytics for SQL
// applications.
func (c *Client) RollbackApplication(ctx context.Context, params *RollbackApplicationInput, optFns ...func(*Options)) (*RollbackApplicationOutput, error) {
	if params == nil {
		params = &RollbackApplicationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RollbackApplication", params, optFns, c.addOperationRollbackApplicationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RollbackApplicationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RollbackApplicationInput struct {

	// The name of the application.
	//
	// This member is required.
	ApplicationName *string

	// The current application version ID. You can retrieve the application version ID
	// using DescribeApplication.
	//
	// This member is required.
	CurrentApplicationVersionId *int64

	noSmithyDocumentSerde
}

type RollbackApplicationOutput struct {

	// Describes the application, including the application Amazon Resource Name (ARN),
	// status, latest version, and input and output configurations.
	//
	// This member is required.
	ApplicationDetail *types.ApplicationDetail

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRollbackApplicationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRollbackApplication{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRollbackApplication{}, middleware.After)
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
	if err = addOpRollbackApplicationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRollbackApplication(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRollbackApplication(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesisanalytics",
		OperationName: "RollbackApplication",
	}
}
