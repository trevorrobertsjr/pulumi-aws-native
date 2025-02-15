// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.ElasticLoadBalancingV2.Outputs
{

    [OutputType]
    public sealed class ListenerRuleAction
    {
        public readonly Outputs.ListenerRuleAuthenticateCognitoConfig? AuthenticateCognitoConfig;
        public readonly Outputs.ListenerRuleAuthenticateOidcConfig? AuthenticateOidcConfig;
        public readonly Outputs.ListenerRuleFixedResponseConfig? FixedResponseConfig;
        public readonly Outputs.ListenerRuleForwardConfig? ForwardConfig;
        public readonly int? Order;
        public readonly Outputs.ListenerRuleRedirectConfig? RedirectConfig;
        public readonly string? TargetGroupArn;
        public readonly string Type;

        [OutputConstructor]
        private ListenerRuleAction(
            Outputs.ListenerRuleAuthenticateCognitoConfig? authenticateCognitoConfig,

            Outputs.ListenerRuleAuthenticateOidcConfig? authenticateOidcConfig,

            Outputs.ListenerRuleFixedResponseConfig? fixedResponseConfig,

            Outputs.ListenerRuleForwardConfig? forwardConfig,

            int? order,

            Outputs.ListenerRuleRedirectConfig? redirectConfig,

            string? targetGroupArn,

            string type)
        {
            AuthenticateCognitoConfig = authenticateCognitoConfig;
            AuthenticateOidcConfig = authenticateOidcConfig;
            FixedResponseConfig = fixedResponseConfig;
            ForwardConfig = forwardConfig;
            Order = order;
            RedirectConfig = redirectConfig;
            TargetGroupArn = targetGroupArn;
            Type = type;
        }
    }
}
