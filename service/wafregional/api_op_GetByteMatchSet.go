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
// global use. Returns the ByteMatchSet specified by ByteMatchSetId.
func (c *Client) GetByteMatchSet(ctx context.Context, params *GetByteMatchSetInput, optFns ...func(*Options)) (*GetByteMatchSetOutput, error) {
	if params == nil {
		params = &GetByteMatchSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetByteMatchSet", params, optFns, c.addOperationGetByteMatchSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetByteMatchSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetByteMatchSetInput struct {

	// The ByteMatchSetId of the ByteMatchSet that you want to get. ByteMatchSetId is
	// returned by CreateByteMatchSet and by ListByteMatchSets.
	//
	// This member is required.
	ByteMatchSetId *string

	noSmithyDocumentSerde
}

type GetByteMatchSetOutput struct {

	// Information about the ByteMatchSet that you specified in the GetByteMatchSet
	// request. For more information, see the following topics:
	//
	// * ByteMatchSet:
	// Contains ByteMatchSetId, ByteMatchTuples, and Name
	//
	// * ByteMatchTuples: Contains
	// an array of ByteMatchTuple objects. Each ByteMatchTuple object contains
	// FieldToMatch, PositionalConstraint, TargetString, and TextTransformation
	//
	// *
	// FieldToMatch: Contains Data and Type
	ByteMatchSet *types.ByteMatchSet

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetByteMatchSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetByteMatchSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetByteMatchSet{}, middleware.After)
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
	if err = addOpGetByteMatchSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetByteMatchSet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetByteMatchSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf-regional",
		OperationName: "GetByteMatchSet",
	}
}
