// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CodeDeploy.Outputs
{

    [OutputType]
    public sealed class DeploymentGroupTargetGroupInfo
    {
        public readonly string? Name;

        [OutputConstructor]
        private DeploymentGroupTargetGroupInfo(string? name)
        {
            Name = name;
        }
    }
}
