// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Msk.Inputs
{

    public sealed class ClusterBrokerNodeGroupInfoArgs : global::Pulumi.ResourceArgs
    {
        [Input("brokerAzDistribution")]
        public Input<string>? BrokerAzDistribution { get; set; }

        [Input("clientSubnets", required: true)]
        private InputList<string>? _clientSubnets;
        public InputList<string> ClientSubnets
        {
            get => _clientSubnets ?? (_clientSubnets = new InputList<string>());
            set => _clientSubnets = value;
        }

        [Input("connectivityInfo")]
        public Input<Inputs.ClusterConnectivityInfoArgs>? ConnectivityInfo { get; set; }

        [Input("instanceType", required: true)]
        public Input<string> InstanceType { get; set; } = null!;

        [Input("securityGroups")]
        private InputList<string>? _securityGroups;
        public InputList<string> SecurityGroups
        {
            get => _securityGroups ?? (_securityGroups = new InputList<string>());
            set => _securityGroups = value;
        }

        [Input("storageInfo")]
        public Input<Inputs.ClusterStorageInfoArgs>? StorageInfo { get; set; }

        public ClusterBrokerNodeGroupInfoArgs()
        {
        }
        public static new ClusterBrokerNodeGroupInfoArgs Empty => new ClusterBrokerNodeGroupInfoArgs();
    }
}
