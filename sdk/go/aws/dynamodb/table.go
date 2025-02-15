// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dynamodb

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Version: None. Resource Type definition for AWS::DynamoDB::Table
type Table struct {
	pulumi.CustomResourceState

	Arn                              pulumi.StringOutput                            `pulumi:"arn"`
	AttributeDefinitions             TableAttributeDefinitionArrayOutput            `pulumi:"attributeDefinitions"`
	BillingMode                      pulumi.StringPtrOutput                         `pulumi:"billingMode"`
	ContributorInsightsSpecification TableContributorInsightsSpecificationPtrOutput `pulumi:"contributorInsightsSpecification"`
	DeletionProtectionEnabled        pulumi.BoolPtrOutput                           `pulumi:"deletionProtectionEnabled"`
	GlobalSecondaryIndexes           TableGlobalSecondaryIndexArrayOutput           `pulumi:"globalSecondaryIndexes"`
	ImportSourceSpecification        TableImportSourceSpecificationPtrOutput        `pulumi:"importSourceSpecification"`
	KeySchema                        pulumi.AnyOutput                               `pulumi:"keySchema"`
	KinesisStreamSpecification       TableKinesisStreamSpecificationPtrOutput       `pulumi:"kinesisStreamSpecification"`
	LocalSecondaryIndexes            TableLocalSecondaryIndexArrayOutput            `pulumi:"localSecondaryIndexes"`
	PointInTimeRecoverySpecification TablePointInTimeRecoverySpecificationPtrOutput `pulumi:"pointInTimeRecoverySpecification"`
	ProvisionedThroughput            TableProvisionedThroughputPtrOutput            `pulumi:"provisionedThroughput"`
	SseSpecification                 TableSseSpecificationPtrOutput                 `pulumi:"sseSpecification"`
	StreamArn                        pulumi.StringOutput                            `pulumi:"streamArn"`
	StreamSpecification              TableStreamSpecificationPtrOutput              `pulumi:"streamSpecification"`
	TableClass                       pulumi.StringPtrOutput                         `pulumi:"tableClass"`
	TableName                        pulumi.StringPtrOutput                         `pulumi:"tableName"`
	Tags                             TableTagArrayOutput                            `pulumi:"tags"`
	TimeToLiveSpecification          TableTimeToLiveSpecificationPtrOutput          `pulumi:"timeToLiveSpecification"`
}

