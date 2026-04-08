import { getBaseUrl } from "@/lib";
import { NextRequest, NextResponse } from "next/server";

export async function PUT(req: NextRequest) {

    try {
        const base = getBaseUrl();
        console.log(base)
        const res = await fetch(base + '/auto/cleanup', {
            method: "PUT",
            headers: {
                "Content-type": "application/json"
            }
        })
        const data = await res.json();
        if (!res.ok) {
            return NextResponse.json({
                data
            }, { status: data.status })
        }
        return NextResponse.json({
            data
        }, { status: data.status })

    } catch (err: any) {
        return NextResponse.json({
            message: err?.response?.message || "Somthing went wrong."
        }, { status: 500 })

    }
}