// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package logs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The AWS::Logs::LogAnomalyDetector resource specifies a CloudWatch Logs LogAnomalyDetector.
type LogAnomalyDetector struct {
	pulumi.CustomResourceState

	// Account ID for owner of detector
	AccountId pulumi.StringPtrOutput `pulumi:"accountId"`
	// ARN of LogAnomalyDetector
	AnomalyDetectorArn pulumi.StringOutput `pulumi:"anomalyDetectorArn"`
	// Current status of detector.
	AnomalyDetectorStatus pulumi.StringOutput     `pulumi:"anomalyDetectorStatus"`
	AnomalyVisibilityTime pulumi.Float64PtrOutput `pulumi:"anomalyVisibilityTime"`
	// When detector was created.
	CreationTimeStamp pulumi.Float64Output `pulumi:"creationTimeStamp"`
	// Name of detector
	DetectorName pulumi.StringPtrOutput `pulumi:"detectorName"`
	// How often log group is evaluated
	EvaluationFrequency LogAnomalyDetectorEvaluationFrequencyPtrOutput `pulumi:"evaluationFrequency"`
	FilterPattern       pulumi.StringPtrOutput                         `pulumi:"filterPattern"`
	// The Amazon Resource Name (ARN) of the CMK to use when encrypting log data.
	KmsKeyId pulumi.StringPtrOutput `pulumi:"kmsKeyId"`
	// When detector was lsat modified.
	LastModifiedTimeStamp pulumi.Float64Output `pulumi:"lastModifiedTimeStamp"`
	// List of Arns for the given log group
	LogGroupArnList pulumi.StringArrayOutput `pulumi:"logGroupArnList"`
}

// NewLogAnomalyDetector registers a new resource with the given unique name, arguments, and options.
func NewLogAnomalyDetector(ctx *pulumi.Context,
	name string, args *LogAnomalyDetectorArgs, opts ...pulumi.ResourceOption) (*LogAnomalyDetector, error) {
	if args == nil {
		args = &LogAnomalyDetectorArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LogAnomalyDetector
	err := ctx.RegisterResource("aws-native:logs:LogAnomalyDetector", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogAnomalyDetector gets an existing LogAnomalyDetector resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogAnomalyDetector(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogAnomalyDetectorState, opts ...pulumi.ResourceOption) (*LogAnomalyDetector, error) {
	var resource LogAnomalyDetector
	err := ctx.ReadResource("aws-native:logs:LogAnomalyDetector", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogAnomalyDetector resources.
type logAnomalyDetectorState struct {
}

type LogAnomalyDetectorState struct {
}

func (LogAnomalyDetectorState) ElementType() reflect.Type {
	return reflect.TypeOf((*logAnomalyDetectorState)(nil)).Elem()
}

type logAnomalyDetectorArgs struct {
	// Account ID for owner of detector
	AccountId             *string  `pulumi:"accountId"`
	AnomalyVisibilityTime *float64 `pulumi:"anomalyVisibilityTime"`
	// Name of detector
	DetectorName *string `pulumi:"detectorName"`
	// How often log group is evaluated
	EvaluationFrequency *LogAnomalyDetectorEvaluationFrequency `pulumi:"evaluationFrequency"`
	FilterPattern       *string                                `pulumi:"filterPattern"`
	// The Amazon Resource Name (ARN) of the CMK to use when encrypting log data.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// List of Arns for the given log group
	LogGroupArnList []string `pulumi:"logGroupArnList"`
}

// The set of arguments for constructing a LogAnomalyDetector resource.
type LogAnomalyDetectorArgs struct {
	// Account ID for owner of detector
	AccountId             pulumi.StringPtrInput
	AnomalyVisibilityTime pulumi.Float64PtrInput
	// Name of detector
	DetectorName pulumi.StringPtrInput
	// How often log group is evaluated
	EvaluationFrequency LogAnomalyDetectorEvaluationFrequencyPtrInput
	FilterPattern       pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of the CMK to use when encrypting log data.
	KmsKeyId pulumi.StringPtrInput
	// List of Arns for the given log group
	LogGroupArnList pulumi.StringArrayInput
}

func (LogAnomalyDetectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logAnomalyDetectorArgs)(nil)).Elem()
}

type LogAnomalyDetectorInput interface {
	pulumi.Input

	ToLogAnomalyDetectorOutput() LogAnomalyDetectorOutput
	ToLogAnomalyDetectorOutputWithContext(ctx context.Context) LogAnomalyDetectorOutput
}

func (*LogAnomalyDetector) ElementType() reflect.Type {
	return reflect.TypeOf((**LogAnomalyDetector)(nil)).Elem()
}

func (i *LogAnomalyDetector) ToLogAnomalyDetectorOutput() LogAnomalyDetectorOutput {
	return i.ToLogAnomalyDetectorOutputWithContext(context.Background())
}

func (i *LogAnomalyDetector) ToLogAnomalyDetectorOutputWithContext(ctx context.Context) LogAnomalyDetectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogAnomalyDetectorOutput)
}

func (i *LogAnomalyDetector) ToOutput(ctx context.Context) pulumix.Output[*LogAnomalyDetector] {
	return pulumix.Output[*LogAnomalyDetector]{
		OutputState: i.ToLogAnomalyDetectorOutputWithContext(ctx).OutputState,
	}
}

type LogAnomalyDetectorOutput struct{ *pulumi.OutputState }

func (LogAnomalyDetectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogAnomalyDetector)(nil)).Elem()
}

