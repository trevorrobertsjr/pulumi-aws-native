// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mwaa

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource schema for AWS::MWAA::Environment
func LookupEnvironment(ctx *pulumi.Context, args *LookupEnvironmentArgs, opts ...pulumi.InvokeOption) (*LookupEnvironmentResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupEnvironmentResult
	err := ctx.Invoke("aws-native:mwaa:getEnvironment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEnvironmentArgs struct {
	Name string `pulumi:"name"`
}

type LookupEnvironmentResult struct {
	// Key/value pairs representing Airflow configuration variables.
	//     Keys are prefixed by their section:
	//
	//     [core]
	//     dags_folder={AIRFLOW_HOME}/dags
	//
	//     Would be represented as
	//
	//     "core.dags_folder": "{AIRFLOW_HOME}/dags"
	AirflowConfigurationOptions  interface{}                      `pulumi:"airflowConfigurationOptions"`
	AirflowVersion               *string                          `pulumi:"airflowVersion"`
	Arn                          *string                          `pulumi:"arn"`
	CeleryExecutorQueue          *string                          `pulumi:"celeryExecutorQueue"`
	DagS3Path                    *string                          `pulumi:"dagS3Path"`
	DatabaseVpcEndpointService   *string                          `pulumi:"databaseVpcEndpointService"`
	EnvironmentClass             *string                          `pulumi:"environmentClass"`
	ExecutionRoleArn             *string                          `pulumi:"executionRoleArn"`
	LoggingConfiguration         *EnvironmentLoggingConfiguration `pulumi:"loggingConfiguration"`
	MaxWorkers                   *int                             `pulumi:"maxWorkers"`
	MinWorkers                   *int                             `pulumi:"minWorkers"`
	NetworkConfiguration         *EnvironmentNetworkConfiguration `pulumi:"networkConfiguration"`
	PluginsS3ObjectVersion       *string                          `pulumi:"pluginsS3ObjectVersion"`
	PluginsS3Path                *string                          `pulumi:"pluginsS3Path"`
	RequirementsS3ObjectVersion  *string                          `pulumi:"requirementsS3ObjectVersion"`
	RequirementsS3Path           *string                          `pulumi:"requirementsS3Path"`
	Schedulers                   *int                             `pulumi:"schedulers"`
	SourceBucketArn              *string                          `pulumi:"sourceBucketArn"`
	StartupScriptS3ObjectVersion *string                          `pulumi:"startupScriptS3ObjectVersion"`
	StartupScriptS3Path          *string                          `pulumi:"startupScriptS3Path"`
	// A map of tags for the environment.
	Tags                         interface{}                     `pulumi:"tags"`
	WebserverAccessMode          *EnvironmentWebserverAccessMode `pulumi:"webserverAccessMode"`
	WebserverUrl                 *string                         `pulumi:"webserverUrl"`
	WebserverVpcEndpointService  *string                         `pulumi:"webserverVpcEndpointService"`
	WeeklyMaintenanceWindowStart *string                         `pulumi:"weeklyMaintenanceWindowStart"`
}

func LookupEnvironmentOutput(ctx *pulumi.Context, args LookupEnvironmentOutputArgs, opts ...pulumi.InvokeOption) LookupEnvironmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEnvironmentResult, error) {
			args := v.(LookupEnvironmentArgs)
			r, err := LookupEnvironment(ctx, &args, opts...)
			var s LookupEnvironmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEnvironmentResultOutput)
}

type LookupEnvironmentOutputArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupEnvironmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnvironmentArgs)(nil)).Elem()
}

type LookupEnvironmentResultOutput struct{ *pulumi.OutputState }

func (LookupEnvironmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnvironmentResult)(nil)).Elem()
}

func (o LookupEnvironmentResultOutput) ToLookupEnvironmentResultOutput() LookupEnvironmentResultOutput {
	return o
}

func (o LookupEnvironmentResultOutput) ToLookupEnvironmentResultOutputWithContext(ctx context.Context) LookupEnvironmentResultOutput {
	return o
}

func (o LookupEnvironmentResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupEnvironmentResult] {
	return pulumix.Output[LookupEnvironmentResult]{
		OutputState: o.OutputState,
	}
}

