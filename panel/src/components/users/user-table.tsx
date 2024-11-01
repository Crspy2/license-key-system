"use client"

import { useEffect, useState } from "react"
import { toast } from "sonner"
import { useDebouncedSearch } from "@/hooks/useDebouncedSearch"
import { searchUsers } from "@/server/services/user"
import { UserObject } from "@/proto/user_pb"
import { DataTable } from "@/components/tables/data-table"
import { userColumns } from "@/components/users/user-columns"
import { Search } from "lucide-react";
import { Input } from "@/components/ui/input"

export const UserTable = () => {
    const [search, setSearch] = useState('')
    const [filteredUsers, setFilteredUsers] = useState<UserObject.AsObject[]>()

    const debouncedQuery = useDebouncedSearch(search, 500);
    useEffect(() => {
        if (!debouncedQuery) {
            setFilteredUsers([]);
            return;
        }
        searchUsers(search).then((usrs) => {
            setFilteredUsers(usrs.data!)
        }).catch((err) => {
            toast.error(err)
        })
    }, [search])
    return (
        <div className="flex flex-col">
            <div className="relative flex-1">
                <Search className="absolute left-3 top-3 h-4 w-4 text-[hsl(215.4_16.3%_46.9%)]"/>
                <Input
                    placeholder="Search users..."
                    value={search}
                    onChange={(e) => setSearch(e.target.value)}
                    className="pl-9"
                />
            </div>
            <DataTable columns={userColumns} data={filteredUsers || []}/>
        </div>
    )
}