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

// Analyzes an input document for relationships between detected items. The types
// of information returned are as follows:
//
// * Form data (key-value pairs). The
// related information is returned in two Block objects, each of type
// KEY_VALUE_SET: a KEY Block object and a VALUE Block object. For example, Name:
// Ana Silva Carolina contains a key and value. Name: is the key. Ana Silva
// Carolina is the value.
//
// * Table and table cell data. A TABLE Block object
// contains information about a detected table. A CELL Block object is returned for
// each cell in a table.
//
// * Lines and words of text. A LINE Block object contains
// one or more WORD Block objects. All lines and words that are detected in the
// document are returned (including text that doesn't have a relationship with the
// value of FeatureTypes).
//
// * Signatures. A SIGNATURE Block object contains the
// location information of a signature in a document. If used in conjunction with
// forms or tables, a signature can be given a Key-Value pairing or be detected in
// the cell of a table.
//
// * Query. A QUERY Block object contains the query text,
// alias and link to the associated Query results block object.
//
// * Query Result. A
// QUERY_RESULT Block object contains the answer to the query and an ID that
// connects it to the query asked. This Block also contains a confidence
// score.
//
// Selection elements such as check boxes and option buttons (radio
// buttons) can be detected in form data and in tables. A SELECTION_ELEMENT Block
// object contains information about a selection element, including the selection
// status. You can choose which type of analysis to perform by specifying the
// FeatureTypes list. The output is returned in a list of Block objects.
// AnalyzeDocument is a synchronous operation. To analyze documents asynchronously,
// use StartDocumentAnalysis. For more information, see Document Text Analysis
// (https://docs.aws.amazon.com/textract/latest/dg/how-it-works-analyzing.html).
func (c *Client) AnalyzeDocument(ctx context.Context, params *AnalyzeDocumentInput, optFns ...func(*Options)) (*AnalyzeDocumentOutput, error) {
	if params == nil {
		params = &AnalyzeDocumentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AnalyzeDocument", params, optFns, c.addOperationAnalyzeDocumentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AnalyzeDocumentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AnalyzeDocumentInput struct {

	// The input document as base64-encoded bytes or an Amazon S3 object. If you use
	// the AWS CLI to call Amazon Textract operations, you can't pass image bytes. The
	// document must be an image in JPEG, PNG, PDF, or TIFF format. If you're using an
	// AWS SDK to call Amazon Textract, you might not need to base64-encode image bytes
	// that are passed using the Bytes field.
	//
	// This member is required.
	Document *types.Document

	// A list of the types of analysis to perform. Add TABLES to the list to return
	// information about the tables that are detected in the input document. Add FORMS
	// to return detected form data. Add SIGNATURES to return the locations of detected
	// signatures. To perform both forms and table analysis, add TABLES and FORMS to
	// FeatureTypes. To detect signatures within form data and table data, add
	// SIGNATURES to either TABLES or FORMS. All lines and words detected in the
	// document are included in the response (including text that isn't related to the
	// value of FeatureTypes).
	//
	// This member is required.
	FeatureTypes []types.FeatureType

	// Sets the configuration for the human in the loop workflow for analyzing
	// documents.
	HumanLoopConfig *types.HumanLoopConfig

	// Contains Queries and the alias for those Queries, as determined by the input.
	QueriesConfig *types.QueriesConfig

	noSmithyDocumentSerde
}

type AnalyzeDocumentOutput struct {

	// The version of the model used to analyze the document.
	AnalyzeDocumentModelVersion *string

	// The items that are detected and analyzed by AnalyzeDocument.
	Blocks []types.Block

	// Metadata about the analyzed document. An example is the number of pages.
	DocumentMetadata *types.DocumentMetadata

	// Shows the results of the human in the loop evaluation.
	HumanLoopActivationOutput *types.HumanLoopActivationOutput

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAnalyzeDocumentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAnalyzeDocument{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAnalyzeDocument{}, middleware.After)
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
	if err = addOpAnalyzeDocumentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAnalyzeDocument(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAnalyzeDocument(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "textract",
		OperationName: "AnalyzeDocument",
	}
}
