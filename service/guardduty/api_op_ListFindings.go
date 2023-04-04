// Code generated by smithy-go-codegen DO NOT EDIT.

package guardduty

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/guardduty/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists Amazon GuardDuty findings for the specified detector ID.
func (c *Client) ListFindings(ctx context.Context, params *ListFindingsInput, optFns ...func(*Options)) (*ListFindingsOutput, error) {
	if params == nil {
		params = &ListFindingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListFindings", params, optFns, c.addOperationListFindingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListFindingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListFindingsInput struct {

	// The ID of the detector that specifies the GuardDuty service whose findings you
	// want to list.
	//
	// This member is required.
	DetectorId *string

	// Represents the criteria used for querying findings. Valid values include:
	//
	// *
	// JSON field name
	//
	// * accountId
	//
	// * region
	//
	// * confidence
	//
	// * id
	//
	// *
	// resource.accessKeyDetails.accessKeyId
	//
	// *
	// resource.accessKeyDetails.principalId
	//
	// * resource.accessKeyDetails.userName
	//
	// *
	// resource.accessKeyDetails.userType
	//
	// *
	// resource.instanceDetails.iamInstanceProfile.id
	//
	// *
	// resource.instanceDetails.imageId
	//
	// * resource.instanceDetails.instanceId
	//
	// *
	// resource.instanceDetails.networkInterfaces.ipv6Addresses
	//
	// *
	// resource.instanceDetails.networkInterfaces.privateIpAddresses.privateIpAddress
	//
	// *
	// resource.instanceDetails.networkInterfaces.publicDnsName
	//
	// *
	// resource.instanceDetails.networkInterfaces.publicIp
	//
	// *
	// resource.instanceDetails.networkInterfaces.securityGroups.groupId
	//
	// *
	// resource.instanceDetails.networkInterfaces.securityGroups.groupName
	//
	// *
	// resource.instanceDetails.networkInterfaces.subnetId
	//
	// *
	// resource.instanceDetails.networkInterfaces.vpcId
	//
	// *
	// resource.instanceDetails.tags.key
	//
	// * resource.instanceDetails.tags.value
	//
	// *
	// resource.resourceType
	//
	// * service.action.actionType
	//
	// *
	// service.action.awsApiCallAction.api
	//
	// *
	// service.action.awsApiCallAction.callerType
	//
	// *
	// service.action.awsApiCallAction.remoteIpDetails.city.cityName
	//
	// *
	// service.action.awsApiCallAction.remoteIpDetails.country.countryName
	//
	// *
	// service.action.awsApiCallAction.remoteIpDetails.ipAddressV4
	//
	// *
	// service.action.awsApiCallAction.remoteIpDetails.organization.asn
	//
	// *
	// service.action.awsApiCallAction.remoteIpDetails.organization.asnOrg
	//
	// *
	// service.action.awsApiCallAction.serviceName
	//
	// *
	// service.action.dnsRequestAction.domain
	//
	// *
	// service.action.networkConnectionAction.blocked
	//
	// *
	// service.action.networkConnectionAction.connectionDirection
	//
	// *
	// service.action.networkConnectionAction.localPortDetails.port
	//
	// *
	// service.action.networkConnectionAction.protocol
	//
	// *
	// service.action.networkConnectionAction.remoteIpDetails.country.countryName
	//
	// *
	// service.action.networkConnectionAction.remoteIpDetails.ipAddressV4
	//
	// *
	// service.action.networkConnectionAction.remoteIpDetails.organization.asn
	//
	// *
	// service.action.networkConnectionAction.remoteIpDetails.organization.asnOrg
	//
	// *
	// service.action.networkConnectionAction.remotePortDetails.port
	//
	// *
	// service.additionalInfo.threatListName
	//
	// * service.archived When this attribute is
	// set to 'true', only archived findings are listed. When it's set to 'false', only
	// unarchived findings are listed. When this attribute is not set, all existing
	// findings are listed.
	//
	// * service.resourceRole
	//
	// * severity
	//
	// * type
	//
	// * updatedAt
	// Type: Timestamp in Unix Epoch millisecond format: 1486685375000
	FindingCriteria *types.FindingCriteria

	// You can use this parameter to indicate the maximum number of items you want in
	// the response. The default value is 50. The maximum value is 50.
	MaxResults int32

	// You can use this parameter when paginating results. Set the value of this
	// parameter to null on your first call to the list action. For subsequent calls to
	// the action, fill nextToken in the request with the value of NextToken from the
	// previous response to continue listing data.
	NextToken *string

	// Represents the criteria used for sorting findings.
	SortCriteria *types.SortCriteria

	noSmithyDocumentSerde
}

type ListFindingsOutput struct {

	// The IDs of the findings that you're listing.
	//
	// This member is required.
	FindingIds []string

	// The pagination parameter to be used on the next list operation to retrieve more
	// items.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListFindingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListFindings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListFindings{}, middleware.After)
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
	if err = addOpListFindingsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListFindings(options.Region), middleware.Before); err != nil {
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

// ListFindingsAPIClient is a client that implements the ListFindings operation.
type ListFindingsAPIClient interface {
	ListFindings(context.Context, *ListFindingsInput, ...func(*Options)) (*ListFindingsOutput, error)
}

var _ ListFindingsAPIClient = (*Client)(nil)

// ListFindingsPaginatorOptions is the paginator options for ListFindings
type ListFindingsPaginatorOptions struct {
	// You can use this parameter to indicate the maximum number of items you want in
	// the response. The default value is 50. The maximum value is 50.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListFindingsPaginator is a paginator for ListFindings
type ListFindingsPaginator struct {
	options   ListFindingsPaginatorOptions
	client    ListFindingsAPIClient
	params    *ListFindingsInput
	nextToken *string
	firstPage bool
}

// NewListFindingsPaginator returns a new ListFindingsPaginator
func NewListFindingsPaginator(client ListFindingsAPIClient, params *ListFindingsInput, optFns ...func(*ListFindingsPaginatorOptions)) *ListFindingsPaginator {
	if params == nil {
		params = &ListFindingsInput{}
	}

	options := ListFindingsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListFindingsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListFindingsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListFindings page.
func (p *ListFindingsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListFindingsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListFindings(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListFindings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "guardduty",
		OperationName: "ListFindings",
	}
}
