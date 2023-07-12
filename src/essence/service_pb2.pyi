from essence import base_pb2 as _base_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Request(_message.Message):
    __slots__ = ["iteration", "space"]
    ITERATION_FIELD_NUMBER: _ClassVar[int]
    SPACE_FIELD_NUMBER: _ClassVar[int]
    iteration: int
    space: _base_pb2.Space
    def __init__(self, iteration: _Optional[int] = ..., space: _Optional[_Union[_base_pb2.Space, _Mapping]] = ...) -> None: ...

class Result(_message.Message):
    __slots__ = ["history"]
    HISTORY_FIELD_NUMBER: _ClassVar[int]
    history: _containers.RepeatedCompositeFieldContainer[_base_pb2.Space]
    def __init__(self, history: _Optional[_Iterable[_Union[_base_pb2.Space, _Mapping]]] = ...) -> None: ...
