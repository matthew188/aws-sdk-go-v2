// Code generated by smithy-go-codegen DO NOT EDIT.

package emrserverless

import (
	"context"
	"fmt"
	"github.com/matthew188/aws-sdk-go-v2/service/emrserverless/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpCancelJobRun struct {
}

func (*validateOpCancelJobRun) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCancelJobRun) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CancelJobRunInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCancelJobRunInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateApplication struct {
}

func (*validateOpCreateApplication) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateApplication) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateApplicationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateApplicationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteApplication struct {
}

func (*validateOpDeleteApplication) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteApplication) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteApplicationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteApplicationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetApplication struct {
}

func (*validateOpGetApplication) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetApplication) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetApplicationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetApplicationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetDashboardForJobRun struct {
}

func (*validateOpGetDashboardForJobRun) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetDashboardForJobRun) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetDashboardForJobRunInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetDashboardForJobRunInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetJobRun struct {
}

func (*validateOpGetJobRun) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetJobRun) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetJobRunInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetJobRunInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListJobRuns struct {
}

func (*validateOpListJobRuns) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListJobRuns) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListJobRunsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListJobRunsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListTagsForResource struct {
}

func (*validateOpListTagsForResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListTagsForResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListTagsForResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListTagsForResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartApplication struct {
}

func (*validateOpStartApplication) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartApplication) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartApplicationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartApplicationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartJobRun struct {
}

func (*validateOpStartJobRun) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartJobRun) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartJobRunInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartJobRunInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStopApplication struct {
}

func (*validateOpStopApplication) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStopApplication) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StopApplicationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStopApplicationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpTagResource struct {
}

func (*validateOpTagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpTagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*TagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpTagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUntagResource struct {
}

func (*validateOpUntagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUntagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UntagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUntagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateApplication struct {
}

func (*validateOpUpdateApplication) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateApplication) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateApplicationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateApplicationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpCancelJobRunValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCancelJobRun{}, middleware.After)
}

func addOpCreateApplicationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateApplication{}, middleware.After)
}

func addOpDeleteApplicationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteApplication{}, middleware.After)
}

func addOpGetApplicationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetApplication{}, middleware.After)
}

func addOpGetDashboardForJobRunValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetDashboardForJobRun{}, middleware.After)
}

func addOpGetJobRunValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetJobRun{}, middleware.After)
}

func addOpListJobRunsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListJobRuns{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpStartApplicationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartApplication{}, middleware.After)
}

func addOpStartJobRunValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartJobRun{}, middleware.After)
}

func addOpStopApplicationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStopApplication{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func addOpUpdateApplicationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateApplication{}, middleware.After)
}

