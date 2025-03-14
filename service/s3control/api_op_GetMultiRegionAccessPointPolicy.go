// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	s3controlcust "github.com/matthew188/aws-sdk-go-v2/service/s3control/internal/customizations"
	"github.com/matthew188/aws-sdk-go-v2/service/s3control/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"strings"
)

// Returns the access control policy of the specified Multi-Region Access Point.
// This action will always be routed to the US West (Oregon) Region. For more
// information about the restrictions around managing Multi-Region Access Points,
// see Managing Multi-Region Access Points
// (https://docs.aws.amazon.com/AmazonS3/latest/userguide/ManagingMultiRegionAccessPoints.html)
// in the Amazon S3 User Guide. The following actions are related to
// GetMultiRegionAccessPointPolicy:
//
// * GetMultiRegionAccessPointPolicyStatus
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetMultiRegionAccessPointPolicyStatus.html)
//
// *
// PutMultiRegionAccessPointPolicy
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutMultiRegionAccessPointPolicy.html)
func (c *Client) GetMultiRegionAccessPointPolicy(ctx context.Context, params *GetMultiRegionAccessPointPolicyInput, optFns ...func(*Options)) (*GetMultiRegionAccessPointPolicyOutput, error) {
	if params == nil {
		params = &GetMultiRegionAccessPointPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetMultiRegionAccessPointPolicy", params, optFns, c.addOperationGetMultiRegionAccessPointPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetMultiRegionAccessPointPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMultiRegionAccessPointPolicyInput struct {

	// The Amazon Web Services account ID for the owner of the Multi-Region Access
	// Point.
	//
	// This member is required.
	AccountId *string

	// Specifies the Multi-Region Access Point. The name of the Multi-Region Access
	// Point is different from the alias. For more information about the distinction
	// between the name and the alias of an Multi-Region Access Point, see Managing
	// Multi-Region Access Points
	// (https://docs.aws.amazon.com/AmazonS3/latest/userguide/CreatingMultiRegionAccessPoints.html#multi-region-access-point-naming)
	// in the Amazon S3 User Guide.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

type GetMultiRegionAccessPointPolicyOutput struct {

	// The policy associated with the specified Multi-Region Access Point.
	Policy *types.MultiRegionAccessPointPolicyDocument

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetMultiRegionAccessPointPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetMultiRegionAccessPointPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetMultiRegionAccessPointPolicy{}, middleware.After)
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
	if err = smithyhttp.AddContentChecksumMiddleware(stack); err != nil {
		return err
	}
	if err = addEndpointPrefix_opGetMultiRegionAccessPointPolicyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetMultiRegionAccessPointPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetMultiRegionAccessPointPolicy(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addGetMultiRegionAccessPointPolicyUpdateEndpoint(stack, options); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

type endpointPrefix_opGetMultiRegionAccessPointPolicyMiddleware struct {
}

func (*endpointPrefix_opGetMultiRegionAccessPointPolicyMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opGetMultiRegionAccessPointPolicyMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	input, ok := in.Parameters.(*GetMultiRegionAccessPointPolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input type %T", in.Parameters)
	}

	var prefix strings.Builder
	if input.AccountId == nil {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so may not be nil")}
	} else if !smithyhttp.ValidHostLabel(*input.AccountId) {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so must match \"[a-zA-Z0-9-]{1,63}\", but was \"%s\"", *input.AccountId)}
	} else {
		prefix.WriteString(*input.AccountId)
	}
	prefix.WriteString(".")
	req.URL.Host = prefix.String() + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opGetMultiRegionAccessPointPolicyMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opGetMultiRegionAccessPointPolicyMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opGetMultiRegionAccessPointPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "GetMultiRegionAccessPointPolicy",
	}
}

func copyGetMultiRegionAccessPointPolicyInputForUpdateEndpoint(params interface{}) (interface{}, error) {
	input, ok := params.(*GetMultiRegionAccessPointPolicyInput)
	if !ok {
		return nil, fmt.Errorf("expect *GetMultiRegionAccessPointPolicyInput type, got %T", params)
	}
	cpy := *input
	return &cpy, nil
}
func backFillGetMultiRegionAccessPointPolicyAccountID(input interface{}, v string) error {
	in := input.(*GetMultiRegionAccessPointPolicyInput)
	if in.AccountId != nil {
		if !strings.EqualFold(*in.AccountId, v) {
			return fmt.Errorf("error backfilling account id")
		}
		return nil
	}
	in.AccountId = &v
	return nil
}
func addGetMultiRegionAccessPointPolicyUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3controlcust.UpdateEndpoint(stack, s3controlcust.UpdateEndpointOptions{
		Accessor: s3controlcust.UpdateEndpointParameterAccessor{GetARNInput: nopGetARNAccessor,
			BackfillAccountID: nopBackfillAccountIDAccessor,
			GetOutpostIDInput: nopGetOutpostIDFromInput,
			UpdateARNField:    nopSetARNAccessor,
			CopyInput:         copyGetMultiRegionAccessPointPolicyInputForUpdateEndpoint,
		},
		EndpointResolver:        options.EndpointResolver,
		EndpointResolverOptions: options.EndpointOptions,
		UseARNRegion:            options.UseARNRegion,
	})
}
