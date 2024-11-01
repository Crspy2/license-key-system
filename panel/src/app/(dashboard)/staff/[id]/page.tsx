import { Error } from "@/components/error"
import {getCurrentSession} from "@/server/services/session";
import {AdminCard} from "@/components/admin/admin-card";
import {getStaff} from "@/server/services/staff";

const StaffInfoPage = async ({ params }: { params: Promise<{ id: string }> }) => {
    const session = await getCurrentSession();
    if (!session.success) return (
        <Error text={session.message! as string} />
    )
    if (!session.data?.staff?.permsList.includes("ManageStaff")) return (
        <Error text="You do not have the correct permissions to view this page" />
    )
    const { id } = await params
    const staff = await getStaff(id)
    if (!staff.success) return (
        <Error text="The requested staff member could not be found" />
    )

    return (
        <div className="container mx-auto py-6 space-y-6 px-8">
            <div className="py-6 space-y-6">
                <h1 className="~text-3xl/4xl font-bold">{staff.data.name}&apos;s Panel Account</h1>
                <AdminCard staff={staff.data}/>
            </div>
        </div>
    )
}

export default StaffInfoPage