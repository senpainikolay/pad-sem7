# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/accident.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x14proto/accident.proto\x12\x05proto\"]\n\x14GetAccidentUserEntry\x12\x11\n\tuser_long\x18\x01 \x01(\x01\x12\x10\n\x08user_lat\x18\x02 \x01(\x01\x12\x12\n\nzoom_index\x18\x03 \x01(\x03\x12\x0c\n\x04\x63ity\x18\x04 \x01(\t\"\xa6\x01\n\x0c\x41\x63\x63identInfo\x12\x15\n\raccident_long\x18\x01 \x01(\x01\x12\x14\n\x0c\x61\x63\x63ident_lat\x18\x02 \x01(\x01\x12*\n\"confirmation_accident_notification\x18\x03 \x01(\x08\x12(\n confirmation_police_notification\x18\x04 \x01(\x08\x12\x13\n\x0b\x63onfirmedBy\x18\x05 \x01(\x03\"A\n\x13GetAccidentResponse\x12*\n\raccident_info\x18\x01 \x03(\x0b\x32\x13.proto.AccidentInfo\"z\n\x11PostAccidentEntry\x12\x15\n\raccident_long\x18\x01 \x01(\x01\x12\x14\n\x0c\x61\x63\x63ident_lat\x18\x02 \x01(\x01\x12\x0c\n\x04\x63ity\x18\x03 \x01(\t\x12\x13\n\x0bstreet_name\x18\x04 \x01(\t\x12\x15\n\rcars_involved\x18\x05 \x01(\x03\"\x7f\n\x14\x43onfirmAccidentEntry\x12\x15\n\raccident_long\x18\x01 \x01(\x01\x12\x14\n\x0c\x61\x63\x63ident_lat\x18\x02 \x01(\x01\x12\x1b\n\x13police_confirmation\x18\x03 \x01(\x08\x12\x1d\n\x15\x61\x63\x63ident_confirmation\x18\x04 \x01(\x08\"-\n\x0fGenericResponse\x12\r\n\x05\x65rror\x18\x01 \x01(\x08\x12\x0b\n\x03msg\x18\x02 \x01(\t\"?\n\x0eHealthResponse\x12\r\n\x05ready\x18\x01 \x01(\x08\x12\x10\n\x08\x64\x61tabase\x18\x02 \x01(\t\x12\x0c\n\x04load\x18\x03 \x01(\x08\"\x0f\n\rHealthRequest\"F\n\x14\x46\x65tchAccidentRequest\x12.\n\tuser_info\x18\x01 \x01(\x0b\x32\x1b.proto.GetAccidentUserEntry\"F\n\x13PostAccidentRequest\x12/\n\raccident_info\x18\x01 \x01(\x0b\x32\x18.proto.PostAccidentEntry\"C\n\x16\x43onfirmAccidentRequest\x12)\n\x04info\x18\x01 \x01(\x0b\x32\x1b.proto.ConfirmAccidentEntry2\xb7\x02\n\x18\x41\x63\x63identReportingService\x12K\n\x0e\x46\x65tchAccidents\x12\x1b.proto.FetchAccidentRequest\x1a\x1a.proto.GetAccidentResponse\"\x00\x12\x44\n\x0cPostAccident\x12\x1a.proto.PostAccidentRequest\x1a\x16.proto.GenericResponse\"\x00\x12J\n\x0f\x43onfirmAccident\x12\x1d.proto.ConfirmAccidentRequest\x1a\x16.proto.GenericResponse\"\x00\x12<\n\x0bHealthCheck\x12\x14.proto.HealthRequest\x1a\x15.proto.HealthResponse\"\x00\x42>Z<github.com/senpainikolay/pad-sem7/accident-reporting-serviceb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'proto.accident_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z<github.com/senpainikolay/pad-sem7/accident-reporting-service'
  _globals['_GETACCIDENTUSERENTRY']._serialized_start=31
  _globals['_GETACCIDENTUSERENTRY']._serialized_end=124
  _globals['_ACCIDENTINFO']._serialized_start=127
  _globals['_ACCIDENTINFO']._serialized_end=293
  _globals['_GETACCIDENTRESPONSE']._serialized_start=295
  _globals['_GETACCIDENTRESPONSE']._serialized_end=360
  _globals['_POSTACCIDENTENTRY']._serialized_start=362
  _globals['_POSTACCIDENTENTRY']._serialized_end=484
  _globals['_CONFIRMACCIDENTENTRY']._serialized_start=486
  _globals['_CONFIRMACCIDENTENTRY']._serialized_end=613
  _globals['_GENERICRESPONSE']._serialized_start=615
  _globals['_GENERICRESPONSE']._serialized_end=660
  _globals['_HEALTHRESPONSE']._serialized_start=662
  _globals['_HEALTHRESPONSE']._serialized_end=725
  _globals['_HEALTHREQUEST']._serialized_start=727
  _globals['_HEALTHREQUEST']._serialized_end=742
  _globals['_FETCHACCIDENTREQUEST']._serialized_start=744
  _globals['_FETCHACCIDENTREQUEST']._serialized_end=814
  _globals['_POSTACCIDENTREQUEST']._serialized_start=816
  _globals['_POSTACCIDENTREQUEST']._serialized_end=886
  _globals['_CONFIRMACCIDENTREQUEST']._serialized_start=888
  _globals['_CONFIRMACCIDENTREQUEST']._serialized_end=955
  _globals['_ACCIDENTREPORTINGSERVICE']._serialized_start=958
  _globals['_ACCIDENTREPORTINGSERVICE']._serialized_end=1269
# @@protoc_insertion_point(module_scope)