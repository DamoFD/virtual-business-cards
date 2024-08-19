import * as yup from 'yup';

const registerSchema = yup.object().shape({
    email: yup.string().required('Email is required').email('Email is invalid').max(320, 'Email must be less than 320 characters'),
    password: yup.string().required('Password is required').min(6, 'Password must be at least 6 characters').max(100, 'Password must be less than 100 characters'),
});

export default registerSchema;
