// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * This schema is for testing purpose only.
 */
export class SimulationApplication extends pulumi.CustomResource {
    /**
     * Get an existing SimulationApplication resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): SimulationApplication {
        return new SimulationApplication(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:robomaker:SimulationApplication';

    /**
     * Returns true if the given object is an instance of SimulationApplication.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SimulationApplication {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SimulationApplication.__pulumiType;
    }

    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * The current revision id.
     */
    public readonly currentRevisionId!: pulumi.Output<string | undefined>;
    /**
     * The URI of the Docker image for the robot application.
     */
    public readonly environment!: pulumi.Output<string | undefined>;
    /**
     * The name of the simulation application.
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * The rendering engine for the simulation application.
     */
    public readonly renderingEngine!: pulumi.Output<outputs.robomaker.SimulationApplicationRenderingEngine | undefined>;
    /**
     * The robot software suite used by the simulation application.
     */
    public readonly robotSoftwareSuite!: pulumi.Output<outputs.robomaker.SimulationApplicationRobotSoftwareSuite>;
    /**
     * The simulation software suite used by the simulation application.
     */
    public readonly simulationSoftwareSuite!: pulumi.Output<outputs.robomaker.SimulationApplicationSimulationSoftwareSuite>;
    /**
     * The sources of the simulation application.
     */
    public readonly sources!: pulumi.Output<outputs.robomaker.SimulationApplicationSourceConfig[] | undefined>;
    public readonly tags!: pulumi.Output<outputs.robomaker.SimulationApplicationTags | undefined>;

    /**
     * Create a SimulationApplication resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SimulationApplicationArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.robotSoftwareSuite === undefined) && !opts.urn) {
                throw new Error("Missing required property 'robotSoftwareSuite'");
            }
            if ((!args || args.simulationSoftwareSuite === undefined) && !opts.urn) {
                throw new Error("Missing required property 'simulationSoftwareSuite'");
            }
            resourceInputs["currentRevisionId"] = args ? args.currentRevisionId : undefined;
            resourceInputs["environment"] = args ? args.environment : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["renderingEngine"] = args ? args.renderingEngine : undefined;
            resourceInputs["robotSoftwareSuite"] = args ? args.robotSoftwareSuite : undefined;
            resourceInputs["simulationSoftwareSuite"] = args ? args.simulationSoftwareSuite : undefined;
            resourceInputs["sources"] = args ? args.sources : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
        } else {
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["currentRevisionId"] = undefined /*out*/;
            resourceInputs["environment"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["renderingEngine"] = undefined /*out*/;
            resourceInputs["robotSoftwareSuite"] = undefined /*out*/;
            resourceInputs["simulationSoftwareSuite"] = undefined /*out*/;
            resourceInputs["sources"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["name"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(SimulationApplication.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a SimulationApplication resource.
 */
export interface SimulationApplicationArgs {
    /**
     * The current revision id.
     */
    currentRevisionId?: pulumi.Input<string>;
    /**
     * The URI of the Docker image for the robot application.
     */
    environment?: pulumi.Input<string>;
    /**
     * The name of the simulation application.
     */
    name?: pulumi.Input<string>;
    /**
     * The rendering engine for the simulation application.
     */
    renderingEngine?: pulumi.Input<inputs.robomaker.SimulationApplicationRenderingEngineArgs>;
    /**
     * The robot software suite used by the simulation application.
     */
    robotSoftwareSuite: pulumi.Input<inputs.robomaker.SimulationApplicationRobotSoftwareSuiteArgs>;
    /**
     * The simulation software suite used by the simulation application.
     */
    simulationSoftwareSuite: pulumi.Input<inputs.robomaker.SimulationApplicationSimulationSoftwareSuiteArgs>;
    /**
     * The sources of the simulation application.
     */
    sources?: pulumi.Input<pulumi.Input<inputs.robomaker.SimulationApplicationSourceConfigArgs>[]>;
    tags?: pulumi.Input<inputs.robomaker.SimulationApplicationTagsArgs>;
}
