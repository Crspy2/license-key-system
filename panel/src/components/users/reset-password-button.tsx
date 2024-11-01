"use client"


import { useTransition } from "react"
import { toast } from "sonner"
import { UserObject } from "@/proto/user_pb"
import { resetPassword } from "@/server/services/user"
import { Button } from "@/components/ui/button"
import { LuLoader2 } from "react-icons/lu"

interface ResetPasswordButtonProps {
    user: UserObject.AsObject
}

export const ResetPasswordButton = ({ user }: ResetPasswordButtonProps) => {
    const [isPending, startTransition] = useTransition()

    const onClick = async () => {
        startTransition(async () => {
            const reset = await resetPassword(user.id)
            if (!reset.success) {
                toast.error(reset.message)
                return
            }

            toast.success(reset.message)
        })
    }
    return (
        <Button onClick={onClick} disabled={isPending} className="flex ~gap-1.5/2.5">
            {isPending && (
                <LuLoader2 className="animate-spin h-7 w-7" />
            )}
            <span>Reset Password</span>
        </Button>
    )
}