import { useEffect, useState, useContext } from 'react'
import { CardContext } from '../../context/CardContext'

const ImageModal = ({isOpen, closeModal, modalName}) => {
    const [visible, setVisible] = useState(false)
    const { updateImage, card } = useContext(CardContext)

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
                updateImage(modalName, reader.result);
            };
            reader.readAsDataURL(file);
        }
    }

    const handleImageDelete = () => {
        updateImage(modalName, null)
    }

    return (
        <div className="fixed w-full h-full z-10 top-0 left-0 flex justify-center items-center backdrop-blur-sm">
            <div onClick={closeModal} className="absolute top-0 left-0 w-full h-full z-10 bg-black opacity-60 justify-center items-center" />

                <div className={`bg-white z-[11] p-10 rounded-lg shadow-lg transition-transform duration-200 ${visible ? 'scale-100' : 'scale-0'}`}>

                    <div className="flex justify-between items-center text-gray-700">
                        <p className="text-xl font-bold">{card.images[modalName] ? 'Edit ' + modalName : 'Add ' + modalName}</p>
                        <p className="cursor-pointer" onClick={closeModal}>X</p>
                    </div>

                    <hr className="my-4 border-gray-300" />

                    <div className="border-2 border-dotted border-gray-300 rounded-lg flex flex-col items-center justify-center p-10">
                    {card.images[modalName] ? (
                        <>
                            <img
                                src={card.images[modalName]}
                                alt="uploaded"
                                className="max-w-full max-h-60 rounded-lg"
                            />
                            <div className="flex justify-between w-full mt-4">
                                <svg onClick={handleImageDelete} className="cursor-pointer text-red-500 size-8" xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 24 24"><path fill="currentColor" d="M6 19a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V7H6zM8 9h8v10H8zm7.5-5l-1-1h-5l-1 1H5v2h14V4z"/></svg>
                                <label for="upload">
                                    <svg className="cursor-pointer text-blue-400 size-8" xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 24 24"><path fill="currentColor" d="m7 17.013l4.413-.015l9.632-9.54c.378-.378.586-.88.586-1.414s-.208-1.036-.586-1.414l-1.586-1.586c-.756-.756-2.075-.752-2.825-.003L7 12.583zM18.045 4.458l1.589 1.583l-1.597 1.582l-1.586-1.585zM9 13.417l6.03-5.973l1.586 1.586l-6.029 5.971L9 15.006z"/><path fill="currentColor" d="M5 21h14c1.103 0 2-.897 2-2v-8.668l-2 2V19H8.158c-.026 0-.053.01-.079.01c-.033 0-.066-.009-.1-.01H5V5h6.847l2-2H5c-1.103 0-2 .897-2 2v14c0 1.103.897 2 2 2"/></svg>
                                </label>
                                <input
                                    id="upload"
                                    type="file"
                                    accept="image/*"
                                    onChange={handleImageUpload}
                                    className="hidden"
                                />
                            </div>
                        </>
                        ) : (
                            <>
                                <label className="cursor-pointer" for="upload">
                                    <p className="font-semibold text-gray-700">Drop your image here, or <span className="text-red-500">browse</span></p>
                                    <p className="text-sm text-gray-500">Supports JPG, JPEG, and PNG</p>
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
