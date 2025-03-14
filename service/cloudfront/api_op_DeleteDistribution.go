// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Delete a distribution.
func (c *Client) DeleteDistribution(ctx context.Context, params *DeleteDistributionInput, optFns ...func(*Options)) (*DeleteDistributionOutput, error) {
	if params == nil {
		params = &DeleteDistributionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteDistribution", params, optFns, c.addOperationDeleteDistributionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteDistributionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// This action deletes a web distribution. To delete a web distribution using the
// CloudFront API, perform the following steps. To delete a web distribution using
// the CloudFront API:
//
// * Disable the web distribution
//
// * Submit a GET Distribution
// Config request to get the current configuration and the Etag header for the
// distribution.
//
// * Update the XML document that was returned in the response to
// your GET Distribution Config request to change the value of Enabled to false.
//
// *
// Submit a PUT Distribution Config request to update the configuration for your
// distribution. In the request body, include the XML document that you updated in
// Step 3. Set the value of the HTTP If-Match header to the value of the ETag
// header that CloudFront returned when you submitted the GET Distribution Config
// request in Step 2.
//
// * Review the response to the PUT Distribution Config request
// to confirm that the distribution was successfully disabled.
//
// * Submit a GET
// Distribution request to confirm that your changes have propagated. When
// propagation is complete, the value of Status is Deployed.
//
// * Submit a DELETE
// Distribution request. Set the value of the HTTP If-Match header to the value of
// the ETag header that CloudFront returned when you submitted the GET Distribution
// Config request in Step 6.
//
// * Review the response to your DELETE Distribution
// request to confirm that the distribution was successfully deleted.
//
// For
// information about deleting a distribution using the CloudFront console, see
// Deleting a Distribution
// (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/HowToDeleteDistribution.html)
// in the Amazon CloudFront Developer Guide.
type DeleteDistributionInput struct {

	// The distribution ID.
	//
	// This member is required.
	Id *string

	// The value of the ETag header that you received when you disabled the
	// distribution. For example: E2QWRUHAPOMQZL.
	IfMatch *string

	noSmithyDocumentSerde
}

type DeleteDistributionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteDistributionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpDeleteDistribution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpDeleteDistribution{}, middleware.After)
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
	if err = addOpDeleteDistributionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteDistribution(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteDistribution(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "DeleteDistribution",
	}
}
