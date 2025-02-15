# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'DomainAdvancedSecurityOptionsInputArgs',
    'DomainClusterConfigArgs',
    'DomainCognitoOptionsArgs',
    'DomainEbsOptionsArgs',
    'DomainEncryptionAtRestOptionsArgs',
    'DomainEndpointOptionsArgs',
    'DomainIdpArgs',
    'DomainMasterUserOptionsArgs',
    'DomainNodeToNodeEncryptionOptionsArgs',
    'DomainOffPeakWindowOptionsArgs',
    'DomainOffPeakWindowArgs',
    'DomainSamlOptionsArgs',
    'DomainSnapshotOptionsArgs',
    'DomainSoftwareUpdateOptionsArgs',
    'DomainTagArgs',
    'DomainVpcOptionsArgs',
    'DomainWindowStartTimeArgs',
    'DomainZoneAwarenessConfigArgs',
]

@pulumi.input_type
class DomainAdvancedSecurityOptionsInputArgs:
    def __init__(__self__, *,
                 anonymous_auth_disable_date: Optional[pulumi.Input[str]] = None,
                 anonymous_auth_enabled: Optional[pulumi.Input[bool]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 internal_user_database_enabled: Optional[pulumi.Input[bool]] = None,
                 master_user_options: Optional[pulumi.Input['DomainMasterUserOptionsArgs']] = None,
                 saml_options: Optional[pulumi.Input['DomainSamlOptionsArgs']] = None):
        if anonymous_auth_disable_date is not None:
            pulumi.set(__self__, "anonymous_auth_disable_date", anonymous_auth_disable_date)
        if anonymous_auth_enabled is not None:
            pulumi.set(__self__, "anonymous_auth_enabled", anonymous_auth_enabled)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if internal_user_database_enabled is not None:
            pulumi.set(__self__, "internal_user_database_enabled", internal_user_database_enabled)
        if master_user_options is not None:
            pulumi.set(__self__, "master_user_options", master_user_options)
        if saml_options is not None:
            pulumi.set(__self__, "saml_options", saml_options)

    @property
    @pulumi.getter(name="anonymousAuthDisableDate")
    def anonymous_auth_disable_date(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "anonymous_auth_disable_date")

    @anonymous_auth_disable_date.setter
    def anonymous_auth_disable_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "anonymous_auth_disable_date", value)

    @property
    @pulumi.getter(name="anonymousAuthEnabled")
    def anonymous_auth_enabled(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "anonymous_auth_enabled")

    @anonymous_auth_enabled.setter
    def anonymous_auth_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "anonymous_auth_enabled", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="internalUserDatabaseEnabled")
    def internal_user_database_enabled(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "internal_user_database_enabled")

    @internal_user_database_enabled.setter
    def internal_user_database_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "internal_user_database_enabled", value)

    @property
    @pulumi.getter(name="masterUserOptions")
    def master_user_options(self) -> Optional[pulumi.Input['DomainMasterUserOptionsArgs']]:
        return pulumi.get(self, "master_user_options")

    @master_user_options.setter
    def master_user_options(self, value: Optional[pulumi.Input['DomainMasterUserOptionsArgs']]):
        pulumi.set(self, "master_user_options", value)

    @property
    @pulumi.getter(name="samlOptions")
    def saml_options(self) -> Optional[pulumi.Input['DomainSamlOptionsArgs']]:
        return pulumi.get(self, "saml_options")

    @saml_options.setter
    def saml_options(self, value: Optional[pulumi.Input['DomainSamlOptionsArgs']]):
        pulumi.set(self, "saml_options", value)


