// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Cancels an instance refresh or rollback that is in progress. If an instance
// refresh or rollback is not in progress, an ActiveInstanceRefreshNotFound error
// occurs. This operation is part of the instance refresh feature
// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-instance-refresh.html)
// in Amazon EC2 Auto Scaling, which helps you update instances in your Auto
// Scaling group after you make configuration changes. When you cancel an instance
// refresh, this does not roll back any changes that it made. Use the
// RollbackInstanceRefresh API to roll back instead.
func (c *Client) CancelInstanceRefresh(ctx context.Context, params *CancelInstanceRefreshInput, optFns ...func(*Options)) (*CancelInstanceRefreshOutput, error) {
	if params == nil {
		params = &CancelInstanceRefreshInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CancelInstanceRefresh", params, optFns, c.addOperationCancelInstanceRefreshMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CancelInstanceRefreshOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CancelInstanceRefreshInput struct {

	// The name of the Auto Scaling group.
	//
	// This member is required.
	AutoScalingGroupName *string

	noSmithyDocumentSerde
}

type CancelInstanceRefreshOutput struct {

	// The instance refresh ID associated with the request. This is the unique ID
	// assigned to the instance refresh when it was started.
	InstanceRefreshId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCancelInstanceRefreshMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCancelInstanceRefresh{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCancelInstanceRefresh{}, middleware.After)
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
	if err = addOpCancelInstanceRefreshValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCancelInstanceRefresh(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCancelInstanceRefresh(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "CancelInstanceRefresh",
	}
}
