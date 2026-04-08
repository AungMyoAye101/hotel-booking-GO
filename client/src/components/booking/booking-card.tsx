"use client";
import { useGetetBookingByUserId } from '@/hooks/use-booking';
import { useAuth } from '@/stores/auth-store'
import { Button, Card, CardBody, Chip, Skeleton } from '@heroui/react';
import Image from 'next/image';
import FiveStars from '../star';
import { BedDouble, Calendar, CreditCard, Hotel, MapPin, UserRound } from 'lucide-react';
import Link from 'next/link';


const BookingCard = () => {
    const userId = useAuth(s => s.user?._id)
    const { data: bookings, isLoading } = useGetetBookingByUserId(userId as string)


    return (
        <section className='py-10'>
            <h1 className='text-2xl font-semibold my-4'>Your Booking Hostory</h1>
            <div className='space-y-6'>

                {
                    isLoading && Array(6).fill(null).map((_, i) => (
                        <Skeleton key={i} className='w-full h-48' />
                    ))
                }
                {


                    bookings && bookings.map(booking => (
                        <Card
                            key={booking._id}
                            className=' border border-slate-300' >

                            <CardBody>
                                <div className='flex flex-col md:flex-row gap-4'>
                                    <div>
                                        <Image
                                            src={booking.hotel?.photo?.secure_url || '/hotel-card.webp'}
                                            alt={booking?.hotel?.name + "photo "}
                                            width={300}
                                            height={260}
                                            className='object-cover aspect-video rounded-lg w-full md:max-w-sm'
                                        />
                                    </div>

                                    <div className='flex-1 flex flex-col gap-3 justify-center'>



                                        {/* Checkin and out */}
                                        <div className='flex flex-wrap items-center gap-1'>
                                            <Hotel />
                                            <span className='font-semibold mr-4'>
                                                Hotel: {booking.hotel?.name}
                                            </span>
                                            <FiveStars count={booking.hotel?.star} />
                                            <Chip size='sm'
                                                color={booking.status === "CONFIRMED" ? "success" : "default"}
                                                className='text-white'
                                            >{booking.status}</Chip>
                                        </div>
                                        <div className='flex items-center gap-1'>
                                            <MapPin />
                                            <span className='font-semibold text-sm'>
                                                Address: {booking.hotel?.adddress} , {booking.hotel?.city}
                                            </span>

                                        </div>
                                        <div className='flex gap-4 flex-wrap '>

                                            <div className='flex  items-center gap-1'>
                                                <Calendar />
                                                <span className='font-semibold text-sm'>
                                                    Check in :  {new Date(booking.checkIn).toDateString()}
                                                </span>
                                            </div>
                                            <div className='flex items-center gap-1'>
                                                <Calendar />
                                                <span className='font-semibold text-sm'>
                                                    Check Out : {new Date(booking.checkOut).toDateString()}
                                                </span>
                                            </div>
                                        </div>


                                        {/* guest*/}
                                        <div className='flex items-center gap-1'>
                                            <UserRound />
                                            <span className='font-semibold text-sm'>
                                                Guest : {3} Adults
                                            </span>
                                        </div>

                                        {/* room */}
                                        <div className='flex items-center gap-1'>
                                            <BedDouble />
                                            <span className='font-semibold text-sm'>
                                                Room : {booking.room?.name}
                                            </span>
                                        </div>


                                        {/* Payment */}

                                        <div className='flex items-center gap-1'>
                                            <CreditCard />
                                            <span className='font-semibold '>
                                                Total : ${booking.totalPrice}
                                            </span>
                                        </div>
                                    </div>



                                    <Button

                                        as={Link}
                                        href={`/user/booking/${booking._id}`}
                                        color='primary'
                                        variant='solid'
                                        radius='sm'
                                        className='self-end'
                                    >
                                        View Details
                                    </Button>


                                </div>
                            </CardBody>



                        </Card>
                    ))
                }

            </div>
        </section>
    )
}

export default BookingCard