// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::LakeFormation::Resource
 *
 * @deprecated Resource is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class Resource extends pulumi.CustomResource {
    /**
     * Get an existing Resource resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Resource {
        pulumi.log.warn("Resource is deprecated: Resource is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new Resource(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:lakeformation:Resource';

    /**
     * Returns true if the given object is an instance of Resource.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Resource {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Resource.__pulumiType;
    }

    public readonly resourceArn!: pulumi.Output<string>;
    public readonly roleArn!: pulumi.Output<string | undefined>;
    public readonly useServiceLinkedRole!: pulumi.Output<boolean>;
    public readonly withFederation!: pulumi.Output<boolean | undefined>;

    /**
     * Create a Resource resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated Resource is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args: ResourceArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("Resource is deprecated: Resource is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.resourceArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceArn'");
            }
            if ((!args || args.useServiceLinkedRole === undefined) && !opts.urn) {
                throw new Error("Missing required property 'useServiceLinkedRole'");
            }
            resourceInputs["resourceArn"] = args ? args.resourceArn : undefined;
            resourceInputs["roleArn"] = args ? args.roleArn : undefined;
            resourceInputs["useServiceLinkedRole"] = args ? args.useServiceLinkedRole : undefined;
            resourceInputs["withFederation"] = args ? args.withFederation : undefined;
        } else {
            resourceInputs["resourceArn"] = undefined /*out*/;
            resourceInputs["roleArn"] = undefined /*out*/;
            resourceInputs["useServiceLinkedRole"] = undefined /*out*/;
            resourceInputs["withFederation"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["resourceArn"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Resource.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Resource resource.
 */
export interface ResourceArgs {
    resourceArn: pulumi.Input<string>;
    roleArn?: pulumi.Input<string>;
    useServiceLinkedRole: pulumi.Input<boolean>;
    withFederation?: pulumi.Input<boolean>;
}
