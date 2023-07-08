import grpc

from essence import service_grpc
from essence.service_pb2 import Request, Result
from essence.base_pb2 import Field


class Parser:
    def __init__(self) -> None:
        grpc.aio.insecure_channel()
        channel: grpc.Channel = grpc.insecure_channel(
            target="localhost:8089",  # 此处可连接serverless驱动的golang服务集群
        )
        self.stub = service_grpc.ContinuumAsyncStub()

    async def time_parser(self, tick: int, field: Field):
        result: Result = await self.stub.TimePass(
            request=Request(
                nest_tick=1,
                field=field,
            ),
        )
