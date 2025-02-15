// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Connect.Inputs
{

    /// <summary>
    /// Contains information about the queue and channel for which priority and delay can be set.
    /// </summary>
    public sealed class RoutingProfileQueueConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("delay", required: true)]
        public Input<int> Delay { get; set; } = null!;

        [Input("priority", required: true)]
        public Input<int> Priority { get; set; } = null!;

        [Input("queueReference", required: true)]
        public Input<Inputs.RoutingProfileQueueReferenceArgs> QueueReference { get; set; } = null!;

        public RoutingProfileQueueConfigArgs()
        {
        }
        public static new RoutingProfileQueueConfigArgs Empty => new RoutingProfileQueueConfigArgs();
    }
}
