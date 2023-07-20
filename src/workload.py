from typing import Iterable, Any
import grpc.aio

from essence import service_grpc
from essence.service_pb2 import Request
from essence.base_pb2 import Space


class Parser:
    def __init__(self, target: str = "127.0.0.1:8089") -> None:
        self.target = target
        self._channel: None | grpc.aio.Channel = None

    async def multi_tick(self, space: Space, epoch: int = 1) -> Iterable[Space]:
        if self._channel is None:
            raise RuntimeError("尚未准备好用于通信的Channel")
        stub: service_grpc.ContinuumAsyncStub
        stub = service_grpc.ContinuumStub(self._channel)  # type:ignore

        return (
            await stub.Tick(  # FIXME: Issue happens here.
                Request(
                    iteration=epoch,
                    space=space,
                )
            )
        ).history

    async def tick(self, space: Space) -> Space:
        return next(iter(await self.multi_tick(space, 1)))

    async def connect(self):
        self._channel = grpc.aio.insecure_channel(
            target=self.target,  # 此处可连接serverless驱动的golang服务集群
        )

    async def disconnect(self):
        if self._channel is not None:
            await self._channel.close(grace=None)
            self._channel = None

    @property
    def is_connected(self):
        return self._channel is not None

    async def __aenter__(self):
        await self.connect()

    async def __aexit__(self, *args: Any):
        if self._channel is not None:
            await self._channel.__aexit__(*args)
