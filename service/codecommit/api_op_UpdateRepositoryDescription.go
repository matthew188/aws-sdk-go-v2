// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Sets or changes the comment or description for a repository. The description
// field for a repository accepts all HTML characters and all valid Unicode
// characters. Applications that do not HTML-encode the description and display it
// in a webpage can expose users to potentially malicious code. Make sure that you
// HTML-encode the description field in any application that uses this API to
// display the repository description on a webpage.
func (c *Client) UpdateRepositoryDescription(ctx context.Context, params *UpdateRepositoryDescriptionInput, optFns ...func(*Options)) (*UpdateRepositoryDescriptionOutput, error) {
	if params == nil {
		params = &UpdateRepositoryDescriptionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateRepositoryDescription", params, optFns, c.addOperationUpdateRepositoryDescriptionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateRepositoryDescriptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of an update repository description operation.
type UpdateRepositoryDescriptionInput struct {

	// The name of the repository to set or change the comment or description for.
	//
	// This member is required.
	RepositoryName *string

	// The new comment or description for the specified repository. Repository
	// descriptions are limited to 1,000 characters.
	RepositoryDescription *string

	noSmithyDocumentSerde
}

type UpdateRepositoryDescriptionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateRepositoryDescriptionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateRepositoryDescription{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateRepositoryDescription{}, middleware.After)
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
	if err = addOpUpdateRepositoryDescriptionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateRepositoryDescription(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateRepositoryDescription(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codecommit",
		OperationName: "UpdateRepositoryDescription",
	}
}
