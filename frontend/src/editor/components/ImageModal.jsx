import { useEffect, useState, useContext } from 'react'
import { ImageContext } from '../../context/ImageContext'

const ImageModal = ({isOpen, closeModal, modalName}) => {
    const [visible, setVisible] = useState(false)
    const { updateImage, images } = useContext(ImageContext)

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

    return (
        <div className="fixed w-full h-full z-10 top-0 left-0 flex justify-center items-center backdrop-blur-sm">
            <div onClick={closeModal} className="absolute top-0 left-0 w-full h-full z-10 bg-black opacity-60 justify-center items-center" />

                <div className={`bg-white z-[11] p-10 rounded-lg shadow-lg transition-transform duration-200 ${visible ? 'scale-100' : 'scale-0'}`}>

                    <div className="flex justify-between items-center text-gray-700">
                        <p className="text-xl font-bold">{'Add ' + modalName}</p>
                        <p className="cursor-pointer" onClick={closeModal}>X</p>
                    </div>

                    <hr className="my-4 border-gray-300" />

                    <div className="border-2 border-dotted border-gray-300 rounded-lg flex flex-col items-center justify-center p-10">
                    {images[modalName] ? (
                        <img
                            src={images[modalName]}
                            alt="uploaded"
                            className="max-w-full max-h-60 rounded-lg"
                        />
                        ) : (
                            <>
                                <label for="upload">
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
