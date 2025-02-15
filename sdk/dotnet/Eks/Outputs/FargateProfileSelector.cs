// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Eks.Outputs
{

    [OutputType]
    public sealed class FargateProfileSelector
    {
        public readonly ImmutableArray<Outputs.FargateProfileLabel> Labels;
        public readonly string Namespace;

        [OutputConstructor]
        private FargateProfileSelector(
            ImmutableArray<Outputs.FargateProfileLabel> labels,

            string @namespace)
        {
            Labels = labels;
            Namespace = @namespace;
        }
    }
}
