// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudsearch

import (
	"context"
	"fmt"
	"github.com/matthew188/aws-sdk-go-v2/service/cloudsearch/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpBuildSuggesters struct {
}

func (*validateOpBuildSuggesters) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpBuildSuggesters) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*BuildSuggestersInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpBuildSuggestersInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateDomain struct {
}

func (*validateOpCreateDomain) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateDomain) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateDomainInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateDomainInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDefineAnalysisScheme struct {
}

func (*validateOpDefineAnalysisScheme) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDefineAnalysisScheme) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DefineAnalysisSchemeInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDefineAnalysisSchemeInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDefineExpression struct {
}

func (*validateOpDefineExpression) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDefineExpression) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DefineExpressionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDefineExpressionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDefineIndexField struct {
}

func (*validateOpDefineIndexField) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDefineIndexField) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DefineIndexFieldInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDefineIndexFieldInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDefineSuggester struct {
}

func (*validateOpDefineSuggester) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDefineSuggester) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DefineSuggesterInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDefineSuggesterInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteAnalysisScheme struct {
}

func (*validateOpDeleteAnalysisScheme) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteAnalysisScheme) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteAnalysisSchemeInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteAnalysisSchemeInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteDomain struct {
}

func (*validateOpDeleteDomain) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteDomain) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteDomainInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteDomainInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteExpression struct {
}

func (*validateOpDeleteExpression) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteExpression) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteExpressionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteExpressionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteIndexField struct {
}

func (*validateOpDeleteIndexField) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteIndexField) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteIndexFieldInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteIndexFieldInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteSuggester struct {
}

func (*validateOpDeleteSuggester) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteSuggester) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteSuggesterInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteSuggesterInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeAnalysisSchemes struct {
}

func (*validateOpDescribeAnalysisSchemes) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeAnalysisSchemes) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeAnalysisSchemesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeAnalysisSchemesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeAvailabilityOptions struct {
}

func (*validateOpDescribeAvailabilityOptions) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeAvailabilityOptions) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeAvailabilityOptionsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeAvailabilityOptionsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeDomainEndpointOptions struct {
}

func (*validateOpDescribeDomainEndpointOptions) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeDomainEndpointOptions) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeDomainEndpointOptionsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeDomainEndpointOptionsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeExpressions struct {
}

func (*validateOpDescribeExpressions) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeExpressions) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeExpressionsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeExpressionsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeIndexFields struct {
}

func (*validateOpDescribeIndexFields) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeIndexFields) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeIndexFieldsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeIndexFieldsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeScalingParameters struct {
}

func (*validateOpDescribeScalingParameters) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeScalingParameters) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeScalingParametersInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeScalingParametersInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeServiceAccessPolicies struct {
}

func (*validateOpDescribeServiceAccessPolicies) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeServiceAccessPolicies) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeServiceAccessPoliciesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeServiceAccessPoliciesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeSuggesters struct {
}

func (*validateOpDescribeSuggesters) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeSuggesters) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeSuggestersInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeSuggestersInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpIndexDocuments struct {
}

func (*validateOpIndexDocuments) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpIndexDocuments) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*IndexDocumentsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpIndexDocumentsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateAvailabilityOptions struct {
}

func (*validateOpUpdateAvailabilityOptions) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateAvailabilityOptions) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateAvailabilityOptionsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateAvailabilityOptionsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateDomainEndpointOptions struct {
}

func (*validateOpUpdateDomainEndpointOptions) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateDomainEndpointOptions) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateDomainEndpointOptionsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateDomainEndpointOptionsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateScalingParameters struct {
}

func (*validateOpUpdateScalingParameters) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateScalingParameters) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateScalingParametersInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateScalingParametersInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateServiceAccessPolicies struct {
}

func (*validateOpUpdateServiceAccessPolicies) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateServiceAccessPolicies) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateServiceAccessPoliciesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateServiceAccessPoliciesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpBuildSuggestersValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpBuildSuggesters{}, middleware.After)
}

func addOpCreateDomainValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateDomain{}, middleware.After)
}

func addOpDefineAnalysisSchemeValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDefineAnalysisScheme{}, middleware.After)
}

func addOpDefineExpressionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDefineExpression{}, middleware.After)
}

func addOpDefineIndexFieldValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDefineIndexField{}, middleware.After)
}

func addOpDefineSuggesterValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDefineSuggester{}, middleware.After)
}

func addOpDeleteAnalysisSchemeValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteAnalysisScheme{}, middleware.After)
}

func addOpDeleteDomainValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteDomain{}, middleware.After)
}

func addOpDeleteExpressionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteExpression{}, middleware.After)
}

func addOpDeleteIndexFieldValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteIndexField{}, middleware.After)
}

func addOpDeleteSuggesterValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteSuggester{}, middleware.After)
}

func addOpDescribeAnalysisSchemesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeAnalysisSchemes{}, middleware.After)
}

func addOpDescribeAvailabilityOptionsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeAvailabilityOptions{}, middleware.After)
}

func addOpDescribeDomainEndpointOptionsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeDomainEndpointOptions{}, middleware.After)
}

func addOpDescribeExpressionsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeExpressions{}, middleware.After)
}

func addOpDescribeIndexFieldsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeIndexFields{}, middleware.After)
}

func addOpDescribeScalingParametersValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeScalingParameters{}, middleware.After)
}

func addOpDescribeServiceAccessPoliciesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeServiceAccessPolicies{}, middleware.After)
}

func addOpDescribeSuggestersValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeSuggesters{}, middleware.After)
}

func addOpIndexDocumentsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpIndexDocuments{}, middleware.After)
}

func addOpUpdateAvailabilityOptionsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateAvailabilityOptions{}, middleware.After)
}

func addOpUpdateDomainEndpointOptionsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateDomainEndpointOptions{}, middleware.After)
}

func addOpUpdateScalingParametersValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateScalingParameters{}, middleware.After)
}

func addOpUpdateServiceAccessPoliciesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateServiceAccessPolicies{}, middleware.After)
}

