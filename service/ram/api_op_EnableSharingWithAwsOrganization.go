// Code generated by smithy-go-codegen DO NOT EDIT.

package ram

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Enables resource sharing within your organization in Organizations. Calling this
// operation enables RAM to retrieve information about the organization and its
// structure. This lets you share resources with all of the accounts in an
// organization by specifying the organization's ID, or all of the accounts in an
// organizational unit (OU) by specifying the OU's ID. Until you enable sharing
// within the organization, you can specify only individual Amazon Web Services
// accounts, or for supported resource types, IAM users and roles. You must call
// this operation from an IAM user or role in the organization's management
// account.
func (c *Client) EnableSharingWithAwsOrganization(ctx context.Context, params *EnableSharingWithAwsOrganizationInput, optFns ...func(*Options)) (*EnableSharingWithAwsOrganizationOutput, error) {
	if params == nil {
		params = &EnableSharingWithAwsOrganizationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "EnableSharingWithAwsOrganization", params, optFns, c.addOperationEnableSharingWithAwsOrganizationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*EnableSharingWithAwsOrganizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EnableSharingWithAwsOrganizationInput struct {
	noSmithyDocumentSerde
}

type EnableSharingWithAwsOrganizationOutput struct {

	// A return value of true indicates that the request succeeded. A value of false
	// indicates that the request failed.
	ReturnValue *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationEnableSharingWithAwsOrganizationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpEnableSharingWithAwsOrganization{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpEnableSharingWithAwsOrganization{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opEnableSharingWithAwsOrganization(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opEnableSharingWithAwsOrganization(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ram",
		OperationName: "EnableSharingWithAwsOrganization",
	}
}
