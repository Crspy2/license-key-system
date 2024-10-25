// package: protofiles
// file: auth.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as staff_pb from "./staff_pb";

export class SessionObject extends jspb.Message { 
    getId(): string;
    setId(value: string): SessionObject;
    getIpaddress(): string;
    setIpaddress(value: string): SessionObject;
    getUseragent(): string;
    setUseragent(value: string): SessionObject;

    hasStaff(): boolean;
    clearStaff(): void;
    getStaff(): staff_pb.StaffObject | undefined;
    setStaff(value?: staff_pb.StaffObject): SessionObject;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): SessionObject.AsObject;
    static toObject(includeInstance: boolean, msg: SessionObject): SessionObject.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: SessionObject, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): SessionObject;
    static deserializeBinaryFromReader(message: SessionObject, reader: jspb.BinaryReader): SessionObject;
}

export namespace SessionObject {
    export type AsObject = {
        id: string,
        ipaddress: string,
        useragent: string,
        staff?: staff_pb.StaffObject.AsObject,
    }
}

export class StandardResponse extends jspb.Message { 
    getMessage(): string;
    setMessage(value: string): StandardResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): StandardResponse.AsObject;
    static toObject(includeInstance: boolean, msg: StandardResponse): StandardResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: StandardResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): StandardResponse;
    static deserializeBinaryFromReader(message: StandardResponse, reader: jspb.BinaryReader): StandardResponse;
}

export namespace StandardResponse {
    export type AsObject = {
        message: string,
    }
}

export class LoginRequest extends jspb.Message { 
    getUsername(): string;
    setUsername(value: string): LoginRequest;
    getPassword(): string;
    setPassword(value: string): LoginRequest;
    getIp(): string;
    setIp(value: string): LoginRequest;
    getUserAgent(): string;
    setUserAgent(value: string): LoginRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): LoginRequest.AsObject;
    static toObject(includeInstance: boolean, msg: LoginRequest): LoginRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: LoginRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): LoginRequest;
    static deserializeBinaryFromReader(message: LoginRequest, reader: jspb.BinaryReader): LoginRequest;
}

export namespace LoginRequest {
    export type AsObject = {
        username: string,
        password: string,
        ip: string,
        userAgent: string,
    }
}

export class RegisterRequest extends jspb.Message { 
    getUsername(): string;
    setUsername(value: string): RegisterRequest;
    getPassword(): string;
    setPassword(value: string): RegisterRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): RegisterRequest.AsObject;
    static toObject(includeInstance: boolean, msg: RegisterRequest): RegisterRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: RegisterRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): RegisterRequest;
    static deserializeBinaryFromReader(message: RegisterRequest, reader: jspb.BinaryReader): RegisterRequest;
}

export namespace RegisterRequest {
    export type AsObject = {
        username: string,
        password: string,
    }
}

export class LoginResponse extends jspb.Message { 
    getMessage(): string;
    setMessage(value: string): LoginResponse;

    hasData(): boolean;
    clearData(): void;
    getData(): LoginResponse.ResponseData | undefined;
    setData(value?: LoginResponse.ResponseData): LoginResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): LoginResponse.AsObject;
    static toObject(includeInstance: boolean, msg: LoginResponse): LoginResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: LoginResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): LoginResponse;
    static deserializeBinaryFromReader(message: LoginResponse, reader: jspb.BinaryReader): LoginResponse;
}

export namespace LoginResponse {
    export type AsObject = {
        message: string,
        data?: LoginResponse.ResponseData.AsObject,
    }


    export class ResponseData extends jspb.Message { 
        getSessionid(): string;
        setSessionid(value: string): ResponseData;

        serializeBinary(): Uint8Array;
        toObject(includeInstance?: boolean): ResponseData.AsObject;
        static toObject(includeInstance: boolean, msg: ResponseData): ResponseData.AsObject;
        static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
        static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
        static serializeBinaryToWriter(message: ResponseData, writer: jspb.BinaryWriter): void;
        static deserializeBinary(bytes: Uint8Array): ResponseData;
        static deserializeBinaryFromReader(message: ResponseData, reader: jspb.BinaryReader): ResponseData;
    }

    export namespace ResponseData {
        export type AsObject = {
            sessionid: string,
        }
    }

}

export class LogoutRequest extends jspb.Message { 
    getSessionid(): string;
    setSessionid(value: string): LogoutRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): LogoutRequest.AsObject;
    static toObject(includeInstance: boolean, msg: LogoutRequest): LogoutRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: LogoutRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): LogoutRequest;
    static deserializeBinaryFromReader(message: LogoutRequest, reader: jspb.BinaryReader): LogoutRequest;
}

export namespace LogoutRequest {
    export type AsObject = {
        sessionid: string,
    }
}

export class SingleSessionRequest extends jspb.Message { 
    getSessionid(): string;
    setSessionid(value: string): SingleSessionRequest;
    getIp(): string;
    setIp(value: string): SingleSessionRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): SingleSessionRequest.AsObject;
    static toObject(includeInstance: boolean, msg: SingleSessionRequest): SingleSessionRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: SingleSessionRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): SingleSessionRequest;
    static deserializeBinaryFromReader(message: SingleSessionRequest, reader: jspb.BinaryReader): SingleSessionRequest;
}

export namespace SingleSessionRequest {
    export type AsObject = {
        sessionid: string,
        ip: string,
    }
}

export class SingleSessionResponse extends jspb.Message { 
    getMessage(): string;
    setMessage(value: string): SingleSessionResponse;

    hasData(): boolean;
    clearData(): void;
    getData(): SessionObject | undefined;
    setData(value?: SessionObject): SingleSessionResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): SingleSessionResponse.AsObject;
    static toObject(includeInstance: boolean, msg: SingleSessionResponse): SingleSessionResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: SingleSessionResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): SingleSessionResponse;
    static deserializeBinaryFromReader(message: SingleSessionResponse, reader: jspb.BinaryReader): SingleSessionResponse;
}

export namespace SingleSessionResponse {
    export type AsObject = {
        message: string,
        data?: SessionObject.AsObject,
    }
}

export class MultiSessionRequest extends jspb.Message { 
    getStaffid(): string;
    setStaffid(value: string): MultiSessionRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): MultiSessionRequest.AsObject;
    static toObject(includeInstance: boolean, msg: MultiSessionRequest): MultiSessionRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: MultiSessionRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): MultiSessionRequest;
    static deserializeBinaryFromReader(message: MultiSessionRequest, reader: jspb.BinaryReader): MultiSessionRequest;
}

export namespace MultiSessionRequest {
    export type AsObject = {
        staffid: string,
    }
}

export class SessionRevokeRequest extends jspb.Message { 
    getId(): string;
    setId(value: string): SessionRevokeRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): SessionRevokeRequest.AsObject;
    static toObject(includeInstance: boolean, msg: SessionRevokeRequest): SessionRevokeRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: SessionRevokeRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): SessionRevokeRequest;
    static deserializeBinaryFromReader(message: SessionRevokeRequest, reader: jspb.BinaryReader): SessionRevokeRequest;
}

export namespace SessionRevokeRequest {
    export type AsObject = {
        id: string,
    }
}
