# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from enum import Enum

__all__ = [
    'LaunchProfileAutomaticTerminationMode',
    'LaunchProfileSessionBackupMode',
    'LaunchProfileSessionPersistenceMode',
    'LaunchProfileStreamingClipboardMode',
    'LaunchProfileStreamingInstanceType',
    'LaunchProfileStreamingSessionStorageMode',
    'StreamingImageEncryptionConfigurationKeyType',
    'StudioComponentInitializationScriptRunContext',
    'StudioComponentLaunchProfilePlatform',
    'StudioComponentSubtype',
    'StudioComponentType',
    'StudioEncryptionConfigurationKeyType',
]


class LaunchProfileAutomaticTerminationMode(str, Enum):
    DEACTIVATED = "DEACTIVATED"
    ACTIVATED = "ACTIVATED"


class LaunchProfileSessionBackupMode(str, Enum):
    AUTOMATIC = "AUTOMATIC"
    DEACTIVATED = "DEACTIVATED"


class LaunchProfileSessionPersistenceMode(str, Enum):
    DEACTIVATED = "DEACTIVATED"
    ACTIVATED = "ACTIVATED"


class LaunchProfileStreamingClipboardMode(str, Enum):
    ENABLED = "ENABLED"
    DISABLED = "DISABLED"


class LaunchProfileStreamingInstanceType(str, Enum):
    G4DN_XLARGE = "g4dn.xlarge"
    G4DN2XLARGE = "g4dn.2xlarge"
    G4DN4XLARGE = "g4dn.4xlarge"
    G4DN8XLARGE = "g4dn.8xlarge"
    G4DN12XLARGE = "g4dn.12xlarge"
    G4DN16XLARGE = "g4dn.16xlarge"
    G34XLARGE = "g3.4xlarge"
    G3S_XLARGE = "g3s.xlarge"
    G5_XLARGE = "g5.xlarge"
    G52XLARGE = "g5.2xlarge"
    G54XLARGE = "g5.4xlarge"
    G58XLARGE = "g5.8xlarge"
    G516XLARGE = "g5.16xlarge"


class LaunchProfileStreamingSessionStorageMode(str, Enum):
    UPLOAD = "UPLOAD"


class StreamingImageEncryptionConfigurationKeyType(str, Enum):
    """
    <p/>
    """
    CUSTOMER_MANAGED_KEY = "CUSTOMER_MANAGED_KEY"


class StudioComponentInitializationScriptRunContext(str, Enum):
    SYSTEM_INITIALIZATION = "SYSTEM_INITIALIZATION"
    USER_INITIALIZATION = "USER_INITIALIZATION"


class StudioComponentLaunchProfilePlatform(str, Enum):
    LINUX = "LINUX"
    WINDOWS = "WINDOWS"


class StudioComponentSubtype(str, Enum):
    AWS_MANAGED_MICROSOFT_AD = "AWS_MANAGED_MICROSOFT_AD"
    AMAZON_FSX_FOR_WINDOWS = "AMAZON_FSX_FOR_WINDOWS"
    AMAZON_FSX_FOR_LUSTRE = "AMAZON_FSX_FOR_LUSTRE"
    CUSTOM = "CUSTOM"


class StudioComponentType(str, Enum):
    ACTIVE_DIRECTORY = "ACTIVE_DIRECTORY"
    SHARED_FILE_SYSTEM = "SHARED_FILE_SYSTEM"
    COMPUTE_FARM = "COMPUTE_FARM"
    LICENSE_SERVICE = "LICENSE_SERVICE"
    CUSTOM = "CUSTOM"


class StudioEncryptionConfigurationKeyType(str, Enum):
    """
    <p>The type of KMS key that is used to encrypt studio data.</p>
    """
    AWS_OWNED_KEY = "AWS_OWNED_KEY"
    CUSTOMER_MANAGED_KEY = "CUSTOMER_MANAGED_KEY"
