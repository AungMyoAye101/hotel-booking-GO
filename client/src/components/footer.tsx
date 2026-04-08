'use client'
import { Button, Input } from '@heroui/react'
import { BadgeDollarSign, FacebookIcon, InstagramIcon, XIcon, YoutubeIcon } from 'lucide-react';
import Image from 'next/image'
import Link from 'next/link'

const Footer = () => {
    return (
        <footer className='bg-slate-800 '>
            <section className='max-w-7xl mx-auto text-white py-12 px-4 flex flex-wrap gap-4'>
                <div className='space-y-4 max-w-xs'>
                    <div className='flex items-center gap-2'>


                        <Image
                            src={'/brand-logo.svg'}
                            alt="Logo"
                            width={80}
                            height={80}
                            className='rounded-full bg-white'
                        />
                        <h1 className='text-xl font-semibold'>BOOKING</h1>
                    </div>
                    <p className='text-balance'>
                        Booking helps you find trusted hotels at the best prices.
                        Secure booking, verified stays, and instant confirmation.
                    </p>
                    <p className='flex gap-2'><BadgeDollarSign /> scurce payment.</p>
                    <div className='flex items-center gap-2'>
                        <Button variant='flat' isIconOnly radius='sm' size='sm'>
                            <FacebookIcon className='text-white/90' />
                        </Button>
                        <Button variant='flat' isIconOnly radius='sm' size='sm'>
                            <InstagramIcon className='text-white/90' />
                        </Button>
                        <Button variant='flat' isIconOnly radius='sm' size='sm'>
                            <XIcon className='text-white/90' />
                        </Button>
                        <Button variant='flat' isIconOnly radius='sm' size='sm'>
                            <YoutubeIcon className='text-white/90' />
                        </Button>

                    </div>
                </div>
                {/* Contact section */}
                <div className='flex-1 max-w-2xl '>
                    <div className='flex flex-wrap justify-between gap-4'>
                        <div className='space-y-4'>
                            <h1 className='text-lg font-semibold'>
                                Company
                            </h1>
                            <Link href='#' className='block hover:underline mt-2'>About Us</Link>
                            <Link href='#' className='block hover:underline mt-2'>Careers</Link>
                            <Link href='#' className='block hover:underline mt-2'>Contact</Link>
                        </div>
                        <div className='space-y-4'>
                            <h1 className='text-lg font-semibold'>
                                Explore
                            </h1>
                            <Link href='#' className='block hover:underline mt-2'>Destinations</Link>
                            <Link href='#' className='block hover:underline mt-2'>Hotels</Link>
                            <Link href='#' className='block hover:underline mt-2'>Deals</Link>
                        </div>
                        <div className='space-y-4'>
                            <h1 className='text-lg font-semibold'>
                                Support
                            </h1>
                            <Link href='#' className='block hover:underline mt-2'>Help center</Link>
                            <Link href='#' className='block hover:underline mt-2'>Cancellation policy</Link>
                            <Link href='#' className='block hover:underline mt-2'>Terms & privacy</Link>
                        </div>
                    </div>
                    <div className='mt-8 space-y-4 '>
                        <h1 className='text-xl font-semibold'>Get hotel deals & travel inspiration</h1>
                        <div className='overflow-hidden rounded-lg flex '>
                            <Input type='email' placeholder='Enter your email' radius='none' />
                            <Button variant='solid' color='primary' radius='none'>Subscribe</Button>
                        </div>

                    </div>
                </div>
            </section>
        </footer>
    )
}

export default Footer