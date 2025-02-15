// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Represents a monitor, which defines the monitoring boundaries for measurements that Internet Monitor publishes information about for an application
 */
export class Monitor extends pulumi.CustomResource {
    /**
     * Get an existing Monitor resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Monitor {
        return new Monitor(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:internetmonitor:Monitor';

    /**
     * Returns true if the given object is an instance of Monitor.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Monitor {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Monitor.__pulumiType;
    }

    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    public readonly healthEventsConfig!: pulumi.Output<outputs.internetmonitor.MonitorHealthEventsConfig | undefined>;
    public readonly internetMeasurementsLogDelivery!: pulumi.Output<outputs.internetmonitor.MonitorInternetMeasurementsLogDelivery | undefined>;
    public readonly maxCityNetworksToMonitor!: pulumi.Output<number | undefined>;
    public /*out*/ readonly modifiedAt!: pulumi.Output<string>;
    public /*out*/ readonly monitorArn!: pulumi.Output<string>;
    public readonly monitorName!: pulumi.Output<string>;
    public /*out*/ readonly processingStatus!: pulumi.Output<enums.internetmonitor.MonitorProcessingStatusCode>;
    public /*out*/ readonly processingStatusInfo!: pulumi.Output<string>;
    public readonly resources!: pulumi.Output<string[] | undefined>;
    public readonly resourcesToAdd!: pulumi.Output<string[] | undefined>;
    public readonly resourcesToRemove!: pulumi.Output<string[] | undefined>;
    public readonly status!: pulumi.Output<enums.internetmonitor.MonitorConfigState | undefined>;
    public readonly tags!: pulumi.Output<outputs.internetmonitor.MonitorTag[] | undefined>;
    public readonly trafficPercentageToMonitor!: pulumi.Output<number | undefined>;

    /**
     * Create a Monitor resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: MonitorArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["healthEventsConfig"] = args ? args.healthEventsConfig : undefined;
            resourceInputs["internetMeasurementsLogDelivery"] = args ? args.internetMeasurementsLogDelivery : undefined;
            resourceInputs["maxCityNetworksToMonitor"] = args ? args.maxCityNetworksToMonitor : undefined;
            resourceInputs["monitorName"] = args ? args.monitorName : undefined;
            resourceInputs["resources"] = args ? args.resources : undefined;
            resourceInputs["resourcesToAdd"] = args ? args.resourcesToAdd : undefined;
            resourceInputs["resourcesToRemove"] = args ? args.resourcesToRemove : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["trafficPercentageToMonitor"] = args ? args.trafficPercentageToMonitor : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["modifiedAt"] = undefined /*out*/;
            resourceInputs["monitorArn"] = undefined /*out*/;
            resourceInputs["processingStatus"] = undefined /*out*/;
            resourceInputs["processingStatusInfo"] = undefined /*out*/;
        } else {
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["healthEventsConfig"] = undefined /*out*/;
            resourceInputs["internetMeasurementsLogDelivery"] = undefined /*out*/;
            resourceInputs["maxCityNetworksToMonitor"] = undefined /*out*/;
            resourceInputs["modifiedAt"] = undefined /*out*/;
            resourceInputs["monitorArn"] = undefined /*out*/;
            resourceInputs["monitorName"] = undefined /*out*/;
            resourceInputs["processingStatus"] = undefined /*out*/;
            resourceInputs["processingStatusInfo"] = undefined /*out*/;
            resourceInputs["resources"] = undefined /*out*/;
            resourceInputs["resourcesToAdd"] = undefined /*out*/;
            resourceInputs["resourcesToRemove"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["trafficPercentageToMonitor"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["monitorName"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Monitor.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Monitor resource.
 */
export interface MonitorArgs {
    healthEventsConfig?: pulumi.Input<inputs.internetmonitor.MonitorHealthEventsConfigArgs>;
    internetMeasurementsLogDelivery?: pulumi.Input<inputs.internetmonitor.MonitorInternetMeasurementsLogDeliveryArgs>;
    maxCityNetworksToMonitor?: pulumi.Input<number>;
    monitorName?: pulumi.Input<string>;
    resources?: pulumi.Input<pulumi.Input<string>[]>;
    resourcesToAdd?: pulumi.Input<pulumi.Input<string>[]>;
    resourcesToRemove?: pulumi.Input<pulumi.Input<string>[]>;
    status?: pulumi.Input<enums.internetmonitor.MonitorConfigState>;
    tags?: pulumi.Input<pulumi.Input<inputs.internetmonitor.MonitorTagArgs>[]>;
    trafficPercentageToMonitor?: pulumi.Input<number>;
}
