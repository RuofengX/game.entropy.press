from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class Time(_message.Message):
    __slots__ = ["age", "speed"]
    AGE_FIELD_NUMBER: _ClassVar[int]
    SPEED_FIELD_NUMBER: _ClassVar[int]
    age: int
    speed: int
    def __init__(self, age: _Optional[int] = ..., speed: _Optional[int] = ...) -> None: ...
