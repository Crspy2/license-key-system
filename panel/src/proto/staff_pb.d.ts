// package: protofiles
// file: staff.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";

export class StaffObject extends jspb.Message { 
    getId(): string;
    setId(value: string): StaffObject;
    getName(): string;
    setName(value: string): StaffObject;
    getPasswordhash(): string;
    setPasswordhash(value: string): StaffObject;
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
        passwordhash: string,
        permsList: Array<string>,
        approved: boolean,
    }
}
