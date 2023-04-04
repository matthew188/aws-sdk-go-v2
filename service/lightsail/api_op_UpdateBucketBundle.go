// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/lightsail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the bundle, or storage plan, of an existing Amazon Lightsail bucket. A
// bucket bundle specifies the monthly cost, storage space, and data transfer quota
// for a bucket. You can update a bucket's bundle only one time within a monthly
// Amazon Web Services billing cycle. To determine if you can update a bucket's
// bundle, use the GetBuckets
// (https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_GetBuckets.html)
// action. The ableToUpdateBundle parameter in the response will indicate whether
// you can currently update a bucket's bundle. Update a bucket's bundle if it's
// consistently going over its storage space or data transfer quota, or if a
// bucket's usage is consistently in the lower range of its storage space or data
// transfer quota. Due to the unpredictable usage fluctuations that a bucket might
// experience, we strongly recommend that you update a bucket's bundle only as a
// long-term strategy, instead of as a short-term, monthly cost-cutting measure.
// Choose a bucket bundle that will provide the bucket with ample storage space and
// data transfer for a long time to come.
func (c *Client) UpdateBucketBundle(ctx context.Context, params *UpdateBucketBundleInput, optFns ...func(*Options)) (*UpdateBucketBundleOutput, error) {
	if params == nil {
		params = &UpdateBucketBundleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateBucketBundle", params, optFns, c.addOperationUpdateBucketBundleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateBucketBundleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateBucketBundleInput struct {

	// The name of the bucket for which to update the bundle.
	//
	// This member is required.
	BucketName *string

	// The ID of the new bundle to apply to the bucket. Use the GetBucketBundles
	// (https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_GetBucketBundles.html)
	// action to get a list of bundle IDs that you can specify.
	//
	// This member is required.
	BundleId *string

	noSmithyDocumentSerde
}

type UpdateBucketBundleOutput struct {

	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected by the
	// request.
	Operations []types.Operation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateBucketBundleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateBucketBundle{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateBucketBundle{}, middleware.After)
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
	if err = addOpUpdateBucketBundleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateBucketBundle(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateBucketBundle(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "UpdateBucketBundle",
	}
}
