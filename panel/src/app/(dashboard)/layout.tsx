import { PropsWithChildren } from "react"
import { redirect } from "next/navigation"
import { getCurrentSession } from "@/server/services/session"
import { SidebarInset, SidebarProvider } from "@/components/ui/sidebar";
import { AppSidebar } from "@/components/app-sidebar"

const DashboardLayout = async ({ children }: PropsWithChildren) => {
    const session = await getCurrentSession()

    if (!session || !session.success)
        redirect("/login")

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