// Code generated by smithy-go-codegen DO NOT EDIT.

package ecrpublic

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/ecrpublic/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates or updates the catalog data for a repository in a public registry.
func (c *Client) PutRepositoryCatalogData(ctx context.Context, params *PutRepositoryCatalogDataInput, optFns ...func(*Options)) (*PutRepositoryCatalogDataOutput, error) {
	if params == nil {
		params = &PutRepositoryCatalogDataInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutRepositoryCatalogData", params, optFns, c.addOperationPutRepositoryCatalogDataMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutRepositoryCatalogDataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutRepositoryCatalogDataInput struct {

	// An object containing the catalog data for a repository. This data is publicly
	// visible in the Amazon ECR Public Gallery.
	//
	// This member is required.
	CatalogData *types.RepositoryCatalogDataInput

	// The name of the repository to create or update the catalog data for.
	//
	// This member is required.
	RepositoryName *string

	// The Amazon Web Services account ID that's associated with the public registry
	// the repository is in. If you do not specify a registry, the default public
	// registry is assumed.
	RegistryId *string

	noSmithyDocumentSerde
}

type PutRepositoryCatalogDataOutput struct {

	// The catalog data for the repository.
	CatalogData *types.RepositoryCatalogData

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutRepositoryCatalogDataMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutRepositoryCatalogData{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutRepositoryCatalogData{}, middleware.After)
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
	if err = addOpPutRepositoryCatalogDataValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutRepositoryCatalogData(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutRepositoryCatalogData(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecr-public",
		OperationName: "PutRepositoryCatalogData",
	}
}
