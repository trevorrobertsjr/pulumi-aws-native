// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.WaFv2.Outputs
{

    /// <summary>
    /// Checks that the request has a valid token with an unexpired challenge timestamp and, if not, returns a browser challenge to the client.
    /// </summary>
    [OutputType]
    public sealed class RuleGroupChallengeAction
    {
        public readonly Outputs.RuleGroupCustomRequestHandling? CustomRequestHandling;

        [OutputConstructor]
        private RuleGroupChallengeAction(Outputs.RuleGroupCustomRequestHandling? customRequestHandling)
        {
            CustomRequestHandling = customRequestHandling;
        }
    }
}
