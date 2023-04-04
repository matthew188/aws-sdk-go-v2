// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitosync

import (
	"context"
	"fmt"
	"github.com/matthew188/aws-sdk-go-v2/service/cognitosync/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpBulkPublish struct {
}

func (*validateOpBulkPublish) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpBulkPublish) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*BulkPublishInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpBulkPublishInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteDataset struct {
}

func (*validateOpDeleteDataset) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteDataset) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteDatasetInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteDatasetInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeDataset struct {
}

func (*validateOpDescribeDataset) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeDataset) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeDatasetInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeDatasetInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeIdentityPoolUsage struct {
}

func (*validateOpDescribeIdentityPoolUsage) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeIdentityPoolUsage) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeIdentityPoolUsageInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeIdentityPoolUsageInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeIdentityUsage struct {
}

func (*validateOpDescribeIdentityUsage) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeIdentityUsage) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeIdentityUsageInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeIdentityUsageInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetBulkPublishDetails struct {
}

func (*validateOpGetBulkPublishDetails) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetBulkPublishDetails) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetBulkPublishDetailsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetBulkPublishDetailsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetCognitoEvents struct {
}

func (*validateOpGetCognitoEvents) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetCognitoEvents) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetCognitoEventsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetCognitoEventsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetIdentityPoolConfiguration struct {
}

func (*validateOpGetIdentityPoolConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetIdentityPoolConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetIdentityPoolConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetIdentityPoolConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListDatasets struct {
}

func (*validateOpListDatasets) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListDatasets) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListDatasetsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListDatasetsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListRecords struct {
}

func (*validateOpListRecords) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListRecords) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListRecordsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListRecordsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpRegisterDevice struct {
}

func (*validateOpRegisterDevice) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpRegisterDevice) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*RegisterDeviceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpRegisterDeviceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpSetCognitoEvents struct {
}

func (*validateOpSetCognitoEvents) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpSetCognitoEvents) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*SetCognitoEventsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpSetCognitoEventsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpSetIdentityPoolConfiguration struct {
}

func (*validateOpSetIdentityPoolConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpSetIdentityPoolConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*SetIdentityPoolConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpSetIdentityPoolConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpSubscribeToDataset struct {
}

func (*validateOpSubscribeToDataset) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpSubscribeToDataset) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*SubscribeToDatasetInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpSubscribeToDatasetInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUnsubscribeFromDataset struct {
}

func (*validateOpUnsubscribeFromDataset) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUnsubscribeFromDataset) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UnsubscribeFromDatasetInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUnsubscribeFromDatasetInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateRecords struct {
}

func (*validateOpUpdateRecords) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateRecords) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateRecordsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateRecordsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpBulkPublishValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpBulkPublish{}, middleware.After)
}

func addOpDeleteDatasetValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteDataset{}, middleware.After)
}

func addOpDescribeDatasetValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeDataset{}, middleware.After)
}

func addOpDescribeIdentityPoolUsageValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeIdentityPoolUsage{}, middleware.After)
}

func addOpDescribeIdentityUsageValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeIdentityUsage{}, middleware.After)
}

func addOpGetBulkPublishDetailsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetBulkPublishDetails{}, middleware.After)
}

func addOpGetCognitoEventsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetCognitoEvents{}, middleware.After)
}

func addOpGetIdentityPoolConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetIdentityPoolConfiguration{}, middleware.After)
}

func addOpListDatasetsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListDatasets{}, middleware.After)
}

func addOpListRecordsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListRecords{}, middleware.After)
}

func addOpRegisterDeviceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpRegisterDevice{}, middleware.After)
}

func addOpSetCognitoEventsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpSetCognitoEvents{}, middleware.After)
}

func addOpSetIdentityPoolConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpSetIdentityPoolConfiguration{}, middleware.After)
}

func addOpSubscribeToDatasetValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpSubscribeToDataset{}, middleware.After)
}

func addOpUnsubscribeFromDatasetValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUnsubscribeFromDataset{}, middleware.After)
}

func addOpUpdateRecordsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateRecords{}, middleware.After)
}

func validateRecordPatch(v *types.RecordPatch) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RecordPatch"}
	if len(v.Op) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Op"))
	}
	if v.Key == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Key"))
	}
	if v.SyncCount == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SyncCount"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateRecordPatchList(v []types.RecordPatch) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RecordPatchList"}
	for i := range v {
		if err := validateRecordPatch(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpBulkPublishInput(v *BulkPublishInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "BulkPublishInput"}
	if v.IdentityPoolId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdentityPoolId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteDatasetInput(v *DeleteDatasetInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteDatasetInput"}
	if v.IdentityPoolId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdentityPoolId"))
	}
	if v.IdentityId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdentityId"))
	}
	if v.DatasetName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DatasetName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeDatasetInput(v *DescribeDatasetInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeDatasetInput"}
	if v.IdentityPoolId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdentityPoolId"))
	}
	if v.IdentityId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdentityId"))
	}
	if v.DatasetName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DatasetName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeIdentityPoolUsageInput(v *DescribeIdentityPoolUsageInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeIdentityPoolUsageInput"}
	if v.IdentityPoolId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdentityPoolId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeIdentityUsageInput(v *DescribeIdentityUsageInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeIdentityUsageInput"}
	if v.IdentityPoolId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdentityPoolId"))
	}
	if v.IdentityId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdentityId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetBulkPublishDetailsInput(v *GetBulkPublishDetailsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetBulkPublishDetailsInput"}
	if v.IdentityPoolId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdentityPoolId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetCognitoEventsInput(v *GetCognitoEventsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetCognitoEventsInput"}
	if v.IdentityPoolId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdentityPoolId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetIdentityPoolConfigurationInput(v *GetIdentityPoolConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetIdentityPoolConfigurationInput"}
	if v.IdentityPoolId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdentityPoolId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListDatasetsInput(v *ListDatasetsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListDatasetsInput"}
	if v.IdentityPoolId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdentityPoolId"))
	}
	if v.IdentityId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdentityId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListRecordsInput(v *ListRecordsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListRecordsInput"}
	if v.IdentityPoolId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdentityPoolId"))
	}
	if v.IdentityId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdentityId"))
	}
	if v.DatasetName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DatasetName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpRegisterDeviceInput(v *RegisterDeviceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RegisterDeviceInput"}
	if v.IdentityPoolId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdentityPoolId"))
	}
	if v.IdentityId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdentityId"))
	}
	if len(v.Platform) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Platform"))
	}
	if v.Token == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Token"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpSetCognitoEventsInput(v *SetCognitoEventsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SetCognitoEventsInput"}
	if v.IdentityPoolId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdentityPoolId"))
	}
	if v.Events == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Events"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpSetIdentityPoolConfigurationInput(v *SetIdentityPoolConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SetIdentityPoolConfigurationInput"}
	if v.IdentityPoolId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdentityPoolId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpSubscribeToDatasetInput(v *SubscribeToDatasetInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SubscribeToDatasetInput"}
	if v.IdentityPoolId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdentityPoolId"))
	}
	if v.IdentityId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdentityId"))
	}
	if v.DatasetName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DatasetName"))
	}
	if v.DeviceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DeviceId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUnsubscribeFromDatasetInput(v *UnsubscribeFromDatasetInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UnsubscribeFromDatasetInput"}
	if v.IdentityPoolId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdentityPoolId"))
	}
	if v.IdentityId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdentityId"))
	}
	if v.DatasetName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DatasetName"))
	}
	if v.DeviceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DeviceId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateRecordsInput(v *UpdateRecordsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateRecordsInput"}
	if v.IdentityPoolId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdentityPoolId"))
	}
	if v.IdentityId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IdentityId"))
	}
	if v.DatasetName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DatasetName"))
	}
	if v.RecordPatches != nil {
		if err := validateRecordPatchList(v.RecordPatches); err != nil {
			invalidParams.AddNested("RecordPatches", err.(smithy.InvalidParamsError))
		}
	}
	if v.SyncSessionToken == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SyncSessionToken"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
