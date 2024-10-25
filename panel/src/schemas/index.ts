import { z } from "zod"

export const LoginSchema = z.object({
    username: z.string().min(3),
    password: z.string().min(8),
})

export const RegisterSchema = z.object({
    username: z.string().min(3),
    password: z.string().min(8),
    confirmPassword: z.string().min(8)
}).refine((data) => data.password === data.confirmPassword, {
        message: 'Passwords must match',
        path: ['confirmPassword'],
});