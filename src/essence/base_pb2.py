# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: essence/base.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from essence import time_pb2 as essence_dot_time__pb2
from essence import velocity_pb2 as essence_dot_velocity__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x12\x65ssence/base.proto\x12\x04\x62\x61se\x1a\x12\x65ssence/time.proto\x1a\x16\x65ssence/velocity.proto\"V\n\x06\x45ntity\x12\n\n\x02ID\x18\x01 \x01(\x04\x12\x1c\n\x04time\x18\x03 \x01(\x0b\x32\x0e.time.Fragment\x12\x1c\n\x04velo\x18\x04 \x01(\x0b\x32\x0e.velo.FragmentJ\x04\x08\x02\x10\x03\"%\n\x05\x46ield\x12\x1c\n\x06\x65ntity\x18\x01 \x03(\x0b\x32\x0c.base.EntityB\x1bZ\x19jormungandr/v2/proto/baseb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'essence.base_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\031jormungandr/v2/proto/base'
  _globals['_ENTITY']._serialized_start=72
  _globals['_ENTITY']._serialized_end=158
  _globals['_FIELD']._serialized_start=160
  _globals['_FIELD']._serialized_end=197
# @@protoc_insertion_point(module_scope)
