// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource Type definition for AWS::EC2::NetworkAclEntry
//
// Deprecated: NetworkAclEntry is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
type NetworkAclEntry struct {
	pulumi.CustomResourceState

	// The IPv4 CIDR range to allow or deny, in CIDR notation (for example, 172.16.0.0/24). Requirement is conditional: You must specify the CidrBlock or Ipv6CidrBlock property
	CidrBlock pulumi.StringPtrOutput `pulumi:"cidrBlock"`
	// Indicates whether this is an egress rule (rule is applied to traffic leaving the subnet)
	Egress pulumi.BoolPtrOutput `pulumi:"egress"`
	// The Internet Control Message Protocol (ICMP) code and type. Requirement is conditional: Required if specifying 1 (ICMP) for the protocol parameter
	Icmp NetworkAclEntryIcmpPtrOutput `pulumi:"icmp"`
	// The IPv6 network range to allow or deny, in CIDR notation (for example 2001:db8:1234:1a00::/64)
	Ipv6CidrBlock pulumi.StringPtrOutput `pulumi:"ipv6CidrBlock"`
	// The ID of the network ACL
	NetworkAclId pulumi.StringOutput `pulumi:"networkAclId"`
	// The IPv4 network range to allow or deny, in CIDR notation (for example 172.16.0.0/24). We modify the specified CIDR block to its canonical form; for example, if you specify 100.68.0.18/18, we modify it to 100.68.0.0/18
	PortRange NetworkAclEntryPortRangePtrOutput `pulumi:"portRange"`
	// The protocol number. A value of "-1" means all protocols. If you specify "-1" or a protocol number other than "6" (TCP), "17" (UDP), or "1" (ICMP), traffic on all ports is allowed, regardless of any ports or ICMP types or codes that you specify. If you specify protocol "58" (ICMPv6) and specify an IPv4 CIDR block, traffic for all ICMP types and codes allowed, regardless of any that you specify. If you specify protocol "58" (ICMPv6) and specify an IPv6 CIDR block, you must specify an ICMP type and code
	Protocol pulumi.IntOutput `pulumi:"protocol"`
	// Indicates whether to allow or deny the traffic that matches the rule
	RuleAction pulumi.StringOutput `pulumi:"ruleAction"`
	// Rule number to assign to the entry, such as 100. ACL entries are processed in ascending order by rule number. Entries can't use the same rule number unless one is an egress rule and the other is an ingress rule
	RuleNumber pulumi.IntOutput `pulumi:"ruleNumber"`
}

