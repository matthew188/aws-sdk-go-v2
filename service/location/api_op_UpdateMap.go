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
	"time"
)

// Updates the specified properties of a given map resource.
func (c *Client) UpdateMap(ctx context.Context, params *UpdateMapInput, optFns ...func(*Options)) (*UpdateMapOutput, error) {
	if params == nil {
		params = &UpdateMapInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateMap", params, optFns, c.addOperationUpdateMapMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateMapOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateMapInput struct {

	// The name of the map resource to update.
	//
	// This member is required.
	MapName *string

	// Updates the description for the map resource.
	Description *string

	// No longer used. If included, the only allowed value is RequestBasedUsage.
	//
	// Deprecated: Deprecated. If included, the only allowed value is
	// RequestBasedUsage.
	PricingPlan types.PricingPlan

	noSmithyDocumentSerde
}

type UpdateMapOutput struct {

	// The Amazon Resource Name (ARN) of the updated map resource. Used to specify a
	// resource across AWS.
	//
	// * Format example:
	// arn:aws:geo:region:account-id:map/ExampleMap
	//
	// This member is required.
	MapArn *string

	// The name of the updated map resource.
	//
	// This member is required.
	MapName *string

	// The timestamp for when the map resource was last updated in  ISO 8601
	// (https://www.iso.org/iso-8601-date-and-time-format.html) format:
	// YYYY-MM-DDThh:mm:ss.sssZ.
	//
	// This member is required.
	UpdateTime *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateMapMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateMap{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateMap{}, middleware.After)
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
	if err = addEndpointPrefix_opUpdateMapMiddleware(stack); err != nil {
		return err
	}
	if err = addOpUpdateMapValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateMap(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opUpdateMapMiddleware struct {
}

func (*endpointPrefix_opUpdateMapMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opUpdateMapMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
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
func addEndpointPrefix_opUpdateMapMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opUpdateMapMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateMap(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "geo",
		OperationName: "UpdateMap",
	}
}
