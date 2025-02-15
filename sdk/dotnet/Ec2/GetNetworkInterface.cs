// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2
{
    public static class GetNetworkInterface
    {
        /// <summary>
        /// The AWS::EC2::NetworkInterface resource creates network interface
        /// </summary>
        public static Task<GetNetworkInterfaceResult> InvokeAsync(GetNetworkInterfaceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetNetworkInterfaceResult>("aws-native:ec2:getNetworkInterface", args ?? new GetNetworkInterfaceArgs(), options.WithDefaults());

        /// <summary>
        /// The AWS::EC2::NetworkInterface resource creates network interface
        /// </summary>
        public static Output<GetNetworkInterfaceResult> Invoke(GetNetworkInterfaceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetNetworkInterfaceResult>("aws-native:ec2:getNetworkInterface", args ?? new GetNetworkInterfaceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNetworkInterfaceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Network interface id.
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetNetworkInterfaceArgs()
        {
        }
        public static new GetNetworkInterfaceArgs Empty => new GetNetworkInterfaceArgs();
    }

    public sealed class GetNetworkInterfaceInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Network interface id.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetNetworkInterfaceInvokeArgs()
        {
        }
        public static new GetNetworkInterfaceInvokeArgs Empty => new GetNetworkInterfaceInvokeArgs();
    }


    [OutputType]
    public sealed class GetNetworkInterfaceResult
    {
        /// <summary>
        /// A description for the network interface.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// A list of security group IDs associated with this network interface.
        /// </summary>
        public readonly ImmutableArray<string> GroupSet;
        /// <summary>
        /// Network interface id.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The number of IPv4 prefixes to assign to a network interface. When you specify a number of IPv4 prefixes, Amazon EC2 selects these prefixes from your existing subnet CIDR reservations, if available, or from free spaces in the subnet. By default, these will be /28 prefixes. You can't specify a count of IPv4 prefixes if you've specified one of the following: specific IPv4 prefixes, specific private IPv4 addresses, or a count of private IPv4 addresses.
        /// </summary>
        public readonly int? Ipv4PrefixCount;
        /// <summary>
        /// Assigns a list of IPv4 prefixes to the network interface. If you want EC2 to automatically assign IPv4 prefixes, use the Ipv4PrefixCount property and do not specify this property. Presently, only /28 prefixes are supported. You can't specify IPv4 prefixes if you've specified one of the following: a count of IPv4 prefixes, specific private IPv4 addresses, or a count of private IPv4 addresses.
        /// </summary>
        public readonly ImmutableArray<Outputs.NetworkInterfaceIpv4PrefixSpecification> Ipv4Prefixes;
        /// <summary>
        /// The number of IPv6 addresses to assign to a network interface. Amazon EC2 automatically selects the IPv6 addresses from the subnet range. To specify specific IPv6 addresses, use the Ipv6Addresses property and don't specify this property.
        /// </summary>
        public readonly int? Ipv6AddressCount;
        /// <summary>
        /// One or more specific IPv6 addresses from the IPv6 CIDR block range of your subnet to associate with the network interface. If you're specifying a number of IPv6 addresses, use the Ipv6AddressCount property and don't specify this property.
        /// </summary>
        public readonly ImmutableArray<Outputs.NetworkInterfaceInstanceIpv6Address> Ipv6Addresses;
        /// <summary>
        /// The number of IPv6 prefixes to assign to a network interface. When you specify a number of IPv6 prefixes, Amazon EC2 selects these prefixes from your existing subnet CIDR reservations, if available, or from free spaces in the subnet. By default, these will be /80 prefixes. You can't specify a count of IPv6 prefixes if you've specified one of the following: specific IPv6 prefixes, specific IPv6 addresses, or a count of IPv6 addresses.
        /// </summary>
        public readonly int? Ipv6PrefixCount;
        /// <summary>
        /// Assigns a list of IPv6 prefixes to the network interface. If you want EC2 to automatically assign IPv6 prefixes, use the Ipv6PrefixCount property and do not specify this property. Presently, only /80 prefixes are supported. You can't specify IPv6 prefixes if you've specified one of the following: a count of IPv6 prefixes, specific IPv6 addresses, or a count of IPv6 addresses.
        /// </summary>
        public readonly ImmutableArray<Outputs.NetworkInterfaceIpv6PrefixSpecification> Ipv6Prefixes;
        /// <summary>
        /// Returns the primary private IP address of the network interface.
        /// </summary>
        public readonly string? PrimaryPrivateIpAddress;
        /// <summary>
        /// Assigns a list of private IP addresses to the network interface. You can specify a primary private IP address by setting the value of the Primary property to true in the PrivateIpAddressSpecification property. If you want EC2 to automatically assign private IP addresses, use the SecondaryPrivateIpAddressCount property and do not specify this property.
        /// </summary>
        public readonly ImmutableArray<Outputs.NetworkInterfacePrivateIpAddressSpecification> PrivateIpAddresses;
        /// <summary>
        /// The number of secondary private IPv4 addresses to assign to a network interface. When you specify a number of secondary IPv4 addresses, Amazon EC2 selects these IP addresses within the subnet's IPv4 CIDR range. You can't specify this option and specify more than one private IP address using privateIpAddresses
        /// </summary>
        public readonly int? SecondaryPrivateIpAddressCount;
        /// <summary>
        /// Returns the secondary private IP addresses of the network interface.
        /// </summary>
        public readonly ImmutableArray<string> SecondaryPrivateIpAddresses;
        /// <summary>
        /// Indicates whether traffic to or from the instance is validated.
        /// </summary>
        public readonly bool? SourceDestCheck;
        /// <summary>
        /// An arbitrary set of tags (key-value pairs) for this network interface.
        /// </summary>
        public readonly ImmutableArray<Outputs.NetworkInterfaceTag> Tags;

        [OutputConstructor]
        private GetNetworkInterfaceResult(
            string? description,

            ImmutableArray<string> groupSet,

            string? id,

            int? ipv4PrefixCount,

            ImmutableArray<Outputs.NetworkInterfaceIpv4PrefixSpecification> ipv4Prefixes,

            int? ipv6AddressCount,

            ImmutableArray<Outputs.NetworkInterfaceInstanceIpv6Address> ipv6Addresses,

            int? ipv6PrefixCount,

            ImmutableArray<Outputs.NetworkInterfaceIpv6PrefixSpecification> ipv6Prefixes,

            string? primaryPrivateIpAddress,

            ImmutableArray<Outputs.NetworkInterfacePrivateIpAddressSpecification> privateIpAddresses,

            int? secondaryPrivateIpAddressCount,

            ImmutableArray<string> secondaryPrivateIpAddresses,

            bool? sourceDestCheck,

            ImmutableArray<Outputs.NetworkInterfaceTag> tags)
        {
            Description = description;
            GroupSet = groupSet;
            Id = id;
            Ipv4PrefixCount = ipv4PrefixCount;
            Ipv4Prefixes = ipv4Prefixes;
            Ipv6AddressCount = ipv6AddressCount;
            Ipv6Addresses = ipv6Addresses;
            Ipv6PrefixCount = ipv6PrefixCount;
            Ipv6Prefixes = ipv6Prefixes;
            PrimaryPrivateIpAddress = primaryPrivateIpAddress;
            PrivateIpAddresses = privateIpAddresses;
            SecondaryPrivateIpAddressCount = secondaryPrivateIpAddressCount;
            SecondaryPrivateIpAddresses = secondaryPrivateIpAddresses;
            SourceDestCheck = sourceDestCheck;
            Tags = tags;
        }
    }
}
