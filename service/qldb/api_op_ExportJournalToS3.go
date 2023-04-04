// Code generated by smithy-go-codegen DO NOT EDIT.

package qldb

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/qldb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Exports journal contents within a date and time range from a ledger into a
// specified Amazon Simple Storage Service (Amazon S3) bucket. A journal export job
// can write the data objects in either the text or binary representation of Amazon
// Ion format, or in JSON Lines text format. In JSON Lines format, each journal
// block in the exported data object is a valid JSON object that is delimited by a
// newline. You can use this format to easily integrate JSON exports with analytics
// tools such as Glue and Amazon Athena because these services can parse
// newline-delimited JSON automatically. For more information about the format, see
// JSON Lines (https://jsonlines.org/). If the ledger with the given Name doesn't
// exist, then throws ResourceNotFoundException. If the ledger with the given Name
// is in CREATING status, then throws ResourcePreconditionNotMetException. You can
// initiate up to two concurrent journal export requests for each ledger. Beyond
// this limit, journal export requests throw LimitExceededException.
func (c *Client) ExportJournalToS3(ctx context.Context, params *ExportJournalToS3Input, optFns ...func(*Options)) (*ExportJournalToS3Output, error) {
	if params == nil {
		params = &ExportJournalToS3Input{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ExportJournalToS3", params, optFns, c.addOperationExportJournalToS3Middlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ExportJournalToS3Output)
	out.ResultMetadata = metadata
	return out, nil
}

type ExportJournalToS3Input struct {

	// The exclusive end date and time for the range of journal contents to export. The
	// ExclusiveEndTime must be in ISO 8601 date and time format and in Universal
	// Coordinated Time (UTC). For example: 2019-06-13T21:36:34Z. The ExclusiveEndTime
	// must be less than or equal to the current UTC date and time.
	//
	// This member is required.
	ExclusiveEndTime *time.Time

	// The inclusive start date and time for the range of journal contents to export.
	// The InclusiveStartTime must be in ISO 8601 date and time format and in Universal
	// Coordinated Time (UTC). For example: 2019-06-13T21:36:34Z. The
	// InclusiveStartTime must be before ExclusiveEndTime. If you provide an
	// InclusiveStartTime that is before the ledger's CreationDateTime, Amazon QLDB
	// defaults it to the ledger's CreationDateTime.
	//
	// This member is required.
	InclusiveStartTime *time.Time

	// The name of the ledger.
	//
	// This member is required.
	Name *string

	// The Amazon Resource Name (ARN) of the IAM role that grants QLDB permissions for
	// a journal export job to do the following:
	//
	// * Write objects into your Amazon
	// Simple Storage Service (Amazon S3) bucket.
	//
	// * (Optional) Use your customer
	// managed key in Key Management Service (KMS) for server-side encryption of your
	// exported data.
	//
	// To pass a role to QLDB when requesting a journal export, you
	// must have permissions to perform the iam:PassRole action on the IAM role
	// resource. This is required for all journal export requests.
	//
	// This member is required.
	RoleArn *string

	// The configuration settings of the Amazon S3 bucket destination for your export
	// request.
	//
	// This member is required.
	S3ExportConfiguration *types.S3ExportConfiguration

	// The output format of your exported journal data. If this parameter is not
	// specified, the exported data defaults to ION_TEXT format.
	OutputFormat types.OutputFormat

	noSmithyDocumentSerde
}

type ExportJournalToS3Output struct {

	// The UUID (represented in Base62-encoded text) that QLDB assigns to each journal
	// export job. To describe your export request and check the status of the job, you
	// can use ExportId to call DescribeJournalS3Export.
	//
	// This member is required.
	ExportId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationExportJournalToS3Middlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpExportJournalToS3{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpExportJournalToS3{}, middleware.After)
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
	if err = addOpExportJournalToS3ValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opExportJournalToS3(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opExportJournalToS3(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "qldb",
		OperationName: "ExportJournalToS3",
	}
}
