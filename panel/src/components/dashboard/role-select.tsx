"use client"

import React from 'react'
import Select, { StylesConfig } from 'react-select'
import makeAnimated from "react-select/animated"
import { useController, Control } from "react-hook-form"

interface RoleSelectProps {
    roles: { label: string; value: number }[]
    control: Control<any>
    name: string
}

export default function RoleSelect({ roles, control, name }: RoleSelectProps) {
    const {
        field: { onChange, value, ref },
    } = useController({
        name,
        control,
    })

    const handleChange = (selectedOption: any) => {
        const value = selectedOption.value
        console.log(value)
        onChange(value)
    }

    const animatedComponents = makeAnimated()

    const customStyles: StylesConfig = {
        control: (provided, state) => ({
            ...provided,
            backgroundColor: '#1a1b1e', // Dark mode background color
            color: '#2c2d31', // Text color
            minHeight: '2.5rem', // Minimum height
            border: '1px solid #2c2d31',
            boxShadow: 'none',
            flexWrap: 'wrap', // Allows wrapping when multiple items are selected
            overflow: 'visible', // Ensures the height grows
            '&:hover': { borderColor: '#3f51b5' },
        }),
        input: (provided) => ({
            ...provided,
            margin: '0px',
            color: "white",
        }),
        indicatorsContainer: (provided) => ({
            ...provided,
            height: '38px'
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
        singleValue: (provided) => ({
            ...provided,
            color: 'white', // Tag background color in dark mode
        }),
        menuPortal: (base) => ({ ...base, zIndex: 9999 }),
    }

    return (
        <Select
            components={animatedComponents}
            options={roles}
            onChange={handleChange}
            value={roles.filter((r) => value === r.value)}
            name={name}
            className="dark text-sm w-36"
            styles={customStyles}
            placeholder="Select role..."
            menuPosition="fixed"
            menuPlacement="auto"
            ref={ref}
        />
    )
}