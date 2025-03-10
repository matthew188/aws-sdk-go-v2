// Code generated by smithy-go-codegen DO NOT EDIT.

package codeguruprofiler

import (
	"context"
	"fmt"
	"github.com/matthew188/aws-sdk-go-v2/service/codeguruprofiler/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpAddNotificationChannels struct {
}

func (*validateOpAddNotificationChannels) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpAddNotificationChannels) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*AddNotificationChannelsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpAddNotificationChannelsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpBatchGetFrameMetricData struct {
}

func (*validateOpBatchGetFrameMetricData) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpBatchGetFrameMetricData) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*BatchGetFrameMetricDataInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpBatchGetFrameMetricDataInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpConfigureAgent struct {
}

func (*validateOpConfigureAgent) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpConfigureAgent) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ConfigureAgentInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpConfigureAgentInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateProfilingGroup struct {
}

func (*validateOpCreateProfilingGroup) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateProfilingGroup) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateProfilingGroupInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateProfilingGroupInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteProfilingGroup struct {
}

func (*validateOpDeleteProfilingGroup) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteProfilingGroup) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteProfilingGroupInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteProfilingGroupInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeProfilingGroup struct {
}

func (*validateOpDescribeProfilingGroup) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeProfilingGroup) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeProfilingGroupInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeProfilingGroupInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetNotificationConfiguration struct {
}

func (*validateOpGetNotificationConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetNotificationConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetNotificationConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetNotificationConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetPolicy struct {
}

func (*validateOpGetPolicy) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetPolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetPolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetPolicyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetProfile struct {
}

func (*validateOpGetProfile) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetProfile) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetProfileInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetProfileInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetRecommendations struct {
}

func (*validateOpGetRecommendations) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetRecommendations) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetRecommendationsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetRecommendationsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListFindingsReports struct {
}

func (*validateOpListFindingsReports) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListFindingsReports) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListFindingsReportsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListFindingsReportsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListProfileTimes struct {
}

func (*validateOpListProfileTimes) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListProfileTimes) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListProfileTimesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListProfileTimesInput(input); err != nil {
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

type validateOpPostAgentProfile struct {
}

func (*validateOpPostAgentProfile) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPostAgentProfile) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PostAgentProfileInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPostAgentProfileInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPutPermission struct {
}

func (*validateOpPutPermission) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPutPermission) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PutPermissionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPutPermissionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpRemoveNotificationChannel struct {
}

func (*validateOpRemoveNotificationChannel) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpRemoveNotificationChannel) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*RemoveNotificationChannelInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpRemoveNotificationChannelInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpRemovePermission struct {
}

func (*validateOpRemovePermission) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpRemovePermission) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*RemovePermissionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpRemovePermissionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpSubmitFeedback struct {
}

func (*validateOpSubmitFeedback) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpSubmitFeedback) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*SubmitFeedbackInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpSubmitFeedbackInput(input); err != nil {
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

type validateOpUpdateProfilingGroup struct {
}

func (*validateOpUpdateProfilingGroup) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateProfilingGroup) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateProfilingGroupInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateProfilingGroupInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpAddNotificationChannelsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpAddNotificationChannels{}, middleware.After)
}

func addOpBatchGetFrameMetricDataValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpBatchGetFrameMetricData{}, middleware.After)
}

func addOpConfigureAgentValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpConfigureAgent{}, middleware.After)
}

func addOpCreateProfilingGroupValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateProfilingGroup{}, middleware.After)
}

func addOpDeleteProfilingGroupValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteProfilingGroup{}, middleware.After)
}

func addOpDescribeProfilingGroupValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeProfilingGroup{}, middleware.After)
}

func addOpGetNotificationConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetNotificationConfiguration{}, middleware.After)
}

func addOpGetPolicyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetPolicy{}, middleware.After)
}

func addOpGetProfileValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetProfile{}, middleware.After)
}

func addOpGetRecommendationsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetRecommendations{}, middleware.After)
}

func addOpListFindingsReportsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListFindingsReports{}, middleware.After)
}

func addOpListProfileTimesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListProfileTimes{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpPostAgentProfileValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPostAgentProfile{}, middleware.After)
}

func addOpPutPermissionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPutPermission{}, middleware.After)
}

func addOpRemoveNotificationChannelValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpRemoveNotificationChannel{}, middleware.After)
}

func addOpRemovePermissionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpRemovePermission{}, middleware.After)
}

func addOpSubmitFeedbackValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpSubmitFeedback{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func addOpUpdateProfilingGroupValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateProfilingGroup{}, middleware.After)
}

