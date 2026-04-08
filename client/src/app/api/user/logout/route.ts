import { getBaseUrl } from "@/lib";
import { cookies } from "next/headers";
import { NextRequest, NextResponse } from "next/server";

export async function POST(req: NextRequest) {
    const BASE_URL = getBaseUrl();
    const cookieStore = await cookies()
    const access_token = cookieStore.get("access_token")?.value;
    if (!access_token) {
        throw new Error("User unauthorized.")
    }

    const res = await fetch(BASE_URL + "/auth/logout", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${access_token}`
        },
        credentials: "include"
    })
    const data = await res.json();

    if (!res.ok) {
        return NextResponse.json({
            success: false,
            message: data.message || "Failed to logout"
        },
            { status: data.status })
    }
    const response = NextResponse.json({
        success: true,
        message: data.message || "user logout."
    }, { status: data.status });


    cookieStore.delete("access_token")
    cookieStore.delete("refresh_token")
    return response;
}