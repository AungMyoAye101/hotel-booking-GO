import { BadgeDollarSign, Building2, ShieldCheck } from 'lucide-react'


const About = () => {
    return (
        <section className='py-12 px-4 flex flex-col gap-6 items-center justify-center'>
            <h1 className='head-1'>Why book with us</h1>
            <div className='flex items-center justify-center gap-12 flex-wrap'>
                <div className='flex flex-col gap-2 items-center '>
                    <Building2 size={60} className='text-primary' />
                    <h2 className='text-lg font-semibold'>Verified hotels</h2>
                    <p className=' text-gray-600'>Handpicked and quality-checked stays</p>
                </div>
                <div className='flex flex-col gap-2 items-center'>
                    <BadgeDollarSign size={60} className='text-primary' />
                    <h2 className='text-lg font-semibold'>Best price guarantee</h2>
                    <p className=' text-gray-600 text-balance'>
                        No hidden fees, no surprises
                    </p>
                </div>
                <div className='flex flex-col gap-2 items-center'>
                    <ShieldCheck size={60} className='text-primary' />
                    <h2 className='text-lg font-semibold'>Instant confirmation</h2>
                    <p className=' text-gray-600'>Book now, stay stress-free</p>
                </div>
            </div>
        </section>
    )
}

export default About