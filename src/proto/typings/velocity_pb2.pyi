from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import (
    ClassVar as _ClassVar,
    Iterable as _Iterable,
    Mapping as _Mapping,
    Optional as _Optional,
    Union as _Union,
)

DESCRIPTOR: _descriptor.FileDescriptor

class Vector(_message.Message):
    __slots__ = ["x", "y", "x_v", "y_v", "x_a", "y_a"]
    X_FIELD_NUMBER: _ClassVar[int]
    Y_FIELD_NUMBER: _ClassVar[int]
    X_V_FIELD_NUMBER: _ClassVar[int]
    Y_V_FIELD_NUMBER: _ClassVar[int]
    X_A_FIELD_NUMBER: _ClassVar[int]
    Y_A_FIELD_NUMBER: _ClassVar[int]
    x: float
    y: float
    x_v: float
    y_v: float
    x_a: float
    y_a: float
    def __init__(
        self,
        x: _Optional[float] = ...,
        y: _Optional[float] = ...,
        x_v: _Optional[float] = ...,
        y_v: _Optional[float] = ...,
        x_a: _Optional[float] = ...,
        y_a: _Optional[float] = ...,
    ) -> None: ...

class VectorList(_message.Message):
    __slots__ = ["vectors"]
    VECTORS_FIELD_NUMBER: _ClassVar[int]
    vectors: _containers.RepeatedCompositeFieldContainer[Vector]
    def __init__(
        self, vectors: _Optional[_Iterable[_Union[Vector, _Mapping]]] = ...
    ) -> None: ...

class Request(_message.Message):
    __slots__ = ["nest_tick", "space"]
    NEST_TICK_FIELD_NUMBER: _ClassVar[int]
    SPACE_FIELD_NUMBER: _ClassVar[int]
    nest_tick: int
    space: VectorList
    def __init__(
        self,
        nest_tick: _Optional[int] = ...,
        space: _Optional[_Union[VectorList, _Mapping]] = ...,
    ) -> None: ...

class Result(_message.Message):
    __slots__ = ["history"]
    HISTORY_FIELD_NUMBER: _ClassVar[int]
    history: _containers.RepeatedCompositeFieldContainer[VectorList]
    def __init__(
        self, history: _Optional[_Iterable[_Union[VectorList, _Mapping]]] = ...
    ) -> None: ...
