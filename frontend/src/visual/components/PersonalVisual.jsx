import { useContext, useState, useRef, useEffect } from 'react'
import { CardContext } from '../../context/CardContext'
import InputModal from '../../editor/components/InputModal'

const PersonalVisual = () => {
    const { card } = useContext(CardContext);
    const [modalName, setModalName] = useState('')
    const [isOpen, setIsOpen] = useState(false)

    const openModal = (name) => {
        setModalName(name)
        setIsOpen(true)
    }

    const closeModal = () => {
        setIsOpen(false)
    }

    return (
        <div className="mx-4 mt-2">
            <div onClick={() => openModal('Name')} className="p-2 group hover:bg-gray-100 rounded-lg cursor-pointer transition-color duration-200 flex justify-between items-center">
                <p className="text-gray-700 text-3xl font-extrabold inline-block">{(card.name['First Name'] || '') + ' ' + (card.name['Last Name'] || '')}</p>
                <svg className="opacity-0 group-hover:opacity-100 transition-opacity duration-200 text-gray-500 size-6" xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 24 24"><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 7H6a2 2 0 0 0-2 2v9a2 2 0 0 0 2 2h9a2 2 0 0 0 2-2v-1"/><path d="M20.385 6.585a2.1 2.1 0 0 0-2.97-2.97L9 12v3h3zM16 5l3 3"/></g></svg>
            </div>
            <div onClick={() => openModal('Job Title')} className="pl-2 group hover:bg-gray-100 rounded-lg cursor-pointer transition-color duration-200 flex justify-between items-center">
                <p className="text-xl font-bold text-gray-700">{card.jobTitle}</p>
                <svg className="opacity-0 group-hover:opacity-100 transition-opacity duration-200 text-gray-500 size-6" xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 24 24"><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 7H6a2 2 0 0 0-2 2v9a2 2 0 0 0 2 2h9a2 2 0 0 0 2-2v-1"/><path d="M20.385 6.585a2.1 2.1 0 0 0-2.97-2.97L9 12v3h3zM16 5l3 3"/></g></svg>
            </div>
            <div className="pl-2 hover:bg-gray-100 rounded-lg cursor-pointer transition-color duration-200">
                <p className="text-xl text-gray-700 font-bold">Department</p>
            </div>
            {isOpen && <InputModal closeModal={closeModal} modalName={modalName} isOpen={isOpen} />}
        </div>
    )
}

export default PersonalVisual
