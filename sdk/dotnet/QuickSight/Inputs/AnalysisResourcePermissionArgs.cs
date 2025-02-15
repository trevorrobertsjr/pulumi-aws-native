// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.QuickSight.Inputs
{

    public sealed class AnalysisResourcePermissionArgs : global::Pulumi.ResourceArgs
    {
        [Input("actions", required: true)]
        private InputList<string>? _actions;
        public InputList<string> Actions
        {
            get => _actions ?? (_actions = new InputList<string>());
            set => _actions = value;
        }

        [Input("principal", required: true)]
        public Input<string> Principal { get; set; } = null!;

        [Input("resource")]
        public Input<string>? Resource { get; set; }

        public AnalysisResourcePermissionArgs()
        {
        }
        public static new AnalysisResourcePermissionArgs Empty => new AnalysisResourcePermissionArgs();
    }
}
