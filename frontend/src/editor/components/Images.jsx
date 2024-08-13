import { useState, useContext } from 'react'
import ImageModal from './ImageModal'
import { CardContext } from '../../context/CardContext'

const Images = () => {
    const [isOpen, setIsOpen] = useState(false)
    const [modalName, setModalName] = useState('')
    const { card } = useContext(CardContext);

    const cardNames = ['Profile Picture', 'Cover Photo']

    const openModal = (name) => {
        setModalName(name)
        setIsOpen(true)
    }

    const closeModal = () => {
        setIsOpen(false)
    }

    return (
        <div className="mt-10">
            <h2 className="text-brand-black font-inter text-2xl font-bold">Add images</h2>
            <div className="flex items-center space-x-4 mt-4 relative">
                {cardNames.map((name, index) => (
                    <div
                        key={index}
                        className="flex flex-col items-center justify-center card-depth bg-white p-8 text-brand-black cursor-pointer"
                        onClick={() => openModal(name)}
                    >
                        {card.images[name] ? (
                            <>
                                <img src={card.images[name]} alt={name} className="max-w-full max-h-16 rounded-lg" />
                                <p className="font-semibold font-hanken">{name}</p>
                            </>
                        ) : (
                            <>
                                <p className="text-xl">+</p>
                                <p className="font-semibold font-hanken">{name}</p>
                            </>
                        )}
                    </div>
                ))}
            </div>
            {isOpen && <ImageModal modalName={modalName} isOpen={isOpen} closeModal={closeModal} />}
        </div>
    )
}

export default Images
