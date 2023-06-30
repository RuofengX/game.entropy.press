import grpc
from proto import service_pb2_grpc

channel: grpc.Channel = grpc.insecure_channel(
    target="localhost:50051",  # 此处可连接serverless驱动的golang服务集群
)
stub = service_pb2_grpc.ContinuumStub(channel)  # grpc调用需要用到stub
