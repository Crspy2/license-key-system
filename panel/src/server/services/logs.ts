import { cache } from "react"
import { Empty } from "google-protobuf/google/protobuf/empty_pb"
import { LogIdRequest, LogObject } from "@/proto/logs_pb"
import type { Safe } from "@/server/safe";
import { SetMetadata } from "@/server/services/helpers"
import { logClient, stream_callback, unary_callback } from "@/server/grpc"

export const getLog = cache(async (logId: string) => {
    const md = await SetMetadata()

    const req = new LogIdRequest();
    req.setLogid(logId);

    const r = await new Promise<Safe<LogObject>>((res) => {
        logClient.getLog(req, md, unary_callback(res))
    });

    if (r.success) {
        return { success: r.success, message: "Retrieved log information", data: r.data.toObject() }
    }

    return { success: r.success, code: r.code, message: r.message }
})

export const listLogs = cache(async () => {
    const md = await SetMetadata()

    const r = await new Promise<Safe<LogObject[]>>((res) => {
        const stream = logClient.listLogsStream(new Empty(), md)
        stream_callback(res)(stream)
    });

    if (r.success) {
        const staffList: LogObject.AsObject[] = []
        for (const staff of r.data) {
            staffList.push(staff.toObject())
        }

        return { success: r.success, message: "Retrieved staff list", data: staffList }
    }

    return { success: r.success, code: r.code, message: r.message }
})
