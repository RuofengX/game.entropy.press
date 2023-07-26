from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Property(_message.Message):
    __slots__ = ["delta", "destroy", "health", "structure", "shield", "shield_recovery"]
    DELTA_FIELD_NUMBER: _ClassVar[int]
    DESTROY_FIELD_NUMBER: _ClassVar[int]
    HEALTH_FIELD_NUMBER: _ClassVar[int]
    STRUCTURE_FIELD_NUMBER: _ClassVar[int]
    SHIELD_FIELD_NUMBER: _ClassVar[int]
    SHIELD_RECOVERY_FIELD_NUMBER: _ClassVar[int]
    delta: Delta
    destroy: bool
    health: float
    structure: float
    shield: float
    shield_recovery: float
    def __init__(self, delta: _Optional[_Union[Delta, _Mapping]] = ..., destroy: bool = ..., health: _Optional[float] = ..., structure: _Optional[float] = ..., shield: _Optional[float] = ..., shield_recovery: _Optional[float] = ...) -> None: ...

class Delta(_message.Message):
    __slots__ = ["health_a", "structure_a", "shield_recovery_a"]
    HEALTH_A_FIELD_NUMBER: _ClassVar[int]
    STRUCTURE_A_FIELD_NUMBER: _ClassVar[int]
    SHIELD_RECOVERY_A_FIELD_NUMBER: _ClassVar[int]
    health_a: float
    structure_a: float
    shield_recovery_a: float
    def __init__(self, health_a: _Optional[float] = ..., structure_a: _Optional[float] = ..., shield_recovery_a: _Optional[float] = ...) -> None: ...
