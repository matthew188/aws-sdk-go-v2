// Code generated by smithy-go-codegen DO NOT EDIT.

package iotevents

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/iotevents/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves runtime information about a detector model analysis. After AWS IoT
// Events starts analyzing your detector model, you have up to 24 hours to retrieve
// the analysis results.
func (c *Client) DescribeDetectorModelAnalysis(ctx context.Context, params *DescribeDetectorModelAnalysisInput, optFns ...func(*Options)) (*DescribeDetectorModelAnalysisOutput, error) {
	if params == nil {
		params = &DescribeDetectorModelAnalysisInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDetectorModelAnalysis", params, optFns, c.addOperationDescribeDetectorModelAnalysisMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDetectorModelAnalysisOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDetectorModelAnalysisInput struct {

	// The ID of the analysis result that you want to retrieve.
	//
	// This member is required.
	AnalysisId *string

	noSmithyDocumentSerde
}

type DescribeDetectorModelAnalysisOutput struct {

	// The status of the analysis activity. The status can be one of the following
	// values:
	//
	// * RUNNING - AWS IoT Events is analyzing your detector model. This
	// process can take several minutes to complete.
	//
	// * COMPLETE - AWS IoT Events
	// finished analyzing your detector model.
	//
	// * FAILED - AWS IoT Events couldn't
	// analyze your detector model. Try again later.
	Status types.AnalysisStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeDetectorModelAnalysisMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeDetectorModelAnalysis{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeDetectorModelAnalysis{}, middleware.After)
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
	if err = addOpDescribeDetectorModelAnalysisValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDetectorModelAnalysis(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeDetectorModelAnalysis(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotevents",
		OperationName: "DescribeDetectorModelAnalysis",
	}
}
