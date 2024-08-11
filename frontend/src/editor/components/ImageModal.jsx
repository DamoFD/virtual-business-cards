import { useEffect, useState } from 'react'

const ImageModal = ({isOpen, closeModal, modalName}) => {
    const [visible, setVisible] = useState(false)

    useEffect(() => {
        if (isOpen) {
            setVisible(true)
        }
    }, [isOpen])

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
                        <p className="font-semibold text-gray-700">Drop your image here, or <span className="text-red-500">browse</span></p>
                        <p className="text-sm text-gray-500">Supports JPG, JPEG, and PNG</p>
                    </div>

                </div>

        </div>
    )
}

export default ImageModal
