import { useContext } from "react"
import { CardContext } from "../context/CardContext"
import GetColors from "../colors/GetColors"
import Stepper from './components/Stepper'
import Images from './components/Images'
import Details from './components/Details'
import Navigation from "./components/Navigation"
import Blotch from '../assets/blotch.svg'
import Blotch2 from '../assets/blotch2.svg'
import Heart from '../assets/heart.svg'
import Mail from '../assets/mail.svg'

const Editor = () => {
    const { card } = useContext(CardContext)
    const colors = GetColors(card.colors['Colors'])

    return (
        <div className="w-full bg-brand-background">
            <Stepper colors={colors} />
            <div className="relative inline-block">
                <h1 className="text-3xl font-extrabold text-brand-black mt-10 font-inter z-[2] relative">Create your first card</h1>
                <div
                    className="absolute w-[245px] h-6 top-[36px] -right-[34px]"
                    style={{
                        background: colors[1]
                    }}
                />
            </div>
            <p className="text-brand-black font-hanken">Ready to design your card? Pick a field below to get started!</p>
            <Images />
            <Details />
            <Navigation colors={colors} />
            <img className="absolute top-24 right-0" src={Blotch} />
            <img className="absolute top-96 right-0" src={Blotch2} />
            <img className="absolute top-72 right-[400px]" src={Heart} />
            <img className="absolute top-[600px] right-[300px]" src={Mail} />
        </div>
    )
}

export default Editor
