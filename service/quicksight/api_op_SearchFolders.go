// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/quicksight/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Searches the subfolders in a folder.
func (c *Client) SearchFolders(ctx context.Context, params *SearchFoldersInput, optFns ...func(*Options)) (*SearchFoldersOutput, error) {
	if params == nil {
		params = &SearchFoldersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchFolders", params, optFns, c.addOperationSearchFoldersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchFoldersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchFoldersInput struct {

	// The ID for the Amazon Web Services account that contains the folder.
	//
	// This member is required.
	AwsAccountId *string

	// The filters to apply to the search. Currently, you can search only by the parent
	// folder ARN. For example, "Filters": [ { "Name": "PARENT_FOLDER_ARN", "Operator":
	// "StringEquals", "Value": "arn:aws:quicksight:us-east-1:1:folder/folderId" } ].
	//
	// This member is required.
	Filters []types.FolderSearchFilter

	// The maximum number of results to be returned per request.
	MaxResults *int32

	// The token for the next set of results, or null if there are no more results.
	NextToken *string

	noSmithyDocumentSerde
}

type SearchFoldersOutput struct {

	// A structure that contains all of the folders in the Amazon Web Services account.
	// This structure provides basic information about the folders.
	FolderSummaryList []types.FolderSummary

	// The token for the next set of results, or null if there are no more results.
	NextToken *string

	// The Amazon Web Services request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSearchFoldersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSearchFolders{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSearchFolders{}, middleware.After)
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
	if err = addOpSearchFoldersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchFolders(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSearchFolders(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "SearchFolders",
	}
}
