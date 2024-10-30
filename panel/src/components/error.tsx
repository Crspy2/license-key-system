export const Error = ({ text }: { text: string }) => {
    return (
        <div className="flex justify-center items-center h-full">
            <div className="py-2 px-4 rounded-xl text-sm font-medium bg-red-950 text-red-500 border border-red-900">
                {text}
            </div>
        </div>
    )
}