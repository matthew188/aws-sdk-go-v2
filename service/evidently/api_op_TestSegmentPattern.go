// Code generated by smithy-go-codegen DO NOT EDIT.

package evidently

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Use this operation to test a rules pattern that you plan to use to create an
// audience segment. For more information about segments, see CreateSegment
// (https://docs.aws.amazon.com/cloudwatchevidently/latest/APIReference/API_CreateSegment.html).
func (c *Client) TestSegmentPattern(ctx context.Context, params *TestSegmentPatternInput, optFns ...func(*Options)) (*TestSegmentPatternOutput, error) {
	if params == nil {
		params = &TestSegmentPatternInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "TestSegmentPattern", params, optFns, c.addOperationTestSegmentPatternMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*TestSegmentPatternOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TestSegmentPatternInput struct {

	// The pattern to test.
	//
	// This value conforms to the media type: application/json
	//
	// This member is required.
	Pattern *string

	// A sample evaluationContext JSON block to test against the specified pattern.
	//
	// This value conforms to the media type: application/json
	//
	// This member is required.
	Payload *string

	noSmithyDocumentSerde
}

type TestSegmentPatternOutput struct {

	// Returns true if the pattern matches the payload.
	//
	// This member is required.
	Match *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationTestSegmentPatternMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpTestSegmentPattern{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpTestSegmentPattern{}, middleware.After)
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
	if err = addOpTestSegmentPatternValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opTestSegmentPattern(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opTestSegmentPattern(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "evidently",
		OperationName: "TestSegmentPattern",
	}
}