@pulumi.input_type
class DomainClusterConfigArgs:
    def __init__(__self__, *,
                 dedicated_master_count: Optional[pulumi.Input[int]] = None,
                 dedicated_master_enabled: Optional[pulumi.Input[bool]] = None,
                 dedicated_master_type: Optional[pulumi.Input[str]] = None,
                 instance_count: Optional[pulumi.Input[int]] = None,
                 instance_type: Optional[pulumi.Input[str]] = None,
                 multi_az_with_standby_enabled: Optional[pulumi.Input[bool]] = None,
                 warm_count: Optional[pulumi.Input[int]] = None,
                 warm_enabled: Optional[pulumi.Input[bool]] = None,
                 warm_type: Optional[pulumi.Input[str]] = None,
                 zone_awareness_config: Optional[pulumi.Input['DomainZoneAwarenessConfigArgs']] = None,
                 zone_awareness_enabled: Optional[pulumi.Input[bool]] = None):
        if dedicated_master_count is not None:
            pulumi.set(__self__, "dedicated_master_count", dedicated_master_count)
        if dedicated_master_enabled is not None:
            pulumi.set(__self__, "dedicated_master_enabled", dedicated_master_enabled)
        if dedicated_master_type is not None:
            pulumi.set(__self__, "dedicated_master_type", dedicated_master_type)
        if instance_count is not None:
            pulumi.set(__self__, "instance_count", instance_count)
        if instance_type is not None:
            pulumi.set(__self__, "instance_type", instance_type)
        if multi_az_with_standby_enabled is not None:
            pulumi.set(__self__, "multi_az_with_standby_enabled", multi_az_with_standby_enabled)
        if warm_count is not None:
            pulumi.set(__self__, "warm_count", warm_count)
        if warm_enabled is not None:
            pulumi.set(__self__, "warm_enabled", warm_enabled)
        if warm_type is not None:
            pulumi.set(__self__, "warm_type", warm_type)
        if zone_awareness_config is not None:
            pulumi.set(__self__, "zone_awareness_config", zone_awareness_config)
        if zone_awareness_enabled is not None:
            pulumi.set(__self__, "zone_awareness_enabled", zone_awareness_enabled)

    @property
    @pulumi.getter(name="dedicatedMasterCount")
    def dedicated_master_count(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "dedicated_master_count")

    @dedicated_master_count.setter
    def dedicated_master_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "dedicated_master_count", value)

    @property
    @pulumi.getter(name="dedicatedMasterEnabled")
    def dedicated_master_enabled(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "dedicated_master_enabled")

    @dedicated_master_enabled.setter
    def dedicated_master_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "dedicated_master_enabled", value)

    @property
    @pulumi.getter(name="dedicatedMasterType")
    def dedicated_master_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "dedicated_master_type")

    @dedicated_master_type.setter
    def dedicated_master_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dedicated_master_type", value)

    @property
    @pulumi.getter(name="instanceCount")
    def instance_count(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "instance_count")

    @instance_count.setter
    def instance_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "instance_count", value)

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "instance_type")

    @instance_type.setter
    def instance_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_type", value)

    @property
    @pulumi.getter(name="multiAzWithStandbyEnabled")
    def multi_az_with_standby_enabled(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "multi_az_with_standby_enabled")

    @multi_az_with_standby_enabled.setter
    def multi_az_with_standby_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "multi_az_with_standby_enabled", value)

    @property
    @pulumi.getter(name="warmCount")
    def warm_count(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "warm_count")

    @warm_count.setter
    def warm_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "warm_count", value)

    @property
    @pulumi.getter(name="warmEnabled")
    def warm_enabled(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "warm_enabled")

    @warm_enabled.setter
    def warm_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "warm_enabled", value)

    @property
    @pulumi.getter(name="warmType")
    def warm_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "warm_type")

    @warm_type.setter
    def warm_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "warm_type", value)

    @property
    @pulumi.getter(name="zoneAwarenessConfig")
    def zone_awareness_config(self) -> Optional[pulumi.Input['DomainZoneAwarenessConfigArgs']]:
        return pulumi.get(self, "zone_awareness_config")

    @zone_awareness_config.setter
    def zone_awareness_config(self, value: Optional[pulumi.Input['DomainZoneAwarenessConfigArgs']]):
        pulumi.set(self, "zone_awareness_config", value)

    @property
    @pulumi.getter(name="zoneAwarenessEnabled")
    def zone_awareness_enabled(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "zone_awareness_enabled")

    @zone_awareness_enabled.setter
    def zone_awareness_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "zone_awareness_enabled", value)


@pulumi.input_type
class DomainCognitoOptionsArgs:
    def __init__(__self__, *,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 identity_pool_id: Optional[pulumi.Input[str]] = None,
                 role_arn: Optional[pulumi.Input[str]] = None,
                 user_pool_id: Optional[pulumi.Input[str]] = None):
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if identity_pool_id is not None:
            pulumi.set(__self__, "identity_pool_id", identity_pool_id)
        if role_arn is not None:
            pulumi.set(__self__, "role_arn", role_arn)
        if user_pool_id is not None:
            pulumi.set(__self__, "user_pool_id", user_pool_id)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="identityPoolId")
    def identity_pool_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "identity_pool_id")

    @identity_pool_id.setter
    def identity_pool_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "identity_pool_id", value)

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "role_arn")

    @role_arn.setter
    def role_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role_arn", value)

    @property
    @pulumi.getter(name="userPoolId")
    def user_pool_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "user_pool_id")

    @user_pool_id.setter
    def user_pool_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_pool_id", value)


