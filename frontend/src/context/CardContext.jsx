import { createContext, useState } from 'react'

export const CardContext = createContext();

export const CardProvider = ({ children }) => {
    const [card, setCard] = useState({
        name: {
            'First Name': null,
            'Last Name': null,
        },
        images: {
            'Company Logo': null,
            'Profile Picture': null,
        },
    });

    const updateImage = (imageName, imageData) => {
        setCard(prevCard => ({
            ...prevCard,
            images: {
                ...prevCard.images,
                [imageName]: imageData,
            }
        }));
    };

    const updateName = (nameType, nameData) => {
        setCard(prevCard => ({
            ...prevCard,
            name: {
                ...prevCard.name,
                [nameType]: nameData,
            }
        }));
    }

    return (
        <CardContext.Provider value={{ card, updateImage, updateName }}>
            {children}
        </CardContext.Provider>
    );
};
