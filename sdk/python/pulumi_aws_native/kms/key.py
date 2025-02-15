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
from ._enums import *
from ._inputs import *

__all__ = ['KeyArgs', 'Key']

@pulumi.input_type
class KeyArgs:
    def __init__(__self__, *,
                 bypass_policy_lockout_safety_check: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enable_key_rotation: Optional[pulumi.Input[bool]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 key_policy: Optional[Any] = None,
                 key_spec: Optional[pulumi.Input['KeySpec']] = None,
                 key_usage: Optional[pulumi.Input['KeyUsage']] = None,
                 multi_region: Optional[pulumi.Input[bool]] = None,
                 origin: Optional[pulumi.Input['KeyOrigin']] = None,
                 pending_window_in_days: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['KeyTagArgs']]]] = None):
        """
        The set of arguments for constructing a Key resource.
        :param pulumi.Input[bool] bypass_policy_lockout_safety_check: Skips ("bypasses") the key policy lockout safety check. The default value is false.
        :param pulumi.Input[str] description: A description of the AWS KMS key. Use a description that helps you to distinguish this AWS KMS key from others in the account, such as its intended use.
        :param pulumi.Input[bool] enable_key_rotation: Enables automatic rotation of the key material for the specified AWS KMS key. By default, automation key rotation is not enabled.
        :param pulumi.Input[bool] enabled: Specifies whether the AWS KMS key is enabled. Disabled AWS KMS keys cannot be used in cryptographic operations.
        :param Any key_policy: The key policy that authorizes use of the AWS KMS key. The key policy must observe the following rules.
        :param pulumi.Input['KeySpec'] key_spec: Specifies the type of AWS KMS key to create. The default value is SYMMETRIC_DEFAULT. This property is required only for asymmetric AWS KMS keys. You can't change the KeySpec value after the AWS KMS key is created.
        :param pulumi.Input['KeyUsage'] key_usage: Determines the cryptographic operations for which you can use the AWS KMS key. The default value is ENCRYPT_DECRYPT. This property is required only for asymmetric AWS KMS keys. You can't change the KeyUsage value after the AWS KMS key is created.
        :param pulumi.Input[bool] multi_region: Specifies whether the AWS KMS key should be Multi-Region. You can't change the MultiRegion value after the AWS KMS key is created.
        :param pulumi.Input['KeyOrigin'] origin: The source of the key material for the KMS key. You cannot change the origin after you create the KMS key. The default is AWS_KMS, which means that AWS KMS creates the key material.
        :param pulumi.Input[int] pending_window_in_days: Specifies the number of days in the waiting period before AWS KMS deletes an AWS KMS key that has been removed from a CloudFormation stack. Enter a value between 7 and 30 days. The default value is 30 days.
        :param pulumi.Input[Sequence[pulumi.Input['KeyTagArgs']]] tags: An array of key-value pairs to apply to this resource.
        """
        if bypass_policy_lockout_safety_check is not None:
            pulumi.set(__self__, "bypass_policy_lockout_safety_check", bypass_policy_lockout_safety_check)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if enable_key_rotation is not None:
            pulumi.set(__self__, "enable_key_rotation", enable_key_rotation)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if key_policy is not None:
            pulumi.set(__self__, "key_policy", key_policy)
        if key_spec is not None:
            pulumi.set(__self__, "key_spec", key_spec)
        if key_usage is not None:
            pulumi.set(__self__, "key_usage", key_usage)
        if multi_region is not None:
            pulumi.set(__self__, "multi_region", multi_region)
        if origin is not None:
            pulumi.set(__self__, "origin", origin)
        if pending_window_in_days is not None:
            pulumi.set(__self__, "pending_window_in_days", pending_window_in_days)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="bypassPolicyLockoutSafetyCheck")
    def bypass_policy_lockout_safety_check(self) -> Optional[pulumi.Input[bool]]:
        """
        Skips ("bypasses") the key policy lockout safety check. The default value is false.
        """
        return pulumi.get(self, "bypass_policy_lockout_safety_check")

    @bypass_policy_lockout_safety_check.setter
    def bypass_policy_lockout_safety_check(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "bypass_policy_lockout_safety_check", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description of the AWS KMS key. Use a description that helps you to distinguish this AWS KMS key from others in the account, such as its intended use.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="enableKeyRotation")
    def enable_key_rotation(self) -> Optional[pulumi.Input[bool]]:
        """
        Enables automatic rotation of the key material for the specified AWS KMS key. By default, automation key rotation is not enabled.
        """
        return pulumi.get(self, "enable_key_rotation")

    @enable_key_rotation.setter
    def enable_key_rotation(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_key_rotation", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether the AWS KMS key is enabled. Disabled AWS KMS keys cannot be used in cryptographic operations.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="keyPolicy")
    def key_policy(self) -> Optional[Any]:
        """
        The key policy that authorizes use of the AWS KMS key. The key policy must observe the following rules.
        """
        return pulumi.get(self, "key_policy")

    @key_policy.setter
    def key_policy(self, value: Optional[Any]):
        pulumi.set(self, "key_policy", value)

    @property
    @pulumi.getter(name="keySpec")
    def key_spec(self) -> Optional[pulumi.Input['KeySpec']]:
        """
        Specifies the type of AWS KMS key to create. The default value is SYMMETRIC_DEFAULT. This property is required only for asymmetric AWS KMS keys. You can't change the KeySpec value after the AWS KMS key is created.
        """
        return pulumi.get(self, "key_spec")

    @key_spec.setter
    def key_spec(self, value: Optional[pulumi.Input['KeySpec']]):
        pulumi.set(self, "key_spec", value)

    @property
    @pulumi.getter(name="keyUsage")
    def key_usage(self) -> Optional[pulumi.Input['KeyUsage']]:
        """
        Determines the cryptographic operations for which you can use the AWS KMS key. The default value is ENCRYPT_DECRYPT. This property is required only for asymmetric AWS KMS keys. You can't change the KeyUsage value after the AWS KMS key is created.
        """
        return pulumi.get(self, "key_usage")

    @key_usage.setter
    def key_usage(self, value: Optional[pulumi.Input['KeyUsage']]):
        pulumi.set(self, "key_usage", value)

    @property
    @pulumi.getter(name="multiRegion")
    def multi_region(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether the AWS KMS key should be Multi-Region. You can't change the MultiRegion value after the AWS KMS key is created.
        """
        return pulumi.get(self, "multi_region")

    @multi_region.setter
    def multi_region(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "multi_region", value)

    @property
    @pulumi.getter
    def origin(self) -> Optional[pulumi.Input['KeyOrigin']]:
        """
        The source of the key material for the KMS key. You cannot change the origin after you create the KMS key. The default is AWS_KMS, which means that AWS KMS creates the key material.
        """
        return pulumi.get(self, "origin")

    @origin.setter
    def origin(self, value: Optional[pulumi.Input['KeyOrigin']]):
        pulumi.set(self, "origin", value)

    @property
    @pulumi.getter(name="pendingWindowInDays")
    def pending_window_in_days(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the number of days in the waiting period before AWS KMS deletes an AWS KMS key that has been removed from a CloudFormation stack. Enter a value between 7 and 30 days. The default value is 30 days.
        """
        return pulumi.get(self, "pending_window_in_days")

    @pending_window_in_days.setter
    def pending_window_in_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "pending_window_in_days", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['KeyTagArgs']]]]:
        """
        An array of key-value pairs to apply to this resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['KeyTagArgs']]]]):
        pulumi.set(self, "tags", value)


class Key(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bypass_policy_lockout_safety_check: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enable_key_rotation: Optional[pulumi.Input[bool]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 key_policy: Optional[Any] = None,
                 key_spec: Optional[pulumi.Input['KeySpec']] = None,
                 key_usage: Optional[pulumi.Input['KeyUsage']] = None,
                 multi_region: Optional[pulumi.Input[bool]] = None,
                 origin: Optional[pulumi.Input['KeyOrigin']] = None,
                 pending_window_in_days: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['KeyTagArgs']]]]] = None,
                 __props__=None):
        """
        The AWS::KMS::Key resource specifies an AWS KMS key in AWS Key Management Service (AWS KMS). Authorized users can use the AWS KMS key to encrypt and decrypt small amounts of data (up to 4096 bytes), but they are more commonly used to generate data keys. You can also use AWS KMS keys to encrypt data stored in AWS services that are integrated with AWS KMS or within their applications.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] bypass_policy_lockout_safety_check: Skips ("bypasses") the key policy lockout safety check. The default value is false.
        :param pulumi.Input[str] description: A description of the AWS KMS key. Use a description that helps you to distinguish this AWS KMS key from others in the account, such as its intended use.
        :param pulumi.Input[bool] enable_key_rotation: Enables automatic rotation of the key material for the specified AWS KMS key. By default, automation key rotation is not enabled.
        :param pulumi.Input[bool] enabled: Specifies whether the AWS KMS key is enabled. Disabled AWS KMS keys cannot be used in cryptographic operations.
        :param Any key_policy: The key policy that authorizes use of the AWS KMS key. The key policy must observe the following rules.
        :param pulumi.Input['KeySpec'] key_spec: Specifies the type of AWS KMS key to create. The default value is SYMMETRIC_DEFAULT. This property is required only for asymmetric AWS KMS keys. You can't change the KeySpec value after the AWS KMS key is created.
        :param pulumi.Input['KeyUsage'] key_usage: Determines the cryptographic operations for which you can use the AWS KMS key. The default value is ENCRYPT_DECRYPT. This property is required only for asymmetric AWS KMS keys. You can't change the KeyUsage value after the AWS KMS key is created.
        :param pulumi.Input[bool] multi_region: Specifies whether the AWS KMS key should be Multi-Region. You can't change the MultiRegion value after the AWS KMS key is created.
        :param pulumi.Input['KeyOrigin'] origin: The source of the key material for the KMS key. You cannot change the origin after you create the KMS key. The default is AWS_KMS, which means that AWS KMS creates the key material.
        :param pulumi.Input[int] pending_window_in_days: Specifies the number of days in the waiting period before AWS KMS deletes an AWS KMS key that has been removed from a CloudFormation stack. Enter a value between 7 and 30 days. The default value is 30 days.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['KeyTagArgs']]]] tags: An array of key-value pairs to apply to this resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[KeyArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The AWS::KMS::Key resource specifies an AWS KMS key in AWS Key Management Service (AWS KMS). Authorized users can use the AWS KMS key to encrypt and decrypt small amounts of data (up to 4096 bytes), but they are more commonly used to generate data keys. You can also use AWS KMS keys to encrypt data stored in AWS services that are integrated with AWS KMS or within their applications.

        :param str resource_name: The name of the resource.
        :param KeyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(KeyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bypass_policy_lockout_safety_check: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enable_key_rotation: Optional[pulumi.Input[bool]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 key_policy: Optional[Any] = None,
                 key_spec: Optional[pulumi.Input['KeySpec']] = None,
                 key_usage: Optional[pulumi.Input['KeyUsage']] = None,
                 multi_region: Optional[pulumi.Input[bool]] = None,
                 origin: Optional[pulumi.Input['KeyOrigin']] = None,
                 pending_window_in_days: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['KeyTagArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = KeyArgs.__new__(KeyArgs)

            __props__.__dict__["bypass_policy_lockout_safety_check"] = bypass_policy_lockout_safety_check
            __props__.__dict__["description"] = description
            __props__.__dict__["enable_key_rotation"] = enable_key_rotation
            __props__.__dict__["enabled"] = enabled
            __props__.__dict__["key_policy"] = key_policy
            __props__.__dict__["key_spec"] = key_spec
            __props__.__dict__["key_usage"] = key_usage
            __props__.__dict__["multi_region"] = multi_region
            __props__.__dict__["origin"] = origin
            __props__.__dict__["pending_window_in_days"] = pending_window_in_days
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["key_id"] = None
        super(Key, __self__).__init__(
            'aws-native:kms:Key',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Key':
        """
        Get an existing Key resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = KeyArgs.__new__(KeyArgs)

        __props__.__dict__["arn"] = None
        __props__.__dict__["bypass_policy_lockout_safety_check"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["enable_key_rotation"] = None
        __props__.__dict__["enabled"] = None
        __props__.__dict__["key_id"] = None
        __props__.__dict__["key_policy"] = None
        __props__.__dict__["key_spec"] = None
        __props__.__dict__["key_usage"] = None
        __props__.__dict__["multi_region"] = None
        __props__.__dict__["origin"] = None
        __props__.__dict__["pending_window_in_days"] = None
        __props__.__dict__["tags"] = None
        return Key(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter(name="bypassPolicyLockoutSafetyCheck")
    def bypass_policy_lockout_safety_check(self) -> pulumi.Output[Optional[bool]]:
        """
        Skips ("bypasses") the key policy lockout safety check. The default value is false.
        """
        return pulumi.get(self, "bypass_policy_lockout_safety_check")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A description of the AWS KMS key. Use a description that helps you to distinguish this AWS KMS key from others in the account, such as its intended use.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="enableKeyRotation")
    def enable_key_rotation(self) -> pulumi.Output[Optional[bool]]:
        """
        Enables automatic rotation of the key material for the specified AWS KMS key. By default, automation key rotation is not enabled.
        """
        return pulumi.get(self, "enable_key_rotation")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies whether the AWS KMS key is enabled. Disabled AWS KMS keys cannot be used in cryptographic operations.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="keyId")
    def key_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "key_id")

    @property
    @pulumi.getter(name="keyPolicy")
    def key_policy(self) -> pulumi.Output[Optional[Any]]:
        """
        The key policy that authorizes use of the AWS KMS key. The key policy must observe the following rules.
        """
        return pulumi.get(self, "key_policy")

    @property
    @pulumi.getter(name="keySpec")
    def key_spec(self) -> pulumi.Output[Optional['KeySpec']]:
        """
        Specifies the type of AWS KMS key to create. The default value is SYMMETRIC_DEFAULT. This property is required only for asymmetric AWS KMS keys. You can't change the KeySpec value after the AWS KMS key is created.
        """
        return pulumi.get(self, "key_spec")

    @property
    @pulumi.getter(name="keyUsage")
    def key_usage(self) -> pulumi.Output[Optional['KeyUsage']]:
        """
        Determines the cryptographic operations for which you can use the AWS KMS key. The default value is ENCRYPT_DECRYPT. This property is required only for asymmetric AWS KMS keys. You can't change the KeyUsage value after the AWS KMS key is created.
        """
        return pulumi.get(self, "key_usage")

    @property
    @pulumi.getter(name="multiRegion")
    def multi_region(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies whether the AWS KMS key should be Multi-Region. You can't change the MultiRegion value after the AWS KMS key is created.
        """
        return pulumi.get(self, "multi_region")

    @property
    @pulumi.getter
    def origin(self) -> pulumi.Output[Optional['KeyOrigin']]:
        """
        The source of the key material for the KMS key. You cannot change the origin after you create the KMS key. The default is AWS_KMS, which means that AWS KMS creates the key material.
        """
        return pulumi.get(self, "origin")

    @property
    @pulumi.getter(name="pendingWindowInDays")
    def pending_window_in_days(self) -> pulumi.Output[Optional[int]]:
        """
        Specifies the number of days in the waiting period before AWS KMS deletes an AWS KMS key that has been removed from a CloudFormation stack. Enter a value between 7 and 30 days. The default value is 30 days.
        """
        return pulumi.get(self, "pending_window_in_days")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.KeyTag']]]:
        """
        An array of key-value pairs to apply to this resource.
        """
        return pulumi.get(self, "tags")

