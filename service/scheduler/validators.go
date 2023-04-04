// Code generated by smithy-go-codegen DO NOT EDIT.

package scheduler

import (
	"context"
	"fmt"
	"github.com/matthew188/aws-sdk-go-v2/service/scheduler/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpCreateScheduleGroup struct {
}

func (*validateOpCreateScheduleGroup) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateScheduleGroup) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateScheduleGroupInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateScheduleGroupInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateSchedule struct {
}

func (*validateOpCreateSchedule) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateSchedule) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateScheduleInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateScheduleInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteScheduleGroup struct {
}

func (*validateOpDeleteScheduleGroup) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteScheduleGroup) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteScheduleGroupInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteScheduleGroupInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteSchedule struct {
}

func (*validateOpDeleteSchedule) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteSchedule) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteScheduleInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteScheduleInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetScheduleGroup struct {
}

func (*validateOpGetScheduleGroup) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetScheduleGroup) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetScheduleGroupInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetScheduleGroupInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetSchedule struct {
}

func (*validateOpGetSchedule) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetSchedule) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetScheduleInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetScheduleInput(input); err != nil {
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

type validateOpUpdateSchedule struct {
}

func (*validateOpUpdateSchedule) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateSchedule) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateScheduleInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateScheduleInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpCreateScheduleGroupValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateScheduleGroup{}, middleware.After)
}

func addOpCreateScheduleValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateSchedule{}, middleware.After)
}

func addOpDeleteScheduleGroupValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteScheduleGroup{}, middleware.After)
}

func addOpDeleteScheduleValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteSchedule{}, middleware.After)
}

func addOpGetScheduleGroupValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetScheduleGroup{}, middleware.After)
}

func addOpGetScheduleValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetSchedule{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func addOpUpdateScheduleValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateSchedule{}, middleware.After)
}

func validateAwsVpcConfiguration(v *types.AwsVpcConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AwsVpcConfiguration"}
	if v.Subnets == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Subnets"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateCapacityProviderStrategy(v []types.CapacityProviderStrategyItem) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CapacityProviderStrategy"}
	for i := range v {
		if err := validateCapacityProviderStrategyItem(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateCapacityProviderStrategyItem(v *types.CapacityProviderStrategyItem) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CapacityProviderStrategyItem"}
	if v.CapacityProvider == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CapacityProvider"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateEcsParameters(v *types.EcsParameters) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "EcsParameters"}
	if v.TaskDefinitionArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TaskDefinitionArn"))
	}
	if v.NetworkConfiguration != nil {
		if err := validateNetworkConfiguration(v.NetworkConfiguration); err != nil {
			invalidParams.AddNested("NetworkConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if v.CapacityProviderStrategy != nil {
		if err := validateCapacityProviderStrategy(v.CapacityProviderStrategy); err != nil {
			invalidParams.AddNested("CapacityProviderStrategy", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateEventBridgeParameters(v *types.EventBridgeParameters) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "EventBridgeParameters"}
	if v.DetailType == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DetailType"))
	}
	if v.Source == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Source"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateFlexibleTimeWindow(v *types.FlexibleTimeWindow) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "FlexibleTimeWindow"}
	if len(v.Mode) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Mode"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateKinesisParameters(v *types.KinesisParameters) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "KinesisParameters"}
	if v.PartitionKey == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PartitionKey"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateNetworkConfiguration(v *types.NetworkConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "NetworkConfiguration"}
	if v.AwsvpcConfiguration != nil {
		if err := validateAwsVpcConfiguration(v.AwsvpcConfiguration); err != nil {
			invalidParams.AddNested("AwsvpcConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateSageMakerPipelineParameter(v *types.SageMakerPipelineParameter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SageMakerPipelineParameter"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.Value == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Value"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateSageMakerPipelineParameterList(v []types.SageMakerPipelineParameter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SageMakerPipelineParameterList"}
	for i := range v {
		if err := validateSageMakerPipelineParameter(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateSageMakerPipelineParameters(v *types.SageMakerPipelineParameters) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SageMakerPipelineParameters"}
	if v.PipelineParameterList != nil {
		if err := validateSageMakerPipelineParameterList(v.PipelineParameterList); err != nil {
			invalidParams.AddNested("PipelineParameterList", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTag(v *types.Tag) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Tag"}
	if v.Key == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Key"))
	}
	if v.Value == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Value"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTagList(v []types.Tag) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagList"}
	for i := range v {
		if err := validateTag(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTarget(v *types.Target) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Target"}
	if v.Arn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Arn"))
	}
	if v.RoleArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RoleArn"))
	}
	if v.EcsParameters != nil {
		if err := validateEcsParameters(v.EcsParameters); err != nil {
			invalidParams.AddNested("EcsParameters", err.(smithy.InvalidParamsError))
		}
	}
	if v.EventBridgeParameters != nil {
		if err := validateEventBridgeParameters(v.EventBridgeParameters); err != nil {
			invalidParams.AddNested("EventBridgeParameters", err.(smithy.InvalidParamsError))
		}
	}
	if v.KinesisParameters != nil {
		if err := validateKinesisParameters(v.KinesisParameters); err != nil {
			invalidParams.AddNested("KinesisParameters", err.(smithy.InvalidParamsError))
		}
	}
	if v.SageMakerPipelineParameters != nil {
		if err := validateSageMakerPipelineParameters(v.SageMakerPipelineParameters); err != nil {
			invalidParams.AddNested("SageMakerPipelineParameters", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateScheduleGroupInput(v *CreateScheduleGroupInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateScheduleGroupInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.Tags != nil {
		if err := validateTagList(v.Tags); err != nil {
			invalidParams.AddNested("Tags", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateScheduleInput(v *CreateScheduleInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateScheduleInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.ScheduleExpression == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ScheduleExpression"))
	}
	if v.Target == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Target"))
	} else if v.Target != nil {
		if err := validateTarget(v.Target); err != nil {
			invalidParams.AddNested("Target", err.(smithy.InvalidParamsError))
		}
	}
	if v.FlexibleTimeWindow == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FlexibleTimeWindow"))
	} else if v.FlexibleTimeWindow != nil {
		if err := validateFlexibleTimeWindow(v.FlexibleTimeWindow); err != nil {
			invalidParams.AddNested("FlexibleTimeWindow", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteScheduleGroupInput(v *DeleteScheduleGroupInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteScheduleGroupInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteScheduleInput(v *DeleteScheduleInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteScheduleInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetScheduleGroupInput(v *GetScheduleGroupInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetScheduleGroupInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetScheduleInput(v *GetScheduleInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetScheduleInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
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
	} else if v.Tags != nil {
		if err := validateTagList(v.Tags); err != nil {
			invalidParams.AddNested("Tags", err.(smithy.InvalidParamsError))
		}
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

func validateOpUpdateScheduleInput(v *UpdateScheduleInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateScheduleInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.ScheduleExpression == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ScheduleExpression"))
	}
	if v.Target == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Target"))
	} else if v.Target != nil {
		if err := validateTarget(v.Target); err != nil {
			invalidParams.AddNested("Target", err.(smithy.InvalidParamsError))
		}
	}
	if v.FlexibleTimeWindow == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FlexibleTimeWindow"))
	} else if v.FlexibleTimeWindow != nil {
		if err := validateFlexibleTimeWindow(v.FlexibleTimeWindow); err != nil {
			invalidParams.AddNested("FlexibleTimeWindow", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
