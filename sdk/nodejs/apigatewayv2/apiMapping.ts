// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * The ``AWS::ApiGatewayV2::ApiMapping`` resource contains an API mapping. An API mapping relates a path of your custom domain name to a stage of your API. A custom domain name can have multiple API mappings, but the paths can't overlap. A custom domain can map only to APIs of the same protocol type. For more information, see [CreateApiMapping](https://docs.aws.amazon.com/apigatewayv2/latest/api-reference/domainnames-domainname-apimappings.html#CreateApiMapping) in the *Amazon API Gateway V2 API Reference*.
 */
export class ApiMapping extends pulumi.CustomResource {
    /**
     * Get an existing ApiMapping resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ApiMapping {
        return new ApiMapping(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:apigatewayv2:ApiMapping';

    /**
     * Returns true if the given object is an instance of ApiMapping.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApiMapping {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApiMapping.__pulumiType;
    }

    /**
     * The identifier of the API.
     */
    public readonly apiId!: pulumi.Output<string>;
    public /*out*/ readonly apiMappingId!: pulumi.Output<string>;
    /**
     * The API mapping key.
     */
    public readonly apiMappingKey!: pulumi.Output<string | undefined>;
    /**
     * The domain name.
     */
    public readonly domainName!: pulumi.Output<string>;
    /**
     * The API stage.
     */
    public readonly stage!: pulumi.Output<string>;

    /**
     * Create a ApiMapping resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApiMappingArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.apiId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'apiId'");
            }
            if ((!args || args.domainName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domainName'");
            }
            if ((!args || args.stage === undefined) && !opts.urn) {
                throw new Error("Missing required property 'stage'");
            }
            resourceInputs["apiId"] = args ? args.apiId : undefined;
            resourceInputs["apiMappingKey"] = args ? args.apiMappingKey : undefined;
            resourceInputs["domainName"] = args ? args.domainName : undefined;
            resourceInputs["stage"] = args ? args.stage : undefined;
            resourceInputs["apiMappingId"] = undefined /*out*/;
        } else {
            resourceInputs["apiId"] = undefined /*out*/;
            resourceInputs["apiMappingId"] = undefined /*out*/;
            resourceInputs["apiMappingKey"] = undefined /*out*/;
            resourceInputs["domainName"] = undefined /*out*/;
            resourceInputs["stage"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["domainName"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(ApiMapping.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ApiMapping resource.
 */
export interface ApiMappingArgs {
    /**
     * The identifier of the API.
     */
    apiId: pulumi.Input<string>;
    /**
     * The API mapping key.
     */
    apiMappingKey?: pulumi.Input<string>;
    /**
     * The domain name.
     */
    domainName: pulumi.Input<string>;
    /**
     * The API stage.
     */
    stage: pulumi.Input<string>;
}
