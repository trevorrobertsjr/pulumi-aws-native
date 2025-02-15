// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::DocDB::DBCluster
 *
 * @deprecated DbCluster is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.
 */
export class DbCluster extends pulumi.CustomResource {
    /**
     * Get an existing DbCluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DbCluster {
        pulumi.log.warn("DbCluster is deprecated: DbCluster is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        return new DbCluster(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:docdb:DbCluster';

    /**
     * Returns true if the given object is an instance of DbCluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DbCluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DbCluster.__pulumiType;
    }

    public readonly availabilityZones!: pulumi.Output<string[] | undefined>;
    public readonly backupRetentionPeriod!: pulumi.Output<number | undefined>;
    public /*out*/ readonly clusterResourceId!: pulumi.Output<string>;
    public readonly copyTagsToSnapshot!: pulumi.Output<boolean | undefined>;
    public readonly dbClusterIdentifier!: pulumi.Output<string | undefined>;
    public readonly dbClusterParameterGroupName!: pulumi.Output<string | undefined>;
    public readonly dbSubnetGroupName!: pulumi.Output<string | undefined>;
    public readonly deletionProtection!: pulumi.Output<boolean | undefined>;
    public readonly enableCloudwatchLogsExports!: pulumi.Output<string[] | undefined>;
    public /*out*/ readonly endpoint!: pulumi.Output<string>;
    public readonly engineVersion!: pulumi.Output<string | undefined>;
    public readonly kmsKeyId!: pulumi.Output<string | undefined>;
    public readonly masterUserPassword!: pulumi.Output<string | undefined>;
    public readonly masterUsername!: pulumi.Output<string | undefined>;
    public readonly port!: pulumi.Output<number | undefined>;
    public readonly preferredBackupWindow!: pulumi.Output<string | undefined>;
    public readonly preferredMaintenanceWindow!: pulumi.Output<string | undefined>;
    public /*out*/ readonly readEndpoint!: pulumi.Output<string>;
    public readonly restoreToTime!: pulumi.Output<string | undefined>;
    public readonly restoreType!: pulumi.Output<string | undefined>;
    public readonly snapshotIdentifier!: pulumi.Output<string | undefined>;
    public readonly sourceDbClusterIdentifier!: pulumi.Output<string | undefined>;
    public readonly storageEncrypted!: pulumi.Output<boolean | undefined>;
    public readonly tags!: pulumi.Output<outputs.docdb.DbClusterTag[] | undefined>;
    public readonly useLatestRestorableTime!: pulumi.Output<boolean | undefined>;
    public readonly vpcSecurityGroupIds!: pulumi.Output<string[] | undefined>;

