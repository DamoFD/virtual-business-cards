const Details = () => {
    return (
        <div className="mt-10">
            <h2 className="text-gray-700 text-2xl font-bold">Add your details</h2>
            <div className="mt-4 text-gray-700">
                <div className="flex flex-col space-y-2">
                    <p>Personal</p>
                    <div className="flex items-center space-x-4">

                        <div className="bg-white rounded-lg shadow-lg flex flex-col items-center justify-center p-4 space-y-2">
                            <svg className="size-8" xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 256 256"><path fill="currentColor" d="M230.92 212c-15.23-26.33-38.7-45.21-66.09-54.16a72 72 0 1 0-73.66 0c-27.39 8.94-50.86 27.82-66.09 54.16a8 8 0 1 0 13.85 8c18.84-32.56 52.14-52 89.07-52s70.23 19.44 89.07 52a8 8 0 1 0 13.85-8M72 96a56 56 0 1 1 56 56a56.06 56.06 0 0 1-56-56"/></svg>
                            <p>Name</p>
                        </div>

                        <div className="bg-white rounded-lg shadow-lg flex flex-col items-center justify-center p-4 space-y-2">
                            <svg className="size-8" xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 256 256"><path fill="currentColor" d="M216 56h-40v-8a24 24 0 0 0-24-24h-48a24 24 0 0 0-24 24v8H40a16 16 0 0 0-16 16v128a16 16 0 0 0 16 16h176a16 16 0 0 0 16-16V72a16 16 0 0 0-16-16M96 48a8 8 0 0 1 8-8h48a8 8 0 0 1 8 8v8H96Zm120 24v41.61A184 184 0 0 1 128 136a184.1 184.1 0 0 1-88-22.38V72Zm0 128H40v-68.36A200.2 200.2 0 0 0 128 152a200.25 200.25 0 0 0 88-20.37zm-112-88a8 8 0 0 1 8-8h32a8 8 0 0 1 0 16h-32a8 8 0 0 1-8-8"/></svg>
                            <p>Job Title</p>
                        </div>

                        <div className="bg-white rounded-lg shadow-lg flex flex-col items-center justify-center p-4 space-y-2">
                            <svg className="size-8" xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 36 36"><path fill="currentColor" d="M17.9 17.3c2.7 0 4.8-2.2 4.8-4.9s-2.2-4.8-4.9-4.8S13 9.8 13 12.4c0 2.7 2.2 4.9 4.9 4.9m-.1-7.7q.15 0 0 0c1.6 0 2.9 1.3 2.9 2.9s-1.3 2.8-2.9 2.8S15 14 15 12.5c0-1.6 1.3-2.9 2.8-2.9" class="clr-i-outline clr-i-outline-path-1"/><path fill="#232c34" d="M32.7 16.7c-1.9-1.7-4.4-2.6-7-2.5h-.8q-.3 1.2-.9 2.1c.6-.1 1.1-.1 1.7-.1c1.9-.1 3.8.5 5.3 1.6V25h2v-8z" class="clr-i-outline clr-i-outline-path-2"/><path fill="#232c34" d="M23.4 7.8c.5-1.2 1.9-1.8 3.2-1.3c1.2.5 1.8 1.9 1.3 3.2c-.4.9-1.3 1.5-2.2 1.5c-.2 0-.5 0-.7-.1c.1.5.1 1 .1 1.4v.6c.2 0 .4.1.6.1c2.5 0 4.5-2 4.5-4.4c0-2.5-2-4.5-4.4-4.5c-1.6 0-3 .8-3.8 2.2c.5.3 1 .7 1.4 1.3" class="clr-i-outline clr-i-outline-path-3"/><path fill="#232c34" d="M12 16.4q-.6-.9-.9-2.1h-.8c-2.6-.1-5.1.8-7 2.4L3 17v8h2v-7.2c1.6-1.1 3.4-1.7 5.3-1.6c.6 0 1.2.1 1.7.2" class="clr-i-outline clr-i-outline-path-4"/><path fill="#232c34" d="M10.3 13.1c.2 0 .4 0 .6-.1v-.6c0-.5 0-1 .1-1.4c-.2.1-.5.1-.7.1c-1.3 0-2.4-1.1-2.4-2.4S9 6.3 10.3 6.3c1 0 1.9.6 2.3 1.5c.4-.5 1-1 1.5-1.4c-1.3-2.1-4-2.8-6.1-1.5s-2.8 4-1.5 6.1c.8 1.3 2.2 2.1 3.8 2.1" class="clr-i-outline clr-i-outline-path-5"/><path fill="#232c34" d="m26.1 22.7l-.2-.3c-2-2.2-4.8-3.5-7.8-3.4c-3-.1-5.9 1.2-7.9 3.4l-.2.3v7.6c0 .9.7 1.7 1.7 1.7h12.8c.9 0 1.7-.8 1.7-1.7v-7.6zm-2 7.3H12v-6.6c1.6-1.6 3.8-2.4 6.1-2.4c2.2-.1 4.4.8 6 2.4z" class="clr-i-outline clr-i-outline-path-6"/><path fill="none" d="M0 0h36v36H0z"/></svg>
                            <p>Department</p>
                        </div>

                        <div className="bg-white rounded-lg shadow-lg flex flex-col items-center justify-center p-4 space-y-2">
                            <svg className="size-8" xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 24 24"><path fill="currentColor" d="m18.36 9l.6 3H5.04l.6-3zM20 4H4v2h16zm0 3H4l-1 5v2h1v6h10v-6h4v6h2v-6h1v-2zM6 18v-4h6v4z"/></svg>
                            <p>Company Name</p>
                        </div>

                        <div className="bg-white rounded-lg shadow-lg flex flex-col items-center justify-center p-4 space-y-2">
                            <svg className="size-8" xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 24 24"><path fill="currentColor" d="M4 2.5a.5.5 0 0 0-.5.5v18a.5.5 0 0 0 .5.5h6a.75.75 0 0 1 0 1.5H4a2 2 0 0 1-2-2V3a2 2 0 0 1 2-2h10.982a2 2 0 0 1 1.414.586l.064.064l.01.008l.31.312a.75.75 0 0 1-1.06 1.06l-.078-.078l-.004-.003l-.093-.093c-.003-.003 0 0 0 0l-.21-.21a.5.5 0 0 0-.353-.146z"/><path fill="#232c34" d="M18 6.25a4.25 4.25 0 1 0 0 8.5a4.25 4.25 0 0 0 0-8.5m-5.75 4.25a5.75 5.75 0 1 1 11.5 0a5.75 5.75 0 0 1-11.5 0"/><path fill="#232c34" d="m21.283 14.866l1.455 8a.75.75 0 0 1-1.002.836l-3.296-1.24a1.25 1.25 0 0 0-.88 0l-3.296 1.24a.75.75 0 0 1-1.002-.836l1.455-8l1.475.268l-1.217 6.698l2.056-.774a2.75 2.75 0 0 1 1.938 0l2.056.774l-1.217-6.698z"/></svg>
                            <p>Accreditations</p>
                        </div>

                        <div className="bg-white rounded-lg shadow-lg flex flex-col items-center justify-center p-4 space-y-2">
                            <svg className="size-8" xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 256 256"><path fill="currentColor" d="M248 120a48.05 48.05 0 0 0-48-48h-39.8c-2.91-.17-53.62-3.74-101.91-44.24A16 16 0 0 0 32 40v160a16 16 0 0 0 26.29 12.25c37.77-31.68 77-40.76 93.71-43.3v31.72a16 16 0 0 0 7.12 13.33l11 7.33A16 16 0 0 0 194.5 212l11.77-44.36A48.07 48.07 0 0 0 248 120M48 199.93V40c42.81 35.91 86.63 45 104 47.24v65.48c-17.35 2.28-61.16 11.35-104 47.21m131 8v.11l-11-7.33V168h21.6ZM200 152h-32V88h32a32 32 0 1 1 0 64"/></svg>
                            <p>Headline</p>
                        </div>

                    </div>
                </div>
            </div>
        </div>
    )
}

export default Details
