// package: protofiles
// file: staff.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import * as staff_pb from "./staff_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";
import * as globals_pb from "./globals_pb";

interface IStaffService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    setStaffAccess: IStaffService_ISetStaffAccess;
    getStaff: IStaffService_IGetStaff;
    getAllStaffStream: IStaffService_IGetAllStaffStream;
    setStaffPermissions: IStaffService_ISetStaffPermissions;
    setStaffRole: IStaffService_ISetStaffRole;
}

interface IStaffService_ISetStaffAccess extends grpc.MethodDefinition<staff_pb.StaffAccessRequest, staff_pb.ApprovalResponse> {
    path: "/protofiles.Staff/SetStaffAccess";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<staff_pb.StaffAccessRequest>;
    requestDeserialize: grpc.deserialize<staff_pb.StaffAccessRequest>;
    responseSerialize: grpc.serialize<staff_pb.ApprovalResponse>;
    responseDeserialize: grpc.deserialize<staff_pb.ApprovalResponse>;
}
interface IStaffService_IGetStaff extends grpc.MethodDefinition<staff_pb.StaffIdRequest, staff_pb.StaffObject> {
    path: "/protofiles.Staff/GetStaff";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<staff_pb.StaffIdRequest>;
    requestDeserialize: grpc.deserialize<staff_pb.StaffIdRequest>;
    responseSerialize: grpc.serialize<staff_pb.StaffObject>;
    responseDeserialize: grpc.deserialize<staff_pb.StaffObject>;
}
interface IStaffService_IGetAllStaffStream extends grpc.MethodDefinition<google_protobuf_empty_pb.Empty, staff_pb.StaffObject> {
    path: "/protofiles.Staff/GetAllStaffStream";
    requestStream: false;
    responseStream: true;
    requestSerialize: grpc.serialize<google_protobuf_empty_pb.Empty>;
    requestDeserialize: grpc.deserialize<google_protobuf_empty_pb.Empty>;
    responseSerialize: grpc.serialize<staff_pb.StaffObject>;
    responseDeserialize: grpc.deserialize<staff_pb.StaffObject>;
}
interface IStaffService_ISetStaffPermissions extends grpc.MethodDefinition<staff_pb.MultiPermissionRequest, globals_pb.StandardResponse> {
    path: "/protofiles.Staff/SetStaffPermissions";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<staff_pb.MultiPermissionRequest>;
    requestDeserialize: grpc.deserialize<staff_pb.MultiPermissionRequest>;
    responseSerialize: grpc.serialize<globals_pb.StandardResponse>;
    responseDeserialize: grpc.deserialize<globals_pb.StandardResponse>;
}
interface IStaffService_ISetStaffRole extends grpc.MethodDefinition<staff_pb.StaffRoleRequest, globals_pb.StandardResponse> {
    path: "/protofiles.Staff/SetStaffRole";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<staff_pb.StaffRoleRequest>;
    requestDeserialize: grpc.deserialize<staff_pb.StaffRoleRequest>;
    responseSerialize: grpc.serialize<globals_pb.StandardResponse>;
    responseDeserialize: grpc.deserialize<globals_pb.StandardResponse>;
}

export const StaffService: IStaffService;

export interface IStaffServer extends grpc.UntypedServiceImplementation {
    setStaffAccess: grpc.handleUnaryCall<staff_pb.StaffAccessRequest, staff_pb.ApprovalResponse>;
    getStaff: grpc.handleUnaryCall<staff_pb.StaffIdRequest, staff_pb.StaffObject>;
    getAllStaffStream: grpc.handleServerStreamingCall<google_protobuf_empty_pb.Empty, staff_pb.StaffObject>;
    setStaffPermissions: grpc.handleUnaryCall<staff_pb.MultiPermissionRequest, globals_pb.StandardResponse>;
    setStaffRole: grpc.handleUnaryCall<staff_pb.StaffRoleRequest, globals_pb.StandardResponse>;
}

