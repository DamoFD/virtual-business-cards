import ImageModal from './ImageModal'
import { useState } from 'react'

const Images = () => {
    const [isOpen, setIsOpen] = useState(false)
    const [modalName, setModalName] = useState('')

    const cardNames = ['Company Logo', 'Profile Picture', 'Cover Photo']

    const openModal = (name) => {
        setModalName(name)
        setIsOpen(true)
    }

    const closeModal = () => {
        setIsOpen(false)
    }

    return (
        <div className="mt-10">
            <div className="flex items-center space-x-4">
                <h2 className="text-gray-700 text-2xl font-bold">Add images</h2>
                <button className="bg-white text-gray-300 px-4 py-2 shadow-lg rounded-lg">Change Layout</button>
            </div>
            <div className="flex items-center space-x-4 mt-4">
                {cardNames.map((name, index) => (
                    <div className="flex flex-col items-center justify-center bg-white rounded-lg shadow-lg p-8 text-gray-700 cursor-pointer" onClick={() => openModal(name)}>
                        <p className="text-xl">+</p>
                        <p className="font-semibold">{name}</p>
                    </div>
                ))}
            </div>
            {isOpen && <ImageModal modalName={modalName} isOpen={isOpen} closeModal={closeModal} />}
        </div>
    )
}

export default Images
