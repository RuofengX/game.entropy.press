from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Fragment(_message.Message):
    __slots__ = ["delta", "x", "y", "x_v", "y_v"]
    DELTA_FIELD_NUMBER: _ClassVar[int]
    X_FIELD_NUMBER: _ClassVar[int]
    Y_FIELD_NUMBER: _ClassVar[int]
    X_V_FIELD_NUMBER: _ClassVar[int]
    Y_V_FIELD_NUMBER: _ClassVar[int]
    delta: Delta
    x: float
    y: float
    x_v: float
    y_v: float
    def __init__(self, delta: _Optional[_Union[Delta, _Mapping]] = ..., x: _Optional[float] = ..., y: _Optional[float] = ..., x_v: _Optional[float] = ..., y_v: _Optional[float] = ...) -> None: ...

class Delta(_message.Message):
    __slots__ = ["x_a", "y_a"]
    X_A_FIELD_NUMBER: _ClassVar[int]
    Y_A_FIELD_NUMBER: _ClassVar[int]
    x_a: float
    y_a: float
    def __init__(self, x_a: _Optional[float] = ..., y_a: _Optional[float] = ...) -> None: ...