    /**
     * Create a DbCluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    /** @deprecated DbCluster is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible. */
    constructor(name: string, args?: DbClusterArgs, opts?: pulumi.CustomResourceOptions) {
        pulumi.log.warn("DbCluster is deprecated: DbCluster is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.")
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["availabilityZones"] = args ? args.availabilityZones : undefined;
            resourceInputs["backupRetentionPeriod"] = args ? args.backupRetentionPeriod : undefined;
            resourceInputs["copyTagsToSnapshot"] = args ? args.copyTagsToSnapshot : undefined;
            resourceInputs["dbClusterIdentifier"] = args ? args.dbClusterIdentifier : undefined;
            resourceInputs["dbClusterParameterGroupName"] = args ? args.dbClusterParameterGroupName : undefined;
            resourceInputs["dbSubnetGroupName"] = args ? args.dbSubnetGroupName : undefined;
            resourceInputs["deletionProtection"] = args ? args.deletionProtection : undefined;
            resourceInputs["enableCloudwatchLogsExports"] = args ? args.enableCloudwatchLogsExports : undefined;
            resourceInputs["engineVersion"] = args ? args.engineVersion : undefined;
            resourceInputs["kmsKeyId"] = args ? args.kmsKeyId : undefined;
            resourceInputs["masterUserPassword"] = args ? args.masterUserPassword : undefined;
            resourceInputs["masterUsername"] = args ? args.masterUsername : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["preferredBackupWindow"] = args ? args.preferredBackupWindow : undefined;
            resourceInputs["preferredMaintenanceWindow"] = args ? args.preferredMaintenanceWindow : undefined;
            resourceInputs["restoreToTime"] = args ? args.restoreToTime : undefined;
            resourceInputs["restoreType"] = args ? args.restoreType : undefined;
            resourceInputs["snapshotIdentifier"] = args ? args.snapshotIdentifier : undefined;
            resourceInputs["sourceDbClusterIdentifier"] = args ? args.sourceDbClusterIdentifier : undefined;
            resourceInputs["storageEncrypted"] = args ? args.storageEncrypted : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["useLatestRestorableTime"] = args ? args.useLatestRestorableTime : undefined;
            resourceInputs["vpcSecurityGroupIds"] = args ? args.vpcSecurityGroupIds : undefined;
            resourceInputs["clusterResourceId"] = undefined /*out*/;
            resourceInputs["endpoint"] = undefined /*out*/;
            resourceInputs["readEndpoint"] = undefined /*out*/;
        } else {
            resourceInputs["availabilityZones"] = undefined /*out*/;
            resourceInputs["backupRetentionPeriod"] = undefined /*out*/;
            resourceInputs["clusterResourceId"] = undefined /*out*/;
            resourceInputs["copyTagsToSnapshot"] = undefined /*out*/;
            resourceInputs["dbClusterIdentifier"] = undefined /*out*/;
            resourceInputs["dbClusterParameterGroupName"] = undefined /*out*/;
            resourceInputs["dbSubnetGroupName"] = undefined /*out*/;
            resourceInputs["deletionProtection"] = undefined /*out*/;
            resourceInputs["enableCloudwatchLogsExports"] = undefined /*out*/;
            resourceInputs["endpoint"] = undefined /*out*/;
            resourceInputs["engineVersion"] = undefined /*out*/;
            resourceInputs["kmsKeyId"] = undefined /*out*/;
            resourceInputs["masterUserPassword"] = undefined /*out*/;
            resourceInputs["masterUsername"] = undefined /*out*/;
            resourceInputs["port"] = undefined /*out*/;
            resourceInputs["preferredBackupWindow"] = undefined /*out*/;
            resourceInputs["preferredMaintenanceWindow"] = undefined /*out*/;
            resourceInputs["readEndpoint"] = undefined /*out*/;
            resourceInputs["restoreToTime"] = undefined /*out*/;
            resourceInputs["restoreType"] = undefined /*out*/;
            resourceInputs["snapshotIdentifier"] = undefined /*out*/;
            resourceInputs["sourceDbClusterIdentifier"] = undefined /*out*/;
            resourceInputs["storageEncrypted"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["useLatestRestorableTime"] = undefined /*out*/;
            resourceInputs["vpcSecurityGroupIds"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["availabilityZones[*]", "dbClusterIdentifier", "dbSubnetGroupName", "engineVersion", "kmsKeyId", "masterUsername", "snapshotIdentifier", "sourceDbClusterIdentifier", "storageEncrypted"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(DbCluster.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a DbCluster resource.
 */
export interface DbClusterArgs {
    availabilityZones?: pulumi.Input<pulumi.Input<string>[]>;
    backupRetentionPeriod?: pulumi.Input<number>;
    copyTagsToSnapshot?: pulumi.Input<boolean>;
    dbClusterIdentifier?: pulumi.Input<string>;
    dbClusterParameterGroupName?: pulumi.Input<string>;
    dbSubnetGroupName?: pulumi.Input<string>;
    deletionProtection?: pulumi.Input<boolean>;
    enableCloudwatchLogsExports?: pulumi.Input<pulumi.Input<string>[]>;
    engineVersion?: pulumi.Input<string>;
    kmsKeyId?: pulumi.Input<string>;
    masterUserPassword?: pulumi.Input<string>;
    masterUsername?: pulumi.Input<string>;
    port?: pulumi.Input<number>;
    preferredBackupWindow?: pulumi.Input<string>;
    preferredMaintenanceWindow?: pulumi.Input<string>;
    restoreToTime?: pulumi.Input<string>;
    restoreType?: pulumi.Input<string>;
    snapshotIdentifier?: pulumi.Input<string>;
    sourceDbClusterIdentifier?: pulumi.Input<string>;
    storageEncrypted?: pulumi.Input<boolean>;
    tags?: pulumi.Input<pulumi.Input<inputs.docdb.DbClusterTagArgs>[]>;
    useLatestRestorableTime?: pulumi.Input<boolean>;
    vpcSecurityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
}
