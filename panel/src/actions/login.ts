'use server';


import { cookies, headers } from "next/headers"
import { userAgent } from "next/server"
import { z } from "zod"
import { Metadata } from "@grpc/grpc-js"
import { LoginRequest, LoginResponse } from "@/proto/auth_pb"
import { authClient, unary_callback } from "@/server/grpc"
import type { Safe } from "@/server/safe"
import { LoginSchema } from "@/schemas"

export const login = async  (values: z.infer<typeof LoginSchema>) => {
    const validatedFields = LoginSchema.safeParse(values)

    if (!validatedFields.success)
        return { status: false, message: validatedFields.error.message }

    const {username, password} = validatedFields.data

    const h = await headers()

    const req = new LoginRequest()
    req.setUsername(username)
    req.setPassword(password)

    const md = new Metadata
    md.set("x-forwarded-for", h.get("x-forwarded-for") || "::1")
    md.set("x-client-user-agent", userAgent({headers: h}).ua)

    const r = await new Promise<Safe<LoginResponse>>((res) => {
        authClient.login(req, md, unary_callback(res));
    });

    if (r.success) {
        const cookieStore = await cookies()
        const hour = 60 * 60 * 1000;
        cookieStore.set('session_token', r.data.toObject().data!.sessionid, {
            secure: true,
            httpOnly: true,
            sameSite: 'lax',
            expires: new Date(Date.now() + hour * 2),
        })
        cookieStore.set('csrf_token', r.data.toObject().data!.csrftoken, {
            secure: true,
            httpOnly: true,
            sameSite: 'lax',
        })
    }

    if (r.success) {
        return { success: r.success, message: r.data.toObject().message }
    }
    return { success: r.success, message: r.message }
}
