'use client'

import { MapPin, Star } from "lucide-react"

import HotelReview from "./hotel-review"

import { hotelType } from "@/types/hotel-types"
import Image from "next/image"
import { AMENITIES } from "@/utils/amentits"
import ReviewCard from "./review-card"



type Props = {
    hotel: hotelType;
}
const HotelDetails = ({ hotel }: Props) => {

    return (

        <section >
            <div className="h-screen max-w-7xl flex items-end  text-white">

                <Image
                    width={1400}
                    height={1024}
                    src={hotel.photo?.secure_url || '/hero.webp'}
                    alt='hotel image'
                    className='  absolute h-screen inset-0 -z-20  object-center object-cover  brightness-60'
                />
                <div className="space-y-6 mb-12 px-4">
                    <h1 className="text-4xl sm:text-5xl text-white font-semibold">
                        {hotel.name}
                    </h1>
                    <div className="flex items-center">
                        {
                            Array(hotel.star).fill(0).map((_, i) => (
                                <Star fill="yellow" strokeWidth={0} key={i} />
                            ))
                        }
                        <span className="">{hotel.rating} (120 reviews)</span>
                    </div>
                    <p className="flex font-medium ">
                        <MapPin /> {hotel.address}
                    </p>
                    <div className="bg-white rounded-lg p-4 text-black grid grid-cols-2 gap-4">
                        {
                            AMENITIES.filter(amenity => hotel.amenities.includes(amenity.value)).slice(0, 4).map((amenity) => (
                                <div key={amenity.value} className="flex gap-2 items-center">
                                    {amenity.icon} {amenity.label}
                                </div>
                            ))
                        }



                    </div>
                </div>


            </div>
            <div className="bg-white mt-4   py-6 flex flex-col md:flex-row gap-12 justify-between max-w-7xl px-4 md:px-0">
                <div className=" space-y-4">
                    <h1 className="head-1">Overview</h1>
                    <p className="text-balance">
                        {hotel.description}
                    </p>

                    <h1 className="head-1">Amenities highlight</h1>
                    <div className="flex flex-wrap gap-4">
                        {
                            AMENITIES.filter(amenity => hotel.amenities.includes(amenity.value)).slice(0, 4).map((amenity) => (
                                <div key={amenity.value} className="flex gap-2 items-center">
                                    {amenity.icon} {amenity.label}
                                </div>
                            ))
                        }

                    </div>
                </div>
                <HotelReview hotelId={hotel.id} />
            </div>
            {/* <ReviewCard hotelId={hotel.id} /> */}
        </section>
    )
}

export default HotelDetails