const Stepper = () => {
    return (
        <div className="flex items-center justify-between w-full mx-auto text-gray-800">

            <div className="flex flex-col items-center justify-center">
                <div className="bg-red-400 rounded-full size-10 text-white flex items-center justify-center">
                    <p className="text-lg font-bold">1</p>
                </div>
                <p className="text-sm">Customize your card</p>
            </div>

            <div className="flex flex-col items-center justify-center">
                <div className="border-red-400 border-2 rounded-full size-10 text-red-400 flex items-center justify-center">
                    <p className="text-lg font-bold">2</p>
                </div>
                <p className="text-sm">Create an account</p>
            </div>

            <div className="flex flex-col items-center justify-center">
                <div className="border-red-400 border-2 rounded-full size-10 text-red-400 flex items-center justify-center">
                    <p className="text-lg font-bold">3</p>
                </div>
                <p className="text-sm">Get your card</p>
            </div>

        </div>
    )
}

export default Stepper
