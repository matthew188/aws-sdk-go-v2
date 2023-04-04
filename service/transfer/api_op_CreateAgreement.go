// Code generated by smithy-go-codegen DO NOT EDIT.

package transfer

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/transfer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an agreement. An agreement is a bilateral trading partner agreement, or
// partnership, between an Transfer Family server and an AS2 process. The agreement
// defines the file and message transfer relationship between the server and the
// AS2 process. To define an agreement, Transfer Family combines a server, local
// profile, partner profile, certificate, and other attributes. The partner is
// identified with the PartnerProfileId, and the AS2 process is identified with the
// LocalProfileId.
func (c *Client) CreateAgreement(ctx context.Context, params *CreateAgreementInput, optFns ...func(*Options)) (*CreateAgreementOutput, error) {
	if params == nil {
		params = &CreateAgreementInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAgreement", params, optFns, c.addOperationCreateAgreementMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAgreementOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAgreementInput struct {

	// With AS2, you can send files by calling StartFileTransfer and specifying the
	// file paths in the request parameter, SendFilePaths. We use the file’s parent
	// directory (for example, for --send-file-paths /bucket/dir/file.txt, parent
	// directory is /bucket/dir/) to temporarily store a processed AS2 message file,
	// store the MDN when we receive them from the partner, and write a final JSON file
	// containing relevant metadata of the transmission. So, the AccessRole needs to
	// provide read and write access to the parent directory of the file location used
	// in the StartFileTransfer request. Additionally, you need to provide read and
	// write access to the parent directory of the files that you intend to send with
	// StartFileTransfer.
	//
	// This member is required.
	AccessRole *string

	// The landing directory (folder) for files transferred by using the AS2 protocol.
	// A BaseDirectory example is /DOC-EXAMPLE-BUCKET/home/mydirectory.
	//
	// This member is required.
	BaseDirectory *string

	// A unique identifier for the AS2 local profile.
	//
	// This member is required.
	LocalProfileId *string

	// A unique identifier for the partner profile used in the agreement.
	//
	// This member is required.
	PartnerProfileId *string

	// A system-assigned unique identifier for a server instance. This is the specific
	// server that the agreement uses.
	//
	// This member is required.
	ServerId *string

	// A name or short description to identify the agreement.
	Description *string

	// The status of the agreement. The agreement can be either ACTIVE or INACTIVE.
	Status types.AgreementStatusType

	// Key-value pairs that can be used to group and search for agreements.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateAgreementOutput struct {

	// The unique identifier for the agreement. Use this ID for deleting, or updating
	// an agreement, as well as in any other API calls that require that you specify
	// the agreement ID.
	//
	// This member is required.
	AgreementId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAgreementMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateAgreement{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateAgreement{}, middleware.After)
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
	if err = addOpCreateAgreementValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAgreement(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateAgreement(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transfer",
		OperationName: "CreateAgreement",
	}
}
