// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies the specified event window. You can define either a set of time ranges
// or a cron expression when modifying the event window, but not both. To modify
// the targets associated with the event window, use the
// AssociateInstanceEventWindow and DisassociateInstanceEventWindow API. If Amazon
// Web Services has already scheduled an event, modifying an event window won't
// change the time of the scheduled event. For more information, see Define event
// windows for scheduled events
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/event-windows.html) in the
// Amazon EC2 User Guide.
func (c *Client) ModifyInstanceEventWindow(ctx context.Context, params *ModifyInstanceEventWindowInput, optFns ...func(*Options)) (*ModifyInstanceEventWindowOutput, error) {
	if params == nil {
		params = &ModifyInstanceEventWindowInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyInstanceEventWindow", params, optFns, c.addOperationModifyInstanceEventWindowMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyInstanceEventWindowOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyInstanceEventWindowInput struct {

	// The ID of the event window.
	//
	// This member is required.
	InstanceEventWindowId *string

	// The cron expression of the event window, for example, * 0-4,20-23 * * 1,5.
	// Constraints:
	//
	// * Only hour and day of the week values are supported.
	//
	// * For day
	// of the week values, you can specify either integers 0 through 6, or alternative
	// single values SUN through SAT.
	//
	// * The minute, month, and year must be specified
	// by *.
	//
	// * The hour value must be one or a multiple range, for example, 0-4 or
	// 0-4,20-23.
	//
	// * Each hour range must be >= 2 hours, for example, 0-2 or 20-23.
	//
	// *
	// The event window must be >= 4 hours. The combined total time ranges in the event
	// window must be >= 4 hours.
	//
	// For more information about cron expressions, see
	// cron (https://en.wikipedia.org/wiki/Cron) on the Wikipedia website.
	CronExpression *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The name of the event window.
	Name *string

	// The time ranges of the event window.
	TimeRanges []types.InstanceEventWindowTimeRangeRequest

	noSmithyDocumentSerde
}

type ModifyInstanceEventWindowOutput struct {

	// Information about the event window.
	InstanceEventWindow *types.InstanceEventWindow

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyInstanceEventWindowMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpModifyInstanceEventWindow{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpModifyInstanceEventWindow{}, middleware.After)
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
	if err = addOpModifyInstanceEventWindowValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyInstanceEventWindow(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyInstanceEventWindow(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ModifyInstanceEventWindow",
	}
}
