// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::AppConfig::Extension
 */
export class Extension extends pulumi.CustomResource {
    /**
     * Get an existing Extension resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Extension {
        return new Extension(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:appconfig:Extension';

    /**
     * Returns true if the given object is an instance of Extension.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Extension {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Extension.__pulumiType;
    }

    public readonly actions!: pulumi.Output<any>;
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Description of the extension.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly latestVersionNumber!: pulumi.Output<number | undefined>;
    /**
     * Name of the extension.
     */
    public readonly name!: pulumi.Output<string>;
    public readonly parameters!: pulumi.Output<any | undefined>;
    /**
     * An array of key-value tags to apply to this resource.
     */
    public readonly tags!: pulumi.Output<outputs.appconfig.ExtensionTag[] | undefined>;
    public /*out*/ readonly versionNumber!: pulumi.Output<number>;

    /**
     * Create a Extension resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ExtensionArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.actions === undefined) && !opts.urn) {
                throw new Error("Missing required property 'actions'");
            }
            resourceInputs["actions"] = args ? args.actions : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["latestVersionNumber"] = args ? args.latestVersionNumber : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["parameters"] = args ? args.parameters : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["versionNumber"] = undefined /*out*/;
        } else {
            resourceInputs["actions"] = undefined /*out*/;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["latestVersionNumber"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["parameters"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["versionNumber"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["name", "tags[*]"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Extension.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Extension resource.
 */
export interface ExtensionArgs {
    actions: any;
    /**
     * Description of the extension.
     */
    description?: pulumi.Input<string>;
    latestVersionNumber?: pulumi.Input<number>;
    /**
     * Name of the extension.
     */
    name?: pulumi.Input<string>;
    parameters?: any;
    /**
     * An array of key-value tags to apply to this resource.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.appconfig.ExtensionTagArgs>[]>;
}
