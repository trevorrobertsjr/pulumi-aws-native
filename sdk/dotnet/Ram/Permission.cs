// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Ram
{
    /// <summary>
    /// Resource type definition for AWS::RAM::Permission
    /// </summary>
    [AwsNativeResourceType("aws-native:ram:Permission")]
    public partial class Permission : global::Pulumi.CustomResource
    {
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// Set to true to use this as the default permission.
        /// </summary>
        [Output("isResourceTypeDefault")]
        public Output<bool> IsResourceTypeDefault { get; private set; } = null!;

        /// <summary>
        /// The name of the permission.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("permissionType")]
        public Output<string> PermissionType { get; private set; } = null!;

        /// <summary>
        /// Policy template for the permission.
        /// </summary>
        [Output("policyTemplate")]
        public Output<object> PolicyTemplate { get; private set; } = null!;

        /// <summary>
        /// The resource type this permission can be used with.
        /// </summary>
        [Output("resourceType")]
        public Output<string> ResourceType { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<Outputs.PermissionTag>> Tags { get; private set; } = null!;

        /// <summary>
        /// Version of the permission.
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;


        /// <summary>
        /// Create a Permission resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Permission(string name, PermissionArgs args, CustomResourceOptions? options = null)
            : base("aws-native:ram:Permission", name, args ?? new PermissionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Permission(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:ram:Permission", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "name",
                    "policyTemplate",
                    "resourceType",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Permission resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Permission Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Permission(name, id, options);
        }
    }

    public sealed class PermissionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the permission.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Policy template for the permission.
        /// </summary>
        [Input("policyTemplate", required: true)]
        public Input<object> PolicyTemplate { get; set; } = null!;

        /// <summary>
        /// The resource type this permission can be used with.
        /// </summary>
        [Input("resourceType", required: true)]
        public Input<string> ResourceType { get; set; } = null!;

        [Input("tags")]
        private InputList<Inputs.PermissionTagArgs>? _tags;
        public InputList<Inputs.PermissionTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.PermissionTagArgs>());
            set => _tags = value;
        }

        public PermissionArgs()
        {
        }
        public static new PermissionArgs Empty => new PermissionArgs();
    }
}
