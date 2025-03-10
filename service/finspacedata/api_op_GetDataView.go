// Code generated by smithy-go-codegen DO NOT EDIT.

package finspacedata

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/finspacedata/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets information about a Dataview.
func (c *Client) GetDataView(ctx context.Context, params *GetDataViewInput, optFns ...func(*Options)) (*GetDataViewOutput, error) {
	if params == nil {
		params = &GetDataViewInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDataView", params, optFns, c.addOperationGetDataViewMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDataViewOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request for retrieving a data view detail. Grouped / accessible within a dataset
// by its dataset id.
type GetDataViewInput struct {

	// The unique identifier for the Dataview.
	//
	// This member is required.
	DataViewId *string

	// The unique identifier for the Dataset used in the Dataview.
	//
	// This member is required.
	DatasetId *string

	noSmithyDocumentSerde
}

// Response from retrieving a dataview, which includes details on the target
// database and table name
type GetDataViewOutput struct {

	// Time range to use for the Dataview. The value is determined as epoch time in
	// milliseconds. For example, the value for Monday, November 1, 2021 12:00:00 PM
	// UTC is specified as 1635768000000.
	AsOfTimestamp *int64

	// Flag to indicate Dataview should be updated automatically.
	AutoUpdate bool

	// The timestamp at which the Dataview was created in FinSpace. The value is
	// determined as epoch time in milliseconds. For example, the value for Monday,
	// November 1, 2021 12:00:00 PM UTC is specified as 1635768000000.
	CreateTime int64

	// The ARN identifier of the Dataview.
	DataViewArn *string

	// The unique identifier for the Dataview.
	DataViewId *string

	// The unique identifier for the Dataset used in the Dataview.
	DatasetId *string

	// Options that define the destination type for the Dataview.
	DestinationTypeParams *types.DataViewDestinationTypeParams

	// Information about an error that occurred for the Dataview.
	ErrorInfo *types.DataViewErrorInfo

	// The last time that a Dataview was modified. The value is determined as epoch
	// time in milliseconds. For example, the value for Monday, November 1, 2021
	// 12:00:00 PM UTC is specified as 1635768000000.
	LastModifiedTime int64

	// Ordered set of column names used to partition data.
	PartitionColumns []string

	// Columns to be used for sorting the data.
	SortColumns []string

	// The status of a Dataview creation.
	//
	// * RUNNING – Dataview creation is running.
	//
	// *
	// STARTING – Dataview creation is starting.
	//
	// * FAILED – Dataview creation has
	// failed.
	//
	// * CANCELLED – Dataview creation has been cancelled.
	//
	// * TIMEOUT –
	// Dataview creation has timed out.
	//
	// * SUCCESS – Dataview creation has
	// succeeded.
	//
	// * PENDING – Dataview creation is pending.
	//
	// * FAILED_CLEANUP_FAILED –
	// Dataview creation failed and resource cleanup failed.
	Status types.DataViewStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetDataViewMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetDataView{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDataView{}, middleware.After)
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
	if err = addRestJsonContentTypeCustomization(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetDataViewValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDataView(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetDataView(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "finspace-api",
		OperationName: "GetDataView",
	}
}
