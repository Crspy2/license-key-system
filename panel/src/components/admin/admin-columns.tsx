"use client"

import Link from "next/link"
import { ColumnDef } from "@tanstack/react-table"
import { StaffObject } from "@/proto/staff_pb"
import { Avatar, AvatarFallback } from "@/components/ui/avatar"
import { PermissionSelect } from "@/components/admin/perm-select"
import { RoleSelect } from "@/components/admin/role-select"
import { AccessSwitch } from "@/components/admin/access-switch"

export const staffColumns: ColumnDef<StaffObject.AsObject>[] = [
    {
        accessorKey: "name",
        header: () => <div className="~text-xs/sm">Name</div>,
        cell: ({ row }) => (
            <Link href={`/staff/${row.original.id}`} className="flex flex-row ~gap-1.5/2.5 justify-start items-center">
                <Avatar className="~h-6/8 ~w-6/8 rounded-lg">
                    <AvatarFallback className="rounded-lg ~text-xs/sm">{row.original.name.toUpperCase().slice(0, 2)}</AvatarFallback>
                </Avatar>
                <span className="~text-xs/sm text-nowrap">{row.original.name}</span>
            </Link>
        ),
    },
    {
        accessorKey: "permsList",
        header: () => <div className="hidden sm:table-row ~text-xs/sm">Permissions</div>,
        cell: ({ row }) => (
            <div className="hidden sm:table-row">
                <PermissionSelect staff={row.original} />
            </div>
        ),
    },
    {
        accessorKey: "role",
        header: () => <div className="~text-xs/sm">Role</div>,
        cell: ({ row }) => <RoleSelect staff={row.original} />,
    },
    {
        accessorKey: "approved",
        header: () => <div className="~text-xs/sm">Approved</div>,
        cell: ({ row }) => <AccessSwitch staff={row.original} />,
    },
];
