// Code generated by smithy-go-codegen DO NOT EDIT.

package codebuild

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/codebuild/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Imports the source repository credentials for an CodeBuild project that has its
// source code stored in a GitHub, GitHub Enterprise, or Bitbucket repository.
func (c *Client) ImportSourceCredentials(ctx context.Context, params *ImportSourceCredentialsInput, optFns ...func(*Options)) (*ImportSourceCredentialsOutput, error) {
	if params == nil {
		params = &ImportSourceCredentialsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ImportSourceCredentials", params, optFns, c.addOperationImportSourceCredentialsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ImportSourceCredentialsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ImportSourceCredentialsInput struct {

	// The type of authentication used to connect to a GitHub, GitHub Enterprise, or
	// Bitbucket repository. An OAUTH connection is not supported by the API and must
	// be created using the CodeBuild console.
	//
	// This member is required.
	AuthType types.AuthType

	// The source provider used for this project.
	//
	// This member is required.
	ServerType types.ServerType

	// For GitHub or GitHub Enterprise, this is the personal access token. For
	// Bitbucket, this is the app password.
	//
	// This member is required.
	Token *string

	// Set to false to prevent overwriting the repository source credentials. Set to
	// true to overwrite the repository source credentials. The default value is true.
	ShouldOverwrite *bool

	// The Bitbucket username when the authType is BASIC_AUTH. This parameter is not
	// valid for other types of source providers or connections.
	Username *string

	noSmithyDocumentSerde
}

type ImportSourceCredentialsOutput struct {

	// The Amazon Resource Name (ARN) of the token.
	Arn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationImportSourceCredentialsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpImportSourceCredentials{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpImportSourceCredentials{}, middleware.After)
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
	if err = addOpImportSourceCredentialsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opImportSourceCredentials(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opImportSourceCredentials(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codebuild",
		OperationName: "ImportSourceCredentials",
	}
}
