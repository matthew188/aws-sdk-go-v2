// Code generated by smithy-go-codegen DO NOT EDIT.

package licensemanagerusersubscriptions

import (
	"context"
	"fmt"
	"github.com/matthew188/aws-sdk-go-v2/service/licensemanagerusersubscriptions/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpAssociateUser struct {
}

func (*validateOpAssociateUser) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpAssociateUser) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*AssociateUserInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpAssociateUserInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeregisterIdentityProvider struct {
}

func (*validateOpDeregisterIdentityProvider) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeregisterIdentityProvider) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeregisterIdentityProviderInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeregisterIdentityProviderInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDisassociateUser struct {
}

func (*validateOpDisassociateUser) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDisassociateUser) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DisassociateUserInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDisassociateUserInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListProductSubscriptions struct {
}

func (*validateOpListProductSubscriptions) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListProductSubscriptions) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListProductSubscriptionsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListProductSubscriptionsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListUserAssociations struct {
}

func (*validateOpListUserAssociations) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListUserAssociations) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListUserAssociationsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListUserAssociationsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpRegisterIdentityProvider struct {
}

func (*validateOpRegisterIdentityProvider) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpRegisterIdentityProvider) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*RegisterIdentityProviderInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpRegisterIdentityProviderInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartProductSubscription struct {
}

func (*validateOpStartProductSubscription) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartProductSubscription) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartProductSubscriptionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartProductSubscriptionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStopProductSubscription struct {
}

func (*validateOpStopProductSubscription) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStopProductSubscription) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StopProductSubscriptionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStopProductSubscriptionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateIdentityProviderSettings struct {
}

func (*validateOpUpdateIdentityProviderSettings) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateIdentityProviderSettings) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateIdentityProviderSettingsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateIdentityProviderSettingsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpAssociateUserValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpAssociateUser{}, middleware.After)
}

func addOpDeregisterIdentityProviderValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeregisterIdentityProvider{}, middleware.After)
}

func addOpDisassociateUserValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDisassociateUser{}, middleware.After)
}

func addOpListProductSubscriptionsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListProductSubscriptions{}, middleware.After)
}

func addOpListUserAssociationsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListUserAssociations{}, middleware.After)
}

func addOpRegisterIdentityProviderValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpRegisterIdentityProvider{}, middleware.After)
}

func addOpStartProductSubscriptionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartProductSubscription{}, middleware.After)
}

func addOpStopProductSubscriptionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStopProductSubscription{}, middleware.After)
}

func addOpUpdateIdentityProviderSettingsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateIdentityProviderSettings{}, middleware.After)
}

func validateSettings(v *types.Settings) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Settings"}
	if v.Subnets == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Subnets"))
	}
	if v.SecurityGroupId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SecurityGroupId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateUpdateSettings(v *types.UpdateSettings) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateSettings"}
	if v.AddSubnets == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AddSubnets"))
	}
	if v.RemoveSubnets == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RemoveSubnets"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpAssociateUserInput(v *AssociateUserInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AssociateUserInput"}
	if v.Username == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Username"))
	}
	if v.InstanceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("InstanceId"))
	}
	if v.IdentityProvider == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdentityProvider"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeregisterIdentityProviderInput(v *DeregisterIdentityProviderInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeregisterIdentityProviderInput"}
	if v.IdentityProvider == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdentityProvider"))
	}
	if v.Product == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Product"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDisassociateUserInput(v *DisassociateUserInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DisassociateUserInput"}
	if v.Username == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Username"))
	}
	if v.InstanceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("InstanceId"))
	}
	if v.IdentityProvider == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdentityProvider"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListProductSubscriptionsInput(v *ListProductSubscriptionsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListProductSubscriptionsInput"}
	if v.Product == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Product"))
	}
	if v.IdentityProvider == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdentityProvider"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListUserAssociationsInput(v *ListUserAssociationsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListUserAssociationsInput"}
	if v.InstanceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("InstanceId"))
	}
	if v.IdentityProvider == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdentityProvider"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpRegisterIdentityProviderInput(v *RegisterIdentityProviderInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RegisterIdentityProviderInput"}
	if v.IdentityProvider == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdentityProvider"))
	}
	if v.Product == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Product"))
	}
	if v.Settings != nil {
		if err := validateSettings(v.Settings); err != nil {
			invalidParams.AddNested("Settings", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartProductSubscriptionInput(v *StartProductSubscriptionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartProductSubscriptionInput"}
	if v.Username == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Username"))
	}
	if v.IdentityProvider == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdentityProvider"))
	}
	if v.Product == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Product"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStopProductSubscriptionInput(v *StopProductSubscriptionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StopProductSubscriptionInput"}
	if v.Username == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Username"))
	}
	if v.IdentityProvider == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdentityProvider"))
	}
	if v.Product == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Product"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateIdentityProviderSettingsInput(v *UpdateIdentityProviderSettingsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateIdentityProviderSettingsInput"}
	if v.IdentityProvider == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdentityProvider"))
	}
	if v.Product == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Product"))
	}
	if v.UpdateSettings == nil {
		invalidParams.Add(smithy.NewErrParamRequired("UpdateSettings"))
	} else if v.UpdateSettings != nil {
		if err := validateUpdateSettings(v.UpdateSettings); err != nil {
			invalidParams.AddNested("UpdateSettings", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
