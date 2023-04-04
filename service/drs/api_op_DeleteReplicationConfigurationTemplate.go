// Code generated by smithy-go-codegen DO NOT EDIT.

package drs

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a single Replication Configuration Template by ID
func (c *Client) DeleteReplicationConfigurationTemplate(ctx context.Context, params *DeleteReplicationConfigurationTemplateInput, optFns ...func(*Options)) (*DeleteReplicationConfigurationTemplateOutput, error) {
	if params == nil {
		params = &DeleteReplicationConfigurationTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteReplicationConfigurationTemplate", params, optFns, c.addOperationDeleteReplicationConfigurationTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteReplicationConfigurationTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteReplicationConfigurationTemplateInput struct {

	// The ID of the Replication Configuration Template to be deleted.
	//
	// This member is required.
	ReplicationConfigurationTemplateID *string

	noSmithyDocumentSerde
}

type DeleteReplicationConfigurationTemplateOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteReplicationConfigurationTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteReplicationConfigurationTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteReplicationConfigurationTemplate{}, middleware.After)
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
	if err = addOpDeleteReplicationConfigurationTemplateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteReplicationConfigurationTemplate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteReplicationConfigurationTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "drs",
		OperationName: "DeleteReplicationConfigurationTemplate",
	}
}
