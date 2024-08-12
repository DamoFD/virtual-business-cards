const Input = ({label, value, onChange}) => {
    return (
        <div className="relative m-2 max-w-[fit-content] group">
            <input type="text" value={value} onChange={onChange} className="outline-none px-3 py-3 peer" placeholder=" "/>

            <label
                className="
                    absolute
                    left-[9px]
                    top-px
                    text-sm
                    text-gray-500
                    transition-all
                    duration-300
                    px-1
                    transform
                    -translate-y-1/2
                    pointer-events-none
                    peer-placeholder-shown:top-1/2
                    peer-placeholder-shown:text-xl
                    group-focus-within:!top-px
                    group-focus-within:!text-sm
                    group-focus-within:!text-red-400"
            >{label}</label>

            <fieldset
                className="
                    inset-0
                    absolute
                    border
                    border-gray-400
                    rounded
                    pointer-events-none
                    mt-[-9px]
                    invisible
                    peer-placeholder-shown:visible
                    group-focus-within:!border-red-400
                    group-focus-within:border-2
                    group-hover:border-gray-700"
            >
                <legend
                    className="
                        ml-2
                        px-0
                        text-sm
                        transition-all
                        duration-300
                        invisible
                        max-w-[0.01px]
                        group-focus-within:max-w-full
                        group-focus-within:px-1
                        whitespace-nowrap"
                >{label}</legend>
            </fieldset>

            <fieldset
                className="
                    inset-0
                    absolute
                    border
                    border-gray-400
                    rounded
                    pointer-events-none
                    mt-[-9px]
                    visible
                    peer-placeholder-shown:invisible
                    group-focus-within:border-2
                    group-focus-within:!border-red-400
                    group-hover:border-gray-700"
            >
                <legend className="ml-2 text-sm invisible px-1 max-w-full whitespace-nowrap">{label}</legend>
            </fieldset>
        </div>
    )
}

export default Input
