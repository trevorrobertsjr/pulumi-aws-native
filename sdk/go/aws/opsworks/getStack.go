// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package opsworks

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource Type definition for AWS::OpsWorks::Stack
func LookupStack(ctx *pulumi.Context, args *LookupStackArgs, opts ...pulumi.InvokeOption) (*LookupStackResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupStackResult
	err := ctx.Invoke("aws-native:opsworks:getStack", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStackArgs struct {
	Id string `pulumi:"id"`
}

type LookupStackResult struct {
	AgentVersion              *string                    `pulumi:"agentVersion"`
	Attributes                interface{}                `pulumi:"attributes"`
	ChefConfiguration         *StackChefConfiguration    `pulumi:"chefConfiguration"`
	ConfigurationManager      *StackConfigurationManager `pulumi:"configurationManager"`
	CustomCookbooksSource     *StackSource               `pulumi:"customCookbooksSource"`
	CustomJson                interface{}                `pulumi:"customJson"`
	DefaultAvailabilityZone   *string                    `pulumi:"defaultAvailabilityZone"`
	DefaultInstanceProfileArn *string                    `pulumi:"defaultInstanceProfileArn"`
	DefaultOs                 *string                    `pulumi:"defaultOs"`
	DefaultRootDeviceType     *string                    `pulumi:"defaultRootDeviceType"`
	DefaultSshKeyName         *string                    `pulumi:"defaultSshKeyName"`
	DefaultSubnetId           *string                    `pulumi:"defaultSubnetId"`
	EcsClusterArn             *string                    `pulumi:"ecsClusterArn"`
	ElasticIps                []StackElasticIp           `pulumi:"elasticIps"`
	HostnameTheme             *string                    `pulumi:"hostnameTheme"`
	Id                        *string                    `pulumi:"id"`
	Name                      *string                    `pulumi:"name"`
	RdsDbInstances            []StackRdsDbInstance       `pulumi:"rdsDbInstances"`
	Tags                      []StackTag                 `pulumi:"tags"`
	UseCustomCookbooks        *bool                      `pulumi:"useCustomCookbooks"`
	UseOpsworksSecurityGroups *bool                      `pulumi:"useOpsworksSecurityGroups"`
}

func LookupStackOutput(ctx *pulumi.Context, args LookupStackOutputArgs, opts ...pulumi.InvokeOption) LookupStackResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStackResult, error) {
			args := v.(LookupStackArgs)
			r, err := LookupStack(ctx, &args, opts...)
			var s LookupStackResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStackResultOutput)
}

type LookupStackOutputArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupStackOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStackArgs)(nil)).Elem()
}

type LookupStackResultOutput struct{ *pulumi.OutputState }

func (LookupStackResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStackResult)(nil)).Elem()
}

func (o LookupStackResultOutput) ToLookupStackResultOutput() LookupStackResultOutput {
	return o
}

func (o LookupStackResultOutput) ToLookupStackResultOutputWithContext(ctx context.Context) LookupStackResultOutput {
	return o
}

func (o LookupStackResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupStackResult] {
	return pulumix.Output[LookupStackResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupStackResultOutput) AgentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStackResult) *string { return v.AgentVersion }).(pulumi.StringPtrOutput)
}

func (o LookupStackResultOutput) Attributes() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupStackResult) interface{} { return v.Attributes }).(pulumi.AnyOutput)
}

func (o LookupStackResultOutput) ChefConfiguration() StackChefConfigurationPtrOutput {
	return o.ApplyT(func(v LookupStackResult) *StackChefConfiguration { return v.ChefConfiguration }).(StackChefConfigurationPtrOutput)
}

func (o LookupStackResultOutput) ConfigurationManager() StackConfigurationManagerPtrOutput {
	return o.ApplyT(func(v LookupStackResult) *StackConfigurationManager { return v.ConfigurationManager }).(StackConfigurationManagerPtrOutput)
}

func (o LookupStackResultOutput) CustomCookbooksSource() StackSourcePtrOutput {
	return o.ApplyT(func(v LookupStackResult) *StackSource { return v.CustomCookbooksSource }).(StackSourcePtrOutput)
}

func (o LookupStackResultOutput) CustomJson() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupStackResult) interface{} { return v.CustomJson }).(pulumi.AnyOutput)
}

func (o LookupStackResultOutput) DefaultAvailabilityZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStackResult) *string { return v.DefaultAvailabilityZone }).(pulumi.StringPtrOutput)
}

func (o LookupStackResultOutput) DefaultInstanceProfileArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStackResult) *string { return v.DefaultInstanceProfileArn }).(pulumi.StringPtrOutput)
}

func (o LookupStackResultOutput) DefaultOs() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStackResult) *string { return v.DefaultOs }).(pulumi.StringPtrOutput)
}

func (o LookupStackResultOutput) DefaultRootDeviceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStackResult) *string { return v.DefaultRootDeviceType }).(pulumi.StringPtrOutput)
}

func (o LookupStackResultOutput) DefaultSshKeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStackResult) *string { return v.DefaultSshKeyName }).(pulumi.StringPtrOutput)
}

func (o LookupStackResultOutput) DefaultSubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStackResult) *string { return v.DefaultSubnetId }).(pulumi.StringPtrOutput)
}

func (o LookupStackResultOutput) EcsClusterArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStackResult) *string { return v.EcsClusterArn }).(pulumi.StringPtrOutput)
}

func (o LookupStackResultOutput) ElasticIps() StackElasticIpArrayOutput {
	return o.ApplyT(func(v LookupStackResult) []StackElasticIp { return v.ElasticIps }).(StackElasticIpArrayOutput)
}

func (o LookupStackResultOutput) HostnameTheme() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStackResult) *string { return v.HostnameTheme }).(pulumi.StringPtrOutput)
}

func (o LookupStackResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStackResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupStackResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStackResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupStackResultOutput) RdsDbInstances() StackRdsDbInstanceArrayOutput {
	return o.ApplyT(func(v LookupStackResult) []StackRdsDbInstance { return v.RdsDbInstances }).(StackRdsDbInstanceArrayOutput)
}

func (o LookupStackResultOutput) Tags() StackTagArrayOutput {
	return o.ApplyT(func(v LookupStackResult) []StackTag { return v.Tags }).(StackTagArrayOutput)
}

func (o LookupStackResultOutput) UseCustomCookbooks() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupStackResult) *bool { return v.UseCustomCookbooks }).(pulumi.BoolPtrOutput)
}

func (o LookupStackResultOutput) UseOpsworksSecurityGroups() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupStackResult) *bool { return v.UseOpsworksSecurityGroups }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStackResultOutput{})
}
