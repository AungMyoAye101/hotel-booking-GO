"use client";

import {
    Avatar,
    Card,
    CardBody,
    CardHeader,
    Divider,
    Skeleton,
    Tab,
    Tabs,
} from '@heroui/react'
import UserInfoInput from './user-input';
import { useGetUserById } from '@/hooks/use-user';

type props = {
    id: string,
}

const Profile = ({ id }: props) => {

    const { data: user, isLoading } = useGetUserById(id);

    if (isLoading) {
        return <Skeleton className='rounded-lg mx-auto w-full h-screen max-w-5xl ' />
    }

    return (
        <Card className="rounded-2xl shadow-2xl border-2 border-gray-200 w-full ">
            {/* Header */}
            <CardHeader className="flex gap-4 items-center py-4">
                <Avatar
                    src={'/user.jpg'}
                    name={user?.name}
                    size='lg'
                    isBordered
                />
                <div>
                    <h1 className="text-2xl font-semibold">{user?.name}</h1>
                    <p className="text-sm text-gray-500">{user?.email}</p>
                </div>
            </CardHeader>

            <Divider />

            {/* Body */}
            <CardBody >
                <Tabs aria-label='Options' color='secondary'>
                    <Tab key='personal info' title='Personal info'>
                        <UserInfoInput user={user!} readonly={true} />
                    </Tab>

                    <Tab key='update' title='Update info'>
                        <UserInfoInput user={user!} readonly={false} />
                    </Tab>
                </Tabs>
            </CardBody>
        </Card>
    )
}

export default Profile