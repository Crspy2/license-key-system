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
                <div className="flex h-10 w-10 items-center justify-center rounded-lg bg-[#313ee0]">
                    {log.object === "User" ? (
                        <FaUser className="h-6 w-6" />
                    ) : log.object === "Staff" ? (
                        <FaUserShield className="h-6 w-6" />
                    ) : log.object === "License Key" ? (
                        <FaKey className="h-6 w-6" />
                    ) : log.object === "Product" ? (
                        <AiFillProduct className="h-6 w-6" />
                    ) : (
                        <FaCube className="h-6 w-6" />
                    )}
                </div>
                <div className="flex-1">
                    <div className="flex items-center justify-between">
                        <h3 className="font-semibold">{log.title}</h3>
                        <time className="text-sm text-muted-foreground">
                            {(new Date(log.occurredAt!.seconds * 1000)).toDateString()}<br />
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
                <p className="text-sm">{log.description}</p>
            </CardContent>
        </Card>
    );
}