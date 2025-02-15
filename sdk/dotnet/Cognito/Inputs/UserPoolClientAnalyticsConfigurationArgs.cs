// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Cognito.Inputs
{

    public sealed class UserPoolClientAnalyticsConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("applicationArn")]
        public Input<string>? ApplicationArn { get; set; }

        [Input("applicationId")]
        public Input<string>? ApplicationId { get; set; }

        [Input("externalId")]
        public Input<string>? ExternalId { get; set; }

        [Input("roleArn")]
        public Input<string>? RoleArn { get; set; }

        [Input("userDataShared")]
        public Input<bool>? UserDataShared { get; set; }

        public UserPoolClientAnalyticsConfigurationArgs()
        {
        }
        public static new UserPoolClientAnalyticsConfigurationArgs Empty => new UserPoolClientAnalyticsConfigurationArgs();
    }
}
