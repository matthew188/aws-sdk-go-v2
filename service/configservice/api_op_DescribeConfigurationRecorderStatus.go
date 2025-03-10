// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/configservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the current status of the specified configuration recorder as well as
// the status of the last recording event for the recorder. If a configuration
// recorder is not specified, this action returns the status of all configuration
// recorders associated with the account. Currently, you can specify only one
// configuration recorder per region in your account. For a detailed status of
// recording events over time, add your Config events to Amazon CloudWatch metrics
// and use CloudWatch metrics.
func (c *Client) DescribeConfigurationRecorderStatus(ctx context.Context, params *DescribeConfigurationRecorderStatusInput, optFns ...func(*Options)) (*DescribeConfigurationRecorderStatusOutput, error) {
	if params == nil {
		params = &DescribeConfigurationRecorderStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeConfigurationRecorderStatus", params, optFns, c.addOperationDescribeConfigurationRecorderStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeConfigurationRecorderStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the DescribeConfigurationRecorderStatus action.
type DescribeConfigurationRecorderStatusInput struct {

	// The name(s) of the configuration recorder. If the name is not specified, the
	// action returns the current status of all the configuration recorders associated
	// with the account.
	ConfigurationRecorderNames []string

	noSmithyDocumentSerde
}

// The output for the DescribeConfigurationRecorderStatus action, in JSON format.
type DescribeConfigurationRecorderStatusOutput struct {

	// A list that contains status of the specified recorders.
	ConfigurationRecordersStatus []types.ConfigurationRecorderStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeConfigurationRecorderStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeConfigurationRecorderStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeConfigurationRecorderStatus{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeConfigurationRecorderStatus(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeConfigurationRecorderStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "DescribeConfigurationRecorderStatus",
	}
}
