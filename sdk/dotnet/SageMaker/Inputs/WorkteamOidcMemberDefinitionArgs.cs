// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.SageMaker.Inputs
{

    public sealed class WorkteamOidcMemberDefinitionArgs : global::Pulumi.ResourceArgs
    {
        [Input("oidcGroups", required: true)]
        private InputList<string>? _oidcGroups;
        public InputList<string> OidcGroups
        {
            get => _oidcGroups ?? (_oidcGroups = new InputList<string>());
            set => _oidcGroups = value;
        }

        public WorkteamOidcMemberDefinitionArgs()
        {
        }
        public static new WorkteamOidcMemberDefinitionArgs Empty => new WorkteamOidcMemberDefinitionArgs();
    }
}
