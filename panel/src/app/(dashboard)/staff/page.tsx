import { listStaff } from "@/server/services/staff"
import {getCurrentSession} from "@/server/services/session"
import { DataTable } from "@/components/tables/data-table"
import { staffColumns } from "@/components/admin/admin-columns"
import { Error } from "@/components/error"

const StaffPage = async () => {
    const session = await getCurrentSession();
    if (!session.success) return (
        <Error text={session.message! as string} />
    )
    if (!session.data?.staff?.permsList.includes("ManageStaff")) return (
        <Error text="You do not have the correct permissions to view this page" />
    )
    const staff = await listStaff();

    return (
        <div className="container mx-auto py-6 space-y-6 px-8">
            <div className="py-6 space-y-6">
                <h1 className="~text-3xl/4xl font-bold">Staff Management</h1>
                <div className="p-4 rounded-md">
                    <DataTable columns={staffColumns} data={staff.data!}/>
                </div>
            </div>
        </div>
    )
}

export default StaffPage