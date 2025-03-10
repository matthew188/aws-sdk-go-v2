// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehend

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a previously created document classifier Only those classifiers that are
// in terminated states (IN_ERROR, TRAINED) will be deleted. If an active inference
// job is using the model, a ResourceInUseException will be returned. This is an
// asynchronous action that puts the classifier into a DELETING state, and it is
// then removed by a background job. Once removed, the classifier disappears from
// your account and is no longer available for use.
func (c *Client) DeleteDocumentClassifier(ctx context.Context, params *DeleteDocumentClassifierInput, optFns ...func(*Options)) (*DeleteDocumentClassifierOutput, error) {
	if params == nil {
		params = &DeleteDocumentClassifierInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteDocumentClassifier", params, optFns, c.addOperationDeleteDocumentClassifierMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteDocumentClassifierOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteDocumentClassifierInput struct {

	// The Amazon Resource Name (ARN) that identifies the document classifier.
	//
	// This member is required.
	DocumentClassifierArn *string

	noSmithyDocumentSerde
}

type DeleteDocumentClassifierOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteDocumentClassifierMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteDocumentClassifier{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteDocumentClassifier{}, middleware.After)
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
	if err = addOpDeleteDocumentClassifierValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteDocumentClassifier(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteDocumentClassifier(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehend",
		OperationName: "DeleteDocumentClassifier",
	}
}
