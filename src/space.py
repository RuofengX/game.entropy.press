import json
from collections import UserDict
from typing import overload

from essence import time_pb2
from essence.base_pb2 import Entity, Field


class Space(UserDict[int, Entity]):
    """宇宙的一帧"""

    _field_list: list[str] = Entity.DESCRIPTOR.fields_by_name  # 全部场

    def __init__(
        self,
    ) -> None:
        """初始的一帧
        self: <实体.ID-实体>的映射
        空间自身没有属性，所有的属性通过其蕴含的实体显现
        """
        super().__init__()

        # 第一推动
        self.new_ent(
            Entity(  # 第一推动，也即原初实体
                ID=1,  # 第一推动是1而不是0，0没有含义
                time=time_pb2.Fragment(delta=time_pb2.Delta(time_a=1)),
            )
        )

    @property
    def age(self) -> int:
        """空间的年龄
        空间的age参照第一个实体(原初实体)的age来计算

        Returns:
            int: 空间的年龄
        """
        return self[0].time.age

    @property
    def capacity(self) -> int:
        """容量属性
        返回空间中已经存在的实体个数

        Returns:
            int: 已存在的实体个数
        """
        return super().__len__()

    @overload
    def new_ent(self, ent: Entity) -> Entity:
        """新增实体
        使用实体数据结构内的ent.ID作为实体在space实例中的索引号

        Args:
            ent (Entity): 实体对象

        Returns:
            Entity: 成功新增的实体

        Raises:
            KeyError: 当实体的ID已经在space实例中存在时，将会返回KeyError

        """
        ...

    @overload
    def new_ent(self, ent: Entity, id: int) -> Entity:
        """新增实体
        将ent添加到space实例中去，使用形参id替换ent.id属性

        Args:
            ent (Entity): 将被添加的实体对象
            id (int): 将被使用的实体索引号

        Returns:
            Entity: 成功添加的实体对象，其中ID属性可能已被替换

        Raises:
            KeyError: 当提供的id已经在space实例中存在时，将会抛出KeyError
        """
        ...

    def new_ent(self, ent: Entity, id: int | None = None) -> Entity:
        if id is None:
            id = ent.ID
        else:
            ent.ID = id
        if ent.ID in self:
            raise KeyError("该实体的ID已经存在")
        self[id] = ent
        return ent

    """
    async def tick(self) -> Self:
        for field_name, field in self._sepreate().items():
            TODO: 加入异步分发任务
            next_field = await GET_PARSER(field_name, f)
            rtn = _merge(next_fields)
        return rtn
    """

    def to_json(self):
        """序列化"""
        rtn = json.dumps(self)
        return rtn

    def save(self):
        """保存到绑定的记录中"""
        ...

    def _sepreate(self) -> dict[str, Field]:
        """将自身包含的实体拆解为多个场
        返回一个包含了这些场的名字和映射的字典
        """
        rtn: dict[str, Field] = {f: Field() for f in self._field_list}
        for ent in self.values():
            for d, _ in ent.ListFields():
                rtn[d.name].entity.append(ent)
        return rtn


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
    s._sepreate()
