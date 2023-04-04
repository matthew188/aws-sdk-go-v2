// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a grant. Typically, you retire a grant when you no longer need its
// permissions. To identify the grant to retire, use a grant token
// (https://docs.aws.amazon.com/kms/latest/developerguide/grants.html#grant_token),
// or both the grant ID and a key identifier (key ID or key ARN) of the KMS key.
// The CreateGrant operation returns both values. This operation can be called by
// the retiring principal for a grant, by the grantee principal if the grant allows
// the RetireGrant operation, and by the Amazon Web Services account in which the
// grant is created. It can also be called by principals to whom permission for
// retiring a grant is delegated. For details, see Retiring and revoking grants
// (https://docs.aws.amazon.com/kms/latest/developerguide/grant-manage.html#grant-delete)
// in the Key Management Service Developer Guide. For detailed information about
// grants, including grant terminology, see Grants in KMS
// (https://docs.aws.amazon.com/kms/latest/developerguide/grants.html) in the Key
// Management Service Developer Guide . For examples of working with grants in
// several programming languages, see Programming grants
// (https://docs.aws.amazon.com/kms/latest/developerguide/programming-grants.html).
// Cross-account use: Yes. You can retire a grant on a KMS key in a different
// Amazon Web Services account. Required permissions::Permission to retire a grant
// is determined primarily by the grant. For details, see Retiring and revoking
// grants
// (https://docs.aws.amazon.com/kms/latest/developerguide/grant-manage.html#grant-delete)
// in the Key Management Service Developer Guide. Related operations:
//
// *
// CreateGrant
//
// * ListGrants
//
// * ListRetirableGrants
//
// * RevokeGrant
func (c *Client) RetireGrant(ctx context.Context, params *RetireGrantInput, optFns ...func(*Options)) (*RetireGrantOutput, error) {
	if params == nil {
		params = &RetireGrantInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RetireGrant", params, optFns, c.addOperationRetireGrantMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RetireGrantOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RetireGrantInput struct {

	// Identifies the grant to retire. To get the grant ID, use CreateGrant,
	// ListGrants, or ListRetirableGrants.
	//
	// * Grant ID Example -
	// 0123456789012345678901234567890123456789012345678901234567890123
	GrantId *string

	// Identifies the grant to be retired. You can use a grant token to identify a new
	// grant even before it has achieved eventual consistency. Only the CreateGrant
	// operation returns a grant token. For details, see Grant token
	// (https://docs.aws.amazon.com/kms/latest/developerguide/grants.html#grant_token)
	// and Eventual consistency
	// (https://docs.aws.amazon.com/kms/latest/developerguide/grants.html#terms-eventual-consistency)
	// in the Key Management Service Developer Guide.
	GrantToken *string

	// The key ARN KMS key associated with the grant. To find the key ARN, use the
	// ListKeys operation. For example:
	// arn:aws:kms:us-east-2:444455556666:key/1234abcd-12ab-34cd-56ef-1234567890ab
	KeyId *string

	noSmithyDocumentSerde
}

type RetireGrantOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRetireGrantMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRetireGrant{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRetireGrant{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRetireGrant(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRetireGrant(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kms",
		OperationName: "RetireGrant",
	}
}
