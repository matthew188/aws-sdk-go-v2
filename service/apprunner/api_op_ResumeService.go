// Code generated by smithy-go-codegen DO NOT EDIT.

package apprunner

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/apprunner/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Resume an active App Runner service. App Runner provisions compute capacity for
// the service. This is an asynchronous operation. On a successful call, you can
// use the returned OperationId and the ListOperations call to track the
// operation's progress.
func (c *Client) ResumeService(ctx context.Context, params *ResumeServiceInput, optFns ...func(*Options)) (*ResumeServiceOutput, error) {
	if params == nil {
		params = &ResumeServiceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ResumeService", params, optFns, c.addOperationResumeServiceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ResumeServiceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ResumeServiceInput struct {

	// The Amazon Resource Name (ARN) of the App Runner service that you want to
	// resume.
	//
	// This member is required.
	ServiceArn *string

	noSmithyDocumentSerde
}

type ResumeServiceOutput struct {

	// A description of the App Runner service that this request just resumed.
	//
	// This member is required.
	Service *types.Service

	// The unique ID of the asynchronous operation that this request started. You can
	// use it combined with the ListOperations call to track the operation's progress.
	OperationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationResumeServiceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpResumeService{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpResumeService{}, middleware.After)
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
	if err = addOpResumeServiceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opResumeService(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opResumeService(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apprunner",
		OperationName: "ResumeService",
	}
}