@pulumi.input_type
class DomainEbsOptionsArgs:
    def __init__(__self__, *,
                 ebs_enabled: Optional[pulumi.Input[bool]] = None,
                 iops: Optional[pulumi.Input[int]] = None,
                 throughput: Optional[pulumi.Input[int]] = None,
                 volume_size: Optional[pulumi.Input[int]] = None,
                 volume_type: Optional[pulumi.Input[str]] = None):
        if ebs_enabled is not None:
            pulumi.set(__self__, "ebs_enabled", ebs_enabled)
        if iops is not None:
            pulumi.set(__self__, "iops", iops)
        if throughput is not None:
            pulumi.set(__self__, "throughput", throughput)
        if volume_size is not None:
            pulumi.set(__self__, "volume_size", volume_size)
        if volume_type is not None:
            pulumi.set(__self__, "volume_type", volume_type)

    @property
    @pulumi.getter(name="ebsEnabled")
    def ebs_enabled(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "ebs_enabled")

    @ebs_enabled.setter
    def ebs_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "ebs_enabled", value)

    @property
    @pulumi.getter
    def iops(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "iops")

    @iops.setter
    def iops(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "iops", value)

    @property
    @pulumi.getter
    def throughput(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "throughput")

    @throughput.setter
    def throughput(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "throughput", value)

    @property
    @pulumi.getter(name="volumeSize")
    def volume_size(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "volume_size")

    @volume_size.setter
    def volume_size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "volume_size", value)

    @property
    @pulumi.getter(name="volumeType")
    def volume_type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "volume_type")

    @volume_type.setter
    def volume_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "volume_type", value)


@pulumi.input_type
class DomainEncryptionAtRestOptionsArgs:
    def __init__(__self__, *,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 kms_key_id: Optional[pulumi.Input[str]] = None):
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if kms_key_id is not None:
            pulumi.set(__self__, "kms_key_id", kms_key_id)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="kmsKeyId")
    def kms_key_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "kms_key_id")

    @kms_key_id.setter
    def kms_key_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kms_key_id", value)


@pulumi.input_type
class DomainEndpointOptionsArgs:
    def __init__(__self__, *,
                 custom_endpoint: Optional[pulumi.Input[str]] = None,
                 custom_endpoint_certificate_arn: Optional[pulumi.Input[str]] = None,
                 custom_endpoint_enabled: Optional[pulumi.Input[bool]] = None,
                 enforce_https: Optional[pulumi.Input[bool]] = None,
                 tls_security_policy: Optional[pulumi.Input[str]] = None):
        if custom_endpoint is not None:
            pulumi.set(__self__, "custom_endpoint", custom_endpoint)
        if custom_endpoint_certificate_arn is not None:
            pulumi.set(__self__, "custom_endpoint_certificate_arn", custom_endpoint_certificate_arn)
        if custom_endpoint_enabled is not None:
            pulumi.set(__self__, "custom_endpoint_enabled", custom_endpoint_enabled)
        if enforce_https is not None:
            pulumi.set(__self__, "enforce_https", enforce_https)
        if tls_security_policy is not None:
            pulumi.set(__self__, "tls_security_policy", tls_security_policy)

    @property
    @pulumi.getter(name="customEndpoint")
    def custom_endpoint(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "custom_endpoint")

    @custom_endpoint.setter
    def custom_endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "custom_endpoint", value)

    @property
    @pulumi.getter(name="customEndpointCertificateArn")
    def custom_endpoint_certificate_arn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "custom_endpoint_certificate_arn")

    @custom_endpoint_certificate_arn.setter
    def custom_endpoint_certificate_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "custom_endpoint_certificate_arn", value)

    @property
    @pulumi.getter(name="customEndpointEnabled")
    def custom_endpoint_enabled(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "custom_endpoint_enabled")

    @custom_endpoint_enabled.setter
    def custom_endpoint_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "custom_endpoint_enabled", value)

    @property
    @pulumi.getter(name="enforceHttps")
    def enforce_https(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "enforce_https")

    @enforce_https.setter
    def enforce_https(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enforce_https", value)

    @property
    @pulumi.getter(name="tlsSecurityPolicy")
    def tls_security_policy(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "tls_security_policy")

    @tls_security_policy.setter
    def tls_security_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tls_security_policy", value)


@pulumi.input_type
class DomainIdpArgs:
    def __init__(__self__, *,
                 entity_id: pulumi.Input[str],
                 metadata_content: pulumi.Input[str]):
        pulumi.set(__self__, "entity_id", entity_id)
        pulumi.set(__self__, "metadata_content", metadata_content)

    @property
    @pulumi.getter(name="entityId")
    def entity_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "entity_id")

    @entity_id.setter
    def entity_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "entity_id", value)

    @property
    @pulumi.getter(name="metadataContent")
    def metadata_content(self) -> pulumi.Input[str]:
        return pulumi.get(self, "metadata_content")

    @metadata_content.setter
    def metadata_content(self, value: pulumi.Input[str]):
        pulumi.set(self, "metadata_content", value)


