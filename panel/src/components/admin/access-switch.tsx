"use client"

import { useState } from "react"
import { toast } from "sonner"
import { StaffObject } from "@/proto/staff_pb"
import { approveStaff } from "@/server/services/staff"
import { Switch } from "@/components/ui/switch"

interface ApproveSwitchProps {
    staff: StaffObject.AsObject
}

export const AccessSwitch = ({ staff }: ApproveSwitchProps) => {
    const [approved, setApproved] = useState(staff.approved)

    const onCheckedChange = async () => {
        setApproved(!approved)
        const status = await approveStaff(staff.id, !approved)
        if (!status.success) {
            return toast.error(`Error code: ${status.code}. Unable to approve staff member!`)
        }
        if (approved) {

        }
        return toast.success(status.message)
    }

    return (
        <Switch
            checked={approved}
            onCheckedChange={onCheckedChange}
        />
    )
}