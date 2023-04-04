// Code generated by smithy-go-codegen DO NOT EDIT.

package sesv2

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/sesv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieve a list of the blacklists that your dedicated IP addresses appear on.
func (c *Client) GetBlacklistReports(ctx context.Context, params *GetBlacklistReportsInput, optFns ...func(*Options)) (*GetBlacklistReportsOutput, error) {
	if params == nil {
		params = &GetBlacklistReportsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetBlacklistReports", params, optFns, c.addOperationGetBlacklistReportsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetBlacklistReportsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to retrieve a list of the blacklists that your dedicated IP addresses
// appear on.
type GetBlacklistReportsInput struct {

	// A list of IP addresses that you want to retrieve blacklist information about.
	// You can only specify the dedicated IP addresses that you use to send email using
	// Amazon SES or Amazon Pinpoint.
	//
	// This member is required.
	BlacklistItemNames []string

	noSmithyDocumentSerde
}

// An object that contains information about blacklist events.
type GetBlacklistReportsOutput struct {

	// An object that contains information about a blacklist that one of your dedicated
	// IP addresses appears on.
	//
	// This member is required.
	BlacklistReport map[string][]types.BlacklistEntry

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetBlacklistReportsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetBlacklistReports{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetBlacklistReports{}, middleware.After)
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
	if err = addOpGetBlacklistReportsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetBlacklistReports(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetBlacklistReports(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "GetBlacklistReports",
	}
}
