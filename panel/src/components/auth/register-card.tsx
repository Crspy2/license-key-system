"use client"

import { useTransition } from "react"
import Link from "next/link"
import Image from "next/image"
import { useRouter } from "next/navigation"
import { toast } from "sonner"
import { z } from "zod"
import { zodResolver } from "@hookform/resolvers/zod"
import { useForm } from "react-hook-form"
import { RegisterSchema } from "@/schemas"
import { register } from "@/actions/register"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { Form, FormControl, FormField, FormItem, FormMessage } from "@/components/ui/form"
import { Label } from "@/components/ui/label"
import { FormInput } from "@/components/ui/form-input"

export const RegisterCard = () => {
    const router = useRouter()
    const [isPending, startTransition] = useTransition()

    const form = useForm<z.infer<typeof RegisterSchema>>({
        resolver: zodResolver(RegisterSchema),
        defaultValues: {
            username: "",
            password: "",
            confirmPassword: "",
        }
    })

    const onSubmit = async (values: z.infer<typeof RegisterSchema>) => {
        startTransition(async () => {
            if (values.password !== values.confirmPassword) {
                form.setError("confirmPassword", {
                    type: "validate",
                    message: "Passwords do not match",
                })
            }

            const data = await register(values)
            if (!data.success) {
                if (data.message.includes("name")) {
                    form.setError("username", {
                        type: "validate",
                        message: data.message,
                    })
                } else {
                    toast.error(data.message)
                }
                return
            }
            toast.success(data.message)
            router.push("/")
            return
        })
    }


    return (
        <Card className="rounded-2xl bg-neutral-900 bg-[radial-gradient(circle_at_50%_0%,theme(colors.white/10%),transparent)] mx-4 py-10 ring-1 ring-inset ring-white/5 sm:w-96 sm:mx-8">
            <CardHeader className="text-center">
                <Image src="/crspy_pfp.webp" alt="Company logo" className="mx-auto size-24 rounded-lg" width={500} height={500}  />
                <CardTitle className="mt-4 text-2xl font-semibold tracking-tight text-white">Register</CardTitle>
            </CardHeader>
            <CardContent>
                <Form {...form}>
                    <form onSubmit={form.handleSubmit(onSubmit)} className="w-full space-y-6">
                        <div className="space-y-4">
                            <FormField
                                control={form.control}
                                name="username"
                                render={({ field, fieldState }) => (
                                    <FormItem className="flex flex-col mb-2.5">
                                        <Label htmlFor="username">Username</Label>
                                        <FormControl>
                                            <FormInput
                                                {...field}
                                                disabled={isPending}
                                                type="text"
                                                autoComplete="username"
                                                className={fieldState.isTouched && fieldState.invalid ? "ring-red-400  focus:ring-red-400" : "focus:ring-blue-400"}
                                            />
                                        </FormControl>
                                        <FormMessage />
                                    </FormItem>
                                )}
                            />
                            <FormField
                                control={form.control}
                                name="password"
                                render={({ field, fieldState }) => (
                                    <FormItem className="flex flex-col mb-2.5">
                                        <Label htmlFor="password">Password</Label>
                                        <FormControl>
                                            <FormInput
                                                {...field}
                                                disabled={isPending}
                                                type="password"
                                                autoComplete="current-password"
                                                className={fieldState.isTouched && fieldState.invalid ? "ring-red-400  focus:ring-red-400" : "focus:ring-blue-400"}
                                            />
                                        </FormControl>
                                        <FormMessage />
                                    </FormItem>
                                )}
                            />
                            <FormField
                                control={form.control}
                                name="confirmPassword"
                                render={({ field, fieldState }) => (
                                    <FormItem className="flex flex-col mb-2.5">
                                        <Label htmlFor="confirmPassword">Verify Password</Label>
                                        <FormControl>
                                            <FormInput
                                                {...field}
                                                disabled={isPending}
                                                type="password"
                                                autoComplete="current-password"
                                                className={fieldState.isTouched && fieldState.invalid ? "ring-red-400  focus:ring-red-400" : "focus:ring-blue-400"}
                                            />
                                        </FormControl>
                                        <FormMessage />
                                    </FormItem>
                                )}
                            />
                        </div>
                        <button type="submit"
                                className="relative isolate w-full rounded-md bg-blue-500 px-3.5 py-1.5 text-center text-sm font-medium text-white shadow-[0_1px_0_0_theme(colors.white/10%)_inset,0_0_0_1px_theme(colors.white/5%)] outline-none before:absolute before:inset-0 before:-z-10 before:rounded-md before:bg-white/5 before:opacity-0 hover:before:opacity-100 focus-visible:outline-[1.5px] focus-visible:outline-offset-2 focus-visible:outline-blue-400 active:text-white/70 active:before:bg-black/10">
                            Sign Up
                        </button>
                        <p className="flex gap-1.5 justify-center text-center text-sm text-zinc-400">
                            Have an account?
                            <Link href="/login"
                               className="font-medium text-white decoration-white/20 underline-offset-4 outline-none hover:underline focus-visible:underline">
                                Sign In
                            </Link>
                        </p>
                    </form>
                </Form>
            </CardContent>
        </Card>
    )
}
