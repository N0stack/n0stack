# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: create_vm_snapshot_request.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='create_vm_snapshot_request.proto',
  package='',
  syntax='proto3',
  serialized_pb=_b('\n create_vm_snapshot_request.proto\"\x19\n\x17\x43reateVMSnapshotRequestb\x06proto3')
)




_CREATEVMSNAPSHOTREQUEST = _descriptor.Descriptor(
  name='CreateVMSnapshotRequest',
  full_name='CreateVMSnapshotRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=36,
  serialized_end=61,
)

DESCRIPTOR.message_types_by_name['CreateVMSnapshotRequest'] = _CREATEVMSNAPSHOTREQUEST
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

CreateVMSnapshotRequest = _reflection.GeneratedProtocolMessageType('CreateVMSnapshotRequest', (_message.Message,), dict(
  DESCRIPTOR = _CREATEVMSNAPSHOTREQUEST,
  __module__ = 'create_vm_snapshot_request_pb2'
  # @@protoc_insertion_point(class_scope:CreateVMSnapshotRequest)
  ))
_sym_db.RegisterMessage(CreateVMSnapshotRequest)


# @@protoc_insertion_point(module_scope)
