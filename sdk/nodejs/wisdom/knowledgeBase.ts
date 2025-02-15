// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Definition of AWS::Wisdom::KnowledgeBase Resource Type
 */
export class KnowledgeBase extends pulumi.CustomResource {
    /**
     * Get an existing KnowledgeBase resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): KnowledgeBase {
        return new KnowledgeBase(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:wisdom:KnowledgeBase';

    /**
     * Returns true if the given object is an instance of KnowledgeBase.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is KnowledgeBase {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === KnowledgeBase.__pulumiType;
    }

    public readonly description!: pulumi.Output<string | undefined>;
    public /*out*/ readonly knowledgeBaseArn!: pulumi.Output<string>;
    public /*out*/ readonly knowledgeBaseId!: pulumi.Output<string>;
    public readonly knowledgeBaseType!: pulumi.Output<enums.wisdom.KnowledgeBaseType>;
    public readonly name!: pulumi.Output<string>;
    public readonly renderingConfiguration!: pulumi.Output<outputs.wisdom.KnowledgeBaseRenderingConfiguration | undefined>;
    public readonly serverSideEncryptionConfiguration!: pulumi.Output<outputs.wisdom.KnowledgeBaseServerSideEncryptionConfiguration | undefined>;
    public readonly sourceConfiguration!: pulumi.Output<outputs.wisdom.KnowledgeBaseSourceConfiguration | undefined>;
    public readonly tags!: pulumi.Output<outputs.wisdom.KnowledgeBaseTag[] | undefined>;

    /**
     * Create a KnowledgeBase resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: KnowledgeBaseArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.knowledgeBaseType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'knowledgeBaseType'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["knowledgeBaseType"] = args ? args.knowledgeBaseType : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["renderingConfiguration"] = args ? args.renderingConfiguration : undefined;
            resourceInputs["serverSideEncryptionConfiguration"] = args ? args.serverSideEncryptionConfiguration : undefined;
            resourceInputs["sourceConfiguration"] = args ? args.sourceConfiguration : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["knowledgeBaseArn"] = undefined /*out*/;
            resourceInputs["knowledgeBaseId"] = undefined /*out*/;
        } else {
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["knowledgeBaseArn"] = undefined /*out*/;
            resourceInputs["knowledgeBaseId"] = undefined /*out*/;
            resourceInputs["knowledgeBaseType"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["renderingConfiguration"] = undefined /*out*/;
            resourceInputs["serverSideEncryptionConfiguration"] = undefined /*out*/;
            resourceInputs["sourceConfiguration"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["description", "knowledgeBaseType", "name", "serverSideEncryptionConfiguration", "sourceConfiguration", "tags[*]"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(KnowledgeBase.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a KnowledgeBase resource.
 */
export interface KnowledgeBaseArgs {
    description?: pulumi.Input<string>;
    knowledgeBaseType: pulumi.Input<enums.wisdom.KnowledgeBaseType>;
    name?: pulumi.Input<string>;
    renderingConfiguration?: pulumi.Input<inputs.wisdom.KnowledgeBaseRenderingConfigurationArgs>;
    serverSideEncryptionConfiguration?: pulumi.Input<inputs.wisdom.KnowledgeBaseServerSideEncryptionConfigurationArgs>;
    sourceConfiguration?: pulumi.Input<inputs.wisdom.KnowledgeBaseSourceConfigurationArgs>;
    tags?: pulumi.Input<pulumi.Input<inputs.wisdom.KnowledgeBaseTagArgs>[]>;
}
