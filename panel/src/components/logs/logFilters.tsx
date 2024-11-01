import { Search, Calendar, User, Box } from 'lucide-react'
import { format } from 'date-fns'
import { StaffObject } from "@/proto/staff_pb"
import { Input } from '@/components/ui/input'
import { Button } from '@/components/ui/button'
import {
    Popover,
    PopoverContent,
    PopoverTrigger,
} from '@/components/ui/popover'
import { Calendar as CalendarComponent } from '@/components/ui/calendar'
import {
    Select,
    SelectContent,
    SelectItem,
    SelectTrigger,
    SelectValue,
} from '@/components/ui/select'

interface LogFiltersProps {
    staff?: StaffObject.AsObject[]
    search: string;
    objectFilter: string;
    staffFilter: string;
    dateFilter: Date | undefined;
    onSearchChange: (value: string) => void;
    onObjectFilterChange: (value: string) => void;
    onStaffFilterChange: (value: string) => void;
    onDateFilterChange: (value: Date | undefined) => void;
    onClearFilters: () => void;
}

export const LogFilters = ({
    staff,
    search,
    objectFilter,
    staffFilter,
    dateFilter,
    onSearchChange,
    onObjectFilterChange,
    onStaffFilterChange,
    onDateFilterChange,
    onClearFilters,
}: LogFiltersProps) => {
    const hasActiveFilters = search || objectFilter !== 'all' || staffFilter !== 'all' || dateFilter

    return (
        <div className="flex flex-col md:flex-row items-center space-y-4 md:space-x-4 md:space-y-0">
            <div className="relative flex-1">
                <Search className="absolute left-3 top-3 h-4 w-4 text-[hsl(215.4_16.3%_46.9%)]" />
                <Input
                    placeholder="Search logs..."
                    value={search}
                    onChange={(e) => onSearchChange(e.target.value)}
                    className="pl-9"
                />
            </div>

            <div className="flex flex-row items-center space-x-4">
                <Select value={objectFilter} onValueChange={onObjectFilterChange}>
                    <SelectTrigger className="w-[180px]">
                        <Box className="mr-2 h-4 w-4" />
                        <SelectValue placeholder="Filter by object" />
                    </SelectTrigger>
                    <SelectContent>
                        <SelectItem value="all">All Objects</SelectItem>
                        <SelectItem value="User">User</SelectItem>
                        <SelectItem value="License Key">License Key</SelectItem>
                        <SelectItem value="Product">Product</SelectItem>
                        <SelectItem value="Staff">Staff</SelectItem>
                    </SelectContent>
                </Select>
                {staff && (
                    <Select value={staffFilter} onValueChange={onStaffFilterChange}>
                        <SelectTrigger className="w-[180px]">
                            <User className="mr-2 h-4 w-4" />
                            <SelectValue placeholder="Filter by staff" />
                        </SelectTrigger>
                        <SelectContent>
                            <SelectItem value="all">All Staff</SelectItem>
                            {(staff || []).map(s => (
                                <SelectItem key={s.id} value={s.name}>{s.name}</SelectItem>
                            ))}
                        </SelectContent>
                    </Select>
                )}
                <Popover>
                    <PopoverTrigger asChild className="hidden lg:inline-flex">
                        <Button variant="outline" className="w-[180px]">
                            <Calendar className="mr-2 h-4 w-4" />
                            {dateFilter ? format(dateFilter, 'PPP') : 'Pick a date'}
                        </Button>
                    </PopoverTrigger>
                    <PopoverContent className="w-auto p-0">
                        <CalendarComponent
                            mode="single"
                            selected={dateFilter}
                            onSelect={onDateFilterChange}
                            initialFocus
                        />
                    </PopoverContent>
                </Popover>
            </div>
            {hasActiveFilters && (
                <Button variant="outline" onClick={onClearFilters}>
                    Clear Filters
                </Button>
            )}
        </div>
    );
}