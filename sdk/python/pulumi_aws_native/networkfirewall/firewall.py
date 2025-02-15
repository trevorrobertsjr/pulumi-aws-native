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

__all__ = ['FirewallArgs', 'Firewall']

@pulumi.input_type
class FirewallArgs:
    def __init__(__self__, *,
                 firewall_policy_arn: pulumi.Input[str],
                 subnet_mappings: pulumi.Input[Sequence[pulumi.Input['FirewallSubnetMappingArgs']]],
                 vpc_id: pulumi.Input[str],
                 delete_protection: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 firewall_name: Optional[pulumi.Input[str]] = None,
                 firewall_policy_change_protection: Optional[pulumi.Input[bool]] = None,
                 subnet_change_protection: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallTagArgs']]]] = None):
        """
        The set of arguments for constructing a Firewall resource.
        """
        pulumi.set(__self__, "firewall_policy_arn", firewall_policy_arn)
        pulumi.set(__self__, "subnet_mappings", subnet_mappings)
        pulumi.set(__self__, "vpc_id", vpc_id)
        if delete_protection is not None:
            pulumi.set(__self__, "delete_protection", delete_protection)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if firewall_name is not None:
            pulumi.set(__self__, "firewall_name", firewall_name)
        if firewall_policy_change_protection is not None:
            pulumi.set(__self__, "firewall_policy_change_protection", firewall_policy_change_protection)
        if subnet_change_protection is not None:
            pulumi.set(__self__, "subnet_change_protection", subnet_change_protection)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="firewallPolicyArn")
    def firewall_policy_arn(self) -> pulumi.Input[str]:
        return pulumi.get(self, "firewall_policy_arn")

    @firewall_policy_arn.setter
    def firewall_policy_arn(self, value: pulumi.Input[str]):
        pulumi.set(self, "firewall_policy_arn", value)

    @property
    @pulumi.getter(name="subnetMappings")
    def subnet_mappings(self) -> pulumi.Input[Sequence[pulumi.Input['FirewallSubnetMappingArgs']]]:
        return pulumi.get(self, "subnet_mappings")

    @subnet_mappings.setter
    def subnet_mappings(self, value: pulumi.Input[Sequence[pulumi.Input['FirewallSubnetMappingArgs']]]):
        pulumi.set(self, "subnet_mappings", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vpc_id", value)

    @property
    @pulumi.getter(name="deleteProtection")
    def delete_protection(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "delete_protection")

    @delete_protection.setter
    def delete_protection(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "delete_protection", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="firewallName")
    def firewall_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "firewall_name")

    @firewall_name.setter
    def firewall_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "firewall_name", value)

    @property
    @pulumi.getter(name="firewallPolicyChangeProtection")
    def firewall_policy_change_protection(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "firewall_policy_change_protection")

    @firewall_policy_change_protection.setter
    def firewall_policy_change_protection(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "firewall_policy_change_protection", value)

    @property
    @pulumi.getter(name="subnetChangeProtection")
    def subnet_change_protection(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "subnet_change_protection")

    @subnet_change_protection.setter
    def subnet_change_protection(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "subnet_change_protection", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FirewallTagArgs']]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FirewallTagArgs']]]]):
        pulumi.set(self, "tags", value)


class Firewall(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 delete_protection: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 firewall_name: Optional[pulumi.Input[str]] = None,
                 firewall_policy_arn: Optional[pulumi.Input[str]] = None,
                 firewall_policy_change_protection: Optional[pulumi.Input[bool]] = None,
                 subnet_change_protection: Optional[pulumi.Input[bool]] = None,
                 subnet_mappings: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallSubnetMappingArgs']]]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallTagArgs']]]]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Resource type definition for AWS::NetworkFirewall::Firewall

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FirewallArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource type definition for AWS::NetworkFirewall::Firewall

        :param str resource_name: The name of the resource.
        :param FirewallArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FirewallArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 delete_protection: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 firewall_name: Optional[pulumi.Input[str]] = None,
                 firewall_policy_arn: Optional[pulumi.Input[str]] = None,
                 firewall_policy_change_protection: Optional[pulumi.Input[bool]] = None,
                 subnet_change_protection: Optional[pulumi.Input[bool]] = None,
                 subnet_mappings: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallSubnetMappingArgs']]]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['FirewallTagArgs']]]]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FirewallArgs.__new__(FirewallArgs)

            __props__.__dict__["delete_protection"] = delete_protection
            __props__.__dict__["description"] = description
            __props__.__dict__["firewall_name"] = firewall_name
            if firewall_policy_arn is None and not opts.urn:
                raise TypeError("Missing required property 'firewall_policy_arn'")
            __props__.__dict__["firewall_policy_arn"] = firewall_policy_arn
            __props__.__dict__["firewall_policy_change_protection"] = firewall_policy_change_protection
            __props__.__dict__["subnet_change_protection"] = subnet_change_protection
            if subnet_mappings is None and not opts.urn:
                raise TypeError("Missing required property 'subnet_mappings'")
            __props__.__dict__["subnet_mappings"] = subnet_mappings
            __props__.__dict__["tags"] = tags
            if vpc_id is None and not opts.urn:
                raise TypeError("Missing required property 'vpc_id'")
            __props__.__dict__["vpc_id"] = vpc_id
            __props__.__dict__["endpoint_ids"] = None
            __props__.__dict__["firewall_arn"] = None
            __props__.__dict__["firewall_id"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["firewall_name", "vpc_id"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(Firewall, __self__).__init__(
            'aws-native:networkfirewall:Firewall',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Firewall':
        """
        Get an existing Firewall resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = FirewallArgs.__new__(FirewallArgs)

        __props__.__dict__["delete_protection"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["endpoint_ids"] = None
        __props__.__dict__["firewall_arn"] = None
        __props__.__dict__["firewall_id"] = None
        __props__.__dict__["firewall_name"] = None
        __props__.__dict__["firewall_policy_arn"] = None
        __props__.__dict__["firewall_policy_change_protection"] = None
        __props__.__dict__["subnet_change_protection"] = None
        __props__.__dict__["subnet_mappings"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["vpc_id"] = None
        return Firewall(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="deleteProtection")
    def delete_protection(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "delete_protection")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="endpointIds")
    def endpoint_ids(self) -> pulumi.Output[Sequence[str]]:
        return pulumi.get(self, "endpoint_ids")

    @property
    @pulumi.getter(name="firewallArn")
    def firewall_arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "firewall_arn")

    @property
    @pulumi.getter(name="firewallId")
    def firewall_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "firewall_id")

    @property
    @pulumi.getter(name="firewallName")
    def firewall_name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "firewall_name")

    @property
    @pulumi.getter(name="firewallPolicyArn")
    def firewall_policy_arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "firewall_policy_arn")

    @property
    @pulumi.getter(name="firewallPolicyChangeProtection")
    def firewall_policy_change_protection(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "firewall_policy_change_protection")

    @property
    @pulumi.getter(name="subnetChangeProtection")
    def subnet_change_protection(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "subnet_change_protection")

    @property
    @pulumi.getter(name="subnetMappings")
    def subnet_mappings(self) -> pulumi.Output[Sequence['outputs.FirewallSubnetMapping']]:
        return pulumi.get(self, "subnet_mappings")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.FirewallTag']]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "vpc_id")

