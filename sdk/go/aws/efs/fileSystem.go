// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package efs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource Type definition for AWS::EFS::FileSystem
type FileSystem struct {
	pulumi.CustomResourceState

	Arn                  pulumi.StringOutput             `pulumi:"arn"`
	AvailabilityZoneName pulumi.StringPtrOutput          `pulumi:"availabilityZoneName"`
	BackupPolicy         FileSystemBackupPolicyPtrOutput `pulumi:"backupPolicy"`
	// Whether to bypass the FileSystemPolicy lockout safety check. The policy lockout safety check determines whether the policy in the request will prevent the principal making the request to be locked out from making future PutFileSystemPolicy requests on the file system. Set BypassPolicyLockoutSafetyCheck to True only when you intend to prevent the principal that is making the request from making a subsequent PutFileSystemPolicy request on the file system. Defaults to false
	BypassPolicyLockoutSafetyCheck pulumi.BoolPtrOutput                        `pulumi:"bypassPolicyLockoutSafetyCheck"`
	Encrypted                      pulumi.BoolPtrOutput                        `pulumi:"encrypted"`
	FileSystemId                   pulumi.StringOutput                         `pulumi:"fileSystemId"`
	FileSystemPolicy               pulumi.AnyOutput                            `pulumi:"fileSystemPolicy"`
	FileSystemProtection           FileSystemProtectionPtrOutput               `pulumi:"fileSystemProtection"`
	FileSystemTags                 FileSystemElasticFileSystemTagArrayOutput   `pulumi:"fileSystemTags"`
	KmsKeyId                       pulumi.StringPtrOutput                      `pulumi:"kmsKeyId"`
	LifecyclePolicies              FileSystemLifecyclePolicyArrayOutput        `pulumi:"lifecyclePolicies"`
	PerformanceMode                pulumi.StringPtrOutput                      `pulumi:"performanceMode"`
	ProvisionedThroughputInMibps   pulumi.Float64PtrOutput                     `pulumi:"provisionedThroughputInMibps"`
	ReplicationConfiguration       FileSystemReplicationConfigurationPtrOutput `pulumi:"replicationConfiguration"`
	ThroughputMode                 pulumi.StringPtrOutput                      `pulumi:"throughputMode"`
}

// NewFileSystem registers a new resource with the given unique name, arguments, and options.
func NewFileSystem(ctx *pulumi.Context,
	name string, args *FileSystemArgs, opts ...pulumi.ResourceOption) (*FileSystem, error) {
	if args == nil {
		args = &FileSystemArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"availabilityZoneName",
		"encrypted",
		"kmsKeyId",
		"performanceMode",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FileSystem
	err := ctx.RegisterResource("aws-native:efs:FileSystem", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFileSystem gets an existing FileSystem resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFileSystem(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FileSystemState, opts ...pulumi.ResourceOption) (*FileSystem, error) {
	var resource FileSystem
	err := ctx.ReadResource("aws-native:efs:FileSystem", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FileSystem resources.
type fileSystemState struct {
}

type FileSystemState struct {
}

func (FileSystemState) ElementType() reflect.Type {
	return reflect.TypeOf((*fileSystemState)(nil)).Elem()
}

type fileSystemArgs struct {
	AvailabilityZoneName *string                 `pulumi:"availabilityZoneName"`
	BackupPolicy         *FileSystemBackupPolicy `pulumi:"backupPolicy"`
	// Whether to bypass the FileSystemPolicy lockout safety check. The policy lockout safety check determines whether the policy in the request will prevent the principal making the request to be locked out from making future PutFileSystemPolicy requests on the file system. Set BypassPolicyLockoutSafetyCheck to True only when you intend to prevent the principal that is making the request from making a subsequent PutFileSystemPolicy request on the file system. Defaults to false
	BypassPolicyLockoutSafetyCheck *bool                               `pulumi:"bypassPolicyLockoutSafetyCheck"`
	Encrypted                      *bool                               `pulumi:"encrypted"`
	FileSystemPolicy               interface{}                         `pulumi:"fileSystemPolicy"`
	FileSystemProtection           *FileSystemProtection               `pulumi:"fileSystemProtection"`
	FileSystemTags                 []FileSystemElasticFileSystemTag    `pulumi:"fileSystemTags"`
	KmsKeyId                       *string                             `pulumi:"kmsKeyId"`
	LifecyclePolicies              []FileSystemLifecyclePolicy         `pulumi:"lifecyclePolicies"`
	PerformanceMode                *string                             `pulumi:"performanceMode"`
	ProvisionedThroughputInMibps   *float64                            `pulumi:"provisionedThroughputInMibps"`
	ReplicationConfiguration       *FileSystemReplicationConfiguration `pulumi:"replicationConfiguration"`
	ThroughputMode                 *string                             `pulumi:"throughputMode"`
}

// The set of arguments for constructing a FileSystem resource.
type FileSystemArgs struct {
	AvailabilityZoneName pulumi.StringPtrInput
	BackupPolicy         FileSystemBackupPolicyPtrInput
	// Whether to bypass the FileSystemPolicy lockout safety check. The policy lockout safety check determines whether the policy in the request will prevent the principal making the request to be locked out from making future PutFileSystemPolicy requests on the file system. Set BypassPolicyLockoutSafetyCheck to True only when you intend to prevent the principal that is making the request from making a subsequent PutFileSystemPolicy request on the file system. Defaults to false
	BypassPolicyLockoutSafetyCheck pulumi.BoolPtrInput
	Encrypted                      pulumi.BoolPtrInput
	FileSystemPolicy               pulumi.Input
	FileSystemProtection           FileSystemProtectionPtrInput
	FileSystemTags                 FileSystemElasticFileSystemTagArrayInput
	KmsKeyId                       pulumi.StringPtrInput
	LifecyclePolicies              FileSystemLifecyclePolicyArrayInput
	PerformanceMode                pulumi.StringPtrInput
	ProvisionedThroughputInMibps   pulumi.Float64PtrInput
	ReplicationConfiguration       FileSystemReplicationConfigurationPtrInput
	ThroughputMode                 pulumi.StringPtrInput
}

func (FileSystemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fileSystemArgs)(nil)).Elem()
}

type FileSystemInput interface {
	pulumi.Input

	ToFileSystemOutput() FileSystemOutput
	ToFileSystemOutputWithContext(ctx context.Context) FileSystemOutput
}

func (*FileSystem) ElementType() reflect.Type {
	return reflect.TypeOf((**FileSystem)(nil)).Elem()
}

func (i *FileSystem) ToFileSystemOutput() FileSystemOutput {
	return i.ToFileSystemOutputWithContext(context.Background())
}

func (i *FileSystem) ToFileSystemOutputWithContext(ctx context.Context) FileSystemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileSystemOutput)
}

func (i *FileSystem) ToOutput(ctx context.Context) pulumix.Output[*FileSystem] {
	return pulumix.Output[*FileSystem]{
		OutputState: i.ToFileSystemOutputWithContext(ctx).OutputState,
	}
}

type FileSystemOutput struct{ *pulumi.OutputState }

func (FileSystemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FileSystem)(nil)).Elem()
}

