// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Logs
{
    /// <summary>
    /// The AWS::Logs::AccountPolicy resource specifies a CloudWatch Logs AccountPolicy.
    /// </summary>
    [AwsNativeResourceType("aws-native:logs:AccountPolicy")]
    public partial class AccountPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// User account id
        /// </summary>
        [Output("accountId")]
        public Output<string> AccountId { get; private set; } = null!;

        /// <summary>
        /// The body of the policy document you want to use for this topic.
        /// 
        /// You can only add one policy per PolicyType.
        /// 
        /// The policy must be in JSON string format.
        /// 
        /// Length Constraints: Maximum length of 30720
        /// </summary>
        [Output("policyDocument")]
        public Output<string> PolicyDocument { get; private set; } = null!;

        /// <summary>
        /// The name of the account policy
        /// </summary>
        [Output("policyName")]
        public Output<string> PolicyName { get; private set; } = null!;

        /// <summary>
        /// Type of the policy.
        /// </summary>
        [Output("policyType")]
        public Output<Pulumi.AwsNative.Logs.AccountPolicyPolicyType> PolicyType { get; private set; } = null!;

        /// <summary>
        /// Scope for policy application
        /// </summary>
        [Output("scope")]
        public Output<Pulumi.AwsNative.Logs.AccountPolicyScope?> Scope { get; private set; } = null!;


        /// <summary>
        /// Create a AccountPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AccountPolicy(string name, AccountPolicyArgs args, CustomResourceOptions? options = null)
            : base("aws-native:logs:AccountPolicy", name, args ?? new AccountPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AccountPolicy(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:logs:AccountPolicy", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "policyName",
                    "policyType",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AccountPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AccountPolicy Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new AccountPolicy(name, id, options);
        }
    }

    public sealed class AccountPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The body of the policy document you want to use for this topic.
        /// 
        /// You can only add one policy per PolicyType.
        /// 
        /// The policy must be in JSON string format.
        /// 
        /// Length Constraints: Maximum length of 30720
        /// </summary>
        [Input("policyDocument", required: true)]
        public Input<string> PolicyDocument { get; set; } = null!;

        /// <summary>
        /// The name of the account policy
        /// </summary>
        [Input("policyName", required: true)]
        public Input<string> PolicyName { get; set; } = null!;

        /// <summary>
        /// Type of the policy.
        /// </summary>
        [Input("policyType", required: true)]
        public Input<Pulumi.AwsNative.Logs.AccountPolicyPolicyType> PolicyType { get; set; } = null!;

        /// <summary>
        /// Scope for policy application
        /// </summary>
        [Input("scope")]
        public Input<Pulumi.AwsNative.Logs.AccountPolicyScope>? Scope { get; set; }

        public AccountPolicyArgs()
        {
        }
        public static new AccountPolicyArgs Empty => new AccountPolicyArgs();
    }
}
