# TODO: 需要重写
import json

from database import DictDatabase
from proto import velocity_pb2 as velo


class Continuum:
    # TODO:
    ...


class Space(velo.Space):
    """宇宙的一帧"""

    recorder = DictDatabase()

    def __init__(
        self,
        time_index: int,
    ) -> None:
        """初始的一帧

        self.entity: 实体列表
        """
        self.time_index = time_index
        super().__init__()

    def tick(self):
        # TODO:
        ...

    def to_json(self):
        """序列化"""
        rtn = json.dumps(self.entity)
        return rtn

    def save(self):
        """保存到绑定的记录中"""
        self.recorder.set_tick(self)
