// Code generated by smithy-go-codegen DO NOT EDIT.

package m2

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a specific application from the specific runtime environment where it
// was previously deployed. You cannot delete a runtime environment using
// DeleteEnvironment if any application has ever been deployed to it. This API
// removes the association of the application with the runtime environment so you
// can delete the environment smoothly.
func (c *Client) DeleteApplicationFromEnvironment(ctx context.Context, params *DeleteApplicationFromEnvironmentInput, optFns ...func(*Options)) (*DeleteApplicationFromEnvironmentOutput, error) {
	if params == nil {
		params = &DeleteApplicationFromEnvironmentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteApplicationFromEnvironment", params, optFns, c.addOperationDeleteApplicationFromEnvironmentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteApplicationFromEnvironmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteApplicationFromEnvironmentInput struct {

	// The unique identifier of the application you want to delete.
	//
	// This member is required.
	ApplicationId *string

	// The unique identifier of the runtime environment where the application was
	// previously deployed.
	//
	// This member is required.
	EnvironmentId *string

	noSmithyDocumentSerde
}

type DeleteApplicationFromEnvironmentOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteApplicationFromEnvironmentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteApplicationFromEnvironment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteApplicationFromEnvironment{}, middleware.After)
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
	if err = addOpDeleteApplicationFromEnvironmentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteApplicationFromEnvironment(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteApplicationFromEnvironment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "m2",
		OperationName: "DeleteApplicationFromEnvironment",
	}
}
