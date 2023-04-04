// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/matthew188/aws-sdk-go-v2/aws/middleware"
	"github.com/matthew188/aws-sdk-go-v2/aws/signer/v4"
	"github.com/matthew188/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns information about an Amazon SageMaker AutoML job.
func (c *Client) DescribeAutoMLJob(ctx context.Context, params *DescribeAutoMLJobInput, optFns ...func(*Options)) (*DescribeAutoMLJobOutput, error) {
	if params == nil {
		params = &DescribeAutoMLJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAutoMLJob", params, optFns, c.addOperationDescribeAutoMLJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAutoMLJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAutoMLJobInput struct {

	// Requests information about an AutoML job using its unique name.
	//
	// This member is required.
	AutoMLJobName *string

	noSmithyDocumentSerde
}

type DescribeAutoMLJobOutput struct {

	// Returns the ARN of the AutoML job.
	//
	// This member is required.
	AutoMLJobArn *string

	// Returns the name of the AutoML job.
	//
	// This member is required.
	AutoMLJobName *string

	// Returns the secondary status of the AutoML job.
	//
	// This member is required.
	AutoMLJobSecondaryStatus types.AutoMLJobSecondaryStatus

	// Returns the status of the AutoML job.
	//
	// This member is required.
	AutoMLJobStatus types.AutoMLJobStatus

	// Returns the creation time of the AutoML job.
	//
	// This member is required.
	CreationTime *time.Time

	// Returns the input data configuration for the AutoML job.
	//
	// This member is required.
	InputDataConfig []types.AutoMLChannel

	// Returns the job's last modified time.
	//
	// This member is required.
	LastModifiedTime *time.Time

	// Returns the job's output data config.
	//
	// This member is required.
	OutputDataConfig *types.AutoMLOutputDataConfig

	// The Amazon Resource Name (ARN) of the Identity and Access Management (IAM) role
	// that has read permission to the input data location and write permission to the
	// output data location in Amazon S3.
	//
	// This member is required.
	RoleArn *string

	// Returns information on the job's artifacts found in AutoMLJobArtifacts.
	AutoMLJobArtifacts *types.AutoMLJobArtifacts

	// Returns the configuration for the AutoML job.
	AutoMLJobConfig *types.AutoMLJobConfig

	// Returns the job's objective.
	AutoMLJobObjective *types.AutoMLJobObjective

	// The best model candidate selected by SageMaker Autopilot using both the best
	// objective metric and lowest InferenceLatency
	// (https://docs.aws.amazon.com/sagemaker/latest/dg/autopilot-metrics-validation.html)
	// for an experiment.
	BestCandidate *types.AutoMLCandidate

	// Returns the end time of the AutoML job.
	EndTime *time.Time

	// Returns the failure reason for an AutoML job, when applicable.
	FailureReason *string

	// Indicates whether the output for an AutoML job generates candidate definitions
	// only.
	GenerateCandidateDefinitionsOnly bool

	// Indicates whether the model was deployed automatically to an endpoint and the
	// name of that endpoint if deployed automatically.
	ModelDeployConfig *types.ModelDeployConfig

	// Provides information about endpoint for the model deployment.
	ModelDeployResult *types.ModelDeployResult

	// Returns a list of reasons for partial failures within an AutoML job.
	PartialFailureReasons []types.AutoMLPartialFailureReason

	// Returns the job's problem type.
	ProblemType types.ProblemType

	// Contains ProblemType, AutoMLJobObjective, and CompletionCriteria. If you do not
	// provide these values, they are auto-inferred. If you do provide them, the values
	// used are the ones you provide.
	ResolvedAttributes *types.ResolvedAttributes

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeAutoMLJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeAutoMLJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeAutoMLJob{}, middleware.After)
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
	if err = addOpDescribeAutoMLJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAutoMLJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeAutoMLJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "DescribeAutoMLJob",
	}
}
