// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/backup/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns the framework details for the specified FrameworkName.
func (c *Client) DescribeFramework(ctx context.Context, params *DescribeFrameworkInput, optFns ...func(*Options)) (*DescribeFrameworkOutput, error) {
	if params == nil {
		params = &DescribeFrameworkInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeFramework", params, optFns, c.addOperationDescribeFrameworkMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeFrameworkOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeFrameworkInput struct {

	// The unique name of a framework.
	//
	// This member is required.
	FrameworkName *string

	noSmithyDocumentSerde
}

type DescribeFrameworkOutput struct {

	// The date and time that a framework is created, in ISO 8601 representation. The
	// value of CreationTime is accurate to milliseconds. For example,
	// 2020-07-10T15:00:00.000-08:00 represents the 10th of July 2020 at 3:00 PM 8
	// hours behind UTC.
	CreationTime *time.Time

	// The deployment status of a framework. The statuses are: CREATE_IN_PROGRESS |
	// UPDATE_IN_PROGRESS | DELETE_IN_PROGRESS | COMPLETED | FAILED
	DeploymentStatus *string

	// An Amazon Resource Name (ARN) that uniquely identifies a resource. The format of
	// the ARN depends on the resource type.
	FrameworkArn *string

	// A list of the controls that make up the framework. Each control in the list has
	// a name, input parameters, and scope.
	FrameworkControls []types.FrameworkControl

	// An optional description of the framework.
	FrameworkDescription *string

	// The unique name of a framework.
	FrameworkName *string

	// A framework consists of one or more controls. Each control governs a resource,
	// such as backup plans, backup selections, backup vaults, or recovery points. You
	// can also turn Config recording on or off for each resource. The statuses are:
	//
	// *
	// ACTIVE when recording is turned on for all resources governed by the
	// framework.
	//
	// * PARTIALLY_ACTIVE when recording is turned off for at least one
	// resource governed by the framework.
	//
	// * INACTIVE when recording is turned off for
	// all resources governed by the framework.
	//
	// * UNAVAILABLE when Backup is unable to
	// validate recording status at this time.
	FrameworkStatus *string

	// A customer-chosen string that you can use to distinguish between otherwise
	// identical calls to DescribeFrameworkOutput. Retrying a successful request with
	// the same idempotency token results in a success message with no action taken.
	IdempotencyToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeFrameworkMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeFramework{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeFramework{}, middleware.After)
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
	if err = addOpDescribeFrameworkValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeFramework(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeFramework(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "backup",
		OperationName: "DescribeFramework",
	}
}
