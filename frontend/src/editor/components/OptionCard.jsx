const OptionCard = ({ name, icon }) => {
    return (
        <div className="bg-white rounded-lg shadow-lg flex flex-col items-center justify-center p-4 space-y-2">
            <svg className="size-8" xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox={icon.view}><path fill="currentColor" d={ icon.path }/></svg>
            <p>{ name }</p>
        </div>
    )
}

export default OptionCard
