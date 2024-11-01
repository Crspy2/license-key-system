"use client"


import {useState, useTransition} from "react"
import Link from "next/link"
import { useForm } from "react-hook-form"
import { z } from "zod"
import { zodResolver } from "@hookform/resolvers/zod"
import { RegisterSchema } from "@/schemas"
import { toast } from "sonner"
import {
    ResponsiveModal,
    ResponsiveModalContent, ResponsiveModalDescription, ResponsiveModalHeader, ResponsiveModalTitle,
    ResponsiveModalTrigger,
} from "@/components/ui/modal"
import { Button } from "@/components/ui/button"
import { Form, FormControl, FormField, FormItem, FormMessage } from "@/components/ui/form"
import { Label } from "@/components/ui/label"
import { FormInput } from "@/components/ui/form-input"
import {createUser} from "@/server/services/user";

export const CreateUserButton = () => {
    const [isPending, startTransition] = useTransition()
    const [isModalOpen, setModalOpen] = useState(false)

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
                return
            }

            const user = await createUser(values.username, values.password)
            if (!user.success) {
                if (user.message.includes("name")) {
                    form.setError("username", {
                        type: "validate",
                        message: user.message,
                    })
                } else {
                    toast.error(user.message)
                }
                return
            }

            toast.success("User account created successfully")
            setModalOpen(false)
        })
    }

    return (
        <ResponsiveModal open={isModalOpen} onOpenChange={setModalOpen}>
            <ResponsiveModalTrigger asChild>
                <Button variant="outline">Create User</Button>
            </ResponsiveModalTrigger>
            <ResponsiveModalContent>
                <ResponsiveModalHeader>
                    <ResponsiveModalTitle>Create a User Account</ResponsiveModalTitle>
                    <ResponsiveModalDescription>
                        Create a user account by specifying a username and password.
                    </ResponsiveModalDescription>
                </ResponsiveModalHeader>
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
                        <Button type="submit" disabled={isPending} className="w-full">
                            Create Account
                        </Button>
                    </form>
                </Form>
            </ResponsiveModalContent>
        </ResponsiveModal>
    )
}