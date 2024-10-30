"use client"

import {useState, useTransition} from "react"
import { ColumnDef, Row } from "@tanstack/react-table"
import { toast } from "sonner"
import { Switch } from "@/components/ui/switch"
import { StaffObject } from "@/proto/staff_pb"
import {approveStaff, setStaffPermissions, setStaffRole} from "@/server/services/staff"
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar"
import PermissionSelect from "@/components/dashboard/permission-select"
import {useForm} from "react-hook-form";
import { IoSave } from "react-icons/io5";
import {z} from "zod";
import {PermissionsSchema, RoleSchema} from "@/schemas";
import {zodResolver} from "@hookform/resolvers/zod";
import {ConvertPermissionsToValues, GetUserRoleText, GetUserRoleValue} from "@/lib/utils";
import {Form, FormControl, FormDescription, FormField, FormItem, FormLabel, FormMessage} from "@/components/ui/form";
import {Button} from "@/components/ui/button";
import {PiSpinner} from "react-icons/pi";
import RoleSelect from "@/components/dashboard/role-select";

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
                    <AvatarImage src={row.original.image} alt={row.original.name} />
                    <AvatarFallback className="rounded-lg">{row.original.name.toUpperCase().slice(0, 2)}</AvatarFallback>
                </Avatar>
                <span className="text-xs lg:text-sm text-nowrap">{row.original.name}</span>
            </div>
        ),
    },
    {
        accessorKey: "permsList",
        header: () => <div className="text-sm">Permissions</div>,
        // cell: ({ row }) => <span>{ConvertPermissionsToValues(row.original.permsList)}</span>,
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
        { label: 'Manage Users', value: 1 << 6 },
        { label: 'Generate Keys', value: 1 << 7 },
        { label: 'Manage Staff', value: 1 << 8 },
    ];

    const [isPending, startTransition] = useTransition()

    const form = useForm<z.infer<typeof PermissionsSchema>>({
        resolver: zodResolver(PermissionsSchema),
        defaultValues: {
            staffId: row.original.id,
            permissions: ConvertPermissionsToValues(row.original.permsList),
        }
    })

    const onSubmit = async (values: z.infer<typeof PermissionsSchema>) => {
        startTransition(async () => {
            const res = await setStaffPermissions(values.staffId, values.permissions)
            if (!res.success) {
                toast.error(res.message)
                return
            }
            toast.success(res.message)
            return
        })
    }

    return (
        <Form {...form}>
            <form onSubmit={form.handleSubmit(onSubmit)} className="flex flex-grow-0 items-center gap-3">
                <FormField
                    control={form.control}
                    name="permissions"
                    render={({ field }) => (
                        <FormItem>
                            <FormControl>
                                <PermissionSelect
                                    permissions={permissions}
                                    control={form.control}
                                    name={field.name}
                                />
                            </FormControl>
                            <FormMessage />
                        </FormItem>
                    )}
                />
                <Button type="submit" size="icon" disabled={isPending}>
                    {isPending ? <PiSpinner className="w-7 h-7"/> :  <IoSave className="w-7 h-7" />}
                </Button>
            </form>
        </Form>
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

    const form = useForm<z.infer<typeof RoleSchema>>({
        resolver: zodResolver(RoleSchema),
        defaultValues: {
            staffId: row.original.id,
            role: row.original.role,
        }
    })

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
        <Form {...form}>
            <form onSubmit={form.handleSubmit(onSubmit)} className="flex flex-grow-0 items-center gap-3">
                <FormField
                    control={form.control}
                    name="role"
                    render={({ field }) => (
                        <FormItem>
                            <FormControl>
                                <RoleSelect
                                    roles={roles}
                                    control={form.control}
                                    name={field.name}
                                />
                            </FormControl>
                            <FormMessage />
                        </FormItem>
                    )}
                />
                <Button type="submit" size="icon" disabled={isPending}>
                    {isPending ? <PiSpinner className="w-7 h-7"/> :  <IoSave className="w-7 h-7" />}
                </Button>
            </form>
        </Form>
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
