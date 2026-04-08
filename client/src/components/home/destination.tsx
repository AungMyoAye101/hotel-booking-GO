"use client"
import { Button, Card, CardBody, Image, } from '@heroui/react';
import { ArrowRight } from 'lucide-react';
import Link from 'next/link';


const DESTINATION_DATA = [
    {
        url: '/ygn.webp',
        name: 'yangon'
    },
    {
        url: '/mdy.webp',
        name: 'mandalay'
    },
    {
        url: '/bagan.webp',
        name: 'bagan'
    },
    {
        url: '/bangkok.webp',
        name: 'bangkok'
    },
]

const Destination = () => {


    return (
        <section className=' py-12 space-y-6 '>
            <h1 className='head-1 text-black'>Popular destinations</h1>


            <div

                className='grid grid-cols-1 sm:grid-cols-2 md:grid-cols-4 gap-4 place-items-center '>
                {
                    DESTINATION_DATA.map((item) => (
                        <Card
                            key={item.name}
                            radius='sm'
                            shadow='sm'
                            className='border border-slate-300 p-0 overflow-hidden w-xs sm:w-full'>
                            <CardBody className='p-0'>


                                <Image
                                    src={item.url}
                                    alt='hotel image'
                                    width={300}
                                    height={260}

                                    className='aspect-square object-cover rounded-none w-full'
                                />

                                <div className='flex flex-col gap-1 p-4'>
                                    <h1 className='text-lg font-bold capitalize'>{item.name}</h1>

                                    <Button
                                        as={Link}
                                        href={`/search?destination=${item.name}`}
                                        variant='bordered'
                                        className='text-primary border-primary bg-blue-50'>
                                        Explore stays   <ArrowRight />
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

export default Destination