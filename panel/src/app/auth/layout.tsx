import {PropsWithChildren} from "react"

const AuthLayout = ({ children }: PropsWithChildren) => {
    return (
        <div className="bg-black min-h-screen w-full">
            {children}
        </div>
    )
}

export default AuthLayout