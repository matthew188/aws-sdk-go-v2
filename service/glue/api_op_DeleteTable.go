// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Removes a table definition from the Data Catalog. After completing this
// operation, you no longer have access to the table versions and partitions that
// belong to the deleted table. Glue deletes these "orphaned" resources
// asynchronously in a timely manner, at the discretion of the service. To ensure
// the immediate deletion of all related resources, before calling DeleteTable, use
// DeleteTableVersion or BatchDeleteTableVersion, and DeletePartition or
// BatchDeletePartition, to delete any resources that belong to the table.
func (c *Client) DeleteTable(ctx context.Context, params *DeleteTableInput, optFns ...func(*Options)) (*DeleteTableOutput, error) {
	if params == nil {
		params = &DeleteTableInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteTable", params, optFns, c.addOperationDeleteTableMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteTableOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteTableInput struct {

	// The name of the catalog database in which the table resides. For Hive
	// compatibility, this name is entirely lowercase.
	//
	// This member is required.
	DatabaseName *string

	// The name of the table to be deleted. For Hive compatibility, this name is
	// entirely lowercase.
	//
	// This member is required.
	Name *string

	// The ID of the Data Catalog where the table resides. If none is provided, the
	// Amazon Web Services account ID is used by default.
	CatalogId *string

	// The transaction ID at which to delete the table contents.
	TransactionId *string

	noSmithyDocumentSerde
}

type DeleteTableOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteTableMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteTable{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteTable{}, middleware.After)
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
	if err = addOpDeleteTableValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteTable(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteTable(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "DeleteTable",
	}
}
