// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lightsail.Inputs
{

    /// <summary>
    /// Describes the origin resource of an Amazon Lightsail content delivery network (CDN) distribution.
    /// </summary>
    public sealed class DistributionInputOriginArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the origin resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The protocol that your Amazon Lightsail distribution uses when establishing a connection with your origin to pull content.
        /// </summary>
        [Input("protocolPolicy")]
        public Input<string>? ProtocolPolicy { get; set; }

        /// <summary>
        /// The AWS Region name of the origin resource.
        /// </summary>
        [Input("regionName")]
        public Input<string>? RegionName { get; set; }

        public DistributionInputOriginArgs()
        {
        }
        public static new DistributionInputOriginArgs Empty => new DistributionInputOriginArgs();
    }
}
