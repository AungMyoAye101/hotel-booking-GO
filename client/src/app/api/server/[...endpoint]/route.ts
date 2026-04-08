import { NextRequest, NextResponse } from "next/server";
import { cookies } from "next/headers";
import { getBaseUrl } from "@/lib";




type RefreshResult = {
    access_token: string;
    setCookie: string[];
};
//=========prevent multiple request from 401 =============

let refreshing: Promise<RefreshResult | null> | null = null;

//============ refresh token logic =========

const refreshToken = async (token: string): Promise<RefreshResult | null> => {
    const BASE_URL = getBaseUrl();
    if (!refreshing) {
        refreshing = (async () => {
            try {
                const res = await fetch(BASE_URL + '/auth/refresh', {
                    method: "POST",
                    headers: {
                        "Content-type": "application/json",
                        "cookie": `refresh_token=${token}`
                    },
                    credentials: "include",
                    cache: "no-store"
                })

                if (!res.ok) return null;
                const data = await res.json();

                return {
                    access_token: data.result.token,
                    setCookie: res.headers.getSetCookie()
                };
            } finally {
                refreshing = null;
            }
        })();


    }
    return refreshing;
}

//============= make api request to server ============
async function forward(req: NextRequest, accessToken?: string, refreshToken?: string) {
    const { pathname, search } = req.nextUrl;
    const BASE_URL = getBaseUrl();
    const endpoint = pathname.replace('/api/server', "") + search;

    const headers: Record<string, string> =
    {
        "Content-Type": "application/json"
    };

    if (accessToken) {
        headers.Authorization = `Bearer ${accessToken}`;
    }
    if (refreshToken) {
        headers.cookie = `refresh_token=${refreshToken}`;
    }

    let body;
    if (!["GET", "HEAD"].includes(req.method)) {

        body = JSON.stringify(await req.json());

    }


    return fetch(BASE_URL + endpoint, {
        method: req.method,
        headers,
        body,
        credentials: "include"
    })

}

//============= route handler ================

const handler = async (req: NextRequest) => {

    const cookieStore = await cookies();
    const refresh_token = cookieStore.get("refresh_token")?.value;
    const access_token = cookieStore.get("access_token")?.value;


    let response = await forward(req, access_token, refresh_token);

    // ===========unauthorized error handle===========

    if (response.status === 401) {
        const refreshed = await refreshToken(refresh_token ?? '');

        if (refreshed) {
            response = await forward(req, refreshed.access_token, refresh_token);

            const data = response.json();


            const res = NextResponse.json(data, { status: 200 });

            res.cookies.set('access_token', refreshed.access_token, {
                httpOnly: true,
                sameSite: 'lax',
                maxAge: 15 * 60 * 1000
            })
            refreshed.setCookie.forEach(cookie => {
                res.headers.append('set-cookie', cookie)
            })
            return res;
        } else {
            return NextResponse.json({
                error: "Session expired",
            }, { status: 401 });

        }
    }

    const data = await response.json();

    if (!response.ok) {
    }

    return NextResponse.json(data, { status: 200 });
}






export { handler as GET, handler as POST, handler as PUT, handler as DELETE }