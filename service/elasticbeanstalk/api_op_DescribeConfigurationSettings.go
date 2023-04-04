// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticbeanstalk

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/elasticbeanstalk/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a description of the settings for the specified configuration set, that
// is, either a configuration template or the configuration set associated with a
// running environment. When describing the settings for the configuration set
// associated with a running environment, it is possible to receive two sets of
// setting descriptions. One is the deployed configuration set, and the other is a
// draft configuration of an environment that is either in the process of
// deployment or that failed to deploy. Related Topics
//
// *
// DeleteEnvironmentConfiguration
func (c *Client) DescribeConfigurationSettings(ctx context.Context, params *DescribeConfigurationSettingsInput, optFns ...func(*Options)) (*DescribeConfigurationSettingsOutput, error) {
	if params == nil {
		params = &DescribeConfigurationSettingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeConfigurationSettings", params, optFns, c.addOperationDescribeConfigurationSettingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeConfigurationSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Result message containing all of the configuration settings for a specified
// solution stack or configuration template.
type DescribeConfigurationSettingsInput struct {

	// The application for the environment or configuration template.
	//
	// This member is required.
	ApplicationName *string

	// The name of the environment to describe. Condition: You must specify either this
	// or a TemplateName, but not both. If you specify both, AWS Elastic Beanstalk
	// returns an InvalidParameterCombination error. If you do not specify either, AWS
	// Elastic Beanstalk returns MissingRequiredParameter error.
	EnvironmentName *string

	// The name of the configuration template to describe. Conditional: You must
	// specify either this parameter or an EnvironmentName, but not both. If you
	// specify both, AWS Elastic Beanstalk returns an InvalidParameterCombination
	// error. If you do not specify either, AWS Elastic Beanstalk returns a
	// MissingRequiredParameter error.
	TemplateName *string

	noSmithyDocumentSerde
}

// The results from a request to change the configuration settings of an
// environment.
type DescribeConfigurationSettingsOutput struct {

	// A list of ConfigurationSettingsDescription.
	ConfigurationSettings []types.ConfigurationSettingsDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeConfigurationSettingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeConfigurationSettings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeConfigurationSettings{}, middleware.After)
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
	if err = addOpDescribeConfigurationSettingsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeConfigurationSettings(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeConfigurationSettings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticbeanstalk",
		OperationName: "DescribeConfigurationSettings",
	}
}
