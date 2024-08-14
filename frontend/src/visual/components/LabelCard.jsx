import getIcon from '../../icons/GetIcon'

const LabelCard = ({data, openModal, name, colors}) => {
    const keys = Object.keys(data)
    const Icon = getIcon(name)

    return (
        <div onClick={() => openModal(name)} className="p-2 group hover:bg-gray-200 rounded-lg cursor-pointer transition-color duration-200 flex justify-between items-center">
            <div className="flex space-x-2 items-center">
                <div
                    className="flex items-center justify-center gradient-background size-10 rounded-full card-shadow"
                    style={{
                        '--from-color': colors[0],
                        '--to-color': colors[1],
                    }}
                >
                    {Icon && <Icon className="text-brand-black size-6" />}
                </div>
                <div>
                    <p className="text-gray-700 font-semibold font-hanken">{(data[keys[0]] || '')}</p>
                    {data[keys[1]] && (
                        <p className="text-sm text-gray-500 font-hanken">{data[keys[1]]}</p>
                    )}
                </div>
            </div>
            <svg className="opacity-0 group-hover:opacity-100 transition-opacity duration-200 text-gray-500 size-6" xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 24 24"><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 7H6a2 2 0 0 0-2 2v9a2 2 0 0 0 2 2h9a2 2 0 0 0 2-2v-1"/><path d="M20.385 6.585a2.1 2.1 0 0 0-2.97-2.97L9 12v3h3zM16 5l3 3"/></g></svg>
        </div>
    )
}

export default LabelCard
