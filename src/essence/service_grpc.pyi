"""
@generated by mypy-protobuf.  Do not edit manually!
isort:skip_file
"""
import abc
import collections.abc
import essence.service_pb2
import grpc
import grpc.aio
import typing

_T = typing.TypeVar('_T')

class _MaybeAsyncIterator(collections.abc.AsyncIterator[_T], collections.abc.Iterator[_T], metaclass=abc.ABCMeta):
    ...

class _ServicerContext(grpc.ServicerContext, grpc.aio.ServicerContext):  # type: ignore
    ...

class ContinuumStub:
    def __init__(self, channel: typing.Union[grpc.Channel, grpc.aio.Channel]) -> None: ...
    TimePass: grpc.UnaryUnaryMultiCallable[
        essence.service_pb2.Request,
        essence.service_pb2.Result,
    ]
    VelocityMove: grpc.UnaryUnaryMultiCallable[
        essence.service_pb2.Request,
        essence.service_pb2.Result,
    ]
    """TODO: 使用stream计数，减少返回延迟"""

class ContinuumAsyncStub:
    TimePass: grpc.aio.UnaryUnaryMultiCallable[
        essence.service_pb2.Request,
        essence.service_pb2.Result,
    ]
    VelocityMove: grpc.aio.UnaryUnaryMultiCallable[
        essence.service_pb2.Request,
        essence.service_pb2.Result,
    ]
    """TODO: 使用stream计数，减少返回延迟"""

class ContinuumServicer(metaclass=abc.ABCMeta):
    @abc.abstractmethod
    def TimePass(
        self,
        request: essence.service_pb2.Request,
        context: _ServicerContext,
    ) -> typing.Union[essence.service_pb2.Result, collections.abc.Awaitable[essence.service_pb2.Result]]: ...
    @abc.abstractmethod
    def VelocityMove(
        self,
        request: essence.service_pb2.Request,
        context: _ServicerContext,
    ) -> typing.Union[essence.service_pb2.Result, collections.abc.Awaitable[essence.service_pb2.Result]]:
        """TODO: 使用stream计数，减少返回延迟"""

def add_ContinuumServicer_to_server(servicer: ContinuumServicer, server: typing.Union[grpc.Server, grpc.aio.Server]) -> None: ...