func validateAgentOrchestrationConfig(v *types.AgentOrchestrationConfig) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AgentOrchestrationConfig"}
	if v.ProfilingEnabled == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProfilingEnabled"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateChannel(v *types.Channel) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Channel"}
	if v.Uri == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Uri"))
	}
	if v.EventPublishers == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EventPublishers"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateChannels(v []types.Channel) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Channels"}
	for i := range v {
		if err := validateChannel(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateFrameMetric(v *types.FrameMetric) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "FrameMetric"}
	if v.FrameName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FrameName"))
	}
	if len(v.Type) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Type"))
	}
	if v.ThreadStates == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ThreadStates"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateFrameMetrics(v []types.FrameMetric) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "FrameMetrics"}
	for i := range v {
		if err := validateFrameMetric(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpAddNotificationChannelsInput(v *AddNotificationChannelsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AddNotificationChannelsInput"}
	if v.ProfilingGroupName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProfilingGroupName"))
	}
	if v.Channels == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Channels"))
	} else if v.Channels != nil {
		if err := validateChannels(v.Channels); err != nil {
			invalidParams.AddNested("Channels", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpBatchGetFrameMetricDataInput(v *BatchGetFrameMetricDataInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "BatchGetFrameMetricDataInput"}
	if v.ProfilingGroupName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProfilingGroupName"))
	}
	if v.FrameMetrics != nil {
		if err := validateFrameMetrics(v.FrameMetrics); err != nil {
			invalidParams.AddNested("FrameMetrics", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpConfigureAgentInput(v *ConfigureAgentInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ConfigureAgentInput"}
	if v.ProfilingGroupName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProfilingGroupName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateProfilingGroupInput(v *CreateProfilingGroupInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateProfilingGroupInput"}
	if v.ProfilingGroupName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProfilingGroupName"))
	}
	if v.ClientToken == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClientToken"))
	}
	if v.AgentOrchestrationConfig != nil {
		if err := validateAgentOrchestrationConfig(v.AgentOrchestrationConfig); err != nil {
			invalidParams.AddNested("AgentOrchestrationConfig", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteProfilingGroupInput(v *DeleteProfilingGroupInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteProfilingGroupInput"}
	if v.ProfilingGroupName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProfilingGroupName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeProfilingGroupInput(v *DescribeProfilingGroupInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeProfilingGroupInput"}
	if v.ProfilingGroupName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProfilingGroupName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetNotificationConfigurationInput(v *GetNotificationConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetNotificationConfigurationInput"}
	if v.ProfilingGroupName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProfilingGroupName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetPolicyInput(v *GetPolicyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetPolicyInput"}
	if v.ProfilingGroupName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProfilingGroupName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetProfileInput(v *GetProfileInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetProfileInput"}
	if v.ProfilingGroupName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProfilingGroupName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetRecommendationsInput(v *GetRecommendationsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetRecommendationsInput"}
	if v.ProfilingGroupName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProfilingGroupName"))
	}
	if v.StartTime == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StartTime"))
	}
	if v.EndTime == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EndTime"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListFindingsReportsInput(v *ListFindingsReportsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListFindingsReportsInput"}
	if v.ProfilingGroupName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProfilingGroupName"))
	}
	if v.StartTime == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StartTime"))
	}
	if v.EndTime == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EndTime"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListProfileTimesInput(v *ListProfileTimesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListProfileTimesInput"}
	if v.ProfilingGroupName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProfilingGroupName"))
	}
	if v.StartTime == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StartTime"))
	}
	if v.EndTime == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EndTime"))
	}
	if len(v.Period) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Period"))
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

func validateOpPostAgentProfileInput(v *PostAgentProfileInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PostAgentProfileInput"}
	if v.ProfilingGroupName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProfilingGroupName"))
	}
	if v.AgentProfile == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AgentProfile"))
	}
	if v.ContentType == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ContentType"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPutPermissionInput(v *PutPermissionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PutPermissionInput"}
	if v.ProfilingGroupName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProfilingGroupName"))
	}
	if len(v.ActionGroup) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("ActionGroup"))
	}
	if v.Principals == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Principals"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpRemoveNotificationChannelInput(v *RemoveNotificationChannelInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RemoveNotificationChannelInput"}
	if v.ProfilingGroupName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProfilingGroupName"))
	}
	if v.ChannelId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ChannelId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpRemovePermissionInput(v *RemovePermissionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RemovePermissionInput"}
	if v.ProfilingGroupName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProfilingGroupName"))
	}
	if len(v.ActionGroup) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("ActionGroup"))
	}
	if v.RevisionId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RevisionId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpSubmitFeedbackInput(v *SubmitFeedbackInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SubmitFeedbackInput"}
	if v.ProfilingGroupName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProfilingGroupName"))
	}
	if v.AnomalyInstanceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AnomalyInstanceId"))
	}
	if len(v.Type) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Type"))
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

func validateOpUpdateProfilingGroupInput(v *UpdateProfilingGroupInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateProfilingGroupInput"}
	if v.ProfilingGroupName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProfilingGroupName"))
	}
	if v.AgentOrchestrationConfig == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AgentOrchestrationConfig"))
	} else if v.AgentOrchestrationConfig != nil {
		if err := validateAgentOrchestrationConfig(v.AgentOrchestrationConfig); err != nil {
			invalidParams.AddNested("AgentOrchestrationConfig", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
