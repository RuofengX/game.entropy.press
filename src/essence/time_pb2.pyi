"""
@generated by mypy-protobuf.  Do not edit manually!
isort:skip_file
"""
import builtins
import google.protobuf.descriptor
import google.protobuf.message
import sys

if sys.version_info >= (3, 8):
    import typing as typing_extensions
else:
    import typing_extensions

DESCRIPTOR: google.protobuf.descriptor.FileDescriptor

@typing_extensions.final
class Fragment(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    DELTA_FIELD_NUMBER: builtins.int
    AGE_FIELD_NUMBER: builtins.int
    SPEED_FIELD_NUMBER: builtins.int
    @property
    def delta(self) -> global___Delta: ...
    age: builtins.int
    speed: builtins.int
    def __init__(
        self,
        *,
        delta: global___Delta | None = ...,
        age: builtins.int = ...,
        speed: builtins.int = ...,
    ) -> None: ...
    def HasField(self, field_name: typing_extensions.Literal["delta", b"delta"]) -> builtins.bool: ...
    def ClearField(self, field_name: typing_extensions.Literal["age", b"age", "delta", b"delta", "speed", b"speed"]) -> None: ...

global___Fragment = Fragment

@typing_extensions.final
class Delta(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    TIME_A_FIELD_NUMBER: builtins.int
    time_a: builtins.int
    def __init__(
        self,
        *,
        time_a: builtins.int = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["time_a", b"time_a"]) -> None: ...

global___Delta = Delta
