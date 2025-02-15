// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.AmplifyUiBuilder
{
    /// <summary>
    /// Definition of AWS::AmplifyUIBuilder::Component Resource Type
    /// </summary>
    [AwsNativeResourceType("aws-native:amplifyuibuilder:Component")]
    public partial class Component : global::Pulumi.CustomResource
    {
        [Output("appId")]
        public Output<string?> AppId { get; private set; } = null!;

        [Output("bindingProperties")]
        public Output<Outputs.ComponentBindingProperties> BindingProperties { get; private set; } = null!;

        [Output("children")]
        public Output<ImmutableArray<Outputs.ComponentChild>> Children { get; private set; } = null!;

        [Output("collectionProperties")]
        public Output<Outputs.ComponentCollectionProperties?> CollectionProperties { get; private set; } = null!;

        [Output("componentType")]
        public Output<string> ComponentType { get; private set; } = null!;

        [Output("environmentName")]
        public Output<string?> EnvironmentName { get; private set; } = null!;

        [Output("events")]
        public Output<Outputs.ComponentEvents?> Events { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("overrides")]
        public Output<Outputs.ComponentOverrides> Overrides { get; private set; } = null!;

        [Output("properties")]
        public Output<Outputs.ComponentProperties> Properties { get; private set; } = null!;

        [Output("schemaVersion")]
        public Output<string?> SchemaVersion { get; private set; } = null!;

        [Output("sourceId")]
        public Output<string?> SourceId { get; private set; } = null!;

        [Output("tags")]
        public Output<Outputs.ComponentTags?> Tags { get; private set; } = null!;

        [Output("variants")]
        public Output<ImmutableArray<Outputs.ComponentVariant>> Variants { get; private set; } = null!;


        /// <summary>
        /// Create a Component resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Component(string name, ComponentArgs args, CustomResourceOptions? options = null)
            : base("aws-native:amplifyuibuilder:Component", name, args ?? new ComponentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Component(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:amplifyuibuilder:Component", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "tags",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Component resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Component Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Component(name, id, options);
        }
    }

    public sealed class ComponentArgs : global::Pulumi.ResourceArgs
    {
        [Input("appId")]
        public Input<string>? AppId { get; set; }

        [Input("bindingProperties", required: true)]
        public Input<Inputs.ComponentBindingPropertiesArgs> BindingProperties { get; set; } = null!;

        [Input("children")]
        private InputList<Inputs.ComponentChildArgs>? _children;
        public InputList<Inputs.ComponentChildArgs> Children
        {
            get => _children ?? (_children = new InputList<Inputs.ComponentChildArgs>());
            set => _children = value;
        }

        [Input("collectionProperties")]
        public Input<Inputs.ComponentCollectionPropertiesArgs>? CollectionProperties { get; set; }

        [Input("componentType", required: true)]
        public Input<string> ComponentType { get; set; } = null!;

        [Input("environmentName")]
        public Input<string>? EnvironmentName { get; set; }

        [Input("events")]
        public Input<Inputs.ComponentEventsArgs>? Events { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("overrides", required: true)]
        public Input<Inputs.ComponentOverridesArgs> Overrides { get; set; } = null!;

        [Input("properties", required: true)]
        public Input<Inputs.ComponentPropertiesArgs> Properties { get; set; } = null!;

        [Input("schemaVersion")]
        public Input<string>? SchemaVersion { get; set; }

        [Input("sourceId")]
        public Input<string>? SourceId { get; set; }

        [Input("tags")]
        public Input<Inputs.ComponentTagsArgs>? Tags { get; set; }

        [Input("variants", required: true)]
        private InputList<Inputs.ComponentVariantArgs>? _variants;
        public InputList<Inputs.ComponentVariantArgs> Variants
        {
            get => _variants ?? (_variants = new InputList<Inputs.ComponentVariantArgs>());
            set => _variants = value;
        }

        public ComponentArgs()
        {
        }
        public static new ComponentArgs Empty => new ComponentArgs();
    }
}
