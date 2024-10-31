import { getCurrentSession } from "@/server/services/session"
import { listLogs } from "@/server/services/logs"
import { listStaff } from "@/server/services/staff"
import { LogList } from "@/components/logs/logList"
import { Error } from "@/components/error"

const LogsPage = async () => {
    const session = await getCurrentSession();
    if (!session.success) return (
        <Error text={session.message! as string} />
    )
    if (!session.data?.staff?.permsList.includes("ViewLogs")) return (
        <Error text="You do not have the correct permissions to view this page" />
    )
    const logs = await listLogs()
    if (!logs.success)
        return <Error text="Unable to fetch logs" />

    const staff = await listStaff()

    return (
        <div className="container mx-auto py-6 space-y-6 px-8">
            <LogList logs={logs.data} staff={staff.data || []} />
        </div>
    );
}

export default LogsPage