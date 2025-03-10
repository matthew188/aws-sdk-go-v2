// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the specified organization conformance pack and all of the Config rules
// and remediation actions from all member accounts in that organization. Only a
// management account or a delegated administrator account can delete an
// organization conformance pack. When calling this API with a delegated
// administrator, you must ensure Organizations ListDelegatedAdministrator
// permissions are added. Config sets the state of a conformance pack to
// DELETE_IN_PROGRESS until the deletion is complete. You cannot update a
// conformance pack while it is in this state.
func (c *Client) DeleteOrganizationConformancePack(ctx context.Context, params *DeleteOrganizationConformancePackInput, optFns ...func(*Options)) (*DeleteOrganizationConformancePackOutput, error) {
	if params == nil {
		params = &DeleteOrganizationConformancePackInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteOrganizationConformancePack", params, optFns, c.addOperationDeleteOrganizationConformancePackMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteOrganizationConformancePackOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteOrganizationConformancePackInput struct {

	// The name of organization conformance pack that you want to delete.
	//
	// This member is required.
	OrganizationConformancePackName *string

	noSmithyDocumentSerde
}

type DeleteOrganizationConformancePackOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteOrganizationConformancePackMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteOrganizationConformancePack{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteOrganizationConformancePack{}, middleware.After)
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
	if err = addOpDeleteOrganizationConformancePackValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteOrganizationConformancePack(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteOrganizationConformancePack(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "DeleteOrganizationConformancePack",
	}
}
