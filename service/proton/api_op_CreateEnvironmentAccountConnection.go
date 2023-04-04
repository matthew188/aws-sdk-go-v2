// Code generated by smithy-go-codegen DO NOT EDIT.

package proton

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/proton/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Create an environment account connection in an environment account so that
// environment infrastructure resources can be provisioned in the environment
// account from a management account. An environment account connection is a secure
// bi-directional connection between a management account and an environment
// account that maintains authorization and permissions. For more information, see
// Environment account connections
// (https://docs.aws.amazon.com/proton/latest/userguide/ag-env-account-connections.html)
// in the Proton User guide.
func (c *Client) CreateEnvironmentAccountConnection(ctx context.Context, params *CreateEnvironmentAccountConnectionInput, optFns ...func(*Options)) (*CreateEnvironmentAccountConnectionOutput, error) {
	if params == nil {
		params = &CreateEnvironmentAccountConnectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateEnvironmentAccountConnection", params, optFns, c.addOperationCreateEnvironmentAccountConnectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateEnvironmentAccountConnectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateEnvironmentAccountConnectionInput struct {

	// The name of the Proton environment that's created in the associated management
	// account.
	//
	// This member is required.
	EnvironmentName *string

	// The ID of the management account that accepts or rejects the environment account
	// connection. You create and manage the Proton environment in this account. If the
	// management account accepts the environment account connection, Proton can use
	// the associated IAM role to provision environment infrastructure resources in the
	// associated environment account.
	//
	// This member is required.
	ManagementAccountId *string

	// When included, if two identical requests are made with the same client token,
	// Proton returns the environment account connection that the first request
	// created.
	ClientToken *string

	// The Amazon Resource Name (ARN) of an IAM service role in the environment
	// account. Proton uses this role to provision infrastructure resources using
	// CodeBuild-based provisioning in the associated environment account.
	CodebuildRoleArn *string

	// The Amazon Resource Name (ARN) of the IAM service role that Proton uses when
	// provisioning directly defined components in the associated environment account.
	// It determines the scope of infrastructure that a component can provision in the
	// account. You must specify componentRoleArn to allow directly defined components
	// to be associated with any environments running in this account. For more
	// information about components, see Proton components
	// (https://docs.aws.amazon.com/proton/latest/userguide/ag-components.html) in the
	// Proton User Guide.
	ComponentRoleArn *string

	// The Amazon Resource Name (ARN) of the IAM service role that's created in the
	// environment account. Proton uses this role to provision infrastructure resources
	// in the associated environment account.
	RoleArn *string

	// An optional list of metadata items that you can associate with the Proton
	// environment account connection. A tag is a key-value pair. For more information,
	// see Proton resources and tagging
	// (https://docs.aws.amazon.com/proton/latest/userguide/resources.html) in the
	// Proton User Guide.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateEnvironmentAccountConnectionOutput struct {

	// The environment account connection detail data that's returned by Proton.
	//
	// This member is required.
	EnvironmentAccountConnection *types.EnvironmentAccountConnection

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateEnvironmentAccountConnectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCreateEnvironmentAccountConnection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCreateEnvironmentAccountConnection{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateEnvironmentAccountConnectionMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateEnvironmentAccountConnectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateEnvironmentAccountConnection(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateEnvironmentAccountConnection struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateEnvironmentAccountConnection) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateEnvironmentAccountConnection) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateEnvironmentAccountConnectionInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateEnvironmentAccountConnectionInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateEnvironmentAccountConnectionMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateEnvironmentAccountConnection{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateEnvironmentAccountConnection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "proton",
		OperationName: "CreateEnvironmentAccountConnection",
	}
}
