
import LoadingSpinner from "@/components/ui/spinner";



const Loading = () => {
    return (
        <section className="h-screen w-full bg-white/80 flex justify-center items-center">
            <LoadingSpinner />
        </section>
    )
}

export default Loading