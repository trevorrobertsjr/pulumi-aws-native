// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.VerifiedPermissions.Inputs
{

    public sealed class IdentitySourceConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("cognitoUserPoolConfiguration", required: true)]
        public Input<Inputs.IdentitySourceCognitoUserPoolConfigurationArgs> CognitoUserPoolConfiguration { get; set; } = null!;

        public IdentitySourceConfigurationArgs()
        {
        }
        public static new IdentitySourceConfigurationArgs Empty => new IdentitySourceConfigurationArgs();
    }
}
