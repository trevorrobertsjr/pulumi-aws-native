// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AmazonMq.Inputs
{

    public sealed class BrokerLogListArgs : global::Pulumi.ResourceArgs
    {
        [Input("audit")]
        public Input<bool>? Audit { get; set; }

        [Input("general")]
        public Input<bool>? General { get; set; }

        public BrokerLogListArgs()
        {
        }
        public static new BrokerLogListArgs Empty => new BrokerLogListArgs();
    }
}
