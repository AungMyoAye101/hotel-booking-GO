"use client"

import { hotelType } from '@/types/hotel-types'
import { Button, Card, CardBody, Chip } from '@heroui/react'
import { ArrowRight, Building, MapPin, Star } from 'lucide-react'
import Image from 'next/image'
import Link from 'next/link'

type Props = {
    hotels: hotelType[]
}
const HotelCard = ({ hotels }: Props) => {
    return (
        <section className='space-y-6'>


            {
                hotels.map(hotel => (

                    <Card key={hotel._id} >
                        <CardBody >

                            <div className='p-2 flex flex-col sm:flex-row gap-4'>

                                <div className="relative w-full sm:w-48 h-48  shrink-0 ">
                                    <Image
                                        src={hotel.photo?.secure_url || '/hotel-card.webp'}
                                        alt="hotel photo"
                                        fill
                                        className=" rounded-lg "
                                    />


                                </div>

                                <div className="flex-1 px-4 py-2 space-y-1 ">
                                    <div className="flex justify-between">
                                        <h3 className="font-semibold truncate text-lg">{hotel.name}</h3>
                                        <Chip color='secondary' radius="sm">
                                            {hotel.rating}
                                        </Chip>

                                    </div>
                                    <div className="flex items-center gap-1">
                                        <Building />      <p>{hotel.type}</p>
                                    </div>
                                    <div className="flex items-center gap-1">
                                        {
                                            Array(hotel.star).fill(null).map((_, i) => (
                                                <Star
                                                    key={i}
                                                    fill=" oklch(79.5% 0.184 86.047)"
                                                    size={20}
                                                    strokeWidth={0}
                                                />
                                            ))
                                        }

                                    </div>

                                    <div className='flex flex-wrap'>
                                        <div className="flex items-center gap-1 text-sm" >
                                            <MapPin />
                                            {hotel.address},
                                        </div>

                                        <span className='font-medium'>{hotel.city}</span>,
                                        <span className='font-semibold'>{hotel.country}</span>,

                                    </div>




                                    <div className="  flex items-center justify-between gap-4 ">

                                        <p>
                                            <span className="text-black font-bold text-xl">${hotel.price}</span>/night

                                        </p>


                                        <Button
                                            as={Link}
                                            href={`/hotel/${hotel._id}`}

                                            radius="sm"
                                            variant="solid"
                                            color="primary"
                                            className="py-2 mb-2 px-6">
                                            Book now <ArrowRight />
                                        </Button>
                                    </div>
                                </div>
                            </div>
                        </CardBody>
                    </Card>
                ))}
        </section>
    )
}

export default HotelCard