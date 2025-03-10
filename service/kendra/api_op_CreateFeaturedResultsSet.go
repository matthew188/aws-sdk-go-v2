// Code generated by smithy-go-codegen DO NOT EDIT.

package kendra

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/kendra/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a set of featured results to display at the top of the search results
// page. Featured results are placed above all other results for certain queries.
// You map specific queries to specific documents for featuring in the results. If
// a query contains an exact match, then one or more specific documents are
// featured in the search results. You can create up to 50 sets of featured results
// per index. You can request to increase this limit by contacting Support
// (http://aws.amazon.com/contact-us/).
func (c *Client) CreateFeaturedResultsSet(ctx context.Context, params *CreateFeaturedResultsSetInput, optFns ...func(*Options)) (*CreateFeaturedResultsSetOutput, error) {
	if params == nil {
		params = &CreateFeaturedResultsSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateFeaturedResultsSet", params, optFns, c.addOperationCreateFeaturedResultsSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateFeaturedResultsSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateFeaturedResultsSetInput struct {

	// A name for the set of featured results.
	//
	// This member is required.
	FeaturedResultsSetName *string

	// The identifier of the index that you want to use for featuring results.
	//
	// This member is required.
	IndexId *string

	// A token that you provide to identify the request to create a set of featured
	// results. Multiple calls to the CreateFeaturedResultsSet API with the same client
	// token will create only one featured results set.
	ClientToken *string

	// A description for the set of featured results.
	Description *string

	// A list of document IDs for the documents you want to feature at the top of the
	// search results page. For more information on the list of documents, see
	// FeaturedResultsSet
	// (https://docs.aws.amazon.com/kendra/latest/dg/API_FeaturedResultsSet.html).
	FeaturedDocuments []types.FeaturedDocument

	// A list of queries for featuring results. For more information on the list of
	// queries, see FeaturedResultsSet
	// (https://docs.aws.amazon.com/kendra/latest/dg/API_FeaturedResultsSet.html).
	QueryTexts []string

	// The current status of the set of featured results. When the value is ACTIVE,
	// featured results are ready for use. You can still configure your settings before
	// setting the status to ACTIVE. You can set the status to ACTIVE or INACTIVE using
	// the UpdateFeaturedResultsSet
	// (https://docs.aws.amazon.com/kendra/latest/dg/API_UpdateFeaturedResultsSet.html)
	// API. The queries you specify for featured results must be unique per featured
	// results set for each index, whether the status is ACTIVE or INACTIVE.
	Status types.FeaturedResultsSetStatus

	// A list of key-value pairs that identify or categorize the featured results set.
	// You can also use tags to help control access to the featured results set. Tag
	// keys and values can consist of Unicode letters, digits, white space, and any of
	// the following symbols:_ . : / = + - @.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateFeaturedResultsSetOutput struct {

	// Information on the set of featured results. This includes the identifier of the
	// featured results set, whether the featured results set is active or inactive,
	// when the featured results set was created, and more.
	FeaturedResultsSet *types.FeaturedResultsSet

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateFeaturedResultsSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateFeaturedResultsSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateFeaturedResultsSet{}, middleware.After)
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
	if err = addOpCreateFeaturedResultsSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateFeaturedResultsSet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateFeaturedResultsSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kendra",
		OperationName: "CreateFeaturedResultsSet",
	}
}
