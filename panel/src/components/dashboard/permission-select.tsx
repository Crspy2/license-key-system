"use client"

import React from 'react'
import Select, { StylesConfig } from 'react-select'
import makeAnimated from "react-select/animated"
import { useController, Control } from "react-hook-form"

interface PermissionSelectProps {
    permissions: { label: string; value: number }[]
    control: Control<any>
    name: string
}

export default function PermissionSelect({ permissions, control, name }: PermissionSelectProps) {
    const {
        field: { onChange, value, ref },
    } = useController({
        name,
        control,
    })

    const handleChange = (selectedOptions: any) => {
        const values = selectedOptions.map((option: { value: number }) => option.value)
        onChange(values)
    }

    const animatedComponents = makeAnimated()

    const customStyles: StylesConfig = {
        control: (provided, state) => ({
            ...provided,
            backgroundColor: '#1a1b1e', // Dark mode background color
            color: '#2c2d31', // Text color
            border: '1px solid #2c2d31',
            boxShadow: 'none',
            flexWrap: "wrap", // Allows wrapping when multiple items are selected
            overflowY: 'scroll', // Ensures the height grows
            overflowX: 'hidden',
            maxHeight: '2.5rem',
            '&:hover': {
                borderColor: '#3f51b5' ,
                maxHeight: "100%",
            },

        }),
        input: (provided) => ({
            ...provided,
            margin: '0px',
        }),
        indicatorsContainer: (provided) => ({
            ...provided,
            height: '38px',

        }),
        menu: (provided) => ({ ...provided, backgroundColor: '#1a1b1e', position: 'absolute' }),
        option: (provided, state) => ({
            ...provided,
            backgroundColor: state.isSelected ? '#3f51b5' : state.isFocused ? '#2c2d31' : '#1a1b1e',
            color: state.isSelected ? 'white' : '#a0aec0',
            '&:hover': {
                backgroundColor: '#2c2d31'
            }
        }),
        multiValue: (provided) => ({
            ...provided,
            backgroundColor: '#2c2d31', // Tag background color in dark mode
        }),
        multiValueLabel: (provided) => ({
            ...provided,
            color: '#a0aec0', // Tag text color in dark mode
        }),
        multiValueRemove: (provided) => ({
            ...provided,
            color: '#a0aec0',
            '&:hover': {
                backgroundColor: 'rgb(248 113 113)',
                color: 'white'
            },
        }),
        menuPortal: (base) => ({ ...base, zIndex: 9999 }),
    }

    return (
        <Select
            isMulti
            components={animatedComponents}
            options={permissions}
            onChange={handleChange}
            value={permissions.filter((perm) => value?.includes(perm.value))}
            name={name}
            className="dark text-sm max-w-[40rem]"
            styles={customStyles}
            placeholder="Select permissions..."
            menuPosition="fixed"
            menuPlacement="auto"
            ref={ref}
        />
    )
}