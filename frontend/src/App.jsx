import Visual from './visual/Visual'
import { CardProvider } from './context/CardContext';
import { AuthProvider } from './context/AuthContext';
import { RouterProvider } from 'react-router-dom';
import router from './routes/router';

function App() {

  return (
    <CardProvider>
        <div className="h-screen w-full flex bg-brand-background relative">
            <Visual />
            <div className="w-2/3 h-full p-14 overflow-y-auto relative overflow-x-hidden">
                <AuthProvider>
                    <RouterProvider router={router} />
                </AuthProvider>
            </div>
        </div>
    </CardProvider>
  )
}

export default App
