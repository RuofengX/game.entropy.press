from essence import time_pb2 as _time_pb2
from essence import velocity_pb2 as _velocity_pb2
from essence import structure_pb2 as _structure_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Entity(_message.Message):
    __slots__ = ["ID", "time", "velo", "structure"]
    ID_FIELD_NUMBER: _ClassVar[int]
    TIME_FIELD_NUMBER: _ClassVar[int]
    VELO_FIELD_NUMBER: _ClassVar[int]
    STRUCTURE_FIELD_NUMBER: _ClassVar[int]
    ID: int
    time: _time_pb2.Property
    velo: _velocity_pb2.Property
    structure: _structure_pb2.Property
    def __init__(self, ID: _Optional[int] = ..., time: _Optional[_Union[_time_pb2.Property, _Mapping]] = ..., velo: _Optional[_Union[_velocity_pb2.Property, _Mapping]] = ..., structure: _Optional[_Union[_structure_pb2.Property, _Mapping]] = ...) -> None: ...

class Space(_message.Message):
    __slots__ = ["entity", "Grid"]
    class EntityEntry(_message.Message):
        __slots__ = ["key", "value"]
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: int
        value: Entity
        def __init__(self, key: _Optional[int] = ..., value: _Optional[_Union[Entity, _Mapping]] = ...) -> None: ...
    class GridEntry(_message.Message):
        __slots__ = ["key", "value"]
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: Chunk
        def __init__(self, key: _Optional[str] = ..., value: _Optional[_Union[Chunk, _Mapping]] = ...) -> None: ...
    ENTITY_FIELD_NUMBER: _ClassVar[int]
    GRID_FIELD_NUMBER: _ClassVar[int]
    entity: _containers.MessageMap[int, Entity]
    Grid: _containers.MessageMap[str, Chunk]
    def __init__(self, entity: _Optional[_Mapping[int, Entity]] = ..., Grid: _Optional[_Mapping[str, Chunk]] = ...) -> None: ...

class Chunk(_message.Message):
    __slots__ = ["entity_index"]
    ENTITY_INDEX_FIELD_NUMBER: _ClassVar[int]
    entity_index: _containers.RepeatedScalarFieldContainer[int]
    def __init__(self, entity_index: _Optional[_Iterable[int]] = ...) -> None: ...
