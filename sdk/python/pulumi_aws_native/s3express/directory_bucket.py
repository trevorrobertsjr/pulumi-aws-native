# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from ._enums import *

__all__ = ['DirectoryBucketArgs', 'DirectoryBucket']

@pulumi.input_type
class DirectoryBucketArgs:
    def __init__(__self__, *,
                 data_redundancy: pulumi.Input['DirectoryBucketDataRedundancy'],
                 location_name: pulumi.Input[str],
                 bucket_name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a DirectoryBucket resource.
        :param pulumi.Input['DirectoryBucketDataRedundancy'] data_redundancy: Specifies the number of Availability Zone that's used for redundancy for the bucket.
        :param pulumi.Input[str] location_name: Specifies the AZ ID of the Availability Zone where the directory bucket will be created. An example AZ ID value is 'use1-az5'.
        :param pulumi.Input[str] bucket_name: Specifies a name for the bucket. The bucket name must contain only lowercase letters, numbers, and hyphens (-). A directory bucket name must be unique in the chosen Availability Zone. The bucket name must also follow the format 'bucket_base_name--az_id--x-s3' (for example, 'DOC-EXAMPLE-BUCKET--usw2-az1--x-s3'). If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the bucket name.
        """
        pulumi.set(__self__, "data_redundancy", data_redundancy)
        pulumi.set(__self__, "location_name", location_name)
        if bucket_name is not None:
            pulumi.set(__self__, "bucket_name", bucket_name)

    @property
    @pulumi.getter(name="dataRedundancy")
    def data_redundancy(self) -> pulumi.Input['DirectoryBucketDataRedundancy']:
        """
        Specifies the number of Availability Zone that's used for redundancy for the bucket.
        """
        return pulumi.get(self, "data_redundancy")

    @data_redundancy.setter
    def data_redundancy(self, value: pulumi.Input['DirectoryBucketDataRedundancy']):
        pulumi.set(self, "data_redundancy", value)

    @property
    @pulumi.getter(name="locationName")
    def location_name(self) -> pulumi.Input[str]:
        """
        Specifies the AZ ID of the Availability Zone where the directory bucket will be created. An example AZ ID value is 'use1-az5'.
        """
        return pulumi.get(self, "location_name")

    @location_name.setter
    def location_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "location_name", value)

    @property
    @pulumi.getter(name="bucketName")
    def bucket_name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies a name for the bucket. The bucket name must contain only lowercase letters, numbers, and hyphens (-). A directory bucket name must be unique in the chosen Availability Zone. The bucket name must also follow the format 'bucket_base_name--az_id--x-s3' (for example, 'DOC-EXAMPLE-BUCKET--usw2-az1--x-s3'). If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the bucket name.
        """
        return pulumi.get(self, "bucket_name")

    @bucket_name.setter
    def bucket_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bucket_name", value)


class DirectoryBucket(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bucket_name: Optional[pulumi.Input[str]] = None,
                 data_redundancy: Optional[pulumi.Input['DirectoryBucketDataRedundancy']] = None,
                 location_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::S3Express::DirectoryBucket.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bucket_name: Specifies a name for the bucket. The bucket name must contain only lowercase letters, numbers, and hyphens (-). A directory bucket name must be unique in the chosen Availability Zone. The bucket name must also follow the format 'bucket_base_name--az_id--x-s3' (for example, 'DOC-EXAMPLE-BUCKET--usw2-az1--x-s3'). If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the bucket name.
        :param pulumi.Input['DirectoryBucketDataRedundancy'] data_redundancy: Specifies the number of Availability Zone that's used for redundancy for the bucket.
        :param pulumi.Input[str] location_name: Specifies the AZ ID of the Availability Zone where the directory bucket will be created. An example AZ ID value is 'use1-az5'.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DirectoryBucketArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::S3Express::DirectoryBucket.

        :param str resource_name: The name of the resource.
        :param DirectoryBucketArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DirectoryBucketArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bucket_name: Optional[pulumi.Input[str]] = None,
                 data_redundancy: Optional[pulumi.Input['DirectoryBucketDataRedundancy']] = None,
                 location_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DirectoryBucketArgs.__new__(DirectoryBucketArgs)

            __props__.__dict__["bucket_name"] = bucket_name
            if data_redundancy is None and not opts.urn:
                raise TypeError("Missing required property 'data_redundancy'")
            __props__.__dict__["data_redundancy"] = data_redundancy
            if location_name is None and not opts.urn:
                raise TypeError("Missing required property 'location_name'")
            __props__.__dict__["location_name"] = location_name
            __props__.__dict__["arn"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["bucket_name", "data_redundancy", "location_name"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(DirectoryBucket, __self__).__init__(
            'aws-native:s3express:DirectoryBucket',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'DirectoryBucket':
        """
        Get an existing DirectoryBucket resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = DirectoryBucketArgs.__new__(DirectoryBucketArgs)

        __props__.__dict__["arn"] = None
        __props__.__dict__["bucket_name"] = None
        __props__.__dict__["data_redundancy"] = None
        __props__.__dict__["location_name"] = None
        return DirectoryBucket(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        Returns the Amazon Resource Name (ARN) of the specified bucket.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="bucketName")
    def bucket_name(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies a name for the bucket. The bucket name must contain only lowercase letters, numbers, and hyphens (-). A directory bucket name must be unique in the chosen Availability Zone. The bucket name must also follow the format 'bucket_base_name--az_id--x-s3' (for example, 'DOC-EXAMPLE-BUCKET--usw2-az1--x-s3'). If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the bucket name.
        """
        return pulumi.get(self, "bucket_name")

    @property
    @pulumi.getter(name="dataRedundancy")
    def data_redundancy(self) -> pulumi.Output['DirectoryBucketDataRedundancy']:
        """
        Specifies the number of Availability Zone that's used for redundancy for the bucket.
        """
        return pulumi.get(self, "data_redundancy")

    @property
    @pulumi.getter(name="locationName")
    def location_name(self) -> pulumi.Output[str]:
        """
        Specifies the AZ ID of the Availability Zone where the directory bucket will be created. An example AZ ID value is 'use1-az5'.
        """
        return pulumi.get(self, "location_name")

