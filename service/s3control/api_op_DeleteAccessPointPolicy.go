// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	s3controlcust "github.com/matthew188/aws-sdk-go-v2/service/s3control/internal/customizations"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"strings"
)

// Deletes the access point policy for the specified access point. All Amazon S3 on
// Outposts REST API requests for this action require an additional parameter of
// x-amz-outpost-id to be passed with the request. In addition, you must use an S3
// on Outposts endpoint hostname prefix instead of s3-control. For an example of
// the request syntax for Amazon S3 on Outposts that uses the S3 on Outposts
// endpoint hostname prefix and the x-amz-outpost-id derived by using the access
// point ARN, see the Examples
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteAccessPointPolicy.html#API_control_DeleteAccessPointPolicy_Examples)
// section. The following actions are related to DeleteAccessPointPolicy:
//
// *
// PutAccessPointPolicy
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutAccessPointPolicy.html)
//
// *
// GetAccessPointPolicy
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetAccessPointPolicy.html)
func (c *Client) DeleteAccessPointPolicy(ctx context.Context, params *DeleteAccessPointPolicyInput, optFns ...func(*Options)) (*DeleteAccessPointPolicyOutput, error) {
	if params == nil {
		params = &DeleteAccessPointPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteAccessPointPolicy", params, optFns, c.addOperationDeleteAccessPointPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteAccessPointPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteAccessPointPolicyInput struct {

	// The account ID for the account that owns the specified access point.
	//
	// This member is required.
	AccountId *string

	// The name of the access point whose policy you want to delete. For using this
	// parameter with Amazon S3 on Outposts with the REST API, you must specify the
	// name and the x-amz-outpost-id as well. For using this parameter with S3 on
	// Outposts with the Amazon Web Services SDK and CLI, you must specify the ARN of
	// the access point accessed in the format
	// arn:aws:s3-outposts:::outpost//accesspoint/. For example, to access the access
	// point reports-ap through Outpost my-outpost owned by account 123456789012 in
	// Region us-west-2, use the URL encoding of
	// arn:aws:s3-outposts:us-west-2:123456789012:outpost/my-outpost/accesspoint/reports-ap.
	// The value must be URL encoded.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

type DeleteAccessPointPolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteAccessPointPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpDeleteAccessPointPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpDeleteAccessPointPolicy{}, middleware.After)
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
	if err = addEndpointPrefix_opDeleteAccessPointPolicyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDeleteAccessPointPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteAccessPointPolicy(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addDeleteAccessPointPolicyUpdateEndpoint(stack, options); err != nil {
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

type endpointPrefix_opDeleteAccessPointPolicyMiddleware struct {
}

func (*endpointPrefix_opDeleteAccessPointPolicyMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opDeleteAccessPointPolicyMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	input, ok := in.Parameters.(*DeleteAccessPointPolicyInput)
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
func addEndpointPrefix_opDeleteAccessPointPolicyMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opDeleteAccessPointPolicyMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteAccessPointPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "DeleteAccessPointPolicy",
	}
}

func copyDeleteAccessPointPolicyInputForUpdateEndpoint(params interface{}) (interface{}, error) {
	input, ok := params.(*DeleteAccessPointPolicyInput)
	if !ok {
		return nil, fmt.Errorf("expect *DeleteAccessPointPolicyInput type, got %T", params)
	}
	cpy := *input
	return &cpy, nil
}
func getDeleteAccessPointPolicyARNMember(input interface{}) (*string, bool) {
	in := input.(*DeleteAccessPointPolicyInput)
	if in.Name == nil {
		return nil, false
	}
	return in.Name, true
}
func setDeleteAccessPointPolicyARNMember(input interface{}, v string) error {
	in := input.(*DeleteAccessPointPolicyInput)
	in.Name = &v
	return nil
}
func backFillDeleteAccessPointPolicyAccountID(input interface{}, v string) error {
	in := input.(*DeleteAccessPointPolicyInput)
	if in.AccountId != nil {
		if !strings.EqualFold(*in.AccountId, v) {
			return fmt.Errorf("error backfilling account id")
		}
		return nil
	}
	in.AccountId = &v
	return nil
}
func addDeleteAccessPointPolicyUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3controlcust.UpdateEndpoint(stack, s3controlcust.UpdateEndpointOptions{
		Accessor: s3controlcust.UpdateEndpointParameterAccessor{GetARNInput: getDeleteAccessPointPolicyARNMember,
			BackfillAccountID: backFillDeleteAccessPointPolicyAccountID,
			GetOutpostIDInput: nopGetOutpostIDFromInput,
			UpdateARNField:    setDeleteAccessPointPolicyARNMember,
			CopyInput:         copyDeleteAccessPointPolicyInputForUpdateEndpoint,
		},
		EndpointResolver:        options.EndpointResolver,
		EndpointResolverOptions: options.EndpointOptions,
		UseARNRegion:            options.UseARNRegion,
	})
}
