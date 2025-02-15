// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::MWAA::Environment
 */
export class Environment extends pulumi.CustomResource {
    /**
     * Get an existing Environment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Environment {
        return new Environment(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:mwaa:Environment';

    /**
     * Returns true if the given object is an instance of Environment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Environment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Environment.__pulumiType;
    }

    /**
     * Key/value pairs representing Airflow configuration variables.
     *     Keys are prefixed by their section:
     *
     *     [core]
     *     dags_folder={AIRFLOW_HOME}/dags
     *
     *     Would be represented as
     *
     *     "core.dags_folder": "{AIRFLOW_HOME}/dags"
     */
    public readonly airflowConfigurationOptions!: pulumi.Output<any | undefined>;
    public readonly airflowVersion!: pulumi.Output<string | undefined>;
    public /*out*/ readonly arn!: pulumi.Output<string>;
    public /*out*/ readonly celeryExecutorQueue!: pulumi.Output<string>;
    public readonly dagS3Path!: pulumi.Output<string | undefined>;
    public /*out*/ readonly databaseVpcEndpointService!: pulumi.Output<string>;
    public readonly endpointManagement!: pulumi.Output<enums.mwaa.EnvironmentEndpointManagement | undefined>;
    public readonly environmentClass!: pulumi.Output<string | undefined>;
    public readonly executionRoleArn!: pulumi.Output<string | undefined>;
    public readonly kmsKey!: pulumi.Output<string | undefined>;
    public readonly loggingConfiguration!: pulumi.Output<outputs.mwaa.EnvironmentLoggingConfiguration | undefined>;
    public readonly maxWorkers!: pulumi.Output<number | undefined>;
    public readonly minWorkers!: pulumi.Output<number | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly networkConfiguration!: pulumi.Output<outputs.mwaa.EnvironmentNetworkConfiguration | undefined>;
    public readonly pluginsS3ObjectVersion!: pulumi.Output<string | undefined>;
    public readonly pluginsS3Path!: pulumi.Output<string | undefined>;
    public readonly requirementsS3ObjectVersion!: pulumi.Output<string | undefined>;
    public readonly requirementsS3Path!: pulumi.Output<string | undefined>;
    public readonly schedulers!: pulumi.Output<number | undefined>;
    public readonly sourceBucketArn!: pulumi.Output<string | undefined>;
    public readonly startupScriptS3ObjectVersion!: pulumi.Output<string | undefined>;
    public readonly startupScriptS3Path!: pulumi.Output<string | undefined>;
    /**
     * A map of tags for the environment.
     */
    public readonly tags!: pulumi.Output<any | undefined>;
    public readonly webserverAccessMode!: pulumi.Output<enums.mwaa.EnvironmentWebserverAccessMode | undefined>;
    public /*out*/ readonly webserverUrl!: pulumi.Output<string>;
    public /*out*/ readonly webserverVpcEndpointService!: pulumi.Output<string>;
    public readonly weeklyMaintenanceWindowStart!: pulumi.Output<string | undefined>;

