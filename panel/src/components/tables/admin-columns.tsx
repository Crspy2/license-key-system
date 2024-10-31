"use client"

import { useState, useTransition } from "react"
import { ColumnDef, Row } from "@tanstack/react-table"
import { z } from "zod"
import { zodResolver } from "@hookform/resolvers/zod"
import { toast } from "sonner"
import { useForm } from "react-hook-form"
import { Switch } from "@/components/ui/switch"
import { StaffObject } from "@/proto/staff_pb"
import { PermissionsSchema, RoleSchema } from "@/schemas";
import { ConvertPermissionsToValues } from "@/lib/utils"
import { approveStaff, setStaffPermissions, setStaffRole } from "@/server/services/staff"
import { Avatar, AvatarFallback } from "@/components/ui/avatar"
import { Form, FormControl, FormField, FormItem, FormMessage } from "@/components/ui/form";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select"
import MultipleSelector, { Option } from '@/components/ui/multiple-selector'

interface CellProps {
    row: Row<StaffObject.AsObject>
}

export const staffColumns: ColumnDef<StaffObject.AsObject>[] = [
    {
        accessorKey: "name",
        header: () => <div className="text-xs lg:text-sm">Name</div>,
        cell: ({ row }) => (
            <div className="flex flex-row gap-1 justify-start items-center">
                <Avatar className="h-8 w-8 rounded-lg">
                    <AvatarFallback className="rounded-lg">{row.original.name.toUpperCase().slice(0, 2)}</AvatarFallback>
                </Avatar>
                <span className="text-xs lg:text-sm text-nowrap">{row.original.name}</span>
            </div>
        ),
    },
    {
        accessorKey: "permsList",
        header: () => <div className="text-sm">Permissions</div>,
        cell: ({ row }) => <PermissionCell row={row} />,
    },
    {
        accessorKey: "role",
        header: () => <div className="text-sm">Role</div>,
        cell: ({ row }) => <RoleCell row={row} />,
    },
    {
        accessorKey: "approved",
        header: () => <div className="text-xs lg:text-sm">Approved</div>,
        cell: ({ row }) => <ApprovedCell row={row} />,
    },
];


const PermissionCell = ({ row }: CellProps) => {
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
    const [previousPerms, setPreviousPerms] = useState(ConvertPermissionsToValues(row.original.permsList))

    const form = useForm<z.infer<typeof PermissionsSchema>>({
        resolver: zodResolver(PermissionsSchema),
        defaultValues: {
            staffId: row.original.id,
            permissions: row.original.permsList.map(p => ({
                label: p,
                value: ConvertPermissionsToValues([p])[0]
            })),
        }
    })

    const onSubmit = async (values: z.infer<typeof PermissionsSchema>) => {
        const permOptions = values.permissions.map(p => p.value)
        if (JSON.stringify(previousPerms) === JSON.stringify(permOptions)) {
            return null
        }

        startTransition(async () => {
            const res = await setStaffPermissions(values.staffId, permOptions)
            if (!res.success) {
                toast.error(res.message)
                return
            }
            setPreviousPerms(permOptions)
            toast.success(res.message)
            return
        })
    }

    return (
        <Form {...form}>
            <form onSubmit={form.handleSubmit(onSubmit)} className="w-2/3 space-y-6">
                <FormField
                    control={form.control}
                    name="permissions"
                    render={({ field }) => (
                        <FormItem>
                            <FormControl>
                                <MultipleSelector
                                    {...field}
                                    defaultOptions={row.original.permsList.map(p => ({
                                        label: p,
                                        value: ConvertPermissionsToValues([p])[0]
                                    }))}
                                    inputProps={{
                                        onBlur: (e) => form.handleSubmit(onSubmit)(e)
                                    }}
                                    options={permissions.filter(p => field.value.map(v => v.value != p.value))}
                                    emptyIndicator={
                                        <p className="bg-zinc-900 text-center text-sm text-gray-600 dark:text-gray-400">
                                            no results found.
                                        </p>
                                    }
                                />
                            </FormControl>
                            <FormMessage />
                        </FormItem>
                    )}
                />
            </form>
        </Form>
        // <Form {...form}>
        //     <form onSubmit={form.handleSubmit(onSubmit)}>
        //         <FormField
        //             control={form.control}
        //             name="permissions"
        //             render={({ field }) => (
        //                 <FormItem>
        //                     <FormControl>
        //                         {/*<PermissionSelect*/}
        //                         {/*    permissions={permissions}*/}
        //                         {/*    control={form.control}*/}
        //                         {/*    name={field.name}*/}
        //                         {/*/>*/}
        //
        //                         <MultiSelector
        //                             onValuesChange={field.onChange}
        //                             values={row.original.permsList}
        //                         >
        //                             <MultiSelectorTrigger>
        //                                 <MultiSelectorInput />
        //                             </MultiSelectorTrigger>
        //                             <MultiSelectorContent>
        //                                 <MultiSelectorList>
        //                                     {permissions.map((perm) => (
        //                                         <MultiSelectorItem key={perm.label} value={perm.label}>
        //                                             <span>{perm.label}</span>
        //                                         </MultiSelectorItem>
        //                                     ))}
        //                                 </MultiSelectorList>
        //                             </MultiSelectorContent>
        //                         </MultiSelector>
        //                     </FormControl>
        //                     <FormMessage />
        //                 </FormItem>
        //             )}
        //         />
        //         <Button type="submit" size="icon" disabled={isPending}>
        //             {isPending ? <PiSpinner className="w-7 h-7"/> :  <IoSave className="w-7 h-7" />}
        //         </Button>
        //     </form>
        // </Form>
    );
}

const RoleCell = ({ row }: CellProps) => {
    const roles = [
        { label: 'Staff', value: 0 },
        { label: 'Senior Staff', value: 1 },
        { label: 'Lead Staff', value: 2 },
        { label: 'Developer', value: 3 },
        { label: 'Owner', value: 4 },
    ];

    const [isPending, startTransition] = useTransition()

    const onSubmit = async (values: z.infer<typeof RoleSchema>) => {
        startTransition(async () => {
            const res = await setStaffRole(values.staffId, values.role)
            if (!res.success) {
                toast.error(res.message)
                return
            }
            toast.success(res.message)
            return
        })
    }

    return (
        <Select onValueChange={(value) => onSubmit({ staffId: row.original.id, role: +value || 0 })} defaultValue={row.original.role.toString()}>
            <SelectTrigger disabled={isPending}>
                <SelectValue placeholder="Select a role" />
            </SelectTrigger>
            <SelectContent>
                {roles.map(role => (
                    <div key={role.value + 1}>
                        <SelectItem value={role.value.toString()}>{role.label}</SelectItem>
                    </div>
                ))}
            </SelectContent>
        </Select>
    );
}

const ApprovedCell = ({ row }: CellProps) => {
    const [approved, setApproved] = useState(row.original.approved)

    const onCheckedChange = async () => {
        setApproved(!approved)
        const status = await approveStaff(row.original.id, !approved)
        if (!status.success) {
            return toast.error(`Error code: ${status.code}. Unable to approve staff member!`)
        }
        if (approved) {

        }
        return toast.success(status.message)
    }

    return (
        <Switch
            checked={approved}
            onCheckedChange={onCheckedChange}
        />
    )
}
