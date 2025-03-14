// Code generated by smithy-go-codegen DO NOT EDIT.

package privatenetworks

import (
	"context"
	"fmt"
	"github.com/matthew188/aws-sdk-go-v2/service/privatenetworks/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpAcknowledgeOrderReceipt struct {
}

func (*validateOpAcknowledgeOrderReceipt) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpAcknowledgeOrderReceipt) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*AcknowledgeOrderReceiptInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpAcknowledgeOrderReceiptInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpActivateDeviceIdentifier struct {
}

func (*validateOpActivateDeviceIdentifier) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpActivateDeviceIdentifier) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ActivateDeviceIdentifierInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpActivateDeviceIdentifierInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpActivateNetworkSite struct {
}

func (*validateOpActivateNetworkSite) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpActivateNetworkSite) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ActivateNetworkSiteInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpActivateNetworkSiteInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpConfigureAccessPoint struct {
}

func (*validateOpConfigureAccessPoint) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpConfigureAccessPoint) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ConfigureAccessPointInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpConfigureAccessPointInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateNetwork struct {
}

func (*validateOpCreateNetwork) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateNetwork) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateNetworkInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateNetworkInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateNetworkSite struct {
}

func (*validateOpCreateNetworkSite) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateNetworkSite) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateNetworkSiteInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateNetworkSiteInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeactivateDeviceIdentifier struct {
}

func (*validateOpDeactivateDeviceIdentifier) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeactivateDeviceIdentifier) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeactivateDeviceIdentifierInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeactivateDeviceIdentifierInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteNetwork struct {
}

func (*validateOpDeleteNetwork) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteNetwork) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteNetworkInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteNetworkInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteNetworkSite struct {
}

func (*validateOpDeleteNetworkSite) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteNetworkSite) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteNetworkSiteInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteNetworkSiteInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetDeviceIdentifier struct {
}

func (*validateOpGetDeviceIdentifier) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetDeviceIdentifier) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetDeviceIdentifierInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetDeviceIdentifierInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetNetwork struct {
}

func (*validateOpGetNetwork) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetNetwork) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetNetworkInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetNetworkInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetNetworkResource struct {
}

func (*validateOpGetNetworkResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetNetworkResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetNetworkResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetNetworkResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetNetworkSite struct {
}

func (*validateOpGetNetworkSite) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetNetworkSite) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetNetworkSiteInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetNetworkSiteInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetOrder struct {
}

func (*validateOpGetOrder) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetOrder) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetOrderInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetOrderInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListDeviceIdentifiers struct {
}

func (*validateOpListDeviceIdentifiers) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListDeviceIdentifiers) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListDeviceIdentifiersInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListDeviceIdentifiersInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListNetworkResources struct {
}

func (*validateOpListNetworkResources) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListNetworkResources) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListNetworkResourcesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListNetworkResourcesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListNetworkSites struct {
}

func (*validateOpListNetworkSites) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListNetworkSites) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListNetworkSitesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListNetworkSitesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListOrders struct {
}

func (*validateOpListOrders) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListOrders) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListOrdersInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListOrdersInput(input); err != nil {
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

type validateOpStartNetworkResourceUpdate struct {
}

func (*validateOpStartNetworkResourceUpdate) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartNetworkResourceUpdate) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartNetworkResourceUpdateInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartNetworkResourceUpdateInput(input); err != nil {
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

type validateOpUpdateNetworkSite struct {
}

func (*validateOpUpdateNetworkSite) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateNetworkSite) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateNetworkSiteInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateNetworkSiteInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateNetworkSitePlan struct {
}

func (*validateOpUpdateNetworkSitePlan) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateNetworkSitePlan) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateNetworkSitePlanInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateNetworkSitePlanInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpAcknowledgeOrderReceiptValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpAcknowledgeOrderReceipt{}, middleware.After)
}

func addOpActivateDeviceIdentifierValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpActivateDeviceIdentifier{}, middleware.After)
}

func addOpActivateNetworkSiteValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpActivateNetworkSite{}, middleware.After)
}

func addOpConfigureAccessPointValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpConfigureAccessPoint{}, middleware.After)
}

func addOpCreateNetworkValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateNetwork{}, middleware.After)
}

func addOpCreateNetworkSiteValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateNetworkSite{}, middleware.After)
}

func addOpDeactivateDeviceIdentifierValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeactivateDeviceIdentifier{}, middleware.After)
}

func addOpDeleteNetworkValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteNetwork{}, middleware.After)
}

func addOpDeleteNetworkSiteValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteNetworkSite{}, middleware.After)
}

func addOpGetDeviceIdentifierValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetDeviceIdentifier{}, middleware.After)
}

func addOpGetNetworkValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetNetwork{}, middleware.After)
}

func addOpGetNetworkResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetNetworkResource{}, middleware.After)
}

func addOpGetNetworkSiteValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetNetworkSite{}, middleware.After)
}

func addOpGetOrderValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetOrder{}, middleware.After)
}

func addOpListDeviceIdentifiersValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListDeviceIdentifiers{}, middleware.After)
}

func addOpListNetworkResourcesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListNetworkResources{}, middleware.After)
}

func addOpListNetworkSitesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListNetworkSites{}, middleware.After)
}

func addOpListOrdersValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListOrders{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpStartNetworkResourceUpdateValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartNetworkResourceUpdate{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func addOpUpdateNetworkSiteValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateNetworkSite{}, middleware.After)
}

func addOpUpdateNetworkSitePlanValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateNetworkSitePlan{}, middleware.After)
}

