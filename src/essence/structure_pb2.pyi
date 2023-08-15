from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Property(_message.Message):
    __slots__ = ["delta", "destroy", "health", "injury"]
    DELTA_FIELD_NUMBER: _ClassVar[int]
    DESTROY_FIELD_NUMBER: _ClassVar[int]
    HEALTH_FIELD_NUMBER: _ClassVar[int]
    INJURY_FIELD_NUMBER: _ClassVar[int]
    delta: Delta
    destroy: bool
    health: Health
    injury: Injury
    def __init__(self, delta: _Optional[_Union[Delta, _Mapping]] = ..., destroy: bool = ..., health: _Optional[_Union[Health, _Mapping]] = ..., injury: _Optional[_Union[Injury, _Mapping]] = ...) -> None: ...

class Health(_message.Message):
    __slots__ = ["body", "armor", "shield"]
    BODY_FIELD_NUMBER: _ClassVar[int]
    ARMOR_FIELD_NUMBER: _ClassVar[int]
    SHIELD_FIELD_NUMBER: _ClassVar[int]
    body: float
    armor: float
    shield: float
    def __init__(self, body: _Optional[float] = ..., armor: _Optional[float] = ..., shield: _Optional[float] = ...) -> None: ...

class Injury(_message.Message):
    __slots__ = ["injurable", "damage", "radius", "direct"]
    INJURABLE_FIELD_NUMBER: _ClassVar[int]
    DAMAGE_FIELD_NUMBER: _ClassVar[int]
    RADIUS_FIELD_NUMBER: _ClassVar[int]
    DIRECT_FIELD_NUMBER: _ClassVar[int]
    injurable: bool
    damage: Health
    radius: float
    direct: int
    def __init__(self, injurable: bool = ..., damage: _Optional[_Union[Health, _Mapping]] = ..., radius: _Optional[float] = ..., direct: _Optional[int] = ...) -> None: ...

class Delta(_message.Message):
    __slots__ = ["health_a"]
    HEALTH_A_FIELD_NUMBER: _ClassVar[int]
    health_a: Health
    def __init__(self, health_a: _Optional[_Union[Health, _Mapping]] = ...) -> None: ...
