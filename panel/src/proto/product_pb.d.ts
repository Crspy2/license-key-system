// package: protofiles
// file: product.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";
import * as google_protobuf_duration_pb from "google-protobuf/google/protobuf/duration_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";
import * as staff_pb from "./staff_pb";
import * as globals_pb from "./globals_pb";

export class FileObject extends jspb.Message { 
    getId(): string;
    setId(value: string): FileObject;
    getName(): string;
    setName(value: string): FileObject;
    getPath(): string;
    setPath(value: string): FileObject;

    hasUploader(): boolean;
    clearUploader(): void;
    getUploader(): staff_pb.StaffObject | undefined;
    setUploader(value?: staff_pb.StaffObject): FileObject;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): FileObject.AsObject;
    static toObject(includeInstance: boolean, msg: FileObject): FileObject.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: FileObject, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): FileObject;
    static deserializeBinaryFromReader(message: FileObject, reader: jspb.BinaryReader): FileObject;
}

export namespace FileObject {
    export type AsObject = {
        id: string,
        name: string,
        path: string,
        uploader?: staff_pb.StaffObject.AsObject,
    }
}

export class ProductObject extends jspb.Message { 
    getId(): string;
    setId(value: string): ProductObject;
    getName(): string;
    setName(value: string): ProductObject;
    getStatus(): string;
    setStatus(value: string): ProductObject;

    hasStatuschangedat(): boolean;
    clearStatuschangedat(): void;
    getStatuschangedat(): google_protobuf_timestamp_pb.Timestamp | undefined;
    setStatuschangedat(value?: google_protobuf_timestamp_pb.Timestamp): ProductObject;

    hasLastfile(): boolean;
    clearLastfile(): void;
    getLastfile(): FileObject | undefined;
    setLastfile(value?: FileObject): ProductObject;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ProductObject.AsObject;
    static toObject(includeInstance: boolean, msg: ProductObject): ProductObject.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ProductObject, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ProductObject;
    static deserializeBinaryFromReader(message: ProductObject, reader: jspb.BinaryReader): ProductObject;
}

export namespace ProductObject {
    export type AsObject = {
        id: string,
        name: string,
        status: string,
        statuschangedat?: google_protobuf_timestamp_pb.Timestamp.AsObject,
        lastfile?: FileObject.AsObject,
    }
}

export class ProductIdRequest extends jspb.Message { 
    getProductid(): string;
    setProductid(value: string): ProductIdRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ProductIdRequest.AsObject;
    static toObject(includeInstance: boolean, msg: ProductIdRequest): ProductIdRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ProductIdRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ProductIdRequest;
    static deserializeBinaryFromReader(message: ProductIdRequest, reader: jspb.BinaryReader): ProductIdRequest;
}

export namespace ProductIdRequest {
    export type AsObject = {
        productid: string,
    }
}

export class ProductCompRequest extends jspb.Message { 
    getProductid(): string;
    setProductid(value: string): ProductCompRequest;

    hasTime(): boolean;
    clearTime(): void;
    getTime(): google_protobuf_duration_pb.Duration | undefined;
    setTime(value?: google_protobuf_duration_pb.Duration): ProductCompRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ProductCompRequest.AsObject;
    static toObject(includeInstance: boolean, msg: ProductCompRequest): ProductCompRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ProductCompRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ProductCompRequest;
    static deserializeBinaryFromReader(message: ProductCompRequest, reader: jspb.BinaryReader): ProductCompRequest;
}

export namespace ProductCompRequest {
    export type AsObject = {
        productid: string,
        time?: google_protobuf_duration_pb.Duration.AsObject,
    }
}

export class ProductStatusRequest extends jspb.Message { 
    getProductid(): string;
    setProductid(value: string): ProductStatusRequest;
    getStatus(): string;
    setStatus(value: string): ProductStatusRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ProductStatusRequest.AsObject;
    static toObject(includeInstance: boolean, msg: ProductStatusRequest): ProductStatusRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ProductStatusRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ProductStatusRequest;
    static deserializeBinaryFromReader(message: ProductStatusRequest, reader: jspb.BinaryReader): ProductStatusRequest;
}

export namespace ProductStatusRequest {
    export type AsObject = {
        productid: string,
        status: string,
    }
}
