// Code generated by smithy-go-codegen DO NOT EDIT.

package inspector

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/inspector/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Removes entire attributes (key and value pairs) from the findings that are
// specified by the ARNs of the findings where an attribute with the specified key
// exists.
func (c *Client) RemoveAttributesFromFindings(ctx context.Context, params *RemoveAttributesFromFindingsInput, optFns ...func(*Options)) (*RemoveAttributesFromFindingsOutput, error) {
	if params == nil {
		params = &RemoveAttributesFromFindingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RemoveAttributesFromFindings", params, optFns, c.addOperationRemoveAttributesFromFindingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RemoveAttributesFromFindingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RemoveAttributesFromFindingsInput struct {

	// The array of attribute keys that you want to remove from specified findings.
	//
	// This member is required.
	AttributeKeys []string

	// The ARNs that specify the findings that you want to remove attributes from.
	//
	// This member is required.
	FindingArns []string

	noSmithyDocumentSerde
}

type RemoveAttributesFromFindingsOutput struct {

	// Attributes details that cannot be described. An error code is provided for each
	// failed item.
	//
	// This member is required.
	FailedItems map[string]types.FailedItemDetails

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRemoveAttributesFromFindingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRemoveAttributesFromFindings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRemoveAttributesFromFindings{}, middleware.After)
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
	if err = addOpRemoveAttributesFromFindingsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRemoveAttributesFromFindings(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRemoveAttributesFromFindings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "inspector",
		OperationName: "RemoveAttributesFromFindings",
	}
}
