// Code generated by smithy-go-codegen DO NOT EDIT.

package braket

import (
	"context"
	"fmt"
	"github.com/matthew188/aws-sdk-go-v2/service/braket/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpCancelJob struct {
}

func (*validateOpCancelJob) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCancelJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CancelJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCancelJobInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCancelQuantumTask struct {
}

func (*validateOpCancelQuantumTask) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCancelQuantumTask) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CancelQuantumTaskInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCancelQuantumTaskInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateJob struct {
}

func (*validateOpCreateJob) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateJobInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateQuantumTask struct {
}

func (*validateOpCreateQuantumTask) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateQuantumTask) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateQuantumTaskInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateQuantumTaskInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetDevice struct {
}

func (*validateOpGetDevice) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetDevice) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetDeviceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetDeviceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetJob struct {
}

func (*validateOpGetJob) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetJobInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetQuantumTask struct {
}

func (*validateOpGetQuantumTask) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetQuantumTask) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetQuantumTaskInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetQuantumTaskInput(input); err != nil {
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

type validateOpSearchDevices struct {
}

func (*validateOpSearchDevices) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpSearchDevices) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*SearchDevicesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpSearchDevicesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpSearchJobs struct {
}

func (*validateOpSearchJobs) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpSearchJobs) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*SearchJobsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpSearchJobsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpSearchQuantumTasks struct {
}

func (*validateOpSearchQuantumTasks) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpSearchQuantumTasks) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*SearchQuantumTasksInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpSearchQuantumTasksInput(input); err != nil {
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

func addOpCancelJobValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCancelJob{}, middleware.After)
}

func addOpCancelQuantumTaskValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCancelQuantumTask{}, middleware.After)
}

func addOpCreateJobValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateJob{}, middleware.After)
}

func addOpCreateQuantumTaskValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateQuantumTask{}, middleware.After)
}

func addOpGetDeviceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetDevice{}, middleware.After)
}

func addOpGetJobValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetJob{}, middleware.After)
}

func addOpGetQuantumTaskValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetQuantumTask{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpSearchDevicesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpSearchDevices{}, middleware.After)
}

func addOpSearchJobsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpSearchJobs{}, middleware.After)
}

func addOpSearchQuantumTasksValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpSearchQuantumTasks{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func validateAlgorithmSpecification(v *types.AlgorithmSpecification) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AlgorithmSpecification"}
	if v.ScriptModeConfig != nil {
		if err := validateScriptModeConfig(v.ScriptModeConfig); err != nil {
			invalidParams.AddNested("ScriptModeConfig", err.(smithy.InvalidParamsError))
		}
	}
	if v.ContainerImage != nil {
		if err := validateContainerImage(v.ContainerImage); err != nil {
			invalidParams.AddNested("ContainerImage", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateContainerImage(v *types.ContainerImage) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ContainerImage"}
	if v.Uri == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Uri"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateDataSource(v *types.DataSource) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DataSource"}
	if v.S3DataSource == nil {
		invalidParams.Add(smithy.NewErrParamRequired("S3DataSource"))
	} else if v.S3DataSource != nil {
		if err := validateS3DataSource(v.S3DataSource); err != nil {
			invalidParams.AddNested("S3DataSource", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateDeviceConfig(v *types.DeviceConfig) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeviceConfig"}
	if v.Device == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Device"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateInputConfigList(v []types.InputFileConfig) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "InputConfigList"}
	for i := range v {
		if err := validateInputFileConfig(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateInputFileConfig(v *types.InputFileConfig) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "InputFileConfig"}
	if v.ChannelName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ChannelName"))
	}
	if v.DataSource == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DataSource"))
	} else if v.DataSource != nil {
		if err := validateDataSource(v.DataSource); err != nil {
			invalidParams.AddNested("DataSource", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateInstanceConfig(v *types.InstanceConfig) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "InstanceConfig"}
	if len(v.InstanceType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("InstanceType"))
	}
	if v.VolumeSizeInGb == nil {
		invalidParams.Add(smithy.NewErrParamRequired("VolumeSizeInGb"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateJobCheckpointConfig(v *types.JobCheckpointConfig) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "JobCheckpointConfig"}
	if v.S3Uri == nil {
		invalidParams.Add(smithy.NewErrParamRequired("S3Uri"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateJobOutputDataConfig(v *types.JobOutputDataConfig) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "JobOutputDataConfig"}
	if v.S3Path == nil {
		invalidParams.Add(smithy.NewErrParamRequired("S3Path"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateS3DataSource(v *types.S3DataSource) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "S3DataSource"}
	if v.S3Uri == nil {
		invalidParams.Add(smithy.NewErrParamRequired("S3Uri"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateScriptModeConfig(v *types.ScriptModeConfig) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ScriptModeConfig"}
	if v.EntryPoint == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EntryPoint"))
	}
	if v.S3Uri == nil {
		invalidParams.Add(smithy.NewErrParamRequired("S3Uri"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateSearchDevicesFilter(v *types.SearchDevicesFilter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SearchDevicesFilter"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.Values == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Values"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateSearchDevicesFilterList(v []types.SearchDevicesFilter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SearchDevicesFilterList"}
	for i := range v {
		if err := validateSearchDevicesFilter(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateSearchJobsFilter(v *types.SearchJobsFilter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SearchJobsFilter"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.Values == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Values"))
	}
	if len(v.Operator) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Operator"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateSearchJobsFilterList(v []types.SearchJobsFilter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SearchJobsFilterList"}
	for i := range v {
		if err := validateSearchJobsFilter(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateSearchQuantumTasksFilter(v *types.SearchQuantumTasksFilter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SearchQuantumTasksFilter"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.Values == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Values"))
	}
	if len(v.Operator) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Operator"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateSearchQuantumTasksFilterList(v []types.SearchQuantumTasksFilter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SearchQuantumTasksFilterList"}
	for i := range v {
		if err := validateSearchQuantumTasksFilter(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCancelJobInput(v *CancelJobInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CancelJobInput"}
	if v.JobArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("JobArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCancelQuantumTaskInput(v *CancelQuantumTaskInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CancelQuantumTaskInput"}
	if v.QuantumTaskArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("QuantumTaskArn"))
	}
	if v.ClientToken == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClientToken"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateJobInput(v *CreateJobInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateJobInput"}
	if v.ClientToken == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClientToken"))
	}
	if v.AlgorithmSpecification == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AlgorithmSpecification"))
	} else if v.AlgorithmSpecification != nil {
		if err := validateAlgorithmSpecification(v.AlgorithmSpecification); err != nil {
			invalidParams.AddNested("AlgorithmSpecification", err.(smithy.InvalidParamsError))
		}
	}
	if v.InputDataConfig != nil {
		if err := validateInputConfigList(v.InputDataConfig); err != nil {
			invalidParams.AddNested("InputDataConfig", err.(smithy.InvalidParamsError))
		}
	}
	if v.OutputDataConfig == nil {
		invalidParams.Add(smithy.NewErrParamRequired("OutputDataConfig"))
	} else if v.OutputDataConfig != nil {
		if err := validateJobOutputDataConfig(v.OutputDataConfig); err != nil {
			invalidParams.AddNested("OutputDataConfig", err.(smithy.InvalidParamsError))
		}
	}
	if v.CheckpointConfig != nil {
		if err := validateJobCheckpointConfig(v.CheckpointConfig); err != nil {
			invalidParams.AddNested("CheckpointConfig", err.(smithy.InvalidParamsError))
		}
	}
	if v.JobName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("JobName"))
	}
	if v.RoleArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RoleArn"))
	}
	if v.InstanceConfig == nil {
		invalidParams.Add(smithy.NewErrParamRequired("InstanceConfig"))
	} else if v.InstanceConfig != nil {
		if err := validateInstanceConfig(v.InstanceConfig); err != nil {
			invalidParams.AddNested("InstanceConfig", err.(smithy.InvalidParamsError))
		}
	}
	if v.DeviceConfig == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DeviceConfig"))
	} else if v.DeviceConfig != nil {
		if err := validateDeviceConfig(v.DeviceConfig); err != nil {
			invalidParams.AddNested("DeviceConfig", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateQuantumTaskInput(v *CreateQuantumTaskInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateQuantumTaskInput"}
	if v.ClientToken == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClientToken"))
	}
	if v.DeviceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DeviceArn"))
	}
	if v.Shots == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Shots"))
	}
	if v.OutputS3Bucket == nil {
		invalidParams.Add(smithy.NewErrParamRequired("OutputS3Bucket"))
	}
	if v.OutputS3KeyPrefix == nil {
		invalidParams.Add(smithy.NewErrParamRequired("OutputS3KeyPrefix"))
	}
	if v.Action == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Action"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetDeviceInput(v *GetDeviceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetDeviceInput"}
	if v.DeviceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DeviceArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetJobInput(v *GetJobInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetJobInput"}
	if v.JobArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("JobArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetQuantumTaskInput(v *GetQuantumTaskInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetQuantumTaskInput"}
	if v.QuantumTaskArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("QuantumTaskArn"))
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

func validateOpSearchDevicesInput(v *SearchDevicesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SearchDevicesInput"}
	if v.Filters == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Filters"))
	} else if v.Filters != nil {
		if err := validateSearchDevicesFilterList(v.Filters); err != nil {
			invalidParams.AddNested("Filters", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpSearchJobsInput(v *SearchJobsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SearchJobsInput"}
	if v.Filters == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Filters"))
	} else if v.Filters != nil {
		if err := validateSearchJobsFilterList(v.Filters); err != nil {
			invalidParams.AddNested("Filters", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpSearchQuantumTasksInput(v *SearchQuantumTasksInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SearchQuantumTasksInput"}
	if v.Filters == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Filters"))
	} else if v.Filters != nil {
		if err := validateSearchQuantumTasksFilterList(v.Filters); err != nil {
			invalidParams.AddNested("Filters", err.(smithy.InvalidParamsError))
		}
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
