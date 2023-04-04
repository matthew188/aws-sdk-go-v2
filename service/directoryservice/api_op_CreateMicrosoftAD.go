// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservice

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/directoryservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a Microsoft AD directory in the Amazon Web Services Cloud. For more
// information, see Managed Microsoft AD
// (https://docs.aws.amazon.com/directoryservice/latest/admin-guide/directory_microsoft_ad.html)
// in the Directory Service Admin Guide. Before you call CreateMicrosoftAD, ensure
// that all of the required permissions have been explicitly granted through a
// policy. For details about what permissions are required to run the
// CreateMicrosoftAD operation, see Directory Service API Permissions: Actions,
// Resources, and Conditions Reference
// (http://docs.aws.amazon.com/directoryservice/latest/admin-guide/UsingWithDS_IAM_ResourcePermissions.html).
func (c *Client) CreateMicrosoftAD(ctx context.Context, params *CreateMicrosoftADInput, optFns ...func(*Options)) (*CreateMicrosoftADOutput, error) {
	if params == nil {
		params = &CreateMicrosoftADInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateMicrosoftAD", params, optFns, c.addOperationCreateMicrosoftADMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateMicrosoftADOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Creates an Managed Microsoft AD directory.
type CreateMicrosoftADInput struct {

	// The fully qualified domain name for the Managed Microsoft AD directory, such as
	// corp.example.com. This name will resolve inside your VPC only. It does not need
	// to be publicly resolvable.
	//
	// This member is required.
	Name *string

	// The password for the default administrative user named Admin. If you need to
	// change the password for the administrator account, you can use the
	// ResetUserPassword API call.
	//
	// This member is required.
	Password *string

	// Contains VPC information for the CreateDirectory or CreateMicrosoftAD operation.
	//
	// This member is required.
	VpcSettings *types.DirectoryVpcSettings

	// A description for the directory. This label will appear on the Amazon Web
	// Services console Directory Details page after the directory is created.
	Description *string

	// Managed Microsoft AD is available in two editions: Standard and Enterprise.
	// Enterprise is the default.
	Edition types.DirectoryEdition

	// The NetBIOS name for your domain, such as CORP. If you don't specify a NetBIOS
	// name, it will default to the first part of your directory DNS. For example, CORP
	// for the directory DNS corp.example.com.
	ShortName *string

	// The tags to be assigned to the Managed Microsoft AD directory.
	Tags []types.Tag

	noSmithyDocumentSerde
}

// Result of a CreateMicrosoftAD request.
type CreateMicrosoftADOutput struct {

	// The identifier of the directory that was created.
	DirectoryId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateMicrosoftADMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateMicrosoftAD{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateMicrosoftAD{}, middleware.After)
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
	if err = addOpCreateMicrosoftADValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateMicrosoftAD(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateMicrosoftAD(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ds",
		OperationName: "CreateMicrosoftAD",
	}
}
