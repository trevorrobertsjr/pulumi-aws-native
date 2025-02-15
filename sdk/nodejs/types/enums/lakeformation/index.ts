// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***


export const PrincipalPermissionsPermission = {
    All: "ALL",
    Select: "SELECT",
    Alter: "ALTER",
    Drop: "DROP",
    Delete: "DELETE",
    Insert: "INSERT",
    Describe: "DESCRIBE",
    CreateDatabase: "CREATE_DATABASE",
    CreateTable: "CREATE_TABLE",
    DataLocationAccess: "DATA_LOCATION_ACCESS",
    CreateTag: "CREATE_TAG",
    Associate: "ASSOCIATE",
} as const;

export type PrincipalPermissionsPermission = (typeof PrincipalPermissionsPermission)[keyof typeof PrincipalPermissionsPermission];

export const PrincipalPermissionsResourceType = {
    Database: "DATABASE",
    Table: "TABLE",
} as const;

export type PrincipalPermissionsResourceType = (typeof PrincipalPermissionsResourceType)[keyof typeof PrincipalPermissionsResourceType];
