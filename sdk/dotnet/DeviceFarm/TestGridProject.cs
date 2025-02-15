// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.DeviceFarm
{
    /// <summary>
    /// AWS::DeviceFarm::TestGridProject creates a new TestGrid Project
    /// </summary>
    [AwsNativeResourceType("aws-native:devicefarm:TestGridProject")]
    public partial class TestGridProject : global::Pulumi.CustomResource
    {
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<Outputs.TestGridProjectTag>> Tags { get; private set; } = null!;

        [Output("vpcConfig")]
        public Output<Outputs.TestGridProjectVpcConfig?> VpcConfig { get; private set; } = null!;


        /// <summary>
        /// Create a TestGridProject resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TestGridProject(string name, TestGridProjectArgs? args = null, CustomResourceOptions? options = null)
            : base("aws-native:devicefarm:TestGridProject", name, args ?? new TestGridProjectArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TestGridProject(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:devicefarm:TestGridProject", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing TestGridProject resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TestGridProject Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new TestGridProject(name, id, options);
        }
    }

    public sealed class TestGridProjectArgs : global::Pulumi.ResourceArgs
    {
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputList<Inputs.TestGridProjectTagArgs>? _tags;
        public InputList<Inputs.TestGridProjectTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.TestGridProjectTagArgs>());
            set => _tags = value;
        }

        [Input("vpcConfig")]
        public Input<Inputs.TestGridProjectVpcConfigArgs>? VpcConfig { get; set; }

        public TestGridProjectArgs()
        {
        }
        public static new TestGridProjectArgs Empty => new TestGridProjectArgs();
    }
}
