// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.PinpointEmail.Inputs
{

    public sealed class ConfigurationSetDeliveryOptionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("sendingPoolName")]
        public Input<string>? SendingPoolName { get; set; }

        public ConfigurationSetDeliveryOptionsArgs()
        {
        }
        public static new ConfigurationSetDeliveryOptionsArgs Empty => new ConfigurationSetDeliveryOptionsArgs();
    }
}
