// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lambda.Inputs
{

    /// <summary>
    /// The VPC security groups and subnets that are attached to a Lambda function. When you connect a function to a VPC, Lambda creates an elastic network interface for each combination of security group and subnet in the function's VPC configuration. The function can only access resources and the internet through that VPC.
    /// </summary>
    public sealed class FunctionVpcConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A boolean indicating whether IPv6 protocols will be allowed for dual stack subnets
        /// </summary>
        [Input("ipv6AllowedForDualStack")]
        public Input<bool>? Ipv6AllowedForDualStack { get; set; }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// A list of VPC security groups IDs.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        [Input("subnetIds")]
        private InputList<string>? _subnetIds;

        /// <summary>
        /// A list of VPC subnet IDs.
        /// </summary>
        public InputList<string> SubnetIds
        {
            get => _subnetIds ?? (_subnetIds = new InputList<string>());
            set => _subnetIds = value;
        }

        public FunctionVpcConfigArgs()
        {
        }
        public static new FunctionVpcConfigArgs Empty => new FunctionVpcConfigArgs();
    }
}
