// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Budgets
{
    public static class GetBudgetsAction
    {
        /// <summary>
        /// An example resource schema demonstrating some basic constructs and validation rules.
        /// </summary>
        public static Task<GetBudgetsActionResult> InvokeAsync(GetBudgetsActionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetBudgetsActionResult>("aws-native:budgets:getBudgetsAction", args ?? new GetBudgetsActionArgs(), options.WithDefaults());

        /// <summary>
        /// An example resource schema demonstrating some basic constructs and validation rules.
        /// </summary>
        public static Output<GetBudgetsActionResult> Invoke(GetBudgetsActionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetBudgetsActionResult>("aws-native:budgets:getBudgetsAction", args ?? new GetBudgetsActionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBudgetsActionArgs : global::Pulumi.InvokeArgs
    {
        [Input("actionId", required: true)]
        public string ActionId { get; set; } = null!;

        [Input("budgetName", required: true)]
        public string BudgetName { get; set; } = null!;

        public GetBudgetsActionArgs()
        {
        }
        public static new GetBudgetsActionArgs Empty => new GetBudgetsActionArgs();
    }

    public sealed class GetBudgetsActionInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("actionId", required: true)]
        public Input<string> ActionId { get; set; } = null!;

        [Input("budgetName", required: true)]
        public Input<string> BudgetName { get; set; } = null!;

        public GetBudgetsActionInvokeArgs()
        {
        }
        public static new GetBudgetsActionInvokeArgs Empty => new GetBudgetsActionInvokeArgs();
    }


    [OutputType]
    public sealed class GetBudgetsActionResult
    {
        public readonly string? ActionId;
        public readonly Outputs.BudgetsActionActionThreshold? ActionThreshold;
        public readonly Pulumi.AwsNative.Budgets.BudgetsActionApprovalModel? ApprovalModel;
        public readonly Outputs.BudgetsActionDefinition? Definition;
        public readonly string? ExecutionRoleArn;
        public readonly Pulumi.AwsNative.Budgets.BudgetsActionNotificationType? NotificationType;
        public readonly ImmutableArray<Outputs.BudgetsActionSubscriber> Subscribers;

        [OutputConstructor]
        private GetBudgetsActionResult(
            string? actionId,

            Outputs.BudgetsActionActionThreshold? actionThreshold,

            Pulumi.AwsNative.Budgets.BudgetsActionApprovalModel? approvalModel,

            Outputs.BudgetsActionDefinition? definition,

            string? executionRoleArn,

            Pulumi.AwsNative.Budgets.BudgetsActionNotificationType? notificationType,

            ImmutableArray<Outputs.BudgetsActionSubscriber> subscribers)
        {
            ActionId = actionId;
            ActionThreshold = actionThreshold;
            ApprovalModel = approvalModel;
            Definition = definition;
            ExecutionRoleArn = executionRoleArn;
            NotificationType = notificationType;
            Subscribers = subscribers;
        }
    }
}
