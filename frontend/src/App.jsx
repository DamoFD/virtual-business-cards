import Editor from './editor/Editor'
import Visual from './visual/Visual'
import { CardProvider } from './context/CardContext';

function App() {

  return (
    <CardProvider>
        <div className="h-screen w-full flex bg-brand-background relative">
            <Visual />
            <div className="w-2/3 h-full p-14 overflow-y-auto relative overflow-x-hidden">
                <Editor />
            </div>
        </div>
    </CardProvider>
  )
}

export default App
