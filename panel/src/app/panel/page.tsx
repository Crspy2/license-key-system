import {getCurrentSession} from "@/data/session";
import {redirect} from "next/navigation";
import toast from "react-hot-toast";

const PanelPage = async () => {
    const session = await getCurrentSession()

    if (!session)
        redirect("/auth/login")

    if (!session.success) {
        toast.error(`The request errored with a code ${session.code}. You are being redirected to login`)
        redirect("/auth/login")
    }

    return (
        <div className="bg-red-600">
            <div className="flex flex-col items-center min-h-screen justify-center text-center max-w-5xl mx-auto">
                <h1 className="text-4xl font-semibold text-white">Red Screen</h1>
                <h2 className="text-lg text-zinc-100">
                    You have red screened because the developer of this panel has enabled REDSCREEN_MODE.
                    Using playwright, the Invision tickets are being monitored, and if a ticket exists with the
                    words &quot;red&quot; and &quot;screen&quot; in the same sentence, this screen appears. The only way to remove the
                    screen is to do tickets. Good luck!
                </h2>
                <h3 className="text-lg text-zinc-100">
                    Please for the love of god do not open a ticket 😫!!!
                </h3>
                <p className="flex flex-col text-sm text-zinc-200 py-4">
                    <span>While you wait, please marvel in your glorious session information:</span>
                    <span>{JSON.stringify(session, null, 2)}</span>
                </p>
            </div>
        </div>
    )
}

export default PanelPage