export interface IStaffClient {
    setStaffAccess(request: staff_pb.StaffAccessRequest, callback: (error: grpc.ServiceError | null, response: staff_pb.ApprovalResponse) => void): grpc.ClientUnaryCall;
    setStaffAccess(request: staff_pb.StaffAccessRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: staff_pb.ApprovalResponse) => void): grpc.ClientUnaryCall;
    setStaffAccess(request: staff_pb.StaffAccessRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: staff_pb.ApprovalResponse) => void): grpc.ClientUnaryCall;
    getStaff(request: staff_pb.StaffIdRequest, callback: (error: grpc.ServiceError | null, response: staff_pb.StaffObject) => void): grpc.ClientUnaryCall;
    getStaff(request: staff_pb.StaffIdRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: staff_pb.StaffObject) => void): grpc.ClientUnaryCall;
    getStaff(request: staff_pb.StaffIdRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: staff_pb.StaffObject) => void): grpc.ClientUnaryCall;
    getAllStaffStream(request: google_protobuf_empty_pb.Empty, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<staff_pb.StaffObject>;
    getAllStaffStream(request: google_protobuf_empty_pb.Empty, metadata?: grpc.Metadata, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<staff_pb.StaffObject>;
    setStaffPermissions(request: staff_pb.MultiPermissionRequest, callback: (error: grpc.ServiceError | null, response: globals_pb.StandardResponse) => void): grpc.ClientUnaryCall;
    setStaffPermissions(request: staff_pb.MultiPermissionRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: globals_pb.StandardResponse) => void): grpc.ClientUnaryCall;
    setStaffPermissions(request: staff_pb.MultiPermissionRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: globals_pb.StandardResponse) => void): grpc.ClientUnaryCall;
    setStaffRole(request: staff_pb.StaffRoleRequest, callback: (error: grpc.ServiceError | null, response: globals_pb.StandardResponse) => void): grpc.ClientUnaryCall;
    setStaffRole(request: staff_pb.StaffRoleRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: globals_pb.StandardResponse) => void): grpc.ClientUnaryCall;
    setStaffRole(request: staff_pb.StaffRoleRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: globals_pb.StandardResponse) => void): grpc.ClientUnaryCall;
}

export class StaffClient extends grpc.Client implements IStaffClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: Partial<grpc.ClientOptions>);
    public setStaffAccess(request: staff_pb.StaffAccessRequest, callback: (error: grpc.ServiceError | null, response: staff_pb.ApprovalResponse) => void): grpc.ClientUnaryCall;
    public setStaffAccess(request: staff_pb.StaffAccessRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: staff_pb.ApprovalResponse) => void): grpc.ClientUnaryCall;
    public setStaffAccess(request: staff_pb.StaffAccessRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: staff_pb.ApprovalResponse) => void): grpc.ClientUnaryCall;
    public getStaff(request: staff_pb.StaffIdRequest, callback: (error: grpc.ServiceError | null, response: staff_pb.StaffObject) => void): grpc.ClientUnaryCall;
    public getStaff(request: staff_pb.StaffIdRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: staff_pb.StaffObject) => void): grpc.ClientUnaryCall;
    public getStaff(request: staff_pb.StaffIdRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: staff_pb.StaffObject) => void): grpc.ClientUnaryCall;
    public getAllStaffStream(request: google_protobuf_empty_pb.Empty, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<staff_pb.StaffObject>;
    public getAllStaffStream(request: google_protobuf_empty_pb.Empty, metadata?: grpc.Metadata, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<staff_pb.StaffObject>;
    public setStaffPermissions(request: staff_pb.MultiPermissionRequest, callback: (error: grpc.ServiceError | null, response: globals_pb.StandardResponse) => void): grpc.ClientUnaryCall;
    public setStaffPermissions(request: staff_pb.MultiPermissionRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: globals_pb.StandardResponse) => void): grpc.ClientUnaryCall;
    public setStaffPermissions(request: staff_pb.MultiPermissionRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: globals_pb.StandardResponse) => void): grpc.ClientUnaryCall;
    public setStaffRole(request: staff_pb.StaffRoleRequest, callback: (error: grpc.ServiceError | null, response: globals_pb.StandardResponse) => void): grpc.ClientUnaryCall;
    public setStaffRole(request: staff_pb.StaffRoleRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: globals_pb.StandardResponse) => void): grpc.ClientUnaryCall;
    public setStaffRole(request: staff_pb.StaffRoleRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: globals_pb.StandardResponse) => void): grpc.ClientUnaryCall;
}
