// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package glue

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource Type definition for AWS::Glue::Job
func LookupJob(ctx *pulumi.Context, args *LookupJobArgs, opts ...pulumi.InvokeOption) (*LookupJobResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupJobResult
	err := ctx.Invoke("aws-native:glue:getJob", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupJobArgs struct {
	Id string `pulumi:"id"`
}

type LookupJobResult struct {
	AllocatedCapacity       *float64                 `pulumi:"allocatedCapacity"`
	Command                 *JobCommand              `pulumi:"command"`
	Connections             *JobConnectionsList      `pulumi:"connections"`
	DefaultArguments        interface{}              `pulumi:"defaultArguments"`
	Description             *string                  `pulumi:"description"`
	ExecutionClass          *string                  `pulumi:"executionClass"`
	ExecutionProperty       *JobExecutionProperty    `pulumi:"executionProperty"`
	GlueVersion             *string                  `pulumi:"glueVersion"`
	Id                      *string                  `pulumi:"id"`
	LogUri                  *string                  `pulumi:"logUri"`
	MaxCapacity             *float64                 `pulumi:"maxCapacity"`
	MaxRetries              *float64                 `pulumi:"maxRetries"`
	NonOverridableArguments interface{}              `pulumi:"nonOverridableArguments"`
	NotificationProperty    *JobNotificationProperty `pulumi:"notificationProperty"`
	NumberOfWorkers         *int                     `pulumi:"numberOfWorkers"`
	Role                    *string                  `pulumi:"role"`
	SecurityConfiguration   *string                  `pulumi:"securityConfiguration"`
	Tags                    interface{}              `pulumi:"tags"`
	Timeout                 *int                     `pulumi:"timeout"`
	WorkerType              *string                  `pulumi:"workerType"`
}

func LookupJobOutput(ctx *pulumi.Context, args LookupJobOutputArgs, opts ...pulumi.InvokeOption) LookupJobResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupJobResult, error) {
			args := v.(LookupJobArgs)
			r, err := LookupJob(ctx, &args, opts...)
			var s LookupJobResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupJobResultOutput)
}

type LookupJobOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupJobOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupJobArgs)(nil)).Elem()
}

type LookupJobResultOutput struct{ *pulumi.OutputState }

func (LookupJobResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupJobResult)(nil)).Elem()
}

func (o LookupJobResultOutput) ToLookupJobResultOutput() LookupJobResultOutput {
	return o
}

func (o LookupJobResultOutput) ToLookupJobResultOutputWithContext(ctx context.Context) LookupJobResultOutput {
	return o
}

func (o LookupJobResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupJobResult] {
	return pulumix.Output[LookupJobResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupJobResultOutput) AllocatedCapacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupJobResult) *float64 { return v.AllocatedCapacity }).(pulumi.Float64PtrOutput)
}

func (o LookupJobResultOutput) Command() JobCommandPtrOutput {
	return o.ApplyT(func(v LookupJobResult) *JobCommand { return v.Command }).(JobCommandPtrOutput)
}

func (o LookupJobResultOutput) Connections() JobConnectionsListPtrOutput {
	return o.ApplyT(func(v LookupJobResult) *JobConnectionsList { return v.Connections }).(JobConnectionsListPtrOutput)
}

func (o LookupJobResultOutput) DefaultArguments() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupJobResult) interface{} { return v.DefaultArguments }).(pulumi.AnyOutput)
}

func (o LookupJobResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupJobResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupJobResultOutput) ExecutionClass() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupJobResult) *string { return v.ExecutionClass }).(pulumi.StringPtrOutput)
}

func (o LookupJobResultOutput) ExecutionProperty() JobExecutionPropertyPtrOutput {
	return o.ApplyT(func(v LookupJobResult) *JobExecutionProperty { return v.ExecutionProperty }).(JobExecutionPropertyPtrOutput)
}

func (o LookupJobResultOutput) GlueVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupJobResult) *string { return v.GlueVersion }).(pulumi.StringPtrOutput)
}

func (o LookupJobResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupJobResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupJobResultOutput) LogUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupJobResult) *string { return v.LogUri }).(pulumi.StringPtrOutput)
}

func (o LookupJobResultOutput) MaxCapacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupJobResult) *float64 { return v.MaxCapacity }).(pulumi.Float64PtrOutput)
}

func (o LookupJobResultOutput) MaxRetries() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupJobResult) *float64 { return v.MaxRetries }).(pulumi.Float64PtrOutput)
}

func (o LookupJobResultOutput) NonOverridableArguments() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupJobResult) interface{} { return v.NonOverridableArguments }).(pulumi.AnyOutput)
}

func (o LookupJobResultOutput) NotificationProperty() JobNotificationPropertyPtrOutput {
	return o.ApplyT(func(v LookupJobResult) *JobNotificationProperty { return v.NotificationProperty }).(JobNotificationPropertyPtrOutput)
}

func (o LookupJobResultOutput) NumberOfWorkers() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupJobResult) *int { return v.NumberOfWorkers }).(pulumi.IntPtrOutput)
}

func (o LookupJobResultOutput) Role() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupJobResult) *string { return v.Role }).(pulumi.StringPtrOutput)
}

func (o LookupJobResultOutput) SecurityConfiguration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupJobResult) *string { return v.SecurityConfiguration }).(pulumi.StringPtrOutput)
}

func (o LookupJobResultOutput) Tags() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupJobResult) interface{} { return v.Tags }).(pulumi.AnyOutput)
}

func (o LookupJobResultOutput) Timeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupJobResult) *int { return v.Timeout }).(pulumi.IntPtrOutput)
}

func (o LookupJobResultOutput) WorkerType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupJobResult) *string { return v.WorkerType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupJobResultOutput{})
}
