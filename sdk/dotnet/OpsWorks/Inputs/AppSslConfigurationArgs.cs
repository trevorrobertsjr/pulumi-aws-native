// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.OpsWorks.Inputs
{

    public sealed class AppSslConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("certificate")]
        public Input<string>? Certificate { get; set; }

        [Input("chain")]
        public Input<string>? Chain { get; set; }

        [Input("privateKey")]
        public Input<string>? PrivateKey { get; set; }

        public AppSslConfigurationArgs()
        {
        }
        public static new AppSslConfigurationArgs Empty => new AppSslConfigurationArgs();
    }
}
