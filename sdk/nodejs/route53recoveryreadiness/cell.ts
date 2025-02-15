// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * The API Schema for AWS Route53 Recovery Readiness Cells.
 */
export class Cell extends pulumi.CustomResource {
    /**
     * Get an existing Cell resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Cell {
        return new Cell(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:route53recoveryreadiness:Cell';

    /**
     * Returns true if the given object is an instance of Cell.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Cell {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Cell.__pulumiType;
    }

    /**
     * The Amazon Resource Name (ARN) of the cell.
     */
    public /*out*/ readonly cellArn!: pulumi.Output<string>;
    /**
     * The name of the cell to create.
     */
    public readonly cellName!: pulumi.Output<string | undefined>;
    /**
     * A list of cell Amazon Resource Names (ARNs) contained within this cell, for use in nested cells. For example, Availability Zones within specific Regions.
     */
    public readonly cells!: pulumi.Output<string[] | undefined>;
    /**
     * The readiness scope for the cell, which can be a cell Amazon Resource Name (ARN) or a recovery group ARN. This is a list but currently can have only one element.
     */
    public /*out*/ readonly parentReadinessScopes!: pulumi.Output<string[]>;
    /**
     * A collection of tags associated with a resource
     */
    public readonly tags!: pulumi.Output<outputs.route53recoveryreadiness.CellTag[] | undefined>;

    /**
     * Create a Cell resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: CellArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["cellName"] = args ? args.cellName : undefined;
            resourceInputs["cells"] = args ? args.cells : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["cellArn"] = undefined /*out*/;
            resourceInputs["parentReadinessScopes"] = undefined /*out*/;
        } else {
            resourceInputs["cellArn"] = undefined /*out*/;
            resourceInputs["cellName"] = undefined /*out*/;
            resourceInputs["cells"] = undefined /*out*/;
            resourceInputs["parentReadinessScopes"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["cellName"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Cell.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Cell resource.
 */
export interface CellArgs {
    /**
     * The name of the cell to create.
     */
    cellName?: pulumi.Input<string>;
    /**
     * A list of cell Amazon Resource Names (ARNs) contained within this cell, for use in nested cells. For example, Availability Zones within specific Regions.
     */
    cells?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A collection of tags associated with a resource
     */
    tags?: pulumi.Input<pulumi.Input<inputs.route53recoveryreadiness.CellTagArgs>[]>;
}
