// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.RedshiftServerless.Outputs
{

    [OutputType]
    public sealed class WorkgroupEndpoint
    {
        public readonly string? Address;
        public readonly int? Port;
        public readonly ImmutableArray<Outputs.WorkgroupVpcEndpoint> VpcEndpoints;

        [OutputConstructor]
        private WorkgroupEndpoint(
            string? address,

            int? port,

            ImmutableArray<Outputs.WorkgroupVpcEndpoint> vpcEndpoints)
        {
            Address = address;
            Port = port;
            VpcEndpoints = vpcEndpoints;
        }
    }
}
