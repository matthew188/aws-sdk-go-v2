// Code generated by smithy-go-codegen DO NOT EDIT.

package finspacedata

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/finspacedata/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Request programmatic credentials to use with FinSpace SDK.
func (c *Client) GetProgrammaticAccessCredentials(ctx context.Context, params *GetProgrammaticAccessCredentialsInput, optFns ...func(*Options)) (*GetProgrammaticAccessCredentialsOutput, error) {
	if params == nil {
		params = &GetProgrammaticAccessCredentialsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetProgrammaticAccessCredentials", params, optFns, c.addOperationGetProgrammaticAccessCredentialsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetProgrammaticAccessCredentialsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request for GetProgrammaticAccessCredentials operation
type GetProgrammaticAccessCredentialsInput struct {

	// The FinSpace environment identifier.
	//
	// This member is required.
	EnvironmentId *string

	// The time duration in which the credentials remain valid.
	DurationInMinutes int64

	noSmithyDocumentSerde
}

// Response for GetProgrammaticAccessCredentials operation
type GetProgrammaticAccessCredentialsOutput struct {

	// Returns the programmatic credentials.
	Credentials *types.Credentials

	// Returns the duration in which the credentials will remain valid.
	DurationInMinutes int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetProgrammaticAccessCredentialsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetProgrammaticAccessCredentials{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetProgrammaticAccessCredentials{}, middleware.After)
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
	if err = addRestJsonContentTypeCustomization(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetProgrammaticAccessCredentialsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetProgrammaticAccessCredentials(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetProgrammaticAccessCredentials(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "finspace-api",
		OperationName: "GetProgrammaticAccessCredentials",
	}
}
