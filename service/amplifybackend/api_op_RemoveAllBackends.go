// Code generated by smithy-go-codegen DO NOT EDIT.

package amplifybackend

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Removes all backend environments from your Amplify project.
func (c *Client) RemoveAllBackends(ctx context.Context, params *RemoveAllBackendsInput, optFns ...func(*Options)) (*RemoveAllBackendsOutput, error) {
	if params == nil {
		params = &RemoveAllBackendsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RemoveAllBackends", params, optFns, c.addOperationRemoveAllBackendsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RemoveAllBackendsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request body for RemoveAllBackends.
type RemoveAllBackendsInput struct {

	// The app ID.
	//
	// This member is required.
	AppId *string

	// Cleans up the Amplify Console app if this value is set to true.
	CleanAmplifyApp bool

	noSmithyDocumentSerde
}

type RemoveAllBackendsOutput struct {

	// The app ID.
	AppId *string

	// If the request fails, this error is returned.
	Error *string

	// The ID for the job.
	JobId *string

	// The name of the operation.
	Operation *string

	// The current status of the request.
	Status *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRemoveAllBackendsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpRemoveAllBackends{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpRemoveAllBackends{}, middleware.After)
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
	if err = addOpRemoveAllBackendsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRemoveAllBackends(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRemoveAllBackends(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "amplifybackend",
		OperationName: "RemoveAllBackends",
	}
}
