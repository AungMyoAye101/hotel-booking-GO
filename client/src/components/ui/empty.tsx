'use client'

import { Card, CardBody } from '@heroui/react'
import { Info } from 'lucide-react'


const Empty = () => {
    return (
        <section>


            <Card className='max-w-sm mx-auto'>
                <CardBody>
                    <div className='flex flex-col items-center gap-2 justify-center p-4 '>
                        <Info size={60} className='text-danger-400' />
                        <h2 className='text-2xl font-bold  '>No Hotels Found!</h2>
                        <p className='text-center text-slate-600'>There is currently no hotels for this filters. Please check back later or try a different filters options.</p>
                    </div>
                </CardBody>
            </Card>
        </section>

    )
}

export default Empty