'use client';
import { useAuth } from '@/stores/auth-store';
import { Avatar, Button, Popover, PopoverContent, PopoverTrigger } from '@heroui/react'
import { LogInIcon } from 'lucide-react'
import Image from 'next/image'
import Link from 'next/link'
import { usePathname } from 'next/navigation';
import { useEffect, useState } from 'react';
import NavUser from './nav-user';

const Header = () => {
    const isAuthenticated = useAuth(s => s.isAuthenticated)
    const user = useAuth(s => s.user);
    const [isScroll, setIsScroll] = useState<boolean>(false);

    const path = usePathname()

    const navBackground = path === '/' ? isScroll ? "bg-blue-950/90" : "bg-blue-800/20" : "bg-blue-950/90";

    useEffect(() => {
        const handleScroll = () => {
            const scroll = window.scrollY;
            setIsScroll(scroll > 100)

        }
        window.addEventListener('scroll', handleScroll);
        return () => window.removeEventListener('scroll', handleScroll);
    }, [])

    return (
        <header className={`${navBackground} w-full z-50 fixed top-0 left-0 right-0 px-4 py-2`}>
            <nav
                className=' flex justify-between items-center  max-w-7xl mx-auto'
            >
                <div>
                    <Link href="/" className='flex items-center gap-4 '>
                        <Image
                            src="/brand-logo.svg"
                            alt="Logo"
                            width={50}
                            height={50}
                            className='bg-white rounded-full'

                        />
                        <h1 className='text-2xl font-bold text-white uppercase'>Booking</h1>
                    </Link>
                </div>

                <div className='flex'>
                    {
                        isAuthenticated

                            ?
                            <Popover>

                                <PopoverTrigger>


                                    <Avatar
                                        src='/user.jpg'
                                        radius='full'
                                        alt='User Avatar'

                                    />
                                </PopoverTrigger>
                                <PopoverContent>
                                    <NavUser id={user?.id!} />
                                </PopoverContent>
                            </Popover>


                            : <Button
                                as={Link}
                                href='/signup'
                                variant='solid'
                                color='primary'

                                radius='sm' >

                                Signup
                                <LogInIcon />

                            </Button>
                    }

                </div>

            </nav >
        </header>
    )
}

export default Header