func validateAnalysisScheme(v *types.AnalysisScheme) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AnalysisScheme"}
	if v.AnalysisSchemeName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AnalysisSchemeName"))
	}
	if len(v.AnalysisSchemeLanguage) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("AnalysisSchemeLanguage"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateDocumentSuggesterOptions(v *types.DocumentSuggesterOptions) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DocumentSuggesterOptions"}
	if v.SourceField == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SourceField"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateExpression(v *types.Expression) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Expression"}
	if v.ExpressionName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ExpressionName"))
	}
	if v.ExpressionValue == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ExpressionValue"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateIndexField(v *types.IndexField) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "IndexField"}
	if v.IndexFieldName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IndexFieldName"))
	}
	if len(v.IndexFieldType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("IndexFieldType"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateSuggester(v *types.Suggester) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Suggester"}
	if v.SuggesterName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SuggesterName"))
	}
	if v.DocumentSuggesterOptions == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DocumentSuggesterOptions"))
	} else if v.DocumentSuggesterOptions != nil {
		if err := validateDocumentSuggesterOptions(v.DocumentSuggesterOptions); err != nil {
			invalidParams.AddNested("DocumentSuggesterOptions", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpBuildSuggestersInput(v *BuildSuggestersInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "BuildSuggestersInput"}
	if v.DomainName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DomainName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateDomainInput(v *CreateDomainInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateDomainInput"}
	if v.DomainName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DomainName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDefineAnalysisSchemeInput(v *DefineAnalysisSchemeInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DefineAnalysisSchemeInput"}
	if v.DomainName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DomainName"))
	}
	if v.AnalysisScheme == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AnalysisScheme"))
	} else if v.AnalysisScheme != nil {
		if err := validateAnalysisScheme(v.AnalysisScheme); err != nil {
			invalidParams.AddNested("AnalysisScheme", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDefineExpressionInput(v *DefineExpressionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DefineExpressionInput"}
	if v.DomainName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DomainName"))
	}
	if v.Expression == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Expression"))
	} else if v.Expression != nil {
		if err := validateExpression(v.Expression); err != nil {
			invalidParams.AddNested("Expression", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDefineIndexFieldInput(v *DefineIndexFieldInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DefineIndexFieldInput"}
	if v.DomainName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DomainName"))
	}
	if v.IndexField == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IndexField"))
	} else if v.IndexField != nil {
		if err := validateIndexField(v.IndexField); err != nil {
			invalidParams.AddNested("IndexField", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDefineSuggesterInput(v *DefineSuggesterInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DefineSuggesterInput"}
	if v.DomainName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DomainName"))
	}
	if v.Suggester == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Suggester"))
	} else if v.Suggester != nil {
		if err := validateSuggester(v.Suggester); err != nil {
			invalidParams.AddNested("Suggester", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteAnalysisSchemeInput(v *DeleteAnalysisSchemeInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteAnalysisSchemeInput"}
	if v.DomainName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DomainName"))
	}
	if v.AnalysisSchemeName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AnalysisSchemeName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteDomainInput(v *DeleteDomainInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteDomainInput"}
	if v.DomainName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DomainName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteExpressionInput(v *DeleteExpressionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteExpressionInput"}
	if v.DomainName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DomainName"))
	}
	if v.ExpressionName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ExpressionName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteIndexFieldInput(v *DeleteIndexFieldInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteIndexFieldInput"}
	if v.DomainName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DomainName"))
	}
	if v.IndexFieldName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IndexFieldName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteSuggesterInput(v *DeleteSuggesterInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteSuggesterInput"}
	if v.DomainName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DomainName"))
	}
	if v.SuggesterName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SuggesterName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeAnalysisSchemesInput(v *DescribeAnalysisSchemesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeAnalysisSchemesInput"}
	if v.DomainName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DomainName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeAvailabilityOptionsInput(v *DescribeAvailabilityOptionsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeAvailabilityOptionsInput"}
	if v.DomainName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DomainName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeDomainEndpointOptionsInput(v *DescribeDomainEndpointOptionsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeDomainEndpointOptionsInput"}
	if v.DomainName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DomainName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeExpressionsInput(v *DescribeExpressionsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeExpressionsInput"}
	if v.DomainName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DomainName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeIndexFieldsInput(v *DescribeIndexFieldsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeIndexFieldsInput"}
	if v.DomainName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DomainName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeScalingParametersInput(v *DescribeScalingParametersInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeScalingParametersInput"}
	if v.DomainName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DomainName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeServiceAccessPoliciesInput(v *DescribeServiceAccessPoliciesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeServiceAccessPoliciesInput"}
	if v.DomainName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DomainName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeSuggestersInput(v *DescribeSuggestersInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeSuggestersInput"}
	if v.DomainName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DomainName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpIndexDocumentsInput(v *IndexDocumentsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "IndexDocumentsInput"}
	if v.DomainName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DomainName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateAvailabilityOptionsInput(v *UpdateAvailabilityOptionsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateAvailabilityOptionsInput"}
	if v.DomainName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DomainName"))
	}
	if v.MultiAZ == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MultiAZ"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateDomainEndpointOptionsInput(v *UpdateDomainEndpointOptionsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateDomainEndpointOptionsInput"}
	if v.DomainName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DomainName"))
	}
	if v.DomainEndpointOptions == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DomainEndpointOptions"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateScalingParametersInput(v *UpdateScalingParametersInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateScalingParametersInput"}
	if v.DomainName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DomainName"))
	}
	if v.ScalingParameters == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ScalingParameters"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateServiceAccessPoliciesInput(v *UpdateServiceAccessPoliciesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateServiceAccessPoliciesInput"}
	if v.DomainName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DomainName"))
	}
	if v.AccessPolicies == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AccessPolicies"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
