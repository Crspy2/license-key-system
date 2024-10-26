// package: protofiles
// file: staff.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import * as staff_pb from "./staff_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";
import * as globals_pb from "./globals_pb";

interface IStaffService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    approveStaff: IStaffService_IApproveStaff;
    getStaff: IStaffService_IGetStaff;
    getAllStaffStream: IStaffService_IGetAllStaffStream;
    setStaffPermissions: IStaffService_ISetStaffPermissions;
    addStaffPermission: IStaffService_IAddStaffPermission;
    removeStaffPermission: IStaffService_IRemoveStaffPermission;
}

interface IStaffService_IApproveStaff extends grpc.MethodDefinition<staff_pb.StaffId, staff_pb.ApprovalResponse> {
    path: "/protofiles.Staff/ApproveStaff";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<staff_pb.StaffId>;
    requestDeserialize: grpc.deserialize<staff_pb.StaffId>;
    responseSerialize: grpc.serialize<staff_pb.ApprovalResponse>;
    responseDeserialize: grpc.deserialize<staff_pb.ApprovalResponse>;
}
interface IStaffService_IGetStaff extends grpc.MethodDefinition<staff_pb.StaffId, staff_pb.StaffObject> {
    path: "/protofiles.Staff/GetStaff";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<staff_pb.StaffId>;
    requestDeserialize: grpc.deserialize<staff_pb.StaffId>;
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
interface IStaffService_IAddStaffPermission extends grpc.MethodDefinition<staff_pb.SinglePermissionRequest, globals_pb.StandardResponse> {
    path: "/protofiles.Staff/AddStaffPermission";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<staff_pb.SinglePermissionRequest>;
    requestDeserialize: grpc.deserialize<staff_pb.SinglePermissionRequest>;
    responseSerialize: grpc.serialize<globals_pb.StandardResponse>;
    responseDeserialize: grpc.deserialize<globals_pb.StandardResponse>;
}
interface IStaffService_IRemoveStaffPermission extends grpc.MethodDefinition<staff_pb.SinglePermissionRequest, globals_pb.StandardResponse> {
    path: "/protofiles.Staff/RemoveStaffPermission";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<staff_pb.SinglePermissionRequest>;
    requestDeserialize: grpc.deserialize<staff_pb.SinglePermissionRequest>;
    responseSerialize: grpc.serialize<globals_pb.StandardResponse>;
    responseDeserialize: grpc.deserialize<globals_pb.StandardResponse>;
}

export const StaffService: IStaffService;

export interface IStaffServer extends grpc.UntypedServiceImplementation {
    approveStaff: grpc.handleUnaryCall<staff_pb.StaffId, staff_pb.ApprovalResponse>;
    getStaff: grpc.handleUnaryCall<staff_pb.StaffId, staff_pb.StaffObject>;
    getAllStaffStream: grpc.handleServerStreamingCall<google_protobuf_empty_pb.Empty, staff_pb.StaffObject>;
    setStaffPermissions: grpc.handleUnaryCall<staff_pb.MultiPermissionRequest, globals_pb.StandardResponse>;
    addStaffPermission: grpc.handleUnaryCall<staff_pb.SinglePermissionRequest, globals_pb.StandardResponse>;
    removeStaffPermission: grpc.handleUnaryCall<staff_pb.SinglePermissionRequest, globals_pb.StandardResponse>;
}

export interface IStaffClient {
    approveStaff(request: staff_pb.StaffId, callback: (error: grpc.ServiceError | null, response: staff_pb.ApprovalResponse) => void): grpc.ClientUnaryCall;
    approveStaff(request: staff_pb.StaffId, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: staff_pb.ApprovalResponse) => void): grpc.ClientUnaryCall;
    approveStaff(request: staff_pb.StaffId, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: staff_pb.ApprovalResponse) => void): grpc.ClientUnaryCall;
    getStaff(request: staff_pb.StaffId, callback: (error: grpc.ServiceError | null, response: staff_pb.StaffObject) => void): grpc.ClientUnaryCall;
    getStaff(request: staff_pb.StaffId, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: staff_pb.StaffObject) => void): grpc.ClientUnaryCall;
    getStaff(request: staff_pb.StaffId, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: staff_pb.StaffObject) => void): grpc.ClientUnaryCall;
    getAllStaffStream(request: google_protobuf_empty_pb.Empty, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<staff_pb.StaffObject>;
    getAllStaffStream(request: google_protobuf_empty_pb.Empty, metadata?: grpc.Metadata, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<staff_pb.StaffObject>;
    setStaffPermissions(request: staff_pb.MultiPermissionRequest, callback: (error: grpc.ServiceError | null, response: globals_pb.StandardResponse) => void): grpc.ClientUnaryCall;
    setStaffPermissions(request: staff_pb.MultiPermissionRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: globals_pb.StandardResponse) => void): grpc.ClientUnaryCall;
    setStaffPermissions(request: staff_pb.MultiPermissionRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: globals_pb.StandardResponse) => void): grpc.ClientUnaryCall;
    addStaffPermission(request: staff_pb.SinglePermissionRequest, callback: (error: grpc.ServiceError | null, response: globals_pb.StandardResponse) => void): grpc.ClientUnaryCall;
    addStaffPermission(request: staff_pb.SinglePermissionRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: globals_pb.StandardResponse) => void): grpc.ClientUnaryCall;
    addStaffPermission(request: staff_pb.SinglePermissionRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: globals_pb.StandardResponse) => void): grpc.ClientUnaryCall;
    removeStaffPermission(request: staff_pb.SinglePermissionRequest, callback: (error: grpc.ServiceError | null, response: globals_pb.StandardResponse) => void): grpc.ClientUnaryCall;
    removeStaffPermission(request: staff_pb.SinglePermissionRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: globals_pb.StandardResponse) => void): grpc.ClientUnaryCall;
    removeStaffPermission(request: staff_pb.SinglePermissionRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: globals_pb.StandardResponse) => void): grpc.ClientUnaryCall;
}

