"use server"

import { cache } from "react"
import type { Safe } from "@/server/safe"
import { Empty } from "google-protobuf/google/protobuf/empty_pb"
import { userClient, stream_callback, unary_callback } from "@/server/grpc"
import {
    UserCreateRequest,
    UserIdRequest,
    UserObject,
} from "@/proto/user_pb"
import { SetMetadata } from "@/server/services/helpers"
import { StandardResponse } from "@/proto/globals_pb"

export const createUser = async (name: string, password: string) => {
    const md = await SetMetadata()
    const req = new UserCreateRequest()
    req.setName(name)
    req.setPassword(password)

    const r = await new Promise<Safe<UserObject>>((res) => {
        userClient.createUser(req, md, unary_callback(res))
    })

    if (r.success) {
        return { success: r.success, message: "User Created", data: r.data.toObject() }
    }

    return { success: r.success, code: r.code, message: r.message }
}

export const getUser = cache(async (userId: number) => {
    const md = await SetMetadata()

    const req = new UserIdRequest()
    req.setUserid(userId)

    const r = await new Promise<Safe<UserObject>>((res) => {
        userClient.getUser(req, md, unary_callback(res))
    })

    if (r.success) {
        return { success: r.success, message: "Retrieved user information", data: r.data.toObject() }
    }

    return { success: r.success, code: r.code, message: r.message }
})

export const listUsers = cache(async () => {
    const md = await SetMetadata()

    const r = await new Promise<Safe<UserObject[]>>((res) => {
        const stream = userClient.listUsersStream(new Empty(), md)
        stream_callback(res)(stream)
    })

    if (r.success) {
        const staffList: UserObject.AsObject[] = []
        for (const staff of r.data) {
            staffList.push(staff.toObject())
        }

        return { success: r.success, message: "Retrieved staff list", data: staffList }
    }

    return { success: r.success, code: r.code, message: r.message }
})

export const resetHardwareId = async (userId: number) => {
    const md = await SetMetadata()

    const req = new UserIdRequest()
    req.setUserid(userId)

    const r = await new Promise<Safe<StandardResponse>>((res) => {
        userClient.resetHardwareId(req, md, unary_callback(res))
    })


    if (r.success) {
        return { success: r.success, message: r.data.getMessage(), data: r.data.toObject() }
    }

    return { success: r.success, code: r.code, message: r.message }
}

export const resetPassword = async (userId: number) => {
    const md = await SetMetadata()

    const req = new UserIdRequest()
    req.setUserid(userId)

    const r = await new Promise<Safe<StandardResponse>>((res) => {
        userClient.resetPassword(req, md, unary_callback(res))
    })


    if (r.success) {
        return { success: r.success, message: r.data.getMessage(), data: r.data.toObject() }
    }

    return { success: r.success, code: r.code, message: r.message }
}

export const banUser = async (userId: number) => {
    const md = await SetMetadata()

    const req = new UserIdRequest()
    req.setUserid(userId)

    const r = await new Promise<Safe<StandardResponse>>((res) => {
        userClient.banUser(req, md, unary_callback(res))
    })


    if (r.success) {
        return { success: r.success, message: r.data.getMessage(), data: r.data.toObject() }
    }

    return { success: r.success, code: r.code, message: r.message }
}

export const revokeUserBan = async (userId: number) => {
    const md = await SetMetadata()

    const req = new UserIdRequest()
    req.setUserid(userId)

    const r = await new Promise<Safe<StandardResponse>>((res) => {
        userClient.revokeBan(req, md, unary_callback(res))
    })


    if (r.success) {
        return { success: r.success, message: r.data.getMessage(), data: r.data.toObject() }
    }

    return { success: r.success, code: r.code, message: r.message }
}