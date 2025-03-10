// Code generated by smithy-go-codegen DO NOT EDIT.

package customerprofiles

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/customerprofiles/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all of the template information for object types.
func (c *Client) ListProfileObjectTypeTemplates(ctx context.Context, params *ListProfileObjectTypeTemplatesInput, optFns ...func(*Options)) (*ListProfileObjectTypeTemplatesOutput, error) {
	if params == nil {
		params = &ListProfileObjectTypeTemplatesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListProfileObjectTypeTemplates", params, optFns, c.addOperationListProfileObjectTypeTemplatesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListProfileObjectTypeTemplatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListProfileObjectTypeTemplatesInput struct {

	// The maximum number of objects returned per page.
	MaxResults *int32

	// The pagination token from the previous ListObjectTypeTemplates API call.
	NextToken *string

	noSmithyDocumentSerde
}

type ListProfileObjectTypeTemplatesOutput struct {

	// The list of ListProfileObjectType template instances.
	Items []types.ListProfileObjectTypeTemplateItem

	// The pagination token from the previous ListObjectTypeTemplates API call.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListProfileObjectTypeTemplatesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListProfileObjectTypeTemplates{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListProfileObjectTypeTemplates{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListProfileObjectTypeTemplates(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListProfileObjectTypeTemplates(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "profile",
		OperationName: "ListProfileObjectTypeTemplates",
	}
}
