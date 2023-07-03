import grpc

from proto import service_pb2_grpc as pb2_grpc
from proto import velocity_pb2 as velo

tron = velo.Space()


channel: grpc.Channel = grpc.insecure_channel(
    target="localhost:8089",  # 此处可连接serverless驱动的golang服务集群
)
stub = pb2_grpc.ContinuumStub(channel)  # grpc调用需要用到stub
