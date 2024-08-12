import { useEffect, useState, useContext } from 'react'
import { CardContext } from '../../context/CardContext'
import Input from './Input'

const InputModal = ({isOpen, closeModal, modalName}) => {
    const [visible, setVisible] = useState(false)
    const { updateName, updateJobTitle, card } = useContext(CardContext)

    const handleFirstNameChange = (e) => {
        updateName('First Name', e.target.value)
    }

    const handleLastNameChange = (e) => {
        updateName('Last Name', e.target.value)
    }

    const handleDeleteName = () => {
        updateName('First Name', null)
        updateName('Last Name', null)
    }

    const handleJobTitleChange = (e) => {
        updateJobTitle(e.target.value)
    }

    const handleDeleteJobTitle = (e) => {
        updateJobTitle(null)
    }

    useEffect(() => {
        if (isOpen) {
            setVisible(true)
        }
    }, [isOpen])

    return (
        <div className="fixed w-full h-full z-10 top-0 left-0 flex justify-center items-center backdrop-blur-sm">
            <div onClick={closeModal} className="absolute top-0 left-0 w-full h-full z-10 bg-black opacity-60 justify-center items-center" />
            <div className={`relative bg-white z-[11] p-10 rounded-lg shadow-lg transition-transform duration-200 ${visible ? 'scale-100' : 'scale-0'}`}>
                <p className="cursor-pointer absolute top-4 right-4" onClick={closeModal}>X</p>
                {modalName == 'Name' && (
                <>
                    <p className="font-bold text-lg text-gray-700">Enter your name</p>
                    <div className="flex flex-col gap-2">
                        <Input
                            label={'First Name'}
                            value={card.name['First Name'] || ''}
                            onChange={handleFirstNameChange}
                        />
                        <Input
                            label={'Last Name'}
                            value={card.name['Last Name'] || ''}
                            onChange={handleLastNameChange}
                        />
                    </div>
                    <hr className="border-gray-300 my-4" />
                    <div className="flex justify-between items-center">
                        <svg onClick={handleDeleteName} className="cursor-pointer text-red-500 size-8" xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 24 24"><path fill="currentColor" d="M6 19a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V7H6zM8 9h8v10H8zm7.5-5l-1-1h-5l-1 1H5v2h14V4z"/></svg>
                        <button onClick={closeModal}>Submit</button>
                    </div>
                </>
                )}
                {modalName == 'Job Title' && (
                    <>
                        <p className="font-bold text-lg text-gray-700">Enter your job title</p>
                        <Input
                            label={'Job Title'}
                            value={card.jobTitle || ''}
                            onChange={handleJobTitleChange}
                        />
                        <hr className="border-gray-300 my-4" />
                        <div className="flex justify-between items-center">
                            <svg onClick={handleDeleteJobTitle} className="cursor-pointer text-red-500 size-8" xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 24 24"><path fill="currentColor" d="M6 19a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V7H6zM8 9h8v10H8zm7.5-5l-1-1h-5l-1 1H5v2h14V4z"/></svg>
                            <button onClick={closeModal}>Submit</button>
                        </div>
                    </>
                )}
            </div>
        </div>
    )
}

export default InputModal
