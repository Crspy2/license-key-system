// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var staff_pb = require('./staff_pb.js');
var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js');
var globals_pb = require('./globals_pb.js');

function serialize_google_protobuf_Empty(arg) {
  if (!(arg instanceof google_protobuf_empty_pb.Empty)) {
    throw new Error('Expected argument of type google.protobuf.Empty');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_google_protobuf_Empty(buffer_arg) {
  return google_protobuf_empty_pb.Empty.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_protofiles_ApprovalResponse(arg) {
  if (!(arg instanceof staff_pb.ApprovalResponse)) {
    throw new Error('Expected argument of type protofiles.ApprovalResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_protofiles_ApprovalResponse(buffer_arg) {
  return staff_pb.ApprovalResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_protofiles_MultiPermissionRequest(arg) {
  if (!(arg instanceof staff_pb.MultiPermissionRequest)) {
    throw new Error('Expected argument of type protofiles.MultiPermissionRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_protofiles_MultiPermissionRequest(buffer_arg) {
  return staff_pb.MultiPermissionRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_protofiles_StaffAccessRequest(arg) {
  if (!(arg instanceof staff_pb.StaffAccessRequest)) {
    throw new Error('Expected argument of type protofiles.StaffAccessRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_protofiles_StaffAccessRequest(buffer_arg) {
  return staff_pb.StaffAccessRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_protofiles_StaffIdRequest(arg) {
  if (!(arg instanceof staff_pb.StaffIdRequest)) {
    throw new Error('Expected argument of type protofiles.StaffIdRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_protofiles_StaffIdRequest(buffer_arg) {
  return staff_pb.StaffIdRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_protofiles_StaffObject(arg) {
  if (!(arg instanceof staff_pb.StaffObject)) {
    throw new Error('Expected argument of type protofiles.StaffObject');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_protofiles_StaffObject(buffer_arg) {
  return staff_pb.StaffObject.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_protofiles_StaffRoleRequest(arg) {
  if (!(arg instanceof staff_pb.StaffRoleRequest)) {
    throw new Error('Expected argument of type protofiles.StaffRoleRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_protofiles_StaffRoleRequest(buffer_arg) {
  return staff_pb.StaffRoleRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_protofiles_StandardResponse(arg) {
  if (!(arg instanceof globals_pb.StandardResponse)) {
    throw new Error('Expected argument of type protofiles.StandardResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_protofiles_StandardResponse(buffer_arg) {
  return globals_pb.StandardResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var StaffService = exports.StaffService = {
  setStaffAccess: {
    path: '/protofiles.Staff/SetStaffAccess',
    requestStream: false,
    responseStream: false,
    requestType: staff_pb.StaffAccessRequest,
    responseType: staff_pb.ApprovalResponse,
    requestSerialize: serialize_protofiles_StaffAccessRequest,
    requestDeserialize: deserialize_protofiles_StaffAccessRequest,
    responseSerialize: serialize_protofiles_ApprovalResponse,
    responseDeserialize: deserialize_protofiles_ApprovalResponse,
  },
  getStaff: {
    path: '/protofiles.Staff/GetStaff',
    requestStream: false,
    responseStream: false,
    requestType: staff_pb.StaffIdRequest,
    responseType: staff_pb.StaffObject,
    requestSerialize: serialize_protofiles_StaffIdRequest,
    requestDeserialize: deserialize_protofiles_StaffIdRequest,
    responseSerialize: serialize_protofiles_StaffObject,
    responseDeserialize: deserialize_protofiles_StaffObject,
  },
  listStaffStream: {
    path: '/protofiles.Staff/ListStaffStream',
    requestStream: false,
    responseStream: true,
    requestType: google_protobuf_empty_pb.Empty,
    responseType: staff_pb.StaffObject,
    requestSerialize: serialize_google_protobuf_Empty,
    requestDeserialize: deserialize_google_protobuf_Empty,
    responseSerialize: serialize_protofiles_StaffObject,
    responseDeserialize: deserialize_protofiles_StaffObject,
  },
  setStaffPermissions: {
    path: '/protofiles.Staff/SetStaffPermissions',
    requestStream: false,
    responseStream: false,
    requestType: staff_pb.MultiPermissionRequest,
    responseType: globals_pb.StandardResponse,
    requestSerialize: serialize_protofiles_MultiPermissionRequest,
    requestDeserialize: deserialize_protofiles_MultiPermissionRequest,
    responseSerialize: serialize_protofiles_StandardResponse,
    responseDeserialize: deserialize_protofiles_StandardResponse,
  },
  setStaffRole: {
    path: '/protofiles.Staff/SetStaffRole',
    requestStream: false,
    responseStream: false,
    requestType: staff_pb.StaffRoleRequest,
    responseType: globals_pb.StandardResponse,
    requestSerialize: serialize_protofiles_StaffRoleRequest,
    requestDeserialize: deserialize_protofiles_StaffRoleRequest,
    responseSerialize: serialize_protofiles_StandardResponse,
    responseDeserialize: deserialize_protofiles_StandardResponse,
  },
};

exports.StaffClient = grpc.makeGenericClientConstructor(StaffService);
