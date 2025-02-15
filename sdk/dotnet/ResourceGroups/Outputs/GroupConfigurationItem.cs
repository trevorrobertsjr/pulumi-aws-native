// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ResourceGroups.Outputs
{

    [OutputType]
    public sealed class GroupConfigurationItem
    {
        public readonly ImmutableArray<Outputs.GroupConfigurationParameter> Parameters;
        public readonly string? Type;

        [OutputConstructor]
        private GroupConfigurationItem(
            ImmutableArray<Outputs.GroupConfigurationParameter> parameters,

            string? type)
        {
            Parameters = parameters;
            Type = type;
        }
    }
}
