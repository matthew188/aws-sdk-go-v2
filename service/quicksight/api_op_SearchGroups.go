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

// Use the SearchGroups operation to search groups in a specified Amazon QuickSight
// namespace using the supplied filters.
func (c *Client) SearchGroups(ctx context.Context, params *SearchGroupsInput, optFns ...func(*Options)) (*SearchGroupsOutput, error) {
	if params == nil {
		params = &SearchGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchGroups", params, optFns, c.addOperationSearchGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchGroupsInput struct {

	// The ID for the Amazon Web Services account that the group is in. Currently, you
	// use the ID for the Amazon Web Services account that contains your Amazon
	// QuickSight account.
	//
	// This member is required.
	AwsAccountId *string

	// The structure for the search filters that you want to apply to your search.
	//
	// This member is required.
	Filters []types.GroupSearchFilter

	// The namespace that you want to search.
	//
	// This member is required.
	Namespace *string

	// The maximum number of results to return from this request.
	MaxResults *int32

	// A pagination token that can be used in a subsequent request.
	NextToken *string

	noSmithyDocumentSerde
}

type SearchGroupsOutput struct {

	// A list of groups in a specified namespace that match the filters you set in your
	// SearchGroups request.
	GroupList []types.Group

	// A pagination token that can be used in a subsequent request.
	NextToken *string

	// The Amazon Web Services request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSearchGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSearchGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSearchGroups{}, middleware.After)
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
	if err = addOpSearchGroupsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchGroups(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSearchGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "SearchGroups",
	}
}
