// Code generated by smithy-go-codegen DO NOT EDIT.

package eks

import (
	"context"
	"fmt"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/eks/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an Fargate profile for your Amazon EKS cluster. You must have at least
// one Fargate profile in a cluster to be able to run pods on Fargate. The Fargate
// profile allows an administrator to declare which pods run on Fargate and specify
// which pods run on which Fargate profile. This declaration is done through the
// profile’s selectors. Each profile can have up to five selectors that contain a
// namespace and labels. A namespace is required for every selector. The label
// field consists of multiple optional key-value pairs. Pods that match the
// selectors are scheduled on Fargate. If a to-be-scheduled pod matches any of the
// selectors in the Fargate profile, then that pod is run on Fargate. When you
// create a Fargate profile, you must specify a pod execution role to use with the
// pods that are scheduled with the profile. This role is added to the cluster's
// Kubernetes Role Based Access Control
// (https://kubernetes.io/docs/admin/authorization/rbac/) (RBAC) for authorization
// so that the kubelet that is running on the Fargate infrastructure can register
// with your Amazon EKS cluster so that it can appear in your cluster as a node.
// The pod execution role also provides IAM permissions to the Fargate
// infrastructure to allow read access to Amazon ECR image repositories. For more
// information, see Pod Execution Role
// (https://docs.aws.amazon.com/eks/latest/userguide/pod-execution-role.html) in
// the Amazon EKS User Guide. Fargate profiles are immutable. However, you can
// create a new updated profile to replace an existing profile and then delete the
// original after the updated profile has finished creating. If any Fargate
// profiles in a cluster are in the DELETING status, you must wait for that Fargate
// profile to finish deleting before you can create any other profiles in that
// cluster. For more information, see Fargate Profile
// (https://docs.aws.amazon.com/eks/latest/userguide/fargate-profile.html) in the
// Amazon EKS User Guide.
func (c *Client) CreateFargateProfile(ctx context.Context, params *CreateFargateProfileInput, optFns ...func(*Options)) (*CreateFargateProfileOutput, error) {
	if params == nil {
		params = &CreateFargateProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateFargateProfile", params, optFns, c.addOperationCreateFargateProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateFargateProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateFargateProfileInput struct {

	// The name of the Amazon EKS cluster to apply the Fargate profile to.
	//
	// This member is required.
	ClusterName *string

	// The name of the Fargate profile.
	//
	// This member is required.
	FargateProfileName *string

	// The Amazon Resource Name (ARN) of the pod execution role to use for pods that
	// match the selectors in the Fargate profile. The pod execution role allows
	// Fargate infrastructure to register with your cluster as a node, and it provides
	// read access to Amazon ECR image repositories. For more information, see Pod
	// Execution Role
	// (https://docs.aws.amazon.com/eks/latest/userguide/pod-execution-role.html) in
	// the Amazon EKS User Guide.
	//
	// This member is required.
	PodExecutionRoleArn *string

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request.
	ClientRequestToken *string

	// The selectors to match for pods to use this Fargate profile. Each selector must
	// have an associated namespace. Optionally, you can also specify labels for a
	// namespace. You may specify up to five selectors in a Fargate profile.
	Selectors []types.FargateProfileSelector

	// The IDs of subnets to launch your pods into. At this time, pods running on
	// Fargate are not assigned public IP addresses, so only private subnets (with no
	// direct route to an Internet Gateway) are accepted for this parameter.
	Subnets []string

	// The metadata to apply to the Fargate profile to assist with categorization and
	// organization. Each tag consists of a key and an optional value. You define both.
	// Fargate profile tags do not propagate to any other resources associated with the
	// Fargate profile, such as the pods that are scheduled with it.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateFargateProfileOutput struct {

	// The full description of your new Fargate profile.
	FargateProfile *types.FargateProfile

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateFargateProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateFargateProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateFargateProfile{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateFargateProfileMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateFargateProfileValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateFargateProfile(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateFargateProfile struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateFargateProfile) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateFargateProfile) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateFargateProfileInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateFargateProfileInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateFargateProfileMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateFargateProfile{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateFargateProfile(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "eks",
		OperationName: "CreateFargateProfile",
	}
}