@pulumi.input_type
class DomainMasterUserOptionsArgs:
    def __init__(__self__, *,
                 master_user_arn: Optional[pulumi.Input[str]] = None,
                 master_user_name: Optional[pulumi.Input[str]] = None,
                 master_user_password: Optional[pulumi.Input[str]] = None):
        if master_user_arn is not None:
            pulumi.set(__self__, "master_user_arn", master_user_arn)
        if master_user_name is not None:
            pulumi.set(__self__, "master_user_name", master_user_name)
        if master_user_password is not None:
            pulumi.set(__self__, "master_user_password", master_user_password)

    @property
    @pulumi.getter(name="masterUserArn")
    def master_user_arn(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "master_user_arn")

    @master_user_arn.setter
    def master_user_arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "master_user_arn", value)

    @property
    @pulumi.getter(name="masterUserName")
    def master_user_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "master_user_name")

    @master_user_name.setter
    def master_user_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "master_user_name", value)

    @property
    @pulumi.getter(name="masterUserPassword")
    def master_user_password(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "master_user_password")

    @master_user_password.setter
    def master_user_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "master_user_password", value)


@pulumi.input_type
class DomainNodeToNodeEncryptionOptionsArgs:
    def __init__(__self__, *,
                 enabled: Optional[pulumi.Input[bool]] = None):
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)


@pulumi.input_type
class DomainOffPeakWindowOptionsArgs:
    def __init__(__self__, *,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 off_peak_window: Optional[pulumi.Input['DomainOffPeakWindowArgs']] = None):
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if off_peak_window is not None:
            pulumi.set(__self__, "off_peak_window", off_peak_window)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="offPeakWindow")
    def off_peak_window(self) -> Optional[pulumi.Input['DomainOffPeakWindowArgs']]:
        return pulumi.get(self, "off_peak_window")

    @off_peak_window.setter
    def off_peak_window(self, value: Optional[pulumi.Input['DomainOffPeakWindowArgs']]):
        pulumi.set(self, "off_peak_window", value)


@pulumi.input_type
class DomainOffPeakWindowArgs:
    def __init__(__self__, *,
                 window_start_time: Optional[pulumi.Input['DomainWindowStartTimeArgs']] = None):
        if window_start_time is not None:
            pulumi.set(__self__, "window_start_time", window_start_time)

    @property
    @pulumi.getter(name="windowStartTime")
    def window_start_time(self) -> Optional[pulumi.Input['DomainWindowStartTimeArgs']]:
        return pulumi.get(self, "window_start_time")

    @window_start_time.setter
    def window_start_time(self, value: Optional[pulumi.Input['DomainWindowStartTimeArgs']]):
        pulumi.set(self, "window_start_time", value)


