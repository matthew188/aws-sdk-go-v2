// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/glue/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Delete the entire registry including schema and all of its versions. To get the
// status of the delete operation, you can call the GetRegistry API after the
// asynchronous call. Deleting a registry will deactivate all online operations for
// the registry such as the UpdateRegistry, CreateSchema, UpdateSchema, and
// RegisterSchemaVersion APIs.
func (c *Client) DeleteRegistry(ctx context.Context, params *DeleteRegistryInput, optFns ...func(*Options)) (*DeleteRegistryOutput, error) {
	if params == nil {
		params = &DeleteRegistryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteRegistry", params, optFns, c.addOperationDeleteRegistryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteRegistryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteRegistryInput struct {

	// This is a wrapper structure that may contain the registry name and Amazon
	// Resource Name (ARN).
	//
	// This member is required.
	RegistryId *types.RegistryId

	noSmithyDocumentSerde
}

type DeleteRegistryOutput struct {

	// The Amazon Resource Name (ARN) of the registry being deleted.
	RegistryArn *string

	// The name of the registry being deleted.
	RegistryName *string

	// The status of the registry. A successful operation will return the Deleting
	// status.
	Status types.RegistryStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteRegistryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteRegistry{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteRegistry{}, middleware.After)
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
	if err = addOpDeleteRegistryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteRegistry(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteRegistry(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "DeleteRegistry",
	}
}
