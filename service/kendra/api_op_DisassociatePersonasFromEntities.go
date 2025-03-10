// Code generated by smithy-go-codegen DO NOT EDIT.

package kendra

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/kendra/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Removes the specific permissions of users or groups in your IAM Identity Center
// identity source with access to your Amazon Kendra experience. You can create an
// Amazon Kendra experience such as a search application. For more information on
// creating a search application experience, see Building a search experience with
// no code
// (https://docs.aws.amazon.com/kendra/latest/dg/deploying-search-experience-no-code.html).
func (c *Client) DisassociatePersonasFromEntities(ctx context.Context, params *DisassociatePersonasFromEntitiesInput, optFns ...func(*Options)) (*DisassociatePersonasFromEntitiesOutput, error) {
	if params == nil {
		params = &DisassociatePersonasFromEntitiesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociatePersonasFromEntities", params, optFns, c.addOperationDisassociatePersonasFromEntitiesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociatePersonasFromEntitiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociatePersonasFromEntitiesInput struct {

	// The identifiers of users or groups in your IAM Identity Center identity source.
	// For example, user IDs could be user emails.
	//
	// This member is required.
	EntityIds []string

	// The identifier of your Amazon Kendra experience.
	//
	// This member is required.
	Id *string

	// The identifier of the index for your Amazon Kendra experience.
	//
	// This member is required.
	IndexId *string

	noSmithyDocumentSerde
}

type DisassociatePersonasFromEntitiesOutput struct {

	// Lists the users or groups in your IAM Identity Center identity source that
	// failed to properly remove access to your Amazon Kendra experience.
	FailedEntityList []types.FailedEntity

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDisassociatePersonasFromEntitiesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDisassociatePersonasFromEntities{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDisassociatePersonasFromEntities{}, middleware.After)
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
	if err = addOpDisassociatePersonasFromEntitiesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociatePersonasFromEntities(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisassociatePersonasFromEntities(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kendra",
		OperationName: "DisassociatePersonasFromEntities",
	}
}
