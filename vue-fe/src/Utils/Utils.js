const axios = require('axios');

const ENDPOINT_MAP = {
    getBlockchain: 'getBlockchain',
    newBlock: 'newBlock'
};

export default {
    getAxios: () => axios,
    getFullEndpoint: (api) => process.env.VUE_APP_ENDPOINT + ENDPOINT_MAP[api],
    getHttpHeaders: () => {
        return {
            'Content-Type': '*',
            'Accept': '*'
        };
    }
};
