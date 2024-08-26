import { useNavigate } from 'react-router-dom';
import { useAuth } from '../context/AuthContext';
import axiosClient from '../axios-client';

const Dashboard = () => {
    const { user, setUser } = useAuth();
    const navigate = useNavigate();

    const handleLogout = async () => {
        try {
            await axiosClient.post('logout');
            setUser(null);
            navigate('/login');
        } catch (error) {
            console.error(error);
        }
    }

    return (
        <div>
            <p>Hello, {user.name}</p>
            <button onClick={handleLogout}>Logout</button>
            Dashboard
        </div>
    )
}

export default Dashboard
