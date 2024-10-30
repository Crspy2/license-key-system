import { listStaff } from "@/server/services/staff";
import {getCurrentSession} from "@/server/services/session";
import { DataTable } from "@/components/tables/data-table";
import { staffColumns } from "@/components/tables/admin-columns";
import { Error } from "@/components/error";

const StaffPage = async () => {
    const session = await getCurrentSession();
    if (!session.success) return (
        <Error text={session.message! as string} />
    )
    if (!session.data?.staff?.permsList.includes("ProductStatus") || !session.data?.staff?.permsList.includes("ManageProducts")) return (
        <Error text="You do not have the correct permissions to view this page" />
    )

    return (
        <Error text="Page is not finished being developed" />
    )
}

export default StaffPage