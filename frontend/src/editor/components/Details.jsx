import { useState } from 'react'
import InputModal from './InputModal'
import OptionCard from './OptionCard'
import options from '../../../optionData.json'

const Details = () => {
    const [modalName, setModalName] = useState('')
    const [fields, setFields] = useState([])
    const [suggestions, setSuggestions] = useState([])
    const [isOpen, setIsOpen] = useState(false)

    const openModal = (name, fields, suggestions) => {
        setFields(fields)
        setSuggestions(suggestions)
        setModalName(name)
        setIsOpen(true)
    }

    const closeModal = () => {
        setIsOpen(false)
    }

    return (
        <div className="mt-10 mb-10">
            <h2 className="text-brand-black text-2xl font-bold font-inter">Add your details</h2>
            <div className="mt-4 text-brand-black">
                <div className="flex flex-col space-y-2">
                    {Object.keys(options).map((category, index) => (
                        <div key={index} className="pt-4">
                            <p className="font-semibold font-hanken">{category.charAt(0).toUpperCase() + category.slice(1)}</p>
                            <div className="flex items-center mt-2 flex-wrap gap-4">
                                {options[category].map((item, idx) => (
                                    <OptionCard openModal={openModal} key={idx} name={item.name} fields={item.fields} suggestions={item.suggestions} icon={item.icon} />
                                ))}
                            </div>
                        </div>
                    ))}
                </div>
            </div>
            {isOpen && <InputModal closeModal={closeModal} modalName={modalName} isOpen={isOpen} fields={fields} suggestions={suggestions} />}
        </div>
    )
}

export default Details
