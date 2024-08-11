import { useContext } from 'react'
import { ImageContext } from '../context/ImageContext'

const Visual = () => {
    const { images } = useContext(ImageContext);

    return (
        <div className="bg-white rounded-lg shadow-lg w-3/4 overflow-hidden">
            <img className="object-cover w-full max-h-32" src={images['Company Logo']} />
        </div>
    )
}

export default Visual
