// Code generated by smithy-go-codegen DO NOT EDIT.

package wafregional

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/wafregional/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This is AWS WAF Classic documentation. For more information, see AWS WAF Classic
// (https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html)
// in the developer guide. For the latest version of AWS WAF, use the AWS WAFV2 API
// and see the AWS WAF Developer Guide
// (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html). With
// the latest version, AWS WAF has a single set of endpoints for regional and
// global use. Creates an GeoMatchSet, which you use to specify which web requests
// you want to allow or block based on the country that the requests originate
// from. For example, if you're receiving a lot of requests from one or more
// countries and you want to block the requests, you can create an GeoMatchSet that
// contains those countries and then configure AWS WAF to block the requests. To
// create and configure a GeoMatchSet, perform the following steps:
//
// * Use
// GetChangeToken to get the change token that you provide in the ChangeToken
// parameter of a CreateGeoMatchSet request.
//
// * Submit a CreateGeoMatchSet
// request.
//
// * Use GetChangeToken to get the change token that you provide in the
// ChangeToken parameter of an UpdateGeoMatchSet request.
//
// * Submit an
// UpdateGeoMatchSetSet request to specify the countries that you want AWS WAF to
// watch for.
//
// For more information about how to use the AWS WAF API to allow or
// block HTTP requests, see the AWS WAF Developer Guide
// (https://docs.aws.amazon.com/waf/latest/developerguide/).
func (c *Client) CreateGeoMatchSet(ctx context.Context, params *CreateGeoMatchSetInput, optFns ...func(*Options)) (*CreateGeoMatchSetOutput, error) {
	if params == nil {
		params = &CreateGeoMatchSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateGeoMatchSet", params, optFns, c.addOperationCreateGeoMatchSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateGeoMatchSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateGeoMatchSetInput struct {

	// The value returned by the most recent call to GetChangeToken.
	//
	// This member is required.
	ChangeToken *string

	// A friendly name or description of the GeoMatchSet. You can't change Name after
	// you create the GeoMatchSet.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

type CreateGeoMatchSetOutput struct {

	// The ChangeToken that you used to submit the CreateGeoMatchSet request. You can
	// also use this value to query the status of the request. For more information,
	// see GetChangeTokenStatus.
	ChangeToken *string

	// The GeoMatchSet returned in the CreateGeoMatchSet response. The GeoMatchSet
	// contains no GeoMatchConstraints.
	GeoMatchSet *types.GeoMatchSet

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateGeoMatchSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateGeoMatchSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateGeoMatchSet{}, middleware.After)
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
	if err = addOpCreateGeoMatchSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateGeoMatchSet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateGeoMatchSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf-regional",
		OperationName: "CreateGeoMatchSet",
	}
}
