from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Fragment(_message.Message):
    __slots__ = ["delta", "age", "speed"]
    DELTA_FIELD_NUMBER: _ClassVar[int]
    AGE_FIELD_NUMBER: _ClassVar[int]
    SPEED_FIELD_NUMBER: _ClassVar[int]
    delta: Delta
    age: int
    speed: int
    def __init__(self, delta: _Optional[_Union[Delta, _Mapping]] = ..., age: _Optional[int] = ..., speed: _Optional[int] = ...) -> None: ...

class Delta(_message.Message):
    __slots__ = ["time_a"]
    TIME_A_FIELD_NUMBER: _ClassVar[int]
    time_a: int
    def __init__(self, time_a: _Optional[int] = ...) -> None: ...
