// Code generated by smithy-go-codegen DO NOT EDIT.

package honeycode

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/honeycode/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// The BatchDeleteTableRows API allows you to delete one or more rows from a table
// in a workbook. You need to specify the ids of the rows that you want to delete
// from the table.
func (c *Client) BatchDeleteTableRows(ctx context.Context, params *BatchDeleteTableRowsInput, optFns ...func(*Options)) (*BatchDeleteTableRowsOutput, error) {
	if params == nil {
		params = &BatchDeleteTableRowsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchDeleteTableRows", params, optFns, c.addOperationBatchDeleteTableRowsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchDeleteTableRowsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchDeleteTableRowsInput struct {

	// The list of row ids to delete from the table. You need to specify at least one
	// row id in this list. Note that if one of the row ids provided in the request
	// does not exist in the table, then the request fails and no rows are deleted from
	// the table.
	//
	// This member is required.
	RowIds []string

	// The ID of the table where the rows are being deleted. If a table with the
	// specified id could not be found, this API throws ResourceNotFoundException.
	//
	// This member is required.
	TableId *string

	// The ID of the workbook where the rows are being deleted. If a workbook with the
	// specified id could not be found, this API throws ResourceNotFoundException.
	//
	// This member is required.
	WorkbookId *string

	// The request token for performing the delete action. Request tokens help to
	// identify duplicate requests. If a call times out or fails due to a transient
	// error like a failed network connection, you can retry the call with the same
	// request token. The service ensures that if the first call using that request
	// token is successfully performed, the second call will not perform the action
	// again. Note that request tokens are valid only for a few minutes. You cannot use
	// request tokens to dedupe requests spanning hours or days.
	ClientRequestToken *string

	noSmithyDocumentSerde
}

type BatchDeleteTableRowsOutput struct {

	// The updated workbook cursor after deleting the rows from the table.
	//
	// This member is required.
	WorkbookCursor int64

	// The list of row ids in the request that could not be deleted from the table.
	// Each element in this list contains one row id from the request that could not be
	// deleted along with the reason why that item could not be deleted.
	FailedBatchItems []types.FailedBatchItem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchDeleteTableRowsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchDeleteTableRows{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchDeleteTableRows{}, middleware.After)
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
	if err = addOpBatchDeleteTableRowsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchDeleteTableRows(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchDeleteTableRows(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "honeycode",
		OperationName: "BatchDeleteTableRows",
	}
}
