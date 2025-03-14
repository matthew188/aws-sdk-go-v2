// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/glue/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a specified batch of versions of a table.
func (c *Client) BatchDeleteTableVersion(ctx context.Context, params *BatchDeleteTableVersionInput, optFns ...func(*Options)) (*BatchDeleteTableVersionOutput, error) {
	if params == nil {
		params = &BatchDeleteTableVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchDeleteTableVersion", params, optFns, c.addOperationBatchDeleteTableVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchDeleteTableVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchDeleteTableVersionInput struct {

	// The database in the catalog in which the table resides. For Hive compatibility,
	// this name is entirely lowercase.
	//
	// This member is required.
	DatabaseName *string

	// The name of the table. For Hive compatibility, this name is entirely lowercase.
	//
	// This member is required.
	TableName *string

	// A list of the IDs of versions to be deleted. A VersionId is a string
	// representation of an integer. Each version is incremented by 1.
	//
	// This member is required.
	VersionIds []string

	// The ID of the Data Catalog where the tables reside. If none is provided, the
	// Amazon Web Services account ID is used by default.
	CatalogId *string

	noSmithyDocumentSerde
}

type BatchDeleteTableVersionOutput struct {

	// A list of errors encountered while trying to delete the specified table
	// versions.
	Errors []types.TableVersionError

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchDeleteTableVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpBatchDeleteTableVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchDeleteTableVersion{}, middleware.After)
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
	if err = addOpBatchDeleteTableVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchDeleteTableVersion(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchDeleteTableVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "BatchDeleteTableVersion",
	}
}
