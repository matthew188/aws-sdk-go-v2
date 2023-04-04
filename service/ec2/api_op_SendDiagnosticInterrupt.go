// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Sends a diagnostic interrupt to the specified Amazon EC2 instance to trigger a
// kernel panic (on Linux instances), or a blue screen/stop error (on Windows
// instances). For instances based on Intel and AMD processors, the interrupt is
// received as a non-maskable interrupt (NMI). In general, the operating system
// crashes and reboots when a kernel panic or stop error is triggered. The
// operating system can also be configured to perform diagnostic tasks, such as
// generating a memory dump file, loading a secondary kernel, or obtaining a call
// trace. Before sending a diagnostic interrupt to your instance, ensure that its
// operating system is configured to perform the required diagnostic tasks. For
// more information about configuring your operating system to generate a crash
// dump when a kernel panic or stop error occurs, see Send a diagnostic interrupt
// (for advanced users)
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/diagnostic-interrupt.html)
// (Linux instances) or Send a diagnostic interrupt (for advanced users)
// (https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/diagnostic-interrupt.html)
// (Windows instances).
func (c *Client) SendDiagnosticInterrupt(ctx context.Context, params *SendDiagnosticInterruptInput, optFns ...func(*Options)) (*SendDiagnosticInterruptOutput, error) {
	if params == nil {
		params = &SendDiagnosticInterruptInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SendDiagnosticInterrupt", params, optFns, c.addOperationSendDiagnosticInterruptMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SendDiagnosticInterruptOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SendDiagnosticInterruptInput struct {

	// The ID of the instance.
	//
	// This member is required.
	InstanceId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	noSmithyDocumentSerde
}

type SendDiagnosticInterruptOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSendDiagnosticInterruptMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpSendDiagnosticInterrupt{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpSendDiagnosticInterrupt{}, middleware.After)
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
	if err = addOpSendDiagnosticInterruptValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSendDiagnosticInterrupt(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSendDiagnosticInterrupt(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "SendDiagnosticInterrupt",
	}
}
