import Editor from './editor/Editor'
import Visual from './visual/Visual'
import { CardProvider } from './context/CardContext';

function App() {

  return (
    <CardProvider>
        <div className="h-screen w-full flex bg-brand-background relative">
            <div className="w-1/3 bg-gradient-to-br from-wc-green to-wc-blue flex items-center justify-center overflow-y-auto inner-shadow">
                <Visual />
            </div>
            <div className="w-2/3 h-full p-14 overflow-y-auto relative overflow-x-hidden">
                <Editor />
            </div>
        </div>
    </CardProvider>
  )
}

export default App
