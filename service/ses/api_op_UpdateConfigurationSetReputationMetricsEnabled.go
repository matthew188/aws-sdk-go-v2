// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Enables or disables the publishing of reputation metrics for emails sent using a
// specific configuration set in a given AWS Region. Reputation metrics include
// bounce and complaint rates. These metrics are published to Amazon CloudWatch. By
// using CloudWatch, you can create alarms when bounce or complaint rates exceed
// certain thresholds. You can execute this operation no more than once per second.
func (c *Client) UpdateConfigurationSetReputationMetricsEnabled(ctx context.Context, params *UpdateConfigurationSetReputationMetricsEnabledInput, optFns ...func(*Options)) (*UpdateConfigurationSetReputationMetricsEnabledOutput, error) {
	if params == nil {
		params = &UpdateConfigurationSetReputationMetricsEnabledInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateConfigurationSetReputationMetricsEnabled", params, optFns, c.addOperationUpdateConfigurationSetReputationMetricsEnabledMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateConfigurationSetReputationMetricsEnabledOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to modify the reputation metric publishing settings for a
// configuration set.
type UpdateConfigurationSetReputationMetricsEnabledInput struct {

	// The name of the configuration set that you want to update.
	//
	// This member is required.
	ConfigurationSetName *string

	// Describes whether or not Amazon SES will publish reputation metrics for the
	// configuration set, such as bounce and complaint rates, to Amazon CloudWatch.
	//
	// This member is required.
	Enabled bool

	noSmithyDocumentSerde
}

type UpdateConfigurationSetReputationMetricsEnabledOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateConfigurationSetReputationMetricsEnabledMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpUpdateConfigurationSetReputationMetricsEnabled{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpUpdateConfigurationSetReputationMetricsEnabled{}, middleware.After)
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
	if err = addOpUpdateConfigurationSetReputationMetricsEnabledValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateConfigurationSetReputationMetricsEnabled(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateConfigurationSetReputationMetricsEnabled(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "UpdateConfigurationSetReputationMetricsEnabled",
	}
}
