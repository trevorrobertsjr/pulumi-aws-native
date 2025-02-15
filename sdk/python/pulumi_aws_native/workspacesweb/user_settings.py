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

__all__ = ['UserSettingsArgs', 'UserSettings']

@pulumi.input_type
class UserSettingsArgs:
    def __init__(__self__, *,
                 copy_allowed: pulumi.Input['UserSettingsEnabledType'],
                 download_allowed: pulumi.Input['UserSettingsEnabledType'],
                 paste_allowed: pulumi.Input['UserSettingsEnabledType'],
                 print_allowed: pulumi.Input['UserSettingsEnabledType'],
                 upload_allowed: pulumi.Input['UserSettingsEnabledType'],
                 additional_encryption_context: Optional[pulumi.Input['UserSettingsEncryptionContextMapArgs']] = None,
                 cookie_synchronization_configuration: Optional[pulumi.Input['UserSettingsCookieSynchronizationConfigurationArgs']] = None,
                 customer_managed_key: Optional[pulumi.Input[str]] = None,
                 disconnect_timeout_in_minutes: Optional[pulumi.Input[float]] = None,
                 idle_disconnect_timeout_in_minutes: Optional[pulumi.Input[float]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['UserSettingsTagArgs']]]] = None):
        """
        The set of arguments for constructing a UserSettings resource.
        """
        pulumi.set(__self__, "copy_allowed", copy_allowed)
        pulumi.set(__self__, "download_allowed", download_allowed)
        pulumi.set(__self__, "paste_allowed", paste_allowed)
        pulumi.set(__self__, "print_allowed", print_allowed)
        pulumi.set(__self__, "upload_allowed", upload_allowed)
        if additional_encryption_context is not None:
            pulumi.set(__self__, "additional_encryption_context", additional_encryption_context)
        if cookie_synchronization_configuration is not None:
            pulumi.set(__self__, "cookie_synchronization_configuration", cookie_synchronization_configuration)
        if customer_managed_key is not None:
            pulumi.set(__self__, "customer_managed_key", customer_managed_key)
        if disconnect_timeout_in_minutes is not None:
            pulumi.set(__self__, "disconnect_timeout_in_minutes", disconnect_timeout_in_minutes)
        if idle_disconnect_timeout_in_minutes is not None:
            pulumi.set(__self__, "idle_disconnect_timeout_in_minutes", idle_disconnect_timeout_in_minutes)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="copyAllowed")
    def copy_allowed(self) -> pulumi.Input['UserSettingsEnabledType']:
        return pulumi.get(self, "copy_allowed")

    @copy_allowed.setter
    def copy_allowed(self, value: pulumi.Input['UserSettingsEnabledType']):
        pulumi.set(self, "copy_allowed", value)

    @property
    @pulumi.getter(name="downloadAllowed")
    def download_allowed(self) -> pulumi.Input['UserSettingsEnabledType']:
        return pulumi.get(self, "download_allowed")

    @download_allowed.setter
    def download_allowed(self, value: pulumi.Input['UserSettingsEnabledType']):
        pulumi.set(self, "download_allowed", value)

    @property
    @pulumi.getter(name="pasteAllowed")
    def paste_allowed(self) -> pulumi.Input['UserSettingsEnabledType']:
        return pulumi.get(self, "paste_allowed")

    @paste_allowed.setter
    def paste_allowed(self, value: pulumi.Input['UserSettingsEnabledType']):
        pulumi.set(self, "paste_allowed", value)

    @property
    @pulumi.getter(name="printAllowed")
    def print_allowed(self) -> pulumi.Input['UserSettingsEnabledType']:
        return pulumi.get(self, "print_allowed")

    @print_allowed.setter
    def print_allowed(self, value: pulumi.Input['UserSettingsEnabledType']):
        pulumi.set(self, "print_allowed", value)

    @property
    @pulumi.getter(name="uploadAllowed")
    def upload_allowed(self) -> pulumi.Input['UserSettingsEnabledType']:
        return pulumi.get(self, "upload_allowed")

    @upload_allowed.setter
    def upload_allowed(self, value: pulumi.Input['UserSettingsEnabledType']):
        pulumi.set(self, "upload_allowed", value)

    @property
    @pulumi.getter(name="additionalEncryptionContext")
    def additional_encryption_context(self) -> Optional[pulumi.Input['UserSettingsEncryptionContextMapArgs']]:
        return pulumi.get(self, "additional_encryption_context")

    @additional_encryption_context.setter
    def additional_encryption_context(self, value: Optional[pulumi.Input['UserSettingsEncryptionContextMapArgs']]):
        pulumi.set(self, "additional_encryption_context", value)

    @property
    @pulumi.getter(name="cookieSynchronizationConfiguration")
    def cookie_synchronization_configuration(self) -> Optional[pulumi.Input['UserSettingsCookieSynchronizationConfigurationArgs']]:
        return pulumi.get(self, "cookie_synchronization_configuration")

    @cookie_synchronization_configuration.setter
    def cookie_synchronization_configuration(self, value: Optional[pulumi.Input['UserSettingsCookieSynchronizationConfigurationArgs']]):
        pulumi.set(self, "cookie_synchronization_configuration", value)

    @property
    @pulumi.getter(name="customerManagedKey")
    def customer_managed_key(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "customer_managed_key")

    @customer_managed_key.setter
    def customer_managed_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "customer_managed_key", value)

    @property
    @pulumi.getter(name="disconnectTimeoutInMinutes")
    def disconnect_timeout_in_minutes(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "disconnect_timeout_in_minutes")

    @disconnect_timeout_in_minutes.setter
    def disconnect_timeout_in_minutes(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "disconnect_timeout_in_minutes", value)

    @property
    @pulumi.getter(name="idleDisconnectTimeoutInMinutes")
    def idle_disconnect_timeout_in_minutes(self) -> Optional[pulumi.Input[float]]:
        return pulumi.get(self, "idle_disconnect_timeout_in_minutes")

    @idle_disconnect_timeout_in_minutes.setter
    def idle_disconnect_timeout_in_minutes(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "idle_disconnect_timeout_in_minutes", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['UserSettingsTagArgs']]]]:
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['UserSettingsTagArgs']]]]):
        pulumi.set(self, "tags", value)


