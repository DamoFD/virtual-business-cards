import { useEffect, useState, useContext } from 'react'
import { CardContext } from '../../context/CardContext'

const ImageModal = ({isOpen, closeModal, modalName}) => {
    const [visible, setVisible] = useState(false)
    const { updateField, card } = useContext(CardContext)

    useEffect(() => {
        if (isOpen) {
            setVisible(true)
        }
    }, [isOpen])

    const handleImageUpload = (event) => {
        const file = event.target.files[0];
        if (file) {
            const reader = new FileReader();
            reader.onload = () => {
                updateField('images', modalName, reader.result);
            };
            reader.readAsDataURL(file);
        }
    }

    const handleImageDelete = () => {
        updateField('images', modalName, null)
    }

    const handleCloseModal = () => {
        setVisible(false)
        setTimeout(() => {
            closeModal()
        }, 200)
    }

    return (
        <div className="fixed w-full h-full z-10 top-0 left-0 flex justify-center items-center backdrop-blur-sm">
            <div onClick={handleCloseModal} className="absolute top-0 left-0 w-full h-full z-10 bg-black opacity-60 justify-center items-center" />

                <div className={`bg-brand-background z-[11] p-10 rounded-lg shadow-lg transition-transform duration-200 ${visible ? 'scale-100' : 'scale-0'}`}>

                    <div className="flex justify-between items-center text-brand-black font-inter">
                        <p className="text-xl font-bold">{card.images[modalName] ? 'Edit ' + modalName : 'Add ' + modalName}</p>
                        <svg className="cursor-pointer size-5" onClick={handleCloseModal} xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 24 24"><g fill="none" fill-rule="evenodd"><path d="M24 0v24H0V0zM12.593 23.258l-.011.002l-.071.035l-.02.004l-.014-.004l-.071-.035q-.016-.005-.024.005l-.004.01l-.017.428l.005.02l.01.013l.104.074l.015.004l.012-.004l.104-.074l.012-.016l.004-.017l-.017-.427q-.004-.016-.017-.018m.265-.113l-.013.002l-.185.093l-.01.01l-.003.011l.018.43l.005.012l.008.007l.201.093q.019.005.029-.008l.004-.014l-.034-.614q-.005-.019-.02-.022m-.715.002a.02.02 0 0 0-.027.006l-.006.014l-.034.614q.001.018.017.024l.015-.002l.201-.093l.01-.008l.004-.011l.017-.43l-.003-.012l-.01-.01z"/><path fill="currentColor" d="m12 14.122l5.303 5.303a1.5 1.5 0 0 0 2.122-2.122L14.12 12l5.304-5.303a1.5 1.5 0 1 0-2.122-2.121L12 9.879L6.697 4.576a1.5 1.5 0 1 0-2.122 2.12L9.88 12l-5.304 5.304a1.5 1.5 0 1 0 2.122 2.12z"/></g></svg>
                    </div>

                    <hr className="my-4 border-gray-300" />

                    <div className="flex flex-col items-center justify-center">
                    {card.images[modalName] ? (
                        <div className="p-10">
                            <img
                                src={card.images[modalName]}
                                alt="uploaded"
                                className="max-w-full max-h-60 rounded-lg"
                            />
                            <div className="flex justify-between w-full mt-4">
                                <div onClick={handleImageDelete} className="cursor-pointer card-depth bg-red-400 p-2">
                                    <svg className="text-brand-black size-8" xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 24 24"><path fill="currentColor" d="M6 19a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V7H6zM8 9h8v10H8zm7.5-5l-1-1h-5l-1 1H5v2h14V4z"/></svg>
                                </div>
                                <label for="upload" className="card-depth bg-blue-400 p-2 cursor-pointer">
                                    <svg className="text-brand-black size-8" xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 24 24"><path fill="currentColor" d="m7 17.013l4.413-.015l9.632-9.54c.378-.378.586-.88.586-1.414s-.208-1.036-.586-1.414l-1.586-1.586c-.756-.756-2.075-.752-2.825-.003L7 12.583zM18.045 4.458l1.589 1.583l-1.597 1.582l-1.586-1.585zM9 13.417l6.03-5.973l1.586 1.586l-6.029 5.971L9 15.006z"/><path fill="currentColor" d="M5 21h14c1.103 0 2-.897 2-2v-8.668l-2 2V19H8.158c-.026 0-.053.01-.079.01c-.033 0-.066-.009-.1-.01H5V5h6.847l2-2H5c-1.103 0-2 .897-2 2v14c0 1.103.897 2 2 2"/></svg>
                                </label>
                                <button onClick={handleCloseModal} className="card-depth bg-green-400 p-2 text-brand-black font-hanken font-bold">Done</button>
                                <input
                                    id="upload"
                                    type="file"
                                    accept="image/*"
                                    onChange={handleImageUpload}
                                    className="hidden"
                                />
                            </div>
                        </div>
                        ) : (
                            <>
                                <label className="cursor-pointer p-10 bg-white card-depth font-hanken" for="upload">
                                    <p className="font-semibold text-brand-black">Drop your image here, or <span className="text-green-600">browse</span></p>
                                    <p className="text-sm text-gray-600">Supports JPG, JPEG, and PNG</p>
                                </label>
                                <input
                                    id="upload"
                                    type="file"
                                    accept="image/*"
                                    onChange={handleImageUpload}
                                    className="hidden"
                                />
                            </>
                        )}
                    </div>

                </div>

        </div>
    )
}

export default ImageModal
