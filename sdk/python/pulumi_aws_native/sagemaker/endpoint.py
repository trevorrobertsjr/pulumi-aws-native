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

__all__ = ['EndpointArgs', 'Endpoint']

@pulumi.input_type
class EndpointArgs:
    def __init__(__self__, *,
                 endpoint_config_name: pulumi.Input[str],
                 deployment_config: Optional[pulumi.Input['EndpointDeploymentConfigArgs']] = None,
                 endpoint_name: Optional[pulumi.Input[str]] = None,
                 exclude_retained_variant_properties: Optional[pulumi.Input[Sequence[pulumi.Input['EndpointVariantPropertyArgs']]]] = None,
                 retain_all_variant_properties: Optional[pulumi.Input[bool]] = None,
                 retain_deployment_config: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['EndpointTagArgs']]]] = None):
        """
        The set of arguments for constructing a Endpoint resource.
        """
        pulumi.set(__self__, "endpoint_config_name", endpoint_config_name)
        if deployment_config is not None:
            pulumi.set(__self__, "deployment_config", deployment_config)
        if endpoint_name is not None:
            pulumi.set(__self__, "endpoint_name", endpoint_name)
        if exclude_retained_variant_properties is not None:
            pulumi.set(__self__, "exclude_retained_variant_properties", exclude_retained_variant_properties)
        if retain_all_variant_properties is not None:
            pulumi.set(__self__, "retain_all_variant_properties", retain_all_variant_properties)
        if retain_deployment_config is not None:
            pulumi.set(__self__, "retain_deployment_config", retain_deployment_config)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="endpointConfigName")
    def endpoint_config_name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "endpoint_config_name")

    @endpoint_config_name.setter
    def endpoint_config_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "endpoint_config_name", value)

    @property
    @pulumi.getter(name="deploymentConfig")
    def deployment_config(self) -> Optional[pulumi.Input['EndpointDeploymentConfigArgs']]:
        return pulumi.get(self, "deployment_config")

    @deployment_config.setter
    def deployment_config(self, value: Optional[pulumi.Input['EndpointDeploymentConfigArgs']]):
        pulumi.set(self, "deployment_config", value)

    @property
    @pulumi.getter(name="endpointName")
    def endpoint_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "endpoint_name")

    @endpoint_name.setter
    def endpoint_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "endpoint_name", value)

    @property
    @pulumi.getter(name="excludeRetainedVariantProperties")
    def exclude_retained_variant_properties(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['EndpointVariantPropertyArgs']]]]:
        return pulumi.get(self, "exclude_retained_variant_properties")

    @exclude_retained_variant_properties.setter
    def exclude_retained_variant_properties(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['EndpointVariantPropertyArgs']]]]):
        pulumi.set(self, "exclude_retained_variant_properties", value)

    @property
    @pulumi.getter(name="retainAllVariantProperties")
    def retain_all_variant_properties(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "retain_all_variant_properties")

    @retain_all_variant_properties.setter
    def retain_all_variant_properties(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "retain_all_variant_properties", value)

    @property
    @pulumi.getter(name="retainDeploymentConfig")
    def retain_deployment_config(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "retain_deployment_config")

    @retain_deployment_config.setter
    def retain_deployment_config(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "retain_deployment_config", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['EndpointTagArgs']]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['EndpointTagArgs']]]]):
        pulumi.set(self, "tags", value)


warnings.warn("""Endpoint is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)


class Endpoint(pulumi.CustomResource):
    warnings.warn("""Endpoint is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 deployment_config: Optional[pulumi.Input[pulumi.InputType['EndpointDeploymentConfigArgs']]] = None,
                 endpoint_config_name: Optional[pulumi.Input[str]] = None,
                 endpoint_name: Optional[pulumi.Input[str]] = None,
                 exclude_retained_variant_properties: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EndpointVariantPropertyArgs']]]]] = None,
                 retain_all_variant_properties: Optional[pulumi.Input[bool]] = None,
                 retain_deployment_config: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EndpointTagArgs']]]]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::SageMaker::Endpoint

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EndpointArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::SageMaker::Endpoint

        :param str resource_name: The name of the resource.
        :param EndpointArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EndpointArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 deployment_config: Optional[pulumi.Input[pulumi.InputType['EndpointDeploymentConfigArgs']]] = None,
                 endpoint_config_name: Optional[pulumi.Input[str]] = None,
                 endpoint_name: Optional[pulumi.Input[str]] = None,
                 exclude_retained_variant_properties: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EndpointVariantPropertyArgs']]]]] = None,
                 retain_all_variant_properties: Optional[pulumi.Input[bool]] = None,
                 retain_deployment_config: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EndpointTagArgs']]]]] = None,
                 __props__=None):
        pulumi.log.warn("""Endpoint is deprecated: Endpoint is not yet supported by AWS Native, so its creation will currently fail. Please use the classic AWS provider, if possible.""")
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EndpointArgs.__new__(EndpointArgs)

            __props__.__dict__["deployment_config"] = deployment_config
            if endpoint_config_name is None and not opts.urn:
                raise TypeError("Missing required property 'endpoint_config_name'")
            __props__.__dict__["endpoint_config_name"] = endpoint_config_name
            __props__.__dict__["endpoint_name"] = endpoint_name
            __props__.__dict__["exclude_retained_variant_properties"] = exclude_retained_variant_properties
            __props__.__dict__["retain_all_variant_properties"] = retain_all_variant_properties
            __props__.__dict__["retain_deployment_config"] = retain_deployment_config
            __props__.__dict__["tags"] = tags
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["endpoint_name"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(Endpoint, __self__).__init__(
            'aws-native:sagemaker:Endpoint',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Endpoint':
        """
        Get an existing Endpoint resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = EndpointArgs.__new__(EndpointArgs)

        __props__.__dict__["deployment_config"] = None
        __props__.__dict__["endpoint_config_name"] = None
        __props__.__dict__["endpoint_name"] = None
        __props__.__dict__["exclude_retained_variant_properties"] = None
        __props__.__dict__["retain_all_variant_properties"] = None
        __props__.__dict__["retain_deployment_config"] = None
        __props__.__dict__["tags"] = None
        return Endpoint(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="deploymentConfig")
    def deployment_config(self) -> pulumi.Output[Optional['outputs.EndpointDeploymentConfig']]:
        return pulumi.get(self, "deployment_config")

    @property
    @pulumi.getter(name="endpointConfigName")
    def endpoint_config_name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "endpoint_config_name")

    @property
    @pulumi.getter(name="endpointName")
    def endpoint_name(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "endpoint_name")

    @property
    @pulumi.getter(name="excludeRetainedVariantProperties")
    def exclude_retained_variant_properties(self) -> pulumi.Output[Optional[Sequence['outputs.EndpointVariantProperty']]]:
        return pulumi.get(self, "exclude_retained_variant_properties")

    @property
    @pulumi.getter(name="retainAllVariantProperties")
    def retain_all_variant_properties(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "retain_all_variant_properties")

    @property
    @pulumi.getter(name="retainDeploymentConfig")
    def retain_deployment_config(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "retain_deployment_config")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.EndpointTag']]]:
        return pulumi.get(self, "tags")

