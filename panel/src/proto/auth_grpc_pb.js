// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var auth_pb = require('./auth_pb.js');
var staff_pb = require('./staff_pb.js');

function serialize_protofiles_LoginRequest(arg) {
  if (!(arg instanceof auth_pb.LoginRequest)) {
    throw new Error('Expected argument of type protofiles.LoginRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_protofiles_LoginRequest(buffer_arg) {
  return auth_pb.LoginRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_protofiles_LoginResponse(arg) {
  if (!(arg instanceof auth_pb.LoginResponse)) {
    throw new Error('Expected argument of type protofiles.LoginResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_protofiles_LoginResponse(buffer_arg) {
  return auth_pb.LoginResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_protofiles_LogoutRequest(arg) {
  if (!(arg instanceof auth_pb.LogoutRequest)) {
    throw new Error('Expected argument of type protofiles.LogoutRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_protofiles_LogoutRequest(buffer_arg) {
  return auth_pb.LogoutRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_protofiles_MultiSessionRequest(arg) {
  if (!(arg instanceof auth_pb.MultiSessionRequest)) {
    throw new Error('Expected argument of type protofiles.MultiSessionRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_protofiles_MultiSessionRequest(buffer_arg) {
  return auth_pb.MultiSessionRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_protofiles_RegisterRequest(arg) {
  if (!(arg instanceof auth_pb.RegisterRequest)) {
    throw new Error('Expected argument of type protofiles.RegisterRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_protofiles_RegisterRequest(buffer_arg) {
  return auth_pb.RegisterRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_protofiles_SessionObject(arg) {
  if (!(arg instanceof auth_pb.SessionObject)) {
    throw new Error('Expected argument of type protofiles.SessionObject');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_protofiles_SessionObject(buffer_arg) {
  return auth_pb.SessionObject.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_protofiles_SessionRevokeRequest(arg) {
  if (!(arg instanceof auth_pb.SessionRevokeRequest)) {
    throw new Error('Expected argument of type protofiles.SessionRevokeRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_protofiles_SessionRevokeRequest(buffer_arg) {
  return auth_pb.SessionRevokeRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_protofiles_SingleSessionRequest(arg) {
  if (!(arg instanceof auth_pb.SingleSessionRequest)) {
    throw new Error('Expected argument of type protofiles.SingleSessionRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_protofiles_SingleSessionRequest(buffer_arg) {
  return auth_pb.SingleSessionRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_protofiles_SingleSessionResponse(arg) {
  if (!(arg instanceof auth_pb.SingleSessionResponse)) {
    throw new Error('Expected argument of type protofiles.SingleSessionResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_protofiles_SingleSessionResponse(buffer_arg) {
  return auth_pb.SingleSessionResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_protofiles_StandardResponse(arg) {
  if (!(arg instanceof auth_pb.StandardResponse)) {
    throw new Error('Expected argument of type protofiles.StandardResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_protofiles_StandardResponse(buffer_arg) {
  return auth_pb.StandardResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var AuthService = exports.AuthService = {
  login: {
    path: '/protofiles.Auth/Login',
    requestStream: false,
    responseStream: false,
    requestType: auth_pb.LoginRequest,
    responseType: auth_pb.LoginResponse,
    requestSerialize: serialize_protofiles_LoginRequest,
    requestDeserialize: deserialize_protofiles_LoginRequest,
    responseSerialize: serialize_protofiles_LoginResponse,
    responseDeserialize: deserialize_protofiles_LoginResponse,
  },
  register: {
    path: '/protofiles.Auth/Register',
    requestStream: false,
    responseStream: false,
    requestType: auth_pb.RegisterRequest,
    responseType: auth_pb.StandardResponse,
    requestSerialize: serialize_protofiles_RegisterRequest,
    requestDeserialize: deserialize_protofiles_RegisterRequest,
    responseSerialize: serialize_protofiles_StandardResponse,
    responseDeserialize: deserialize_protofiles_StandardResponse,
  },
  logout: {
    path: '/protofiles.Auth/Logout',
    requestStream: false,
    responseStream: false,
    requestType: auth_pb.LogoutRequest,
    responseType: auth_pb.StandardResponse,
    requestSerialize: serialize_protofiles_LogoutRequest,
    requestDeserialize: deserialize_protofiles_LogoutRequest,
    responseSerialize: serialize_protofiles_StandardResponse,
    responseDeserialize: deserialize_protofiles_StandardResponse,
  },
  getSessionInfo: {
    path: '/protofiles.Auth/GetSessionInfo',
    requestStream: false,
    responseStream: false,
    requestType: auth_pb.SingleSessionRequest,
    responseType: auth_pb.SingleSessionResponse,
    requestSerialize: serialize_protofiles_SingleSessionRequest,
    requestDeserialize: deserialize_protofiles_SingleSessionRequest,
    responseSerialize: serialize_protofiles_SingleSessionResponse,
    responseDeserialize: deserialize_protofiles_SingleSessionResponse,
  },
  getUserSessionsStream: {
    path: '/protofiles.Auth/GetUserSessionsStream',
    requestStream: false,
    responseStream: true,
    requestType: auth_pb.MultiSessionRequest,
    responseType: auth_pb.SessionObject,
    requestSerialize: serialize_protofiles_MultiSessionRequest,
    requestDeserialize: deserialize_protofiles_MultiSessionRequest,
    responseSerialize: serialize_protofiles_SessionObject,
    responseDeserialize: deserialize_protofiles_SessionObject,
  },
  revokeSession: {
    path: '/protofiles.Auth/RevokeSession',
    requestStream: false,
    responseStream: false,
    requestType: auth_pb.SessionRevokeRequest,
    responseType: auth_pb.StandardResponse,
    requestSerialize: serialize_protofiles_SessionRevokeRequest,
    requestDeserialize: deserialize_protofiles_SessionRevokeRequest,
    responseSerialize: serialize_protofiles_StandardResponse,
    responseDeserialize: deserialize_protofiles_StandardResponse,
  },
};

exports.AuthClient = grpc.makeGenericClientConstructor(AuthService);
