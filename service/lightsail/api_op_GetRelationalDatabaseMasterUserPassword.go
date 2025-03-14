// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/lightsail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns the current, previous, or pending versions of the master user password
// for a Lightsail database. The GetRelationalDatabaseMasterUserPassword operation
// supports tag-based access control via resource tags applied to the resource
// identified by relationalDatabaseName.
func (c *Client) GetRelationalDatabaseMasterUserPassword(ctx context.Context, params *GetRelationalDatabaseMasterUserPasswordInput, optFns ...func(*Options)) (*GetRelationalDatabaseMasterUserPasswordOutput, error) {
	if params == nil {
		params = &GetRelationalDatabaseMasterUserPasswordInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetRelationalDatabaseMasterUserPassword", params, optFns, c.addOperationGetRelationalDatabaseMasterUserPasswordMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetRelationalDatabaseMasterUserPasswordOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRelationalDatabaseMasterUserPasswordInput struct {

	// The name of your database for which to get the master user password.
	//
	// This member is required.
	RelationalDatabaseName *string

	// The password version to return. Specifying CURRENT or PREVIOUS returns the
	// current or previous passwords respectively. Specifying PENDING returns the
	// newest version of the password that will rotate to CURRENT. After the PENDING
	// password rotates to CURRENT, the PENDING password is no longer available.
	// Default: CURRENT
	PasswordVersion types.RelationalDatabasePasswordVersion

	noSmithyDocumentSerde
}

type GetRelationalDatabaseMasterUserPasswordOutput struct {

	// The timestamp when the specified version of the master user password was
	// created.
	CreatedAt *time.Time

	// The master user password for the password version specified.
	MasterUserPassword *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetRelationalDatabaseMasterUserPasswordMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetRelationalDatabaseMasterUserPassword{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetRelationalDatabaseMasterUserPassword{}, middleware.After)
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
	if err = addOpGetRelationalDatabaseMasterUserPasswordValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetRelationalDatabaseMasterUserPassword(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetRelationalDatabaseMasterUserPassword(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "GetRelationalDatabaseMasterUserPassword",
	}
}
