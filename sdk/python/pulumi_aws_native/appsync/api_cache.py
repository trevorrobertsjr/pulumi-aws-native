# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ApiCacheArgs', 'ApiCache']

@pulumi.input_type
class ApiCacheArgs:
    def __init__(__self__, *,
                 api_caching_behavior: pulumi.Input[str],
                 api_id: pulumi.Input[str],
                 ttl: pulumi.Input[float],
                 type: pulumi.Input[str],
                 at_rest_encryption_enabled: Optional[pulumi.Input[bool]] = None,
                 transit_encryption_enabled: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a ApiCache resource.
        """
        pulumi.set(__self__, "api_caching_behavior", api_caching_behavior)
        pulumi.set(__self__, "api_id", api_id)
        pulumi.set(__self__, "ttl", ttl)
        pulumi.set(__self__, "type", type)
        if at_rest_encryption_enabled is not None:
            pulumi.set(__self__, "at_rest_encryption_enabled", at_rest_encryption_enabled)
        if transit_encryption_enabled is not None:
            pulumi.set(__self__, "transit_encryption_enabled", transit_encryption_enabled)

    @property
    @pulumi.getter(name="apiCachingBehavior")
    def api_caching_behavior(self) -> pulumi.Input[str]:
        return pulumi.get(self, "api_caching_behavior")

    @api_caching_behavior.setter
    def api_caching_behavior(self, value: pulumi.Input[str]):
        pulumi.set(self, "api_caching_behavior", value)

    @property
    @pulumi.getter(name="apiId")
    def api_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "api_id")

    @api_id.setter
    def api_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "api_id", value)

    @property
    @pulumi.getter
    def ttl(self) -> pulumi.Input[float]:
        return pulumi.get(self, "ttl")

    @ttl.setter
    def ttl(self, value: pulumi.Input[float]):
        pulumi.set(self, "ttl", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="atRestEncryptionEnabled")
    def at_rest_encryption_enabled(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "at_rest_encryption_enabled")

    @at_rest_encryption_enabled.setter
    def at_rest_encryption_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "at_rest_encryption_enabled", value)

    @property
    @pulumi.getter(name="transitEncryptionEnabled")
    def transit_encryption_enabled(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "transit_encryption_enabled")

    @transit_encryption_enabled.setter
    def transit_encryption_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "transit_encryption_enabled", value)


warnings.warn("""ApiCache is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)


class ApiCache(pulumi.CustomResource):
    warnings.warn("""ApiCache is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_caching_behavior: Optional[pulumi.Input[str]] = None,
                 api_id: Optional[pulumi.Input[str]] = None,
                 at_rest_encryption_enabled: Optional[pulumi.Input[bool]] = None,
                 transit_encryption_enabled: Optional[pulumi.Input[bool]] = None,
                 ttl: Optional[pulumi.Input[float]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::AppSync::ApiCache

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ApiCacheArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::AppSync::ApiCache

        :param str resource_name: The name of the resource.
        :param ApiCacheArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ApiCacheArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_caching_behavior: Optional[pulumi.Input[str]] = None,
                 api_id: Optional[pulumi.Input[str]] = None,
                 at_rest_encryption_enabled: Optional[pulumi.Input[bool]] = None,
                 transit_encryption_enabled: Optional[pulumi.Input[bool]] = None,
                 ttl: Optional[pulumi.Input[float]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        pulumi.log.warn("""ApiCache is deprecated: ApiCache is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""")
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ApiCacheArgs.__new__(ApiCacheArgs)

            if api_caching_behavior is None and not opts.urn:
                raise TypeError("Missing required property 'api_caching_behavior'")
            __props__.__dict__["api_caching_behavior"] = api_caching_behavior
            if api_id is None and not opts.urn:
                raise TypeError("Missing required property 'api_id'")
            __props__.__dict__["api_id"] = api_id
            __props__.__dict__["at_rest_encryption_enabled"] = at_rest_encryption_enabled
            __props__.__dict__["transit_encryption_enabled"] = transit_encryption_enabled
            if ttl is None and not opts.urn:
                raise TypeError("Missing required property 'ttl'")
            __props__.__dict__["ttl"] = ttl
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["api_id"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(ApiCache, __self__).__init__(
            'aws-native:appsync:ApiCache',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ApiCache':
        """
        Get an existing ApiCache resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ApiCacheArgs.__new__(ApiCacheArgs)

        __props__.__dict__["api_caching_behavior"] = None
        __props__.__dict__["api_id"] = None
        __props__.__dict__["at_rest_encryption_enabled"] = None
        __props__.__dict__["transit_encryption_enabled"] = None
        __props__.__dict__["ttl"] = None
        __props__.__dict__["type"] = None
        return ApiCache(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="apiCachingBehavior")
    def api_caching_behavior(self) -> pulumi.Output[str]:
        return pulumi.get(self, "api_caching_behavior")

    @property
    @pulumi.getter(name="apiId")
    def api_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "api_id")

    @property
    @pulumi.getter(name="atRestEncryptionEnabled")
    def at_rest_encryption_enabled(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "at_rest_encryption_enabled")

    @property
    @pulumi.getter(name="transitEncryptionEnabled")
    def transit_encryption_enabled(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "transit_encryption_enabled")

    @property
    @pulumi.getter
    def ttl(self) -> pulumi.Output[float]:
        return pulumi.get(self, "ttl")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        return pulumi.get(self, "type")

