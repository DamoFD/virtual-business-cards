import { useContext, useState } from 'react'
import { CardContext } from '../../context/CardContext'
import InputModal from '../../editor/components/InputModal'
import TitleCard from './TitleCard'
import LabelCard from './LabelCard'
import GetColors from '../../colors/GetColors'

const PersonalVisual = () => {
    const { card, jsonData } = useContext(CardContext);
    const [modalName, setModalName] = useState('')
    const [isOpen, setIsOpen] = useState(false)
    const [fields, setFields] = useState([])
    const [suggestions, setSuggestions] = useState([])

    const openModal = (name) => {
        setModalName(name)
        setIsOpen(true)

        const data = Object.values(jsonData)
            .flatMap(category => category)
            .find(item => item.name === name)

        if (data) {
            setFields(data.fields || []);
            setSuggestions(data.suggestions || []);
        }
    }

    const closeModal = () => {
        setIsOpen(false)
    }

    const colors = GetColors(card.colors['Colors']);

    return (
        <div className="mx-4 mt-2">
            {(card.Name['First Name'] || card.Name['Last Name']) && (
                <div onClick={() => openModal('Name')} className="p-2 group hover:bg-gray-200 rounded-lg cursor-pointer transition-color duration-200 flex justify-between items-center">
                    <p className="text-brand-black text-3xl font-extrabold inline-block font-inter">{(card.Name['First Name'] || '') + ' ' + (card.Name['Last Name'] || '')}</p>
                    <svg className="opacity-0 group-hover:opacity-100 transition-opacity duration-200 text-gray-500 size-6" xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 24 24"><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 7H6a2 2 0 0 0-2 2v9a2 2 0 0 0 2 2h9a2 2 0 0 0 2-2v-1"/><path d="M20.385 6.585a2.1 2.1 0 0 0-2.97-2.97L9 12v3h3zM16 5l3 3"/></g></svg>
                </div>
            )}
            {card['Job Title']['Job Title'] && (
                <div onClick={() => openModal('Job Title')} className="pl-2 group hover:bg-gray-200 rounded-lg cursor-pointer transition-color duration-200 flex justify-between items-center">
                    <p className="text-xl font-bold text-gray-700 font-hanken">{card['Job Title']['Job Title'] || ''}</p>
                    <svg className="opacity-0 group-hover:opacity-100 transition-opacity duration-200 text-gray-500 size-6" xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 24 24"><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 7H6a2 2 0 0 0-2 2v9a2 2 0 0 0 2 2h9a2 2 0 0 0 2-2v-1"/><path d="M20.385 6.585a2.1 2.1 0 0 0-2.97-2.97L9 12v3h3zM16 5l3 3"/></g></svg>
                </div>
            )}
            {card.Department.Department && (
                <div onClick={() => openModal('Department')} className="pl-2 group hover:bg-gray-200 rounded-lg cursor-pointer transition-color duration-200 flex justify-between items-center">
                    <p className="text-xl font-bold text-gray-700 font-hanken">{card.Department.Department || ''}</p>
                    <svg className="opacity-0 group-hover:opacity-100 transition-opacity duration-200 text-gray-500 size-6" xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 24 24"><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 7H6a2 2 0 0 0-2 2v9a2 2 0 0 0 2 2h9a2 2 0 0 0 2-2v-1"/><path d="M20.385 6.585a2.1 2.1 0 0 0-2.97-2.97L9 12v3h3zM16 5l3 3"/></g></svg>
                </div>
            )}
            {card['Company Name']['Company Name'] && (
                <div onClick={() => openModal('Company Name')} className="pl-2 group hover:bg-gray-200 rounded-lg cursor-pointer transition-color duration-200 flex justify-between items-center">
                    <p className="text-xl font-bold text-gray-700 font-hanken">{card['Company Name']['Company Name'] || ''}</p>
                    <svg className="opacity-0 group-hover:opacity-100 transition-opacity duration-200 text-gray-500 size-6" xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 24 24"><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 7H6a2 2 0 0 0-2 2v9a2 2 0 0 0 2 2h9a2 2 0 0 0 2-2v-1"/><path d="M20.385 6.585a2.1 2.1 0 0 0-2.97-2.97L9 12v3h3zM16 5l3 3"/></g></svg>
                </div>
            )}
            {Object.keys(card).map((key) => {
                const fieldData = card[key];

                if (fieldData && typeof fieldData === 'object') {
                    const fields = Object.keys(fieldData);

                    if (fields.includes('Label') && fieldData[fields[0]] !== null) {
                        return (
                            <LabelCard colors={colors} name={key} key={key} data={fieldData} openModal={openModal} />
                        )
                    } else if (fields.includes('Title') && fieldData[fields[0]] !== null) {
                        return (
                            <TitleCard colors={colors} name={key} key={key} data={fieldData} openModal={openModal} />
                        )
                    }
                }
                return null;
            })}
            {isOpen && <InputModal closeModal={closeModal} modalName={modalName} isOpen={isOpen} fields={fields} suggestions={suggestions} />}
        </div>
    )
}

export default PersonalVisual
