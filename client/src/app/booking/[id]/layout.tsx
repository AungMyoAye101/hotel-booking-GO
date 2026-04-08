
import { ReactNode } from 'react'

type Props = {
    children: ReactNode
}

const layout = ({ children }: Props) => {
    return <section className="py-24 space-y-6 px-4">
        {children}
    </section>
}

export default layout