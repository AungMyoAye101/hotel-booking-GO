import { getBaseUrl } from "@/lib";
import { NextRequest, NextResponse } from "next/server";


export async function POST(req: NextRequest) {
    const { pathname } = req.nextUrl;
    const BASE_URL = getBaseUrl();

    let endpoint: string = '/';
    if (pathname === "/api/auth/login") {
        endpoint = BASE_URL + '/auth/login'
    } else if (pathname === "/api/auth/signup") {
        endpoint = BASE_URL + '/auth/register'
    }

    const fields = await req.json();


    const res = await fetch(endpoint, {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(fields),
        credentials: "include"
    })
    const data = await res.json();


    if (!res.ok && !data.success) {

        return NextResponse.json(
            {
                status: data.status,
                message: data.message,
            }
            , { status: data.status })

    }
    const response = NextResponse.json(data.result, { status: data.statusCode })
    response.cookies.set(
        "access_token",
        data.result.token,
        {
            httpOnly: true,
            sameSite: "lax",
            maxAge: 15 * 60 * 1000
        })

    const cookies = res.headers.getSetCookie();
    cookies.forEach(cookie => {
        response.headers.append("set-cookie", cookie)
    })


    return response;
}