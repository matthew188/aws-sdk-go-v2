// Code generated by smithy-go-codegen DO NOT EDIT.

package pricing

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This feature is in preview release and is subject to change. Your use of Amazon
// Web Services Price List API is subject to the Beta Service Participation terms
// of the Amazon Web Services Service Terms (https://aws.amazon.com/service-terms/)
// (Section 1.10). This returns the URL that you can retrieve your Price List file
// from. This URL is based on the PriceListArn and FileFormat that you retrieve
// from the ListPriceLists
// (https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_pricing_ListPriceLists.html)
// response.
func (c *Client) GetPriceListFileUrl(ctx context.Context, params *GetPriceListFileUrlInput, optFns ...func(*Options)) (*GetPriceListFileUrlOutput, error) {
	if params == nil {
		params = &GetPriceListFileUrlInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetPriceListFileUrl", params, optFns, c.addOperationGetPriceListFileUrlMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetPriceListFileUrlOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetPriceListFileUrlInput struct {

	// The format that you want to retrieve your Price List files in. The FileFormat
	// can be obtained from the ListPriceLists
	// (https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_pricing_ListPriceLists.html)
	// response.
	//
	// This member is required.
	FileFormat *string

	// The unique identifier that maps to where your Price List files are located.
	// PriceListArn can be obtained from the ListPriceLists
	// (https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_pricing_ListPriceLists.html)
	// response.
	//
	// This member is required.
	PriceListArn *string

	noSmithyDocumentSerde
}

type GetPriceListFileUrlOutput struct {

	// The URL to download your Price List file from.
	Url *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetPriceListFileUrlMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetPriceListFileUrl{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetPriceListFileUrl{}, middleware.After)
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
	if err = addOpGetPriceListFileUrlValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetPriceListFileUrl(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetPriceListFileUrl(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "pricing",
		OperationName: "GetPriceListFileUrl",
	}
}
