// Code generated by smithy-go-codegen DO NOT EDIT.

package location

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves glyphs used to display labels on a map.
func (c *Client) GetMapGlyphs(ctx context.Context, params *GetMapGlyphsInput, optFns ...func(*Options)) (*GetMapGlyphsOutput, error) {
	if params == nil {
		params = &GetMapGlyphsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetMapGlyphs", params, optFns, c.addOperationGetMapGlyphsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetMapGlyphsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMapGlyphsInput struct {

	// A comma-separated list of fonts to load glyphs from in order of preference. For
	// example, Noto Sans Regular, Arial Unicode. Valid fonts stacks for Esri
	// (https://docs.aws.amazon.com/location/latest/developerguide/esri.html)
	// styles:
	//
	// * VectorEsriDarkGrayCanvas – Ubuntu Medium Italic | Ubuntu Medium |
	// Ubuntu Italic | Ubuntu Regular | Ubuntu Bold
	//
	// * VectorEsriLightGrayCanvas –
	// Ubuntu Italic | Ubuntu Regular | Ubuntu Light | Ubuntu Bold
	//
	// *
	// VectorEsriTopographic – Noto Sans Italic | Noto Sans Regular | Noto Sans Bold |
	// Noto Serif Regular | Roboto Condensed Light Italic
	//
	// * VectorEsriStreets – Arial
	// Regular | Arial Italic | Arial Bold
	//
	// * VectorEsriNavigation – Arial Regular |
	// Arial Italic | Arial Bold
	//
	// Valid font stacks for HERE Technologies
	// (https://docs.aws.amazon.com/location/latest/developerguide/HERE.html)
	// styles:
	//
	// * VectorHereContrast – Fira GO Regular | Fira GO Bold
	//
	// *
	// VectorHereExplore, VectorHereExploreTruck, HybridHereExploreSatellite – Fira GO
	// Italic | Fira GO Map | Fira GO Map Bold | Noto Sans CJK JP Bold | Noto Sans CJK
	// JP Light | Noto Sans CJK JP Regular
	//
	// Valid font stacks for GrabMaps
	// (https://docs.aws.amazon.com/location/latest/developerguide/grab.html)
	// styles:
	//
	// * VectorGrabStandardLight, VectorGrabStandardDark – Noto Sans Regular |
	// Noto Sans Medium | Noto Sans Bold
	//
	// Valid font stacks for Open Data
	// (https://docs.aws.amazon.com/location/latest/developerguide/open-data.html)
	// styles:
	//
	// * VectorOpenDataStandardLight, VectorOpenDataStandardDark,
	// VectorOpenDataVisualizationLight, VectorOpenDataVisualizationDark – Amazon Ember
	// Regular,Noto Sans Regular | Amazon Ember Bold,Noto Sans Bold | Amazon Ember
	// Medium,Noto Sans Medium | Amazon Ember Regular Italic,Noto Sans Italic | Amazon
	// Ember Condensed RC Regular,Noto Sans Regular | Amazon Ember Condensed RC
	// Bold,Noto Sans Bold
	//
	// The fonts used by the Open Data map styles are combined
	// fonts that use Amazon Ember for most glyphs but Noto Sans for glyphs unsupported
	// by Amazon Ember.
	//
	// This member is required.
	FontStack *string

	// A Unicode range of characters to download glyphs for. Each response will contain
	// 256 characters. For example, 0–255 includes all characters from range U+0000 to
	// 00FF. Must be aligned to multiples of 256.
	//
	// This member is required.
	FontUnicodeRange *string

	// The map resource associated with the glyph ﬁle.
	//
	// This member is required.
	MapName *string

	// The optional API key
	// (https://docs.aws.amazon.com/location/latest/developerguide/using-apikeys.html)
	// to authorize the request.
	Key *string

	noSmithyDocumentSerde
}

type GetMapGlyphsOutput struct {

	// The glyph, as binary blob.
	Blob []byte

	// The HTTP Cache-Control directive for the value.
	CacheControl *string

	// The map glyph content type. For example, application/octet-stream.
	ContentType *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetMapGlyphsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetMapGlyphs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetMapGlyphs{}, middleware.After)
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
	if err = addEndpointPrefix_opGetMapGlyphsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetMapGlyphsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetMapGlyphs(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opGetMapGlyphsMiddleware struct {
}

func (*endpointPrefix_opGetMapGlyphsMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opGetMapGlyphsMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "maps." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opGetMapGlyphsMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opGetMapGlyphsMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opGetMapGlyphs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "geo",
		OperationName: "GetMapGlyphs",
	}
}
