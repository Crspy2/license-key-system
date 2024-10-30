import { listStaff } from "@/server/services/staff"
import {getCurrentSession} from "@/server/services/session"
import { DataTable } from "@/components/tables/data-table"
import { staffColumns } from "@/components/tables/admin-columns"
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
        <div className="m-8 gap-4 space-y-5">
            <h1 className="text-4xl font-bold">Staff Management</h1>
            <div className="p-4 rounded-md">
                <DataTable columns={staffColumns} data={staff.data!}/>
            </div>
        </div>
    )
}

export default StaffPage