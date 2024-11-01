import {
    type ServiceError,
    credentials,
    ClientReadableStream,
} from '@grpc/grpc-js'
import { AuthClient } from '@/proto/auth_grpc_pb'
import { StaffClient } from "@/proto/staff_grpc_pb"
import { LogClient } from "@/proto/logs_grpc_pb"
import { UserClient } from "@/proto/user_grpc_pb"
import { Safe } from "@/server/safe"
import { Status } from "@grpc/grpc-js/build/src/constants"
import path from "node:path"
import * as fs from "node:fs"
import { StaffObject } from "@/proto/staff_pb"

const loadSSLCertificate = () => {
    const cert = Buffer.from(process.env.SSL_CERTIFICATE || "", 'base64').toString('utf-8');
    const caPath = path.join('/tmp', 'ca.crt');
    fs.writeFileSync(caPath, cert);
    return caPath;
}

// Read the server certificate file
export const authClient = new AuthClient(
    process.env.GRPC_SERVER_ADDRESS || 'localhost:8080',
    credentials.createSsl(fs.readFileSync(loadSSLCertificate()))
);

export const staffClient = new StaffClient(
    process.env.GRPC_SERVER_ADDRESS || 'localhost:8080',
    credentials.createSsl(fs.readFileSync(loadSSLCertificate()))
);

export const logClient = new LogClient(
    process.env.GRPC_SERVER_ADDRESS || 'localhost:8080',
    credentials.createSsl(fs.readFileSync(loadSSLCertificate()))
);

export const userClient = new UserClient(
    process.env.GRPC_SERVER_ADDRESS || 'localhost:8080',
    credentials.createSsl(fs.readFileSync(loadSSLCertificate()))
);


const map_error_code = (code: number): number => {
    switch (code) {
        case Status.OK:
            return 200;
        case Status.INVALID_ARGUMENT:
            return 422;
        case Status.NOT_FOUND:
            return 404;
        case Status.UNAUTHENTICATED:
            return 401;
        default:
            return 500;
    }
}

export function unary_callback<T>(
    res: (value: Safe<T>) => void,
): (err: ServiceError | null, data: T | undefined) => void {
    return (err: ServiceError | null, data: T | undefined) => {
        if (err) {
            if (err.code === Status.INVALID_ARGUMENT) {
                try {
                    const fields = JSON.parse(err.details) as { field: string; tag: string }[];
                    return res({
                        success: false,
                        fields,
                        message: "Validation error",
                        code: map_error_code(err.code),
                    });
                } catch (_) {
                    //
                }
            }
            return res({
                success: false,
                message: err.details,
                code: map_error_code(err.code),
            });
        }
        if (!data) {
            return res({
                success: false,
                message: "No data returned",
                code: 500,
            });
        }

        return res({ data, success: true, message: "Success" });
    };
}


export function stream_callback<T>(
    res: (value: Safe<T[]>) => void,
): (stream: ClientReadableStream<T>) => void {
    return (stream: ClientReadableStream<T>) => {
        const dataBuffer: T[] = [];

        stream.on("data", (data: any) => {
            dataBuffer.push(data);
        });

        stream.on("end", () => {
            res({
                success: true,
                data: dataBuffer,
                message: "Stream ended successfully",
            });
        });

        stream.on("error", (err: ServiceError) => {
            res({
                success: false,
                message: err.details || "Stream error",
                code: map_error_code(err.code),
            });
        });
    };
}
