// Code generated by smithy-go-codegen DO NOT EDIT.

package grafana

import (
	"context"
	"fmt"
	"github.com/matthew188/aws-sdk-go-v2/service/grafana/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpAssociateLicense struct {
}

func (*validateOpAssociateLicense) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpAssociateLicense) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*AssociateLicenseInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpAssociateLicenseInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateWorkspaceApiKey struct {
}

func (*validateOpCreateWorkspaceApiKey) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateWorkspaceApiKey) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateWorkspaceApiKeyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateWorkspaceApiKeyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateWorkspace struct {
}

func (*validateOpCreateWorkspace) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateWorkspace) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateWorkspaceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateWorkspaceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteWorkspaceApiKey struct {
}

func (*validateOpDeleteWorkspaceApiKey) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteWorkspaceApiKey) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteWorkspaceApiKeyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteWorkspaceApiKeyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteWorkspace struct {
}

func (*validateOpDeleteWorkspace) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteWorkspace) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteWorkspaceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteWorkspaceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeWorkspaceAuthentication struct {
}

func (*validateOpDescribeWorkspaceAuthentication) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeWorkspaceAuthentication) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeWorkspaceAuthenticationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeWorkspaceAuthenticationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeWorkspaceConfiguration struct {
}

func (*validateOpDescribeWorkspaceConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeWorkspaceConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeWorkspaceConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeWorkspaceConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeWorkspace struct {
}

func (*validateOpDescribeWorkspace) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeWorkspace) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeWorkspaceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeWorkspaceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDisassociateLicense struct {
}

func (*validateOpDisassociateLicense) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDisassociateLicense) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DisassociateLicenseInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDisassociateLicenseInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListPermissions struct {
}

func (*validateOpListPermissions) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListPermissions) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListPermissionsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListPermissionsInput(input); err != nil {
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

type validateOpUpdatePermissions struct {
}

func (*validateOpUpdatePermissions) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdatePermissions) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdatePermissionsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdatePermissionsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateWorkspaceAuthentication struct {
}

func (*validateOpUpdateWorkspaceAuthentication) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateWorkspaceAuthentication) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateWorkspaceAuthenticationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateWorkspaceAuthenticationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateWorkspaceConfiguration struct {
}

func (*validateOpUpdateWorkspaceConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateWorkspaceConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateWorkspaceConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateWorkspaceConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateWorkspace struct {
}

func (*validateOpUpdateWorkspace) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateWorkspace) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateWorkspaceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateWorkspaceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpAssociateLicenseValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpAssociateLicense{}, middleware.After)
}

func addOpCreateWorkspaceApiKeyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateWorkspaceApiKey{}, middleware.After)
}

func addOpCreateWorkspaceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateWorkspace{}, middleware.After)
}

func addOpDeleteWorkspaceApiKeyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteWorkspaceApiKey{}, middleware.After)
}

func addOpDeleteWorkspaceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteWorkspace{}, middleware.After)
}

func addOpDescribeWorkspaceAuthenticationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeWorkspaceAuthentication{}, middleware.After)
}

func addOpDescribeWorkspaceConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeWorkspaceConfiguration{}, middleware.After)
}

func addOpDescribeWorkspaceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeWorkspace{}, middleware.After)
}

func addOpDisassociateLicenseValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDisassociateLicense{}, middleware.After)
}

func addOpListPermissionsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListPermissions{}, middleware.After)
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

func addOpUpdatePermissionsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdatePermissions{}, middleware.After)
}

func addOpUpdateWorkspaceAuthenticationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateWorkspaceAuthentication{}, middleware.After)
}

func addOpUpdateWorkspaceConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateWorkspaceConfiguration{}, middleware.After)
}

func addOpUpdateWorkspaceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateWorkspace{}, middleware.After)
}

