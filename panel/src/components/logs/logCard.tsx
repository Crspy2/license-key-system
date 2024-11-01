import Link from "next/link"
import { formatRelative } from 'date-fns'
import { titleCase } from "title-case";
import { LogObject } from "@/proto/logs_pb"
import { Card, CardContent, CardHeader } from '@/components/ui/card'
import { FaCube, FaKey, FaUser, FaUserShield } from "react-icons/fa6"
import { AiFillProduct } from "react-icons/ai"

interface LogCardProps {
    log: LogObject.AsObject
}

export const LogCard = ({ log }: LogCardProps) => {
    return (
        <Card>
            <CardHeader className="flex flex-row items-center gap-4 space-y-0">
                <div className="flex ~h-8/10 ~w-8/10 items-center justify-center rounded-lg bg-[#313ee0]">
                    {log.object === "User" ? (
                        <FaUser className="~h-5/6 ~w-5/6" />
                    ) : log.object === "Staff" ? (
                        <FaUserShield className="~h-5/6 ~w-5/6" />
                    ) : log.object === "License Key" ? (
                        <FaKey className="~h-5/6 ~w-5/6" />
                    ) : log.object === "Product" ? (
                        <AiFillProduct className="~h-5/6 ~w-5/6" />
                    ) : (
                        <FaCube className="~h-5/6 ~w-5/6" />
                    )}
                </div>
                <div className="flex-1">
                    <div className="flex items-center justify-between">
                        <h3 className="~text-sm/base font-semibold">{log.title}</h3>
                        <time className="~text-xs/sm text-[hsl(215.4_16.3%_46.9%)]">
                            {titleCase(formatRelative(new Date(log.occurredAt!.seconds * 1000), new Date()))}
                        </time>
                    </div>
                    <Link
                        href={`/staff/${log.staff!.id}`}
                        target="_blank"
                        className="flex gap-1 text-sm"
                    >
                        by
                        <span className="font-semibold">
                            {log.staff!.name}
                        </span>
                    </Link>
                </div>
            </CardHeader>
            <CardContent>
                <p className="~text-xs/sm">{log.description}</p>
            </CardContent>
        </Card>
    );
}