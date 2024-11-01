"use client"

import { useTransition } from "react"
import { z } from "zod"
import { toast } from "sonner"
import { StaffObject } from "@/proto/staff_pb"
import { RoleSchema } from "@/schemas"
import { setStaffRole } from "@/server/services/staff"
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select"

interface RoleSelectProps {
    staff: StaffObject.AsObject
}

export const RoleSelect = ({ staff }: RoleSelectProps) => {
    const roles = [
        { label: 'Staff', value: 0 },
        { label: 'Senior Staff', value: 1 },
        { label: 'Lead Staff', value: 2 },
        { label: 'Developer', value: 3 },
        { label: 'Owner', value: 4 },
    ];

    const [isPending, startTransition] = useTransition()

    const onSubmit = async (values: z.infer<typeof RoleSchema>) => {
        startTransition(async () => {
            const res = await setStaffRole(values.staffId, values.role)
            if (!res.success) {
                toast.error(res.message)
                return
            }
            toast.success(res.message)
            return
        })
    }

    return (
        <Select onValueChange={(value) => onSubmit({ staffId: staff.id, role: +value || 0 })} defaultValue={staff.role.toString()}>
            <SelectTrigger disabled={isPending}>
                <SelectValue placeholder="Select a role" />
            </SelectTrigger>
            <SelectContent>
                {roles.map(role => (
                    <div key={role.value + 1}>
                        <SelectItem value={role.value.toString()}>{role.label}</SelectItem>
                    </div>
                ))}
            </SelectContent>
        </Select>
    );
}
