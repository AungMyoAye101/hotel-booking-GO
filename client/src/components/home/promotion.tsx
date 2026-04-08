'use client'
import { Button, Card, CardBody } from '@heroui/react'
import { ArrowRight } from 'lucide-react'
import Image from 'next/image'
import Link from 'next/link'


const Promotion = () => {
    return (
        <section className='py-20 px-4 flex items-center justify-center'>
            <Card>
                <CardBody>
                    <div className='flex flex-col sm:flex-row gap-6'>
                        <div className='relative aspect-square min-w-xs max-w-lg'>

                            <Image
                                src={'/hotel-bg.webp'}
                                alt='hotel image'
                                fill
                                className=' object-cover rounded-md '

                            />
                        </div>

                        <div className='flex flex-col gap-4 justify-center p-4'>
                            <h1 className='text-3xl font-semibold text-black/90'>Exclusive member deals</h1>
                            <p className='text-balance text-lg font-medium text-black/90 bg-green-100 p-4 rounded-md  '>
                                Save up to <b>20%</b> on selected hotels.<br />
                                Early access to seasonal discounts.

                            </p>
                            <Button
                                as={Link}
                                href='/login'
                                variant='solid'
                                color='primary'
                                radius='sm'
                                className='w-fit'>Get early access <ArrowRight /></Button>
                        </div>
                    </div>
                </CardBody>
            </Card>


        </section >
    )
}

export default Promotion