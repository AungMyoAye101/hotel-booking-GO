import AvaliableRooms from "@/components/hotel/avaliable-rooms";
import HotelDetails from "@/components/hotel/hotel-details";
import { serverFetch } from "@/hooks/api";
import { APIResponse } from "@/types";
import { hotelType } from "@/types/hotel-types";
import { Metadata } from "next";

type Params = {
    hotelId: string;
}

const getHotel = async (hotelId: string): Promise<hotelType> => {
    const data = await serverFetch<APIResponse<hotelType>>('/hotels/' + hotelId)

    return data.result;
}

export const generateMetadata = async ({ params }: { params: Promise<Params> }): Promise<Metadata> => {
    const hotelId = (await params).hotelId;
    const hotel = await getHotel(hotelId)

    return {
        title: `${hotel.name || "Hotel details"} `,
        description: `${hotel.description || "hotel description"}`
    }

}

const page = async ({ params }: { params: Promise<Params> }) => {

    const hotelId = (await params).hotelId;

    const hotel = await getHotel(hotelId)


    return (
        <section className="p-4">
            <HotelDetails hotel={hotel} />
            <AvaliableRooms hotel_id={hotelId} />
        </section>
    )
}

export default page;





