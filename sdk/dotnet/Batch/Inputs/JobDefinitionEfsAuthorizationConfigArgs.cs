// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Batch.Inputs
{

    public sealed class JobDefinitionEfsAuthorizationConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("accessPointId")]
        public Input<string>? AccessPointId { get; set; }

        [Input("iam")]
        public Input<string>? Iam { get; set; }

        public JobDefinitionEfsAuthorizationConfigArgs()
        {
        }
        public static new JobDefinitionEfsAuthorizationConfigArgs Empty => new JobDefinitionEfsAuthorizationConfigArgs();
    }
}
