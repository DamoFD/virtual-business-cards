const Images = () => {
    return (
        <div className="mt-10">
            <div className="flex items-center space-x-4">
                <h2 className="text-gray-700 text-2xl font-bold">Add images</h2>
                <button className="bg-white text-gray-300 px-4 py-2 shadow-lg rounded-lg">Change Layout</button>
            </div>
            <div className="flex items-center space-x-4 mt-4">
                <div className="flex flex-col items-center justify-center bg-white rounded-lg shadow-lg p-8 text-gray-700">
                    <p className="text-xl">+</p>
                    <p className="font-semibold">Company Logo</p>
                </div>
                <div className="flex flex-col items-center justify-center bg-white rounded-lg shadow-lg p-8">
                    <p className="text-xl">+</p>
                    <p className="font-semibold">Profile Picture</p>
                </div>
                <div className="flex flex-col items-center justify-center bg-white rounded-lg shadow-lg p-8">
                    <p className="text-xl">+</p>
                    <p className="font-semibold">Cover Photo</p>
                </div>
            </div>
        </div>
    )
}

export default Images