// NewTable registers a new resource with the given unique name, arguments, and options.
func NewTable(ctx *pulumi.Context,
	name string, args *TableArgs, opts ...pulumi.ResourceOption) (*Table, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.KeySchema == nil {
		return nil, errors.New("invalid value for required argument 'KeySchema'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"importSourceSpecification",
		"tableName",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Table
	err := ctx.RegisterResource("aws-native:dynamodb:Table", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTable gets an existing Table resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TableState, opts ...pulumi.ResourceOption) (*Table, error) {
	var resource Table
	err := ctx.ReadResource("aws-native:dynamodb:Table", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Table resources.
type tableState struct {
}

type TableState struct {
}

func (TableState) ElementType() reflect.Type {
	return reflect.TypeOf((*tableState)(nil)).Elem()
}

type tableArgs struct {
	AttributeDefinitions             []TableAttributeDefinition             `pulumi:"attributeDefinitions"`
	BillingMode                      *string                                `pulumi:"billingMode"`
	ContributorInsightsSpecification *TableContributorInsightsSpecification `pulumi:"contributorInsightsSpecification"`
	DeletionProtectionEnabled        *bool                                  `pulumi:"deletionProtectionEnabled"`
	GlobalSecondaryIndexes           []TableGlobalSecondaryIndex            `pulumi:"globalSecondaryIndexes"`
	ImportSourceSpecification        *TableImportSourceSpecification        `pulumi:"importSourceSpecification"`
	KeySchema                        interface{}                            `pulumi:"keySchema"`
	KinesisStreamSpecification       *TableKinesisStreamSpecification       `pulumi:"kinesisStreamSpecification"`
	LocalSecondaryIndexes            []TableLocalSecondaryIndex             `pulumi:"localSecondaryIndexes"`
	PointInTimeRecoverySpecification *TablePointInTimeRecoverySpecification `pulumi:"pointInTimeRecoverySpecification"`
	ProvisionedThroughput            *TableProvisionedThroughput            `pulumi:"provisionedThroughput"`
	SseSpecification                 *TableSseSpecification                 `pulumi:"sseSpecification"`
	StreamSpecification              *TableStreamSpecification              `pulumi:"streamSpecification"`
	TableClass                       *string                                `pulumi:"tableClass"`
	TableName                        *string                                `pulumi:"tableName"`
	Tags                             []TableTag                             `pulumi:"tags"`
	TimeToLiveSpecification          *TableTimeToLiveSpecification          `pulumi:"timeToLiveSpecification"`
}

// The set of arguments for constructing a Table resource.
type TableArgs struct {
	AttributeDefinitions             TableAttributeDefinitionArrayInput
	BillingMode                      pulumi.StringPtrInput
	ContributorInsightsSpecification TableContributorInsightsSpecificationPtrInput
	DeletionProtectionEnabled        pulumi.BoolPtrInput
	GlobalSecondaryIndexes           TableGlobalSecondaryIndexArrayInput
	ImportSourceSpecification        TableImportSourceSpecificationPtrInput
	KeySchema                        pulumi.Input
	KinesisStreamSpecification       TableKinesisStreamSpecificationPtrInput
	LocalSecondaryIndexes            TableLocalSecondaryIndexArrayInput
	PointInTimeRecoverySpecification TablePointInTimeRecoverySpecificationPtrInput
	ProvisionedThroughput            TableProvisionedThroughputPtrInput
	SseSpecification                 TableSseSpecificationPtrInput
	StreamSpecification              TableStreamSpecificationPtrInput
	TableClass                       pulumi.StringPtrInput
	TableName                        pulumi.StringPtrInput
	Tags                             TableTagArrayInput
	TimeToLiveSpecification          TableTimeToLiveSpecificationPtrInput
}

func (TableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tableArgs)(nil)).Elem()
}

type TableInput interface {
	pulumi.Input

	ToTableOutput() TableOutput
	ToTableOutputWithContext(ctx context.Context) TableOutput
}

func (*Table) ElementType() reflect.Type {
	return reflect.TypeOf((**Table)(nil)).Elem()
}

func (i *Table) ToTableOutput() TableOutput {
	return i.ToTableOutputWithContext(context.Background())
}

func (i *Table) ToTableOutputWithContext(ctx context.Context) TableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableOutput)
}

func (i *Table) ToOutput(ctx context.Context) pulumix.Output[*Table] {
	return pulumix.Output[*Table]{
		OutputState: i.ToTableOutputWithContext(ctx).OutputState,
	}
}

type TableOutput struct{ *pulumi.OutputState }

func (TableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Table)(nil)).Elem()
}

func (o TableOutput) ToTableOutput() TableOutput {
	return o
}

func (o TableOutput) ToTableOutputWithContext(ctx context.Context) TableOutput {
	return o
}

func (o TableOutput) ToOutput(ctx context.Context) pulumix.Output[*Table] {
	return pulumix.Output[*Table]{
		OutputState: o.OutputState,
	}
}

func (o TableOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *Table) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o TableOutput) AttributeDefinitions() TableAttributeDefinitionArrayOutput {
	return o.ApplyT(func(v *Table) TableAttributeDefinitionArrayOutput { return v.AttributeDefinitions }).(TableAttributeDefinitionArrayOutput)
}

func (o TableOutput) BillingMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Table) pulumi.StringPtrOutput { return v.BillingMode }).(pulumi.StringPtrOutput)
}

func (o TableOutput) ContributorInsightsSpecification() TableContributorInsightsSpecificationPtrOutput {
	return o.ApplyT(func(v *Table) TableContributorInsightsSpecificationPtrOutput {
		return v.ContributorInsightsSpecification
	}).(TableContributorInsightsSpecificationPtrOutput)
}

