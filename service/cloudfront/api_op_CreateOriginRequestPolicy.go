// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an origin request policy. After you create an origin request policy, you
// can attach it to one or more cache behaviors. When it's attached to a cache
// behavior, the origin request policy determines the values that CloudFront
// includes in requests that it sends to the origin. Each request that CloudFront
// sends to the origin includes the following:
//
// * The request body and the URL path
// (without the domain name) from the viewer request.
//
// * The headers that
// CloudFront automatically includes in every origin request, including Host,
// User-Agent, and X-Amz-Cf-Id.
//
// * All HTTP headers, cookies, and URL query strings
// that are specified in the cache policy or the origin request policy. These can
// include items from the viewer request and, in the case of headers, additional
// ones that are added by CloudFront.
//
// CloudFront sends a request when it can't
// find a valid object in its cache that matches the request. If you want to send
// values to the origin and also include them in the cache key, use CachePolicy.
// For more information about origin request policies, see Controlling origin
// requests
// (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-origin-requests.html)
// in the Amazon CloudFront Developer Guide.
func (c *Client) CreateOriginRequestPolicy(ctx context.Context, params *CreateOriginRequestPolicyInput, optFns ...func(*Options)) (*CreateOriginRequestPolicyOutput, error) {
	if params == nil {
		params = &CreateOriginRequestPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateOriginRequestPolicy", params, optFns, c.addOperationCreateOriginRequestPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateOriginRequestPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateOriginRequestPolicyInput struct {

	// An origin request policy configuration.
	//
	// This member is required.
	OriginRequestPolicyConfig *types.OriginRequestPolicyConfig

	noSmithyDocumentSerde
}

type CreateOriginRequestPolicyOutput struct {

	// The current version of the origin request policy.
	ETag *string

	// The fully qualified URI of the origin request policy just created.
	Location *string

	// An origin request policy.
	OriginRequestPolicy *types.OriginRequestPolicy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateOriginRequestPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpCreateOriginRequestPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpCreateOriginRequestPolicy{}, middleware.After)
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
	if err = addOpCreateOriginRequestPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateOriginRequestPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateOriginRequestPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "CreateOriginRequestPolicy",
	}
}
