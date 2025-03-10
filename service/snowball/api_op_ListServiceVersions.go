// Code generated by smithy-go-codegen DO NOT EDIT.

package snowball

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/snowball/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all supported versions for Snow on-device services. Returns an array of
// ServiceVersion object containing the supported versions for a particular
// service.
func (c *Client) ListServiceVersions(ctx context.Context, params *ListServiceVersionsInput, optFns ...func(*Options)) (*ListServiceVersionsOutput, error) {
	if params == nil {
		params = &ListServiceVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListServiceVersions", params, optFns, c.addOperationListServiceVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListServiceVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListServiceVersionsInput struct {

	// The name of the service for which you're requesting supported versions.
	//
	// This member is required.
	ServiceName types.ServiceName

	// A list of names and versions of dependant services of the requested service.
	DependentServices []types.DependentService

	// The maximum number of ListServiceVersions objects to return.
	MaxResults *int32

	// Because HTTP requests are stateless, this is the starting point for the next
	// list of returned ListServiceVersionsRequest versions.
	NextToken *string

	noSmithyDocumentSerde
}

type ListServiceVersionsOutput struct {

	// The name of the service for which the system provided supported versions.
	//
	// This member is required.
	ServiceName types.ServiceName

	// A list of supported versions.
	//
	// This member is required.
	ServiceVersions []types.ServiceVersion

	// A list of names and versions of dependant services of the service for which the
	// system provided supported versions.
	DependentServices []types.DependentService

	// Because HTTP requests are stateless, this is the starting point of the next list
	// of returned ListServiceVersionsResult results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListServiceVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListServiceVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListServiceVersions{}, middleware.After)
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
	if err = addOpListServiceVersionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListServiceVersions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListServiceVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "snowball",
		OperationName: "ListServiceVersions",
	}
}
