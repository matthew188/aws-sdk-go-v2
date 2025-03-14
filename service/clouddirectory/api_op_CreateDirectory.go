// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a Directory by copying the published schema into the directory. A
// directory cannot be created without a schema. You can also quickly create a
// directory using a managed schema, called the QuickStartSchema. For more
// information, see Managed Schema
// (https://docs.aws.amazon.com/clouddirectory/latest/developerguide/schemas_managed.html)
// in the Amazon Cloud Directory Developer Guide.
func (c *Client) CreateDirectory(ctx context.Context, params *CreateDirectoryInput, optFns ...func(*Options)) (*CreateDirectoryOutput, error) {
	if params == nil {
		params = &CreateDirectoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDirectory", params, optFns, c.addOperationCreateDirectoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDirectoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDirectoryInput struct {

	// The name of the Directory. Should be unique per account, per region.
	//
	// This member is required.
	Name *string

	// The Amazon Resource Name (ARN) of the published schema that will be copied into
	// the data Directory. For more information, see arns.
	//
	// This member is required.
	SchemaArn *string

	noSmithyDocumentSerde
}

type CreateDirectoryOutput struct {

	// The ARN of the published schema in the Directory. Once a published schema is
	// copied into the directory, it has its own ARN, which is referred to applied
	// schema ARN. For more information, see arns.
	//
	// This member is required.
	AppliedSchemaArn *string

	// The ARN that is associated with the Directory. For more information, see arns.
	//
	// This member is required.
	DirectoryArn *string

	// The name of the Directory.
	//
	// This member is required.
	Name *string

	// The root object node of the created directory.
	//
	// This member is required.
	ObjectIdentifier *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDirectoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateDirectory{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateDirectory{}, middleware.After)
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
	if err = addOpCreateDirectoryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDirectory(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDirectory(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "clouddirectory",
		OperationName: "CreateDirectory",
	}
}
