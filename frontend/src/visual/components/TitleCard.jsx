import getIcon from '../../icons/GetIcon'

const TitleCard = ({data, openModal, name}) => {
    const keys = Object.keys(data)
    const Icon = getIcon(name)

    return (
        <div onClick={() => openModal(name)} className="p-2 group hover:bg-gray-100 rounded-lg cursor-pointer transition-color duration-200 flex justify-between items-center">
            <div className="flex space-x-2 items-center">
                <div className="flex items-center justify-center bg-red-400 size-10 rounded-full shadow-lg">
                    {Icon && <Icon className="text-white size-6" />}
                </div>
                <div>
                    {data[keys[1]] && (
                        <p className="text-gray-700 font-semibold">{data[keys[1]]}</p>
                    )}
                </div>
            </div>
            <svg className="opacity-0 group-hover:opacity-100 transition-opacity duration-200 text-gray-500 size-6" xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 24 24"><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 7H6a2 2 0 0 0-2 2v9a2 2 0 0 0 2 2h9a2 2 0 0 0 2-2v-1"/><path d="M20.385 6.585a2.1 2.1 0 0 0-2.97-2.97L9 12v3h3zM16 5l3 3"/></g></svg>
        </div>
    )
}

export default TitleCard
