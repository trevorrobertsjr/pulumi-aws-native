// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of AWS::VerifiedPermissions::IdentitySource Resource Type
 */
export class IdentitySource extends pulumi.CustomResource {
    /**
     * Get an existing IdentitySource resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): IdentitySource {
        return new IdentitySource(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:verifiedpermissions:IdentitySource';

    /**
     * Returns true if the given object is an instance of IdentitySource.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IdentitySource {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IdentitySource.__pulumiType;
    }

    public readonly configuration!: pulumi.Output<outputs.verifiedpermissions.IdentitySourceConfiguration>;
    public /*out*/ readonly details!: pulumi.Output<outputs.verifiedpermissions.IdentitySourceDetails>;
    public /*out*/ readonly identitySourceId!: pulumi.Output<string>;
    public readonly policyStoreId!: pulumi.Output<string | undefined>;
    public readonly principalEntityType!: pulumi.Output<string | undefined>;

    /**
     * Create a IdentitySource resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IdentitySourceArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.configuration === undefined) && !opts.urn) {
                throw new Error("Missing required property 'configuration'");
            }
            resourceInputs["configuration"] = args ? args.configuration : undefined;
            resourceInputs["policyStoreId"] = args ? args.policyStoreId : undefined;
            resourceInputs["principalEntityType"] = args ? args.principalEntityType : undefined;
            resourceInputs["details"] = undefined /*out*/;
            resourceInputs["identitySourceId"] = undefined /*out*/;
        } else {
            resourceInputs["configuration"] = undefined /*out*/;
            resourceInputs["details"] = undefined /*out*/;
            resourceInputs["identitySourceId"] = undefined /*out*/;
            resourceInputs["policyStoreId"] = undefined /*out*/;
            resourceInputs["principalEntityType"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["policyStoreId"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(IdentitySource.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a IdentitySource resource.
 */
export interface IdentitySourceArgs {
    configuration: pulumi.Input<inputs.verifiedpermissions.IdentitySourceConfigurationArgs>;
    policyStoreId?: pulumi.Input<string>;
    principalEntityType?: pulumi.Input<string>;
}
