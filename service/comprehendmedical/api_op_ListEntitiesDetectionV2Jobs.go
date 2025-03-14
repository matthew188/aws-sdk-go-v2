// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehendmedical

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/comprehendmedical/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a list of medical entity detection jobs that you have submitted.
func (c *Client) ListEntitiesDetectionV2Jobs(ctx context.Context, params *ListEntitiesDetectionV2JobsInput, optFns ...func(*Options)) (*ListEntitiesDetectionV2JobsOutput, error) {
	if params == nil {
		params = &ListEntitiesDetectionV2JobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListEntitiesDetectionV2Jobs", params, optFns, c.addOperationListEntitiesDetectionV2JobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListEntitiesDetectionV2JobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListEntitiesDetectionV2JobsInput struct {

	// Filters the jobs that are returned. You can filter jobs based on their names,
	// status, or the date and time that they were submitted. You can only set one
	// filter at a time.
	Filter *types.ComprehendMedicalAsyncJobFilter

	// The maximum number of results to return in each page. The default is 100.
	MaxResults *int32

	// Identifies the next page of results to return.
	NextToken *string

	noSmithyDocumentSerde
}

type ListEntitiesDetectionV2JobsOutput struct {

	// A list containing the properties of each job returned.
	ComprehendMedicalAsyncJobPropertiesList []types.ComprehendMedicalAsyncJobProperties

	// Identifies the next page of results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListEntitiesDetectionV2JobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListEntitiesDetectionV2Jobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListEntitiesDetectionV2Jobs{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListEntitiesDetectionV2Jobs(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListEntitiesDetectionV2Jobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehendmedical",
		OperationName: "ListEntitiesDetectionV2Jobs",
	}
}
