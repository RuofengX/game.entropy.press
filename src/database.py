from collections import UserDict
from typing import TYPE_CHECKING, MutableSequence

from utils import SingletonMixin

if TYPE_CHECKING:
    from space import Space

History = MutableSequence


class Database(SingletonMixin):
    """用于记录<时间-帧>的对象"""

    def __init__(self) -> None:
        super().__init__()

    def get_tick(self, time_index: int) -> Space:
        raise NotImplementedError

    def set_tick(self, space: Space) -> None:
        raise NotImplementedError


class DictDatabase(Database, UserDict[int, Space]):
    """字典存储"""

    def __init__(self) -> None:
        super().__init__()

    def get_tick(self, time_index: int) -> Space:
        return self[time_index]

    def set_tick(self, space: Space) -> None:
        self[space.time_index] = space
