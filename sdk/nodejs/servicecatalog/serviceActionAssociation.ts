// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Schema for AWS::ServiceCatalog::ServiceActionAssociation
 */
export class ServiceActionAssociation extends pulumi.CustomResource {
    /**
     * Get an existing ServiceActionAssociation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ServiceActionAssociation {
        return new ServiceActionAssociation(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:servicecatalog:ServiceActionAssociation';

    /**
     * Returns true if the given object is an instance of ServiceActionAssociation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServiceActionAssociation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServiceActionAssociation.__pulumiType;
    }

    public readonly productId!: pulumi.Output<string>;
    public readonly provisioningArtifactId!: pulumi.Output<string>;
    public readonly serviceActionId!: pulumi.Output<string>;

    /**
     * Create a ServiceActionAssociation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceActionAssociationArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.productId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'productId'");
            }
            if ((!args || args.provisioningArtifactId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'provisioningArtifactId'");
            }
            if ((!args || args.serviceActionId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceActionId'");
            }
            resourceInputs["productId"] = args ? args.productId : undefined;
            resourceInputs["provisioningArtifactId"] = args ? args.provisioningArtifactId : undefined;
            resourceInputs["serviceActionId"] = args ? args.serviceActionId : undefined;
        } else {
            resourceInputs["productId"] = undefined /*out*/;
            resourceInputs["provisioningArtifactId"] = undefined /*out*/;
            resourceInputs["serviceActionId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["productId", "provisioningArtifactId", "serviceActionId"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(ServiceActionAssociation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ServiceActionAssociation resource.
 */
export interface ServiceActionAssociationArgs {
    productId: pulumi.Input<string>;
    provisioningArtifactId: pulumi.Input<string>;
    serviceActionId: pulumi.Input<string>;
}
