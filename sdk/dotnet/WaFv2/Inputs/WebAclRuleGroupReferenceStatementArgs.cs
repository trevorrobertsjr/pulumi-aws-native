// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WaFv2.Inputs
{

    public sealed class WebAclRuleGroupReferenceStatementArgs : global::Pulumi.ResourceArgs
    {
        [Input("arn", required: true)]
        public Input<string> Arn { get; set; } = null!;

        [Input("excludedRules")]
        private InputList<Inputs.WebAclExcludedRuleArgs>? _excludedRules;
        public InputList<Inputs.WebAclExcludedRuleArgs> ExcludedRules
        {
            get => _excludedRules ?? (_excludedRules = new InputList<Inputs.WebAclExcludedRuleArgs>());
            set => _excludedRules = value;
        }

        [Input("ruleActionOverrides")]
        private InputList<Inputs.WebAclRuleActionOverrideArgs>? _ruleActionOverrides;

        /// <summary>
        /// Action overrides for rules in the rule group.
        /// </summary>
        public InputList<Inputs.WebAclRuleActionOverrideArgs> RuleActionOverrides
        {
            get => _ruleActionOverrides ?? (_ruleActionOverrides = new InputList<Inputs.WebAclRuleActionOverrideArgs>());
            set => _ruleActionOverrides = value;
        }

        public WebAclRuleGroupReferenceStatementArgs()
        {
        }
        public static new WebAclRuleGroupReferenceStatementArgs Empty => new WebAclRuleGroupReferenceStatementArgs();
    }
}
