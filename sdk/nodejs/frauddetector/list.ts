// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * A resource schema for a List in Amazon Fraud Detector.
 */
export class List extends pulumi.CustomResource {
    /**
     * Get an existing List resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): List {
        return new List(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:frauddetector:List';

    /**
     * Returns true if the given object is an instance of List.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is List {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === List.__pulumiType;
    }

    /**
     * The list ARN.
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The time when the list was created.
     */
    public /*out*/ readonly createdTime!: pulumi.Output<string>;
    /**
     * The description of the list.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The elements in this list.
     */
    public readonly elements!: pulumi.Output<string[] | undefined>;
    /**
     * The time when the list was last updated.
     */
    public /*out*/ readonly lastUpdatedTime!: pulumi.Output<string>;
    /**
     * The name of the list.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Tags associated with this list.
     */
    public readonly tags!: pulumi.Output<outputs.frauddetector.ListTag[] | undefined>;
    /**
     * The variable type of the list.
     */
    public readonly variableType!: pulumi.Output<string | undefined>;

    /**
     * Create a List resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ListArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["elements"] = args ? args.elements : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["variableType"] = args ? args.variableType : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["createdTime"] = undefined /*out*/;
            resourceInputs["lastUpdatedTime"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["createdTime"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["elements"] = undefined /*out*/;
            resourceInputs["lastUpdatedTime"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["variableType"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["name"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(List.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a List resource.
 */
export interface ListArgs {
    /**
     * The description of the list.
     */
    description?: pulumi.Input<string>;
    /**
     * The elements in this list.
     */
    elements?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the list.
     */
    name?: pulumi.Input<string>;
    /**
     * Tags associated with this list.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.frauddetector.ListTagArgs>[]>;
    /**
     * The variable type of the list.
     */
    variableType?: pulumi.Input<string>;
}
