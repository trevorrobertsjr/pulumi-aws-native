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

__all__ = ['WorkspaceArgs', 'Workspace']

@pulumi.input_type
class WorkspaceArgs:
    def __init__(__self__, *,
                 alert_manager_definition: Optional[pulumi.Input[str]] = None,
                 alias: Optional[pulumi.Input[str]] = None,
                 logging_configuration: Optional[pulumi.Input['WorkspaceLoggingConfigurationArgs']] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['WorkspaceTagArgs']]]] = None):
        """
        The set of arguments for constructing a Workspace resource.
        :param pulumi.Input[str] alert_manager_definition: The AMP Workspace alert manager definition data
        :param pulumi.Input[str] alias: AMP Workspace alias.
        :param pulumi.Input[Sequence[pulumi.Input['WorkspaceTagArgs']]] tags: An array of key-value pairs to apply to this resource.
        """
        if alert_manager_definition is not None:
            pulumi.set(__self__, "alert_manager_definition", alert_manager_definition)
        if alias is not None:
            pulumi.set(__self__, "alias", alias)
        if logging_configuration is not None:
            pulumi.set(__self__, "logging_configuration", logging_configuration)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="alertManagerDefinition")
    def alert_manager_definition(self) -> Optional[pulumi.Input[str]]:
        """
        The AMP Workspace alert manager definition data
        """
        return pulumi.get(self, "alert_manager_definition")

    @alert_manager_definition.setter
    def alert_manager_definition(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "alert_manager_definition", value)

    @property
    @pulumi.getter
    def alias(self) -> Optional[pulumi.Input[str]]:
        """
        AMP Workspace alias.
        """
        return pulumi.get(self, "alias")

    @alias.setter
    def alias(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "alias", value)

    @property
    @pulumi.getter(name="loggingConfiguration")
    def logging_configuration(self) -> Optional[pulumi.Input['WorkspaceLoggingConfigurationArgs']]:
        return pulumi.get(self, "logging_configuration")

    @logging_configuration.setter
    def logging_configuration(self, value: Optional[pulumi.Input['WorkspaceLoggingConfigurationArgs']]):
        pulumi.set(self, "logging_configuration", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['WorkspaceTagArgs']]]]:
        """
        An array of key-value pairs to apply to this resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['WorkspaceTagArgs']]]]):
        pulumi.set(self, "tags", value)


class Workspace(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alert_manager_definition: Optional[pulumi.Input[str]] = None,
                 alias: Optional[pulumi.Input[str]] = None,
                 logging_configuration: Optional[pulumi.Input[pulumi.InputType['WorkspaceLoggingConfigurationArgs']]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WorkspaceTagArgs']]]]] = None,
                 __props__=None):
        """
        Resource Type definition for AWS::APS::Workspace

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alert_manager_definition: The AMP Workspace alert manager definition data
        :param pulumi.Input[str] alias: AMP Workspace alias.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WorkspaceTagArgs']]]] tags: An array of key-value pairs to apply to this resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[WorkspaceArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource Type definition for AWS::APS::Workspace

        :param str resource_name: The name of the resource.
        :param WorkspaceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WorkspaceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alert_manager_definition: Optional[pulumi.Input[str]] = None,
                 alias: Optional[pulumi.Input[str]] = None,
                 logging_configuration: Optional[pulumi.Input[pulumi.InputType['WorkspaceLoggingConfigurationArgs']]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['WorkspaceTagArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = WorkspaceArgs.__new__(WorkspaceArgs)

            __props__.__dict__["alert_manager_definition"] = alert_manager_definition
            __props__.__dict__["alias"] = alias
            __props__.__dict__["logging_configuration"] = logging_configuration
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["prometheus_endpoint"] = None
            __props__.__dict__["workspace_id"] = None
        super(Workspace, __self__).__init__(
            'aws-native:aps:Workspace',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Workspace':
        """
        Get an existing Workspace resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = WorkspaceArgs.__new__(WorkspaceArgs)

        __props__.__dict__["alert_manager_definition"] = None
        __props__.__dict__["alias"] = None
        __props__.__dict__["arn"] = None
        __props__.__dict__["logging_configuration"] = None
        __props__.__dict__["prometheus_endpoint"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["workspace_id"] = None
        return Workspace(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="alertManagerDefinition")
    def alert_manager_definition(self) -> pulumi.Output[Optional[str]]:
        """
        The AMP Workspace alert manager definition data
        """
        return pulumi.get(self, "alert_manager_definition")

    @property
    @pulumi.getter
    def alias(self) -> pulumi.Output[Optional[str]]:
        """
        AMP Workspace alias.
        """
        return pulumi.get(self, "alias")

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        Workspace arn.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="loggingConfiguration")
    def logging_configuration(self) -> pulumi.Output[Optional['outputs.WorkspaceLoggingConfiguration']]:
        return pulumi.get(self, "logging_configuration")

    @property
    @pulumi.getter(name="prometheusEndpoint")
    def prometheus_endpoint(self) -> pulumi.Output[str]:
        """
        AMP Workspace prometheus endpoint
        """
        return pulumi.get(self, "prometheus_endpoint")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.WorkspaceTag']]]:
        """
        An array of key-value pairs to apply to this resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="workspaceId")
    def workspace_id(self) -> pulumi.Output[str]:
        """
        Required to identify a specific APS Workspace.
        """
        return pulumi.get(self, "workspace_id")

