import { useContext, useState } from 'react';
import { useNavigate, Link } from 'react-router-dom';
import { CardContext } from '../context/CardContext';
import { useAuth } from '../context/AuthContext';
import GetColors from '../colors/GetColors';
import Stepper from '../editor/components/Stepper';
import Input from '../editor/components/Input';
import axiosClient from '../axios-client';
import loginSchema from '../validation/loginSchema';

const Login = () => {
    const [email, setEmail] = useState('')
    const [password, setPassword] = useState('') 
    const [errors, setErrors] = useState({})

    const { card } = useContext(CardContext)
    const colors = GetColors(card.colors['Colors'])
    const { fetchUser } = useAuth();
    const navigate = useNavigate();

    const handleLogin = async () => {
        try {
            // reset errors
            setErrors({});

            // Validate the data
            await loginSchema.validate({ email, password }, { abortEarly: false });

            // If validation passes, prodceed with the API call
            await axiosClient.post('login', {
                'email': email,
                'password': password,
            });

            // set user
            await fetchUser();

            // redirect to Dashboard
            navigate('/dashboard');
        } catch (error) {
            //setError(error.response.data.error);
            if (error.name === 'ValidationError') {
                const validationErrors = {};
                error.inner.forEach(err => {
                    validationErrors[err.path] = err.message;
                });
                setErrors(validationErrors);
            } else if (error.response && error.response.status === 429) {
                setErrors({ api: 'Too many requests. Please try again later.' });
            } else {
                // Handle API errors
                setErrors({ api: error.response?.data?.error || 'An error occured. Please try again later.' });
            }
        }
    }

    return (
        <div className="w-full bg-brand-background flex flex-col items-center">
            <Stepper colors={colors} />
            <div className="mt-10">
                <h1 className="text-3xl font-extrabold text-brand-black mt-10 font-inter z-[2] relative">Login to your account</h1>
                <p className="text-brand-black font-hanken text-lg">Time to bring your card to life! Login to save your changes and share it with others.</p>
                {errors.api && <p className="text-red-500 text-lg font-semibold font-hanken">{errors.api}</p>}

                <div>

                    <Input autoFocus={true} label="Email" value={email} onChange={(e) => setEmail(e.target.value)} color={errors.email ? 'red-500' : null} />
                    {errors.email && <p className="text-red-500 text-lg font-semibold font-hanken">{errors.email}</p>}

                    <Input label="Password" value={password} onChange={(e) => setPassword(e.target.value)} color={errors.password ? 'red-500' : null} />
                    {errors.password && <p className="text-red-500 text-lg font-semibold font-hanken">{errors.password}</p>}

                    <div className="flex items-center justify-between mt-8">
                        <button
                            className="card-depth px-4 py-2 font-hanken"
                            style={{ background: colors[1] }}
                            onClick={handleLogin}
                        >Continue</button>
                        <p className="text-brand-black font-hanken">Need an account? <Link to="/register" className="text-green-600 underline">Sign up</Link></p>
                    </div>
                </div>

                <div className="flex items-center w-full justify-center space-x-4 mt-10">
                    <hr className="w-1/4 border-brand-black" />
                    <p className="text-brand-black font-hanken font-bold">or</p>
                    <hr className="w-1/4 border-brand-black" />
                </div>

                <div className="w-full flex justify-center">
                    <button className="card-depth py-6 font-hanken flex items-center space-x-2 mt-10 w-1/2 justify-center">
                        <svg xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 128 128"><path fill="#fff" d="M44.59 4.21a63.28 63.28 0 0 0 4.33 120.9a67.6 67.6 0 0 0 32.36.35a57.13 57.13 0 0 0 25.9-13.46a57.44 57.44 0 0 0 16-26.26a74.3 74.3 0 0 0 1.61-33.58H65.27v24.69h34.47a29.72 29.72 0 0 1-12.66 19.52a36.2 36.2 0 0 1-13.93 5.5a41.3 41.3 0 0 1-15.1 0A37.2 37.2 0 0 1 44 95.74a39.3 39.3 0 0 1-14.5-19.42a38.3 38.3 0 0 1 0-24.63a39.25 39.25 0 0 1 9.18-14.91A37.17 37.17 0 0 1 76.13 27a34.3 34.3 0 0 1 13.64 8q5.83-5.8 11.64-11.63c2-2.09 4.18-4.08 6.15-6.22A61.2 61.2 0 0 0 87.2 4.59a64 64 0 0 0-42.61-.38"/><path fill="#e33629" d="M44.59 4.21a64 64 0 0 1 42.61.37a61.2 61.2 0 0 1 20.35 12.62c-2 2.14-4.11 4.14-6.15 6.22Q95.58 29.23 89.77 35a34.3 34.3 0 0 0-13.64-8a37.17 37.17 0 0 0-37.46 9.74a39.25 39.25 0 0 0-9.18 14.91L8.76 35.6A63.53 63.53 0 0 1 44.59 4.21"/><path fill="#f8bd00" d="M3.26 51.5a63 63 0 0 1 5.5-15.9l20.73 16.09a38.3 38.3 0 0 0 0 24.63q-10.36 8-20.73 16.08a63.33 63.33 0 0 1-5.5-40.9"/><path fill="#587dbd" d="M65.27 52.15h59.52a74.3 74.3 0 0 1-1.61 33.58a57.44 57.44 0 0 1-16 26.26c-6.69-5.22-13.41-10.4-20.1-15.62a29.72 29.72 0 0 0 12.66-19.54H65.27c-.01-8.22 0-16.45 0-24.68"/><path fill="#319f43" d="M8.75 92.4q10.37-8 20.73-16.08A39.3 39.3 0 0 0 44 95.74a37.2 37.2 0 0 0 14.08 6.08a41.3 41.3 0 0 0 15.1 0a36.2 36.2 0 0 0 13.93-5.5c6.69 5.22 13.41 10.4 20.1 15.62a57.13 57.13 0 0 1-25.9 13.47a67.6 67.6 0 0 1-32.36-.35a63 63 0 0 1-23-11.59A63.7 63.7 0 0 1 8.75 92.4"/></svg>
                        <p>Continue with Google</p>
                    </button>
                </div>
            </div>
        </div>
    );
};

export default Login;
