// Code generated by smithy-go-codegen DO NOT EDIT.

package wisdom

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the knowledge base. When you use this API to delete an external
// knowledge base such as Salesforce or ServiceNow, you must also delete the Amazon
// AppIntegrations
// (https://docs.aws.amazon.com/appintegrations/latest/APIReference/Welcome.html)
// DataIntegration. This is because you can't reuse the DataIntegration after it's
// been associated with an external knowledge base. However, you can delete and
// recreate it. See DeleteDataIntegration
// (https://docs.aws.amazon.com/appintegrations/latest/APIReference/API_DeleteDataIntegration.html)
// and CreateDataIntegration
// (https://docs.aws.amazon.com/appintegrations/latest/APIReference/API_CreateDataIntegration.html)
// in the Amazon AppIntegrations API Reference.
func (c *Client) DeleteKnowledgeBase(ctx context.Context, params *DeleteKnowledgeBaseInput, optFns ...func(*Options)) (*DeleteKnowledgeBaseOutput, error) {
	if params == nil {
		params = &DeleteKnowledgeBaseInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteKnowledgeBase", params, optFns, c.addOperationDeleteKnowledgeBaseMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteKnowledgeBaseOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteKnowledgeBaseInput struct {

	// The knowledge base to delete content from. Can be either the ID or the ARN. URLs
	// cannot contain the ARN.
	//
	// This member is required.
	KnowledgeBaseId *string

	noSmithyDocumentSerde
}

type DeleteKnowledgeBaseOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteKnowledgeBaseMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteKnowledgeBase{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteKnowledgeBase{}, middleware.After)
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
	if err = addOpDeleteKnowledgeBaseValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteKnowledgeBase(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteKnowledgeBase(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "wisdom",
		OperationName: "DeleteKnowledgeBase",
	}
}
