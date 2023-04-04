// Code generated by smithy-go-codegen DO NOT EDIT.

package devopsguru

import (
	"context"
	"fmt"
	"github.com/matthew188/aws-sdk-go-v2/service/devopsguru/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpAddNotificationChannel struct {
}

func (*validateOpAddNotificationChannel) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpAddNotificationChannel) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*AddNotificationChannelInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpAddNotificationChannelInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteInsight struct {
}

func (*validateOpDeleteInsight) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteInsight) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteInsightInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteInsightInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeAccountOverview struct {
}

func (*validateOpDescribeAccountOverview) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeAccountOverview) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeAccountOverviewInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeAccountOverviewInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeAnomaly struct {
}

func (*validateOpDescribeAnomaly) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeAnomaly) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeAnomalyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeAnomalyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeInsight struct {
}

func (*validateOpDescribeInsight) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeInsight) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeInsightInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeInsightInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeOrganizationOverview struct {
}

func (*validateOpDescribeOrganizationOverview) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeOrganizationOverview) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeOrganizationOverviewInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeOrganizationOverviewInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeOrganizationResourceCollectionHealth struct {
}

func (*validateOpDescribeOrganizationResourceCollectionHealth) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeOrganizationResourceCollectionHealth) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeOrganizationResourceCollectionHealthInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeOrganizationResourceCollectionHealthInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeResourceCollectionHealth struct {
}

func (*validateOpDescribeResourceCollectionHealth) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeResourceCollectionHealth) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeResourceCollectionHealthInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeResourceCollectionHealthInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetResourceCollection struct {
}

func (*validateOpGetResourceCollection) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetResourceCollection) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetResourceCollectionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetResourceCollectionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListAnomaliesForInsight struct {
}

func (*validateOpListAnomaliesForInsight) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListAnomaliesForInsight) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListAnomaliesForInsightInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListAnomaliesForInsightInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListAnomalousLogGroups struct {
}

func (*validateOpListAnomalousLogGroups) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListAnomalousLogGroups) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListAnomalousLogGroupsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListAnomalousLogGroupsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListEvents struct {
}

func (*validateOpListEvents) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListEvents) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListEventsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListEventsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListInsights struct {
}

func (*validateOpListInsights) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListInsights) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListInsightsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListInsightsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListMonitoredResources struct {
}

func (*validateOpListMonitoredResources) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListMonitoredResources) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListMonitoredResourcesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListMonitoredResourcesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListOrganizationInsights struct {
}

func (*validateOpListOrganizationInsights) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListOrganizationInsights) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListOrganizationInsightsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListOrganizationInsightsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListRecommendations struct {
}

func (*validateOpListRecommendations) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListRecommendations) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListRecommendationsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListRecommendationsInput(input); err != nil {
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

type validateOpSearchInsights struct {
}

func (*validateOpSearchInsights) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpSearchInsights) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*SearchInsightsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpSearchInsightsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpSearchOrganizationInsights struct {
}

func (*validateOpSearchOrganizationInsights) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpSearchOrganizationInsights) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*SearchOrganizationInsightsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpSearchOrganizationInsightsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartCostEstimation struct {
}

func (*validateOpStartCostEstimation) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartCostEstimation) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartCostEstimationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartCostEstimationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateResourceCollection struct {
}

func (*validateOpUpdateResourceCollection) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateResourceCollection) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateResourceCollectionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateResourceCollectionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateServiceIntegration struct {
}

func (*validateOpUpdateServiceIntegration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateServiceIntegration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateServiceIntegrationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateServiceIntegrationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpAddNotificationChannelValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpAddNotificationChannel{}, middleware.After)
}

func addOpDeleteInsightValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteInsight{}, middleware.After)
}

func addOpDescribeAccountOverviewValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeAccountOverview{}, middleware.After)
}

func addOpDescribeAnomalyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeAnomaly{}, middleware.After)
}

func addOpDescribeInsightValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeInsight{}, middleware.After)
}

func addOpDescribeOrganizationOverviewValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeOrganizationOverview{}, middleware.After)
}

func addOpDescribeOrganizationResourceCollectionHealthValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeOrganizationResourceCollectionHealth{}, middleware.After)
}

func addOpDescribeResourceCollectionHealthValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeResourceCollectionHealth{}, middleware.After)
}

func addOpGetResourceCollectionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetResourceCollection{}, middleware.After)
}

func addOpListAnomaliesForInsightValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListAnomaliesForInsight{}, middleware.After)
}

func addOpListAnomalousLogGroupsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListAnomalousLogGroups{}, middleware.After)
}

func addOpListEventsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListEvents{}, middleware.After)
}

func addOpListInsightsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListInsights{}, middleware.After)
}

func addOpListMonitoredResourcesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListMonitoredResources{}, middleware.After)
}

func addOpListOrganizationInsightsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListOrganizationInsights{}, middleware.After)
}

func addOpListRecommendationsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListRecommendations{}, middleware.After)
}

func addOpRemoveNotificationChannelValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpRemoveNotificationChannel{}, middleware.After)
}

func addOpSearchInsightsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpSearchInsights{}, middleware.After)
}

func addOpSearchOrganizationInsightsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpSearchOrganizationInsights{}, middleware.After)
}

func addOpStartCostEstimationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartCostEstimation{}, middleware.After)
}

func addOpUpdateResourceCollectionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateResourceCollection{}, middleware.After)
}

func addOpUpdateServiceIntegrationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateServiceIntegration{}, middleware.After)
}

func validateCostEstimationResourceCollectionFilter(v *types.CostEstimationResourceCollectionFilter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CostEstimationResourceCollectionFilter"}
	if v.Tags != nil {
		if err := validateTagCostEstimationResourceCollectionFilters(v.Tags); err != nil {
			invalidParams.AddNested("Tags", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateEventTimeRange(v *types.EventTimeRange) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "EventTimeRange"}
	if v.FromTime == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FromTime"))
	}
	if v.ToTime == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ToTime"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateListEventsFilters(v *types.ListEventsFilters) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListEventsFilters"}
	if v.EventTimeRange != nil {
		if err := validateEventTimeRange(v.EventTimeRange); err != nil {
			invalidParams.AddNested("EventTimeRange", err.(smithy.InvalidParamsError))
		}
	}
	if v.ResourceCollection != nil {
		if err := validateResourceCollection(v.ResourceCollection); err != nil {
			invalidParams.AddNested("ResourceCollection", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateListInsightsAnyStatusFilter(v *types.ListInsightsAnyStatusFilter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListInsightsAnyStatusFilter"}
	if len(v.Type) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Type"))
	}
	if v.StartTimeRange == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StartTimeRange"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateListInsightsClosedStatusFilter(v *types.ListInsightsClosedStatusFilter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListInsightsClosedStatusFilter"}
	if len(v.Type) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Type"))
	}
	if v.EndTimeRange == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EndTimeRange"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateListInsightsOngoingStatusFilter(v *types.ListInsightsOngoingStatusFilter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListInsightsOngoingStatusFilter"}
	if len(v.Type) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Type"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateListInsightsStatusFilter(v *types.ListInsightsStatusFilter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListInsightsStatusFilter"}
	if v.Ongoing != nil {
		if err := validateListInsightsOngoingStatusFilter(v.Ongoing); err != nil {
			invalidParams.AddNested("Ongoing", err.(smithy.InvalidParamsError))
		}
	}
	if v.Closed != nil {
		if err := validateListInsightsClosedStatusFilter(v.Closed); err != nil {
			invalidParams.AddNested("Closed", err.(smithy.InvalidParamsError))
		}
	}
	if v.Any != nil {
		if err := validateListInsightsAnyStatusFilter(v.Any); err != nil {
			invalidParams.AddNested("Any", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateListMonitoredResourcesFilters(v *types.ListMonitoredResourcesFilters) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListMonitoredResourcesFilters"}
	if len(v.ResourcePermission) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("ResourcePermission"))
	}
	if v.ResourceTypeFilters == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceTypeFilters"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateNotificationChannelConfig(v *types.NotificationChannelConfig) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "NotificationChannelConfig"}
	if v.Sns == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Sns"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateResourceCollection(v *types.ResourceCollection) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ResourceCollection"}
	if v.Tags != nil {
		if err := validateTagCollections(v.Tags); err != nil {
			invalidParams.AddNested("Tags", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateSearchInsightsFilters(v *types.SearchInsightsFilters) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SearchInsightsFilters"}
	if v.ResourceCollection != nil {
		if err := validateResourceCollection(v.ResourceCollection); err != nil {
			invalidParams.AddNested("ResourceCollection", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateSearchOrganizationInsightsFilters(v *types.SearchOrganizationInsightsFilters) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SearchOrganizationInsightsFilters"}
	if v.ResourceCollection != nil {
		if err := validateResourceCollection(v.ResourceCollection); err != nil {
			invalidParams.AddNested("ResourceCollection", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTagCollection(v *types.TagCollection) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagCollection"}
	if v.AppBoundaryKey == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AppBoundaryKey"))
	}
	if v.TagValues == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TagValues"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTagCollections(v []types.TagCollection) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagCollections"}
	for i := range v {
		if err := validateTagCollection(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTagCostEstimationResourceCollectionFilter(v *types.TagCostEstimationResourceCollectionFilter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagCostEstimationResourceCollectionFilter"}
	if v.AppBoundaryKey == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AppBoundaryKey"))
	}
	if v.TagValues == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TagValues"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTagCostEstimationResourceCollectionFilters(v []types.TagCostEstimationResourceCollectionFilter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagCostEstimationResourceCollectionFilters"}
	for i := range v {
		if err := validateTagCostEstimationResourceCollectionFilter(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateUpdateResourceCollectionFilter(v *types.UpdateResourceCollectionFilter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateResourceCollectionFilter"}
	if v.Tags != nil {
		if err := validateUpdateTagCollectionFilters(v.Tags); err != nil {
			invalidParams.AddNested("Tags", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateUpdateTagCollectionFilter(v *types.UpdateTagCollectionFilter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateTagCollectionFilter"}
	if v.AppBoundaryKey == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AppBoundaryKey"))
	}
	if v.TagValues == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TagValues"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateUpdateTagCollectionFilters(v []types.UpdateTagCollectionFilter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateTagCollectionFilters"}
	for i := range v {
		if err := validateUpdateTagCollectionFilter(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpAddNotificationChannelInput(v *AddNotificationChannelInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AddNotificationChannelInput"}
	if v.Config == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Config"))
	} else if v.Config != nil {
		if err := validateNotificationChannelConfig(v.Config); err != nil {
			invalidParams.AddNested("Config", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteInsightInput(v *DeleteInsightInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteInsightInput"}
	if v.Id == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeAccountOverviewInput(v *DescribeAccountOverviewInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeAccountOverviewInput"}
	if v.FromTime == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FromTime"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeAnomalyInput(v *DescribeAnomalyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeAnomalyInput"}
	if v.Id == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeInsightInput(v *DescribeInsightInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeInsightInput"}
	if v.Id == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeOrganizationOverviewInput(v *DescribeOrganizationOverviewInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeOrganizationOverviewInput"}
	if v.FromTime == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FromTime"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeOrganizationResourceCollectionHealthInput(v *DescribeOrganizationResourceCollectionHealthInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeOrganizationResourceCollectionHealthInput"}
	if len(v.OrganizationResourceCollectionType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("OrganizationResourceCollectionType"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeResourceCollectionHealthInput(v *DescribeResourceCollectionHealthInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeResourceCollectionHealthInput"}
	if len(v.ResourceCollectionType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceCollectionType"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetResourceCollectionInput(v *GetResourceCollectionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetResourceCollectionInput"}
	if len(v.ResourceCollectionType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceCollectionType"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListAnomaliesForInsightInput(v *ListAnomaliesForInsightInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListAnomaliesForInsightInput"}
	if v.InsightId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("InsightId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListAnomalousLogGroupsInput(v *ListAnomalousLogGroupsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListAnomalousLogGroupsInput"}
	if v.InsightId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("InsightId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListEventsInput(v *ListEventsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListEventsInput"}
	if v.Filters == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Filters"))
	} else if v.Filters != nil {
		if err := validateListEventsFilters(v.Filters); err != nil {
			invalidParams.AddNested("Filters", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListInsightsInput(v *ListInsightsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListInsightsInput"}
	if v.StatusFilter == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StatusFilter"))
	} else if v.StatusFilter != nil {
		if err := validateListInsightsStatusFilter(v.StatusFilter); err != nil {
			invalidParams.AddNested("StatusFilter", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListMonitoredResourcesInput(v *ListMonitoredResourcesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListMonitoredResourcesInput"}
	if v.Filters != nil {
		if err := validateListMonitoredResourcesFilters(v.Filters); err != nil {
			invalidParams.AddNested("Filters", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListOrganizationInsightsInput(v *ListOrganizationInsightsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListOrganizationInsightsInput"}
	if v.StatusFilter == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StatusFilter"))
	} else if v.StatusFilter != nil {
		if err := validateListInsightsStatusFilter(v.StatusFilter); err != nil {
			invalidParams.AddNested("StatusFilter", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListRecommendationsInput(v *ListRecommendationsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListRecommendationsInput"}
	if v.InsightId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("InsightId"))
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
	if v.Id == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpSearchInsightsInput(v *SearchInsightsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SearchInsightsInput"}
	if v.StartTimeRange == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StartTimeRange"))
	}
	if v.Filters != nil {
		if err := validateSearchInsightsFilters(v.Filters); err != nil {
			invalidParams.AddNested("Filters", err.(smithy.InvalidParamsError))
		}
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

func validateOpSearchOrganizationInsightsInput(v *SearchOrganizationInsightsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SearchOrganizationInsightsInput"}
	if v.AccountIds == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AccountIds"))
	}
	if v.StartTimeRange == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StartTimeRange"))
	}
	if v.Filters != nil {
		if err := validateSearchOrganizationInsightsFilters(v.Filters); err != nil {
			invalidParams.AddNested("Filters", err.(smithy.InvalidParamsError))
		}
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

func validateOpStartCostEstimationInput(v *StartCostEstimationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartCostEstimationInput"}
	if v.ResourceCollection == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceCollection"))
	} else if v.ResourceCollection != nil {
		if err := validateCostEstimationResourceCollectionFilter(v.ResourceCollection); err != nil {
			invalidParams.AddNested("ResourceCollection", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateResourceCollectionInput(v *UpdateResourceCollectionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateResourceCollectionInput"}
	if len(v.Action) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Action"))
	}
	if v.ResourceCollection == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceCollection"))
	} else if v.ResourceCollection != nil {
		if err := validateUpdateResourceCollectionFilter(v.ResourceCollection); err != nil {
			invalidParams.AddNested("ResourceCollection", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateServiceIntegrationInput(v *UpdateServiceIntegrationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateServiceIntegrationInput"}
	if v.ServiceIntegration == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ServiceIntegration"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
