import { Error } from "@/components/error"
import {getCurrentSession} from "@/server/services/session";

const LogsPage = async () => {
    const session = await getCurrentSession();
    if (!session.success) return (
        <Error text={session.message! as string} />
    )
    if (!session.data?.staff?.permsList.includes("ManageStaff")) return (
        <Error text="You do not have the correct permissions to view this page" />
    )

    return (
        <Error text="Page is not finished being developed" />
    )
}

export default LogsPage