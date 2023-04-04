// Code generated by smithy-go-codegen DO NOT EDIT.

package supportapp

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a Slack workspace configuration from your Amazon Web Services account.
// This operation doesn't delete your Slack workspace.
func (c *Client) DeleteSlackWorkspaceConfiguration(ctx context.Context, params *DeleteSlackWorkspaceConfigurationInput, optFns ...func(*Options)) (*DeleteSlackWorkspaceConfigurationOutput, error) {
	if params == nil {
		params = &DeleteSlackWorkspaceConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteSlackWorkspaceConfiguration", params, optFns, c.addOperationDeleteSlackWorkspaceConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteSlackWorkspaceConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteSlackWorkspaceConfigurationInput struct {

	// The team ID in Slack. This ID uniquely identifies a Slack workspace, such as
	// T012ABCDEFG.
	//
	// This member is required.
	TeamId *string

	noSmithyDocumentSerde
}

type DeleteSlackWorkspaceConfigurationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteSlackWorkspaceConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteSlackWorkspaceConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteSlackWorkspaceConfiguration{}, middleware.After)
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
	if err = addOpDeleteSlackWorkspaceConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteSlackWorkspaceConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteSlackWorkspaceConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "supportapp",
		OperationName: "DeleteSlackWorkspaceConfiguration",
	}
}
