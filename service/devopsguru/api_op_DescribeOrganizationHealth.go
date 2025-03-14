// Code generated by smithy-go-codegen DO NOT EDIT.

package devopsguru

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns active insights, predictive insights, and resource hours analyzed in
// last hour.
func (c *Client) DescribeOrganizationHealth(ctx context.Context, params *DescribeOrganizationHealthInput, optFns ...func(*Options)) (*DescribeOrganizationHealthOutput, error) {
	if params == nil {
		params = &DescribeOrganizationHealthInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeOrganizationHealth", params, optFns, c.addOperationDescribeOrganizationHealthMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeOrganizationHealthOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeOrganizationHealthInput struct {

	// The ID of the Amazon Web Services account.
	AccountIds []string

	// The ID of the organizational unit.
	OrganizationalUnitIds []string

	noSmithyDocumentSerde
}

type DescribeOrganizationHealthOutput struct {

	// An integer that specifies the number of metrics that have been analyzed in your
	// organization.
	//
	// This member is required.
	MetricsAnalyzed int32

	// An integer that specifies the number of open proactive insights in your Amazon
	// Web Services account.
	//
	// This member is required.
	OpenProactiveInsights int32

	// An integer that specifies the number of open reactive insights in your Amazon
	// Web Services account.
	//
	// This member is required.
	OpenReactiveInsights int32

	// The number of Amazon DevOps Guru resource analysis hours billed to the current
	// Amazon Web Services account in the last hour.
	//
	// This member is required.
	ResourceHours *int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeOrganizationHealthMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeOrganizationHealth{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeOrganizationHealth{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeOrganizationHealth(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeOrganizationHealth(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devops-guru",
		OperationName: "DescribeOrganizationHealth",
	}
}
