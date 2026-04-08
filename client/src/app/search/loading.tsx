import LoadingSpinner from "@/components/ui/spinner"



const Loading = () => {
    return (
        <section>
            <div className="flex justify-center items-center min-h-screen">
                <LoadingSpinner />
            </div>
        </section>
    )
}

export default Loading