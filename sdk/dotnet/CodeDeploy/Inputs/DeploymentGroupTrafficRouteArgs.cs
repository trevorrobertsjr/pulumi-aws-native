// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.CodeDeploy.Inputs
{

    public sealed class DeploymentGroupTrafficRouteArgs : global::Pulumi.ResourceArgs
    {
        [Input("listenerArns")]
        private InputList<string>? _listenerArns;
        public InputList<string> ListenerArns
        {
            get => _listenerArns ?? (_listenerArns = new InputList<string>());
            set => _listenerArns = value;
        }

        public DeploymentGroupTrafficRouteArgs()
        {
        }
        public static new DeploymentGroupTrafficRouteArgs Empty => new DeploymentGroupTrafficRouteArgs();
    }
}
