import Image from 'next/image'
import InputCon from './input-con'


const Hero = () => {

    return (
        <section className='h-screen flex items-center justify-center'>
            <Image
                width={1400}
                height={1024}
                src={'/hero.webp'}
                alt='hotel image'
                className='h-screen w-full  absolute -z-20  object-center object-cover  brightness-60'
            />
            {/* <div className="absolute -z-10 inset-0 bg-black/60 " /> */}
            <div className='space-y-4' >
                <h1 className='text-3xl sm:text-4xl md:text-5xl font-bold text-white'>
                    Find your perfect stay

                </h1>
                <h2 className='text-base sm:text-2xl  text-white ' >
                    Compare prices, check availability, book with confidence.
                </h2>
                <InputCon />
                <div>




                </div>
            </div>

        </section>
    )
}

export default Hero