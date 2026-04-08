'use client'
import { Card, CardBody, Checkbox, CheckboxGroup, Input, Radio, RadioGroup, Slider } from '@heroui/react'
import { useRouter, useSearchParams } from 'next/navigation'
import { useDebouncedCallback } from 'use-debounce'


// import { SearchPropsType } from './SideBar'



export const priceRange = [{
    name: "minPrice",
    label: "Min price"
}, {
    name: "maxPrice",
    label: "Max price"
}]

export const starCheckBoxes = {
    name: "rating",
    fields: [{ value: 5, label: "5 Stars" }, { value: 4, label: "4 Stars" }, { value: 3, label: "3 Stars" }, { value: 2, label: "2 Stars" }, { value: 1, label: "1 Stars" }]
}
export const typeCheckBoxes = {
    name: "rating",
    fields: [
        { value: "hotel", label: "Hotel" },
        { value: 'motel', label: "Motel" },
        { value: 'guest-house', label: "Guest house" }
    ]
}

export const ratingOrder = [
    {

        label: "Highest to Lowest rating",
        value: "desc",
    }, {

        label: "Lowest to Highest rating",

        value: "asc",


    }

]

export const priceOrder = [{
    label: "Highest to Lowest price",
    value: 'desc',
}, {

    label: "Lowest to Highest price",

    value: 'asc',


}]
export const typeRadio = [
    { value: "hotel", label: "Hotel" },
    { value: 'motel', label: "Motel" },
    { value: 'guest-house', label: "Guest house" }
]





const SideBar = () => {
    const searchParams = useSearchParams()
    const router = useRouter();
    const updateParams = (key: string, value?: string | string[]) => {
        const params = new URLSearchParams(searchParams.toString());

        if (!value || (Array.isArray(value) && value.length === 0)) {
            params.delete(key);
        } else if (Array.isArray(value)) {
            params.delete(key);
            value.forEach(v => params.append(key, v));
        } else {
            params.set(key, value);
        }
        router.replace(`/search?${params.toString()}`, { scroll: false });
    }

    const handleDestinationChange = useDebouncedCallback((value: string) => {
        updateParams("destination", value)
    }, 500)

    return (
        <Card radius='sm' shadow='md' className='border-2 border-slate-200  max-w-xs hidden md:block ' >
            <CardBody  >
                <div className='flex flex-col gap-6 p-4 max-w-sm'>

                    <Input
                        type='text'
                        label='Search by Destination'
                        labelPlacement='outside'
                        placeholder='destination'
                        name='destination'
                        radius='sm'
                        defaultValue={searchParams.get('destination') || ''}
                        onChange={(e) => handleDestinationChange(e.target.value)}
                    />




                    <div>
                        <Slider
                            className='w-60'
                            defaultValue={[
                                Number(searchParams.get("minPrice")) ?? 100,
                                Number(searchParams.get("maxPrice")) ?? 300]}
                            formatOptions={{ style: "currency", currency: "USD" }}
                            label="Price Range"
                            maxValue={1000}
                            minValue={0}
                            onChangeEnd={(value) => {
                                if (Array.isArray(value)) {
                                    const [min, max] = value;
                                    updateParams("minPrice", String(min));
                                    updateParams("maxPrice", String(max));
                                } else {
                                    updateParams("minPrice", String(value));
                                    updateParams("maxPrice", String(value));
                                }
                            }}
                            step={10}
                            color='secondary'

                        />
                    </div>

                    {/* sort by price order */}
                    <RadioGroup
                        defaultValue={searchParams.get('priceOrder') ?? 'asc'}
                        onChange={(e) => updateParams('priceOrder', e.target.value)

                        }
                        label='Sort by price order'
                    >
                        {
                            priceOrder.map(field => (
                                <Radio value={field.value} key={field.value} color='secondary' >
                                    <span className='text-sm ml-1'>
                                        {field.label}
                                    </span>
                                </Radio>
                            ))
                        }
                    </RadioGroup>


                    {/* Sort by stars */}
                    <div>
                        <CheckboxGroup
                            defaultValue={searchParams.getAll('star')}
                            onChange={(values) => updateParams('star', values)}
                            label="Sort by stars">

                            {
                                starCheckBoxes.fields.map((field) => (
                                    <Checkbox
                                        key={field.value}
                                        value={field.value.toString()}
                                        color='secondary'
                                    >
                                        <span className='text-sm ml-1'>
                                            {field.label}
                                        </span>
                                    </Checkbox>
                                ))
                            }
                        </CheckboxGroup>
                    </div>

                    <RadioGroup
                        defaultValue={searchParams.get('type') ?? ""}
                        onChange={(e) => updateParams("type", e.target.value)}
                        label='Property Typex   '
                    >
                        {
                            typeRadio.map(field => (
                                <Radio
                                    key={field.value}
                                    value={field.value}
                                    color='secondary'
                                >
                                    <span className='text-sm ml-1'>
                                        {field.label}
                                    </span>
                                </Radio>
                            ))
                        }
                    </RadioGroup>
                    {/* Sorting by rating */}


                    <RadioGroup
                        defaultValue={searchParams.get('ratingOrder') ?? "asc"}
                        onChange={(e) => updateParams("ratingOrder", e.target.value)}
                        label='Sort by rating order'
                    >
                        {
                            ratingOrder.map(field => (
                                <Radio
                                    key={field.value}
                                    value={field.value}
                                    color='secondary'
                                >
                                    <span className='text-sm ml-1'>
                                        {field.label}
                                    </span>
                                </Radio>
                            ))
                        }
                    </RadioGroup>

                </div>
            </CardBody>
        </Card>
    )
}

export default SideBar


