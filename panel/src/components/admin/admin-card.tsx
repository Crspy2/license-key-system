import { StaffObject } from "@/proto/staff_pb"
import { Avatar, AvatarFallback } from "@/components/ui/avatar"
import { Card, CardContent, CardHeader } from "@/components/ui/card"
import { listLogs } from "@/server/services/logs"
import { Error } from "@/components/error"
import { LogList } from "@/components/logs/logList"
import { RoleSelect } from "@/components/admin/role-select"
import { PermissionSelect } from "@/components/admin/perm-select"
import { AccessSwitch } from "@/components/admin/access-switch"

interface AdminCardParams {
    staff: StaffObject.AsObject
}


export const AdminCard = async ({ staff }: AdminCardParams) => {
    const logs = await listLogs()
    if (!logs.success)
        return <Error text="Unable to fetch logs" />

    const userLogs = logs.data.filter(l => l.staff?.id === staff.id)

    return (
        <Card>
            <CardHeader className="flex flex-row">
                {/*<div className="flex flex-col gap-4">*/}
                {/*    <div className="flex flex-row ~gap-1.5/2.5 ml-12 mt-12">*/}
                {/*        <div>*/}
                {/*            <Avatar className="h-10 w-10">*/}
                {/*                <AvatarFallback className="h-8 w-8 rounded-md rounded-tl-md">{staff.name.toUpperCase().slice(0, 2)}</AvatarFallback>*/}
                {/*            </Avatar>*/}
                {/*            <span className="text-xs lg:text-sm text-nowrap my-auto">{staff.name}</span>*/}
                {/*        </div>*/}
                {/*        <div className="flex flex-row items-center ~gap-1.5/2.5">*/}
                {/*            <RoleSelect staff={staff}/>*/}
                {/*            <AccessSwitch staff={staff}/>*/}
                {/*        </div>*/}
                {/*    </div>*/}
                {/*    <PermissionSelect staff={staff}/>*/}
                {/*</div>*/}

                <div className="flex flex-col space-y-4">
                    <div className="flex flex-row justify-between">
                        <div className="flex flex-row ~gap-1.5/2.5 justify-start items-center">
                            <Avatar className="~h-10/16 ~w-10/16 rounded-lg">
                                <AvatarFallback
                                    className="rounded-lg ~text-base/2xl">{staff.name.toUpperCase().slice(0, 2)}</AvatarFallback>
                            </Avatar>
                            <span className="~text-base/3xl font-semibold text-nowrap">{staff.name}</span>
                        </div>
                        <div className="flex flex-row items-center ~gap-1.5/2.5">
                            <div className="flex flex-row items-center ~gap-1.5/2.5">
                                <h3 className="hidden sm:inline-flex ~text-sm/base">Role: </h3>
                                <RoleSelect staff={staff}/>
                            </div>
                            <div className="flex flex-row items-center ~gap-1.5/2.5">
                                <h3 className="hidden sm:inline-flex ~text-sm/base">Staff Access:</h3>
                                <AccessSwitch staff={staff}/>
                            </div>
                        </div>
                    </div>
                    <PermissionSelect staff={staff}/>
                    </div>
            </CardHeader>
            <CardContent className="py-6 space-y-6">
                <LogList logs={userLogs}/>
            </CardContent>
        </Card>
    )
}
