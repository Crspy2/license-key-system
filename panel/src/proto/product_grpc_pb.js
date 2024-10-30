// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var product_pb = require('./product_pb.js');
var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js');
var google_protobuf_duration_pb = require('google-protobuf/google/protobuf/duration_pb.js');
var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js');
var staff_pb = require('./staff_pb.js');
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

function serialize_protofiles_FileObject(arg) {
  if (!(arg instanceof product_pb.FileObject)) {
    throw new Error('Expected argument of type protofiles.FileObject');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_protofiles_FileObject(buffer_arg) {
  return product_pb.FileObject.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_protofiles_ProductIdRequest(arg) {
  if (!(arg instanceof product_pb.ProductIdRequest)) {
    throw new Error('Expected argument of type protofiles.ProductIdRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_protofiles_ProductIdRequest(buffer_arg) {
  return product_pb.ProductIdRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_protofiles_ProductObject(arg) {
  if (!(arg instanceof product_pb.ProductObject)) {
    throw new Error('Expected argument of type protofiles.ProductObject');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_protofiles_ProductObject(buffer_arg) {
  return product_pb.ProductObject.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_protofiles_ProductStatusRequest(arg) {
  if (!(arg instanceof product_pb.ProductStatusRequest)) {
    throw new Error('Expected argument of type protofiles.ProductStatusRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_protofiles_ProductStatusRequest(buffer_arg) {
  return product_pb.ProductStatusRequest.deserializeBinary(new Uint8Array(buffer_arg));
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


var ProductService = exports.ProductService = {
  getProduct: {
    path: '/protofiles.Product/GetProduct',
    requestStream: false,
    responseStream: false,
    requestType: product_pb.ProductIdRequest,
    responseType: product_pb.ProductObject,
    requestSerialize: serialize_protofiles_ProductIdRequest,
    requestDeserialize: deserialize_protofiles_ProductIdRequest,
    responseSerialize: serialize_protofiles_ProductObject,
    responseDeserialize: deserialize_protofiles_ProductObject,
  },
  listProductStream: {
    path: '/protofiles.Product/ListProductStream',
    requestStream: false,
    responseStream: true,
    requestType: google_protobuf_empty_pb.Empty,
    responseType: product_pb.ProductObject,
    requestSerialize: serialize_google_protobuf_Empty,
    requestDeserialize: deserialize_google_protobuf_Empty,
    responseSerialize: serialize_protofiles_ProductObject,
    responseDeserialize: deserialize_protofiles_ProductObject,
  },
  createProduct: {
    path: '/protofiles.Product/CreateProduct',
    requestStream: false,
    responseStream: false,
    requestType: product_pb.ProductObject,
    responseType: product_pb.ProductObject,
    requestSerialize: serialize_protofiles_ProductObject,
    requestDeserialize: deserialize_protofiles_ProductObject,
    responseSerialize: serialize_protofiles_ProductObject,
    responseDeserialize: deserialize_protofiles_ProductObject,
  },
  deleteProduct: {
    path: '/protofiles.Product/DeleteProduct',
    requestStream: false,
    responseStream: false,
    requestType: product_pb.ProductIdRequest,
    responseType: globals_pb.StandardResponse,
    requestSerialize: serialize_protofiles_ProductIdRequest,
    requestDeserialize: deserialize_protofiles_ProductIdRequest,
    responseSerialize: serialize_protofiles_StandardResponse,
    responseDeserialize: deserialize_protofiles_StandardResponse,
  },
  //  rpc CompensateProduct(ProductCompRequest) returns (StandardResponse);
setProductStatus: {
    path: '/protofiles.Product/SetProductStatus',
    requestStream: false,
    responseStream: false,
    requestType: product_pb.ProductStatusRequest,
    responseType: globals_pb.StandardResponse,
    requestSerialize: serialize_protofiles_ProductStatusRequest,
    requestDeserialize: deserialize_protofiles_ProductStatusRequest,
    responseSerialize: serialize_protofiles_StandardResponse,
    responseDeserialize: deserialize_protofiles_StandardResponse,
  },
  setProductFile: {
    path: '/protofiles.Product/SetProductFile',
    requestStream: false,
    responseStream: false,
    requestType: product_pb.FileObject,
    responseType: globals_pb.StandardResponse,
    requestSerialize: serialize_protofiles_FileObject,
    requestDeserialize: deserialize_protofiles_FileObject,
    responseSerialize: serialize_protofiles_StandardResponse,
    responseDeserialize: deserialize_protofiles_StandardResponse,
  },
};

exports.ProductClient = grpc.makeGenericClientConstructor(ProductService);
