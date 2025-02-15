// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of AWS::Location::APIKey Resource Type
 */
export class ApiKey extends pulumi.CustomResource {
    /**
     * Get an existing ApiKey resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ApiKey {
        return new ApiKey(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:location:ApiKey';

    /**
     * Returns true if the given object is an instance of ApiKey.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApiKey {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApiKey.__pulumiType;
    }

    public /*out*/ readonly arn!: pulumi.Output<string>;
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly expireTime!: pulumi.Output<string | undefined>;
    public readonly forceDelete!: pulumi.Output<boolean | undefined>;
    public readonly forceUpdate!: pulumi.Output<boolean | undefined>;
    public /*out*/ readonly keyArn!: pulumi.Output<string>;
    public readonly keyName!: pulumi.Output<string>;
    public readonly noExpiry!: pulumi.Output<boolean | undefined>;
    public readonly restrictions!: pulumi.Output<outputs.location.ApiKeyRestrictions>;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    public readonly tags!: pulumi.Output<outputs.location.ApiKeyTag[] | undefined>;
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a ApiKey resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApiKeyArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.keyName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'keyName'");
            }
            if ((!args || args.restrictions === undefined) && !opts.urn) {
                throw new Error("Missing required property 'restrictions'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["expireTime"] = args ? args.expireTime : undefined;
            resourceInputs["forceDelete"] = args ? args.forceDelete : undefined;
            resourceInputs["forceUpdate"] = args ? args.forceUpdate : undefined;
            resourceInputs["keyName"] = args ? args.keyName : undefined;
            resourceInputs["noExpiry"] = args ? args.noExpiry : undefined;
            resourceInputs["restrictions"] = args ? args.restrictions : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["keyArn"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["expireTime"] = undefined /*out*/;
            resourceInputs["forceDelete"] = undefined /*out*/;
            resourceInputs["forceUpdate"] = undefined /*out*/;
            resourceInputs["keyArn"] = undefined /*out*/;
            resourceInputs["keyName"] = undefined /*out*/;
            resourceInputs["noExpiry"] = undefined /*out*/;
            resourceInputs["restrictions"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["keyName"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(ApiKey.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ApiKey resource.
 */
export interface ApiKeyArgs {
    description?: pulumi.Input<string>;
    expireTime?: pulumi.Input<string>;
    forceDelete?: pulumi.Input<boolean>;
    forceUpdate?: pulumi.Input<boolean>;
    keyName: pulumi.Input<string>;
    noExpiry?: pulumi.Input<boolean>;
    restrictions: pulumi.Input<inputs.location.ApiKeyRestrictionsArgs>;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.location.ApiKeyTagArgs>[]>;
}
