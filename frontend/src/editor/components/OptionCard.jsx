import getIcon from '../../icons/GetIcon'

const OptionCard = ({ name, openModal, fields, suggestions }) => {
    const Icon = getIcon(name)

    return (
        <div onClick={() => openModal(name, fields, suggestions)} className="cursor-pointer card-depth bg-white flex flex-col items-center justify-center p-4 space-y-2">
            {Icon && <Icon className="size-8" />}
            <p className="font-hanken">{ name }</p>
        </div>
    )
}

export default OptionCard
