import { useEffect, useState, useContext } from 'react'
import { CardContext } from '../../context/CardContext'
import Input from './Input'

const InputModal = ({isOpen, closeModal, modalName, fields, suggestions}) => {
    const [visible, setVisible] = useState(false)
    const { updateField, card } = useContext(CardContext)

    const handleInputChange = (field) => (e) => {
        const value = e.target.value;
        updateField(field, value);
    }

    const handleDeleteField = () => {
        fields.forEach(field => {
            updateField(field, null);
        })
    }

    useEffect(() => {
        if (isOpen) {
            setVisible(true)
        }
    }, [isOpen])

    return (
        <div className="fixed w-full h-full z-10 top-0 left-0 flex justify-center items-center backdrop-blur-sm">
            <div onClick={closeModal} className="absolute top-0 left-0 w-full h-full z-10 bg-black opacity-60 justify-center items-center" />
            <div className={`relative bg-white z-[11] w-1/2 p-10 rounded-lg shadow-lg transition-transform duration-200 ${visible ? 'scale-100' : 'scale-0'}`}>
                <p className="cursor-pointer absolute top-4 right-4" onClick={closeModal}>X</p>
                    <p className="font-bold text-lg text-gray-700">{'Enter Your ' + modalName}</p>
                    <div className="flex flex-col gap-2">
                        {fields.map((field, idx) => (
                            <div key={idx}>
                                <Input
                                    label={field}
                                    value={card[modalName][field]}
                                    onChange={handleInputChange(field)}
                                />
                            </div>
                        ))}
                        {suggestions && (
                            <div>
                                <p>Here are some suggestions for your {fields[1]}</p>
                                {suggestions.map((suggestion, i) => (
                                    <p key={i}>{suggestion}</p>
                                ))}
                            </div>
                        )}
                    </div>
                    <hr className="border-gray-300 my-4" />
                    <div className="flex justify-between items-center">
                        <svg onClick={handleDeleteField} className="cursor-pointer text-red-500 size-8" xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 24 24"><path fill="currentColor" d="M6 19a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V7H6zM8 9h8v10H8zm7.5-5l-1-1h-5l-1 1H5v2h14V4z"/></svg>
                        <button onClick={closeModal}>Submit</button>
                    </div>
            </div>
        </div>
    )
}

export default InputModal
