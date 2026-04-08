
import Link from "next/link"

const NotFound = () => {
    return (
        <section className='h-screen px-4 py-12 flex justify-center items-center '>
            <div className="flex flex-col items-center gap-4">
                <h1 className='text-6xl font-bold text-center'>404</h1>
                <h2 className='text-2xl font-bold text-center'>Page Not found!</h2>
                <p>It happens! Let’s get you back on track.</p>
                <Link
                    href={'/'}
                    className="px-4 py-2 bg-primary rounded-lg text-white text-sm"
                >
                    Back to home page
                </Link>
            </div>
        </section>
    )
}

export default NotFound