
import Profile from "@/components/user/profile";

type Params = {
    id: string
}

const page = async ({ params }: { params: Promise<Params> }) => {
    const id = (await params).id;


    return (
        <div className="min-h-screen px-4 py-20 max-w-4xl mx-auto">
            <Profile id={id} />
        </div>
    )

}

export default page





