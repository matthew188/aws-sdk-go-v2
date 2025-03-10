// Code generated by smithy-go-codegen DO NOT EDIT.

package machinelearning

import (
	"context"
	"fmt"
	"github.com/matthew188/aws-sdk-go-v2/service/machinelearning/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpAddTags struct {
}

func (*validateOpAddTags) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpAddTags) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*AddTagsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpAddTagsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateBatchPrediction struct {
}

func (*validateOpCreateBatchPrediction) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateBatchPrediction) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateBatchPredictionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateBatchPredictionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateDataSourceFromRDS struct {
}

func (*validateOpCreateDataSourceFromRDS) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateDataSourceFromRDS) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateDataSourceFromRDSInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateDataSourceFromRDSInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateDataSourceFromRedshift struct {
}

func (*validateOpCreateDataSourceFromRedshift) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateDataSourceFromRedshift) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateDataSourceFromRedshiftInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateDataSourceFromRedshiftInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateDataSourceFromS3 struct {
}

func (*validateOpCreateDataSourceFromS3) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateDataSourceFromS3) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateDataSourceFromS3Input)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateDataSourceFromS3Input(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateEvaluation struct {
}

func (*validateOpCreateEvaluation) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateEvaluation) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateEvaluationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateEvaluationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateMLModel struct {
}

func (*validateOpCreateMLModel) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateMLModel) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateMLModelInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateMLModelInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateRealtimeEndpoint struct {
}

func (*validateOpCreateRealtimeEndpoint) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateRealtimeEndpoint) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateRealtimeEndpointInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateRealtimeEndpointInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteBatchPrediction struct {
}

func (*validateOpDeleteBatchPrediction) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteBatchPrediction) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteBatchPredictionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteBatchPredictionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteDataSource struct {
}

func (*validateOpDeleteDataSource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteDataSource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteDataSourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteDataSourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteEvaluation struct {
}

func (*validateOpDeleteEvaluation) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteEvaluation) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteEvaluationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteEvaluationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteMLModel struct {
}

func (*validateOpDeleteMLModel) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteMLModel) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteMLModelInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteMLModelInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteRealtimeEndpoint struct {
}

func (*validateOpDeleteRealtimeEndpoint) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteRealtimeEndpoint) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteRealtimeEndpointInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteRealtimeEndpointInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteTags struct {
}

func (*validateOpDeleteTags) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteTags) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteTagsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteTagsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeTags struct {
}

func (*validateOpDescribeTags) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeTags) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeTagsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeTagsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetBatchPrediction struct {
}

func (*validateOpGetBatchPrediction) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetBatchPrediction) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetBatchPredictionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetBatchPredictionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetDataSource struct {
}

func (*validateOpGetDataSource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetDataSource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetDataSourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetDataSourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetEvaluation struct {
}

func (*validateOpGetEvaluation) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetEvaluation) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetEvaluationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetEvaluationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetMLModel struct {
}

func (*validateOpGetMLModel) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetMLModel) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetMLModelInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetMLModelInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPredict struct {
}

func (*validateOpPredict) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPredict) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PredictInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPredictInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateBatchPrediction struct {
}

func (*validateOpUpdateBatchPrediction) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateBatchPrediction) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateBatchPredictionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateBatchPredictionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateDataSource struct {
}

func (*validateOpUpdateDataSource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateDataSource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateDataSourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateDataSourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateEvaluation struct {
}

func (*validateOpUpdateEvaluation) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateEvaluation) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateEvaluationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateEvaluationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateMLModel struct {
}

func (*validateOpUpdateMLModel) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateMLModel) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateMLModelInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateMLModelInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpAddTagsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpAddTags{}, middleware.After)
}

func addOpCreateBatchPredictionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateBatchPrediction{}, middleware.After)
}

func addOpCreateDataSourceFromRDSValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateDataSourceFromRDS{}, middleware.After)
}

func addOpCreateDataSourceFromRedshiftValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateDataSourceFromRedshift{}, middleware.After)
}

func addOpCreateDataSourceFromS3ValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateDataSourceFromS3{}, middleware.After)
}

func addOpCreateEvaluationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateEvaluation{}, middleware.After)
}

func addOpCreateMLModelValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateMLModel{}, middleware.After)
}

func addOpCreateRealtimeEndpointValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateRealtimeEndpoint{}, middleware.After)
}

func addOpDeleteBatchPredictionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteBatchPrediction{}, middleware.After)
}

func addOpDeleteDataSourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteDataSource{}, middleware.After)
}

func addOpDeleteEvaluationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteEvaluation{}, middleware.After)
}

func addOpDeleteMLModelValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteMLModel{}, middleware.After)
}

func addOpDeleteRealtimeEndpointValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteRealtimeEndpoint{}, middleware.After)
}

func addOpDeleteTagsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteTags{}, middleware.After)
}

func addOpDescribeTagsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeTags{}, middleware.After)
}

func addOpGetBatchPredictionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetBatchPrediction{}, middleware.After)
}

func addOpGetDataSourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetDataSource{}, middleware.After)
}

func addOpGetEvaluationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetEvaluation{}, middleware.After)
}

func addOpGetMLModelValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetMLModel{}, middleware.After)
}

func addOpPredictValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPredict{}, middleware.After)
}

func addOpUpdateBatchPredictionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateBatchPrediction{}, middleware.After)
}

func addOpUpdateDataSourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateDataSource{}, middleware.After)
}

func addOpUpdateEvaluationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateEvaluation{}, middleware.After)
}

func addOpUpdateMLModelValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateMLModel{}, middleware.After)
}

func validateRDSDatabase(v *types.RDSDatabase) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RDSDatabase"}
	if v.InstanceIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("InstanceIdentifier"))
	}
	if v.DatabaseName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DatabaseName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateRDSDatabaseCredentials(v *types.RDSDatabaseCredentials) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RDSDatabaseCredentials"}
	if v.Username == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Username"))
	}
	if v.Password == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Password"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateRDSDataSpec(v *types.RDSDataSpec) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RDSDataSpec"}
	if v.DatabaseInformation == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DatabaseInformation"))
	} else if v.DatabaseInformation != nil {
		if err := validateRDSDatabase(v.DatabaseInformation); err != nil {
			invalidParams.AddNested("DatabaseInformation", err.(smithy.InvalidParamsError))
		}
	}
	if v.SelectSqlQuery == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SelectSqlQuery"))
	}
	if v.DatabaseCredentials == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DatabaseCredentials"))
	} else if v.DatabaseCredentials != nil {
		if err := validateRDSDatabaseCredentials(v.DatabaseCredentials); err != nil {
			invalidParams.AddNested("DatabaseCredentials", err.(smithy.InvalidParamsError))
		}
	}
	if v.S3StagingLocation == nil {
		invalidParams.Add(smithy.NewErrParamRequired("S3StagingLocation"))
	}
	if v.ResourceRole == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceRole"))
	}
	if v.ServiceRole == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ServiceRole"))
	}
	if v.SubnetId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SubnetId"))
	}
	if v.SecurityGroupIds == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SecurityGroupIds"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateRedshiftDatabase(v *types.RedshiftDatabase) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RedshiftDatabase"}
	if v.DatabaseName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DatabaseName"))
	}
	if v.ClusterIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClusterIdentifier"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateRedshiftDatabaseCredentials(v *types.RedshiftDatabaseCredentials) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RedshiftDatabaseCredentials"}
	if v.Username == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Username"))
	}
	if v.Password == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Password"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateRedshiftDataSpec(v *types.RedshiftDataSpec) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RedshiftDataSpec"}
	if v.DatabaseInformation == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DatabaseInformation"))
	} else if v.DatabaseInformation != nil {
		if err := validateRedshiftDatabase(v.DatabaseInformation); err != nil {
			invalidParams.AddNested("DatabaseInformation", err.(smithy.InvalidParamsError))
		}
	}
	if v.SelectSqlQuery == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SelectSqlQuery"))
	}
	if v.DatabaseCredentials == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DatabaseCredentials"))
	} else if v.DatabaseCredentials != nil {
		if err := validateRedshiftDatabaseCredentials(v.DatabaseCredentials); err != nil {
			invalidParams.AddNested("DatabaseCredentials", err.(smithy.InvalidParamsError))
		}
	}
	if v.S3StagingLocation == nil {
		invalidParams.Add(smithy.NewErrParamRequired("S3StagingLocation"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateS3DataSpec(v *types.S3DataSpec) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "S3DataSpec"}
	if v.DataLocationS3 == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DataLocationS3"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpAddTagsInput(v *AddTagsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AddTagsInput"}
	if v.Tags == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Tags"))
	}
	if v.ResourceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceId"))
	}
	if len(v.ResourceType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceType"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateBatchPredictionInput(v *CreateBatchPredictionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateBatchPredictionInput"}
	if v.BatchPredictionId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BatchPredictionId"))
	}
	if v.MLModelId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MLModelId"))
	}
	if v.BatchPredictionDataSourceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BatchPredictionDataSourceId"))
	}
	if v.OutputUri == nil {
		invalidParams.Add(smithy.NewErrParamRequired("OutputUri"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateDataSourceFromRDSInput(v *CreateDataSourceFromRDSInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateDataSourceFromRDSInput"}
	if v.DataSourceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DataSourceId"))
	}
	if v.RDSData == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RDSData"))
	} else if v.RDSData != nil {
		if err := validateRDSDataSpec(v.RDSData); err != nil {
			invalidParams.AddNested("RDSData", err.(smithy.InvalidParamsError))
		}
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

func validateOpCreateDataSourceFromRedshiftInput(v *CreateDataSourceFromRedshiftInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateDataSourceFromRedshiftInput"}
	if v.DataSourceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DataSourceId"))
	}
	if v.DataSpec == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DataSpec"))
	} else if v.DataSpec != nil {
		if err := validateRedshiftDataSpec(v.DataSpec); err != nil {
			invalidParams.AddNested("DataSpec", err.(smithy.InvalidParamsError))
		}
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

func validateOpCreateDataSourceFromS3Input(v *CreateDataSourceFromS3Input) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateDataSourceFromS3Input"}
	if v.DataSourceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DataSourceId"))
	}
	if v.DataSpec == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DataSpec"))
	} else if v.DataSpec != nil {
		if err := validateS3DataSpec(v.DataSpec); err != nil {
			invalidParams.AddNested("DataSpec", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateEvaluationInput(v *CreateEvaluationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateEvaluationInput"}
	if v.EvaluationId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EvaluationId"))
	}
	if v.MLModelId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MLModelId"))
	}
	if v.EvaluationDataSourceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EvaluationDataSourceId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateMLModelInput(v *CreateMLModelInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateMLModelInput"}
	if v.MLModelId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MLModelId"))
	}
	if len(v.MLModelType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("MLModelType"))
	}
	if v.TrainingDataSourceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TrainingDataSourceId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateRealtimeEndpointInput(v *CreateRealtimeEndpointInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateRealtimeEndpointInput"}
	if v.MLModelId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MLModelId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteBatchPredictionInput(v *DeleteBatchPredictionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteBatchPredictionInput"}
	if v.BatchPredictionId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BatchPredictionId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteDataSourceInput(v *DeleteDataSourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteDataSourceInput"}
	if v.DataSourceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DataSourceId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteEvaluationInput(v *DeleteEvaluationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteEvaluationInput"}
	if v.EvaluationId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EvaluationId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteMLModelInput(v *DeleteMLModelInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteMLModelInput"}
	if v.MLModelId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MLModelId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteRealtimeEndpointInput(v *DeleteRealtimeEndpointInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteRealtimeEndpointInput"}
	if v.MLModelId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MLModelId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteTagsInput(v *DeleteTagsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteTagsInput"}
	if v.TagKeys == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TagKeys"))
	}
	if v.ResourceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceId"))
	}
	if len(v.ResourceType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceType"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeTagsInput(v *DescribeTagsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeTagsInput"}
	if v.ResourceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceId"))
	}
	if len(v.ResourceType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceType"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetBatchPredictionInput(v *GetBatchPredictionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetBatchPredictionInput"}
	if v.BatchPredictionId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BatchPredictionId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetDataSourceInput(v *GetDataSourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetDataSourceInput"}
	if v.DataSourceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DataSourceId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetEvaluationInput(v *GetEvaluationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetEvaluationInput"}
	if v.EvaluationId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EvaluationId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetMLModelInput(v *GetMLModelInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetMLModelInput"}
	if v.MLModelId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MLModelId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPredictInput(v *PredictInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PredictInput"}
	if v.MLModelId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MLModelId"))
	}
	if v.Record == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Record"))
	}
	if v.PredictEndpoint == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PredictEndpoint"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateBatchPredictionInput(v *UpdateBatchPredictionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateBatchPredictionInput"}
	if v.BatchPredictionId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BatchPredictionId"))
	}
	if v.BatchPredictionName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BatchPredictionName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateDataSourceInput(v *UpdateDataSourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateDataSourceInput"}
	if v.DataSourceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DataSourceId"))
	}
	if v.DataSourceName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DataSourceName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateEvaluationInput(v *UpdateEvaluationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateEvaluationInput"}
	if v.EvaluationId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EvaluationId"))
	}
	if v.EvaluationName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EvaluationName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateMLModelInput(v *UpdateMLModelInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateMLModelInput"}
	if v.MLModelId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MLModelId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
