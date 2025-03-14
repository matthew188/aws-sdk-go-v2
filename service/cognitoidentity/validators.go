// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentity

import (
	"context"
	"fmt"
	"github.com/matthew188/aws-sdk-go-v2/service/cognitoidentity/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpCreateIdentityPool struct {
}

func (*validateOpCreateIdentityPool) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateIdentityPool) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateIdentityPoolInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateIdentityPoolInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteIdentities struct {
}

func (*validateOpDeleteIdentities) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteIdentities) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteIdentitiesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteIdentitiesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteIdentityPool struct {
}

func (*validateOpDeleteIdentityPool) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteIdentityPool) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteIdentityPoolInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteIdentityPoolInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeIdentity struct {
}

func (*validateOpDescribeIdentity) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeIdentity) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeIdentityInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeIdentityInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeIdentityPool struct {
}

func (*validateOpDescribeIdentityPool) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeIdentityPool) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeIdentityPoolInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeIdentityPoolInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetCredentialsForIdentity struct {
}

func (*validateOpGetCredentialsForIdentity) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetCredentialsForIdentity) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetCredentialsForIdentityInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetCredentialsForIdentityInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetIdentityPoolRoles struct {
}

func (*validateOpGetIdentityPoolRoles) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetIdentityPoolRoles) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetIdentityPoolRolesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetIdentityPoolRolesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetId struct {
}

func (*validateOpGetId) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetId) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetIdInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetIdInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetOpenIdTokenForDeveloperIdentity struct {
}

func (*validateOpGetOpenIdTokenForDeveloperIdentity) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetOpenIdTokenForDeveloperIdentity) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetOpenIdTokenForDeveloperIdentityInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetOpenIdTokenForDeveloperIdentityInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetOpenIdToken struct {
}

func (*validateOpGetOpenIdToken) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetOpenIdToken) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetOpenIdTokenInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetOpenIdTokenInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetPrincipalTagAttributeMap struct {
}

func (*validateOpGetPrincipalTagAttributeMap) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetPrincipalTagAttributeMap) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetPrincipalTagAttributeMapInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetPrincipalTagAttributeMapInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListIdentities struct {
}

func (*validateOpListIdentities) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListIdentities) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListIdentitiesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListIdentitiesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListIdentityPools struct {
}

func (*validateOpListIdentityPools) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListIdentityPools) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListIdentityPoolsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListIdentityPoolsInput(input); err != nil {
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

type validateOpLookupDeveloperIdentity struct {
}

func (*validateOpLookupDeveloperIdentity) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpLookupDeveloperIdentity) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*LookupDeveloperIdentityInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpLookupDeveloperIdentityInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpMergeDeveloperIdentities struct {
}

func (*validateOpMergeDeveloperIdentities) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpMergeDeveloperIdentities) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*MergeDeveloperIdentitiesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpMergeDeveloperIdentitiesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpSetIdentityPoolRoles struct {
}

func (*validateOpSetIdentityPoolRoles) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpSetIdentityPoolRoles) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*SetIdentityPoolRolesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpSetIdentityPoolRolesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpSetPrincipalTagAttributeMap struct {
}

func (*validateOpSetPrincipalTagAttributeMap) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpSetPrincipalTagAttributeMap) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*SetPrincipalTagAttributeMapInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpSetPrincipalTagAttributeMapInput(input); err != nil {
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

type validateOpUnlinkDeveloperIdentity struct {
}

func (*validateOpUnlinkDeveloperIdentity) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUnlinkDeveloperIdentity) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UnlinkDeveloperIdentityInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUnlinkDeveloperIdentityInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUnlinkIdentity struct {
}

func (*validateOpUnlinkIdentity) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUnlinkIdentity) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UnlinkIdentityInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUnlinkIdentityInput(input); err != nil {
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

type validateOpUpdateIdentityPool struct {
}

func (*validateOpUpdateIdentityPool) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateIdentityPool) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateIdentityPoolInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateIdentityPoolInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpCreateIdentityPoolValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateIdentityPool{}, middleware.After)
}

func addOpDeleteIdentitiesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteIdentities{}, middleware.After)
}

func addOpDeleteIdentityPoolValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteIdentityPool{}, middleware.After)
}

func addOpDescribeIdentityValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeIdentity{}, middleware.After)
}

func addOpDescribeIdentityPoolValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeIdentityPool{}, middleware.After)
}

func addOpGetCredentialsForIdentityValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetCredentialsForIdentity{}, middleware.After)
}

func addOpGetIdentityPoolRolesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetIdentityPoolRoles{}, middleware.After)
}

func addOpGetIdValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetId{}, middleware.After)
}

func addOpGetOpenIdTokenForDeveloperIdentityValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetOpenIdTokenForDeveloperIdentity{}, middleware.After)
}

func addOpGetOpenIdTokenValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetOpenIdToken{}, middleware.After)
}

func addOpGetPrincipalTagAttributeMapValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetPrincipalTagAttributeMap{}, middleware.After)
}

func addOpListIdentitiesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListIdentities{}, middleware.After)
}

func addOpListIdentityPoolsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListIdentityPools{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpLookupDeveloperIdentityValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpLookupDeveloperIdentity{}, middleware.After)
}

func addOpMergeDeveloperIdentitiesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpMergeDeveloperIdentities{}, middleware.After)
}

func addOpSetIdentityPoolRolesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpSetIdentityPoolRoles{}, middleware.After)
}

func addOpSetPrincipalTagAttributeMapValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpSetPrincipalTagAttributeMap{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUnlinkDeveloperIdentityValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUnlinkDeveloperIdentity{}, middleware.After)
}

func addOpUnlinkIdentityValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUnlinkIdentity{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func addOpUpdateIdentityPoolValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateIdentityPool{}, middleware.After)
}

func validateMappingRule(v *types.MappingRule) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "MappingRule"}
	if v.Claim == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Claim"))
	}
	if len(v.MatchType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("MatchType"))
	}
	if v.Value == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Value"))
	}
	if v.RoleARN == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RoleARN"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateMappingRulesList(v []types.MappingRule) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "MappingRulesList"}
	for i := range v {
		if err := validateMappingRule(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateRoleMapping(v *types.RoleMapping) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RoleMapping"}
	if len(v.Type) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Type"))
	}
	if v.RulesConfiguration != nil {
		if err := validateRulesConfigurationType(v.RulesConfiguration); err != nil {
			invalidParams.AddNested("RulesConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateRoleMappingMap(v map[string]types.RoleMapping) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RoleMappingMap"}
	for key := range v {
		value := v[key]
		if err := validateRoleMapping(&value); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%q]", key), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateRulesConfigurationType(v *types.RulesConfigurationType) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RulesConfigurationType"}
	if v.Rules == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Rules"))
	} else if v.Rules != nil {
		if err := validateMappingRulesList(v.Rules); err != nil {
			invalidParams.AddNested("Rules", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateIdentityPoolInput(v *CreateIdentityPoolInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateIdentityPoolInput"}
	if v.IdentityPoolName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdentityPoolName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteIdentitiesInput(v *DeleteIdentitiesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteIdentitiesInput"}
	if v.IdentityIdsToDelete == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdentityIdsToDelete"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteIdentityPoolInput(v *DeleteIdentityPoolInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteIdentityPoolInput"}
	if v.IdentityPoolId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdentityPoolId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeIdentityInput(v *DescribeIdentityInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeIdentityInput"}
	if v.IdentityId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdentityId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeIdentityPoolInput(v *DescribeIdentityPoolInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeIdentityPoolInput"}
	if v.IdentityPoolId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdentityPoolId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetCredentialsForIdentityInput(v *GetCredentialsForIdentityInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetCredentialsForIdentityInput"}
	if v.IdentityId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdentityId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetIdentityPoolRolesInput(v *GetIdentityPoolRolesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetIdentityPoolRolesInput"}
	if v.IdentityPoolId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdentityPoolId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetIdInput(v *GetIdInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetIdInput"}
	if v.IdentityPoolId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdentityPoolId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetOpenIdTokenForDeveloperIdentityInput(v *GetOpenIdTokenForDeveloperIdentityInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetOpenIdTokenForDeveloperIdentityInput"}
	if v.IdentityPoolId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdentityPoolId"))
	}
	if v.Logins == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Logins"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetOpenIdTokenInput(v *GetOpenIdTokenInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetOpenIdTokenInput"}
	if v.IdentityId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdentityId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetPrincipalTagAttributeMapInput(v *GetPrincipalTagAttributeMapInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetPrincipalTagAttributeMapInput"}
	if v.IdentityPoolId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdentityPoolId"))
	}
	if v.IdentityProviderName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdentityProviderName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListIdentitiesInput(v *ListIdentitiesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListIdentitiesInput"}
	if v.IdentityPoolId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdentityPoolId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListIdentityPoolsInput(v *ListIdentityPoolsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListIdentityPoolsInput"}
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

func validateOpLookupDeveloperIdentityInput(v *LookupDeveloperIdentityInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "LookupDeveloperIdentityInput"}
	if v.IdentityPoolId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdentityPoolId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpMergeDeveloperIdentitiesInput(v *MergeDeveloperIdentitiesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "MergeDeveloperIdentitiesInput"}
	if v.SourceUserIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SourceUserIdentifier"))
	}
	if v.DestinationUserIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DestinationUserIdentifier"))
	}
	if v.DeveloperProviderName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DeveloperProviderName"))
	}
	if v.IdentityPoolId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdentityPoolId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpSetIdentityPoolRolesInput(v *SetIdentityPoolRolesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SetIdentityPoolRolesInput"}
	if v.IdentityPoolId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdentityPoolId"))
	}
	if v.Roles == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Roles"))
	}
	if v.RoleMappings != nil {
		if err := validateRoleMappingMap(v.RoleMappings); err != nil {
			invalidParams.AddNested("RoleMappings", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpSetPrincipalTagAttributeMapInput(v *SetPrincipalTagAttributeMapInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SetPrincipalTagAttributeMapInput"}
	if v.IdentityPoolId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdentityPoolId"))
	}
	if v.IdentityProviderName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdentityProviderName"))
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

func validateOpUnlinkDeveloperIdentityInput(v *UnlinkDeveloperIdentityInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UnlinkDeveloperIdentityInput"}
	if v.IdentityId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdentityId"))
	}
	if v.IdentityPoolId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdentityPoolId"))
	}
	if v.DeveloperProviderName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DeveloperProviderName"))
	}
	if v.DeveloperUserIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DeveloperUserIdentifier"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUnlinkIdentityInput(v *UnlinkIdentityInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UnlinkIdentityInput"}
	if v.IdentityId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdentityId"))
	}
	if v.Logins == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Logins"))
	}
	if v.LoginsToRemove == nil {
		invalidParams.Add(smithy.NewErrParamRequired("LoginsToRemove"))
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

func validateOpUpdateIdentityPoolInput(v *UpdateIdentityPoolInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateIdentityPoolInput"}
	if v.IdentityPoolId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdentityPoolId"))
	}
	if v.IdentityPoolName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdentityPoolName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
