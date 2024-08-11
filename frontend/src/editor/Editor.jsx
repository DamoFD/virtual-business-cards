import Stepper from './components/Stepper'
import Images from './components/Images'
import Details from './components/Details'

const Editor = () => {
    return (
        <div className="w-full h-full bg-gray-100">
            <Stepper />
            <h1 className="text-3xl font-bold text-gray-800 mt-10">Create your first card</h1>
            <p className="text-gray-700">Ready to design your card? Pick a field below to get started!</p>
            <Images />
            <Details />
        </div>
    )
}

export default Editor
