// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppFlow.Inputs
{

    public sealed class FlowLookoutMetricsDestinationPropertiesArgs : global::Pulumi.ResourceArgs
    {
        [Input("object")]
        public Input<string>? Object { get; set; }

        public FlowLookoutMetricsDestinationPropertiesArgs()
        {
        }
        public static new FlowLookoutMetricsDestinationPropertiesArgs Empty => new FlowLookoutMetricsDestinationPropertiesArgs();
    }
}
