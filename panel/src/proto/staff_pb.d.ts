// package: protofiles
// file: staff.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";
import * as globals_pb from "./globals_pb";

export class StaffObject extends jspb.Message { 
    getId(): string;
    setId(value: string): StaffObject;
    getName(): string;
    setName(value: string): StaffObject;

    hasImage(): boolean;
    clearImage(): void;
    getImage(): string | undefined;
    setImage(value: string): StaffObject;
    getPasswordhash(): string;
    setPasswordhash(value: string): StaffObject;
    getRole(): string;
    setRole(value: string): StaffObject;
    clearPermsList(): void;
    getPermsList(): Array<string>;
    setPermsList(value: Array<string>): StaffObject;
    addPerms(value: string, index?: number): string;
    getApproved(): boolean;
    setApproved(value: boolean): StaffObject;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): StaffObject.AsObject;
    static toObject(includeInstance: boolean, msg: StaffObject): StaffObject.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: StaffObject, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): StaffObject;
    static deserializeBinaryFromReader(message: StaffObject, reader: jspb.BinaryReader): StaffObject;
}

export namespace StaffObject {
    export type AsObject = {
        id: string,
        name: string,
        image?: string,
        passwordhash: string,
        role: string,
        permsList: Array<string>,
        approved: boolean,
    }
}

export class StaffIdRequest extends jspb.Message { 
    getStaffid(): string;
    setStaffid(value: string): StaffIdRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): StaffIdRequest.AsObject;
    static toObject(includeInstance: boolean, msg: StaffIdRequest): StaffIdRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: StaffIdRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): StaffIdRequest;
    static deserializeBinaryFromReader(message: StaffIdRequest, reader: jspb.BinaryReader): StaffIdRequest;
}

export namespace StaffIdRequest {
    export type AsObject = {
        staffid: string,
    }
}

export class ApprovalResponse extends jspb.Message { 
    getMessage(): string;
    setMessage(value: string): ApprovalResponse;

    hasStaff(): boolean;
    clearStaff(): void;
    getStaff(): StaffObject | undefined;
    setStaff(value?: StaffObject): ApprovalResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ApprovalResponse.AsObject;
    static toObject(includeInstance: boolean, msg: ApprovalResponse): ApprovalResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ApprovalResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ApprovalResponse;
    static deserializeBinaryFromReader(message: ApprovalResponse, reader: jspb.BinaryReader): ApprovalResponse;
}

export namespace ApprovalResponse {
    export type AsObject = {
        message: string,
        staff?: StaffObject.AsObject,
    }
}

export class MultiPermissionRequest extends jspb.Message { 
    getStaffid(): string;
    setStaffid(value: string): MultiPermissionRequest;
    clearPermissionsList(): void;
    getPermissionsList(): Array<number>;
    setPermissionsList(value: Array<number>): MultiPermissionRequest;
    addPermissions(value: number, index?: number): number;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): MultiPermissionRequest.AsObject;
    static toObject(includeInstance: boolean, msg: MultiPermissionRequest): MultiPermissionRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: MultiPermissionRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): MultiPermissionRequest;
    static deserializeBinaryFromReader(message: MultiPermissionRequest, reader: jspb.BinaryReader): MultiPermissionRequest;
}

export namespace MultiPermissionRequest {
    export type AsObject = {
        staffid: string,
        permissionsList: Array<number>,
    }
}

export class SinglePermissionRequest extends jspb.Message { 
    getStaffid(): string;
    setStaffid(value: string): SinglePermissionRequest;
    getPermission(): number;
    setPermission(value: number): SinglePermissionRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): SinglePermissionRequest.AsObject;
    static toObject(includeInstance: boolean, msg: SinglePermissionRequest): SinglePermissionRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: SinglePermissionRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): SinglePermissionRequest;
    static deserializeBinaryFromReader(message: SinglePermissionRequest, reader: jspb.BinaryReader): SinglePermissionRequest;
}

export namespace SinglePermissionRequest {
    export type AsObject = {
        staffid: string,
        permission: number,
    }
}
