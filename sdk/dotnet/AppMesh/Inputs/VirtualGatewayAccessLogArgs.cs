// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AppMesh.Inputs
{

    public sealed class VirtualGatewayAccessLogArgs : global::Pulumi.ResourceArgs
    {
        [Input("file")]
        public Input<Inputs.VirtualGatewayFileAccessLogArgs>? File { get; set; }

        public VirtualGatewayAccessLogArgs()
        {
        }
        public static new VirtualGatewayAccessLogArgs Empty => new VirtualGatewayAccessLogArgs();
    }
}
