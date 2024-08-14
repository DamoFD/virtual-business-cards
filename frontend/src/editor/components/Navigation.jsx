import { Link } from "react-router-dom"

const Navigation = ({colors}) => {
    return (
        <div className="bg-gray-100 bg-opacity-90 rounded-t-lg fixed bottom-0 right-[20%] w-1/3 p-4 flex justify-between items-center z-[4]">
            <p>By continuing, you agree to our <a className="text-green-600 underline" href="#">Privacy Policy</a> and <a className="text-green-600 underline" href="#">Terms of Service</a></p>
            <Link to="/register">
                <button
                    className="card-depth px-6 py-2 font-hanken"
                    style={{background: colors[1]}}
                >Next</button>
            </Link>
        </div>
    )
}

export default Navigation