// NewNetworkAclEntry registers a new resource with the given unique name, arguments, and options.
func NewNetworkAclEntry(ctx *pulumi.Context,
	name string, args *NetworkAclEntryArgs, opts ...pulumi.ResourceOption) (*NetworkAclEntry, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NetworkAclId == nil {
		return nil, errors.New("invalid value for required argument 'NetworkAclId'")
	}
	if args.Protocol == nil {
		return nil, errors.New("invalid value for required argument 'Protocol'")
	}
	if args.RuleAction == nil {
		return nil, errors.New("invalid value for required argument 'RuleAction'")
	}
	if args.RuleNumber == nil {
		return nil, errors.New("invalid value for required argument 'RuleNumber'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"egress",
		"networkAclId",
		"ruleNumber",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource NetworkAclEntry
	err := ctx.RegisterResource("aws-native:ec2:NetworkAclEntry", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkAclEntry gets an existing NetworkAclEntry resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkAclEntry(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkAclEntryState, opts ...pulumi.ResourceOption) (*NetworkAclEntry, error) {
	var resource NetworkAclEntry
	err := ctx.ReadResource("aws-native:ec2:NetworkAclEntry", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkAclEntry resources.
type networkAclEntryState struct {
}

type NetworkAclEntryState struct {
}

func (NetworkAclEntryState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkAclEntryState)(nil)).Elem()
}

type networkAclEntryArgs struct {
	// The IPv4 CIDR range to allow or deny, in CIDR notation (for example, 172.16.0.0/24). Requirement is conditional: You must specify the CidrBlock or Ipv6CidrBlock property
	CidrBlock *string `pulumi:"cidrBlock"`
	// Indicates whether this is an egress rule (rule is applied to traffic leaving the subnet)
	Egress *bool `pulumi:"egress"`
	// The Internet Control Message Protocol (ICMP) code and type. Requirement is conditional: Required if specifying 1 (ICMP) for the protocol parameter
	Icmp *NetworkAclEntryIcmp `pulumi:"icmp"`
	// The IPv6 network range to allow or deny, in CIDR notation (for example 2001:db8:1234:1a00::/64)
	Ipv6CidrBlock *string `pulumi:"ipv6CidrBlock"`
	// The ID of the network ACL
	NetworkAclId string `pulumi:"networkAclId"`
	// The IPv4 network range to allow or deny, in CIDR notation (for example 172.16.0.0/24). We modify the specified CIDR block to its canonical form; for example, if you specify 100.68.0.18/18, we modify it to 100.68.0.0/18
	PortRange *NetworkAclEntryPortRange `pulumi:"portRange"`
	// The protocol number. A value of "-1" means all protocols. If you specify "-1" or a protocol number other than "6" (TCP), "17" (UDP), or "1" (ICMP), traffic on all ports is allowed, regardless of any ports or ICMP types or codes that you specify. If you specify protocol "58" (ICMPv6) and specify an IPv4 CIDR block, traffic for all ICMP types and codes allowed, regardless of any that you specify. If you specify protocol "58" (ICMPv6) and specify an IPv6 CIDR block, you must specify an ICMP type and code
	Protocol int `pulumi:"protocol"`
	// Indicates whether to allow or deny the traffic that matches the rule
	RuleAction string `pulumi:"ruleAction"`
	// Rule number to assign to the entry, such as 100. ACL entries are processed in ascending order by rule number. Entries can't use the same rule number unless one is an egress rule and the other is an ingress rule
	RuleNumber int `pulumi:"ruleNumber"`
}

// The set of arguments for constructing a NetworkAclEntry resource.
type NetworkAclEntryArgs struct {
	// The IPv4 CIDR range to allow or deny, in CIDR notation (for example, 172.16.0.0/24). Requirement is conditional: You must specify the CidrBlock or Ipv6CidrBlock property
	CidrBlock pulumi.StringPtrInput
	// Indicates whether this is an egress rule (rule is applied to traffic leaving the subnet)
	Egress pulumi.BoolPtrInput
	// The Internet Control Message Protocol (ICMP) code and type. Requirement is conditional: Required if specifying 1 (ICMP) for the protocol parameter
	Icmp NetworkAclEntryIcmpPtrInput
	// The IPv6 network range to allow or deny, in CIDR notation (for example 2001:db8:1234:1a00::/64)
	Ipv6CidrBlock pulumi.StringPtrInput
	// The ID of the network ACL
	NetworkAclId pulumi.StringInput
	// The IPv4 network range to allow or deny, in CIDR notation (for example 172.16.0.0/24). We modify the specified CIDR block to its canonical form; for example, if you specify 100.68.0.18/18, we modify it to 100.68.0.0/18
	PortRange NetworkAclEntryPortRangePtrInput
	// The protocol number. A value of "-1" means all protocols. If you specify "-1" or a protocol number other than "6" (TCP), "17" (UDP), or "1" (ICMP), traffic on all ports is allowed, regardless of any ports or ICMP types or codes that you specify. If you specify protocol "58" (ICMPv6) and specify an IPv4 CIDR block, traffic for all ICMP types and codes allowed, regardless of any that you specify. If you specify protocol "58" (ICMPv6) and specify an IPv6 CIDR block, you must specify an ICMP type and code
	Protocol pulumi.IntInput
	// Indicates whether to allow or deny the traffic that matches the rule
	RuleAction pulumi.StringInput
	// Rule number to assign to the entry, such as 100. ACL entries are processed in ascending order by rule number. Entries can't use the same rule number unless one is an egress rule and the other is an ingress rule
	RuleNumber pulumi.IntInput
}

func (NetworkAclEntryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkAclEntryArgs)(nil)).Elem()
}

type NetworkAclEntryInput interface {
	pulumi.Input

	ToNetworkAclEntryOutput() NetworkAclEntryOutput
	ToNetworkAclEntryOutputWithContext(ctx context.Context) NetworkAclEntryOutput
}

func (*NetworkAclEntry) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkAclEntry)(nil)).Elem()
}

func (i *NetworkAclEntry) ToNetworkAclEntryOutput() NetworkAclEntryOutput {
	return i.ToNetworkAclEntryOutputWithContext(context.Background())
}

func (i *NetworkAclEntry) ToNetworkAclEntryOutputWithContext(ctx context.Context) NetworkAclEntryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkAclEntryOutput)
}

