// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/apigateway/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets information about a Deployments collection.
func (c *Client) GetDeployments(ctx context.Context, params *GetDeploymentsInput, optFns ...func(*Options)) (*GetDeploymentsOutput, error) {
	if params == nil {
		params = &GetDeploymentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDeployments", params, optFns, c.addOperationGetDeploymentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDeploymentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Requests API Gateway to get information about a Deployments collection.
type GetDeploymentsInput struct {

	// The string identifier of the associated RestApi.
	//
	// This member is required.
	RestApiId *string

	// The maximum number of returned results per page. The default value is 25 and the
	// maximum value is 500.
	Limit *int32

	// The current pagination position in the paged result set.
	Position *string

	noSmithyDocumentSerde
}

// Represents a collection resource that contains zero or more references to your
// existing deployments, and links that guide you on how to interact with your
// collection. The collection offers a paginated view of the contained deployments.
type GetDeploymentsOutput struct {

	// The current page of elements from this collection.
	Items []types.Deployment

	// The current pagination position in the paged result set.
	Position *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetDeploymentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetDeployments{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDeployments{}, middleware.After)
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
	if err = addOpGetDeploymentsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDeployments(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addAcceptHeader(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

// GetDeploymentsAPIClient is a client that implements the GetDeployments
// operation.
type GetDeploymentsAPIClient interface {
	GetDeployments(context.Context, *GetDeploymentsInput, ...func(*Options)) (*GetDeploymentsOutput, error)
}

var _ GetDeploymentsAPIClient = (*Client)(nil)

// GetDeploymentsPaginatorOptions is the paginator options for GetDeployments
type GetDeploymentsPaginatorOptions struct {
	// The maximum number of returned results per page. The default value is 25 and the
	// maximum value is 500.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetDeploymentsPaginator is a paginator for GetDeployments
type GetDeploymentsPaginator struct {
	options   GetDeploymentsPaginatorOptions
	client    GetDeploymentsAPIClient
	params    *GetDeploymentsInput
	nextToken *string
	firstPage bool
}

// NewGetDeploymentsPaginator returns a new GetDeploymentsPaginator
func NewGetDeploymentsPaginator(client GetDeploymentsAPIClient, params *GetDeploymentsInput, optFns ...func(*GetDeploymentsPaginatorOptions)) *GetDeploymentsPaginator {
	if params == nil {
		params = &GetDeploymentsInput{}
	}

	options := GetDeploymentsPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetDeploymentsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Position,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetDeploymentsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetDeployments page.
func (p *GetDeploymentsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetDeploymentsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Position = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	result, err := p.client.GetDeployments(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Position

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opGetDeployments(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "GetDeployments",
	}
}
