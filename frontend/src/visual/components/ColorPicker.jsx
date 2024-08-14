import { useContext } from 'react'
import GetColors from '../../colors/GetColors'
import { CardContext } from '../../context/CardContext';

const ColorPicker = () => {
    const { updateField, card } = useContext(CardContext);

    const colors = [];
    for (let i = 0; i < 10; i++) {
        colors.push(GetColors(i))
    }

    const handleColorChange = (value) => {
        updateField('colors', 'Colors', value);
    }

    const selectedClass = 'border border-brand-black';

    return (
        <div className="flex items-center justify-center gap-4 flex-wrap w-3/4 mt-8">
            {colors.map((color, index) => (
                <div key={index}>
                    <div
                        style={{backgroundImage: `linear-gradient(to bottom right, ${color[0]}, ${color[1]})`}}
                        className={`w-10 h-10 rounded-full card-shadow hover:-translate-y-1 transition-transform duration-200 cursor-pointer ${card.colors['Colors'] === index ? selectedClass : ''}`}
                        onClick={() => handleColorChange(index)}
                    />
                </div>
            ))}
        </div>
    )
}

export default ColorPicker
