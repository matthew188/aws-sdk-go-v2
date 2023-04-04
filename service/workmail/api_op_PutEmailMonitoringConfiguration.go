// Code generated by smithy-go-codegen DO NOT EDIT.

package workmail

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates or updates the email monitoring configuration for a specified
// organization.
func (c *Client) PutEmailMonitoringConfiguration(ctx context.Context, params *PutEmailMonitoringConfigurationInput, optFns ...func(*Options)) (*PutEmailMonitoringConfigurationOutput, error) {
	if params == nil {
		params = &PutEmailMonitoringConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutEmailMonitoringConfiguration", params, optFns, c.addOperationPutEmailMonitoringConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutEmailMonitoringConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutEmailMonitoringConfigurationInput struct {

	// The Amazon Resource Name (ARN) of the CloudWatch Log group associated with the
	// email monitoring configuration.
	//
	// This member is required.
	LogGroupArn *string

	// The ID of the organization for which the email monitoring configuration is set.
	//
	// This member is required.
	OrganizationId *string

	// The Amazon Resource Name (ARN) of the IAM Role associated with the email
	// monitoring configuration.
	//
	// This member is required.
	RoleArn *string

	noSmithyDocumentSerde
}

type PutEmailMonitoringConfigurationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutEmailMonitoringConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutEmailMonitoringConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutEmailMonitoringConfiguration{}, middleware.After)
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
	if err = addOpPutEmailMonitoringConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutEmailMonitoringConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutEmailMonitoringConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workmail",
		OperationName: "PutEmailMonitoringConfiguration",
	}
}
