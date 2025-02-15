// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Sqs
{
    public static class GetQueuePolicy
    {
        /// <summary>
        /// The ``AWS::SQS::QueuePolicy`` type applies a policy to SQS queues. For an example snippet, see [Declaring an policy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/quickref-iam.html#scenario-sqs-policy) in the *User Guide*.
        /// </summary>
        public static Task<GetQueuePolicyResult> InvokeAsync(GetQueuePolicyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetQueuePolicyResult>("aws-native:sqs:getQueuePolicy", args ?? new GetQueuePolicyArgs(), options.WithDefaults());

        /// <summary>
        /// The ``AWS::SQS::QueuePolicy`` type applies a policy to SQS queues. For an example snippet, see [Declaring an policy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/quickref-iam.html#scenario-sqs-policy) in the *User Guide*.
        /// </summary>
        public static Output<GetQueuePolicyResult> Invoke(GetQueuePolicyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetQueuePolicyResult>("aws-native:sqs:getQueuePolicy", args ?? new GetQueuePolicyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetQueuePolicyArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetQueuePolicyArgs()
        {
        }
        public static new GetQueuePolicyArgs Empty => new GetQueuePolicyArgs();
    }

    public sealed class GetQueuePolicyInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetQueuePolicyInvokeArgs()
        {
        }
        public static new GetQueuePolicyInvokeArgs Empty => new GetQueuePolicyInvokeArgs();
    }


    [OutputType]
    public sealed class GetQueuePolicyResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// A policy document that contains the permissions for the specified SQS queues. For more information about SQS policies, see [Using custom policies with the access policy language](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-creating-custom-policies.html) in the *Developer Guide*.
        /// </summary>
        public readonly object? PolicyDocument;
        /// <summary>
        /// The URLs of the queues to which you want to add the policy. You can use the ``Ref`` function to specify an ``AWS::SQS::Queue`` resource.
        /// </summary>
        public readonly ImmutableArray<string> Queues;

        [OutputConstructor]
        private GetQueuePolicyResult(
            string? id,

            object? policyDocument,

            ImmutableArray<string> queues)
        {
            Id = id;
            PolicyDocument = policyDocument;
            Queues = queues;
        }
    }
}
