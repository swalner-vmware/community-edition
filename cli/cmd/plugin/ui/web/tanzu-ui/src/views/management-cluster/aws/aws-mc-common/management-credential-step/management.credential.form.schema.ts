// Library imports
import * as yup from 'yup';

export const managementCredentialFormSchema = yup
    .object({
        PROFILE: yup.string(),
        REGION: yup.string().required('Please select an AWS region'),
        SECRET_ACCESS_KEY: yup.string(),
        SESSION_TOKEN: yup.string(),
        ACCESS_KEY_ID: yup.string(),
        EC2_KEY_PAIR: yup.string().required('Please select an EC2 key pair'),
    })
    .required();