import grpc

from essence import service_grpc

channel: grpc.Channel = grpc.insecure_channel(
    target="localhost:8089",  # 此处可连接serverless驱动的golang服务集群
)
stub = service_grpc.ContinuumStub(channel)  # grpc调用需要用到stub


stub.TimePass(
    request=
)