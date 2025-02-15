// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.GameLift.Inputs
{

    /// <summary>
    /// A remote location where a multi-location fleet can deploy EC2 instances for game hosting.
    /// </summary>
    public sealed class FleetLocationConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("locationCapacity")]
        public Input<Inputs.FleetLocationCapacityArgs>? LocationCapacity { get; set; }

        public FleetLocationConfigurationArgs()
        {
        }
        public static new FleetLocationConfigurationArgs Empty => new FleetLocationConfigurationArgs();
    }
}
