// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Glue.Inputs
{

    public sealed class JobExecutionPropertyArgs : global::Pulumi.ResourceArgs
    {
        [Input("maxConcurrentRuns")]
        public Input<double>? MaxConcurrentRuns { get; set; }

        public JobExecutionPropertyArgs()
        {
        }
        public static new JobExecutionPropertyArgs Empty => new JobExecutionPropertyArgs();
    }
}
