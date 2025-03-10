// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/apigateway/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the GatewayResponses collection on the given RestApi. If an API developer
// has not added any definitions for gateway responses, the result will be the API
// Gateway-generated default GatewayResponses collection for the supported response
// types.
func (c *Client) GetGatewayResponses(ctx context.Context, params *GetGatewayResponsesInput, optFns ...func(*Options)) (*GetGatewayResponsesOutput, error) {
	if params == nil {
		params = &GetGatewayResponsesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetGatewayResponses", params, optFns, c.addOperationGetGatewayResponsesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetGatewayResponsesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Gets the GatewayResponses collection on the given RestApi. If an API developer
// has not added any definitions for gateway responses, the result will be the API
// Gateway-generated default GatewayResponses collection for the supported response
// types.
type GetGatewayResponsesInput struct {

	// The string identifier of the associated RestApi.
	//
	// This member is required.
	RestApiId *string

	// The maximum number of returned results per page. The default value is 25 and the
	// maximum value is 500. The GatewayResponses collection does not support
	// pagination and the limit does not apply here.
	Limit *int32

	// The current pagination position in the paged result set. The GatewayResponse
	// collection does not support pagination and the position does not apply here.
	Position *string

	noSmithyDocumentSerde
}

// The collection of the GatewayResponse instances of a RestApi as a
// responseType-to-GatewayResponse object map of key-value pairs. As such,
// pagination is not supported for querying this collection.
type GetGatewayResponsesOutput struct {

	// Returns the entire collection, because of no pagination support.
	Items []types.GatewayResponse

	// The current pagination position in the paged result set. The GatewayResponse
	// collection does not support pagination and the position does not apply here.
	Position *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetGatewayResponsesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetGatewayResponses{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetGatewayResponses{}, middleware.After)
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
	if err = addOpGetGatewayResponsesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetGatewayResponses(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addAcceptHeader(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opGetGatewayResponses(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "GetGatewayResponses",
	}
}
