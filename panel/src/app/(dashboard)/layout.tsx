import { PropsWithChildren } from "react"
import { redirect } from "next/navigation"
import { getCurrentSession } from "@/server/services/session"
import { SidebarInset, SidebarProvider } from "@/components/ui/sidebar";
import { AppSidebar } from "@/components/app-sidebar"
import {Error} from "@/components/error";

const DashboardLayout = async ({ children }: PropsWithChildren) => {
    const session = await getCurrentSession()

    if (!session || !session.success)
        redirect("/login")

    if (!session.data?.staff?.approved) {
        return (
            <Error text="Your account has not been approved! Please contact an administrator for support." />
        )
    }
    return (
        <SidebarProvider>
            <AppSidebar session={session.data!} />
            <SidebarInset className="text-neutral-200">
                <>
                    {children}
                </>
            </SidebarInset>
        </SidebarProvider>
    )
}

export default DashboardLayout