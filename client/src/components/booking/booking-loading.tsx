'use client';

import { Skeleton } from "@heroui/react";

const BookingLoading = () => {
    return (
        <section className=" space-y-4">

            <Skeleton className="w-60 h-6 rounded-lg" />
            <Skeleton className="w-60 h-6 rounded-lg " />
            <div className="flex flex-col md:flex-row gap-12 ">
                <div className="space-y-6 w-full max-w-lg">
                    <Skeleton className=" h-96 rounded-lg" />
                    <Skeleton className=" h-96 rounded-lg" />
                    <Skeleton className=" h-40 rounded-lg" />
                </div>
                <Skeleton className="w-full max-w-md h-120 rounded-lg" />
            </div>
        </section>
    )
}

export default BookingLoading