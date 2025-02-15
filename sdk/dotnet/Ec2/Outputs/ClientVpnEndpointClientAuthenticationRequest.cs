// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ec2.Outputs
{

    [OutputType]
    public sealed class ClientVpnEndpointClientAuthenticationRequest
    {
        public readonly Outputs.ClientVpnEndpointDirectoryServiceAuthenticationRequest? ActiveDirectory;
        public readonly Outputs.ClientVpnEndpointFederatedAuthenticationRequest? FederatedAuthentication;
        public readonly Outputs.ClientVpnEndpointCertificateAuthenticationRequest? MutualAuthentication;
        public readonly string Type;

        [OutputConstructor]
        private ClientVpnEndpointClientAuthenticationRequest(
            Outputs.ClientVpnEndpointDirectoryServiceAuthenticationRequest? activeDirectory,

            Outputs.ClientVpnEndpointFederatedAuthenticationRequest? federatedAuthentication,

            Outputs.ClientVpnEndpointCertificateAuthenticationRequest? mutualAuthentication,

            string type)
        {
            ActiveDirectory = activeDirectory;
            FederatedAuthentication = federatedAuthentication;
            MutualAuthentication = mutualAuthentication;
            Type = type;
        }
    }
}
