// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Outputs
{

    [OutputType]
    public sealed class AnalysisResourcePermission
    {
        public readonly ImmutableArray<string> Actions;
        public readonly string Principal;
        public readonly string? Resource;

        [OutputConstructor]
        private AnalysisResourcePermission(
            ImmutableArray<string> actions,

            string principal,

            string? resource)
        {
            Actions = actions;
            Principal = principal;
            Resource = resource;
        }
    }
}