func validateAddress(v *types.Address) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Address"}
	if v.City == nil {
		invalidParams.Add(smithy.NewErrParamRequired("City"))
	}
	if v.Country == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Country"))
	}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.PostalCode == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PostalCode"))
	}
	if v.StateOrProvince == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StateOrProvince"))
	}
	if v.Street1 == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Street1"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateNameValuePair(v *types.NameValuePair) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "NameValuePair"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateNetworkResourceDefinition(v *types.NetworkResourceDefinition) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "NetworkResourceDefinition"}
	if len(v.Type) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Type"))
	}
	if v.Options != nil {
		if err := validateOptions(v.Options); err != nil {
			invalidParams.AddNested("Options", err.(smithy.InvalidParamsError))
		}
	}
	if v.Count == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Count"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateNetworkResourceDefinitions(v []types.NetworkResourceDefinition) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "NetworkResourceDefinitions"}
	for i := range v {
		if err := validateNetworkResourceDefinition(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOptions(v []types.NameValuePair) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Options"}
	for i := range v {
		if err := validateNameValuePair(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateSitePlan(v *types.SitePlan) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SitePlan"}
	if v.ResourceDefinitions != nil {
		if err := validateNetworkResourceDefinitions(v.ResourceDefinitions); err != nil {
			invalidParams.AddNested("ResourceDefinitions", err.(smithy.InvalidParamsError))
		}
	}
	if v.Options != nil {
		if err := validateOptions(v.Options); err != nil {
			invalidParams.AddNested("Options", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpAcknowledgeOrderReceiptInput(v *AcknowledgeOrderReceiptInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AcknowledgeOrderReceiptInput"}
	if v.OrderArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("OrderArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpActivateDeviceIdentifierInput(v *ActivateDeviceIdentifierInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ActivateDeviceIdentifierInput"}
	if v.DeviceIdentifierArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DeviceIdentifierArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpActivateNetworkSiteInput(v *ActivateNetworkSiteInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ActivateNetworkSiteInput"}
	if v.NetworkSiteArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NetworkSiteArn"))
	}
	if v.ShippingAddress == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ShippingAddress"))
	} else if v.ShippingAddress != nil {
		if err := validateAddress(v.ShippingAddress); err != nil {
			invalidParams.AddNested("ShippingAddress", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpConfigureAccessPointInput(v *ConfigureAccessPointInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ConfigureAccessPointInput"}
	if v.AccessPointArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AccessPointArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateNetworkInput(v *CreateNetworkInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateNetworkInput"}
	if v.NetworkName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NetworkName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateNetworkSiteInput(v *CreateNetworkSiteInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateNetworkSiteInput"}
	if v.NetworkSiteName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NetworkSiteName"))
	}
	if v.NetworkArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NetworkArn"))
	}
	if v.PendingPlan != nil {
		if err := validateSitePlan(v.PendingPlan); err != nil {
			invalidParams.AddNested("PendingPlan", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeactivateDeviceIdentifierInput(v *DeactivateDeviceIdentifierInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeactivateDeviceIdentifierInput"}
	if v.DeviceIdentifierArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DeviceIdentifierArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteNetworkInput(v *DeleteNetworkInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteNetworkInput"}
	if v.NetworkArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NetworkArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteNetworkSiteInput(v *DeleteNetworkSiteInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteNetworkSiteInput"}
	if v.NetworkSiteArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NetworkSiteArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetDeviceIdentifierInput(v *GetDeviceIdentifierInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetDeviceIdentifierInput"}
	if v.DeviceIdentifierArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DeviceIdentifierArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetNetworkInput(v *GetNetworkInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetNetworkInput"}
	if v.NetworkArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NetworkArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetNetworkResourceInput(v *GetNetworkResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetNetworkResourceInput"}
	if v.NetworkResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NetworkResourceArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetNetworkSiteInput(v *GetNetworkSiteInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetNetworkSiteInput"}
	if v.NetworkSiteArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NetworkSiteArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetOrderInput(v *GetOrderInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetOrderInput"}
	if v.OrderArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("OrderArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListDeviceIdentifiersInput(v *ListDeviceIdentifiersInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListDeviceIdentifiersInput"}
	if v.NetworkArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NetworkArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListNetworkResourcesInput(v *ListNetworkResourcesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListNetworkResourcesInput"}
	if v.NetworkArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NetworkArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListNetworkSitesInput(v *ListNetworkSitesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListNetworkSitesInput"}
	if v.NetworkArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NetworkArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListOrdersInput(v *ListOrdersInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListOrdersInput"}
	if v.NetworkArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NetworkArn"))
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

func validateOpStartNetworkResourceUpdateInput(v *StartNetworkResourceUpdateInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartNetworkResourceUpdateInput"}
	if v.NetworkResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NetworkResourceArn"))
	}
	if len(v.UpdateType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("UpdateType"))
	}
	if v.ShippingAddress != nil {
		if err := validateAddress(v.ShippingAddress); err != nil {
			invalidParams.AddNested("ShippingAddress", err.(smithy.InvalidParamsError))
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

func validateOpUpdateNetworkSiteInput(v *UpdateNetworkSiteInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateNetworkSiteInput"}
	if v.NetworkSiteArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NetworkSiteArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateNetworkSitePlanInput(v *UpdateNetworkSitePlanInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateNetworkSitePlanInput"}
	if v.NetworkSiteArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NetworkSiteArn"))
	}
	if v.PendingPlan == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PendingPlan"))
	} else if v.PendingPlan != nil {
		if err := validateSitePlan(v.PendingPlan); err != nil {
			invalidParams.AddNested("PendingPlan", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