export class StaffClient extends grpc.Client implements IStaffClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: Partial<grpc.ClientOptions>);
    public approveStaff(request: staff_pb.StaffId, callback: (error: grpc.ServiceError | null, response: staff_pb.ApprovalResponse) => void): grpc.ClientUnaryCall;
    public approveStaff(request: staff_pb.StaffId, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: staff_pb.ApprovalResponse) => void): grpc.ClientUnaryCall;
    public approveStaff(request: staff_pb.StaffId, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: staff_pb.ApprovalResponse) => void): grpc.ClientUnaryCall;
    public getStaff(request: staff_pb.StaffId, callback: (error: grpc.ServiceError | null, response: staff_pb.StaffObject) => void): grpc.ClientUnaryCall;
    public getStaff(request: staff_pb.StaffId, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: staff_pb.StaffObject) => void): grpc.ClientUnaryCall;
    public getStaff(request: staff_pb.StaffId, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: staff_pb.StaffObject) => void): grpc.ClientUnaryCall;
    public getAllStaffStream(request: google_protobuf_empty_pb.Empty, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<staff_pb.StaffObject>;
    public getAllStaffStream(request: google_protobuf_empty_pb.Empty, metadata?: grpc.Metadata, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<staff_pb.StaffObject>;
    public setStaffPermissions(request: staff_pb.MultiPermissionRequest, callback: (error: grpc.ServiceError | null, response: globals_pb.StandardResponse) => void): grpc.ClientUnaryCall;
    public setStaffPermissions(request: staff_pb.MultiPermissionRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: globals_pb.StandardResponse) => void): grpc.ClientUnaryCall;
    public setStaffPermissions(request: staff_pb.MultiPermissionRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: globals_pb.StandardResponse) => void): grpc.ClientUnaryCall;
    public addStaffPermission(request: staff_pb.SinglePermissionRequest, callback: (error: grpc.ServiceError | null, response: globals_pb.StandardResponse) => void): grpc.ClientUnaryCall;
    public addStaffPermission(request: staff_pb.SinglePermissionRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: globals_pb.StandardResponse) => void): grpc.ClientUnaryCall;
    public addStaffPermission(request: staff_pb.SinglePermissionRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: globals_pb.StandardResponse) => void): grpc.ClientUnaryCall;
    public removeStaffPermission(request: staff_pb.SinglePermissionRequest, callback: (error: grpc.ServiceError | null, response: globals_pb.StandardResponse) => void): grpc.ClientUnaryCall;
    public removeStaffPermission(request: staff_pb.SinglePermissionRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: globals_pb.StandardResponse) => void): grpc.ClientUnaryCall;
    public removeStaffPermission(request: staff_pb.SinglePermissionRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: globals_pb.StandardResponse) => void): grpc.ClientUnaryCall;
}
