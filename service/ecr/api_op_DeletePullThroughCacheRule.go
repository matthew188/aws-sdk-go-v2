// Code generated by smithy-go-codegen DO NOT EDIT.

package ecr

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Deletes a pull through cache rule.
func (c *Client) DeletePullThroughCacheRule(ctx context.Context, params *DeletePullThroughCacheRuleInput, optFns ...func(*Options)) (*DeletePullThroughCacheRuleOutput, error) {
	if params == nil {
		params = &DeletePullThroughCacheRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeletePullThroughCacheRule", params, optFns, c.addOperationDeletePullThroughCacheRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeletePullThroughCacheRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeletePullThroughCacheRuleInput struct {

	// The Amazon ECR repository prefix associated with the pull through cache rule to
	// delete.
	//
	// This member is required.
	EcrRepositoryPrefix *string

	// The Amazon Web Services account ID associated with the registry that contains
	// the pull through cache rule. If you do not specify a registry, the default
	// registry is assumed.
	RegistryId *string

	noSmithyDocumentSerde
}

type DeletePullThroughCacheRuleOutput struct {

	// The timestamp associated with the pull through cache rule.
	CreatedAt *time.Time

	// The Amazon ECR repository prefix associated with the request.
	EcrRepositoryPrefix *string

	// The registry ID associated with the request.
	RegistryId *string

	// The upstream registry URL associated with the pull through cache rule.
	UpstreamRegistryUrl *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeletePullThroughCacheRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeletePullThroughCacheRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeletePullThroughCacheRule{}, middleware.After)
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
	if err = addOpDeletePullThroughCacheRuleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeletePullThroughCacheRule(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeletePullThroughCacheRule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecr",
		OperationName: "DeletePullThroughCacheRule",
	}
}
