import { createContext, useState } from 'react'

export const CardContext = createContext();

export const CardProvider = ({ children }) => {
    const [card, setCard] = useState({
        images: {
            'Company Logo': null,
            'Profile Picture': null,
        },
        name: {
            'First Name': null,
            'Last Name': null,
        },
        jobTitle: '',
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

    const updateJobTitle = (newJobTitle) => {
        setCard(prevCard => ({
            ...prevCard,
            jobTitle: newJobTitle
        }));
    }

    return (
        <CardContext.Provider value={{ card, updateImage, updateName, updateJobTitle }}>
            {children}
        </CardContext.Provider>
    );
};
