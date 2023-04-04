// Code generated by smithy-go-codegen DO NOT EDIT.

package iotwireless

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/iotwireless/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Disassociates your AWS account from a partner account. If PartnerAccountId and
// PartnerType are null, disassociates your AWS account from all partner accounts.
func (c *Client) DisassociateAwsAccountFromPartnerAccount(ctx context.Context, params *DisassociateAwsAccountFromPartnerAccountInput, optFns ...func(*Options)) (*DisassociateAwsAccountFromPartnerAccountOutput, error) {
	if params == nil {
		params = &DisassociateAwsAccountFromPartnerAccountInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociateAwsAccountFromPartnerAccount", params, optFns, c.addOperationDisassociateAwsAccountFromPartnerAccountMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociateAwsAccountFromPartnerAccountOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateAwsAccountFromPartnerAccountInput struct {

	// The partner account ID to disassociate from the AWS account.
	//
	// This member is required.
	PartnerAccountId *string

	// The partner type.
	//
	// This member is required.
	PartnerType types.PartnerType

	noSmithyDocumentSerde
}

type DisassociateAwsAccountFromPartnerAccountOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDisassociateAwsAccountFromPartnerAccountMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDisassociateAwsAccountFromPartnerAccount{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDisassociateAwsAccountFromPartnerAccount{}, middleware.After)
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
	if err = addOpDisassociateAwsAccountFromPartnerAccountValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateAwsAccountFromPartnerAccount(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisassociateAwsAccountFromPartnerAccount(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotwireless",
		OperationName: "DisassociateAwsAccountFromPartnerAccount",
	}
}
