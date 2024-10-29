export const Error = ({ text }: { text: string }) => {
    return (
        <div className="flex justify-center items-center bg-zinc-900 min-h-screen">
            <div className="py-2 px-4 rounded-xl text-sm font-medium bg-red-950 text-red-500 border border-red-900">
                {text}
            </div>
        </div>
    )
}