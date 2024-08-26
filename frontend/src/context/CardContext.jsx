import { createContext, useState } from 'react'
import jsonData from '../../optionData.json'

export const CardContext = createContext();

export const CardProvider = ({ children }) => {

    const initializeCardState = () => {
        const savedCard = localStorage.getItem('card');
        if (savedCard) {
            return JSON.parse(savedCard);
        }

        const initialState = {
            images: {
                'Cover Photo': null,
                'Profile Picture': null,
            },
            colors: {
                'Colors': 0,
            }
        };

        Object.keys(jsonData).forEach(category => {
            jsonData[category].forEach(item => {
                // If the item has multiple fields, create an object for those fields
                initialState[item.name] = {};
                item.fields.forEach(field => {
                    initialState[item.name][field] = null;
                });
            });
        });
        return initialState;
    }

    const [card, setCard] = useState(initializeCardState);

    const updateField = (name, fieldName, value) => {
        setCard(prevCard => {
            const updatedCard = {
                ...prevCard,
                [name]: prevCard[name]
                    ? {
                          ...prevCard[name],
                          [fieldName]: value,
                      }
                    : value,
            };
            localStorage.setItem('card', JSON.stringify(updatedCard)); // Save the updated card state to localStorage
            return updatedCard;
        });
    };

    return (
        <CardContext.Provider value={{ card, updateField, jsonData }}>
            {children}
        </CardContext.Provider>
    );
};
