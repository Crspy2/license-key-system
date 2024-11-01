"use client"


import { useState, useTransition } from "react"
import { useForm } from "react-hook-form"
import { z } from "zod"
import { zodResolver } from "@hookform/resolvers/zod"
import { toast } from "sonner"
import { PermissionsSchema } from "@/schemas"
import { setStaffPermissions } from "@/server/services/staff"
import MultiSelect, { Option } from "@/components/ui/multi-select"
import { ConvertPermissionsToValues } from "@/lib/utils"
import { Form, FormControl, FormField, FormItem } from "@/components/ui/form"
import { StaffObject } from "@/proto/staff_pb"

interface PermissionSelectProps {
    staff: StaffObject.AsObject
}

export const PermissionSelect = ({ staff }: PermissionSelectProps) => {
    const permissions = [
        { label: 'Default', value: 1 << 0 },
        { label: 'HWID Reset', value: 1 << 1 },
        { label: 'Password Reset', value: 1 << 2 },
        { label: 'Compensate Keys', value: 1 << 3 },
        { label: 'Change Status', value: 1 << 4 },
        { label: 'Manage Products', value: 1 << 5 },
        { label: 'Generate Keys', value: 1 << 6 },
        { label: 'View Logs', value: 1 << 7 },
        { label: 'Manage Users', value: 1 << 8 },
        { label: 'Manage Staff', value: 1 << 9 },
    ] as Option[];

    const [isPending, startTransition] = useTransition()
    const [previousPerms, setPreviousPerms] = useState(staff.permsList.map(p => ({
        label: p,
        value: ConvertPermissionsToValues([p])[0]
    })))

    const form = useForm<z.infer<typeof PermissionsSchema>>({
        resolver: zodResolver(PermissionsSchema),
        defaultValues: {
            staffId: staff.id,
            permissions: staff.permsList.map(p => ({
                label: p,
                value: ConvertPermissionsToValues([p])[0]
            })),
        }
    })

    const onSubmit = async (values: z.infer<typeof PermissionsSchema>) => {
        const permOptions = values.permissions.map(p => p.value)
        const previousPermValues = previousPerms.map(p => p.value)
        if (JSON.stringify(previousPermValues) === JSON.stringify(permOptions)) {
            return null
        }

        startTransition(async () => {
            const res = await setStaffPermissions(values.staffId, permOptions)
            if (!res.success) {
                toast.error(res.message)
                form.setValue("permissions", previousPerms)
                return
            }
            setPreviousPerms(values.permissions)
            toast.success(res.message)
            return
        })
    }

    return (
        <Form {...form}>
            <form onSubmit={form.handleSubmit(onSubmit)}>
                <FormField
                    control={form.control}
                    name="permissions"
                    render={({ field }) => (
                        <FormItem>
                            <FormControl>
                                <MultiSelect
                                    {...field}
                                    defaultOptions={staff.permsList.map(p => ({
                                        label: p,
                                        value: ConvertPermissionsToValues([p])[0]
                                    }))}
                                    inputProps={{
                                        onBlur: (e) => form.handleSubmit(onSubmit)(e)
                                    }}
                                    options={permissions.filter(p => field.value.map(v => v.value != p.value))}
                                    emptyIndicator={
                                        <p className="text-center ~text-xs/sm text-gray-600 dark:text-gray-400">
                                            no results found.
                                        </p>
                                    }
                                />
                            </FormControl>
                        </FormItem>
                    )}
                />
            </form>
        </Form>
    );
}
