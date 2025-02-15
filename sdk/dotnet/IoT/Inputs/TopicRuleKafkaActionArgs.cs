// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoT.Inputs
{

    public sealed class TopicRuleKafkaActionArgs : global::Pulumi.ResourceArgs
    {
        [Input("clientProperties", required: true)]
        public Input<object> ClientProperties { get; set; } = null!;

        [Input("destinationArn", required: true)]
        public Input<string> DestinationArn { get; set; } = null!;

        [Input("headers")]
        private InputList<Inputs.TopicRuleKafkaActionHeaderArgs>? _headers;
        public InputList<Inputs.TopicRuleKafkaActionHeaderArgs> Headers
        {
            get => _headers ?? (_headers = new InputList<Inputs.TopicRuleKafkaActionHeaderArgs>());
            set => _headers = value;
        }

        [Input("key")]
        public Input<string>? Key { get; set; }

        [Input("partition")]
        public Input<string>? Partition { get; set; }

        [Input("topic", required: true)]
        public Input<string> Topic { get; set; } = null!;

        public TopicRuleKafkaActionArgs()
        {
        }
        public static new TopicRuleKafkaActionArgs Empty => new TopicRuleKafkaActionArgs();
    }
}