func (o FileSystemOutput) ToFileSystemOutput() FileSystemOutput {
	return o
}

func (o FileSystemOutput) ToFileSystemOutputWithContext(ctx context.Context) FileSystemOutput {
	return o
}

func (o FileSystemOutput) ToOutput(ctx context.Context) pulumix.Output[*FileSystem] {
	return pulumix.Output[*FileSystem]{
		OutputState: o.OutputState,
	}
}

func (o FileSystemOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v *FileSystem) pulumi.StringOutput { return v.Arn }).(pulumi.StringOutput)
}

func (o FileSystemOutput) AvailabilityZoneName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FileSystem) pulumi.StringPtrOutput { return v.AvailabilityZoneName }).(pulumi.StringPtrOutput)
}

func (o FileSystemOutput) BackupPolicy() FileSystemBackupPolicyPtrOutput {
	return o.ApplyT(func(v *FileSystem) FileSystemBackupPolicyPtrOutput { return v.BackupPolicy }).(FileSystemBackupPolicyPtrOutput)
}

// Whether to bypass the FileSystemPolicy lockout safety check. The policy lockout safety check determines whether the policy in the request will prevent the principal making the request to be locked out from making future PutFileSystemPolicy requests on the file system. Set BypassPolicyLockoutSafetyCheck to True only when you intend to prevent the principal that is making the request from making a subsequent PutFileSystemPolicy request on the file system. Defaults to false
func (o FileSystemOutput) BypassPolicyLockoutSafetyCheck() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FileSystem) pulumi.BoolPtrOutput { return v.BypassPolicyLockoutSafetyCheck }).(pulumi.BoolPtrOutput)
}

func (o FileSystemOutput) Encrypted() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FileSystem) pulumi.BoolPtrOutput { return v.Encrypted }).(pulumi.BoolPtrOutput)
}

func (o FileSystemOutput) FileSystemId() pulumi.StringOutput {
	return o.ApplyT(func(v *FileSystem) pulumi.StringOutput { return v.FileSystemId }).(pulumi.StringOutput)
}

func (o FileSystemOutput) FileSystemPolicy() pulumi.AnyOutput {
	return o.ApplyT(func(v *FileSystem) pulumi.AnyOutput { return v.FileSystemPolicy }).(pulumi.AnyOutput)
}

func (o FileSystemOutput) FileSystemProtection() FileSystemProtectionPtrOutput {
	return o.ApplyT(func(v *FileSystem) FileSystemProtectionPtrOutput { return v.FileSystemProtection }).(FileSystemProtectionPtrOutput)
}

func (o FileSystemOutput) FileSystemTags() FileSystemElasticFileSystemTagArrayOutput {
	return o.ApplyT(func(v *FileSystem) FileSystemElasticFileSystemTagArrayOutput { return v.FileSystemTags }).(FileSystemElasticFileSystemTagArrayOutput)
}

func (o FileSystemOutput) KmsKeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FileSystem) pulumi.StringPtrOutput { return v.KmsKeyId }).(pulumi.StringPtrOutput)
}

func (o FileSystemOutput) LifecyclePolicies() FileSystemLifecyclePolicyArrayOutput {
	return o.ApplyT(func(v *FileSystem) FileSystemLifecyclePolicyArrayOutput { return v.LifecyclePolicies }).(FileSystemLifecyclePolicyArrayOutput)
}

func (o FileSystemOutput) PerformanceMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FileSystem) pulumi.StringPtrOutput { return v.PerformanceMode }).(pulumi.StringPtrOutput)
}

func (o FileSystemOutput) ProvisionedThroughputInMibps() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *FileSystem) pulumi.Float64PtrOutput { return v.ProvisionedThroughputInMibps }).(pulumi.Float64PtrOutput)
}

func (o FileSystemOutput) ReplicationConfiguration() FileSystemReplicationConfigurationPtrOutput {
	return o.ApplyT(func(v *FileSystem) FileSystemReplicationConfigurationPtrOutput { return v.ReplicationConfiguration }).(FileSystemReplicationConfigurationPtrOutput)
}

func (o FileSystemOutput) ThroughputMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FileSystem) pulumi.StringPtrOutput { return v.ThroughputMode }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FileSystemInput)(nil)).Elem(), &FileSystem{})
	pulumi.RegisterOutputType(FileSystemOutput{})
}
