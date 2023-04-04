// Code generated by smithy-go-codegen DO NOT EDIT.

package securitylake

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Designates the Amazon Security Lake delegated administrator account for the
// organization. This API can only be called by the organization management
// account. The organization management account cannot be the delegated
// administrator account.
func (c *Client) CreateDatalakeDelegatedAdmin(ctx context.Context, params *CreateDatalakeDelegatedAdminInput, optFns ...func(*Options)) (*CreateDatalakeDelegatedAdminOutput, error) {
	if params == nil {
		params = &CreateDatalakeDelegatedAdminInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDatalakeDelegatedAdmin", params, optFns, c.addOperationCreateDatalakeDelegatedAdminMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDatalakeDelegatedAdminOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDatalakeDelegatedAdminInput struct {

	// The Amazon Web Services account ID of the Security Lake delegated administrator.
	//
	// This member is required.
	Account *string

	noSmithyDocumentSerde
}

type CreateDatalakeDelegatedAdminOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDatalakeDelegatedAdminMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateDatalakeDelegatedAdmin{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateDatalakeDelegatedAdmin{}, middleware.After)
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
	if err = addOpCreateDatalakeDelegatedAdminValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDatalakeDelegatedAdmin(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDatalakeDelegatedAdmin(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securitylake",
		OperationName: "CreateDatalakeDelegatedAdmin",
	}
}
