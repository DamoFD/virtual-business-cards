import Stepper from './components/Stepper'
import Images from './components/Images'
import Details from './components/Details'
import Blotch from '../assets/blotch.svg'
import Blotch2 from '../assets/blotch2.svg'
import Heart from '../assets/heart.svg'
import Mail from '../assets/mail.svg'

const Editor = () => {
    return (
        <div className="w-full h-full bg-gray-100">
            <Stepper />
            <div className="relative inline-block">
                <h1 className="text-3xl font-extrabold text-brand-black mt-10 font-inter z-[2] relative">Create your first card</h1>
                <div className="absolute bg-wc-green w-[245px] h-6 top-[36px] -right-[34px]" />
            </div>
            <p className="text-brand-black font-hanken">Ready to design your card? Pick a field below to get started!</p>
            <Images />
            <Details />
            <img className="absolute top-24 right-0" src={Blotch} />
            <img className="absolute top-96 right-0" src={Blotch2} />
            <img className="absolute top-72 right-[400px]" src={Heart} />
            <img className="absolute top-[600px] right-[300px]" src={Mail} />
        </div>
    )
}

export default Editor
