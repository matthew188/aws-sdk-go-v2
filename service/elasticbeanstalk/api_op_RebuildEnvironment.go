// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticbeanstalk

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes and recreates all of the AWS resources (for example: the Auto Scaling
// group, load balancer, etc.) for a specified environment and forces a restart.
func (c *Client) RebuildEnvironment(ctx context.Context, params *RebuildEnvironmentInput, optFns ...func(*Options)) (*RebuildEnvironmentOutput, error) {
	if params == nil {
		params = &RebuildEnvironmentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RebuildEnvironment", params, optFns, c.addOperationRebuildEnvironmentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RebuildEnvironmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RebuildEnvironmentInput struct {

	// The ID of the environment to rebuild. Condition: You must specify either this or
	// an EnvironmentName, or both. If you do not specify either, AWS Elastic Beanstalk
	// returns MissingRequiredParameter error.
	EnvironmentId *string

	// The name of the environment to rebuild. Condition: You must specify either this
	// or an EnvironmentId, or both. If you do not specify either, AWS Elastic
	// Beanstalk returns MissingRequiredParameter error.
	EnvironmentName *string

	noSmithyDocumentSerde
}

type RebuildEnvironmentOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRebuildEnvironmentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpRebuildEnvironment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpRebuildEnvironment{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRebuildEnvironment(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRebuildEnvironment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticbeanstalk",
		OperationName: "RebuildEnvironment",
	}
}