func validateNetworkAccessConfiguration(v *types.NetworkAccessConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "NetworkAccessConfiguration"}
	if v.PrefixListIds == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PrefixListIds"))
	}
	if v.VpceIds == nil {
		invalidParams.Add(smithy.NewErrParamRequired("VpceIds"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateSamlConfiguration(v *types.SamlConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SamlConfiguration"}
	if v.IdpMetadata == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdpMetadata"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateUpdateInstruction(v *types.UpdateInstruction) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateInstruction"}
	if len(v.Action) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Action"))
	}
	if len(v.Role) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Role"))
	}
	if v.Users == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Users"))
	} else if v.Users != nil {
		if err := validateUserList(v.Users); err != nil {
			invalidParams.AddNested("Users", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateUpdateInstructionBatch(v []types.UpdateInstruction) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateInstructionBatch"}
	for i := range v {
		if err := validateUpdateInstruction(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateUser(v *types.User) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "User"}
	if v.Id == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Id"))
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

func validateUserList(v []types.User) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UserList"}
	for i := range v {
		if err := validateUser(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateVpcConfiguration(v *types.VpcConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "VpcConfiguration"}
	if v.SecurityGroupIds == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SecurityGroupIds"))
	}
	if v.SubnetIds == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SubnetIds"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpAssociateLicenseInput(v *AssociateLicenseInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AssociateLicenseInput"}
	if v.WorkspaceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("WorkspaceId"))
	}
	if len(v.LicenseType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("LicenseType"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateWorkspaceApiKeyInput(v *CreateWorkspaceApiKeyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateWorkspaceApiKeyInput"}
	if v.KeyName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("KeyName"))
	}
	if v.KeyRole == nil {
		invalidParams.Add(smithy.NewErrParamRequired("KeyRole"))
	}
	if v.SecondsToLive == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SecondsToLive"))
	}
	if v.WorkspaceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("WorkspaceId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateWorkspaceInput(v *CreateWorkspaceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateWorkspaceInput"}
	if len(v.AccountAccessType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("AccountAccessType"))
	}
	if len(v.PermissionType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("PermissionType"))
	}
	if v.AuthenticationProviders == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AuthenticationProviders"))
	}
	if v.VpcConfiguration != nil {
		if err := validateVpcConfiguration(v.VpcConfiguration); err != nil {
			invalidParams.AddNested("VpcConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if v.NetworkAccessControl != nil {
		if err := validateNetworkAccessConfiguration(v.NetworkAccessControl); err != nil {
			invalidParams.AddNested("NetworkAccessControl", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteWorkspaceApiKeyInput(v *DeleteWorkspaceApiKeyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteWorkspaceApiKeyInput"}
	if v.KeyName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("KeyName"))
	}
	if v.WorkspaceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("WorkspaceId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteWorkspaceInput(v *DeleteWorkspaceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteWorkspaceInput"}
	if v.WorkspaceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("WorkspaceId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeWorkspaceAuthenticationInput(v *DescribeWorkspaceAuthenticationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeWorkspaceAuthenticationInput"}
	if v.WorkspaceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("WorkspaceId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeWorkspaceConfigurationInput(v *DescribeWorkspaceConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeWorkspaceConfigurationInput"}
	if v.WorkspaceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("WorkspaceId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeWorkspaceInput(v *DescribeWorkspaceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeWorkspaceInput"}
	if v.WorkspaceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("WorkspaceId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDisassociateLicenseInput(v *DisassociateLicenseInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DisassociateLicenseInput"}
	if v.WorkspaceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("WorkspaceId"))
	}
	if len(v.LicenseType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("LicenseType"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListPermissionsInput(v *ListPermissionsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListPermissionsInput"}
	if v.WorkspaceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("WorkspaceId"))
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

func validateOpUpdatePermissionsInput(v *UpdatePermissionsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdatePermissionsInput"}
	if v.UpdateInstructionBatch == nil {
		invalidParams.Add(smithy.NewErrParamRequired("UpdateInstructionBatch"))
	} else if v.UpdateInstructionBatch != nil {
		if err := validateUpdateInstructionBatch(v.UpdateInstructionBatch); err != nil {
			invalidParams.AddNested("UpdateInstructionBatch", err.(smithy.InvalidParamsError))
		}
	}
	if v.WorkspaceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("WorkspaceId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateWorkspaceAuthenticationInput(v *UpdateWorkspaceAuthenticationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateWorkspaceAuthenticationInput"}
	if v.WorkspaceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("WorkspaceId"))
	}
	if v.AuthenticationProviders == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AuthenticationProviders"))
	}
	if v.SamlConfiguration != nil {
		if err := validateSamlConfiguration(v.SamlConfiguration); err != nil {
			invalidParams.AddNested("SamlConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateWorkspaceConfigurationInput(v *UpdateWorkspaceConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateWorkspaceConfigurationInput"}
	if v.Configuration == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Configuration"))
	}
	if v.WorkspaceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("WorkspaceId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateWorkspaceInput(v *UpdateWorkspaceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateWorkspaceInput"}
	if v.WorkspaceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("WorkspaceId"))
	}
	if v.VpcConfiguration != nil {
		if err := validateVpcConfiguration(v.VpcConfiguration); err != nil {
			invalidParams.AddNested("VpcConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if v.NetworkAccessControl != nil {
		if err := validateNetworkAccessConfiguration(v.NetworkAccessControl); err != nil {
			invalidParams.AddNested("NetworkAccessControl", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
