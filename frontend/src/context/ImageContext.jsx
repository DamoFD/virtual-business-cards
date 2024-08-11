import { createContext, useState } from 'react'

export const ImageContext = createContext();

export const ImageProvider = ({ children }) => {
    const [images, setImages] = useState({
        'Company Logo': null,
        'Profile Picture': null,
        'Cover Photo': null,
    });

    const updateImage = (name, data) => {
        setImages(prevImages => ({
            ...prevImages,
            [name]: data,
        }));
    };

    return (
        <ImageContext.Provider value={{ images, updateImage }}>
            {children}
        </ImageContext.Provider>
    );
};
