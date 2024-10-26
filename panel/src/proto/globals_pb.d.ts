// package: protofiles
// file: globals.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";

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
