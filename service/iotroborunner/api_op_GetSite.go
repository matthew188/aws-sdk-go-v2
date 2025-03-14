// Code generated by smithy-go-codegen DO NOT EDIT.

package iotroborunner

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Grants permission to get a site
func (c *Client) GetSite(ctx context.Context, params *GetSiteInput, optFns ...func(*Options)) (*GetSiteOutput, error) {
	if params == nil {
		params = &GetSiteInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSite", params, optFns, c.addOperationGetSiteMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSiteOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSiteInput struct {

	// Site ARN.
	//
	// This member is required.
	Id *string

	noSmithyDocumentSerde
}

type GetSiteOutput struct {

	// Site ARN.
	//
	// This member is required.
	Arn *string

	// A valid ISO 3166-1 alpha-2 code for the country in which the site resides. e.g.,
	// US.
	//
	// This member is required.
	CountryCode *string

	// Timestamp at which the resource was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// Filters access by the site's identifier
	//
	// This member is required.
	Id *string

	// Human friendly name of the resource.
	//
	// This member is required.
	Name *string

	// Timestamp at which the resource was last updated.
	//
	// This member is required.
	UpdatedAt *time.Time

	// A high-level description of the site.
	Description *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetSiteMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetSite{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetSite{}, middleware.After)
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
	if err = addOpGetSiteValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSite(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetSite(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotroborunner",
		OperationName: "GetSite",
	}
}
