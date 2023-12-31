# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from proto import health_pb2 as proto_dot_health__pb2
from proto import police_pb2 as proto_dot_police__pb2


class PoliceReportingServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.FetchPolice = channel.unary_unary(
                '/proto.PoliceReportingService/FetchPolice',
                request_serializer=proto_dot_police__pb2.FetchPoliceRequest.SerializeToString,
                response_deserializer=proto_dot_police__pb2.GetPoliceResponse.FromString,
                )
        self.PostPolice = channel.unary_unary(
                '/proto.PoliceReportingService/PostPolice',
                request_serializer=proto_dot_police__pb2.PostPoliceRequest.SerializeToString,
                response_deserializer=proto_dot_police__pb2.PoliceResponse.FromString,
                )
        self.ConfirmPolice = channel.unary_unary(
                '/proto.PoliceReportingService/ConfirmPolice',
                request_serializer=proto_dot_police__pb2.ConfirmPoliceRequest.SerializeToString,
                response_deserializer=proto_dot_police__pb2.PoliceResponse.FromString,
                )
        self.HealthCheck = channel.unary_unary(
                '/proto.PoliceReportingService/HealthCheck',
                request_serializer=proto_dot_health__pb2.HealthRequest.SerializeToString,
                response_deserializer=proto_dot_health__pb2.HealthResponse.FromString,
                )


class PoliceReportingServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def FetchPolice(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def PostPolice(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ConfirmPolice(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def HealthCheck(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_PoliceReportingServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'FetchPolice': grpc.unary_unary_rpc_method_handler(
                    servicer.FetchPolice,
                    request_deserializer=proto_dot_police__pb2.FetchPoliceRequest.FromString,
                    response_serializer=proto_dot_police__pb2.GetPoliceResponse.SerializeToString,
            ),
            'PostPolice': grpc.unary_unary_rpc_method_handler(
                    servicer.PostPolice,
                    request_deserializer=proto_dot_police__pb2.PostPoliceRequest.FromString,
                    response_serializer=proto_dot_police__pb2.PoliceResponse.SerializeToString,
            ),
            'ConfirmPolice': grpc.unary_unary_rpc_method_handler(
                    servicer.ConfirmPolice,
                    request_deserializer=proto_dot_police__pb2.ConfirmPoliceRequest.FromString,
                    response_serializer=proto_dot_police__pb2.PoliceResponse.SerializeToString,
            ),
            'HealthCheck': grpc.unary_unary_rpc_method_handler(
                    servicer.HealthCheck,
                    request_deserializer=proto_dot_health__pb2.HealthRequest.FromString,
                    response_serializer=proto_dot_health__pb2.HealthResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'proto.PoliceReportingService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class PoliceReportingService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def FetchPolice(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/proto.PoliceReportingService/FetchPolice',
            proto_dot_police__pb2.FetchPoliceRequest.SerializeToString,
            proto_dot_police__pb2.GetPoliceResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def PostPolice(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/proto.PoliceReportingService/PostPolice',
            proto_dot_police__pb2.PostPoliceRequest.SerializeToString,
            proto_dot_police__pb2.PoliceResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ConfirmPolice(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/proto.PoliceReportingService/ConfirmPolice',
            proto_dot_police__pb2.ConfirmPoliceRequest.SerializeToString,
            proto_dot_police__pb2.PoliceResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def HealthCheck(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/proto.PoliceReportingService/HealthCheck',
            proto_dot_health__pb2.HealthRequest.SerializeToString,
            proto_dot_health__pb2.HealthResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
