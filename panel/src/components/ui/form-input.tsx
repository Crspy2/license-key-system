import * as React from "react"

import { cn } from "@/lib/utils"

export interface InputProps
    extends React.InputHTMLAttributes<HTMLInputElement> {}

const FormInput = React.forwardRef<HTMLInputElement, InputProps>(
    ({ className, type, ...props }, ref) => {
        return (
            <input
                type={type}
                className={cn(
                    "w-full rounded-md bg-neutral-900 px-3.5 py-2 text-sm text-white outline-none ring-1 ring-inset ring-zinc-700 hover:ring-zinc-600 focus:bg-transparent focus:ring-[1.5px] focus:ring-blue-400",
                    className
                )}
                ref={ref}
                {...props}
            />
        )
    }
)
FormInput.displayName = "Input"

export { FormInput }
