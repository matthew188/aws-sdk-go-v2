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

// Inspects the text of a batch of documents for named entities and returns
// information about them. For more information about named entities, see Entities
// (https://docs.aws.amazon.com/comprehend/latest/dg/how-entities.html) in the
// Comprehend Developer Guide.
func (c *Client) BatchDetectEntities(ctx context.Context, params *BatchDetectEntitiesInput, optFns ...func(*Options)) (*BatchDetectEntitiesOutput, error) {
	if params == nil {
		params = &BatchDetectEntitiesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchDetectEntities", params, optFns, c.addOperationBatchDetectEntitiesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchDetectEntitiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchDetectEntitiesInput struct {

	// The language of the input documents. You can specify any of the primary
	// languages supported by Amazon Comprehend. All documents must be in the same
	// language.
	//
	// This member is required.
	LanguageCode types.LanguageCode

	// A list containing the UTF-8 encoded text of the input documents. The list can
	// contain a maximum of 25 documents. The maximum size of each document is 5 KB.
	//
	// This member is required.
	TextList []string

	noSmithyDocumentSerde
}

type BatchDetectEntitiesOutput struct {

	// A list containing one object for each document that contained an error. The
	// results are sorted in ascending order by the Index field and match the order of
	// the documents in the input list. If there are no errors in the batch, the
	// ErrorList is empty.
	//
	// This member is required.
	ErrorList []types.BatchItemError

	// A list of objects containing the results of the operation. The results are
	// sorted in ascending order by the Index field and match the order of the
	// documents in the input list. If all of the documents contain an error, the
	// ResultList is empty.
	//
	// This member is required.
	ResultList []types.BatchDetectEntitiesItemResult

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchDetectEntitiesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpBatchDetectEntities{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchDetectEntities{}, middleware.After)
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
	if err = addOpBatchDetectEntitiesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchDetectEntities(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchDetectEntities(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehend",
		OperationName: "BatchDetectEntities",
	}
}
