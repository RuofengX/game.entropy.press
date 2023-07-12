from essence import time_pb2 as _time_pb2
from essence import velocity_pb2 as _velocity_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Entity(_message.Message):
    __slots__ = ["ID", "time", "time_d", "velo", "velo_d"]
    ID_FIELD_NUMBER: _ClassVar[int]
    TIME_FIELD_NUMBER: _ClassVar[int]
    TIME_D_FIELD_NUMBER: _ClassVar[int]
    VELO_FIELD_NUMBER: _ClassVar[int]
    VELO_D_FIELD_NUMBER: _ClassVar[int]
    ID: int
    time: _time_pb2.Property
    time_d: _time_pb2.Delta
    velo: _velocity_pb2.Property
    velo_d: _velocity_pb2.Delta
    def __init__(self, ID: _Optional[int] = ..., time: _Optional[_Union[_time_pb2.Property, _Mapping]] = ..., time_d: _Optional[_Union[_time_pb2.Delta, _Mapping]] = ..., velo: _Optional[_Union[_velocity_pb2.Property, _Mapping]] = ..., velo_d: _Optional[_Union[_velocity_pb2.Delta, _Mapping]] = ...) -> None: ...

class Space(_message.Message):
    __slots__ = ["entity"]
    ENTITY_FIELD_NUMBER: _ClassVar[int]
    entity: _containers.RepeatedCompositeFieldContainer[Entity]
    def __init__(self, entity: _Optional[_Iterable[_Union[Entity, _Mapping]]] = ...) -> None: ...
