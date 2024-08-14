import { createBrowserRouter } from 'react-router-dom';

import Editor from '../editor/Editor';
import Register from '../auth/Register';

const router = createBrowserRouter([
    {
        path: '/',
        element: <Editor />,
    },
    {
        path: '/register',
        element: <Register />,
    },
]);

export default router;
