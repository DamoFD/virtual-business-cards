import getIcon from '../../icons/GetIcon'

const OptionCard = ({ name, openModal, fields, suggestions }) => {
    const Icon = getIcon(name)

    return (
        <div onClick={() => openModal(name, fields, suggestions)} className="cursor-pointer hover:-translate-y-1 transition-translate duration-200 bg-white rounded-lg shadow-lg flex flex-col items-center justify-center p-4 space-y-2">
            {Icon && <Icon className="size-8" />}
            <p>{ name }</p>
        </div>
    )
}

export default OptionCard
