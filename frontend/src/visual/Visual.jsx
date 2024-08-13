import { useContext, useState } from 'react'
import { CardContext } from '../context/CardContext'
import ImageModal from '../editor/components/ImageModal'
import PersonalVisual from './components/PersonalVisual'

const Visual = () => {
    const [isOpen, setIsOpen] = useState(false)
    const [modalName, setModalName] = useState('')
    const { card } = useContext(CardContext);

    const openModal = (name) => {
        setModalName(name)
        setIsOpen(true)
    }

    const closeModal = () => {
        setIsOpen(false)
    }

    return (
        <div className="bg-brand-background rounded-lg card-shadow w-3/4 overflow-hidden pb-8">
            {(card.images['Cover Photo'] || card.images['Profile Picture']) && (
                <div className="max-h-56 w-full relative cursor-pointer overflow-hidden group">
                    <img className="object-cover w-full max-h-56 group-hover:scale-105 transition-transform duration-200" src={card.images['Cover Photo'] ? card.images['Cover Photo'] : card.images['Profile Picture']} />
                    <div
                        onClick={() => openModal('Cover Photo')}
                        className="inset-0 bg-black opacity-0 group-hover:opacity-20 absolute top-0 left-0 rounded-t-lg transition-opacity duration-200"
                    >
                        <div className="opacity-0 group-hover:opacity-100 bg-gray-800 flex items-center justify-center p-2 rounded-full text-white float-right mr-2 mt-2">
                            <svg className="size-6" xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 24 24"><path fill="currentColor" d="M20.71 7.04c.39-.39.39-1.04 0-1.41l-2.34-2.34c-.37-.39-1.02-.39-1.41 0l-1.84 1.83l3.75 3.75M3 17.25V21h3.75L17.81 9.93l-3.75-3.75z"/></svg>
                        </div>
                    </div>
                </div>
            )}
            {(card.images['Cover Photo'] && card.images['Profile Picture']) && (
                <div
                    className="size-32 p-2 bg-brand-background rounded-full -mt-16 z-[3] relative ml-4 group cursor-pointer"
                    onClick={() => openModal('Profile Picture')}
                >
                    <div className="z-[4] absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 bg-black opacity-0 group-hover:opacity-50 rounded-full size-10 transition-opacity duration-200" />
                    <svg className="z-[4] absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 text-white size-6 opacity-0 group-hover:opacity-100 transition-opacity duration-200" xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 24 24"><path fill="currentColor" d="M20.71 7.04c.39-.39.39-1.04 0-1.41l-2.34-2.34c-.37-.39-1.02-.39-1.41 0l-1.84 1.83l3.75 3.75M3 17.25V21h3.75L17.81 9.93l-3.75-3.75z"/></svg>
                    <div className="size-full rounded-full overflow-hidden">
                        <img className="object-cover rounded-full size-full group-hover:scale-105 transition-transform duration-200" src={card.images['Profile Picture']} />
                    </div>
                </div>
            )}
            <PersonalVisual />
            {isOpen && <ImageModal modalName={modalName} isOpen={isOpen} closeModal={closeModal} />}
        </div>
    )
}

export default Visual
