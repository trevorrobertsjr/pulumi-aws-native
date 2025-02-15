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
    'AllowListCriteriaArgs',
    'AllowListTagArgs',
    'CustomDataIdentifierTagArgs',
    'FindingsFilterCriterionArgs',
    'FindingsFilterFindingCriteriaArgs',
    'FindingsFilterTagArgs',
]

@pulumi.input_type
class AllowListCriteriaArgs:
    def __init__(__self__):
        """
        The regex or s3 object to use for the AllowList.
        """
        pass


@pulumi.input_type
class AllowListTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        """
        A key-value pair to associate with a resource.
        :param pulumi.Input[str] key: The tag's key.
        :param pulumi.Input[str] value: The tag's value.
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        """
        The tag's key.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        """
        The tag's value.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class CustomDataIdentifierTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        """
        A key-value pair to associate with a resource.
        :param pulumi.Input[str] key: The tag's key.
        :param pulumi.Input[str] value: The tag's value.
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        """
        The tag's key.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        """
        The tag's value.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class FindingsFilterCriterionArgs:
    def __init__(__self__):
        """
        Map of filter criteria.
        """
        pass


@pulumi.input_type
class FindingsFilterFindingCriteriaArgs:
    def __init__(__self__, *,
                 criterion: Optional[pulumi.Input['FindingsFilterCriterionArgs']] = None):
        if criterion is not None:
            pulumi.set(__self__, "criterion", criterion)

    @property
    @pulumi.getter
    def criterion(self) -> Optional[pulumi.Input['FindingsFilterCriterionArgs']]:
        return pulumi.get(self, "criterion")

    @criterion.setter
    def criterion(self, value: Optional[pulumi.Input['FindingsFilterCriterionArgs']]):
        pulumi.set(self, "criterion", value)


@pulumi.input_type
class FindingsFilterTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        """
        A key-value pair to associate with a resource.
        :param pulumi.Input[str] key: The tag's key.
        :param pulumi.Input[str] value: The tag's value.
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        """
        The tag's key.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        """
        The tag's value.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


