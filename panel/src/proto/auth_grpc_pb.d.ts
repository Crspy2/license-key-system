// package: protofiles
// file: auth.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import * as auth_pb from "./auth_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";
import * as staff_pb from "./staff_pb";
import * as globals_pb from "./globals_pb";

interface IAuthService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    login: IAuthService_ILogin;
    register: IAuthService_IRegister;
    logout: IAuthService_ILogout;
    getSessionInfo: IAuthService_IGetSessionInfo;
    getUserSessionsStream: IAuthService_IGetUserSessionsStream;
    revokeSession: IAuthService_IRevokeSession;
}

interface IAuthService_ILogin extends grpc.MethodDefinition<auth_pb.LoginRequest, auth_pb.LoginResponse> {
    path: "/protofiles.Auth/Login";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<auth_pb.LoginRequest>;
    requestDeserialize: grpc.deserialize<auth_pb.LoginRequest>;
    responseSerialize: grpc.serialize<auth_pb.LoginResponse>;
    responseDeserialize: grpc.deserialize<auth_pb.LoginResponse>;
}
interface IAuthService_IRegister extends grpc.MethodDefinition<auth_pb.RegisterRequest, globals_pb.StandardResponse> {
    path: "/protofiles.Auth/Register";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<auth_pb.RegisterRequest>;
    requestDeserialize: grpc.deserialize<auth_pb.RegisterRequest>;
    responseSerialize: grpc.serialize<globals_pb.StandardResponse>;
    responseDeserialize: grpc.deserialize<globals_pb.StandardResponse>;
}
interface IAuthService_ILogout extends grpc.MethodDefinition<google_protobuf_empty_pb.Empty, globals_pb.StandardResponse> {
    path: "/protofiles.Auth/Logout";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<google_protobuf_empty_pb.Empty>;
    requestDeserialize: grpc.deserialize<google_protobuf_empty_pb.Empty>;
    responseSerialize: grpc.serialize<globals_pb.StandardResponse>;
    responseDeserialize: grpc.deserialize<globals_pb.StandardResponse>;
}
interface IAuthService_IGetSessionInfo extends grpc.MethodDefinition<google_protobuf_empty_pb.Empty, auth_pb.SingleSessionResponse> {
    path: "/protofiles.Auth/GetSessionInfo";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<google_protobuf_empty_pb.Empty>;
    requestDeserialize: grpc.deserialize<google_protobuf_empty_pb.Empty>;
    responseSerialize: grpc.serialize<auth_pb.SingleSessionResponse>;
    responseDeserialize: grpc.deserialize<auth_pb.SingleSessionResponse>;
}
interface IAuthService_IGetUserSessionsStream extends grpc.MethodDefinition<auth_pb.MultiSessionRequest, auth_pb.SessionObject> {
    path: "/protofiles.Auth/GetUserSessionsStream";
    requestStream: false;
    responseStream: true;
    requestSerialize: grpc.serialize<auth_pb.MultiSessionRequest>;
    requestDeserialize: grpc.deserialize<auth_pb.MultiSessionRequest>;
    responseSerialize: grpc.serialize<auth_pb.SessionObject>;
    responseDeserialize: grpc.deserialize<auth_pb.SessionObject>;
}
interface IAuthService_IRevokeSession extends grpc.MethodDefinition<auth_pb.SessionRevokeRequest, globals_pb.StandardResponse> {
    path: "/protofiles.Auth/RevokeSession";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<auth_pb.SessionRevokeRequest>;
    requestDeserialize: grpc.deserialize<auth_pb.SessionRevokeRequest>;
    responseSerialize: grpc.serialize<globals_pb.StandardResponse>;
    responseDeserialize: grpc.deserialize<globals_pb.StandardResponse>;
}

export const AuthService: IAuthService;

export interface IAuthServer extends grpc.UntypedServiceImplementation {
    login: grpc.handleUnaryCall<auth_pb.LoginRequest, auth_pb.LoginResponse>;
    register: grpc.handleUnaryCall<auth_pb.RegisterRequest, globals_pb.StandardResponse>;
    logout: grpc.handleUnaryCall<google_protobuf_empty_pb.Empty, globals_pb.StandardResponse>;
    getSessionInfo: grpc.handleUnaryCall<google_protobuf_empty_pb.Empty, auth_pb.SingleSessionResponse>;
    getUserSessionsStream: grpc.handleServerStreamingCall<auth_pb.MultiSessionRequest, auth_pb.SessionObject>;
    revokeSession: grpc.handleUnaryCall<auth_pb.SessionRevokeRequest, globals_pb.StandardResponse>;
}

