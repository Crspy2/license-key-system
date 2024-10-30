// package: protofiles
// file: product.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import * as product_pb from "./product_pb";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";
import * as google_protobuf_duration_pb from "google-protobuf/google/protobuf/duration_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";
import * as staff_pb from "./staff_pb";
import * as globals_pb from "./globals_pb";

interface IProductService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    getProduct: IProductService_IGetProduct;
    listProductStream: IProductService_IListProductStream;
    createProduct: IProductService_ICreateProduct;
    deleteProduct: IProductService_IDeleteProduct;
    setProductStatus: IProductService_ISetProductStatus;
    setProductFile: IProductService_ISetProductFile;
}

interface IProductService_IGetProduct extends grpc.MethodDefinition<product_pb.ProductIdRequest, product_pb.ProductObject> {
    path: "/protofiles.Product/GetProduct";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<product_pb.ProductIdRequest>;
    requestDeserialize: grpc.deserialize<product_pb.ProductIdRequest>;
    responseSerialize: grpc.serialize<product_pb.ProductObject>;
    responseDeserialize: grpc.deserialize<product_pb.ProductObject>;
}
interface IProductService_IListProductStream extends grpc.MethodDefinition<google_protobuf_empty_pb.Empty, product_pb.ProductObject> {
    path: "/protofiles.Product/ListProductStream";
    requestStream: false;
    responseStream: true;
    requestSerialize: grpc.serialize<google_protobuf_empty_pb.Empty>;
    requestDeserialize: grpc.deserialize<google_protobuf_empty_pb.Empty>;
    responseSerialize: grpc.serialize<product_pb.ProductObject>;
    responseDeserialize: grpc.deserialize<product_pb.ProductObject>;
}
interface IProductService_ICreateProduct extends grpc.MethodDefinition<product_pb.ProductObject, product_pb.ProductObject> {
    path: "/protofiles.Product/CreateProduct";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<product_pb.ProductObject>;
    requestDeserialize: grpc.deserialize<product_pb.ProductObject>;
    responseSerialize: grpc.serialize<product_pb.ProductObject>;
    responseDeserialize: grpc.deserialize<product_pb.ProductObject>;
}
interface IProductService_IDeleteProduct extends grpc.MethodDefinition<product_pb.ProductIdRequest, globals_pb.StandardResponse> {
    path: "/protofiles.Product/DeleteProduct";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<product_pb.ProductIdRequest>;
    requestDeserialize: grpc.deserialize<product_pb.ProductIdRequest>;
    responseSerialize: grpc.serialize<globals_pb.StandardResponse>;
    responseDeserialize: grpc.deserialize<globals_pb.StandardResponse>;
}
interface IProductService_ISetProductStatus extends grpc.MethodDefinition<product_pb.ProductStatusRequest, globals_pb.StandardResponse> {
    path: "/protofiles.Product/SetProductStatus";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<product_pb.ProductStatusRequest>;
    requestDeserialize: grpc.deserialize<product_pb.ProductStatusRequest>;
    responseSerialize: grpc.serialize<globals_pb.StandardResponse>;
    responseDeserialize: grpc.deserialize<globals_pb.StandardResponse>;
}
interface IProductService_ISetProductFile extends grpc.MethodDefinition<product_pb.FileObject, globals_pb.StandardResponse> {
    path: "/protofiles.Product/SetProductFile";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<product_pb.FileObject>;
    requestDeserialize: grpc.deserialize<product_pb.FileObject>;
    responseSerialize: grpc.serialize<globals_pb.StandardResponse>;
    responseDeserialize: grpc.deserialize<globals_pb.StandardResponse>;
}

export const ProductService: IProductService;

export interface IProductServer extends grpc.UntypedServiceImplementation {
    getProduct: grpc.handleUnaryCall<product_pb.ProductIdRequest, product_pb.ProductObject>;
    listProductStream: grpc.handleServerStreamingCall<google_protobuf_empty_pb.Empty, product_pb.ProductObject>;
    createProduct: grpc.handleUnaryCall<product_pb.ProductObject, product_pb.ProductObject>;
    deleteProduct: grpc.handleUnaryCall<product_pb.ProductIdRequest, globals_pb.StandardResponse>;
    setProductStatus: grpc.handleUnaryCall<product_pb.ProductStatusRequest, globals_pb.StandardResponse>;
    setProductFile: grpc.handleUnaryCall<product_pb.FileObject, globals_pb.StandardResponse>;
}