// Key/value pairs representing Airflow configuration variables.
//
//	Keys are prefixed by their section:
//
//	[core]
//	dags_folder={AIRFLOW_HOME}/dags
//
//	Would be represented as
//
//	"core.dags_folder": "{AIRFLOW_HOME}/dags"
func (o LookupEnvironmentResultOutput) AirflowConfigurationOptions() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) interface{} { return v.AirflowConfigurationOptions }).(pulumi.AnyOutput)
}

func (o LookupEnvironmentResultOutput) AirflowVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) *string { return v.AirflowVersion }).(pulumi.StringPtrOutput)
}

func (o LookupEnvironmentResultOutput) Arn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) *string { return v.Arn }).(pulumi.StringPtrOutput)
}

func (o LookupEnvironmentResultOutput) CeleryExecutorQueue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) *string { return v.CeleryExecutorQueue }).(pulumi.StringPtrOutput)
}

func (o LookupEnvironmentResultOutput) DagS3Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) *string { return v.DagS3Path }).(pulumi.StringPtrOutput)
}

func (o LookupEnvironmentResultOutput) DatabaseVpcEndpointService() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) *string { return v.DatabaseVpcEndpointService }).(pulumi.StringPtrOutput)
}

func (o LookupEnvironmentResultOutput) EnvironmentClass() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) *string { return v.EnvironmentClass }).(pulumi.StringPtrOutput)
}

func (o LookupEnvironmentResultOutput) ExecutionRoleArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) *string { return v.ExecutionRoleArn }).(pulumi.StringPtrOutput)
}

func (o LookupEnvironmentResultOutput) LoggingConfiguration() EnvironmentLoggingConfigurationPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) *EnvironmentLoggingConfiguration { return v.LoggingConfiguration }).(EnvironmentLoggingConfigurationPtrOutput)
}

func (o LookupEnvironmentResultOutput) MaxWorkers() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) *int { return v.MaxWorkers }).(pulumi.IntPtrOutput)
}

func (o LookupEnvironmentResultOutput) MinWorkers() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) *int { return v.MinWorkers }).(pulumi.IntPtrOutput)
}

func (o LookupEnvironmentResultOutput) NetworkConfiguration() EnvironmentNetworkConfigurationPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) *EnvironmentNetworkConfiguration { return v.NetworkConfiguration }).(EnvironmentNetworkConfigurationPtrOutput)
}

func (o LookupEnvironmentResultOutput) PluginsS3ObjectVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) *string { return v.PluginsS3ObjectVersion }).(pulumi.StringPtrOutput)
}

func (o LookupEnvironmentResultOutput) PluginsS3Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) *string { return v.PluginsS3Path }).(pulumi.StringPtrOutput)
}

func (o LookupEnvironmentResultOutput) RequirementsS3ObjectVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) *string { return v.RequirementsS3ObjectVersion }).(pulumi.StringPtrOutput)
}

func (o LookupEnvironmentResultOutput) RequirementsS3Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) *string { return v.RequirementsS3Path }).(pulumi.StringPtrOutput)
}

func (o LookupEnvironmentResultOutput) Schedulers() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) *int { return v.Schedulers }).(pulumi.IntPtrOutput)
}

func (o LookupEnvironmentResultOutput) SourceBucketArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) *string { return v.SourceBucketArn }).(pulumi.StringPtrOutput)
}

func (o LookupEnvironmentResultOutput) StartupScriptS3ObjectVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) *string { return v.StartupScriptS3ObjectVersion }).(pulumi.StringPtrOutput)
}

func (o LookupEnvironmentResultOutput) StartupScriptS3Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) *string { return v.StartupScriptS3Path }).(pulumi.StringPtrOutput)
}

// A map of tags for the environment.
func (o LookupEnvironmentResultOutput) Tags() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) interface{} { return v.Tags }).(pulumi.AnyOutput)
}

func (o LookupEnvironmentResultOutput) WebserverAccessMode() EnvironmentWebserverAccessModePtrOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) *EnvironmentWebserverAccessMode { return v.WebserverAccessMode }).(EnvironmentWebserverAccessModePtrOutput)
}

func (o LookupEnvironmentResultOutput) WebserverUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) *string { return v.WebserverUrl }).(pulumi.StringPtrOutput)
}

func (o LookupEnvironmentResultOutput) WebserverVpcEndpointService() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) *string { return v.WebserverVpcEndpointService }).(pulumi.StringPtrOutput)
}

func (o LookupEnvironmentResultOutput) WeeklyMaintenanceWindowStart() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) *string { return v.WeeklyMaintenanceWindowStart }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEnvironmentResultOutput{})
}
