// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CleanRooms.Inputs
{

    public sealed class CollaborationPaymentConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("queryCompute", required: true)]
        public Input<Inputs.CollaborationQueryComputePaymentConfigArgs> QueryCompute { get; set; } = null!;

        public CollaborationPaymentConfigurationArgs()
        {
        }
        public static new CollaborationPaymentConfigurationArgs Empty => new CollaborationPaymentConfigurationArgs();
    }
}
