"use server"

import { cache } from "react"
import type { Safe } from "@/server/safe"
import { Empty } from "google-protobuf/google/protobuf/empty_pb"
import { staffClient, stream_callback, unary_callback } from "@/server/grpc"
import {
    ApprovalResponse,
    MultiPermissionRequest,
    StaffAccessRequest,
    StaffIdRequest,
    StaffObject, StaffRoleRequest,
} from "@/proto/staff_pb"
import { SetMetadata } from "@/server/services/helpers"
import { StandardResponse } from "@/proto/globals_pb"

export const approveStaff = cache(async (staffId: string, approved: boolean) => {
    const md = await SetMetadata()
    const req = new StaffAccessRequest();
    req.setStaffid(staffId);
    req.setApproved(approved)

    const r = await new Promise<Safe<ApprovalResponse>>((res) => {
        staffClient.setStaffAccess(req, md, unary_callback(res));
    });

    if (r.success) {
        return { success: r.success, message: r.data.getMessage() }
    }

    return { success: r.success, code: r.code, message: r.message }
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

    return { success: r.success, code: r.code, message: r.message }
})

export const listStaff = cache(async () => {
    const md = await SetMetadata()

    const r = await new Promise<Safe<StaffObject[]>>((res) => {
        const stream = staffClient.listStaffStream(new Empty(), md)
        stream_callback(res)(stream)
    });

    if (r.success) {
        const staffList: StaffObject.AsObject[] = []
        for (const staff of r.data) {
            staffList.push(staff.toObject())
        }

        return { success: r.success, message: "Retrieved staff list", data: staffList }
    }

    return { success: r.success, code: r.code, message: r.message }
})

export const setStaffPermissions = async (staffId: string, permissions: number[]) => {
    const md = await SetMetadata()

    const req = new MultiPermissionRequest()
    req.setStaffid(staffId)
    req.setPermissionsList(permissions)

    const r = await new Promise<Safe<StandardResponse>>((res) => {
        staffClient.setStaffPermissions(req, md, unary_callback(res))
    })


    if (r.success) {
        return { success: r.success, message: r.data.getMessage(), data: r.data.toObject() }
    }

    return { success: r.success, code: r.code, message: r.message }
}

export const setStaffRole = async (staffId: string, role: number) => {
    const md = await SetMetadata()

    const req = new StaffRoleRequest()
    req.setStaffid(staffId)
    req.setRole(role)

    const r = await new Promise<Safe<StandardResponse>>((res) => {
        staffClient.setStaffRole(req, md, unary_callback(res))
    })


    if (r.success) {
        return { success: r.success, message: r.data.getMessage(), data: r.data.toObject() }
    }

    return { success: r.success, code: r.code, message: r.message }
}