// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::SageMaker::InferenceComponent
 */
export class InferenceComponent extends pulumi.CustomResource {
    /**
     * Get an existing InferenceComponent resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): InferenceComponent {
        return new InferenceComponent(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:sagemaker:InferenceComponent';

    /**
     * Returns true if the given object is an instance of InferenceComponent.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is InferenceComponent {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === InferenceComponent.__pulumiType;
    }

    public /*out*/ readonly creationTime!: pulumi.Output<string>;
    public readonly endpointArn!: pulumi.Output<string | undefined>;
    public readonly endpointName!: pulumi.Output<string>;
    public /*out*/ readonly failureReason!: pulumi.Output<string>;
    public /*out*/ readonly inferenceComponentArn!: pulumi.Output<string>;
    public readonly inferenceComponentName!: pulumi.Output<string | undefined>;
    public /*out*/ readonly inferenceComponentStatus!: pulumi.Output<enums.sagemaker.InferenceComponentStatus>;
    public /*out*/ readonly lastModifiedTime!: pulumi.Output<string>;
    public readonly runtimeConfig!: pulumi.Output<outputs.sagemaker.InferenceComponentRuntimeConfig>;
    public readonly specification!: pulumi.Output<outputs.sagemaker.InferenceComponentSpecification>;
    public readonly tags!: pulumi.Output<outputs.sagemaker.InferenceComponentTag[] | undefined>;
    public readonly variantName!: pulumi.Output<string>;

    /**
     * Create a InferenceComponent resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InferenceComponentArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.endpointName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'endpointName'");
            }
            if ((!args || args.runtimeConfig === undefined) && !opts.urn) {
                throw new Error("Missing required property 'runtimeConfig'");
            }
            if ((!args || args.specification === undefined) && !opts.urn) {
                throw new Error("Missing required property 'specification'");
            }
            if ((!args || args.variantName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'variantName'");
            }
            resourceInputs["endpointArn"] = args ? args.endpointArn : undefined;
            resourceInputs["endpointName"] = args ? args.endpointName : undefined;
            resourceInputs["inferenceComponentName"] = args ? args.inferenceComponentName : undefined;
            resourceInputs["runtimeConfig"] = args ? args.runtimeConfig : undefined;
            resourceInputs["specification"] = args ? args.specification : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["variantName"] = args ? args.variantName : undefined;
            resourceInputs["creationTime"] = undefined /*out*/;
            resourceInputs["failureReason"] = undefined /*out*/;
            resourceInputs["inferenceComponentArn"] = undefined /*out*/;
            resourceInputs["inferenceComponentStatus"] = undefined /*out*/;
            resourceInputs["lastModifiedTime"] = undefined /*out*/;
        } else {
            resourceInputs["creationTime"] = undefined /*out*/;
            resourceInputs["endpointArn"] = undefined /*out*/;
            resourceInputs["endpointName"] = undefined /*out*/;
            resourceInputs["failureReason"] = undefined /*out*/;
            resourceInputs["inferenceComponentArn"] = undefined /*out*/;
            resourceInputs["inferenceComponentName"] = undefined /*out*/;
            resourceInputs["inferenceComponentStatus"] = undefined /*out*/;
            resourceInputs["lastModifiedTime"] = undefined /*out*/;
            resourceInputs["runtimeConfig"] = undefined /*out*/;
            resourceInputs["specification"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
            resourceInputs["variantName"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(InferenceComponent.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a InferenceComponent resource.
 */
export interface InferenceComponentArgs {
    endpointArn?: pulumi.Input<string>;
    endpointName: pulumi.Input<string>;
    inferenceComponentName?: pulumi.Input<string>;
    runtimeConfig: pulumi.Input<inputs.sagemaker.InferenceComponentRuntimeConfigArgs>;
    specification: pulumi.Input<inputs.sagemaker.InferenceComponentSpecificationArgs>;
    tags?: pulumi.Input<pulumi.Input<inputs.sagemaker.InferenceComponentTagArgs>[]>;
    variantName: pulumi.Input<string>;
}
