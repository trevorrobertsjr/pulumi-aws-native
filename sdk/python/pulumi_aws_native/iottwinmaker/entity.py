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

__all__ = ['EntityArgs', 'Entity']

@pulumi.input_type
class EntityArgs:
    def __init__(__self__, *,
                 workspace_id: pulumi.Input[str],
                 components: Optional[Any] = None,
                 composite_components: Optional[Any] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 entity_id: Optional[pulumi.Input[str]] = None,
                 entity_name: Optional[pulumi.Input[str]] = None,
                 parent_entity_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[Any] = None):
        """
        The set of arguments for constructing a Entity resource.
        :param pulumi.Input[str] workspace_id: The ID of the workspace.
        :param Any components: A map that sets information about a component type.
        :param Any composite_components: A map that sets information about a composite component.
        :param pulumi.Input[str] description: The description of the entity.
        :param pulumi.Input[str] entity_id: The ID of the entity.
        :param pulumi.Input[str] entity_name: The name of the entity.
        :param pulumi.Input[str] parent_entity_id: The ID of the parent entity.
        :param Any tags: A key-value pair to associate with a resource.
        """
        pulumi.set(__self__, "workspace_id", workspace_id)
        if components is not None:
            pulumi.set(__self__, "components", components)
        if composite_components is not None:
            pulumi.set(__self__, "composite_components", composite_components)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if entity_id is not None:
            pulumi.set(__self__, "entity_id", entity_id)
        if entity_name is not None:
            pulumi.set(__self__, "entity_name", entity_name)
        if parent_entity_id is not None:
            pulumi.set(__self__, "parent_entity_id", parent_entity_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="workspaceId")
    def workspace_id(self) -> pulumi.Input[str]:
        """
        The ID of the workspace.
        """
        return pulumi.get(self, "workspace_id")

    @workspace_id.setter
    def workspace_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "workspace_id", value)

    @property
    @pulumi.getter
    def components(self) -> Optional[Any]:
        """
        A map that sets information about a component type.
        """
        return pulumi.get(self, "components")

    @components.setter
    def components(self, value: Optional[Any]):
        pulumi.set(self, "components", value)

    @property
    @pulumi.getter(name="compositeComponents")
    def composite_components(self) -> Optional[Any]:
        """
        A map that sets information about a composite component.
        """
        return pulumi.get(self, "composite_components")

    @composite_components.setter
    def composite_components(self, value: Optional[Any]):
        pulumi.set(self, "composite_components", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the entity.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="entityId")
    def entity_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the entity.
        """
        return pulumi.get(self, "entity_id")

    @entity_id.setter
    def entity_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "entity_id", value)

    @property
    @pulumi.getter(name="entityName")
    def entity_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the entity.
        """
        return pulumi.get(self, "entity_name")

    @entity_name.setter
    def entity_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "entity_name", value)

    @property
    @pulumi.getter(name="parentEntityId")
    def parent_entity_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the parent entity.
        """
        return pulumi.get(self, "parent_entity_id")

    @parent_entity_id.setter
    def parent_entity_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "parent_entity_id", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[Any]:
        """
        A key-value pair to associate with a resource.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[Any]):
        pulumi.set(self, "tags", value)


