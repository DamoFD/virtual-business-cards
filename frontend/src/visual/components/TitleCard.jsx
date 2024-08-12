const TitleCard = ({data, openModal}) => {
    const keys = Object.keys(data)

    return (
        <div onClick={() => openModal(keys[0])} className="p-2 group hover:bg-gray-100 rounded-lg cursor-pointer transition-color duration-200 flex justify-between items-center">
            <div className="flex space-x-2 items-center">
                <div className="flex items-center justify-center bg-red-400 size-10 rounded-full">
                    <svg className="text-white size-6" xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 24 24"><path fill="currentColor" d="m10 17.55l-1.77 1.72a2.47 2.47 0 0 1-3.5-3.5l4.54-4.55a2.46 2.46 0 0 1 3.39-.09l.12.1a1 1 0 0 0 1.4-1.43a3 3 0 0 0-.18-.21a4.46 4.46 0 0 0-6.09.22l-4.6 4.55a4.48 4.48 0 0 0 6.33 6.33L11.37 19A1 1 0 0 0 10 17.55M20.69 3.31a4.49 4.49 0 0 0-6.33 0L12.63 5A1 1 0 0 0 14 6.45l1.73-1.72a2.47 2.47 0 0 1 3.5 3.5l-4.54 4.55a2.46 2.46 0 0 1-3.39.09l-.12-.1a1 1 0 0 0-1.4 1.43a3 3 0 0 0 .23.21a4.47 4.47 0 0 0 6.09-.22l4.55-4.55a4.49 4.49 0 0 0 .04-6.33"/></svg>
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
