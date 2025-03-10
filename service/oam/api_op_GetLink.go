// Code generated by smithy-go-codegen DO NOT EDIT.

package oam

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns complete information about one link. To use this operation, provide the
// link ARN. To retrieve a list of link ARNs, use ListLinks
// (https://docs.aws.amazon.com/OAM/latest/APIReference/API_ListLinks.html).
func (c *Client) GetLink(ctx context.Context, params *GetLinkInput, optFns ...func(*Options)) (*GetLinkOutput, error) {
	if params == nil {
		params = &GetLinkInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetLink", params, optFns, c.addOperationGetLinkMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetLinkOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetLinkInput struct {

	// The ARN of the link to retrieve information for.
	//
	// This member is required.
	Identifier *string

	noSmithyDocumentSerde
}

type GetLinkOutput struct {

	// The ARN of the link.
	Arn *string

	// The random ID string that Amazon Web Services generated as part of the link ARN.
	Id *string

	// The label that you assigned to this link, with the variables resolved to their
	// actual values.
	Label *string

	// The exact label template that was specified when the link was created, with the
	// template variables not resolved.
	LabelTemplate *string

	// The resource types supported by this link.
	ResourceTypes []string

	// The ARN of the sink that is used for this link.
	SinkArn *string

	// The tags assigned to the link.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetLinkMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetLink{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetLink{}, middleware.After)
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
	if err = addOpGetLinkValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetLink(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetLink(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "oam",
		OperationName: "GetLink",
	}
}
