// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WaFv2.Outputs
{

    [OutputType]
    public sealed class RuleGroupRegexMatchStatement
    {
        public readonly Outputs.RuleGroupFieldToMatch FieldToMatch;
        public readonly string RegexString;
        public readonly ImmutableArray<Outputs.RuleGroupTextTransformation> TextTransformations;

        [OutputConstructor]
        private RuleGroupRegexMatchStatement(
            Outputs.RuleGroupFieldToMatch fieldToMatch,

            string regexString,

            ImmutableArray<Outputs.RuleGroupTextTransformation> textTransformations)
        {
            FieldToMatch = fieldToMatch;
            RegexString = regexString;
            TextTransformations = textTransformations;
        }
    }
}
