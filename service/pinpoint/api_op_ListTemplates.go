// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpoint

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/pinpoint/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves information about all the message templates that are associated with
// your Amazon Pinpoint account.
func (c *Client) ListTemplates(ctx context.Context, params *ListTemplatesInput, optFns ...func(*Options)) (*ListTemplatesOutput, error) {
	if params == nil {
		params = &ListTemplatesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTemplates", params, optFns, c.addOperationListTemplatesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTemplatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTemplatesInput struct {

	// The string that specifies which page of results to return in a paginated
	// response. This parameter is not supported for application, campaign, and journey
	// metrics.
	NextToken *string

	// The maximum number of items to include in each page of a paginated response.
	// This parameter is not supported for application, campaign, and journey metrics.
	PageSize *string

	// The substring to match in the names of the message templates to include in the
	// results. If you specify this value, Amazon Pinpoint returns only those templates
	// whose names begin with the value that you specify.
	Prefix *string

	// The type of message template to include in the results. Valid values are: EMAIL,
	// PUSH, SMS, and VOICE. To include all types of templates in the results, don't
	// include this parameter in your request.
	TemplateType *string

	noSmithyDocumentSerde
}

type ListTemplatesOutput struct {

	// Provides information about all the message templates that are associated with
	// your Amazon Pinpoint account.
	//
	// This member is required.
	TemplatesResponse *types.TemplatesResponse

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListTemplatesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListTemplates{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListTemplates{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTemplates(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListTemplates(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mobiletargeting",
		OperationName: "ListTemplates",
	}
}
