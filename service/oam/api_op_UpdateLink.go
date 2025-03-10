// Code generated by smithy-go-codegen DO NOT EDIT.

package oam

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/oam/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Use this operation to change what types of data are shared from a source account
// to its linked monitoring account sink. You can't change the sink or change the
// monitoring account with this operation. To update the list of tags associated
// with the sink, use TagResource
// (https://docs.aws.amazon.com/OAM/latest/APIReference/API_TagResource.html).
func (c *Client) UpdateLink(ctx context.Context, params *UpdateLinkInput, optFns ...func(*Options)) (*UpdateLinkOutput, error) {
	if params == nil {
		params = &UpdateLinkInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateLink", params, optFns, c.addOperationUpdateLinkMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateLinkOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateLinkInput struct {

	// The ARN of the link that you want to update.
	//
	// This member is required.
	Identifier *string

	// An array of strings that define which types of data that the source account will
	// send to the monitoring account. Your input here replaces the current set of data
	// types that are shared.
	//
	// This member is required.
	ResourceTypes []types.ResourceType

	noSmithyDocumentSerde
}

type UpdateLinkOutput struct {

	// The ARN of the link that you have updated.
	Arn *string

	// The random ID string that Amazon Web Services generated as part of the sink ARN.
	Id *string

	// The label assigned to this link, with the variables resolved to their actual
	// values.
	Label *string

	// The exact label template that was specified when the link was created, with the
	// template variables not resolved.
	LabelTemplate *string

	// The resource types now supported by this link.
	ResourceTypes []string

	// The ARN of the sink that is used for this link.
	SinkArn *string

	// The tags assigned to the link.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateLinkMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateLink{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateLink{}, middleware.After)
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
	if err = addOpUpdateLinkValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateLink(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateLink(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "oam",
		OperationName: "UpdateLink",
	}
}
