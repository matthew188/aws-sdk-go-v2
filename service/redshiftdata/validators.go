// Code generated by smithy-go-codegen DO NOT EDIT.

package redshiftdata

import (
	"context"
	"fmt"
	"github.com/matthew188/aws-sdk-go-v2/service/redshiftdata/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpBatchExecuteStatement struct {
}

func (*validateOpBatchExecuteStatement) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpBatchExecuteStatement) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*BatchExecuteStatementInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpBatchExecuteStatementInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCancelStatement struct {
}

func (*validateOpCancelStatement) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCancelStatement) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CancelStatementInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCancelStatementInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeStatement struct {
}

func (*validateOpDescribeStatement) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeStatement) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeStatementInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeStatementInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeTable struct {
}

func (*validateOpDescribeTable) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeTable) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeTableInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeTableInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpExecuteStatement struct {
}

func (*validateOpExecuteStatement) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpExecuteStatement) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ExecuteStatementInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpExecuteStatementInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetStatementResult struct {
}

func (*validateOpGetStatementResult) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetStatementResult) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetStatementResultInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetStatementResultInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListDatabases struct {
}

func (*validateOpListDatabases) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListDatabases) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListDatabasesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListDatabasesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListSchemas struct {
}

func (*validateOpListSchemas) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListSchemas) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListSchemasInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListSchemasInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListTables struct {
}

func (*validateOpListTables) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListTables) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListTablesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListTablesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpBatchExecuteStatementValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpBatchExecuteStatement{}, middleware.After)
}

func addOpCancelStatementValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCancelStatement{}, middleware.After)
}

func addOpDescribeStatementValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeStatement{}, middleware.After)
}

func addOpDescribeTableValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeTable{}, middleware.After)
}

func addOpExecuteStatementValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpExecuteStatement{}, middleware.After)
}

func addOpGetStatementResultValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetStatementResult{}, middleware.After)
}

func addOpListDatabasesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListDatabases{}, middleware.After)
}

func addOpListSchemasValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListSchemas{}, middleware.After)
}

func addOpListTablesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTables{}, middleware.After)
}

func validateSqlParameter(v *types.SqlParameter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SqlParameter"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.Value == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Value"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateSqlParametersList(v []types.SqlParameter) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SqlParametersList"}
	for i := range v {
		if err := validateSqlParameter(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpBatchExecuteStatementInput(v *BatchExecuteStatementInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "BatchExecuteStatementInput"}
	if v.Sqls == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Sqls"))
	}
	if v.Database == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Database"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCancelStatementInput(v *CancelStatementInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CancelStatementInput"}
	if v.Id == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeStatementInput(v *DescribeStatementInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeStatementInput"}
	if v.Id == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeTableInput(v *DescribeTableInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeTableInput"}
	if v.Database == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Database"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpExecuteStatementInput(v *ExecuteStatementInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ExecuteStatementInput"}
	if v.Sql == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Sql"))
	}
	if v.Database == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Database"))
	}
	if v.Parameters != nil {
		if err := validateSqlParametersList(v.Parameters); err != nil {
			invalidParams.AddNested("Parameters", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetStatementResultInput(v *GetStatementResultInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetStatementResultInput"}
	if v.Id == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Id"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListDatabasesInput(v *ListDatabasesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListDatabasesInput"}
	if v.Database == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Database"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListSchemasInput(v *ListSchemasInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListSchemasInput"}
	if v.Database == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Database"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListTablesInput(v *ListTablesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListTablesInput"}
	if v.Database == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Database"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
