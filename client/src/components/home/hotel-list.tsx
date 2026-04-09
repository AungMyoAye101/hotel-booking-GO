'use client';

import { useGetAllHotels } from "@/hooks/use-hotel";
import { Card, CardBody, Button, Chip, Skeleton } from "@heroui/react";
import { ArrowLeft, ArrowRight, MapPin, Star } from "lucide-react";
import Image from "next/image";
import Link from "next/link";
import { useRef } from "react";



export default function HotelCardList() {

    const { data: hotels, isLoading, isError, error } = useGetAllHotels();

    const containerRef = useRef<HTMLDivElement>(null);

    const handleSlide = (isRight: boolean) => {
        const container = containerRef.current;
        if (container) {
            container.scrollBy({
                left: isRight ? 300 : -300,
                behavior: 'smooth'
            })
        }

    }

    if (isLoading) {
        return <section className="flex overflow-hidden flex-nowrap gap-4">
            {
                Array(6).fill(null).map((_, i) => (
                    <Card key={i} className=" min-w-xs w-full max-w-sm ">
                        <CardBody className="p-0">
                            <Skeleton className="w-full h-52 rounded-lg" />
                            <div className="space-y-2 p-4">
                                <div className="flex justify-between">
                                    <Skeleton className="w-32 h-4 rounded-lg" />
                                    <Skeleton className="w-24 h-4 rounded-lg" />
                                </div>
                                <Skeleton className="w-32 h-4 rounded-lg" />
                                <Skeleton className="w-32 h-4 rounded-lg" />
                                <div className="flex justify-between">
                                    <Skeleton className="w-32 h-4 rounded-lg" />
                                    <Skeleton className="w-20 h-4 rounded-lg" />
                                </div>
                            </div>
                        </CardBody>
                    </Card>
                ))
            }
        </section>
    }

    if (isError || !hotels) {
        console.warn(error?.message)
        throw new Error("Failed to load hotels.")
    }
    return (
        <section className="py-12 space-y-6 ">
            <h2 className="head-1 text-black">Featured stays</h2>
            <div

                className="relative "
            >
                <Button isIconOnly
                    variant='solid'
                    radius="full"
                    className="absolute left-0 top-1/2 bottom-1/2 z-10"
                    onPress={() => handleSlide(false)}
                ><ArrowLeft /></Button>
                <div
                    ref={containerRef}
                    className="flex gap-4 overflow-hidden overflow-x-scroll no-scrollbar px-4">

                    {
                        hotels.map(hotel => (
                            <Card
                                className="min-w-xs border-2 border-slate-200"
                                shadow='sm'
                                radius='sm'
                                key={hotel.id}
                            >

                                <CardBody className="p-0">
                                    <div className="relative w-full">

                                        <Image
                                            src={hotel.photo?.secure_url || '/hotel-card.webp'}
                                            alt="hotel photo"
                                            height={160}
                                            width={240}
                                            className="w-full rounded-none aspect-video h-auto"
                                        />
                                    </div>
                                    <div className="px-4 py-2 flex flex-col gap-2">
                                        <div className="flex justify-between">
                                            <h3 className="font-semibold truncate text-lg">{hotel.name}</h3>
                                            <Chip color='secondary' radius="sm">
                                                {hotel.rating}
                                            </Chip>

                                        </div>
                                        <p className="line-clamp-1 flex gap-1 items-center">
                                            <MapPin className="size-5" />
                                            <span className="text-sm">

                                                {hotel.address}
                                            </span>
                                        </p>
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


                                        <div className="   flex items-center justify-between gap-4 ">
                                            <p>
                                                <span className="text-black font-bold text-xl">${hotel.price}</span>/night

                                            </p>
                                            <Button
                                                as={Link}
                                                href={`/hotel/${hotel.id}`}
                                                radius="sm"
                                                variant="solid"
                                                color="primary"
                                                className="py-2 mb-2 px-4">
                                                Book now <ArrowRight />
                                            </Button>
                                        </div>
                                    </div>

                                </CardBody>


                            </Card>
                        ))
                    }

                </div>
                <Button
                    isIconOnly
                    variant='solid'
                    radius="full"
                    className="absolute right-0 top-1/2 bottom-1/2 z-10"
                    onPress={() => handleSlide(true)}
                ><ArrowRight /></Button>
            </div >
        </section >


    );
}
