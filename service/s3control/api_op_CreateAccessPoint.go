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

// Creates an access point and associates it with the specified bucket. For more
// information, see Managing Data Access with Amazon S3 Access Points
// (https://docs.aws.amazon.com/AmazonS3/latest/userguide/access-points.html) in
// the Amazon S3 User Guide. S3 on Outposts only supports VPC-style access points.
// For more information, see  Accessing Amazon S3 on Outposts using virtual private
// cloud (VPC) only access points
// (https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3onOutposts.html) in the
// Amazon S3 User Guide. All Amazon S3 on Outposts REST API requests for this
// action require an additional parameter of x-amz-outpost-id to be passed with the
// request. In addition, you must use an S3 on Outposts endpoint hostname prefix
// instead of s3-control. For an example of the request syntax for Amazon S3 on
// Outposts that uses the S3 on Outposts endpoint hostname prefix and the
// x-amz-outpost-id derived by using the access point ARN, see the Examples
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateAccessPoint.html#API_control_CreateAccessPoint_Examples)
// section. The following actions are related to CreateAccessPoint:
//
// *
// GetAccessPoint
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetAccessPoint.html)
//
// *
// DeleteAccessPoint
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteAccessPoint.html)
//
// *
// ListAccessPoints
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_ListAccessPoints.html)
func (c *Client) CreateAccessPoint(ctx context.Context, params *CreateAccessPointInput, optFns ...func(*Options)) (*CreateAccessPointOutput, error) {
	if params == nil {
		params = &CreateAccessPointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAccessPoint", params, optFns, c.addOperationCreateAccessPointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAccessPointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAccessPointInput struct {

	// The Amazon Web Services account ID for the account that owns the specified
	// access point.
	//
	// This member is required.
	AccountId *string

	// The name of the bucket that you want to associate this access point with. For
	// using this parameter with Amazon S3 on Outposts with the REST API, you must
	// specify the name and the x-amz-outpost-id as well. For using this parameter with
	// S3 on Outposts with the Amazon Web Services SDK and CLI, you must specify the
	// ARN of the bucket accessed in the format arn:aws:s3-outposts:::outpost//bucket/.
	// For example, to access the bucket reports through Outpost my-outpost owned by
	// account 123456789012 in Region us-west-2, use the URL encoding of
	// arn:aws:s3-outposts:us-west-2:123456789012:outpost/my-outpost/bucket/reports.
	// The value must be URL encoded.
	//
	// This member is required.
	Bucket *string

	// The name you want to assign to this access point.
	//
	// This member is required.
	Name *string

	// The Amazon Web Services account ID associated with the S3 bucket associated with
	// this access point.
	BucketAccountId *string

	// The PublicAccessBlock configuration that you want to apply to the access point.
	PublicAccessBlockConfiguration *types.PublicAccessBlockConfiguration

	// If you include this field, Amazon S3 restricts access to this access point to
	// requests from the specified virtual private cloud (VPC). This is required for
	// creating an access point for Amazon S3 on Outposts buckets.
	VpcConfiguration *types.VpcConfiguration

	noSmithyDocumentSerde
}

type CreateAccessPointOutput struct {

	// The ARN of the access point. This is only supported by Amazon S3 on Outposts.
	AccessPointArn *string

	// The name or alias of the access point.
	Alias *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAccessPointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpCreateAccessPoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpCreateAccessPoint{}, middleware.After)
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
	if err = addEndpointPrefix_opCreateAccessPointMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateAccessPointValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAccessPoint(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addCreateAccessPointUpdateEndpoint(stack, options); err != nil {
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

type endpointPrefix_opCreateAccessPointMiddleware struct {
}

func (*endpointPrefix_opCreateAccessPointMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opCreateAccessPointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	input, ok := in.Parameters.(*CreateAccessPointInput)
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
func addEndpointPrefix_opCreateAccessPointMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opCreateAccessPointMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opCreateAccessPoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "CreateAccessPoint",
	}
}

func copyCreateAccessPointInputForUpdateEndpoint(params interface{}) (interface{}, error) {
	input, ok := params.(*CreateAccessPointInput)
	if !ok {
		return nil, fmt.Errorf("expect *CreateAccessPointInput type, got %T", params)
	}
	cpy := *input
	return &cpy, nil
}
func getCreateAccessPointARNMember(input interface{}) (*string, bool) {
	in := input.(*CreateAccessPointInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func setCreateAccessPointARNMember(input interface{}, v string) error {
	in := input.(*CreateAccessPointInput)
	in.Bucket = &v
	return nil
}
func backFillCreateAccessPointAccountID(input interface{}, v string) error {
	in := input.(*CreateAccessPointInput)
	if in.AccountId != nil {
		if !strings.EqualFold(*in.AccountId, v) {
			return fmt.Errorf("error backfilling account id")
		}
		return nil
	}
	in.AccountId = &v
	return nil
}
func addCreateAccessPointUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3controlcust.UpdateEndpoint(stack, s3controlcust.UpdateEndpointOptions{
		Accessor: s3controlcust.UpdateEndpointParameterAccessor{GetARNInput: getCreateAccessPointARNMember,
			BackfillAccountID: backFillCreateAccessPointAccountID,
			GetOutpostIDInput: nopGetOutpostIDFromInput,
			UpdateARNField:    setCreateAccessPointARNMember,
			CopyInput:         copyCreateAccessPointInputForUpdateEndpoint,
		},
		EndpointResolver:        options.EndpointResolver,
		EndpointResolverOptions: options.EndpointOptions,
		UseARNRegion:            options.UseARNRegion,
	})
}
