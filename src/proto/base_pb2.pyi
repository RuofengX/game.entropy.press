import time_pb2 as _time_pb2
import velocity_pb2 as _velocity_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Entity(_message.Message):
    __slots__ = ["ID", "time", "velo"]
    ID_FIELD_NUMBER: _ClassVar[int]
    TIME_FIELD_NUMBER: _ClassVar[int]
    VELO_FIELD_NUMBER: _ClassVar[int]
    ID: int
    time: _time_pb2.Fragment
    velo: _velocity_pb2.Fragment
    def __init__(self, ID: _Optional[int] = ..., time: _Optional[_Union[_time_pb2.Fragment, _Mapping]] = ..., velo: _Optional[_Union[_velocity_pb2.Fragment, _Mapping]] = ...) -> None: ...

class Field(_message.Message):
    __slots__ = ["entity"]
    ENTITY_FIELD_NUMBER: _ClassVar[int]
    entity: _containers.RepeatedCompositeFieldContainer[Entity]
    def __init__(self, entity: _Optional[_Iterable[_Union[Entity, _Mapping]]] = ...) -> None: ...
