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

__all__ = [
    'DevicePoolRuleArgs',
    'DevicePoolTagArgs',
    'InstanceProfileTagArgs',
    'NetworkProfileTagArgs',
    'ProjectTagArgs',
    'ProjectVpcConfigArgs',
    'TestGridProjectTagArgs',
    'TestGridProjectVpcConfigArgs',
    'VpceConfigurationTagArgs',
]

@pulumi.input_type
class DevicePoolRuleArgs:
    def __init__(__self__, *,
                 attribute: Optional[pulumi.Input['DevicePoolRuleAttribute']] = None,
                 operator: Optional[pulumi.Input['DevicePoolRuleOperator']] = None,
                 value: Optional[pulumi.Input[str]] = None):
        """
        Represents a condition for a device pool.
        :param pulumi.Input['DevicePoolRuleAttribute'] attribute: The rule's stringified attribute.
        :param pulumi.Input['DevicePoolRuleOperator'] operator: Specifies how Device Farm compares the rule's attribute to the value.
        :param pulumi.Input[str] value: The rule's value.
        """
        if attribute is not None:
            pulumi.set(__self__, "attribute", attribute)
        if operator is not None:
            pulumi.set(__self__, "operator", operator)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def attribute(self) -> Optional[pulumi.Input['DevicePoolRuleAttribute']]:
        """
        The rule's stringified attribute.
        """
        return pulumi.get(self, "attribute")

    @attribute.setter
    def attribute(self, value: Optional[pulumi.Input['DevicePoolRuleAttribute']]):
        pulumi.set(self, "attribute", value)

    @property
    @pulumi.getter
    def operator(self) -> Optional[pulumi.Input['DevicePoolRuleOperator']]:
        """
        Specifies how Device Farm compares the rule's attribute to the value.
        """
        return pulumi.get(self, "operator")

    @operator.setter
    def operator(self, value: Optional[pulumi.Input['DevicePoolRuleOperator']]):
        pulumi.set(self, "operator", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        """
        The rule's value.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class DevicePoolTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class InstanceProfileTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class NetworkProfileTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class ProjectTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class ProjectVpcConfigArgs:
    def __init__(__self__, *,
                 security_group_ids: pulumi.Input[Sequence[pulumi.Input[str]]],
                 subnet_ids: pulumi.Input[Sequence[pulumi.Input[str]]],
                 vpc_id: pulumi.Input[str]):
        """
        The VPC security groups and subnets that are attached to a project
        :param pulumi.Input[Sequence[pulumi.Input[str]]] security_group_ids: An array of security group Ids in your Amazon VPC
        :param pulumi.Input[Sequence[pulumi.Input[str]]] subnet_ids: A array of subnet IDs in your Amazon VPC.
        :param pulumi.Input[str] vpc_id: The ID of the Amazon VPC
        """
        pulumi.set(__self__, "security_group_ids", security_group_ids)
        pulumi.set(__self__, "subnet_ids", subnet_ids)
        pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        An array of security group Ids in your Amazon VPC
        """
        return pulumi.get(self, "security_group_ids")

    @security_group_ids.setter
    def security_group_ids(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "security_group_ids", value)

    @property
    @pulumi.getter(name="subnetIds")
    def subnet_ids(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        A array of subnet IDs in your Amazon VPC.
        """
        return pulumi.get(self, "subnet_ids")

    @subnet_ids.setter
    def subnet_ids(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "subnet_ids", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Input[str]:
        """
        The ID of the Amazon VPC
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vpc_id", value)


@pulumi.input_type
class TestGridProjectTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class TestGridProjectVpcConfigArgs:
    def __init__(__self__, *,
                 security_group_ids: pulumi.Input[Sequence[pulumi.Input[str]]],
                 subnet_ids: pulumi.Input[Sequence[pulumi.Input[str]]],
                 vpc_id: pulumi.Input[str]):
        """
        The VPC security groups and subnets that are attached to a TestGrid project.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] security_group_ids: A list of VPC security group IDs in your Amazon VPC.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] subnet_ids: A list of VPC subnet IDs in your Amazon VPC.
        """
        pulumi.set(__self__, "security_group_ids", security_group_ids)
        pulumi.set(__self__, "subnet_ids", subnet_ids)
        pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        A list of VPC security group IDs in your Amazon VPC.
        """
        return pulumi.get(self, "security_group_ids")

    @security_group_ids.setter
    def security_group_ids(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "security_group_ids", value)

    @property
    @pulumi.getter(name="subnetIds")
    def subnet_ids(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        A list of VPC subnet IDs in your Amazon VPC.
        """
        return pulumi.get(self, "subnet_ids")

    @subnet_ids.setter
    def subnet_ids(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "subnet_ids", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vpc_id", value)


@pulumi.input_type
class VpceConfigurationTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