func validateConfiguration(v *types.Configuration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Configuration"}
	if v.Classification == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Classification"))
	}
	if v.Configurations != nil {
		if err := validateConfigurationList(v.Configurations); err != nil {
			invalidParams.AddNested("Configurations", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateConfigurationList(v []types.Configuration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ConfigurationList"}
	for i := range v {
		if err := validateConfiguration(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateConfigurationOverrides(v *types.ConfigurationOverrides) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ConfigurationOverrides"}
	if v.ApplicationConfiguration != nil {
		if err := validateConfigurationList(v.ApplicationConfiguration); err != nil {
			invalidParams.AddNested("ApplicationConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateHive(v *types.Hive) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Hive"}
	if v.Query == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Query"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateInitialCapacityConfig(v *types.InitialCapacityConfig) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "InitialCapacityConfig"}
	if v.WorkerConfiguration != nil {
		if err := validateWorkerResourceConfig(v.WorkerConfiguration); err != nil {
			invalidParams.AddNested("WorkerConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateInitialCapacityConfigMap(v map[string]types.InitialCapacityConfig) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "InitialCapacityConfigMap"}
	for key := range v {
		value := v[key]
		if err := validateInitialCapacityConfig(&value); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%q]", key), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateJobDriver(v types.JobDriver) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "JobDriver"}
	switch uv := v.(type) {
	case *types.JobDriverMemberHive:
		if err := validateHive(&uv.Value); err != nil {
			invalidParams.AddNested("[hive]", err.(smithy.InvalidParamsError))
		}

	case *types.JobDriverMemberSparkSubmit:
		if err := validateSparkSubmit(&uv.Value); err != nil {
			invalidParams.AddNested("[sparkSubmit]", err.(smithy.InvalidParamsError))
		}

	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateMaximumAllowedResources(v *types.MaximumAllowedResources) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "MaximumAllowedResources"}
	if v.Cpu == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Cpu"))
	}
	if v.Memory == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Memory"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateSparkSubmit(v *types.SparkSubmit) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SparkSubmit"}
	if v.EntryPoint == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EntryPoint"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateWorkerResourceConfig(v *types.WorkerResourceConfig) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "WorkerResourceConfig"}
	if v.Cpu == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Cpu"))
	}
	if v.Memory == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Memory"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCancelJobRunInput(v *CancelJobRunInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CancelJobRunInput"}
	if v.ApplicationId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ApplicationId"))
	}
	if v.JobRunId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("JobRunId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateApplicationInput(v *CreateApplicationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateApplicationInput"}
	if v.ReleaseLabel == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ReleaseLabel"))
	}
	if v.Type == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Type"))
	}
	if v.ClientToken == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClientToken"))
	}
	if v.InitialCapacity != nil {
		if err := validateInitialCapacityConfigMap(v.InitialCapacity); err != nil {
			invalidParams.AddNested("InitialCapacity", err.(smithy.InvalidParamsError))
		}
	}
	if v.MaximumCapacity != nil {
		if err := validateMaximumAllowedResources(v.MaximumCapacity); err != nil {
			invalidParams.AddNested("MaximumCapacity", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteApplicationInput(v *DeleteApplicationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteApplicationInput"}
	if v.ApplicationId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ApplicationId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetApplicationInput(v *GetApplicationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetApplicationInput"}
	if v.ApplicationId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ApplicationId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetDashboardForJobRunInput(v *GetDashboardForJobRunInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetDashboardForJobRunInput"}
	if v.ApplicationId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ApplicationId"))
	}
	if v.JobRunId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("JobRunId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetJobRunInput(v *GetJobRunInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetJobRunInput"}
	if v.ApplicationId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ApplicationId"))
	}
	if v.JobRunId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("JobRunId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListJobRunsInput(v *ListJobRunsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListJobRunsInput"}
	if v.ApplicationId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ApplicationId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListTagsForResourceInput(v *ListTagsForResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListTagsForResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartApplicationInput(v *StartApplicationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartApplicationInput"}
	if v.ApplicationId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ApplicationId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartJobRunInput(v *StartJobRunInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartJobRunInput"}
	if v.ApplicationId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ApplicationId"))
	}
	if v.ClientToken == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClientToken"))
	}
	if v.ExecutionRoleArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ExecutionRoleArn"))
	}
	if v.JobDriver != nil {
		if err := validateJobDriver(v.JobDriver); err != nil {
			invalidParams.AddNested("JobDriver", err.(smithy.InvalidParamsError))
		}
	}
	if v.ConfigurationOverrides != nil {
		if err := validateConfigurationOverrides(v.ConfigurationOverrides); err != nil {
			invalidParams.AddNested("ConfigurationOverrides", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStopApplicationInput(v *StopApplicationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StopApplicationInput"}
	if v.ApplicationId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ApplicationId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpTagResourceInput(v *TagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if v.Tags == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Tags"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUntagResourceInput(v *UntagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UntagResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if v.TagKeys == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TagKeys"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateApplicationInput(v *UpdateApplicationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateApplicationInput"}
	if v.ApplicationId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ApplicationId"))
	}
	if v.ClientToken == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClientToken"))
	}
	if v.InitialCapacity != nil {
		if err := validateInitialCapacityConfigMap(v.InitialCapacity); err != nil {
			invalidParams.AddNested("InitialCapacity", err.(smithy.InvalidParamsError))
		}
	}
	if v.MaximumCapacity != nil {
		if err := validateMaximumAllowedResources(v.MaximumCapacity); err != nil {
			invalidParams.AddNested("MaximumCapacity", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