func (i *NetworkAclEntry) ToOutput(ctx context.Context) pulumix.Output[*NetworkAclEntry] {
	return pulumix.Output[*NetworkAclEntry]{
		OutputState: i.ToNetworkAclEntryOutputWithContext(ctx).OutputState,
	}
}

type NetworkAclEntryOutput struct{ *pulumi.OutputState }

func (NetworkAclEntryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkAclEntry)(nil)).Elem()
}

func (o NetworkAclEntryOutput) ToNetworkAclEntryOutput() NetworkAclEntryOutput {
	return o
}

func (o NetworkAclEntryOutput) ToNetworkAclEntryOutputWithContext(ctx context.Context) NetworkAclEntryOutput {
	return o
}

func (o NetworkAclEntryOutput) ToOutput(ctx context.Context) pulumix.Output[*NetworkAclEntry] {
	return pulumix.Output[*NetworkAclEntry]{
		OutputState: o.OutputState,
	}
}

// The IPv4 CIDR range to allow or deny, in CIDR notation (for example, 172.16.0.0/24). Requirement is conditional: You must specify the CidrBlock or Ipv6CidrBlock property
func (o NetworkAclEntryOutput) CidrBlock() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkAclEntry) pulumi.StringPtrOutput { return v.CidrBlock }).(pulumi.StringPtrOutput)
}

// Indicates whether this is an egress rule (rule is applied to traffic leaving the subnet)
func (o NetworkAclEntryOutput) Egress() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NetworkAclEntry) pulumi.BoolPtrOutput { return v.Egress }).(pulumi.BoolPtrOutput)
}

// The Internet Control Message Protocol (ICMP) code and type. Requirement is conditional: Required if specifying 1 (ICMP) for the protocol parameter
func (o NetworkAclEntryOutput) Icmp() NetworkAclEntryIcmpPtrOutput {
	return o.ApplyT(func(v *NetworkAclEntry) NetworkAclEntryIcmpPtrOutput { return v.Icmp }).(NetworkAclEntryIcmpPtrOutput)
}

// The IPv6 network range to allow or deny, in CIDR notation (for example 2001:db8:1234:1a00::/64)
func (o NetworkAclEntryOutput) Ipv6CidrBlock() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkAclEntry) pulumi.StringPtrOutput { return v.Ipv6CidrBlock }).(pulumi.StringPtrOutput)
}

// The ID of the network ACL
func (o NetworkAclEntryOutput) NetworkAclId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkAclEntry) pulumi.StringOutput { return v.NetworkAclId }).(pulumi.StringOutput)
}

// The IPv4 network range to allow or deny, in CIDR notation (for example 172.16.0.0/24). We modify the specified CIDR block to its canonical form; for example, if you specify 100.68.0.18/18, we modify it to 100.68.0.0/18
func (o NetworkAclEntryOutput) PortRange() NetworkAclEntryPortRangePtrOutput {
	return o.ApplyT(func(v *NetworkAclEntry) NetworkAclEntryPortRangePtrOutput { return v.PortRange }).(NetworkAclEntryPortRangePtrOutput)
}

// The protocol number. A value of "-1" means all protocols. If you specify "-1" or a protocol number other than "6" (TCP), "17" (UDP), or "1" (ICMP), traffic on all ports is allowed, regardless of any ports or ICMP types or codes that you specify. If you specify protocol "58" (ICMPv6) and specify an IPv4 CIDR block, traffic for all ICMP types and codes allowed, regardless of any that you specify. If you specify protocol "58" (ICMPv6) and specify an IPv6 CIDR block, you must specify an ICMP type and code
func (o NetworkAclEntryOutput) Protocol() pulumi.IntOutput {
	return o.ApplyT(func(v *NetworkAclEntry) pulumi.IntOutput { return v.Protocol }).(pulumi.IntOutput)
}

// Indicates whether to allow or deny the traffic that matches the rule
func (o NetworkAclEntryOutput) RuleAction() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkAclEntry) pulumi.StringOutput { return v.RuleAction }).(pulumi.StringOutput)
}

// Rule number to assign to the entry, such as 100. ACL entries are processed in ascending order by rule number. Entries can't use the same rule number unless one is an egress rule and the other is an ingress rule
func (o NetworkAclEntryOutput) RuleNumber() pulumi.IntOutput {
	return o.ApplyT(func(v *NetworkAclEntry) pulumi.IntOutput { return v.RuleNumber }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkAclEntryInput)(nil)).Elem(), &NetworkAclEntry{})
	pulumi.RegisterOutputType(NetworkAclEntryOutput{})
}
