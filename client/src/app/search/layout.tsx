import SideBar from '@/components/hotel/filter-sidebar'
import MobileSideBar from '@/components/hotel/mobile-sidebar'
import React from 'react'

const Layout = ({ children }: { children: React.ReactNode }) => {
    return (
        <section className="min-h-screen py-20 px-4">
            <div className="flex flex-col md:flex-row gap-6">
                <div>
                    <MobileSideBar />
                    <SideBar />

                </div>
                <div className="flex-1 ">
                    {children}

                </div>
            </div>
        </section>
    )
}

export default Layout