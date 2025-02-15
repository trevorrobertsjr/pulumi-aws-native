// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::EC2::TransitGateway
 */
export class TransitGateway extends pulumi.CustomResource {
    /**
     * Get an existing TransitGateway resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): TransitGateway {
        return new TransitGateway(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:ec2:TransitGateway';

    /**
     * Returns true if the given object is an instance of TransitGateway.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TransitGateway {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TransitGateway.__pulumiType;
    }

    public readonly amazonSideAsn!: pulumi.Output<number | undefined>;
    public readonly associationDefaultRouteTableId!: pulumi.Output<string | undefined>;
    public readonly autoAcceptSharedAttachments!: pulumi.Output<string | undefined>;
    public readonly defaultRouteTableAssociation!: pulumi.Output<string | undefined>;
    public readonly defaultRouteTablePropagation!: pulumi.Output<string | undefined>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly dnsSupport!: pulumi.Output<string | undefined>;
    public readonly multicastSupport!: pulumi.Output<string | undefined>;
    public readonly propagationDefaultRouteTableId!: pulumi.Output<string | undefined>;
    public readonly tags!: pulumi.Output<outputs.ec2.TransitGatewayTag[] | undefined>;
    public /*out*/ readonly transitGatewayArn!: pulumi.Output<string>;
    public readonly transitGatewayCidrBlocks!: pulumi.Output<string[] | undefined>;
    public readonly vpnEcmpSupport!: pulumi.Output<string | undefined>;

    /**
     * Create a TransitGateway resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: TransitGatewayArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["amazonSideAsn"] = args ? args.amazonSideAsn : undefined;
            resourceInputs["associationDefaultRouteTableId"] = args ? args.associationDefaultRouteTableId : undefined;
            resourceInputs["autoAcceptSharedAttachments"] = args ? args.autoAcceptSharedAttachments : undefined;
            resourceInputs["defaultRouteTableAssociation"] = args ? args.defaultRouteTableAssociation : undefined;
            resourceInputs["defaultRouteTablePropagation"] = args ? args.defaultRouteTablePropagation : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["dnsSupport"] = args ? args.dnsSupport : undefined;
            resourceInputs["multicastSupport"] = args ? args.multicastSupport : undefined;
            resourceInputs["propagationDefaultRouteTableId"] = args ? args.propagationDefaultRouteTableId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["transitGatewayCidrBlocks"] = args ? args.transitGatewayCidrBlocks : undefined;
            resourceInputs["vpnEcmpSupport"] = args ? args.vpnEcmpSupport : undefined;
            resourceInputs["transitGatewayArn"] = undefined /*out*/;
        } else {
            resourceInputs["amazonSideAsn"] = undefined /*out*/;
            resourceInputs["associationDefaultRouteTableId"] = undefined /*out*/;
            resourceInputs["autoAcceptSharedAttachments"] = undefined /*out*/;
            resourceInputs["defaultRouteTableAssociation"] = undefined /*out*/;
            resourceInputs["defaultRouteTablePropagation"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["dnsSupport"] = undefined /*out*/;
            resourceInputs["multicastSupport"] = undefined /*out*/;
            resourceInputs["propagationDefaultRouteTableId"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["transitGatewayArn"] = undefined /*out*/;
            resourceInputs["transitGatewayCidrBlocks"] = undefined /*out*/;
            resourceInputs["vpnEcmpSupport"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["amazonSideAsn", "multicastSupport"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(TransitGateway.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a TransitGateway resource.
 */
export interface TransitGatewayArgs {
    amazonSideAsn?: pulumi.Input<number>;
    associationDefaultRouteTableId?: pulumi.Input<string>;
    autoAcceptSharedAttachments?: pulumi.Input<string>;
    defaultRouteTableAssociation?: pulumi.Input<string>;
    defaultRouteTablePropagation?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    dnsSupport?: pulumi.Input<string>;
    multicastSupport?: pulumi.Input<string>;
    propagationDefaultRouteTableId?: pulumi.Input<string>;
    tags?: pulumi.Input<pulumi.Input<inputs.ec2.TransitGatewayTagArgs>[]>;
    transitGatewayCidrBlocks?: pulumi.Input<pulumi.Input<string>[]>;
    vpnEcmpSupport?: pulumi.Input<string>;
}
