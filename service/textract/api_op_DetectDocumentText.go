// Code generated by smithy-go-codegen DO NOT EDIT.

package textract

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/textract/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Detects text in the input document. Amazon Textract can detect lines of text and
// the words that make up a line of text. The input document must be in one of the
// following image formats: JPEG, PNG, PDF, or TIFF. DetectDocumentText returns the
// detected text in an array of Block objects. Each document page has as an
// associated Block of type PAGE. Each PAGE Block object is the parent of LINE
// Block objects that represent the lines of detected text on a page. A LINE Block
// object is a parent for each word that makes up the line. Words are represented
// by Block objects of type WORD. DetectDocumentText is a synchronous operation. To
// analyze documents asynchronously, use StartDocumentTextDetection. For more
// information, see Document Text Detection
// (https://docs.aws.amazon.com/textract/latest/dg/how-it-works-detecting.html).
func (c *Client) DetectDocumentText(ctx context.Context, params *DetectDocumentTextInput, optFns ...func(*Options)) (*DetectDocumentTextOutput, error) {
	if params == nil {
		params = &DetectDocumentTextInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DetectDocumentText", params, optFns, c.addOperationDetectDocumentTextMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DetectDocumentTextOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DetectDocumentTextInput struct {

	// The input document as base64-encoded bytes or an Amazon S3 object. If you use
	// the AWS CLI to call Amazon Textract operations, you can't pass image bytes. The
	// document must be an image in JPEG or PNG format. If you're using an AWS SDK to
	// call Amazon Textract, you might not need to base64-encode image bytes that are
	// passed using the Bytes field.
	//
	// This member is required.
	Document *types.Document

	noSmithyDocumentSerde
}

type DetectDocumentTextOutput struct {

	// An array of Block objects that contain the text that's detected in the document.
	Blocks []types.Block

	//
	DetectDocumentTextModelVersion *string

	// Metadata about the document. It contains the number of pages that are detected
	// in the document.
	DocumentMetadata *types.DocumentMetadata

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDetectDocumentTextMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDetectDocumentText{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDetectDocumentText{}, middleware.After)
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
	if err = addOpDetectDocumentTextValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDetectDocumentText(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDetectDocumentText(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "textract",
		OperationName: "DetectDocumentText",
	}
}
