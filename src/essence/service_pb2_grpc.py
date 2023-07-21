# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from essence import service_pb2 as essence_dot_service__pb2


class ContinuumStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Tick = channel.unary_unary(
                '/continuum.Continuum/Tick',
                request_serializer=essence_dot_service__pb2.TickRequest.SerializeToString,
                response_deserializer=essence_dot_service__pb2.HistoryResult.FromString,
                )


class ContinuumServicer(object):
    """Missing associated documentation comment in .proto file."""

    def Tick(self, request, context):
        """单个Tick，一次性发送/返回
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_ContinuumServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'Tick': grpc.unary_unary_rpc_method_handler(
                    servicer.Tick,
                    request_deserializer=essence_dot_service__pb2.TickRequest.FromString,
                    response_serializer=essence_dot_service__pb2.HistoryResult.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'continuum.Continuum', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class Continuum(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def Tick(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/continuum.Continuum/Tick',
            essence_dot_service__pb2.TickRequest.SerializeToString,
            essence_dot_service__pb2.HistoryResult.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
