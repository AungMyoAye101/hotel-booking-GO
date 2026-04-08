"use client";
import { Skeleton } from '@heroui/react'


const HotelCardLoading = () => {
    return (
        <div className='flex flex-col gap-4'>

            {
                Array.from({ length: 3 }).map((_, index) => (

                    <div
                        key={index}
                        className='flex flex-col md:flex-row rounded-lg bg-background shadow-lg border-2 border-slate-200 gap-4 p-4'>
                        <Skeleton className='aspect-video md:w-72 w-full rounded-lg' />
                        <div className='flex-1 flex flex-col gap-4'>
                            <div className='flex justify-between gap-4'>
                                <Skeleton className='h-6 w-3/4 rounded-md' />
                                <Skeleton className='h-6 w-1/4 rounded-md' />
                            </div>
                            <Skeleton className='h-6 w-1/4 rounded-md' />
                            <Skeleton className='h-6 w-1/4 rounded-md' />


                            <div className='flex gap-4 justify-between '>
                                <div className='flex flex-col gap-1'>
                                    <Skeleton className='h-6 w-40 rounded-md' />
                                    <Skeleton className='h-6 w-40 rounded-md' />

                                </div>
                                <Skeleton className='h-8 w-1/4 rounded-md' />
                            </div>

                        </div>
                    </div>
                ))
            }
        </div>
    )
}

export default HotelCardLoading