// Code generated by smithy-go-codegen DO NOT EDIT.

package costexplorer

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/costexplorer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// An API operation for adding one or more tags (key-value pairs) to a resource.
// You can use the TagResource operation with a resource that already has tags. If
// you specify a new tag key for the resource, this tag is appended to the list of
// tags associated with the resource. If you specify a tag key that is already
// associated with the resource, the new tag value you specify replaces the
// previous value for that tag. Although the maximum number of array members is
// 200, user-tag maximum is 50. The remaining are reserved for Amazon Web Services
// use.
func (c *Client) TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) {
	if params == nil {
		params = &TagResourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "TagResource", params, optFns, c.addOperationTagResourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*TagResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TagResourceInput struct {

	// The Amazon Resource Name (ARN) of the resource. For a list of supported
	// resources, see ResourceTag
	// (https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_ResourceTag.html).
	//
	// This member is required.
	ResourceArn *string

	// A list of tag key-value pairs to be added to the resource. Each tag consists of
	// a key and a value, and each key must be unique for the resource. The following
	// restrictions apply to resource tags:
	//
	// * Although the maximum number of array
	// members is 200, you can assign a maximum of 50 user-tags to one resource. The
	// remaining are reserved for Amazon Web Services use
	//
	// * The maximum length of a
	// key is 128 characters
	//
	// * The maximum length of a value is 256 characters
	//
	// * Keys
	// and values can only contain alphanumeric characters, spaces, and any of the
	// following: _.:/=+@-
	//
	// * Keys and values are case sensitive
	//
	// * Keys and values are
	// trimmed for any leading or trailing whitespaces
	//
	// * Don’t use aws: as a prefix
	// for your keys. This prefix is reserved for Amazon Web Services use
	//
	// This member is required.
	ResourceTags []types.ResourceTag

	noSmithyDocumentSerde
}

type TagResourceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationTagResourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpTagResource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpTagResource{}, middleware.After)
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
	if err = addOpTagResourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opTagResource(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opTagResource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ce",
		OperationName: "TagResource",
	}
}
