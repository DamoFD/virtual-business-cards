import Editor from './editor/Editor'

function App() {

  return (
    <div className="min-h-screen w-full flex bg-gray-100">
        <div className="w-2/5 bg-red-200">
            <h1 className="text-red-500">Test</h1>
        </div>
        <div className="w-3/5 h-full p-14">
            <Editor />
        </div>
    </div>
  )
}

export default App