export interface IAuthClient {
    login(request: auth_pb.LoginRequest, callback: (error: grpc.ServiceError | null, response: auth_pb.LoginResponse) => void): grpc.ClientUnaryCall;
    login(request: auth_pb.LoginRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: auth_pb.LoginResponse) => void): grpc.ClientUnaryCall;
    login(request: auth_pb.LoginRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: auth_pb.LoginResponse) => void): grpc.ClientUnaryCall;
    register(request: auth_pb.RegisterRequest, callback: (error: grpc.ServiceError | null, response: globals_pb.StandardResponse) => void): grpc.ClientUnaryCall;
    register(request: auth_pb.RegisterRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: globals_pb.StandardResponse) => void): grpc.ClientUnaryCall;
    register(request: auth_pb.RegisterRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: globals_pb.StandardResponse) => void): grpc.ClientUnaryCall;
    logout(request: google_protobuf_empty_pb.Empty, callback: (error: grpc.ServiceError | null, response: globals_pb.StandardResponse) => void): grpc.ClientUnaryCall;
    logout(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: globals_pb.StandardResponse) => void): grpc.ClientUnaryCall;
    logout(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: globals_pb.StandardResponse) => void): grpc.ClientUnaryCall;
    getSessionInfo(request: google_protobuf_empty_pb.Empty, callback: (error: grpc.ServiceError | null, response: auth_pb.SingleSessionResponse) => void): grpc.ClientUnaryCall;
    getSessionInfo(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: auth_pb.SingleSessionResponse) => void): grpc.ClientUnaryCall;
    getSessionInfo(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: auth_pb.SingleSessionResponse) => void): grpc.ClientUnaryCall;
    getUserSessionsStream(request: auth_pb.MultiSessionRequest, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<auth_pb.SessionObject>;
    getUserSessionsStream(request: auth_pb.MultiSessionRequest, metadata?: grpc.Metadata, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<auth_pb.SessionObject>;
    revokeSession(request: auth_pb.SessionRevokeRequest, callback: (error: grpc.ServiceError | null, response: globals_pb.StandardResponse) => void): grpc.ClientUnaryCall;
    revokeSession(request: auth_pb.SessionRevokeRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: globals_pb.StandardResponse) => void): grpc.ClientUnaryCall;
    revokeSession(request: auth_pb.SessionRevokeRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: globals_pb.StandardResponse) => void): grpc.ClientUnaryCall;
}

export class AuthClient extends grpc.Client implements IAuthClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: Partial<grpc.ClientOptions>);
    public login(request: auth_pb.LoginRequest, callback: (error: grpc.ServiceError | null, response: auth_pb.LoginResponse) => void): grpc.ClientUnaryCall;
    public login(request: auth_pb.LoginRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: auth_pb.LoginResponse) => void): grpc.ClientUnaryCall;
    public login(request: auth_pb.LoginRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: auth_pb.LoginResponse) => void): grpc.ClientUnaryCall;
    public register(request: auth_pb.RegisterRequest, callback: (error: grpc.ServiceError | null, response: globals_pb.StandardResponse) => void): grpc.ClientUnaryCall;
    public register(request: auth_pb.RegisterRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: globals_pb.StandardResponse) => void): grpc.ClientUnaryCall;
    public register(request: auth_pb.RegisterRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: globals_pb.StandardResponse) => void): grpc.ClientUnaryCall;
    public logout(request: google_protobuf_empty_pb.Empty, callback: (error: grpc.ServiceError | null, response: globals_pb.StandardResponse) => void): grpc.ClientUnaryCall;
    public logout(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: globals_pb.StandardResponse) => void): grpc.ClientUnaryCall;
    public logout(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: globals_pb.StandardResponse) => void): grpc.ClientUnaryCall;
    public getSessionInfo(request: google_protobuf_empty_pb.Empty, callback: (error: grpc.ServiceError | null, response: auth_pb.SingleSessionResponse) => void): grpc.ClientUnaryCall;
    public getSessionInfo(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: auth_pb.SingleSessionResponse) => void): grpc.ClientUnaryCall;
    public getSessionInfo(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: auth_pb.SingleSessionResponse) => void): grpc.ClientUnaryCall;
    public getUserSessionsStream(request: auth_pb.MultiSessionRequest, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<auth_pb.SessionObject>;
    public getUserSessionsStream(request: auth_pb.MultiSessionRequest, metadata?: grpc.Metadata, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<auth_pb.SessionObject>;
    public revokeSession(request: auth_pb.SessionRevokeRequest, callback: (error: grpc.ServiceError | null, response: globals_pb.StandardResponse) => void): grpc.ClientUnaryCall;
    public revokeSession(request: auth_pb.SessionRevokeRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: globals_pb.StandardResponse) => void): grpc.ClientUnaryCall;
    public revokeSession(request: auth_pb.SessionRevokeRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: globals_pb.StandardResponse) => void): grpc.ClientUnaryCall;
}
