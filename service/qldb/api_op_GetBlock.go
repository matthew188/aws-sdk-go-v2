// Code generated by smithy-go-codegen DO NOT EDIT.

package qldb

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/qldb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a block object at a specified address in a journal. Also returns a proof
// of the specified block for verification if DigestTipAddress is provided. For
// information about the data contents in a block, see Journal contents
// (https://docs.aws.amazon.com/qldb/latest/developerguide/journal-contents.html)
// in the Amazon QLDB Developer Guide. If the specified ledger doesn't exist or is
// in DELETING status, then throws ResourceNotFoundException. If the specified
// ledger is in CREATING status, then throws ResourcePreconditionNotMetException.
// If no block exists with the specified address, then throws
// InvalidParameterException.
func (c *Client) GetBlock(ctx context.Context, params *GetBlockInput, optFns ...func(*Options)) (*GetBlockOutput, error) {
	if params == nil {
		params = &GetBlockInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetBlock", params, optFns, c.addOperationGetBlockMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetBlockOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBlockInput struct {

	// The location of the block that you want to request. An address is an Amazon Ion
	// structure that has two fields: strandId and sequenceNo. For example:
	// {strandId:"BlFTjlSXze9BIh1KOszcE3",sequenceNo:14}.
	//
	// This member is required.
	BlockAddress *types.ValueHolder

	// The name of the ledger.
	//
	// This member is required.
	Name *string

	// The latest block location covered by the digest for which to request a proof. An
	// address is an Amazon Ion structure that has two fields: strandId and sequenceNo.
	// For example: {strandId:"BlFTjlSXze9BIh1KOszcE3",sequenceNo:49}.
	DigestTipAddress *types.ValueHolder

	noSmithyDocumentSerde
}

type GetBlockOutput struct {

	// The block data object in Amazon Ion format.
	//
	// This member is required.
	Block *types.ValueHolder

	// The proof object in Amazon Ion format returned by a GetBlock request. A proof
	// contains the list of hash values required to recalculate the specified digest
	// using a Merkle tree, starting with the specified block.
	Proof *types.ValueHolder

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetBlockMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetBlock{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetBlock{}, middleware.After)
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
	if err = addOpGetBlockValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetBlock(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetBlock(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "qldb",
		OperationName: "GetBlock",
	}
}
