import { Star } from 'lucide-react'


const FiveStars = ({ count }: { count: number }) => {

    return (
        <div className='flex items-center gap-0.5'>
            {
                Array.from({ length: 5 }).map((_, i) => (

                    <Star
                        key={i}
                        fill={i + 1 <= count ? "oklch(79.5% 0.184 86.047)" : "white"}
                        size={20}
                        className="text-yellow-500" />

                ))
            }
        </div>
    )
}

export default FiveStars