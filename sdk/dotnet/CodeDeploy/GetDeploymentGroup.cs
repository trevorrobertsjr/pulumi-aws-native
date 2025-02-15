// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CodeDeploy
{
    public static class GetDeploymentGroup
    {
        /// <summary>
        /// Resource Type definition for AWS::CodeDeploy::DeploymentGroup
        /// </summary>
        public static Task<GetDeploymentGroupResult> InvokeAsync(GetDeploymentGroupArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDeploymentGroupResult>("aws-native:codedeploy:getDeploymentGroup", args ?? new GetDeploymentGroupArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::CodeDeploy::DeploymentGroup
        /// </summary>
        public static Output<GetDeploymentGroupResult> Invoke(GetDeploymentGroupInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDeploymentGroupResult>("aws-native:codedeploy:getDeploymentGroup", args ?? new GetDeploymentGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDeploymentGroupArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetDeploymentGroupArgs()
        {
        }
        public static new GetDeploymentGroupArgs Empty => new GetDeploymentGroupArgs();
    }

    public sealed class GetDeploymentGroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetDeploymentGroupInvokeArgs()
        {
        }
        public static new GetDeploymentGroupInvokeArgs Empty => new GetDeploymentGroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetDeploymentGroupResult
    {
        public readonly Outputs.DeploymentGroupAlarmConfiguration? AlarmConfiguration;
        public readonly Outputs.DeploymentGroupAutoRollbackConfiguration? AutoRollbackConfiguration;
        public readonly ImmutableArray<string> AutoScalingGroups;
        public readonly Outputs.DeploymentGroupBlueGreenDeploymentConfiguration? BlueGreenDeploymentConfiguration;
        public readonly Outputs.DeploymentGroupDeployment? Deployment;
        public readonly string? DeploymentConfigName;
        public readonly Outputs.DeploymentGroupDeploymentStyle? DeploymentStyle;
        public readonly ImmutableArray<Outputs.DeploymentGroupEc2TagFilter> Ec2TagFilters;
        public readonly Outputs.DeploymentGroupEc2TagSet? Ec2TagSet;
        public readonly ImmutableArray<Outputs.DeploymentGroupEcsService> EcsServices;
        public readonly string? Id;
        public readonly Outputs.DeploymentGroupLoadBalancerInfo? LoadBalancerInfo;
        public readonly ImmutableArray<Outputs.DeploymentGroupTagFilter> OnPremisesInstanceTagFilters;
        public readonly Outputs.DeploymentGroupOnPremisesTagSet? OnPremisesTagSet;
        public readonly string? OutdatedInstancesStrategy;
        public readonly string? ServiceRoleArn;
        public readonly ImmutableArray<Outputs.DeploymentGroupTag> Tags;
        public readonly bool? TerminationHookEnabled;
        public readonly ImmutableArray<Outputs.DeploymentGroupTriggerConfig> TriggerConfigurations;

        [OutputConstructor]
        private GetDeploymentGroupResult(
            Outputs.DeploymentGroupAlarmConfiguration? alarmConfiguration,

            Outputs.DeploymentGroupAutoRollbackConfiguration? autoRollbackConfiguration,

            ImmutableArray<string> autoScalingGroups,

            Outputs.DeploymentGroupBlueGreenDeploymentConfiguration? blueGreenDeploymentConfiguration,

            Outputs.DeploymentGroupDeployment? deployment,

            string? deploymentConfigName,

            Outputs.DeploymentGroupDeploymentStyle? deploymentStyle,

            ImmutableArray<Outputs.DeploymentGroupEc2TagFilter> ec2TagFilters,

            Outputs.DeploymentGroupEc2TagSet? ec2TagSet,

            ImmutableArray<Outputs.DeploymentGroupEcsService> ecsServices,

            string? id,

            Outputs.DeploymentGroupLoadBalancerInfo? loadBalancerInfo,

            ImmutableArray<Outputs.DeploymentGroupTagFilter> onPremisesInstanceTagFilters,

            Outputs.DeploymentGroupOnPremisesTagSet? onPremisesTagSet,

            string? outdatedInstancesStrategy,

            string? serviceRoleArn,

            ImmutableArray<Outputs.DeploymentGroupTag> tags,

            bool? terminationHookEnabled,

            ImmutableArray<Outputs.DeploymentGroupTriggerConfig> triggerConfigurations)
        {
            AlarmConfiguration = alarmConfiguration;
            AutoRollbackConfiguration = autoRollbackConfiguration;
            AutoScalingGroups = autoScalingGroups;
            BlueGreenDeploymentConfiguration = blueGreenDeploymentConfiguration;
            Deployment = deployment;
            DeploymentConfigName = deploymentConfigName;
            DeploymentStyle = deploymentStyle;
            Ec2TagFilters = ec2TagFilters;
            Ec2TagSet = ec2TagSet;
            EcsServices = ecsServices;
            Id = id;
            LoadBalancerInfo = loadBalancerInfo;
            OnPremisesInstanceTagFilters = onPremisesInstanceTagFilters;
            OnPremisesTagSet = onPremisesTagSet;
            OutdatedInstancesStrategy = outdatedInstancesStrategy;
            ServiceRoleArn = serviceRoleArn;
            Tags = tags;
            TerminationHookEnabled = terminationHookEnabled;
            TriggerConfigurations = triggerConfigurations;
        }
    }
}
