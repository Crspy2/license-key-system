"use server"
import "server-only"

import { cache } from "react"
import { cookies } from "next/headers"
import { SingleSessionResponse } from "@/proto/auth_pb"
import { StandardResponse } from "@/proto/globals_pb"
import type { Safe } from "@/server/safe"
import { Empty } from "google-protobuf/google/protobuf/empty_pb"
import { authClient, unary_callback } from "@/server/grpc"
import { Metadata } from "@grpc/grpc-js"

export const getCurrentSession = cache(async () => {
    const cookieStore = await cookies()
    const session_token = cookieStore.get("session_token")
    const csrf_token = cookieStore.get("csrf_token")

    const md = new Metadata
    md.set("session_token", session_token?.value || "")
    md.set("csrf_token", csrf_token?.value || "")

    const r = await new Promise<Safe<SingleSessionResponse>>((res) => {
        authClient.getSessionInfo(new Empty(), md, unary_callback(res));
    });

    if (r.success) {
        return { success: r.success, message: r.data.getMessage(), data: r.data.getData()?.toObject() }
    }

    return { success: r.success, code: r.code }
})


export const logOut = async () => {
    const cookieStore = await cookies()
    const session_token = cookieStore.get("session_token")
    const csrf_token = cookieStore.get("csrf_token")

    const md = new Metadata
    md.set("session_token", session_token?.value || "")
    md.set("csrf_token", csrf_token?.value || "")

    const r = await new Promise<Safe<StandardResponse>>((res) => {
        authClient.logout(new Empty(), md, unary_callback(res));
    });

    if (r.success) {
        return { success: r.success, message: r.data.getMessage(), data: r.data.toObject() }
    }

    return { success: r.success, code: r.code }
}