export interface IProductClient {
    getProduct(request: product_pb.ProductIdRequest, callback: (error: grpc.ServiceError | null, response: product_pb.ProductObject) => void): grpc.ClientUnaryCall;
    getProduct(request: product_pb.ProductIdRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: product_pb.ProductObject) => void): grpc.ClientUnaryCall;
    getProduct(request: product_pb.ProductIdRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: product_pb.ProductObject) => void): grpc.ClientUnaryCall;
    listProductStream(request: google_protobuf_empty_pb.Empty, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<product_pb.ProductObject>;
    listProductStream(request: google_protobuf_empty_pb.Empty, metadata?: grpc.Metadata, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<product_pb.ProductObject>;
    createProduct(request: product_pb.ProductObject, callback: (error: grpc.ServiceError | null, response: product_pb.ProductObject) => void): grpc.ClientUnaryCall;
    createProduct(request: product_pb.ProductObject, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: product_pb.ProductObject) => void): grpc.ClientUnaryCall;
    createProduct(request: product_pb.ProductObject, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: product_pb.ProductObject) => void): grpc.ClientUnaryCall;
    deleteProduct(request: product_pb.ProductIdRequest, callback: (error: grpc.ServiceError | null, response: globals_pb.StandardResponse) => void): grpc.ClientUnaryCall;
    deleteProduct(request: product_pb.ProductIdRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: globals_pb.StandardResponse) => void): grpc.ClientUnaryCall;
    deleteProduct(request: product_pb.ProductIdRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: globals_pb.StandardResponse) => void): grpc.ClientUnaryCall;
    setProductStatus(request: product_pb.ProductStatusRequest, callback: (error: grpc.ServiceError | null, response: globals_pb.StandardResponse) => void): grpc.ClientUnaryCall;
    setProductStatus(request: product_pb.ProductStatusRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: globals_pb.StandardResponse) => void): grpc.ClientUnaryCall;
    setProductStatus(request: product_pb.ProductStatusRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: globals_pb.StandardResponse) => void): grpc.ClientUnaryCall;
    setProductFile(request: product_pb.FileObject, callback: (error: grpc.ServiceError | null, response: globals_pb.StandardResponse) => void): grpc.ClientUnaryCall;
    setProductFile(request: product_pb.FileObject, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: globals_pb.StandardResponse) => void): grpc.ClientUnaryCall;
    setProductFile(request: product_pb.FileObject, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: globals_pb.StandardResponse) => void): grpc.ClientUnaryCall;
}

export class ProductClient extends grpc.Client implements IProductClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: Partial<grpc.ClientOptions>);
    public getProduct(request: product_pb.ProductIdRequest, callback: (error: grpc.ServiceError | null, response: product_pb.ProductObject) => void): grpc.ClientUnaryCall;
    public getProduct(request: product_pb.ProductIdRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: product_pb.ProductObject) => void): grpc.ClientUnaryCall;
    public getProduct(request: product_pb.ProductIdRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: product_pb.ProductObject) => void): grpc.ClientUnaryCall;
    public listProductStream(request: google_protobuf_empty_pb.Empty, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<product_pb.ProductObject>;
    public listProductStream(request: google_protobuf_empty_pb.Empty, metadata?: grpc.Metadata, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<product_pb.ProductObject>;
    public createProduct(request: product_pb.ProductObject, callback: (error: grpc.ServiceError | null, response: product_pb.ProductObject) => void): grpc.ClientUnaryCall;
    public createProduct(request: product_pb.ProductObject, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: product_pb.ProductObject) => void): grpc.ClientUnaryCall;
    public createProduct(request: product_pb.ProductObject, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: product_pb.ProductObject) => void): grpc.ClientUnaryCall;
    public deleteProduct(request: product_pb.ProductIdRequest, callback: (error: grpc.ServiceError | null, response: globals_pb.StandardResponse) => void): grpc.ClientUnaryCall;
    public deleteProduct(request: product_pb.ProductIdRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: globals_pb.StandardResponse) => void): grpc.ClientUnaryCall;
    public deleteProduct(request: product_pb.ProductIdRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: globals_pb.StandardResponse) => void): grpc.ClientUnaryCall;
    public setProductStatus(request: product_pb.ProductStatusRequest, callback: (error: grpc.ServiceError | null, response: globals_pb.StandardResponse) => void): grpc.ClientUnaryCall;
    public setProductStatus(request: product_pb.ProductStatusRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: globals_pb.StandardResponse) => void): grpc.ClientUnaryCall;
    public setProductStatus(request: product_pb.ProductStatusRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: globals_pb.StandardResponse) => void): grpc.ClientUnaryCall;
    public setProductFile(request: product_pb.FileObject, callback: (error: grpc.ServiceError | null, response: globals_pb.StandardResponse) => void): grpc.ClientUnaryCall;
    public setProductFile(request: product_pb.FileObject, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: globals_pb.StandardResponse) => void): grpc.ClientUnaryCall;
    public setProductFile(request: product_pb.FileObject, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: globals_pb.StandardResponse) => void): grpc.ClientUnaryCall;
}