@pulumi.input_type
class DomainSamlOptionsArgs:
    def __init__(__self__, *,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 idp: Optional[pulumi.Input['DomainIdpArgs']] = None,
                 master_backend_role: Optional[pulumi.Input[str]] = None,
                 master_user_name: Optional[pulumi.Input[str]] = None,
                 roles_key: Optional[pulumi.Input[str]] = None,
                 session_timeout_minutes: Optional[pulumi.Input[int]] = None,
                 subject_key: Optional[pulumi.Input[str]] = None):
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if idp is not None:
            pulumi.set(__self__, "idp", idp)
        if master_backend_role is not None:
            pulumi.set(__self__, "master_backend_role", master_backend_role)
        if master_user_name is not None:
            pulumi.set(__self__, "master_user_name", master_user_name)
        if roles_key is not None:
            pulumi.set(__self__, "roles_key", roles_key)
        if session_timeout_minutes is not None:
            pulumi.set(__self__, "session_timeout_minutes", session_timeout_minutes)
        if subject_key is not None:
            pulumi.set(__self__, "subject_key", subject_key)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter
    def idp(self) -> Optional[pulumi.Input['DomainIdpArgs']]:
        return pulumi.get(self, "idp")

    @idp.setter
    def idp(self, value: Optional[pulumi.Input['DomainIdpArgs']]):
        pulumi.set(self, "idp", value)

    @property
    @pulumi.getter(name="masterBackendRole")
    def master_backend_role(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "master_backend_role")

    @master_backend_role.setter
    def master_backend_role(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "master_backend_role", value)

    @property
    @pulumi.getter(name="masterUserName")
    def master_user_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "master_user_name")

    @master_user_name.setter
    def master_user_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "master_user_name", value)

    @property
    @pulumi.getter(name="rolesKey")
    def roles_key(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "roles_key")

    @roles_key.setter
    def roles_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "roles_key", value)

    @property
    @pulumi.getter(name="sessionTimeoutMinutes")
    def session_timeout_minutes(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "session_timeout_minutes")

    @session_timeout_minutes.setter
    def session_timeout_minutes(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "session_timeout_minutes", value)

    @property
    @pulumi.getter(name="subjectKey")
    def subject_key(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "subject_key")

    @subject_key.setter
    def subject_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subject_key", value)


@pulumi.input_type
class DomainSnapshotOptionsArgs:
    def __init__(__self__, *,
                 automated_snapshot_start_hour: Optional[pulumi.Input[int]] = None):
        if automated_snapshot_start_hour is not None:
            pulumi.set(__self__, "automated_snapshot_start_hour", automated_snapshot_start_hour)

    @property
    @pulumi.getter(name="automatedSnapshotStartHour")
    def automated_snapshot_start_hour(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "automated_snapshot_start_hour")

    @automated_snapshot_start_hour.setter
    def automated_snapshot_start_hour(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "automated_snapshot_start_hour", value)


@pulumi.input_type
class DomainSoftwareUpdateOptionsArgs:
    def __init__(__self__, *,
                 auto_software_update_enabled: Optional[pulumi.Input[bool]] = None):
        if auto_software_update_enabled is not None:
            pulumi.set(__self__, "auto_software_update_enabled", auto_software_update_enabled)

    @property
    @pulumi.getter(name="autoSoftwareUpdateEnabled")
    def auto_software_update_enabled(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "auto_software_update_enabled")

    @auto_software_update_enabled.setter
    def auto_software_update_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_software_update_enabled", value)


@pulumi.input_type
class DomainTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        """
        :param pulumi.Input[str] key: The value of the tag.
        :param pulumi.Input[str] value: The key of the tag.
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        """
        The value of the tag.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        """
        The key of the tag.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class DomainVpcOptionsArgs:
    def __init__(__self__, *,
                 security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 subnet_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        if security_group_ids is not None:
            pulumi.set(__self__, "security_group_ids", security_group_ids)
        if subnet_ids is not None:
            pulumi.set(__self__, "subnet_ids", subnet_ids)

    @property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "security_group_ids")

    @security_group_ids.setter
    def security_group_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "security_group_ids", value)

    @property
    @pulumi.getter(name="subnetIds")
    def subnet_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "subnet_ids")

    @subnet_ids.setter
    def subnet_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "subnet_ids", value)


@pulumi.input_type
class DomainWindowStartTimeArgs:
    def __init__(__self__, *,
                 hours: pulumi.Input[int],
                 minutes: pulumi.Input[int]):
        pulumi.set(__self__, "hours", hours)
        pulumi.set(__self__, "minutes", minutes)

    @property
    @pulumi.getter
    def hours(self) -> pulumi.Input[int]:
        return pulumi.get(self, "hours")

    @hours.setter
    def hours(self, value: pulumi.Input[int]):
        pulumi.set(self, "hours", value)

    @property
    @pulumi.getter
    def minutes(self) -> pulumi.Input[int]:
        return pulumi.get(self, "minutes")

    @minutes.setter
    def minutes(self, value: pulumi.Input[int]):
        pulumi.set(self, "minutes", value)


@pulumi.input_type
class DomainZoneAwarenessConfigArgs:
    def __init__(__self__, *,
                 availability_zone_count: Optional[pulumi.Input[int]] = None):
        if availability_zone_count is not None:
            pulumi.set(__self__, "availability_zone_count", availability_zone_count)

    @property
    @pulumi.getter(name="availabilityZoneCount")
    def availability_zone_count(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "availability_zone_count")

    @availability_zone_count.setter
    def availability_zone_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "availability_zone_count", value)


