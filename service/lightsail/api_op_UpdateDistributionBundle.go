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

// Updates the bundle of your Amazon Lightsail content delivery network (CDN)
// distribution. A distribution bundle specifies the monthly network transfer quota
// and monthly cost of your distribution. Update your distribution's bundle if your
// distribution is going over its monthly network transfer quota and is incurring
// an overage fee. You can update your distribution's bundle only one time within
// your monthly Amazon Web Services billing cycle. To determine if you can update
// your distribution's bundle, use the GetDistributions action. The
// ableToUpdateBundle parameter in the result will indicate whether you can
// currently update your distribution's bundle.
func (c *Client) UpdateDistributionBundle(ctx context.Context, params *UpdateDistributionBundleInput, optFns ...func(*Options)) (*UpdateDistributionBundleOutput, error) {
	if params == nil {
		params = &UpdateDistributionBundleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDistributionBundle", params, optFns, c.addOperationUpdateDistributionBundleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDistributionBundleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDistributionBundleInput struct {

	// The bundle ID of the new bundle to apply to your distribution. Use the
	// GetDistributionBundles action to get a list of distribution bundle IDs that you
	// can specify.
	BundleId *string

	// The name of the distribution for which to update the bundle. Use the
	// GetDistributions action to get a list of distribution names that you can
	// specify.
	DistributionName *string

	noSmithyDocumentSerde
}

type UpdateDistributionBundleOutput struct {

	// An object that describes the result of the action, such as the status of the
	// request, the timestamp of the request, and the resources affected by the
	// request.
	Operation *types.Operation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateDistributionBundleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateDistributionBundle{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateDistributionBundle{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDistributionBundle(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateDistributionBundle(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "UpdateDistributionBundle",
	}
}
