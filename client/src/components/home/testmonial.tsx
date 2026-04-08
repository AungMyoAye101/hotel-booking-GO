'use client'

import { Avatar, Card, CardBody, CardFooter, } from '@heroui/react'
import FiveStars from '../star'
const REVIEW_DATA = [
    {
        name: "Jordan Smith",
        rating: 5,
        text: "Absolutely loved every second of my stay here.The staff went above and beyond to make us feel welcome.Already planning my return trip for next summer!"
    },
    {
        name: "Riley Vance",
        rating: 3,
        text: "The location was perfect for sightseeing and local eats.However, the room was much smaller than the photos suggested.Decent for a quick sleep, but don't expect much luxury."
    },
    {
        name: "Casey Morgan",
        rating: 4,
        text: "Very disappointed with the cleanliness of the bathroom.I had to call maintenance twice just to get fresh towels.I would definitely recommend looking at other options nearby."
    }
]

const color: ("primary" | "success" | "warning" | "default" | "secondary" | "danger")[] = ['primary', 'success', 'warning']

const Testmonial = () => {
    return (
        <section className='py-20 px-4 flex flex-col items-center  justify-center'>
            <h1 className='head-1'>Trusted by thousands of travelers</h1>
            <p>
                Real reviews from verified bookings
            </p>
            <div className='flex flex-wrap items-center justify-center gap-6 my-10'>
                {
                    REVIEW_DATA.map((v, i) => (
                        <Card key={i} className='hover:shadow-xl min-w-xs max-w-sm '>

                            <CardBody className='space-y-2' >
                                <FiveStars count={v.rating} />


                                <p className='text-balance'>
                                    {v.text}

                                </p>
                            </CardBody>
                            <CardFooter className='flex gap-4 items-center'>
                                <Avatar src='/user.jpg' alt='user photo' isBordered color={color[i]} size='sm' />
                                <h2 className='text-lg font-semibold '>{v.name}</h2>

                            </CardFooter>
                        </Card>
                    ))
                }
            </div>

        </section>
    )
}

export default Testmonial