'use server';

import { z } from "zod"
import { Metadata } from "@grpc/grpc-js"
import { RegisterRequest, StandardResponse } from "@/proto/auth_pb"
import { authClient, unary_callback } from "@/server/grpc"
import type { Safe } from "@/server/safe"
import { RegisterSchema } from "@/schemas"

export const register = async  (values: z.infer<typeof RegisterSchema>) => {
    const validatedFields = RegisterSchema.safeParse(values)

    if (!validatedFields.success)
        return { status: false, message: validatedFields.error.message }

    const { username, password, confirmPassword } = validatedFields.data

    if (password !== confirmPassword) {
        return { status: false, message: "Passwords do not match"}
    }

    const req = new RegisterRequest()
    req.setUsername(username)
    req.setPassword(password)

    const r = await new Promise<Safe<StandardResponse>>((res) => {
        authClient.register(req, new Metadata, unary_callback(res));
    });

    if (r.success) {
        return { success: r.success, message: r.data.toObject().message }
    }
    return { success: r.success, message: r.message }
}
