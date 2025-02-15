// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.KinesisFirehose.Inputs
{

    public sealed class DeliveryStreamProcessorParameterArgs : global::Pulumi.ResourceArgs
    {
        [Input("parameterName", required: true)]
        public Input<string> ParameterName { get; set; } = null!;

        [Input("parameterValue", required: true)]
        public Input<string> ParameterValue { get; set; } = null!;

        public DeliveryStreamProcessorParameterArgs()
        {
        }
        public static new DeliveryStreamProcessorParameterArgs Empty => new DeliveryStreamProcessorParameterArgs();
    }
}
