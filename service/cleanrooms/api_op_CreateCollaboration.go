// Code generated by smithy-go-codegen DO NOT EDIT.

package cleanrooms

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/cleanrooms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new collaboration.
func (c *Client) CreateCollaboration(ctx context.Context, params *CreateCollaborationInput, optFns ...func(*Options)) (*CreateCollaborationOutput, error) {
	if params == nil {
		params = &CreateCollaborationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCollaboration", params, optFns, c.addOperationCreateCollaborationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateCollaborationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateCollaborationInput struct {

	// The display name of the collaboration creator.
	//
	// This member is required.
	CreatorDisplayName *string

	// The abilities granted to the collaboration creator.
	//
	// This member is required.
	CreatorMemberAbilities []types.MemberAbility

	// A description of the collaboration provided by the collaboration owner.
	//
	// This member is required.
	Description *string

	// A list of initial members, not including the creator. This list is immutable.
	//
	// This member is required.
	Members []types.MemberSpecification

	// The display name for a collaboration.
	//
	// This member is required.
	Name *string

	// An indicator as to whether query logging has been enabled or disabled for the
	// collaboration.
	//
	// This member is required.
	QueryLogStatus types.CollaborationQueryLogStatus

	// The settings for client-side encryption with Cryptographic Computing for Clean
	// Rooms.
	DataEncryptionMetadata *types.DataEncryptionMetadata

	// An optional label that you can assign to a resource when you create it. Each tag
	// consists of a key and an optional value, both of which you define. When you use
	// tagging, you can also use tag-based access control in IAM policies to control
	// access to this resource.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateCollaborationOutput struct {

	// The entire created collaboration object.
	//
	// This member is required.
	Collaboration *types.Collaboration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateCollaborationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateCollaboration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateCollaboration{}, middleware.After)
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
	if err = addOpCreateCollaborationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCollaboration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateCollaboration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cleanrooms",
		OperationName: "CreateCollaboration",
	}
}
