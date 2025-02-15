// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AmazonMq
{
    public static class GetBroker
    {
        /// <summary>
        /// Resource Type definition for AWS::AmazonMQ::Broker
        /// </summary>
        public static Task<GetBrokerResult> InvokeAsync(GetBrokerArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetBrokerResult>("aws-native:amazonmq:getBroker", args ?? new GetBrokerArgs(), options.WithDefaults());

        /// <summary>
        /// Resource Type definition for AWS::AmazonMQ::Broker
        /// </summary>
        public static Output<GetBrokerResult> Invoke(GetBrokerInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetBrokerResult>("aws-native:amazonmq:getBroker", args ?? new GetBrokerInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBrokerArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetBrokerArgs()
        {
        }
        public static new GetBrokerArgs Empty => new GetBrokerArgs();
    }

    public sealed class GetBrokerInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetBrokerInvokeArgs()
        {
        }
        public static new GetBrokerInvokeArgs Empty => new GetBrokerInvokeArgs();
    }


    [OutputType]
    public sealed class GetBrokerResult
    {
        public readonly ImmutableArray<string> AmqpEndpoints;
        public readonly string? Arn;
        public readonly bool? AutoMinorVersionUpgrade;
        public readonly Outputs.BrokerConfigurationId? Configuration;
        public readonly string? ConfigurationId;
        public readonly int? ConfigurationRevision;
        public readonly string? DataReplicationMode;
        public readonly string? DataReplicationPrimaryBrokerArn;
        public readonly string? EngineVersion;
        public readonly string? HostInstanceType;
        public readonly string? Id;
        public readonly ImmutableArray<string> IpAddresses;
        public readonly Outputs.BrokerLdapServerMetadata? LdapServerMetadata;
        public readonly Outputs.BrokerLogList? Logs;
        public readonly Outputs.BrokerMaintenanceWindow? MaintenanceWindowStartTime;
        public readonly ImmutableArray<string> MqttEndpoints;
        public readonly ImmutableArray<string> OpenWireEndpoints;
        public readonly ImmutableArray<string> SecurityGroups;
        public readonly ImmutableArray<string> StompEndpoints;
        public readonly ImmutableArray<Outputs.BrokerTagsEntry> Tags;
        public readonly ImmutableArray<Outputs.BrokerUser> Users;
        public readonly ImmutableArray<string> WssEndpoints;

        [OutputConstructor]
        private GetBrokerResult(
            ImmutableArray<string> amqpEndpoints,

            string? arn,

            bool? autoMinorVersionUpgrade,

            Outputs.BrokerConfigurationId? configuration,

            string? configurationId,

            int? configurationRevision,

            string? dataReplicationMode,

            string? dataReplicationPrimaryBrokerArn,

            string? engineVersion,

            string? hostInstanceType,

            string? id,

            ImmutableArray<string> ipAddresses,

            Outputs.BrokerLdapServerMetadata? ldapServerMetadata,

            Outputs.BrokerLogList? logs,

            Outputs.BrokerMaintenanceWindow? maintenanceWindowStartTime,

            ImmutableArray<string> mqttEndpoints,

            ImmutableArray<string> openWireEndpoints,

            ImmutableArray<string> securityGroups,

            ImmutableArray<string> stompEndpoints,

            ImmutableArray<Outputs.BrokerTagsEntry> tags,

            ImmutableArray<Outputs.BrokerUser> users,

            ImmutableArray<string> wssEndpoints)
        {
            AmqpEndpoints = amqpEndpoints;
            Arn = arn;
            AutoMinorVersionUpgrade = autoMinorVersionUpgrade;
            Configuration = configuration;
            ConfigurationId = configurationId;
            ConfigurationRevision = configurationRevision;
            DataReplicationMode = dataReplicationMode;
            DataReplicationPrimaryBrokerArn = dataReplicationPrimaryBrokerArn;
            EngineVersion = engineVersion;
            HostInstanceType = hostInstanceType;
            Id = id;
            IpAddresses = ipAddresses;
            LdapServerMetadata = ldapServerMetadata;
            Logs = logs;
            MaintenanceWindowStartTime = maintenanceWindowStartTime;
            MqttEndpoints = mqttEndpoints;
            OpenWireEndpoints = openWireEndpoints;
            SecurityGroups = securityGroups;
            StompEndpoints = stompEndpoints;
            Tags = tags;
            Users = users;
            WssEndpoints = wssEndpoints;
        }
    }
}
