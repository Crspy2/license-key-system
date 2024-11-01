import { getCurrentSession } from "@/server/services/session"
import { listUsers } from "@/server/services/user"
import { DataTable } from "@/components/tables/data-table"
import { CreateUserButton } from "@/components/users/create-user-button"
import { userColumns } from "@/components/users/user-columns"
import { Error } from "@/components/error"
import {UserTable} from "@/components/users/user-table";

const UsersPage = async () => {
    const session = await getCurrentSession();
    if (!session.success) return (
        <Error text={session.message! as string} />
    )
    if (!session.data?.staff?.permsList.includes("Default")) return (
        <Error text="You do not have the correct permissions to view this page" />
    )

    return (
        <div className="container mx-auto py-6 space-y-6 px-8">
            <div className="py-6 space-y-6">
                <div className="flex flex-row items-center justify-between">
                    <h1 className="~text-3xl/4xl font-bold">User List</h1>
                    <CreateUserButton/>
                </div>
                <div className="p-4 rounded-md">
                    <UserTable />
                </div>
            </div>
        </div>
    )
}

export default UsersPage