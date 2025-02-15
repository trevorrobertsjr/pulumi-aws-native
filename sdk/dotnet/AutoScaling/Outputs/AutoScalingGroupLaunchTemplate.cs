// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AutoScaling.Outputs
{

    [OutputType]
    public sealed class AutoScalingGroupLaunchTemplate
    {
        public readonly Outputs.AutoScalingGroupLaunchTemplateSpecification LaunchTemplateSpecification;
        public readonly ImmutableArray<Outputs.AutoScalingGroupLaunchTemplateOverrides> Overrides;

        [OutputConstructor]
        private AutoScalingGroupLaunchTemplate(
            Outputs.AutoScalingGroupLaunchTemplateSpecification launchTemplateSpecification,

            ImmutableArray<Outputs.AutoScalingGroupLaunchTemplateOverrides> overrides)
        {
            LaunchTemplateSpecification = launchTemplateSpecification;
            Overrides = overrides;
        }
    }
}
