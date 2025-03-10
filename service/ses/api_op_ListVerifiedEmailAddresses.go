// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deprecated. Use the ListIdentities operation to list the email addresses and
// domains associated with your account.
func (c *Client) ListVerifiedEmailAddresses(ctx context.Context, params *ListVerifiedEmailAddressesInput, optFns ...func(*Options)) (*ListVerifiedEmailAddressesOutput, error) {
	if params == nil {
		params = &ListVerifiedEmailAddressesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListVerifiedEmailAddresses", params, optFns, c.addOperationListVerifiedEmailAddressesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListVerifiedEmailAddressesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListVerifiedEmailAddressesInput struct {
	noSmithyDocumentSerde
}

// A list of email addresses that you have verified with Amazon SES under your AWS
// account.
type ListVerifiedEmailAddressesOutput struct {

	// A list of email addresses that have been verified.
	VerifiedEmailAddresses []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListVerifiedEmailAddressesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpListVerifiedEmailAddresses{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpListVerifiedEmailAddresses{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListVerifiedEmailAddresses(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListVerifiedEmailAddresses(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "ListVerifiedEmailAddresses",
	}
}
