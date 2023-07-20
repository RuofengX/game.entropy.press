import asyncio
import json
from typing import overload


from essence import base_pb2
from essence.base_pb2 import Entity

from workload import Parser


class Space:
    """宇宙的一帧"""

    @staticmethod
    def from_pb2(pb: base_pb2.Space) -> "Space":
        return Space(data=pb)

    def to_pb2(self) -> base_pb2.Space:
        return self.data

    def __init__(
        self,
        data: base_pb2.Space | None = None,
    ) -> None:
        """初始的一帧
        self: <实体.ID-实体>的映射
        空间自身没有属性，所有的属性通过其蕴含的实体显现
        """
        if data is None:
            # 第一推动
            data = base_pb2.Space(
                entity={
                    1: Entity(  # 第一推动，也即原初实体
                        ID=1,  # 第一推动是1而不是0，0没有含义
                        time={
                            "age": 0,
                            "speed": 0,
                            "delta": {"time_a": 1},
                        },
                    ),
                }
            )

        self.data = data

    @property
    def age(self) -> int:
        """空间的年龄
        空间的age参照第一个实体(原初实体)的age来计算

        Returns:
            int: 空间的年龄
        """
        return self.data.entity[0].time.age

    @property
    def capacity(self) -> int:
        """容量属性
        返回空间中已经存在的实体个数

        Returns:
            int: 已存在的实体个数
        """
        return len(self.data.entity)

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
        if ent.ID in self.data.entity:
            raise KeyError("该实体的ID已经存在")
        self.data.entity[id] = ent
        return ent

    def to_json(self):
        """序列化"""
        rtn = json.dumps(self)
        return rtn

    def save(self):
        """保存到绑定的记录中"""
        ...


class Continuum:
    # TODO:
    def __init__(self):
        self.age = 0
        self.history: dict[int, Space] = {0: Space()}

        self._parser = Parser()

    def get_latest(self) -> Space:
        return self.history[self.age]

    async def tick(self):

        _space = await self._parser.tick(self.get_latest().to_pb2())
        new_space = Space.from_pb2(pb=_space)

        self.age += 1
        self.history[self.age] = new_space

    async def run(self):
        async with self._parser:
            while 1:
                await self.tick()

    def save(self):
        ...
        # TODO


if __name__ == "__main__":
    ct = Continuum()
    asyncio.run(ct.run())
