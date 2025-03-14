// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the instance protection settings of the specified instances. This
// operation cannot be called on instances in a warm pool. For more information
// about preventing instances that are part of an Auto Scaling group from
// terminating on scale in, see Using instance scale-in protection
// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-instance-protection.html)
// in the Amazon EC2 Auto Scaling User Guide. If you exceed your maximum limit of
// instance IDs, which is 50 per Auto Scaling group, the call fails.
func (c *Client) SetInstanceProtection(ctx context.Context, params *SetInstanceProtectionInput, optFns ...func(*Options)) (*SetInstanceProtectionOutput, error) {
	if params == nil {
		params = &SetInstanceProtectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetInstanceProtection", params, optFns, c.addOperationSetInstanceProtectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SetInstanceProtectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SetInstanceProtectionInput struct {

	// The name of the Auto Scaling group.
	//
	// This member is required.
	AutoScalingGroupName *string

	// One or more instance IDs. You can specify up to 50 instances.
	//
	// This member is required.
	InstanceIds []string

	// Indicates whether the instance is protected from termination by Amazon EC2 Auto
	// Scaling when scaling in.
	//
	// This member is required.
	ProtectedFromScaleIn *bool

	noSmithyDocumentSerde
}

type SetInstanceProtectionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSetInstanceProtectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpSetInstanceProtection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpSetInstanceProtection{}, middleware.After)
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
	if err = addOpSetInstanceProtectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSetInstanceProtection(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSetInstanceProtection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "SetInstanceProtection",
	}
}
