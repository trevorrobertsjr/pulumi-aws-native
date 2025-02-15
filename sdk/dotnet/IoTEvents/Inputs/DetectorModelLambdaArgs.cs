// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.IoTEvents.Inputs
{

    public sealed class DetectorModelLambdaArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the Lambda function that is executed.
        /// </summary>
        [Input("functionArn", required: true)]
        public Input<string> FunctionArn { get; set; } = null!;

        [Input("payload")]
        public Input<Inputs.DetectorModelPayloadArgs>? Payload { get; set; }

        public DetectorModelLambdaArgs()
        {
        }
        public static new DetectorModelLambdaArgs Empty => new DetectorModelLambdaArgs();
    }
}
