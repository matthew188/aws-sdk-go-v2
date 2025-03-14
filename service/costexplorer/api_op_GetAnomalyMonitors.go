// Code generated by smithy-go-codegen DO NOT EDIT.

package costexplorer

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/costexplorer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the cost anomaly monitor definitions for your account. You can filter
// using a list of cost anomaly monitor Amazon Resource Names (ARNs).
func (c *Client) GetAnomalyMonitors(ctx context.Context, params *GetAnomalyMonitorsInput, optFns ...func(*Options)) (*GetAnomalyMonitorsOutput, error) {
	if params == nil {
		params = &GetAnomalyMonitorsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAnomalyMonitors", params, optFns, c.addOperationGetAnomalyMonitorsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAnomalyMonitorsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAnomalyMonitorsInput struct {

	// The number of entries that a paginated response contains.
	MaxResults *int32

	// A list of cost anomaly monitor ARNs.
	MonitorArnList []string

	// The token to retrieve the next set of results. Amazon Web Services provides the
	// token when the response from a previous call has more results than the maximum
	// page size.
	NextPageToken *string

	noSmithyDocumentSerde
}

type GetAnomalyMonitorsOutput struct {

	// A list of cost anomaly monitors that includes the detailed metadata for each
	// monitor.
	//
	// This member is required.
	AnomalyMonitors []types.AnomalyMonitor

	// The token to retrieve the next set of results. Amazon Web Services provides the
	// token when the response from a previous call has more results than the maximum
	// page size.
	NextPageToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetAnomalyMonitorsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetAnomalyMonitors{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetAnomalyMonitors{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetAnomalyMonitors(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetAnomalyMonitors(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ce",
		OperationName: "GetAnomalyMonitors",
	}
}
