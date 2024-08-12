import { createContext, useState } from 'react'
import jsonData from '../../optionData.json'

export const CardContext = createContext();

export const CardProvider = ({ children }) => {

    const initializeCardState = () => {
        const initialState = {
            images: {
                'Cover Photo': null,
                'Profile Picture': null,
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
                if (prevCard[name]) {
                    return {
                        ...prevCard,
                        [name]: {
                            ...prevCard[name],
                            [fieldName]: value,
                        }
                    }
                } else {
                    return {
                    ...prevCard,
                    [name]: value,
                };
            }
        })
    }

    return (
        <CardContext.Provider value={{ card, updateField, jsonData }}>
            {children}
        </CardContext.Provider>
    );
};
