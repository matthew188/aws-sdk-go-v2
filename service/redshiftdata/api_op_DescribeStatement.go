// Code generated by smithy-go-codegen DO NOT EDIT.

package redshiftdata

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/redshiftdata/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Describes the details about a specific instance when a query was run by the
// Amazon Redshift Data API. The information includes when the query started, when
// it finished, the query status, the number of rows returned, and the SQL
// statement. For more information about the Amazon Redshift Data API and CLI usage
// examples, see Using the Amazon Redshift Data API
// (https://docs.aws.amazon.com/redshift/latest/mgmt/data-api.html) in the Amazon
// Redshift Management Guide.
func (c *Client) DescribeStatement(ctx context.Context, params *DescribeStatementInput, optFns ...func(*Options)) (*DescribeStatementOutput, error) {
	if params == nil {
		params = &DescribeStatementInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeStatement", params, optFns, c.addOperationDescribeStatementMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeStatementOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeStatementInput struct {

	// The identifier of the SQL statement to describe. This value is a universally
	// unique identifier (UUID) generated by Amazon Redshift Data API. A suffix
	// indicates the number of the SQL statement. For example,
	// d9b6c0c9-0747-4bf4-b142-e8883122f766:2 has a suffix of :2 that indicates the
	// second SQL statement of a batch query. This identifier is returned by
	// BatchExecuteStatment, ExecuteStatement, and ListStatements.
	//
	// This member is required.
	Id *string

	noSmithyDocumentSerde
}

type DescribeStatementOutput struct {

	// The identifier of the SQL statement described. This value is a universally
	// unique identifier (UUID) generated by Amazon Redshift Data API.
	//
	// This member is required.
	Id *string

	// The cluster identifier.
	ClusterIdentifier *string

	// The date and time (UTC) when the SQL statement was submitted to run.
	CreatedAt *time.Time

	// The name of the database.
	Database *string

	// The database user name.
	DbUser *string

	// The amount of time in nanoseconds that the statement ran.
	Duration int64

	// The error message from the cluster if the SQL statement encountered an error
	// while running.
	Error *string

	// A value that indicates whether the statement has a result set. The result set
	// can be empty. The value is true for an empty result set. The value is true if
	// any substatement returns a result set.
	HasResultSet *bool

	// The parameters for the SQL statement.
	QueryParameters []types.SqlParameter

	// The SQL statement text.
	QueryString *string

	// The process identifier from Amazon Redshift.
	RedshiftPid int64

	// The identifier of the query generated by Amazon Redshift. These identifiers are
	// also available in the query column of the STL_QUERY system view.
	RedshiftQueryId int64

	// Either the number of rows returned from the SQL statement or the number of rows
	// affected. If result size is greater than zero, the result rows can be the number
	// of rows affected by SQL statements such as INSERT, UPDATE, DELETE, COPY, and
	// others. A -1 indicates the value is null.
	ResultRows int64

	// The size in bytes of the returned results. A -1 indicates the value is null.
	ResultSize int64

	// The name or Amazon Resource Name (ARN) of the secret that enables access to the
	// database.
	SecretArn *string

	// The status of the SQL statement being described. Status values are defined as
	// follows:
	//
	// * ABORTED - The query run was stopped by the user.
	//
	// * ALL - A status
	// value that includes all query statuses. This value can be used to filter
	// results.
	//
	// * FAILED - The query run failed.
	//
	// * FINISHED - The query has finished
	// running.
	//
	// * PICKED - The query has been chosen to be run.
	//
	// * STARTED - The query
	// run has started.
	//
	// * SUBMITTED - The query was submitted, but not yet processed.
	Status types.StatusString

	// The SQL statements from a multiple statement run.
	SubStatements []types.SubStatementData

	// The date and time (UTC) that the metadata for the SQL statement was last
	// updated. An example is the time the status last changed.
	UpdatedAt *time.Time

	// The serverless workgroup name or Amazon Resource Name (ARN).
	WorkgroupName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeStatementMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeStatement{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeStatement{}, middleware.After)
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
	if err = addOpDescribeStatementValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeStatement(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeStatement(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift-data",
		OperationName: "DescribeStatement",
	}
}
