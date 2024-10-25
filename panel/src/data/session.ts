import "server-only"

import { SingleSessionRequest, SingleSessionResponse } from "@/proto/auth_pb";
import type { Safe } from "@/server/safe";
import { authClient, unary_callback } from "@/server/grpc";
import { Metadata } from "@grpc/grpc-js";
import { cookies, headers } from "next/headers";

export const getCurrentSession = async () => {
    const cookieStore = await cookies()
    const session_token = cookieStore.get("session_token")
    if (!session_token)
        return null

    const h = await headers()
    const req = new SingleSessionRequest()
    req.setSessionid(session_token?.value)
    req.setIp(h.get("x-forwarded-for") || "::1")

    const r = await new Promise<Safe<SingleSessionResponse>>((res) => {
        authClient.getSessionInfo(req, new Metadata, unary_callback(res));
    });

    if (r.success) {
        return { success: r.success, message: r.data.getMessage(), data: r.data.getData()?.toObject() }
    }

    return { success: r.success, code: r.code }
}
