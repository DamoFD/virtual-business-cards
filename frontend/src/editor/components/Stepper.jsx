const Stepper = ({ colors }) => {

    return (
        <div className="flex items-center justify-between w-full mx-auto text-gray-800">

            <div className="flex flex-col items-center justify-center space-y-2">
                <div
                    className="gradient-background rounded-full size-12 text-brand-black flex items-center justify-center card-shadow"
                    style={{
                        '--from-color': colors[0],
                        '--to-color': colors[1],
                    }}
                >
                    <p className="text-lg font-bold font-inter">1</p>
                </div>
                <p className="text-sm font-hanken text-brand-black">Customize your card</p>
            </div>

            <svg width="111" height="64" viewBox="0 0 148 85" fill="none" xmlns="http://www.w3.org/2000/svg">
                <g clip-path="url(#clip0_6_31)">
                    <path d="M125.982 42.124C127.041 41.8191 129.112 41.5359 130.014 40.7297C126.741 38.2471 119.467 37.983 115.47 37.1106C109.396 35.785 103.317 34.544 97.1777 33.54C91.4042 32.5955 88.4926 31.4919 93.1263 37.1248C95.7285 40.2881 100.011 44.56 99.5269 49.0618C97.8909 64.2897 47.5775 49.0629 38.7472 47.4251C28.0457 45.4407 17.4038 43.5454 6.52786 42.9485C4.94343 42.8617 1.79775 43.6129 0.945193 41.7448C0.0114547 39.6995 2.03003 38.5386 3.51499 37.7628C7.96812 35.4359 15.1968 36.7701 19.9632 37.0302C31.263 37.6466 42.3396 38.9209 53.4808 40.7877C56.7032 41.3277 89.2666 48.5511 89.7916 46.4367C90.1511 44.9887 85.2292 40.6664 84.426 39.7485C82.5099 37.5585 79.5507 34.9846 78.7455 32.0862C77.1297 26.2703 82.8423 23.6177 87.636 23.5726C94.9166 23.5044 102.932 25.9334 109.877 27.7927C113.822 28.8488 117.779 29.8471 121.709 30.9552C124.487 31.7385 129.135 34.144 131.935 34.0904C131.517 29.3886 125.144 23.0711 122.272 19.2246C121.071 17.6164 118.281 14.5739 120.454 12.5548C122.206 10.9263 123.804 12.299 125.285 13.4024C129.765 16.7397 133.086 22.709 135.974 27.354C138.819 31.93 143.337 36.4115 145.699 41.13C148.552 46.8267 141.578 46.2927 137.825 47.286C131.853 48.8666 126.062 51.0977 120.261 53.1981C117.683 54.1313 109.257 58.0935 111.41 52.3627C113.017 48.0865 121.688 42.7831 125.982 42.124C127.476 41.6938 124.05 42.4205 125.982 42.124Z" fill="#222221"/>
                </g>
                <defs>
                    <clipPath id="clip0_6_31">
                        <rect width="47.6696" height="139.443" fill="white" transform="translate(134) rotate(73.9385)"/>
                    </clipPath>
                </defs>
            </svg>


            <div className="flex flex-col items-center justify-center space-y-2">
                <div
                    className="gradient-background rounded-full size-12 text-brand-black flex items-center justify-center card-shadow"
                    style={{
                        '--from-color': colors[0],
                        '--to-color': colors[1],
                    }}
                >
                    <p className="text-lg font-bold font-inter">2</p>
                </div>
                <p className="text-sm font-hanken text-brand-black">Create an account</p>
            </div>

            <svg width="102" height="40" viewBox="0 0 136 53" fill="none" xmlns="http://www.w3.org/2000/svg">
                <g clip-path="url(#clip0_6_23)">
                    <path d="M85.8796 28.8191C91.9909 28.5069 98.0157 28.0121 104.019 26.9533C106.882 26.4483 114.633 26.2624 116.194 23.8505C111.106 22.1265 105.542 22.4216 100.47 20.0317C98.2856 19.0025 95.3039 17.5516 95.017 14.8411C94.5906 10.8125 98.238 12.1835 100.829 12.6552C106.4 13.6695 112.197 13.9958 117.837 14.5844C122.154 15.0352 132.342 13.9642 134.374 18.9709C136.414 23.9978 128.684 28.5561 125.413 31.1911C120.113 35.4617 114.859 39.8085 109.522 44.0287C108.051 45.1917 105.275 47.9619 103.087 47.3123C100.032 46.4053 101.274 43.6006 102.499 41.6689C103.29 40.4215 114.195 28.9639 114.22 28.9733C113.275 28.6122 107.82 30.5969 106.922 30.8137C103.254 31.6987 99.4156 32.2842 95.6029 32.7652C85.6346 34.0227 75.2145 35.4704 65.1514 34.3201C60.4491 33.7827 57.0794 31.8534 52.9989 29.8479C49.1241 27.9436 45.4988 30.1139 41.2425 30.8248C32.4499 32.2937 23.2996 31.3762 14.8847 28.5982C9.13897 26.7011 -3.80945 20.5532 1.94747 12.5415C3.4496 10.45 6.13197 9.14718 5.81351 12.1115C5.59728 14.1214 4.07114 14.3483 4.80647 16.7219C5.33569 18.4313 6.88599 19.4078 8.33491 20.2593C15.5332 24.4901 26.109 26.1345 34.2937 25.8329C37.4134 25.718 43.6278 24.9524 46.2527 23.1393C48.5868 21.5268 47.2209 19.0124 47.8562 15.9089C49.5033 7.86289 59.2956 0.960712 67.7589 4.66008C80.0207 10.0196 64.4627 22.4245 58.3141 25.2992C63.9383 30.999 78.8279 29.1793 85.8796 28.8191ZM61.8312 10.003C55.5347 10.5824 53.4395 14.3588 53.9077 20.0104C55.9667 18.9269 71.5734 9.50532 61.8312 10.003C60.7511 10.1024 62.689 9.95919 61.8312 10.003Z" fill="#222221"/>
                </g>
                <defs>
                    <clipPath id="clip0_6_23">
                        <rect width="133.836" height="45.8269" fill="white" transform="translate(136 45.7672) rotate(177.076)"/>
                    </clipPath>
                </defs>
            </svg>


            <div className="flex flex-col items-center justify-center space-y-2">
                <div
                    className="gradient-background rounded-full size-12 text-brand-black flex items-center justify-center card-shadow"
                    style={{
                        '--from-color': colors[0],
                        '--to-color': colors[1],
                    }}
                >
                    <p className="text-lg font-bold font-inter">3</p>
                </div>
                <p className="text-sm font-hanken text-brand-black">Get your card</p>
            </div>

        </div>
    )
}

export default Stepper
