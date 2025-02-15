# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['FunctionArgs', 'Function']

@pulumi.input_type
class FunctionArgs:
    def __init__(__self__, *,
                 function_code: pulumi.Input[str],
                 function_config: pulumi.Input['FunctionConfigArgs'],
                 auto_publish: Optional[pulumi.Input[bool]] = None,
                 function_metadata: Optional[pulumi.Input['FunctionMetadataArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Function resource.
        """
        pulumi.set(__self__, "function_code", function_code)
        pulumi.set(__self__, "function_config", function_config)
        if auto_publish is not None:
            pulumi.set(__self__, "auto_publish", auto_publish)
        if function_metadata is not None:
            pulumi.set(__self__, "function_metadata", function_metadata)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="functionCode")
    def function_code(self) -> pulumi.Input[str]:
        return pulumi.get(self, "function_code")

    @function_code.setter
    def function_code(self, value: pulumi.Input[str]):
        pulumi.set(self, "function_code", value)

    @property
    @pulumi.getter(name="functionConfig")
    def function_config(self) -> pulumi.Input['FunctionConfigArgs']:
        return pulumi.get(self, "function_config")

    @function_config.setter
    def function_config(self, value: pulumi.Input['FunctionConfigArgs']):
        pulumi.set(self, "function_config", value)

    @property
    @pulumi.getter(name="autoPublish")
    def auto_publish(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "auto_publish")

    @auto_publish.setter
    def auto_publish(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_publish", value)

    @property
    @pulumi.getter(name="functionMetadata")
    def function_metadata(self) -> Optional[pulumi.Input['FunctionMetadataArgs']]:
        return pulumi.get(self, "function_metadata")

    @function_metadata.setter
    def function_metadata(self, value: Optional[pulumi.Input['FunctionMetadataArgs']]):
        pulumi.set(self, "function_metadata", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class Function(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_publish: Optional[pulumi.Input[bool]] = None,
                 function_code: Optional[pulumi.Input[str]] = None,
                 function_config: Optional[pulumi.Input[pulumi.InputType['FunctionConfigArgs']]] = None,
                 function_metadata: Optional[pulumi.Input[pulumi.InputType['FunctionMetadataArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::CloudFront::Function

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FunctionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::CloudFront::Function

        :param str resource_name: The name of the resource.
        :param FunctionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FunctionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_publish: Optional[pulumi.Input[bool]] = None,
                 function_code: Optional[pulumi.Input[str]] = None,
                 function_config: Optional[pulumi.Input[pulumi.InputType['FunctionConfigArgs']]] = None,
                 function_metadata: Optional[pulumi.Input[pulumi.InputType['FunctionMetadataArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FunctionArgs.__new__(FunctionArgs)

            __props__.__dict__["auto_publish"] = auto_publish
            if function_code is None and not opts.urn:
                raise TypeError("Missing required property 'function_code'")
            __props__.__dict__["function_code"] = function_code
            if function_config is None and not opts.urn:
                raise TypeError("Missing required property 'function_config'")
            __props__.__dict__["function_config"] = function_config
            __props__.__dict__["function_metadata"] = function_metadata
            __props__.__dict__["name"] = name
            __props__.__dict__["function_arn"] = None
            __props__.__dict__["stage"] = None
        super(Function, __self__).__init__(
            'aws-native:cloudfront:Function',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Function':
        """
        Get an existing Function resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = FunctionArgs.__new__(FunctionArgs)

        __props__.__dict__["auto_publish"] = None
        __props__.__dict__["function_arn"] = None
        __props__.__dict__["function_code"] = None
        __props__.__dict__["function_config"] = None
        __props__.__dict__["function_metadata"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["stage"] = None
        return Function(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoPublish")
    def auto_publish(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "auto_publish")

    @property
    @pulumi.getter(name="functionArn")
    def function_arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "function_arn")

    @property
    @pulumi.getter(name="functionCode")
    def function_code(self) -> pulumi.Output[str]:
        return pulumi.get(self, "function_code")

    @property
    @pulumi.getter(name="functionConfig")
    def function_config(self) -> pulumi.Output['outputs.FunctionConfig']:
        return pulumi.get(self, "function_config")

    @property
    @pulumi.getter(name="functionMetadata")
    def function_metadata(self) -> pulumi.Output[Optional['outputs.FunctionMetadata']]:
        return pulumi.get(self, "function_metadata")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def stage(self) -> pulumi.Output[str]:
        return pulumi.get(self, "stage")

