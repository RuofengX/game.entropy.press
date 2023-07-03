from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class Velocity(_message.Message):
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
    def __init__(self, x: _Optional[float] = ..., y: _Optional[float] = ..., x_v: _Optional[float] = ..., y_v: _Optional[float] = ..., x_a: _Optional[float] = ..., y_a: _Optional[float] = ...) -> None: ...
