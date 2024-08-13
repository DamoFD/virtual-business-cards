import { useEffect, useState, useContext } from 'react'
import { CardContext } from '../../context/CardContext'
import Input from './Input'

const InputModal = ({isOpen, closeModal, modalName, fields, suggestions}) => {
    const [visible, setVisible] = useState(false)
    const { updateField, card } = useContext(CardContext)

    const handleInputChange = (field) => (e) => {
        const value = e.target.value;
        updateField(modalName, field, value);
    }

    const handleDeleteField = () => {
        fields.forEach(field => {
            updateField(modalName, field, null);
        })
    }

    useEffect(() => {
        if (isOpen) {
            setVisible(true)
        }
    }, [isOpen])

    const handleCloseModal = () => {
        setVisible(false)
        setTimeout(() => {
            closeModal()
        }, 200)
    }

    return (
        <div className="fixed w-full h-full z-10 top-0 left-0 flex justify-center items-center backdrop-blur-sm">
            <div onClick={handleCloseModal} className="absolute top-0 left-0 w-full h-full z-10 bg-black opacity-60 justify-center items-center" />
            <div className={`text-brand-black relative bg-white z-[11] w-1/3 p-10 rounded-lg shadow-lg transition-transform duration-200 ${visible ? 'scale-100' : 'scale-0'}`}>
                <svg className="cursor-pointer size-5 absolute top-6 right-6" onClick={handleCloseModal} xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 24 24"><g fill="none" fill-rule="evenodd"><path d="M24 0v24H0V0zM12.593 23.258l-.011.002l-.071.035l-.02.004l-.014-.004l-.071-.035q-.016-.005-.024.005l-.004.01l-.017.428l.005.02l.01.013l.104.074l.015.004l.012-.004l.104-.074l.012-.016l.004-.017l-.017-.427q-.004-.016-.017-.018m.265-.113l-.013.002l-.185.093l-.01.01l-.003.011l.018.43l.005.012l.008.007l.201.093q.019.005.029-.008l.004-.014l-.034-.614q-.005-.019-.02-.022m-.715.002a.02.02 0 0 0-.027.006l-.006.014l-.034.614q.001.018.017.024l.015-.002l.201-.093l.01-.008l.004-.011l.017-.43l-.003-.012l-.01-.01z"/><path fill="currentColor" d="m12 14.122l5.303 5.303a1.5 1.5 0 0 0 2.122-2.122L14.12 12l5.304-5.303a1.5 1.5 0 1 0-2.122-2.121L12 9.879L6.697 4.576a1.5 1.5 0 1 0-2.122 2.12L9.88 12l-5.304 5.304a1.5 1.5 0 1 0 2.122 2.12z"/></g></svg>
                    <p className="font-bold text-xl font-inter">{'Enter Your ' + modalName}</p>
                    <div className="flex flex-col">
                        {fields.map((field, idx) => (
                            <div key={idx}>
                                <Input
                                    label={field}
                                    value={card[modalName][field] || ''}
                                    onChange={handleInputChange(field)}
                                />
                            </div>
                        ))}
                        {(suggestions && suggestions.length > 0) && (
                            <div className="mt-2">
                                <p className="text-sm text-gray-600 font-hasken">Here are some suggestions for your {fields[1]}</p>
                                <div className="flex space-x-4 mt-2">
                                    {suggestions.map((suggestion, i) => (
                                        <p
                                            className="card-depth text-sm text-center text-brand-black font-hasken px-2 py-1 cursor-pointer"
                                            key={i}
                                            onClick={() => handleInputChange(fields[1])({ target: { value: suggestion } })}
                                        >{suggestion}</p>
                                    ))}
                                </div>
                            </div>
                        )}
                    </div>
                    <hr className="border-gray-300 my-4" />
                    <div className="flex justify-between items-center">
                        <div className="card-depth bg-red-400 p-2 cursor-pointer text-brand-black" onClick={handleDeleteField}>
                            <svg className="size-8" xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 24 24"><path fill="currentColor" d="M6 19a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V7H6zM8 9h8v10H8zm7.5-5l-1-1h-5l-1 1H5v2h14V4z"/></svg>
                        </div>
                        <button className="card-depth bg-green-400 p-2 cursor-pointer text-brand-black font-hasken font-bold text-lg" onClick={closeModal}>Submit</button>
                    </div>
            </div>
        </div>
    )
}

export default InputModal
