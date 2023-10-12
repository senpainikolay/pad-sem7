# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from proto import accident_pb2


class AccidentReportingServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.FetchAccidents = channel.unary_unary(
                '/proto.AccidentReportingService/FetchAccidents',
                request_serializer=accident_pb2.FetchAccidentRequest.SerializeToString,
                response_deserializer=accident_pb2.GetAccidentResponse.FromString,
                )
        self.PostAccident = channel.unary_unary(
                '/proto.AccidentReportingService/PostAccident',
                request_serializer=accident_pb2.PostAccidentRequest.SerializeToString,
                response_deserializer=accident_pb2.GenericResponse.FromString,
                )
        self.ConfirmAccident = channel.unary_unary(
                '/proto.AccidentReportingService/ConfirmAccident',
                request_serializer=accident_pb2.ConfirmAccidentRequest.SerializeToString,
                response_deserializer=accident_pb2.GenericResponse.FromString,
                )


class AccidentReportingServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def FetchAccidents(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def PostAccident(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ConfirmAccident(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_AccidentReportingServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'FetchAccidents': grpc.unary_unary_rpc_method_handler(
                    servicer.FetchAccidents,
                    request_deserializer=accident_pb2.FetchAccidentRequest.FromString,
                    response_serializer=accident_pb2.GetAccidentResponse.SerializeToString,
            ),
            'PostAccident': grpc.unary_unary_rpc_method_handler(
                    servicer.PostAccident,
                    request_deserializer=accident_pb2.PostAccidentRequest.FromString,
                    response_serializer=accident_pb2.GenericResponse.SerializeToString,
            ),
            'ConfirmAccident': grpc.unary_unary_rpc_method_handler(
                    servicer.ConfirmAccident,
                    request_deserializer=accident_pb2.ConfirmAccidentRequest.FromString,
                    response_serializer=accident_pb2.GenericResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'proto.AccidentReportingService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class AccidentReportingService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def FetchAccidents(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/proto.AccidentReportingService/FetchAccidents',
            accident_pb2.FetchAccidentRequest.SerializeToString,
            accident_pb2.GetAccidentResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def PostAccident(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/proto.AccidentReportingService/PostAccident',
            accident_pb2.PostAccidentRequest.SerializeToString,
            accident_pb2.GenericResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ConfirmAccident(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/proto.AccidentReportingService/ConfirmAccident',
            accident_pb2.ConfirmAccidentRequest.SerializeToString,
            accident_pb2.GenericResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
