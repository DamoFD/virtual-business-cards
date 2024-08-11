import OptionCard from './OptionCard'
import options from '../../../optionData.json'

const Details = () => {
    return (
        <div className="mt-10">
            <h2 className="text-gray-700 text-2xl font-bold">Add your details</h2>
            <div className="mt-4 text-gray-700">
                <div className="flex flex-col space-y-2">
                    {Object.keys(options).map((category, index) => (
                        <div key={index} className="pt-4">
                            <p className="font-semibold">{category.charAt(0).toUpperCase() + category.slice(1)}</p>
                            <div className="flex items-center mt-2 flex-wrap gap-4">
                                {options[category].map((item, idx) => (
                                    <OptionCard key={idx} name={item.name} icon={item.icon} />
                                ))}
                            </div>
                        </div>
                    ))}
                </div>
            </div>
        </div>
    )
}

export default Details
