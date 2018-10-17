import axios from 'axios';

axios.defaults.timeout = 5000;

axios.defaults.baseURL = 'https://gateway.kuipmake.com/api';

export default axios;