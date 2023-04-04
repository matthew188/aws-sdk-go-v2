// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List the distributions that are associated with a specified WAF web ACL.
func (c *Client) ListDistributionsByWebACLId(ctx context.Context, params *ListDistributionsByWebACLIdInput, optFns ...func(*Options)) (*ListDistributionsByWebACLIdOutput, error) {
	if params == nil {
		params = &ListDistributionsByWebACLIdInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDistributionsByWebACLId", params, optFns, c.addOperationListDistributionsByWebACLIdMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDistributionsByWebACLIdOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request to list distributions that are associated with a specified WAF web
// ACL.
type ListDistributionsByWebACLIdInput struct {

	// The ID of the WAF web ACL that you want to list the associated distributions. If
	// you specify "null" for the ID, the request returns a list of the distributions
	// that aren't associated with a web ACL.
	//
	// This member is required.
	WebACLId *string

	// Use Marker and MaxItems to control pagination of results. If you have more than
	// MaxItems distributions that satisfy the request, the response includes a
	// NextMarker element. To get the next page of results, submit another request. For
	// the value of Marker, specify the value of NextMarker from the last response.
	// (For the first request, omit Marker.)
	Marker *string

	// The maximum number of distributions that you want CloudFront to return in the
	// response body. The maximum and default values are both 100.
	MaxItems *int32

	noSmithyDocumentSerde
}

// The response to a request to list the distributions that are associated with a
// specified WAF web ACL.
type ListDistributionsByWebACLIdOutput struct {

	// The DistributionList type.
	DistributionList *types.DistributionList

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDistributionsByWebACLIdMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpListDistributionsByWebACLId{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListDistributionsByWebACLId{}, middleware.After)
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
	if err = addOpListDistributionsByWebACLIdValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDistributionsByWebACLId(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListDistributionsByWebACLId(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "ListDistributionsByWebACLId",
	}
}
