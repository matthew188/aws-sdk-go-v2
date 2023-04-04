// Code generated by smithy-go-codegen DO NOT EDIT.

package managedblockchain

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/managedblockchain/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new blockchain network using Amazon Managed Blockchain. Applies only
// to Hyperledger Fabric.
func (c *Client) CreateNetwork(ctx context.Context, params *CreateNetworkInput, optFns ...func(*Options)) (*CreateNetworkOutput, error) {
	if params == nil {
		params = &CreateNetworkInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateNetwork", params, optFns, c.addOperationCreateNetworkMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateNetworkOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateNetworkInput struct {

	// This is a unique, case-sensitive identifier that you provide to ensure the
	// idempotency of the operation. An idempotent operation completes no more than
	// once. This identifier is required only if you make a service request directly
	// using an HTTP client. It is generated automatically if you use an Amazon Web
	// Services SDK or the Amazon Web Services CLI.
	//
	// This member is required.
	ClientRequestToken *string

	// The blockchain framework that the network uses.
	//
	// This member is required.
	Framework types.Framework

	// The version of the blockchain framework that the network uses.
	//
	// This member is required.
	FrameworkVersion *string

	// Configuration properties for the first member within the network.
	//
	// This member is required.
	MemberConfiguration *types.MemberConfiguration

	// The name of the network.
	//
	// This member is required.
	Name *string

	// The voting rules used by the network to determine if a proposal is approved.
	//
	// This member is required.
	VotingPolicy *types.VotingPolicy

	// An optional description for the network.
	Description *string

	// Configuration properties of the blockchain framework relevant to the network
	// configuration.
	FrameworkConfiguration *types.NetworkFrameworkConfiguration

	// Tags to assign to the network. Each tag consists of a key and an optional value.
	// You can specify multiple key-value pairs in a single request with an overall
	// maximum of 50 tags allowed per resource. For more information about tags, see
	// Tagging Resources
	// (https://docs.aws.amazon.com/managed-blockchain/latest/ethereum-dev/tagging-resources.html)
	// in the Amazon Managed Blockchain Ethereum Developer Guide, or Tagging Resources
	// (https://docs.aws.amazon.com/managed-blockchain/latest/hyperledger-fabric-dev/tagging-resources.html)
	// in the Amazon Managed Blockchain Hyperledger Fabric Developer Guide.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateNetworkOutput struct {

	// The unique identifier for the first member within the network.
	MemberId *string

	// The unique identifier for the network.
	NetworkId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateNetworkMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateNetwork{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateNetwork{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateNetworkMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateNetworkValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateNetwork(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateNetwork struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateNetwork) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateNetwork) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateNetworkInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateNetworkInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateNetworkMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateNetwork{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateNetwork(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "managedblockchain",
		OperationName: "CreateNetwork",
	}
}
