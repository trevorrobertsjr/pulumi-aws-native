// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Forecast.Inputs
{

    public sealed class EncryptionConfigPropertiesArgs : global::Pulumi.ResourceArgs
    {
        [Input("kmsKeyArn")]
        public Input<string>? KmsKeyArn { get; set; }

        [Input("roleArn")]
        public Input<string>? RoleArn { get; set; }

        public EncryptionConfigPropertiesArgs()
        {
        }
        public static new EncryptionConfigPropertiesArgs Empty => new EncryptionConfigPropertiesArgs();
    }
}
