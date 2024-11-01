"use client"


import {useState, useTransition} from "react"
import { toast } from "sonner"
import { UserObject } from "@/proto/user_pb"
import { banUser, revokeUserBan } from "@/server/services/user"
import { Button } from "@/components/ui/button"
import { LuLoader2 } from "react-icons/lu"
import {
    ResponsiveModal,
    ResponsiveModalContent, ResponsiveModalDescription,
    ResponsiveModalHeader, ResponsiveModalTitle,
    ResponsiveModalTrigger
} from "@/components/ui/modal";
import {Form, FormControl, FormField, FormItem, FormMessage} from "@/components/ui/form";
import {Label} from "@/components/ui/label";
import {FormInput} from "@/components/ui/form-input";

interface UserBanButtonProps {
    user: UserObject.AsObject
}

export const UserBanButton = ({ user }: UserBanButtonProps) => {
    const [isPending, startTransition] = useTransition()
    const [isBanned, setIsBanned] = useState(user.banned)
    const [isModalOpen, setModalOpen] = useState(false);

    const onClick = async () => {
        startTransition(async () => {
            let reset
            if (isBanned) {
                reset = await revokeUserBan(user.id)
            } else {
                reset = await banUser(user.id)
            }

            setModalOpen(false)
            if (!reset.success) {
                toast.error(reset.message)
                return
            }
            toast.success(reset.message)
            setIsBanned(!isBanned)
        })
    }
    return (
    <ResponsiveModal open={isModalOpen} onOpenChange={setModalOpen}>
        <ResponsiveModalTrigger asChild>
            <Button variant="destructive" disabled={isPending} className="flex ~gap-1.5/2.5">
                {isPending && (
                    <LuLoader2 className="animate-spin h-7 w-7" />
                )}
                <span>{isBanned ? "Unban User" : "Ban User"}</span>
            </Button>
        </ResponsiveModalTrigger>
        <ResponsiveModalContent>
            <ResponsiveModalHeader>
                <ResponsiveModalTitle>{isBanned ? "Unban User Confirmation" : "Ban User Confirmation"}</ResponsiveModalTitle>
                <ResponsiveModalDescription>
                    Please press the button below to confirm your action
                </ResponsiveModalDescription>
            </ResponsiveModalHeader>
                <Button variant="destructive" onClick={onClick} disabled={isPending} className="flex ~gap-1.5/2.5">
                    {isPending && (
                        <LuLoader2 className="animate-spin h-7 w-7" />
                    )}
                    <span>Confirm</span>
                </Button>
        </ResponsiveModalContent>
    </ResponsiveModal>
    )
}