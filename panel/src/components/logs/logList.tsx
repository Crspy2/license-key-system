"use client"

import { useState } from "react"
import { format } from "date-fns"
import { LogCard } from '@/components/logs/logCard'
import { LogObject } from "@/proto/logs_pb"
import { StaffObject } from "@/proto/staff_pb"
import { LogFilters } from "@/components/logs/logFilters"

interface LogListProps {
    logs: LogObject.AsObject[]
    staff: StaffObject.AsObject[]
}

const isSameDay = (date1: Date, date2: Date) => {
    return (
        date1.getFullYear() === date2.getFullYear() &&
        date1.getMonth() === date2.getMonth() &&
        date1.getDate() === date2.getDate()
    );
}

export const LogList = ({ logs, staff }: LogListProps) => {
    const [search, setSearch] = useState('')
    const [objectFilter, setObjectFilter] = useState<string>('all')
    const [staffFilter, setStaffFilter] = useState<string>('all')
    const [dateFilter, setDateFilter] = useState<Date>()

    const filteredLogs = logs.filter((log) => {
        const matchesSearch = search === '' ||
            log.title.toLowerCase().includes(search.toLowerCase()) ||
            log.description.toLowerCase().includes(search.toLowerCase())

        const matchesObject = objectFilter === 'all' || log.object === objectFilter
        const matchesStaff = staffFilter === 'all' || log.staff!.name === staffFilter
        const matchesDate = !dateFilter ||
            isSameDay(new Date(log.occurredAt!.seconds * 1000) ,dateFilter)

        return matchesSearch && matchesObject && matchesStaff && matchesDate;
    })

    const clearFilters = () => {
        setSearch('')
        setObjectFilter('all')
        setStaffFilter('all')
        setDateFilter(undefined)
    }

    return (
        <div className="container mx-auto py-6 space-y-6">
            <h1 className="text-3xl font-bold">Audit Logs</h1>
            <LogFilters
                staff={staff}
                search={search}
                objectFilter={objectFilter}
                staffFilter={staffFilter}
                dateFilter={dateFilter}
                onSearchChange={setSearch}
                onObjectFilterChange={setObjectFilter}
                onStaffFilterChange={setStaffFilter}
                onDateFilterChange={setDateFilter}
                onClearFilters={clearFilters}
            />
            <div className="space-y-4">
                {filteredLogs.map((log) => (
                    <LogCard key={log.id} log={log}/>
                ))}
            </div>
        </div>
    );
}