// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Emr.Inputs
{

    public sealed class ClusterInstanceFleetConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("instanceTypeConfigs")]
        private InputList<Inputs.ClusterInstanceTypeConfigArgs>? _instanceTypeConfigs;
        public InputList<Inputs.ClusterInstanceTypeConfigArgs> InstanceTypeConfigs
        {
            get => _instanceTypeConfigs ?? (_instanceTypeConfigs = new InputList<Inputs.ClusterInstanceTypeConfigArgs>());
            set => _instanceTypeConfigs = value;
        }

        [Input("launchSpecifications")]
        public Input<Inputs.ClusterInstanceFleetProvisioningSpecificationsArgs>? LaunchSpecifications { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("targetOnDemandCapacity")]
        public Input<int>? TargetOnDemandCapacity { get; set; }

        [Input("targetSpotCapacity")]
        public Input<int>? TargetSpotCapacity { get; set; }

        public ClusterInstanceFleetConfigArgs()
        {
        }
        public static new ClusterInstanceFleetConfigArgs Empty => new ClusterInstanceFleetConfigArgs();
    }
}
