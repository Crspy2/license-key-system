"use client"

import * as React from "react"
import Link from "next/link"
import { SessionObject } from "@/proto/auth_pb"

import { MainNav } from "@/components/main-nav"
import { NavSecondary } from "@/components/nav-secondary"
import { NavUser } from "@/components/nav-user"
import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarHeader,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
} from "@/components/ui/sidebar"
import Image from "next/image"

import { IoIosSend } from "react-icons/io"
import { PiLifebuoyBold } from "react-icons/pi"
import { FaBox, FaFileInvoice, FaKey, FaUsers, FaUsersGear } from "react-icons/fa6"
import { IoTerminal } from "react-icons/io5"
import { AiOutlineAudit } from "react-icons/ai"
import {StaffNav} from "@/components/staff-nav";
import {ConvertPermissionsToValues, GetUserRoleText, GetUserRoleValue} from "@/lib/utils";

const data = {
  user: {
    name: "shadcn",
    email: "m@example.com",
    avatar: "/avatars/shadcn.jpg",
  },
  mainNav: [
    {
      title: "Panel",
      url: "/",
      icon: IoTerminal,
    },
    {
      title: "Users",
      url: "/users",
      icon: FaUsers,
    },
    {
      title: "Products",
      url: "/products",
      icon: FaBox,
    },
    {
      title: "Keys",
      url: "/licenses",
      icon: FaKey,
    },
    {
      title: "Subscriptions",
      url: "/subscriptions",
      icon: FaFileInvoice,
    },
  ],
  staffNav:[
    {
      title: "Staff",
      url: "/staff",
      icon: FaUsersGear,
    },
    {
      title: "Logs",
      url: "/logs",
      icon: AiOutlineAudit,
    },
  ],
  secondaryNav: [
    {
      title: "Support",
      url: "https://github.com/crspy2",
      icon: PiLifebuoyBold,
    },
    {
      title: "Feedback",
      url: "https://discord.com/users/385568884511473664",
      icon: IoIosSend,
    },
  ],
}

export function AppSidebar({ session, ...props }: React.ComponentProps<typeof Sidebar> & { session: SessionObject.AsObject }) {
  if (!session.staff)
    return null
  return (
    <Sidebar variant="inset" {...props}>
      <SidebarHeader>
        <SidebarMenu>
          <SidebarMenuItem>
            <SidebarMenuButton size="lg" asChild>
              <Link href="/">
                <div className="flex aspect-square size-8 items-center justify-center rounded-lg bg-sidebar-primary text-sidebar-primary-foreground">
                  <Image src="/crspy_pfp.webp" alt="company logo" width={500} height={500} />
                </div>
                <div className="grid flex-1 text-left text-sm leading-tight">
                  <span className="truncate font-semibold">Altera</span>
                  <span className="truncate text-xs">Licences</span>
                </div>
              </Link>
            </SidebarMenuButton>
          </SidebarMenuItem>
        </SidebarMenu>
      </SidebarHeader>
      <SidebarContent>
        <MainNav items={data.mainNav} />
        {session.staff.role >= 1 && session.staff.permsList.includes("ManageUsers") && (
            <StaffNav items={data.staffNav} />
        )}
        <NavSecondary items={data.secondaryNav} className="mt-auto" />
      </SidebarContent>
      <SidebarFooter>
        <NavUser staff={session.staff!} />
      </SidebarFooter>
    </Sidebar>
  )
}
