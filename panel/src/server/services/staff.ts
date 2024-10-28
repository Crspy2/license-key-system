import { cache } from "react"
import { cookies } from "next/headers"
import { Metadata } from "@grpc/grpc-js"
import type { Safe } from "@/server/safe"
import { Empty } from "google-protobuf/google/protobuf/empty_pb"
import {staffClient, stream_callback, unary_callback} from "@/server/grpc"
import {ApprovalResponse, MultiPermissionRequest, StaffIdRequest, StaffObject} from "@/proto/staff_pb"
import { SetMetadata } from "@/server/services/helpers";
import {MultiSessionRequest} from "@/proto/auth_pb";
import {StandardResponse} from "@/proto/globals_pb";

export const approveStaff = cache(async (staffId: string) => {
    const md = await SetMetadata()
    const req = new StaffIdRequest();
    req.setStaffid(staffId);

    const r = await new Promise<Safe<ApprovalResponse>>((res) => {
        staffClient.approveStaff(req, md, unary_callback(res));
    });

    if (r.success) {
        return { success: r.success, message: r.data.getMessage(), data: r.data.getStaff() }
    }

    return { success: r.success, code: r.code }
})

export const getStaff = cache(async (staffId: string) => {
    const md = await SetMetadata()

    const req = new StaffIdRequest();
    req.setStaffid(staffId);

    const r = await new Promise<Safe<StaffObject>>((res) => {
        staffClient.getStaff(req, md, unary_callback(res))
    });

    if (r.success) {
        return { success: r.success, message: "Retrieved staff information", data: r.data.toObject() }
    }

    return { success: r.success, code: r.code }
})

export const listStaff = cache(async () => {
    const md = await SetMetadata()

    const r = await new Promise<Safe<StaffObject[]>>((res) => {
        const stream = staffClient.getAllStaffStream(new Empty(), md)
        stream_callback(res)(stream)
    });

    if (r.success) {
        return { success: r.success, message: "Retrieved staff list", data: r.data }
    }

    return { success: r.success, code: r.code }
})

export const setStaffPermissions = cache(async (staffId: string, permissions: number[]) => {
    const md = await SetMetadata()

    const req = new MultiPermissionRequest();
    req.setStaffid(staffId);
    req.setPermissionsList(permissions)

    const r = await new Promise<Safe<StandardResponse>>((res) => {
        staffClient.setStaffPermissions(req, md, unary_callback(res))
    });

    if (r.success) {
        return { success: r.success, message: "Retrieved staff information", data: r.data.toObject() }
    }

    return { success: r.success, code: r.code }
})