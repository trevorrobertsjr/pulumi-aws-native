// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Configuration.Outputs
{

    [OutputType]
    public sealed class OrganizationConfigRuleOrganizationManagedRuleMetadata
    {
        public readonly string? Description;
        public readonly string? InputParameters;
        public readonly string? MaximumExecutionFrequency;
        public readonly string? ResourceIdScope;
        public readonly ImmutableArray<string> ResourceTypesScope;
        public readonly string RuleIdentifier;
        public readonly string? TagKeyScope;
        public readonly string? TagValueScope;

        [OutputConstructor]
        private OrganizationConfigRuleOrganizationManagedRuleMetadata(
            string? description,

            string? inputParameters,

            string? maximumExecutionFrequency,

            string? resourceIdScope,

            ImmutableArray<string> resourceTypesScope,

            string ruleIdentifier,

            string? tagKeyScope,

            string? tagValueScope)
        {
            Description = description;
            InputParameters = inputParameters;
            MaximumExecutionFrequency = maximumExecutionFrequency;
            ResourceIdScope = resourceIdScope;
            ResourceTypesScope = resourceTypesScope;
            RuleIdentifier = ruleIdentifier;
            TagKeyScope = tagKeyScope;
            TagValueScope = tagValueScope;
        }
    }
}