func (o TableOutput) DeletionProtectionEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Table) pulumi.BoolPtrOutput { return v.DeletionProtectionEnabled }).(pulumi.BoolPtrOutput)
}

func (o TableOutput) GlobalSecondaryIndexes() TableGlobalSecondaryIndexArrayOutput {
	return o.ApplyT(func(v *Table) TableGlobalSecondaryIndexArrayOutput { return v.GlobalSecondaryIndexes }).(TableGlobalSecondaryIndexArrayOutput)
}

func (o TableOutput) ImportSourceSpecification() TableImportSourceSpecificationPtrOutput {
	return o.ApplyT(func(v *Table) TableImportSourceSpecificationPtrOutput { return v.ImportSourceSpecification }).(TableImportSourceSpecificationPtrOutput)
}

func (o TableOutput) KeySchema() pulumi.AnyOutput {
	return o.ApplyT(func(v *Table) pulumi.AnyOutput { return v.KeySchema }).(pulumi.AnyOutput)
}

func (o TableOutput) KinesisStreamSpecification() TableKinesisStreamSpecificationPtrOutput {
	return o.ApplyT(func(v *Table) TableKinesisStreamSpecificationPtrOutput { return v.KinesisStreamSpecification }).(TableKinesisStreamSpecificationPtrOutput)
}

func (o TableOutput) LocalSecondaryIndexes() TableLocalSecondaryIndexArrayOutput {
	return o.ApplyT(func(v *Table) TableLocalSecondaryIndexArrayOutput { return v.LocalSecondaryIndexes }).(TableLocalSecondaryIndexArrayOutput)
}

func (o TableOutput) PointInTimeRecoverySpecification() TablePointInTimeRecoverySpecificationPtrOutput {
	return o.ApplyT(func(v *Table) TablePointInTimeRecoverySpecificationPtrOutput {
		return v.PointInTimeRecoverySpecification
	}).(TablePointInTimeRecoverySpecificationPtrOutput)
}

func (o TableOutput) ProvisionedThroughput() TableProvisionedThroughputPtrOutput {
	return o.ApplyT(func(v *Table) TableProvisionedThroughputPtrOutput { return v.ProvisionedThroughput }).(TableProvisionedThroughputPtrOutput)
}

func (o TableOutput) SseSpecification() TableSseSpecificationPtrOutput {
	return o.ApplyT(func(v *Table) TableSseSpecificationPtrOutput { return v.SseSpecification }).(TableSseSpecificationPtrOutput)
}

func (o TableOutput) StreamArn() pulumi.StringOutput {
	return o.ApplyT(func(v *Table) pulumi.StringOutput { return v.StreamArn }).(pulumi.StringOutput)
}

func (o TableOutput) StreamSpecification() TableStreamSpecificationPtrOutput {
	return o.ApplyT(func(v *Table) TableStreamSpecificationPtrOutput { return v.StreamSpecification }).(TableStreamSpecificationPtrOutput)
}

func (o TableOutput) TableClass() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Table) pulumi.StringPtrOutput { return v.TableClass }).(pulumi.StringPtrOutput)
}

func (o TableOutput) TableName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Table) pulumi.StringPtrOutput { return v.TableName }).(pulumi.StringPtrOutput)
}

func (o TableOutput) Tags() TableTagArrayOutput {
	return o.ApplyT(func(v *Table) TableTagArrayOutput { return v.Tags }).(TableTagArrayOutput)
}

func (o TableOutput) TimeToLiveSpecification() TableTimeToLiveSpecificationPtrOutput {
	return o.ApplyT(func(v *Table) TableTimeToLiveSpecificationPtrOutput { return v.TimeToLiveSpecification }).(TableTimeToLiveSpecificationPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TableInput)(nil)).Elem(), &Table{})
	pulumi.RegisterOutputType(TableOutput{})
}
