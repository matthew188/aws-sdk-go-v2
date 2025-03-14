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

// Gets a description of a hyperparameter tuning job.
func (c *Client) DescribeHyperParameterTuningJob(ctx context.Context, params *DescribeHyperParameterTuningJobInput, optFns ...func(*Options)) (*DescribeHyperParameterTuningJobOutput, error) {
	if params == nil {
		params = &DescribeHyperParameterTuningJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeHyperParameterTuningJob", params, optFns, c.addOperationDescribeHyperParameterTuningJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeHyperParameterTuningJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeHyperParameterTuningJobInput struct {

	// The name of the tuning job.
	//
	// This member is required.
	HyperParameterTuningJobName *string

	noSmithyDocumentSerde
}

type DescribeHyperParameterTuningJobOutput struct {

	// The date and time that the tuning job started.
	//
	// This member is required.
	CreationTime *time.Time

	// The Amazon Resource Name (ARN) of the tuning job.
	//
	// This member is required.
	HyperParameterTuningJobArn *string

	// The HyperParameterTuningJobConfig object that specifies the configuration of the
	// tuning job.
	//
	// This member is required.
	HyperParameterTuningJobConfig *types.HyperParameterTuningJobConfig

	// The name of the tuning job.
	//
	// This member is required.
	HyperParameterTuningJobName *string

	// The status of the tuning job: InProgress, Completed, Failed, Stopping, or
	// Stopped.
	//
	// This member is required.
	HyperParameterTuningJobStatus types.HyperParameterTuningJobStatus

	// The ObjectiveStatusCounters object that specifies the number of training jobs,
	// categorized by the status of their final objective metric, that this tuning job
	// launched.
	//
	// This member is required.
	ObjectiveStatusCounters *types.ObjectiveStatusCounters

	// The TrainingJobStatusCounters object that specifies the number of training jobs,
	// categorized by status, that this tuning job launched.
	//
	// This member is required.
	TrainingJobStatusCounters *types.TrainingJobStatusCounters

	// A TrainingJobSummary object that describes the training job that completed with
	// the best current HyperParameterTuningJobObjective.
	BestTrainingJob *types.HyperParameterTrainingJobSummary

	// The total resources consumed by your hyperparameter tuning job.
	ConsumedResources *types.HyperParameterTuningJobConsumedResources

	// If the tuning job failed, the reason it failed.
	FailureReason *string

	// The date and time that the tuning job ended.
	HyperParameterTuningEndTime *time.Time

	// The date and time that the status of the tuning job was modified.
	LastModifiedTime *time.Time

	// If the hyperparameter tuning job is an warm start tuning job with a
	// WarmStartType of IDENTICAL_DATA_AND_ALGORITHM, this is the TrainingJobSummary
	// for the training job with the best objective metric value of all training jobs
	// launched by this tuning job and all parent jobs specified for the warm start
	// tuning job.
	OverallBestTrainingJob *types.HyperParameterTrainingJobSummary

	// The HyperParameterTrainingJobDefinition object that specifies the definition of
	// the training jobs that this tuning job launches.
	TrainingJobDefinition *types.HyperParameterTrainingJobDefinition

	// A list of the HyperParameterTrainingJobDefinition objects launched for this
	// tuning job.
	TrainingJobDefinitions []types.HyperParameterTrainingJobDefinition

	// Tuning job completion information returned as the response from a hyperparameter
	// tuning job. This information tells if your tuning job has or has not converged.
	// It also includes the number of training jobs that have not improved model
	// performance as evaluated against the objective function.
	TuningJobCompletionDetails *types.HyperParameterTuningJobCompletionDetails

	// The configuration for starting the hyperparameter parameter tuning job using one
	// or more previous tuning jobs as a starting point. The results of previous tuning
	// jobs are used to inform which combinations of hyperparameters to search over in
	// the new tuning job.
	WarmStartConfig *types.HyperParameterTuningJobWarmStartConfig

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeHyperParameterTuningJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeHyperParameterTuningJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeHyperParameterTuningJob{}, middleware.After)
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
	if err = addOpDescribeHyperParameterTuningJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeHyperParameterTuningJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeHyperParameterTuningJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "DescribeHyperParameterTuningJob",
	}
}
