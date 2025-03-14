// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/ssm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Removes tag keys from the specified resource.
func (c *Client) RemoveTagsFromResource(ctx context.Context, params *RemoveTagsFromResourceInput, optFns ...func(*Options)) (*RemoveTagsFromResourceOutput, error) {
	if params == nil {
		params = &RemoveTagsFromResourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RemoveTagsFromResource", params, optFns, c.addOperationRemoveTagsFromResourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RemoveTagsFromResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RemoveTagsFromResourceInput struct {

	// The ID of the resource from which you want to remove tags. For example:
	// ManagedInstance: mi-012345abcde MaintenanceWindow: mw-012345abcde Automation:
	// example-c160-4567-8519-012345abcde PatchBaseline: pb-012345abcde OpsMetadata
	// object: ResourceID for tagging is created from the Amazon Resource Name (ARN)
	// for the object. Specifically, ResourceID is created from the strings that come
	// after the word opsmetadata in the ARN. For example, an OpsMetadata object with
	// an ARN of
	// arn:aws:ssm:us-east-2:1234567890:opsmetadata/aws/ssm/MyGroup/appmanager has a
	// ResourceID of either aws/ssm/MyGroup/appmanager or /aws/ssm/MyGroup/appmanager.
	// For the Document and Parameter values, use the name of the resource. The
	// ManagedInstance type for this API operation is only for on-premises managed
	// nodes. Specify the name of the managed node in the following format:
	// mi-ID_number. For example, mi-1a2b3c4d5e6f.
	//
	// This member is required.
	ResourceId *string

	// The type of resource from which you want to remove a tag. The ManagedInstance
	// type for this API operation is only for on-premises managed nodes. Specify the
	// name of the managed node in the following format: mi-ID_number . For example,
	// mi-1a2b3c4d5e6f.
	//
	// This member is required.
	ResourceType types.ResourceTypeForTagging

	// Tag keys that you want to remove from the specified resource.
	//
	// This member is required.
	TagKeys []string

	noSmithyDocumentSerde
}

type RemoveTagsFromResourceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRemoveTagsFromResourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRemoveTagsFromResource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRemoveTagsFromResource{}, middleware.After)
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
	if err = addOpRemoveTagsFromResourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRemoveTagsFromResource(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRemoveTagsFromResource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "RemoveTagsFromResource",
	}
}