    /**
     * Create a Environment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: EnvironmentArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["airflowConfigurationOptions"] = args ? args.airflowConfigurationOptions : undefined;
            resourceInputs["airflowVersion"] = args ? args.airflowVersion : undefined;
            resourceInputs["dagS3Path"] = args ? args.dagS3Path : undefined;
            resourceInputs["endpointManagement"] = args ? args.endpointManagement : undefined;
            resourceInputs["environmentClass"] = args ? args.environmentClass : undefined;
            resourceInputs["executionRoleArn"] = args ? args.executionRoleArn : undefined;
            resourceInputs["kmsKey"] = args ? args.kmsKey : undefined;
            resourceInputs["loggingConfiguration"] = args ? args.loggingConfiguration : undefined;
            resourceInputs["maxWorkers"] = args ? args.maxWorkers : undefined;
            resourceInputs["minWorkers"] = args ? args.minWorkers : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["networkConfiguration"] = args ? args.networkConfiguration : undefined;
            resourceInputs["pluginsS3ObjectVersion"] = args ? args.pluginsS3ObjectVersion : undefined;
            resourceInputs["pluginsS3Path"] = args ? args.pluginsS3Path : undefined;
            resourceInputs["requirementsS3ObjectVersion"] = args ? args.requirementsS3ObjectVersion : undefined;
            resourceInputs["requirementsS3Path"] = args ? args.requirementsS3Path : undefined;
            resourceInputs["schedulers"] = args ? args.schedulers : undefined;
            resourceInputs["sourceBucketArn"] = args ? args.sourceBucketArn : undefined;
            resourceInputs["startupScriptS3ObjectVersion"] = args ? args.startupScriptS3ObjectVersion : undefined;
            resourceInputs["startupScriptS3Path"] = args ? args.startupScriptS3Path : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["webserverAccessMode"] = args ? args.webserverAccessMode : undefined;
            resourceInputs["weeklyMaintenanceWindowStart"] = args ? args.weeklyMaintenanceWindowStart : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["celeryExecutorQueue"] = undefined /*out*/;
            resourceInputs["databaseVpcEndpointService"] = undefined /*out*/;
            resourceInputs["webserverUrl"] = undefined /*out*/;
            resourceInputs["webserverVpcEndpointService"] = undefined /*out*/;
        } else {
            resourceInputs["airflowConfigurationOptions"] = undefined /*out*/;
            resourceInputs["airflowVersion"] = undefined /*out*/;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["celeryExecutorQueue"] = undefined /*out*/;
            resourceInputs["dagS3Path"] = undefined /*out*/;
            resourceInputs["databaseVpcEndpointService"] = undefined /*out*/;
            resourceInputs["endpointManagement"] = undefined /*out*/;
            resourceInputs["environmentClass"] = undefined /*out*/;
            resourceInputs["executionRoleArn"] = undefined /*out*/;
            resourceInputs["kmsKey"] = undefined /*out*/;
            resourceInputs["loggingConfiguration"] = undefined /*out*/;
            resourceInputs["maxWorkers"] = undefined /*out*/;
            resourceInputs["minWorkers"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["networkConfiguration"] = undefined /*out*/;
            resourceInputs["pluginsS3ObjectVersion"] = undefined /*out*/;
            resourceInputs["pluginsS3Path"] = undefined /*out*/;
            resourceInputs["requirementsS3ObjectVersion"] = undefined /*out*/;
            resourceInputs["requirementsS3Path"] = undefined /*out*/;
            resourceInputs["schedulers"] = undefined /*out*/;
            resourceInputs["sourceBucketArn"] = undefined /*out*/;
            resourceInputs["startupScriptS3ObjectVersion"] = undefined /*out*/;
            resourceInputs["startupScriptS3Path"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["webserverAccessMode"] = undefined /*out*/;
            resourceInputs["webserverUrl"] = undefined /*out*/;
            resourceInputs["webserverVpcEndpointService"] = undefined /*out*/;
            resourceInputs["weeklyMaintenanceWindowStart"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["endpointManagement", "kmsKey", "name", "networkConfiguration.subnetIds[*]"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Environment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Environment resource.
 */
export interface EnvironmentArgs {
    /**
     * Key/value pairs representing Airflow configuration variables.
     *     Keys are prefixed by their section:
     *
     *     [core]
     *     dags_folder={AIRFLOW_HOME}/dags
     *
     *     Would be represented as
     *
     *     "core.dags_folder": "{AIRFLOW_HOME}/dags"
     */
    airflowConfigurationOptions?: any;
    airflowVersion?: pulumi.Input<string>;
    dagS3Path?: pulumi.Input<string>;
    endpointManagement?: pulumi.Input<enums.mwaa.EnvironmentEndpointManagement>;
    environmentClass?: pulumi.Input<string>;
    executionRoleArn?: pulumi.Input<string>;
    kmsKey?: pulumi.Input<string>;
    loggingConfiguration?: pulumi.Input<inputs.mwaa.EnvironmentLoggingConfigurationArgs>;
    maxWorkers?: pulumi.Input<number>;
    minWorkers?: pulumi.Input<number>;
    name?: pulumi.Input<string>;
    networkConfiguration?: pulumi.Input<inputs.mwaa.EnvironmentNetworkConfigurationArgs>;
    pluginsS3ObjectVersion?: pulumi.Input<string>;
    pluginsS3Path?: pulumi.Input<string>;
    requirementsS3ObjectVersion?: pulumi.Input<string>;
    requirementsS3Path?: pulumi.Input<string>;
    schedulers?: pulumi.Input<number>;
    sourceBucketArn?: pulumi.Input<string>;
    startupScriptS3ObjectVersion?: pulumi.Input<string>;
    startupScriptS3Path?: pulumi.Input<string>;
    /**
     * A map of tags for the environment.
     */
    tags?: any;
    webserverAccessMode?: pulumi.Input<enums.mwaa.EnvironmentWebserverAccessMode>;
    weeklyMaintenanceWindowStart?: pulumi.Input<string>;
}
