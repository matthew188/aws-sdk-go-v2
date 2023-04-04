// Code generated by smithy-go-codegen DO NOT EDIT.

package securityhub

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This method is deprecated. Instead, use DisassociateFromAdministratorAccount.
// The Security Hub console continues to use DisassociateFromMasterAccount. It will
// eventually change to use DisassociateFromAdministratorAccount. Any IAM policies
// that specifically control access to this function must continue to use
// DisassociateFromMasterAccount. You should also add
// DisassociateFromAdministratorAccount to your policies to ensure that the correct
// permissions are in place after the console begins to use
// DisassociateFromAdministratorAccount. Disassociates the current Security Hub
// member account from the associated administrator account. This operation is only
// used by accounts that are not part of an organization. For organization
// accounts, only the administrator account can disassociate a member account.
//
// Deprecated: This API has been deprecated, use
// DisassociateFromAdministratorAccount API instead.
func (c *Client) DisassociateFromMasterAccount(ctx context.Context, params *DisassociateFromMasterAccountInput, optFns ...func(*Options)) (*DisassociateFromMasterAccountOutput, error) {
	if params == nil {
		params = &DisassociateFromMasterAccountInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociateFromMasterAccount", params, optFns, c.addOperationDisassociateFromMasterAccountMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociateFromMasterAccountOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateFromMasterAccountInput struct {
	noSmithyDocumentSerde
}

type DisassociateFromMasterAccountOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDisassociateFromMasterAccountMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDisassociateFromMasterAccount{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDisassociateFromMasterAccount{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateFromMasterAccount(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisassociateFromMasterAccount(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securityhub",
		OperationName: "DisassociateFromMasterAccount",
	}
}
