// Code generated by smithy-go-codegen DO NOT EDIT.

package wisdom

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/wisdom/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves summary information about the content.
func (c *Client) GetContentSummary(ctx context.Context, params *GetContentSummaryInput, optFns ...func(*Options)) (*GetContentSummaryOutput, error) {
	if params == nil {
		params = &GetContentSummaryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetContentSummary", params, optFns, c.addOperationGetContentSummaryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetContentSummaryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetContentSummaryInput struct {

	// The identifier of the content. Can be either the ID or the ARN. URLs cannot
	// contain the ARN.
	//
	// This member is required.
	ContentId *string

	// The identifier of the knowledge base. Can be either the ID or the ARN. URLs
	// cannot contain the ARN.
	//
	// This member is required.
	KnowledgeBaseId *string

	noSmithyDocumentSerde
}

type GetContentSummaryOutput struct {

	// The content summary.
	ContentSummary *types.ContentSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetContentSummaryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetContentSummary{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetContentSummary{}, middleware.After)
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
	if err = addOpGetContentSummaryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetContentSummary(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetContentSummary(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "wisdom",
		OperationName: "GetContentSummary",
	}
}
