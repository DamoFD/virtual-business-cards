import Editor from './editor/Editor'
import Visual from './visual/Visual'
import { CardProvider } from './context/CardContext';

function App() {

  return (
    <CardProvider>
        <div className="h-screen w-full flex bg-gray-100 relative">
            <div className="w-1/3 bg-red-200 flex items-center justify-center">
                <Visual />
            </div>
            <div className="w-2/3 h-full p-14 overflow-y-auto">
                <Editor />
            </div>
        </div>
    </CardProvider>
  )
}

export default App
