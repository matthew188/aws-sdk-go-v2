// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/connect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the traffic distribution for a given traffic distribution group. For
// more information about updating a traffic distribution group, see Update
// telephony traffic distribution across Amazon Web Services Regions
// (https://docs.aws.amazon.com/connect/latest/adminguide/update-telephony-traffic-distribution.html)
// in the Amazon Connect Administrator Guide.
func (c *Client) UpdateTrafficDistribution(ctx context.Context, params *UpdateTrafficDistributionInput, optFns ...func(*Options)) (*UpdateTrafficDistributionOutput, error) {
	if params == nil {
		params = &UpdateTrafficDistributionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateTrafficDistribution", params, optFns, c.addOperationUpdateTrafficDistributionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateTrafficDistributionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateTrafficDistributionInput struct {

	// The identifier of the traffic distribution group. This can be the ID or the ARN
	// if the API is being called in the Region where the traffic distribution group
	// was created. The ARN must be provided if the call is from the replicated Region.
	//
	// This member is required.
	Id *string

	// The distribution of traffic between the instance and its replica(s).
	TelephonyConfig *types.TelephonyConfig

	noSmithyDocumentSerde
}

type UpdateTrafficDistributionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateTrafficDistributionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateTrafficDistribution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateTrafficDistribution{}, middleware.After)
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
	if err = addOpUpdateTrafficDistributionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateTrafficDistribution(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateTrafficDistribution(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "UpdateTrafficDistribution",
	}
}
