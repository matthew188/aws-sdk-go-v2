// Code generated by smithy-go-codegen DO NOT EDIT.

package location

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/location/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Finds a place by its unique ID. A PlaceId is returned by other search
// operations. A PlaceId is valid only if all of the following are the same in the
// original search request and the call to GetPlace.
//
// * Customer Amazon Web
// Services account
//
// * Amazon Web Services Region
//
// * Data provider specified in the
// place index resource
func (c *Client) GetPlace(ctx context.Context, params *GetPlaceInput, optFns ...func(*Options)) (*GetPlaceOutput, error) {
	if params == nil {
		params = &GetPlaceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetPlace", params, optFns, c.addOperationGetPlaceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetPlaceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetPlaceInput struct {

	// The name of the place index resource that you want to use for the search.
	//
	// This member is required.
	IndexName *string

	// The identifier of the place to find.
	//
	// This member is required.
	PlaceId *string

	// The preferred language used to return results. The value must be a valid BCP 47
	// (https://tools.ietf.org/search/bcp47) language tag, for example, en for English.
	// This setting affects the languages used in the results, but not the results
	// themselves. If no language is specified, or not supported for a particular
	// result, the partner automatically chooses a language for the result. For an
	// example, we'll use the Greek language. You search for a location around Athens,
	// Greece, with the language parameter set to en. The city in the results will most
	// likely be returned as Athens. If you set the language parameter to el, for
	// Greek, then the city in the results will more likely be returned as Αθήνα. If
	// the data provider does not have a value for Greek, the result will be in a
	// language that the provider does support.
	Language *string

	noSmithyDocumentSerde
}

type GetPlaceOutput struct {

	// Details about the result, such as its address and position.
	//
	// This member is required.
	Place *types.Place

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetPlaceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetPlace{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetPlace{}, middleware.After)
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
	if err = addEndpointPrefix_opGetPlaceMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetPlaceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetPlace(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opGetPlaceMiddleware struct {
}

func (*endpointPrefix_opGetPlaceMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opGetPlaceMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "places." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opGetPlaceMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opGetPlaceMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opGetPlace(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "geo",
		OperationName: "GetPlace",
	}
}
