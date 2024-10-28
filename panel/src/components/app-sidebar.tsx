"use client"

import * as React from "react"
import Link from "next/link"
import { SessionObject } from "@/proto/auth_pb"

import { NavMain } from "@/components/nav-main"
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
import { FaBox, FaFileInvoice, FaKey, FaUserSecret, FaUsers, FaUsersGear } from "react-icons/fa6"
import { IoTerminal } from "react-icons/io5"

const data = {
  user: {
    name: "shadcn",
    email: "m@example.com",
    avatar: "/avatars/shadcn.jpg",
  },
  navMain: [
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
    {
      title: "Offsets",
      url: "/offsets",
      icon: FaUserSecret,
    },
    {
      title: "Staff",
      url: "/staff",
      icon: FaUsersGear,
    },
  ],
  navSecondary: [
    {
      title: "Support",
      url: "#",
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
                  <Image src="/altera_dark.png" alt="company logo" width={500} height={500} />
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
        <NavMain items={data.navMain} />
        <NavSecondary items={data.navSecondary} className="mt-auto" />
      </SidebarContent>
      <SidebarFooter>
        <NavUser staff={session.staff!} />
      </SidebarFooter>
    </Sidebar>
  )
}
