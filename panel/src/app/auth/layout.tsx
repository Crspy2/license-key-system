import { PropsWithChildren } from "react"
import { redirect } from "next/navigation"
import { getCurrentSession } from "@/server/services/session"

const AuthLayout = async ({ children }: PropsWithChildren) => {
    const session = await getCurrentSession()

    if (session && session.success)
        redirect("/panel")

    return (
        <div className="bg-black min-h-screen w-full">
            {children}
        </div>
    )
}

export default AuthLayout