class UserSettings(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 additional_encryption_context: Optional[pulumi.Input[pulumi.InputType['UserSettingsEncryptionContextMapArgs']]] = None,
                 cookie_synchronization_configuration: Optional[pulumi.Input[pulumi.InputType['UserSettingsCookieSynchronizationConfigurationArgs']]] = None,
                 copy_allowed: Optional[pulumi.Input['UserSettingsEnabledType']] = None,
                 customer_managed_key: Optional[pulumi.Input[str]] = None,
                 disconnect_timeout_in_minutes: Optional[pulumi.Input[float]] = None,
                 download_allowed: Optional[pulumi.Input['UserSettingsEnabledType']] = None,
                 idle_disconnect_timeout_in_minutes: Optional[pulumi.Input[float]] = None,
                 paste_allowed: Optional[pulumi.Input['UserSettingsEnabledType']] = None,
                 print_allowed: Optional[pulumi.Input['UserSettingsEnabledType']] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['UserSettingsTagArgs']]]]] = None,
                 upload_allowed: Optional[pulumi.Input['UserSettingsEnabledType']] = None,
                 __props__=None):
        """
        Definition of AWS::WorkSpacesWeb::UserSettings Resource Type

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: UserSettingsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Definition of AWS::WorkSpacesWeb::UserSettings Resource Type

        :param str resource_name: The name of the resource.
        :param UserSettingsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(UserSettingsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 additional_encryption_context: Optional[pulumi.Input[pulumi.InputType['UserSettingsEncryptionContextMapArgs']]] = None,
                 cookie_synchronization_configuration: Optional[pulumi.Input[pulumi.InputType['UserSettingsCookieSynchronizationConfigurationArgs']]] = None,
                 copy_allowed: Optional[pulumi.Input['UserSettingsEnabledType']] = None,
                 customer_managed_key: Optional[pulumi.Input[str]] = None,
                 disconnect_timeout_in_minutes: Optional[pulumi.Input[float]] = None,
                 download_allowed: Optional[pulumi.Input['UserSettingsEnabledType']] = None,
                 idle_disconnect_timeout_in_minutes: Optional[pulumi.Input[float]] = None,
                 paste_allowed: Optional[pulumi.Input['UserSettingsEnabledType']] = None,
                 print_allowed: Optional[pulumi.Input['UserSettingsEnabledType']] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['UserSettingsTagArgs']]]]] = None,
                 upload_allowed: Optional[pulumi.Input['UserSettingsEnabledType']] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = UserSettingsArgs.__new__(UserSettingsArgs)

            __props__.__dict__["additional_encryption_context"] = additional_encryption_context
            __props__.__dict__["cookie_synchronization_configuration"] = cookie_synchronization_configuration
            if copy_allowed is None and not opts.urn:
                raise TypeError("Missing required property 'copy_allowed'")
            __props__.__dict__["copy_allowed"] = copy_allowed
            __props__.__dict__["customer_managed_key"] = customer_managed_key
            __props__.__dict__["disconnect_timeout_in_minutes"] = disconnect_timeout_in_minutes
            if download_allowed is None and not opts.urn:
                raise TypeError("Missing required property 'download_allowed'")
            __props__.__dict__["download_allowed"] = download_allowed
            __props__.__dict__["idle_disconnect_timeout_in_minutes"] = idle_disconnect_timeout_in_minutes
            if paste_allowed is None and not opts.urn:
                raise TypeError("Missing required property 'paste_allowed'")
            __props__.__dict__["paste_allowed"] = paste_allowed
            if print_allowed is None and not opts.urn:
                raise TypeError("Missing required property 'print_allowed'")
            __props__.__dict__["print_allowed"] = print_allowed
            __props__.__dict__["tags"] = tags
            if upload_allowed is None and not opts.urn:
                raise TypeError("Missing required property 'upload_allowed'")
            __props__.__dict__["upload_allowed"] = upload_allowed
            __props__.__dict__["associated_portal_arns"] = None
            __props__.__dict__["user_settings_arn"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["additional_encryption_context", "customer_managed_key"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(UserSettings, __self__).__init__(
            'aws-native:workspacesweb:UserSettings',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'UserSettings':
        """
        Get an existing UserSettings resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = UserSettingsArgs.__new__(UserSettingsArgs)

        __props__.__dict__["additional_encryption_context"] = None
        __props__.__dict__["associated_portal_arns"] = None
        __props__.__dict__["cookie_synchronization_configuration"] = None
        __props__.__dict__["copy_allowed"] = None
        __props__.__dict__["customer_managed_key"] = None
        __props__.__dict__["disconnect_timeout_in_minutes"] = None
        __props__.__dict__["download_allowed"] = None
        __props__.__dict__["idle_disconnect_timeout_in_minutes"] = None
        __props__.__dict__["paste_allowed"] = None
        __props__.__dict__["print_allowed"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["upload_allowed"] = None
        __props__.__dict__["user_settings_arn"] = None
        return UserSettings(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="additionalEncryptionContext")
    def additional_encryption_context(self) -> pulumi.Output[Optional['outputs.UserSettingsEncryptionContextMap']]:
        return pulumi.get(self, "additional_encryption_context")

    @property
    @pulumi.getter(name="associatedPortalArns")
    def associated_portal_arns(self) -> pulumi.Output[Sequence[str]]:
        return pulumi.get(self, "associated_portal_arns")

    @property
    @pulumi.getter(name="cookieSynchronizationConfiguration")
    def cookie_synchronization_configuration(self) -> pulumi.Output[Optional['outputs.UserSettingsCookieSynchronizationConfiguration']]:
        return pulumi.get(self, "cookie_synchronization_configuration")

    @property
    @pulumi.getter(name="copyAllowed")
    def copy_allowed(self) -> pulumi.Output['UserSettingsEnabledType']:
        return pulumi.get(self, "copy_allowed")

    @property
    @pulumi.getter(name="customerManagedKey")
    def customer_managed_key(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "customer_managed_key")

    @property
    @pulumi.getter(name="disconnectTimeoutInMinutes")
    def disconnect_timeout_in_minutes(self) -> pulumi.Output[Optional[float]]:
        return pulumi.get(self, "disconnect_timeout_in_minutes")

    @property
    @pulumi.getter(name="downloadAllowed")
    def download_allowed(self) -> pulumi.Output['UserSettingsEnabledType']:
        return pulumi.get(self, "download_allowed")

    @property
    @pulumi.getter(name="idleDisconnectTimeoutInMinutes")
    def idle_disconnect_timeout_in_minutes(self) -> pulumi.Output[Optional[float]]:
        return pulumi.get(self, "idle_disconnect_timeout_in_minutes")

    @property
    @pulumi.getter(name="pasteAllowed")
    def paste_allowed(self) -> pulumi.Output['UserSettingsEnabledType']:
        return pulumi.get(self, "paste_allowed")

    @property
    @pulumi.getter(name="printAllowed")
    def print_allowed(self) -> pulumi.Output['UserSettingsEnabledType']:
        return pulumi.get(self, "print_allowed")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.UserSettingsTag']]]:
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="uploadAllowed")
    def upload_allowed(self) -> pulumi.Output['UserSettingsEnabledType']:
        return pulumi.get(self, "upload_allowed")

    @property
    @pulumi.getter(name="userSettingsArn")
    def user_settings_arn(self) -> pulumi.Output[str]:
        return pulumi.get(self, "user_settings_arn")

