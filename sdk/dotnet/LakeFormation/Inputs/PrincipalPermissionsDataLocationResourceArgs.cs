// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.LakeFormation.Inputs
{

    public sealed class PrincipalPermissionsDataLocationResourceArgs : global::Pulumi.ResourceArgs
    {
        [Input("catalogId", required: true)]
        public Input<string> CatalogId { get; set; } = null!;

        [Input("resourceArn", required: true)]
        public Input<string> ResourceArn { get; set; } = null!;

        public PrincipalPermissionsDataLocationResourceArgs()
        {
        }
        public static new PrincipalPermissionsDataLocationResourceArgs Empty => new PrincipalPermissionsDataLocationResourceArgs();
    }
}