func (o LogAnomalyDetectorOutput) ToLogAnomalyDetectorOutput() LogAnomalyDetectorOutput {
	return o
}

func (o LogAnomalyDetectorOutput) ToLogAnomalyDetectorOutputWithContext(ctx context.Context) LogAnomalyDetectorOutput {
	return o
}

func (o LogAnomalyDetectorOutput) ToOutput(ctx context.Context) pulumix.Output[*LogAnomalyDetector] {
	return pulumix.Output[*LogAnomalyDetector]{
		OutputState: o.OutputState,
	}
}

// Account ID for owner of detector
func (o LogAnomalyDetectorOutput) AccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogAnomalyDetector) pulumi.StringPtrOutput { return v.AccountId }).(pulumi.StringPtrOutput)
}

// ARN of LogAnomalyDetector
func (o LogAnomalyDetectorOutput) AnomalyDetectorArn() pulumi.StringOutput {
	return o.ApplyT(func(v *LogAnomalyDetector) pulumi.StringOutput { return v.AnomalyDetectorArn }).(pulumi.StringOutput)
}

// Current status of detector.
func (o LogAnomalyDetectorOutput) AnomalyDetectorStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *LogAnomalyDetector) pulumi.StringOutput { return v.AnomalyDetectorStatus }).(pulumi.StringOutput)
}

func (o LogAnomalyDetectorOutput) AnomalyVisibilityTime() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *LogAnomalyDetector) pulumi.Float64PtrOutput { return v.AnomalyVisibilityTime }).(pulumi.Float64PtrOutput)
}

// When detector was created.
func (o LogAnomalyDetectorOutput) CreationTimeStamp() pulumi.Float64Output {
	return o.ApplyT(func(v *LogAnomalyDetector) pulumi.Float64Output { return v.CreationTimeStamp }).(pulumi.Float64Output)
}

// Name of detector
func (o LogAnomalyDetectorOutput) DetectorName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogAnomalyDetector) pulumi.StringPtrOutput { return v.DetectorName }).(pulumi.StringPtrOutput)
}

// How often log group is evaluated
func (o LogAnomalyDetectorOutput) EvaluationFrequency() LogAnomalyDetectorEvaluationFrequencyPtrOutput {
	return o.ApplyT(func(v *LogAnomalyDetector) LogAnomalyDetectorEvaluationFrequencyPtrOutput {
		return v.EvaluationFrequency
	}).(LogAnomalyDetectorEvaluationFrequencyPtrOutput)
}

func (o LogAnomalyDetectorOutput) FilterPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogAnomalyDetector) pulumi.StringPtrOutput { return v.FilterPattern }).(pulumi.StringPtrOutput)
}

// The Amazon Resource Name (ARN) of the CMK to use when encrypting log data.
func (o LogAnomalyDetectorOutput) KmsKeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogAnomalyDetector) pulumi.StringPtrOutput { return v.KmsKeyId }).(pulumi.StringPtrOutput)
}

// When detector was lsat modified.
func (o LogAnomalyDetectorOutput) LastModifiedTimeStamp() pulumi.Float64Output {
	return o.ApplyT(func(v *LogAnomalyDetector) pulumi.Float64Output { return v.LastModifiedTimeStamp }).(pulumi.Float64Output)
}

// List of Arns for the given log group
func (o LogAnomalyDetectorOutput) LogGroupArnList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LogAnomalyDetector) pulumi.StringArrayOutput { return v.LogGroupArnList }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LogAnomalyDetectorInput)(nil)).Elem(), &LogAnomalyDetector{})
	pulumi.RegisterOutputType(LogAnomalyDetectorOutput{})
}
