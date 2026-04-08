"use client"

import { Button, Checkbox, CheckboxGroup, Drawer, DrawerBody, DrawerContent, DrawerFooter, DrawerHeader, Input, Radio, RadioGroup, Slider, useDisclosure } from "@heroui/react";
import { useDebouncedCallback } from "use-debounce";
import { priceOrder, ratingOrder, starCheckBoxes, typeRadio } from "./filter-sidebar";
import { useRouter, useSearchParams } from "next/navigation";

const MobileSideBar = () => {
    const { isOpen, onOpen, onOpenChange } = useDisclosure();

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
        <>
            <Button
                variant='bordered'
                color='primary'
                radius='sm'
                onPress={() => onOpen()}
                className='block md:hidden'
            >
                Filter
            </Button>
            <Drawer isOpen={isOpen} onOpenChange={onOpenChange} placement="left" >
                <DrawerContent>
                    {
                        onclose => (
                            <>
                                <DrawerHeader>
                                    Filter Options
                                </DrawerHeader>
                                <DrawerBody>
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
                                </DrawerBody>
                                <DrawerFooter>
                                    <Button color="danger" variant='bordered' onPress={onclose}>
                                        Close
                                    </Button>

                                </DrawerFooter>
                            </>
                        )
                    }
                </DrawerContent>

            </Drawer>
        </>
    )
}

export default MobileSideBar