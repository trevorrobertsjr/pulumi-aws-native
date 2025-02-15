// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ecs.Inputs
{

    /// <summary>
    /// The health check command and associated configuration parameters for the container.
    /// </summary>
    public sealed class TaskDefinitionHealthCheckArgs : global::Pulumi.ResourceArgs
    {
        [Input("command")]
        private InputList<string>? _command;

        /// <summary>
        /// A string array representing the command that the container runs to determine if it is healthy.
        /// </summary>
        public InputList<string> Command
        {
            get => _command ?? (_command = new InputList<string>());
            set => _command = value;
        }

        /// <summary>
        /// The time period in seconds between each health check execution. You may specify between 5 and 300 seconds. The default value is 30 seconds.
        /// </summary>
        [Input("interval")]
        public Input<int>? Interval { get; set; }

        /// <summary>
        /// The number of times to retry a failed health check before the container is considered unhealthy. You may specify between 1 and 10 retries. The default value is three retries.
        /// </summary>
        [Input("retries")]
        public Input<int>? Retries { get; set; }

        /// <summary>
        /// The optional grace period within which to provide containers time to bootstrap before failed health checks count towards the maximum number of retries. You may specify between 0 and 300 seconds. The startPeriod is disabled by default.
        /// </summary>
        [Input("startPeriod")]
        public Input<int>? StartPeriod { get; set; }

        /// <summary>
        /// The time period in seconds to wait for a health check to succeed before it is considered a failure. You may specify between 2 and 60 seconds. The default value is 5 seconds.
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        public TaskDefinitionHealthCheckArgs()
        {
        }
        public static new TaskDefinitionHealthCheckArgs Empty => new TaskDefinitionHealthCheckArgs();
    }
}
