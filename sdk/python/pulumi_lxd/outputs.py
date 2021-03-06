# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'ContainerDevice',
    'ContainerFile',
    'ProfileDevice',
    'ProviderLXDRemote',
]

@pulumi.output_type
class ContainerDevice(dict):
    def __init__(__self__, *,
                 name: str,
                 properties: Mapping[str, Any],
                 type: str):
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "properties", properties)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def properties(self) -> Mapping[str, Any]:
        return pulumi.get(self, "properties")

    @property
    @pulumi.getter
    def type(self) -> str:
        return pulumi.get(self, "type")


@pulumi.output_type
class ContainerFile(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "targetFile":
            suggest = "target_file"
        elif key == "createDirectories":
            suggest = "create_directories"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ContainerFile. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ContainerFile.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ContainerFile.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 target_file: str,
                 content: Optional[str] = None,
                 create_directories: Optional[bool] = None,
                 gid: Optional[int] = None,
                 mode: Optional[str] = None,
                 source: Optional[str] = None,
                 uid: Optional[int] = None):
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
    def target_file(self) -> str:
        return pulumi.get(self, "target_file")

    @property
    @pulumi.getter
    def content(self) -> Optional[str]:
        return pulumi.get(self, "content")

    @property
    @pulumi.getter(name="createDirectories")
    def create_directories(self) -> Optional[bool]:
        return pulumi.get(self, "create_directories")

    @property
    @pulumi.getter
    def gid(self) -> Optional[int]:
        return pulumi.get(self, "gid")

    @property
    @pulumi.getter
    def mode(self) -> Optional[str]:
        return pulumi.get(self, "mode")

    @property
    @pulumi.getter
    def source(self) -> Optional[str]:
        return pulumi.get(self, "source")

    @property
    @pulumi.getter
    def uid(self) -> Optional[int]:
        return pulumi.get(self, "uid")


@pulumi.output_type
class ProfileDevice(dict):
    def __init__(__self__, *,
                 name: str,
                 properties: Mapping[str, Any],
                 type: str):
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "properties", properties)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def properties(self) -> Mapping[str, Any]:
        return pulumi.get(self, "properties")

    @property
    @pulumi.getter
    def type(self) -> str:
        return pulumi.get(self, "type")


@pulumi.output_type
class ProviderLXDRemote(dict):
    def __init__(__self__, *,
                 name: str,
                 address: Optional[str] = None,
                 default: Optional[bool] = None,
                 password: Optional[str] = None,
                 port: Optional[str] = None,
                 scheme: Optional[str] = None):
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
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def address(self) -> Optional[str]:
        return pulumi.get(self, "address")

    @property
    @pulumi.getter
    def default(self) -> Optional[bool]:
        return pulumi.get(self, "default")

    @property
    @pulumi.getter
    def password(self) -> Optional[str]:
        return pulumi.get(self, "password")

    @property
    @pulumi.getter
    def port(self) -> Optional[str]:
        return pulumi.get(self, "port")

    @property
    @pulumi.getter
    def scheme(self) -> Optional[str]:
        return pulumi.get(self, "scheme")


