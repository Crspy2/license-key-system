"use client"

import { useTheme } from "next-themes"
import { Toaster as Sonner } from "sonner"

type ToasterProps = React.ComponentProps<typeof Sonner>

const Toaster = ({ ...props }: ToasterProps) => {
  const { theme = "system" } = useTheme()

  return (
    <Sonner
      theme={theme as ToasterProps["theme"]}
      className="toaster group"
      toastOptions={{
          unstyled: true,
          className: "mx-auto  w-fit items-center justify-center",
          classNames: {
              error: 'flex items-center gap-1.5 py-2 px-4 rounded-xl text-sm font-medium bg-red-950 text-red-500 border border-red-900',
              success: 'flex py-2 gap-1.5 px-4 rounded-xl text-sm font-medium bg-green-950 text-green-500 border border-green-900',
              warning: 'flex py-2 gap-1.5 px-4 rounded-xl text-sm font-medium bg-yellow-950 text-yellow-500 border border-yellow-900',
              info: 'flex py-2 gap-1.5 px-4 rounded-xl text-sm font-medium bg-sky-950 text-sky-500 border border-sky-900',
          },
      }}
      {...props}
    />
  )
}

export { Toaster }
