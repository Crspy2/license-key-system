import { cookies } from "next/headers";
import { Metadata } from "@grpc/grpc-js";

export const SetMetadata = async () => {
    const cookieStore = await cookies()
    const session_token = cookieStore.get("session_token")
    const csrf_token = cookieStore.get("csrf_token")

    const md = new Metadata
    md.set("session_token", session_token?.value || "")
    md.set("csrf_token", csrf_token?.value || "")
    return md
}
