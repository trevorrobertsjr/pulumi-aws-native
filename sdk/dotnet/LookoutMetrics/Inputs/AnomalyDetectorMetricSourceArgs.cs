// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.LookoutMetrics.Inputs
{

    public sealed class AnomalyDetectorMetricSourceArgs : global::Pulumi.ResourceArgs
    {
        [Input("appFlowConfig")]
        public Input<Inputs.AnomalyDetectorAppFlowConfigArgs>? AppFlowConfig { get; set; }

        [Input("cloudwatchConfig")]
        public Input<Inputs.AnomalyDetectorCloudwatchConfigArgs>? CloudwatchConfig { get; set; }

        [Input("rdsSourceConfig")]
        public Input<Inputs.AnomalyDetectorRdsSourceConfigArgs>? RdsSourceConfig { get; set; }

        [Input("redshiftSourceConfig")]
        public Input<Inputs.AnomalyDetectorRedshiftSourceConfigArgs>? RedshiftSourceConfig { get; set; }

        [Input("s3SourceConfig")]
        public Input<Inputs.AnomalyDetectorS3SourceConfigArgs>? S3SourceConfig { get; set; }

        public AnomalyDetectorMetricSourceArgs()
        {
        }
        public static new AnomalyDetectorMetricSourceArgs Empty => new AnomalyDetectorMetricSourceArgs();
    }
}
