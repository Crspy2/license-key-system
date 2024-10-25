import { type ServiceError, credentials } from '@grpc/grpc-js';
import { AuthClient } from '@/proto/auth_grpc_pb';
import { Safe } from "@/server/safe";
import { Status } from "@grpc/grpc-js/build/src/constants"

export const authClient = new AuthClient(
    process.env.GRPC_SERVER_ADDRESS || 'localhost:8080',
    credentials.createInsecure()
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
            console.log("NO DATA")
            return res({
                success: false,
                message: "No data returned",
                code: 500,
            });
        }

        console.log("RES: ", JSON.stringify(res));
        return res({ data, success: true, message: "Success" });
    };
}
