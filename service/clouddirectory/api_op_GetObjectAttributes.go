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

// Retrieves attributes within a facet that are associated with an object.
func (c *Client) GetObjectAttributes(ctx context.Context, params *GetObjectAttributesInput, optFns ...func(*Options)) (*GetObjectAttributesOutput, error) {
	if params == nil {
		params = &GetObjectAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetObjectAttributes", params, optFns, c.addOperationGetObjectAttributesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetObjectAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetObjectAttributesInput struct {

	// List of attribute names whose values will be retrieved.
	//
	// This member is required.
	AttributeNames []string

	// The Amazon Resource Name (ARN) that is associated with the Directory where the
	// object resides.
	//
	// This member is required.
	DirectoryArn *string

	// Reference that identifies the object whose attributes will be retrieved.
	//
	// This member is required.
	ObjectReference *types.ObjectReference

	// Identifier for the facet whose attributes will be retrieved. See SchemaFacet for
	// details.
	//
	// This member is required.
	SchemaFacet *types.SchemaFacet

	// The consistency level at which to retrieve the attributes on an object.
	ConsistencyLevel types.ConsistencyLevel

	noSmithyDocumentSerde
}

type GetObjectAttributesOutput struct {

	// The attributes that are associated with the object.
	Attributes []types.AttributeKeyAndValue

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetObjectAttributesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetObjectAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetObjectAttributes{}, middleware.After)
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
	if err = addOpGetObjectAttributesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetObjectAttributes(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetObjectAttributes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "clouddirectory",
		OperationName: "GetObjectAttributes",
	}
}
