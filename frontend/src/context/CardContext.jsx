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

    const updateField = (fieldName, value) => {
        setCard(prevCard => {
            for (let key in prevCard) {
                if (typeof prevCard[key] === 'object' && fieldName in prevCard[key]) {
                    return {
                        ...prevCard,
                        [key]: {
                            ...prevCard[key],
                            [fieldName]: value,
                        }
                    }
                }
            }
            return prevCard
        })
    }

    return (
        <CardContext.Provider value={{ card, updateField, jsonData }}>
            {children}
        </CardContext.Provider>
    );
};
