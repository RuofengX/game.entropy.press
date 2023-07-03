import 10_time_pb2 as _10_time_pb2
import 11_velocity_pb2 as _11_velocity_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Entity(_message.Message):
    __slots__ = ["id", "time", "velo"]
    ID_FIELD_NUMBER: _ClassVar[int]
    TIME_FIELD_NUMBER: _ClassVar[int]
    VELO_FIELD_NUMBER: _ClassVar[int]
    id: int
    time: _10_time_pb2.Time
    velo: _11_velocity_pb2.Velocity
    def __init__(self, id: _Optional[int] = ..., time: _Optional[_Union[_10_time_pb2.Time, _Mapping]] = ..., velo: _Optional[_Union[_11_velocity_pb2.Velocity, _Mapping]] = ...) -> None: ...

class Field(_message.Message):
    __slots__ = ["entity"]
    ENTITY_FIELD_NUMBER: _ClassVar[int]
    entity: _containers.RepeatedCompositeFieldContainer[Entity]
    def __init__(self, entity: _Optional[_Iterable[_Union[Entity, _Mapping]]] = ...) -> None: ...
