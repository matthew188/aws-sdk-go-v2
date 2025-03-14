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
// global use. Creates a SizeConstraintSet. You then use UpdateSizeConstraintSet to
// identify the part of a web request that you want AWS WAF to check for length,
// such as the length of the User-Agent header or the length of the query string.
// For example, you can create a SizeConstraintSet that matches any requests that
// have a query string that is longer than 100 bytes. You can then configure AWS
// WAF to reject those requests. To create and configure a SizeConstraintSet,
// perform the following steps:
//
// * Use GetChangeToken to get the change token that
// you provide in the ChangeToken parameter of a CreateSizeConstraintSet
// request.
//
// * Submit a CreateSizeConstraintSet request.
//
// * Use GetChangeToken to
// get the change token that you provide in the ChangeToken parameter of an
// UpdateSizeConstraintSet request.
//
// * Submit an UpdateSizeConstraintSet request to
// specify the part of the request that you want AWS WAF to inspect (for example,
// the header or the URI) and the value that you want AWS WAF to watch for.
//
// For
// more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the AWS WAF Developer Guide
// (https://docs.aws.amazon.com/waf/latest/developerguide/).
func (c *Client) CreateSizeConstraintSet(ctx context.Context, params *CreateSizeConstraintSetInput, optFns ...func(*Options)) (*CreateSizeConstraintSetOutput, error) {
	if params == nil {
		params = &CreateSizeConstraintSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSizeConstraintSet", params, optFns, c.addOperationCreateSizeConstraintSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSizeConstraintSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSizeConstraintSetInput struct {

	// The value returned by the most recent call to GetChangeToken.
	//
	// This member is required.
	ChangeToken *string

	// A friendly name or description of the SizeConstraintSet. You can't change Name
	// after you create a SizeConstraintSet.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

type CreateSizeConstraintSetOutput struct {

	// The ChangeToken that you used to submit the CreateSizeConstraintSet request. You
	// can also use this value to query the status of the request. For more
	// information, see GetChangeTokenStatus.
	ChangeToken *string

	// A SizeConstraintSet that contains no SizeConstraint objects.
	SizeConstraintSet *types.SizeConstraintSet

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateSizeConstraintSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateSizeConstraintSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateSizeConstraintSet{}, middleware.After)
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
	if err = addOpCreateSizeConstraintSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSizeConstraintSet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateSizeConstraintSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf-regional",
		OperationName: "CreateSizeConstraintSet",
	}
}
