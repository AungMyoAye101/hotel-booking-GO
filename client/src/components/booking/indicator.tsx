'use client'
import { BookingStatus } from '@/types'
import { Chip, Progress } from '@heroui/react'

type Prop = {
    status: BookingStatus
}

const Indicator = ({ status }: Prop) => {
    const stage = status === "PENDING" ? 2 : status === "CONFIRMED" ? 3 : 1

    return (
        <div className=' relative h-fit '>


            <Progress aria-label='progress' value={stage === 2 ? 50 : stage === 3 ? 100 : 0} color='primary' radius='sm' />


            <Chip radius='full' color={stage === 1 || stage > 1 ? "primary" : "default"} size={'lg'}
                className='absolute z-10 left-0 top-1/2 transform -translate-y-1/2' >1</Chip>
            <Chip radius='full' color={stage === 2 || stage > 2 ? "primary" : "default"} size={'lg'}
                className='absolute z-10 right-1/2 top-1/2 transform -translate-y-1/2'  >2</Chip>



            <Chip radius='full' color={stage === 3 ? 'primary' : "default"} size={'lg'}
                className='absolute z-10 right-0 top-1/2 transform -translate-y-1/2'
            >3</Chip>
        </div>
    )
}

export default Indicator