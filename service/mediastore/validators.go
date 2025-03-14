// Code generated by smithy-go-codegen DO NOT EDIT.

package mediastore

import (
	"context"
	"fmt"
	"github.com/matthew188/aws-sdk-go-v2/service/mediastore/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpCreateContainer struct {
}

func (*validateOpCreateContainer) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateContainer) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateContainerInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateContainerInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteContainer struct {
}

func (*validateOpDeleteContainer) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteContainer) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteContainerInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteContainerInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteContainerPolicy struct {
}

func (*validateOpDeleteContainerPolicy) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteContainerPolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteContainerPolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteContainerPolicyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteCorsPolicy struct {
}

func (*validateOpDeleteCorsPolicy) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteCorsPolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteCorsPolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteCorsPolicyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteLifecyclePolicy struct {
}

func (*validateOpDeleteLifecyclePolicy) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteLifecyclePolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteLifecyclePolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteLifecyclePolicyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteMetricPolicy struct {
}

func (*validateOpDeleteMetricPolicy) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteMetricPolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteMetricPolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteMetricPolicyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetContainerPolicy struct {
}

func (*validateOpGetContainerPolicy) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetContainerPolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetContainerPolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetContainerPolicyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetCorsPolicy struct {
}

func (*validateOpGetCorsPolicy) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetCorsPolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetCorsPolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetCorsPolicyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetLifecyclePolicy struct {
}

func (*validateOpGetLifecyclePolicy) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetLifecyclePolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetLifecyclePolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetLifecyclePolicyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetMetricPolicy struct {
}

func (*validateOpGetMetricPolicy) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetMetricPolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetMetricPolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetMetricPolicyInput(input); err != nil {
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

type validateOpPutContainerPolicy struct {
}

func (*validateOpPutContainerPolicy) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPutContainerPolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PutContainerPolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPutContainerPolicyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPutCorsPolicy struct {
}

func (*validateOpPutCorsPolicy) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPutCorsPolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PutCorsPolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPutCorsPolicyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPutLifecyclePolicy struct {
}

func (*validateOpPutLifecyclePolicy) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPutLifecyclePolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PutLifecyclePolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPutLifecyclePolicyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPutMetricPolicy struct {
}

func (*validateOpPutMetricPolicy) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPutMetricPolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PutMetricPolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPutMetricPolicyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartAccessLogging struct {
}

func (*validateOpStartAccessLogging) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartAccessLogging) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartAccessLoggingInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartAccessLoggingInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStopAccessLogging struct {
}

func (*validateOpStopAccessLogging) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStopAccessLogging) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StopAccessLoggingInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStopAccessLoggingInput(input); err != nil {
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

func addOpCreateContainerValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateContainer{}, middleware.After)
}

func addOpDeleteContainerValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteContainer{}, middleware.After)
}

func addOpDeleteContainerPolicyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteContainerPolicy{}, middleware.After)
}

func addOpDeleteCorsPolicyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteCorsPolicy{}, middleware.After)
}

func addOpDeleteLifecyclePolicyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteLifecyclePolicy{}, middleware.After)
}

func addOpDeleteMetricPolicyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteMetricPolicy{}, middleware.After)
}

func addOpGetContainerPolicyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetContainerPolicy{}, middleware.After)
}

func addOpGetCorsPolicyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetCorsPolicy{}, middleware.After)
}

func addOpGetLifecyclePolicyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetLifecyclePolicy{}, middleware.After)
}

func addOpGetMetricPolicyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetMetricPolicy{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpPutContainerPolicyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPutContainerPolicy{}, middleware.After)
}

func addOpPutCorsPolicyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPutCorsPolicy{}, middleware.After)
}

func addOpPutLifecyclePolicyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPutLifecyclePolicy{}, middleware.After)
}

func addOpPutMetricPolicyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPutMetricPolicy{}, middleware.After)
}

func addOpStartAccessLoggingValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartAccessLogging{}, middleware.After)
}

func addOpStopAccessLoggingValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStopAccessLogging{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func validateCorsPolicy(v []types.CorsRule) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CorsPolicy"}
	for i := range v {
		if err := validateCorsRule(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateCorsRule(v *types.CorsRule) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CorsRule"}
	if v.AllowedOrigins == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AllowedOrigins"))
	}
	if v.AllowedHeaders == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AllowedHeaders"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateMetricPolicy(v *types.MetricPolicy) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "MetricPolicy"}
	if len(v.ContainerLevelMetrics) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("ContainerLevelMetrics"))
	}
	if v.MetricPolicyRules != nil {
		if err := validateMetricPolicyRules(v.MetricPolicyRules); err != nil {
			invalidParams.AddNested("MetricPolicyRules", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateMetricPolicyRule(v *types.MetricPolicyRule) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "MetricPolicyRule"}
	if v.ObjectGroup == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ObjectGroup"))
	}
	if v.ObjectGroupName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ObjectGroupName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateMetricPolicyRules(v []types.MetricPolicyRule) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "MetricPolicyRules"}
	for i := range v {
		if err := validateMetricPolicyRule(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
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

func validateOpCreateContainerInput(v *CreateContainerInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateContainerInput"}
	if v.ContainerName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ContainerName"))
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

func validateOpDeleteContainerInput(v *DeleteContainerInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteContainerInput"}
	if v.ContainerName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ContainerName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteContainerPolicyInput(v *DeleteContainerPolicyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteContainerPolicyInput"}
	if v.ContainerName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ContainerName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteCorsPolicyInput(v *DeleteCorsPolicyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteCorsPolicyInput"}
	if v.ContainerName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ContainerName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteLifecyclePolicyInput(v *DeleteLifecyclePolicyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteLifecyclePolicyInput"}
	if v.ContainerName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ContainerName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteMetricPolicyInput(v *DeleteMetricPolicyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteMetricPolicyInput"}
	if v.ContainerName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ContainerName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetContainerPolicyInput(v *GetContainerPolicyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetContainerPolicyInput"}
	if v.ContainerName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ContainerName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetCorsPolicyInput(v *GetCorsPolicyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetCorsPolicyInput"}
	if v.ContainerName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ContainerName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetLifecyclePolicyInput(v *GetLifecyclePolicyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetLifecyclePolicyInput"}
	if v.ContainerName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ContainerName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetMetricPolicyInput(v *GetMetricPolicyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetMetricPolicyInput"}
	if v.ContainerName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ContainerName"))
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
	if v.Resource == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Resource"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPutContainerPolicyInput(v *PutContainerPolicyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PutContainerPolicyInput"}
	if v.ContainerName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ContainerName"))
	}
	if v.Policy == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Policy"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPutCorsPolicyInput(v *PutCorsPolicyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PutCorsPolicyInput"}
	if v.ContainerName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ContainerName"))
	}
	if v.CorsPolicy == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CorsPolicy"))
	} else if v.CorsPolicy != nil {
		if err := validateCorsPolicy(v.CorsPolicy); err != nil {
			invalidParams.AddNested("CorsPolicy", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPutLifecyclePolicyInput(v *PutLifecyclePolicyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PutLifecyclePolicyInput"}
	if v.ContainerName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ContainerName"))
	}
	if v.LifecyclePolicy == nil {
		invalidParams.Add(smithy.NewErrParamRequired("LifecyclePolicy"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPutMetricPolicyInput(v *PutMetricPolicyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PutMetricPolicyInput"}
	if v.ContainerName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ContainerName"))
	}
	if v.MetricPolicy == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MetricPolicy"))
	} else if v.MetricPolicy != nil {
		if err := validateMetricPolicy(v.MetricPolicy); err != nil {
			invalidParams.AddNested("MetricPolicy", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartAccessLoggingInput(v *StartAccessLoggingInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartAccessLoggingInput"}
	if v.ContainerName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ContainerName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStopAccessLoggingInput(v *StopAccessLoggingInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StopAccessLoggingInput"}
	if v.ContainerName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ContainerName"))
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
	if v.Resource == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Resource"))
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
	if v.Resource == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Resource"))
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
