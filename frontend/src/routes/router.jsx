import { createBrowserRouter, Navigate } from 'react-router-dom';

import Editor from '../editor/Editor';
import Register from '../auth/Register';
import Login from '../auth/Login';
import Dashboard from '../dashboard/Dashboard';
import { useAuth } from '../context/AuthContext';

const ProtectedRoute = ({ element }) => {
    const { user, loading } = useAuth();

    if (loading) {
        return <div>Loading...</div>;
    }

    if (!user) {
        return <Navigate to="/register" />;
    }

    return element;
};

const router = createBrowserRouter([
    {
        path: '/',
        element: <Editor />,
    },
    {
        path: '/register',
        element: <Register />,
    },
    {
        path: '/login',
        element: <Login />,
    },
    {
        path: '/dashboard',
        element: (
            <ProtectedRoute element={<Dashboard />} />
        ),
    },
]);

export default router;
