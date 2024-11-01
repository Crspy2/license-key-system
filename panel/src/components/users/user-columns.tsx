"use client"

import { ColumnDef } from "@tanstack/react-table"
import { UserObject } from "@/proto/user_pb"
import { ResetPasswordButton } from "@/components/users/reset-password-button"
import { ResetHwidButton } from "@/components/users/reset-hwid-button"
import { UserBanButton } from "@/components/users/user-ban-button"

export const userColumns: ColumnDef<UserObject.AsObject>[] = [
    {
        accessorKey: "id",
        header: () => <div className="~text-xs/sm">UID</div>,
        cell: ({ row }) => <span className="~text-xs/sm text-nowrap">{row.original.id}</span>,
    },
    {
        accessorKey: "name",
        header: () => <div className="~text-xs/sm">Name</div>,
        cell: ({ row }) => <span className="~text-xs/sm text-nowrap">{row.original.name}</span>,
    },
    {
        accessorKey: "hwid",
        header: () => <div className="hidden sm:table-row ~text-xs/sm">Reset HWID</div>,
        cell: ({ row }) => <ResetHwidButton user={row.original} />,
    },
    {
        accessorKey: "pswd",
        header: () => <div className="~text-xs/sm">Reset Password</div>,
        cell: ({ row }) => <ResetPasswordButton user={row.original} />,
    },
    {
        accessorKey: "banned",
        header: () => <div className="~text-xs/sm">Ban User</div>,
        cell: ({ row }) => <UserBanButton user={row.original} />,
    },
];
