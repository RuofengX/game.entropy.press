"""
@generated by mypy-protobuf.  Do not edit manually!
isort:skip_file
"""
import builtins
import collections.abc
import essence.time_pb2
import essence.velocity_pb2
import google.protobuf.descriptor
import google.protobuf.internal.containers
import google.protobuf.message
import sys

if sys.version_info >= (3, 8):
    import typing as typing_extensions
else:
    import typing_extensions

DESCRIPTOR: google.protobuf.descriptor.FileDescriptor

@typing_extensions.final
class Entity(google.protobuf.message.Message):
    """一个模块计算所需的所有数据"""

    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    ID_FIELD_NUMBER: builtins.int
    TIME_FIELD_NUMBER: builtins.int
    VELO_FIELD_NUMBER: builtins.int
    ID: builtins.int
    """哈希字段
    ID
    """
    @property
    def time(self) -> essence.time_pb2.Fragment:
        """承载字段
        时间grpc_out
        """
    @property
    def velo(self) -> essence.velocity_pb2.Fragment:
        """空间"""
    def __init__(
        self,
        *,
        ID: builtins.int = ...,
        time: essence.time_pb2.Fragment | None = ...,
        velo: essence.velocity_pb2.Fragment | None = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["time", b"time", "velo", b"velo"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["ID", b"ID", "time", b"time", "velo", b"velo"]) -> None: ...

global___Entity = Entity

@typing_extensions.final
class Field(google.protobuf.message.Message):
    """对应一个特定的包含很多实体的场"""

    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    ENTITY_FIELD_NUMBER: builtins.int
    @property
    def entity(self) -> google.protobuf.internal.containers.RepeatedCompositeFieldContainer[global___Entity]: ...
    def __init__(
        self,
        *,
        entity: collections.abc.Iterable[global___Entity] | None = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["entity", b"entity"]) -> None: ...

global___Field = Field