import * as yup from 'yup';

const registerSchema = yup.object().shape({
    name: yup.string().required('Name is required').min(2, 'Name must be at least 2 characters').max(100, 'Name must be less than 100 characters'),
    email: yup.string().required('Email is required').email('Email is invalid').max(320, 'Email must be less than 320 characters'),
    password: yup.string().required('Password is required').min(6, 'Password must be at least 6 characters').max(100, 'Password must be less than 100 characters'),
    confirmPassword: yup.string().required('Confirm password is required')
        .oneOf([yup.ref('password'), null], 'Passwords must match')
        .min(6, 'Password must be at least 6 characters')
        .max(100, 'Password must be less than 100 characters')
});

export default registerSchema;
