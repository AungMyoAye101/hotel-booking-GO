
import HotelCard from "@/components/hotel/hotel-card"
import HotelCardLoading from "@/components/loading/hotel-loading"
import Meta from "@/components/pagination"
import Empty from "@/components/ui/empty"
import { serverFetch } from "@/hooks/api"
import { APIResponse, MetaType } from "@/types"
import { hotelType } from "@/types/hotel-types"
import { Suspense } from "react"


type HotelQueryProps = {
    destination?: string,
    minPrice?: number,
    maxPrice?: number,
    priceOrder?: 'asc' | 'desc'
    ratingOrder?: 'asc' | 'desc',
    star?: string | string[]
    page: number
}


type HotelWithMeta = {
    hotels: hotelType[],
    meta: MetaType
}

const page = async ({ searchParams }: { searchParams: Promise<HotelQueryProps> }) => {
    const params = await searchParams;


    const urlParams = new URLSearchParams();
    if (params.destination?.trim()) {
        urlParams.set('destination', params.destination);
    } else {
        urlParams.delete('destination');
    }

    urlParams.set('minPrice', String(params.minPrice ?? 100),);
    urlParams.set('maxPrice', String((params.maxPrice ?? 500)));
    urlParams.set('priceOrder', params.priceOrder ?? "asc");
    urlParams.set('ratingOrder', params.ratingOrder ?? "asc");
    urlParams.set('page', params.page ? params.page.toString() : "1");
    urlParams.delete('star');
    if (Array.isArray(params.star)) {
        params.star.forEach(star => {
            urlParams.append("stars", star.toString())
        })
    }

    const data = await serverFetch<APIResponse<HotelWithMeta>>(`/hotel?${urlParams.toString()}`);

    if (data.result.hotels.length === 0) {
        return <Empty />;
    }

    return (
        <Suspense fallback={<HotelCardLoading />}>
            <HotelCard hotels={data.result.hotels} />
            <Meta meta={data.result.meta} />
        </Suspense>
    )
}

export default page