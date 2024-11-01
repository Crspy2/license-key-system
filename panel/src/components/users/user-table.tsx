"use client"

import {UserObject} from "@/proto/user_pb"
import {DataTable} from "@/components/tables/data-table"
import {userColumns} from "@/components/users/user-columns"
import {Search} from "lucide-react";
import {Input} from "@/components/ui/input";
import {useState} from "react";

interface UserTableProps {
    users: UserObject.AsObject[]
    search: string
}

export const UserTable = ({
       users,
}: UserTableProps) => {
    const [search, setSearch] = useState('')

    const filteredUsers = users.filter((user) => {
        return search === '' ||
            user.name.toLowerCase().includes(search.toLowerCase()) ;
    })

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
            <DataTable columns={userColumns} data={filteredUsers}/>
        </div>
    )
}