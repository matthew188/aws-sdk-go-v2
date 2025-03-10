// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehend

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/comprehend/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Analyzes input text for the presence of personally identifiable information
// (PII) and returns the labels of identified PII entity types such as name,
// address, bank account number, or phone number.
func (c *Client) ContainsPiiEntities(ctx context.Context, params *ContainsPiiEntitiesInput, optFns ...func(*Options)) (*ContainsPiiEntitiesOutput, error) {
	if params == nil {
		params = &ContainsPiiEntitiesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ContainsPiiEntities", params, optFns, c.addOperationContainsPiiEntitiesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ContainsPiiEntitiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ContainsPiiEntitiesInput struct {

	// The language of the input documents. Currently, English is the only valid
	// language.
	//
	// This member is required.
	LanguageCode types.LanguageCode

	// A UTF-8 text string. The maximum string size is 100 KB.
	//
	// This member is required.
	Text *string

	noSmithyDocumentSerde
}

type ContainsPiiEntitiesOutput struct {

	// The labels used in the document being analyzed. Individual labels represent
	// personally identifiable information (PII) entity types.
	Labels []types.EntityLabel

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationContainsPiiEntitiesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpContainsPiiEntities{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpContainsPiiEntities{}, middleware.After)
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
	if err = addOpContainsPiiEntitiesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opContainsPiiEntities(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opContainsPiiEntities(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehend",
		OperationName: "ContainsPiiEntities",
	}
}
