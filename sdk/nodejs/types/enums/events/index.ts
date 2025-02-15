// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const ApiDestinationHttpMethod = {
    Get: "GET",
    Head: "HEAD",
    Post: "POST",
    Options: "OPTIONS",
    Put: "PUT",
    Delete: "DELETE",
    Patch: "PATCH",
} as const;

export type ApiDestinationHttpMethod = (typeof ApiDestinationHttpMethod)[keyof typeof ApiDestinationHttpMethod];

export const ConnectionAuthorizationType = {
    ApiKey: "API_KEY",
    Basic: "BASIC",
    OauthClientCredentials: "OAUTH_CLIENT_CREDENTIALS",
} as const;

export type ConnectionAuthorizationType = (typeof ConnectionAuthorizationType)[keyof typeof ConnectionAuthorizationType];

export const ConnectionOAuthParametersHttpMethod = {
    Get: "GET",
    Post: "POST",
    Put: "PUT",
} as const;

export type ConnectionOAuthParametersHttpMethod = (typeof ConnectionOAuthParametersHttpMethod)[keyof typeof ConnectionOAuthParametersHttpMethod];

export const EndpointReplicationState = {
    Enabled: "ENABLED",
    Disabled: "DISABLED",
} as const;

export type EndpointReplicationState = (typeof EndpointReplicationState)[keyof typeof EndpointReplicationState];

export const EndpointState = {
    Active: "ACTIVE",
    Creating: "CREATING",
    Updating: "UPDATING",
    Deleting: "DELETING",
    CreateFailed: "CREATE_FAILED",
    UpdateFailed: "UPDATE_FAILED",
} as const;

export type EndpointState = (typeof EndpointState)[keyof typeof EndpointState];

export const RuleState = {
    Disabled: "DISABLED",
    Enabled: "ENABLED",
    EnabledWithAllCloudtrailManagementEvents: "ENABLED_WITH_ALL_CLOUDTRAIL_MANAGEMENT_EVENTS",
} as const;

/**
 * The state of the rule.
 */
export type RuleState = (typeof RuleState)[keyof typeof RuleState];