class Entity(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 components: Optional[Any] = None,
                 composite_components: Optional[Any] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 entity_id: Optional[pulumi.Input[str]] = None,
                 entity_name: Optional[pulumi.Input[str]] = None,
                 parent_entity_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[Any] = None,
                 workspace_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Resource schema for AWS::IoTTwinMaker::Entity

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param Any components: A map that sets information about a component type.
        :param Any composite_components: A map that sets information about a composite component.
        :param pulumi.Input[str] description: The description of the entity.
        :param pulumi.Input[str] entity_id: The ID of the entity.
        :param pulumi.Input[str] entity_name: The name of the entity.
        :param pulumi.Input[str] parent_entity_id: The ID of the parent entity.
        :param Any tags: A key-value pair to associate with a resource.
        :param pulumi.Input[str] workspace_id: The ID of the workspace.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EntityArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource schema for AWS::IoTTwinMaker::Entity

        :param str resource_name: The name of the resource.
        :param EntityArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EntityArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 components: Optional[Any] = None,
                 composite_components: Optional[Any] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 entity_id: Optional[pulumi.Input[str]] = None,
                 entity_name: Optional[pulumi.Input[str]] = None,
                 parent_entity_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[Any] = None,
                 workspace_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EntityArgs.__new__(EntityArgs)

            __props__.__dict__["components"] = components
            __props__.__dict__["composite_components"] = composite_components
            __props__.__dict__["description"] = description
            __props__.__dict__["entity_id"] = entity_id
            __props__.__dict__["entity_name"] = entity_name
            __props__.__dict__["parent_entity_id"] = parent_entity_id
            __props__.__dict__["tags"] = tags
            if workspace_id is None and not opts.urn:
                raise TypeError("Missing required property 'workspace_id'")
            __props__.__dict__["workspace_id"] = workspace_id
            __props__.__dict__["arn"] = None
            __props__.__dict__["creation_date_time"] = None
            __props__.__dict__["has_child_entities"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["update_date_time"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["entity_id", "workspace_id"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(Entity, __self__).__init__(
            'aws-native:iottwinmaker:Entity',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Entity':
        """
        Get an existing Entity resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = EntityArgs.__new__(EntityArgs)

        __props__.__dict__["arn"] = None
        __props__.__dict__["components"] = None
        __props__.__dict__["composite_components"] = None
        __props__.__dict__["creation_date_time"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["entity_id"] = None
        __props__.__dict__["entity_name"] = None
        __props__.__dict__["has_child_entities"] = None
        __props__.__dict__["parent_entity_id"] = None
        __props__.__dict__["status"] = None
        __props__.__dict__["tags"] = None
        __props__.__dict__["update_date_time"] = None
        __props__.__dict__["workspace_id"] = None
        return Entity(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        The ARN of the entity.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def components(self) -> pulumi.Output[Optional[Any]]:
        """
        A map that sets information about a component type.
        """
        return pulumi.get(self, "components")

    @property
    @pulumi.getter(name="compositeComponents")
    def composite_components(self) -> pulumi.Output[Optional[Any]]:
        """
        A map that sets information about a composite component.
        """
        return pulumi.get(self, "composite_components")

    @property
    @pulumi.getter(name="creationDateTime")
    def creation_date_time(self) -> pulumi.Output[str]:
        """
        The date and time when the entity was created.
        """
        return pulumi.get(self, "creation_date_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the entity.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="entityId")
    def entity_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the entity.
        """
        return pulumi.get(self, "entity_id")

    @property
    @pulumi.getter(name="entityName")
    def entity_name(self) -> pulumi.Output[str]:
        """
        The name of the entity.
        """
        return pulumi.get(self, "entity_name")

    @property
    @pulumi.getter(name="hasChildEntities")
    def has_child_entities(self) -> pulumi.Output[bool]:
        """
        A Boolean value that specifies whether the entity has child entities or not.
        """
        return pulumi.get(self, "has_child_entities")

    @property
    @pulumi.getter(name="parentEntityId")
    def parent_entity_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the parent entity.
        """
        return pulumi.get(self, "parent_entity_id")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output['outputs.EntityStatus']:
        """
        The current status of the entity.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Any]]:
        """
        A key-value pair to associate with a resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="updateDateTime")
    def update_date_time(self) -> pulumi.Output[str]:
        """
        The last date and time when the entity was updated.
        """
        return pulumi.get(self, "update_date_time")

    @property
    @pulumi.getter(name="workspaceId")
    def workspace_id(self) -> pulumi.Output[str]:
        """
        The ID of the workspace.
        """
        return pulumi.get(self, "workspace_id")

