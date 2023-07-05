import json
from collections import UserDict
from typing import overload

from essence import time_pb2
from essence.base_pb2 import Entity


class Space(UserDict[int, Entity]):
    """宇宙的一帧，
    类似Field的实现
    """

    def __init__(
        self,
    ) -> None:
        """初始的一帧
        self: <实体.ID-实体>的映射
        """
        super().__init__()
        # 第一推动
        self.new_ent(
            Entity(
                ID=0,
                time=time_pb2.Fragment(delta=time_pb2.Delta(time_a=1)),  # 第零个  # 第一推动
            ),
        )
        self.all_field = dict(self[0].ListFields())

    @property
    def age(self) -> int:
        """空间的age参照第一个实体的age来计算

        Returns:
            int: _description_
        """
        return self[0].time.age

    @overload
    def new_ent(self, ent: Entity) -> Entity:
        ...

    @overload
    def new_ent(self, ent: Entity, id: int) -> Entity:
        ...

    def new_ent(self, ent: Entity, id: int | None = None) -> Entity:
        if id is None:
            id = ent.ID
        if ent.ID in self:
            raise KeyError("该实体的ID已经存在")
        self[id] = ent
        return ent

    def tick(self):
        # TODO:
        ...

    def to_json(self):
        """序列化"""
        rtn = json.dumps(self)
        return rtn

    def save(self):
        """保存到绑定的记录中"""
        ...

    # def _sepreate(self) -> Iterable[Field]:
    #     """将自身包含的实体拆解为多个场
    #     返回一个包含了这些场的迭代器
    #     """
    #     f_list: list[Field] = []
    #     for i, ent in self.items():
    #         ent.ListFields()


class Continuum:
    # TODO:
    ...

    def run(self):
        ...

    def save(self):
        ...

    # 其他表象系统定义在这里

    def _run_from(self, tick: int):
        ...

    def _merge_two(self, tick, other_tick):
        ...

    def _merge_many(self, *tick):
        ...


if __name__ == "__main__":
    s = Space()
    print(s.all_field)
