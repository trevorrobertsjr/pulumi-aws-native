// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DynamoDb.Inputs
{

    public sealed class GlobalTableReadProvisionedThroughputSettingsArgs : global::Pulumi.ResourceArgs
    {
        [Input("readCapacityAutoScalingSettings")]
        public Input<Inputs.GlobalTableCapacityAutoScalingSettingsArgs>? ReadCapacityAutoScalingSettings { get; set; }

        [Input("readCapacityUnits")]
        public Input<int>? ReadCapacityUnits { get; set; }

        public GlobalTableReadProvisionedThroughputSettingsArgs()
        {
        }
        public static new GlobalTableReadProvisionedThroughputSettingsArgs Empty => new GlobalTableReadProvisionedThroughputSettingsArgs();
    }
}
