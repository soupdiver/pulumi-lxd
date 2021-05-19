# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'ContainerDeviceArgs',
    'ContainerFileArgs',
    'ProfileDeviceArgs',
    'ProviderLXDRemoteArgs',
]

@pulumi.input_type
class ContainerDeviceArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 properties: pulumi.Input[Mapping[str, Any]],
                 type: pulumi.Input[str]):
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "properties", properties)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def properties(self) -> pulumi.Input[Mapping[str, Any]]:
        return pulumi.get(self, "properties")

    @properties.setter
    def properties(self, value: pulumi.Input[Mapping[str, Any]]):
        pulumi.set(self, "properties", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class ContainerFileArgs:
    def __init__(__self__, *,
                 target_file: pulumi.Input[str],
                 content: Optional[pulumi.Input[str]] = None,
                 create_directories: Optional[pulumi.Input[bool]] = None,
                 gid: Optional[pulumi.Input[int]] = None,
                 mode: Optional[pulumi.Input[str]] = None,
                 source: Optional[pulumi.Input[str]] = None,
                 uid: Optional[pulumi.Input[int]] = None):
        pulumi.set(__self__, "target_file", target_file)
        if content is not None:
            pulumi.set(__self__, "content", content)
        if create_directories is not None:
            pulumi.set(__self__, "create_directories", create_directories)
        if gid is not None:
            pulumi.set(__self__, "gid", gid)
        if mode is not None:
            pulumi.set(__self__, "mode", mode)
        if source is not None:
            pulumi.set(__self__, "source", source)
        if uid is not None:
            pulumi.set(__self__, "uid", uid)

    @property
    @pulumi.getter(name="targetFile")
    def target_file(self) -> pulumi.Input[str]:
        return pulumi.get(self, "target_file")

    @target_file.setter
    def target_file(self, value: pulumi.Input[str]):
        pulumi.set(self, "target_file", value)

    @property
    @pulumi.getter
    def content(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "content")

    @content.setter
    def content(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "content", value)

    @property
    @pulumi.getter(name="createDirectories")
    def create_directories(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "create_directories")

    @create_directories.setter
    def create_directories(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "create_directories", value)

    @property
    @pulumi.getter
    def gid(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "gid")

    @gid.setter
    def gid(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "gid", value)

    @property
    @pulumi.getter
    def mode(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mode", value)

    @property
    @pulumi.getter
    def source(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "source")

    @source.setter
    def source(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source", value)

    @property
    @pulumi.getter
    def uid(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "uid")

    @uid.setter
    def uid(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "uid", value)


@pulumi.input_type
class ProfileDeviceArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 properties: pulumi.Input[Mapping[str, Any]],
                 type: pulumi.Input[str]):
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "properties", properties)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def properties(self) -> pulumi.Input[Mapping[str, Any]]:
        return pulumi.get(self, "properties")

    @properties.setter
    def properties(self, value: pulumi.Input[Mapping[str, Any]]):
        pulumi.set(self, "properties", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class ProviderLXDRemoteArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 address: Optional[pulumi.Input[str]] = None,
                 default: Optional[pulumi.Input[bool]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[str]] = None,
                 scheme: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "name", name)
        if address is not None:
            pulumi.set(__self__, "address", address)
        if default is not None:
            pulumi.set(__self__, "default", default)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if scheme is not None:
            pulumi.set(__self__, "scheme", scheme)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def address(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address", value)

    @property
    @pulumi.getter
    def default(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "default")

    @default.setter
    def default(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "default", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter
    def scheme(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "scheme")

    @scheme.setter
    def scheme(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scheme", value)

