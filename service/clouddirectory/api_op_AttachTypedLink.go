// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/clouddirectory/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Attaches a typed link to a specified source and target object. For more
// information, see Typed Links
// (https://docs.aws.amazon.com/clouddirectory/latest/developerguide/directory_objects_links.html#directory_objects_links_typedlink).
func (c *Client) AttachTypedLink(ctx context.Context, params *AttachTypedLinkInput, optFns ...func(*Options)) (*AttachTypedLinkOutput, error) {
	if params == nil {
		params = &AttachTypedLinkInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AttachTypedLink", params, optFns, c.addOperationAttachTypedLinkMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AttachTypedLinkOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AttachTypedLinkInput struct {

	// A set of attributes that are associated with the typed link.
	//
	// This member is required.
	Attributes []types.AttributeNameAndValue

	// The Amazon Resource Name (ARN) of the directory where you want to attach the
	// typed link.
	//
	// This member is required.
	DirectoryArn *string

	// Identifies the source object that the typed link will attach to.
	//
	// This member is required.
	SourceObjectReference *types.ObjectReference

	// Identifies the target object that the typed link will attach to.
	//
	// This member is required.
	TargetObjectReference *types.ObjectReference

	// Identifies the typed link facet that is associated with the typed link.
	//
	// This member is required.
	TypedLinkFacet *types.TypedLinkSchemaAndFacetName

	noSmithyDocumentSerde
}

type AttachTypedLinkOutput struct {

	// Returns a typed link specifier as output.
	TypedLinkSpecifier *types.TypedLinkSpecifier

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAttachTypedLinkMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpAttachTypedLink{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpAttachTypedLink{}, middleware.After)
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
	if err = addOpAttachTypedLinkValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAttachTypedLink(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAttachTypedLink(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "clouddirectory",
		OperationName: "AttachTypedLink",
	}
}
