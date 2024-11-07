// Code generated by smithy-go-codegen DO NOT EDIT.

package amp

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/amp/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpCreateAlertManagerDefinition struct {
}

func (*validateOpCreateAlertManagerDefinition) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateAlertManagerDefinition) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateAlertManagerDefinitionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateAlertManagerDefinitionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateLoggingConfiguration struct {
}

func (*validateOpCreateLoggingConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateLoggingConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateLoggingConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateLoggingConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateRuleGroupsNamespace struct {
}

func (*validateOpCreateRuleGroupsNamespace) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateRuleGroupsNamespace) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateRuleGroupsNamespaceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateRuleGroupsNamespaceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateScraper struct {
}

func (*validateOpCreateScraper) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateScraper) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateScraperInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateScraperInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteAlertManagerDefinition struct {
}

func (*validateOpDeleteAlertManagerDefinition) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteAlertManagerDefinition) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteAlertManagerDefinitionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteAlertManagerDefinitionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteLoggingConfiguration struct {
}

func (*validateOpDeleteLoggingConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteLoggingConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteLoggingConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteLoggingConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteRuleGroupsNamespace struct {
}

func (*validateOpDeleteRuleGroupsNamespace) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteRuleGroupsNamespace) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteRuleGroupsNamespaceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteRuleGroupsNamespaceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteScraper struct {
}

func (*validateOpDeleteScraper) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteScraper) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteScraperInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteScraperInput(input); err != nil {
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

type validateOpDescribeAlertManagerDefinition struct {
}

func (*validateOpDescribeAlertManagerDefinition) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeAlertManagerDefinition) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeAlertManagerDefinitionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeAlertManagerDefinitionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeLoggingConfiguration struct {
}

func (*validateOpDescribeLoggingConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeLoggingConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeLoggingConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeLoggingConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeRuleGroupsNamespace struct {
}

func (*validateOpDescribeRuleGroupsNamespace) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeRuleGroupsNamespace) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeRuleGroupsNamespaceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeRuleGroupsNamespaceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeScraper struct {
}

func (*validateOpDescribeScraper) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeScraper) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeScraperInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeScraperInput(input); err != nil {
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

type validateOpListRuleGroupsNamespaces struct {
}

func (*validateOpListRuleGroupsNamespaces) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListRuleGroupsNamespaces) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListRuleGroupsNamespacesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListRuleGroupsNamespacesInput(input); err != nil {
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

type validateOpPutAlertManagerDefinition struct {
}

func (*validateOpPutAlertManagerDefinition) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPutAlertManagerDefinition) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PutAlertManagerDefinitionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPutAlertManagerDefinitionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPutRuleGroupsNamespace struct {
}

func (*validateOpPutRuleGroupsNamespace) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPutRuleGroupsNamespace) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PutRuleGroupsNamespaceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPutRuleGroupsNamespaceInput(input); err != nil {
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

type validateOpUpdateLoggingConfiguration struct {
}

func (*validateOpUpdateLoggingConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateLoggingConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateLoggingConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateLoggingConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateScraper struct {
}

func (*validateOpUpdateScraper) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateScraper) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateScraperInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateScraperInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateWorkspaceAlias struct {
}

func (*validateOpUpdateWorkspaceAlias) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateWorkspaceAlias) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateWorkspaceAliasInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateWorkspaceAliasInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpCreateAlertManagerDefinitionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateAlertManagerDefinition{}, middleware.After)
}

func addOpCreateLoggingConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateLoggingConfiguration{}, middleware.After)
}

func addOpCreateRuleGroupsNamespaceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateRuleGroupsNamespace{}, middleware.After)
}

func addOpCreateScraperValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateScraper{}, middleware.After)
}

func addOpDeleteAlertManagerDefinitionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteAlertManagerDefinition{}, middleware.After)
}

func addOpDeleteLoggingConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteLoggingConfiguration{}, middleware.After)
}

func addOpDeleteRuleGroupsNamespaceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteRuleGroupsNamespace{}, middleware.After)
}

func addOpDeleteScraperValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteScraper{}, middleware.After)
}

func addOpDeleteWorkspaceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteWorkspace{}, middleware.After)
}

func addOpDescribeAlertManagerDefinitionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeAlertManagerDefinition{}, middleware.After)
}

func addOpDescribeLoggingConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeLoggingConfiguration{}, middleware.After)
}

func addOpDescribeRuleGroupsNamespaceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeRuleGroupsNamespace{}, middleware.After)
}

func addOpDescribeScraperValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeScraper{}, middleware.After)
}

func addOpDescribeWorkspaceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeWorkspace{}, middleware.After)
}

func addOpListRuleGroupsNamespacesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListRuleGroupsNamespaces{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpPutAlertManagerDefinitionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPutAlertManagerDefinition{}, middleware.After)
}

func addOpPutRuleGroupsNamespaceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPutRuleGroupsNamespace{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func addOpUpdateLoggingConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateLoggingConfiguration{}, middleware.After)
}

func addOpUpdateScraperValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateScraper{}, middleware.After)
}

func addOpUpdateWorkspaceAliasValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateWorkspaceAlias{}, middleware.After)
}

func validateAmpConfiguration(v *types.AmpConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AmpConfiguration"}
	if v.WorkspaceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("WorkspaceArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateDestination(v types.Destination) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Destination"}
	switch uv := v.(type) {
	case *types.DestinationMemberAmpConfiguration:
		if err := validateAmpConfiguration(&uv.Value); err != nil {
			invalidParams.AddNested("[ampConfiguration]", err.(smithy.InvalidParamsError))
		}

	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateEksConfiguration(v *types.EksConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "EksConfiguration"}
	if v.ClusterArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClusterArn"))
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

func validateSource(v types.Source) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Source"}
	switch uv := v.(type) {
	case *types.SourceMemberEksConfiguration:
		if err := validateEksConfiguration(&uv.Value); err != nil {
			invalidParams.AddNested("[eksConfiguration]", err.(smithy.InvalidParamsError))
		}

	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateAlertManagerDefinitionInput(v *CreateAlertManagerDefinitionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateAlertManagerDefinitionInput"}
	if v.WorkspaceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("WorkspaceId"))
	}
	if v.Data == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Data"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateLoggingConfigurationInput(v *CreateLoggingConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateLoggingConfigurationInput"}
	if v.WorkspaceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("WorkspaceId"))
	}
	if v.LogGroupArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("LogGroupArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateRuleGroupsNamespaceInput(v *CreateRuleGroupsNamespaceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateRuleGroupsNamespaceInput"}
	if v.WorkspaceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("WorkspaceId"))
	}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.Data == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Data"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateScraperInput(v *CreateScraperInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateScraperInput"}
	if v.ScrapeConfiguration == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ScrapeConfiguration"))
	}
	if v.Source == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Source"))
	} else if v.Source != nil {
		if err := validateSource(v.Source); err != nil {
			invalidParams.AddNested("Source", err.(smithy.InvalidParamsError))
		}
	}
	if v.Destination == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Destination"))
	} else if v.Destination != nil {
		if err := validateDestination(v.Destination); err != nil {
			invalidParams.AddNested("Destination", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteAlertManagerDefinitionInput(v *DeleteAlertManagerDefinitionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteAlertManagerDefinitionInput"}
	if v.WorkspaceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("WorkspaceId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteLoggingConfigurationInput(v *DeleteLoggingConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteLoggingConfigurationInput"}
	if v.WorkspaceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("WorkspaceId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteRuleGroupsNamespaceInput(v *DeleteRuleGroupsNamespaceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteRuleGroupsNamespaceInput"}
	if v.WorkspaceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("WorkspaceId"))
	}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteScraperInput(v *DeleteScraperInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteScraperInput"}
	if v.ScraperId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ScraperId"))
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

func validateOpDescribeAlertManagerDefinitionInput(v *DescribeAlertManagerDefinitionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeAlertManagerDefinitionInput"}
	if v.WorkspaceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("WorkspaceId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeLoggingConfigurationInput(v *DescribeLoggingConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeLoggingConfigurationInput"}
	if v.WorkspaceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("WorkspaceId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeRuleGroupsNamespaceInput(v *DescribeRuleGroupsNamespaceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeRuleGroupsNamespaceInput"}
	if v.WorkspaceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("WorkspaceId"))
	}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeScraperInput(v *DescribeScraperInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeScraperInput"}
	if v.ScraperId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ScraperId"))
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

func validateOpListRuleGroupsNamespacesInput(v *ListRuleGroupsNamespacesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListRuleGroupsNamespacesInput"}
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

func validateOpPutAlertManagerDefinitionInput(v *PutAlertManagerDefinitionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PutAlertManagerDefinitionInput"}
	if v.WorkspaceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("WorkspaceId"))
	}
	if v.Data == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Data"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPutRuleGroupsNamespaceInput(v *PutRuleGroupsNamespaceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PutRuleGroupsNamespaceInput"}
	if v.WorkspaceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("WorkspaceId"))
	}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.Data == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Data"))
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

func validateOpUpdateLoggingConfigurationInput(v *UpdateLoggingConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateLoggingConfigurationInput"}
	if v.WorkspaceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("WorkspaceId"))
	}
	if v.LogGroupArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("LogGroupArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateScraperInput(v *UpdateScraperInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateScraperInput"}
	if v.ScraperId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ScraperId"))
	}
	if v.Destination != nil {
		if err := validateDestination(v.Destination); err != nil {
			invalidParams.AddNested("Destination", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateWorkspaceAliasInput(v *UpdateWorkspaceAliasInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateWorkspaceAliasInput"}
	if v.WorkspaceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("WorkspaceId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
